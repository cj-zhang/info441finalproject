package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"info441finalproject/server/gateway/handlers"
	"info441finalproject/server/gateway/indexes"
	"info441finalproject/server/gateway/models"
	"info441finalproject/server/gateway/sessions"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//main is the main entry point for the server
func main() {
	tournamentAddr := os.Getenv("TOURNAMENTADDR")
	addr := os.Getenv("ADDR")
	sessionkey := os.Getenv("SESSIONKEY")
	redisaddr := os.Getenv("REDISADDR")
	dsn := os.Getenv("DSN")
	tlscert := os.Getenv("TLSCERT")
	tlskey := os.Getenv("TLSKEY")
	if len(addr) == 0 {
		addr = ":443"
	}
	if len(tlskey) == 0 || len(tlscert) == 0 {
		fmt.Fprintf(os.Stderr, "error: tlskey and tlscert env variables must be set")
		os.Exit(1)
	}

	// create sessionStore through Redis
	redisDb := redis.NewClient(&redis.Options{
		Addr:     redisaddr,
		Password: "",
		DB:       0,
	})
	pong, err := redisDb.Ping().Result()
	fmt.Println(pong, err)
	redisStore := sessions.NewRedisStore(redisDb, 150*time.Second)
	// create userStore through mySQL
	db, err := sql.Open("mysql", dsn)
	if err := db.Ping(); err != nil {
		fmt.Printf("error pinging the db: %v\n", err)
	}

	userStore := &models.MySQLStore{
		Client: db,
	}
	defer db.Close()

	// instantiate handler context
	ctx := &handlers.HandlerContext{
		SigningKey:  sessionkey,
		SessStore:   redisStore,
		UserStore:   userStore,
		SearchTrie:  indexes.NewTrie(),
		SocketStore: new(handlers.SocketStore),
	}

	//Connect to RabbitMQ
	rabbitAddr := os.Getenv("RABBITADDR")
	conn, err := amqp.Dial(rabbitAddr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notify", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for m := range msgs {
			log.Printf("Received a message: %s", m.Body)
			var message map[string]interface{}
			err := json.Unmarshal(m.Body, message)
			if err != nil {
				fmt.Errorf("Error receiving message from queue: %v", err)
			}
			if message["userIDs"] == nil {
				err = ctx.SocketStore.WriteToAllConnections(handlers.TextMessage, m.Body)
				if err != nil {
					fmt.Errorf("Error writing queue message to connections: %v", err)
				}
			} else {
				err = ctx.SocketStore.WriteToConnection(message["userIDs"].([]int64), handlers.TextMessage, m.Body)
				if err != nil {
					fmt.Errorf("Error writing queue message to connections: %v", err)
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	mux := http.NewServeMux()
	director := func(r *http.Request) {
		//Check if authenticated
		sessionState := new(handlers.SessionState)
		sid, err := sessions.GetState(r, ctx.SigningKey, ctx.SessStore, sessionState)
		if sid != sessions.InvalidSessionID && err == nil {
			user, err := json.Marshal(sessionState.User)
			if err == nil {
				r.Header.Del("X-User")
				r.Header.Add("X-User", string(user)) // User is authenticated
			}
		}
		r.Host = tournamentAddr
		r.URL.Host = tournamentAddr
		r.URL.Scheme = "http"
	}
	tournamentProxy := &httputil.ReverseProxy{Director: director}
	mux.HandleFunc("/v1/ws", ctx.WebSocketConnectionHandler)
	mux.Handle("/v1/tournaments", tournamentProxy)
	mux.Handle("/v1/tournaments/players", tournamentProxy)
	mux.Handle("/v1/tournaments/organizers", tournamentProxy)
	mux.Handle("/v1/tournaments/games", tournamentProxy)
	mux.Handle("/v1/tournaments/standings", tournamentProxy)
	mux.HandleFunc("/v1/users", ctx.UsersHandler)
	mux.HandleFunc("/v1/users/", ctx.SpecificUserHandler)
	mux.HandleFunc("/v1/sessions", ctx.SessionsHandler)
	mux.HandleFunc("/v1/sessions/", ctx.SpecificSessionHandler)

	wrappedMux := &handlers.CorsMiddleWare{Handler: mux}

	log.Printf("server is listening at %s...", addr)

	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, wrappedMux))
}

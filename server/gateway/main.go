package main

import (
	"database/sql"
	"fmt"
	"info441finalproject/server/gateway/handlers"
	"info441finalproject/server/gateway/indexes"
	"info441finalproject/server/gateway/models"
	"info441finalproject/server/gateway/sessions"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//main is the main entry point for the server
func main() {
	//summaryaddr := os.Getenv("SUMMARYADDR")
	//msgaddrs := strings.Split(os.Getenv("MESSAGESADDR"), ",")
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
	// if len(summaryaddr) == 0 {
	// 	summaryaddr = "summary:80"
	// }
	// if len(msgaddrs) == 0 {
	// 	msgaddrs = []string{"messaging:80"}
	// }

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

	// handle RabbitMQ connections and instantiate queues
	// conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672/")
	// failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	// ch, err := conn.Channel()
	// failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	"events", // name
	// 	false,    // durable
	// 	false,    // delete when unused
	// 	false,    // exclusive
	// 	false,    // no-wait
	// 	nil,      // arguments
	// )

	// failOnError(err, "Failed to create a queue")

	// events, err := ch.Consume(
	// 	q.Name, // queue
	// 	"",     // consumer
	// 	true,   // auto-ack
	// 	false,  // exclusive
	// 	false,  // no-local
	// 	false,  // no-wait
	// 	nil,    // args
	// )
	// failOnError(err, "Failed to register a consumer")

	//forever := make(chan bool)

	// go func() {
	// 	for event := range events {
	// 		result := &msgevents.Event{}
	// 		json.Unmarshal(event.Body, result)

	// 		// do some processing
	// 		if result.UserIDs != nil {
	// 			ctx.SocketStore.WriteToAllConnections(handlers.TextMessage, event.Body)
	// 		} else {
	// 			ctx.SocketStore.WriteToConnection(result.UserIDs, handlers.TextMessage, event.Body)
	// 		}
	// 	}
	// }()

	// log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-forever

	// change all the string addresses to urls
	// urlMsgAddrs := make([]*url.URL, len(msgaddrs))
	// for _, addr := range msgaddrs {
	// 	urlAddr, _ := url.Parse(addr)
	// 	urlMsgAddrs = append(urlMsgAddrs, urlAddr)
	// }
	// msgProxy := &httputil.ReverseProxy{Director: ctx.CustomDirector(urlMsgAddrs)}
	// summaryProxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: summaryaddr})
	mux := http.NewServeMux()
	// replace summary handler w proxy
	// mux.Handle("/v1/channels/", msgProxy)
	// mux.Handle("/v1/messages/", msgProxy)
	// mux.Handle("/v1/summary", summaryProxy)

	//mux.HandleFunc("/v1/ws", ctx.WebSocketConnectionHandler)
	mux.HandleFunc("/v1/users", ctx.UsersHandler)
	mux.HandleFunc("/v1/users/", ctx.SpecificUserHandler)
	mux.HandleFunc("/v1/sessions", ctx.SessionsHandler)
	mux.HandleFunc("/v1/sessions/", ctx.SpecificSessionHandler)

	wrappedMux := &handlers.CorsMiddleWare{Handler: mux}

	log.Printf("server is listening at %s...", addr)

	log.Fatal(http.ListenAndServeTLS(addr, tlscert, tlskey, wrappedMux))
}

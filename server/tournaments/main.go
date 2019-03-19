package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"info441finalproject/server/gateway/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	tournamentAddr := os.Getenv("TOURNAMENTADDR")
	if len(tournamentAddr) == 0 {
		tournamentAddr = ":80"
	}
	dsn := os.Getenv("DSN")
	sqlDb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sqlDb.Close()
	if err := sqlDb.Ping(); err != nil {
		fmt.Printf("error pinging database: %v\n", err)
	} else {
		fmt.Printf("successfully connected!\n")
	}
	sqlStore := models.NewMySQLStore(sqlDb)

	//Connect to rabbit server
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

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
	ctx := &TournamentContext{
		UserStore:     sqlStore,
		RabbitChannel: ch,
		QueueName:     "notify",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/smashqq/tournaments", ctx.TourneyHandler)
	mux.HandleFunc("/smashqq/tournaments/", ctx.TourneyHandler)
	mux.HandleFunc("/smashqq/tournaments/{tournamentID}/players", ctx.PlayerHandler)
	mux.HandleFunc("/smashqq/tournaments/{tournamentID}/organizers", ctx.OrganizerHandler)
	mux.HandleFunc("/smashqq/tournaments/{tournamentID}/games", ctx.GamesHandler)
	mux.HandleFunc("/smashqq/tournaments/{tournamentID}/standings", ctx.StandingsHandler)
	log.Printf("server is listening at %s...", tournamentAddr)
	log.Fatal(http.ListenAndServe(tournamentAddr, mux))
}

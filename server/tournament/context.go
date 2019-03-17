package main

import (
	"github.com/streadway/amqp"
)

//TournamentContext holds access to session variables and
//is a receiver on various HTTP handler functions
type TournamentContext struct {
	UserStore     users.Store
	RabbitChannel *amqp.Channel
	QueueName     string
}
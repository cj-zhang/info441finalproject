package main

import (
	"info441finalproject/server/gateway/models"

	"github.com/streadway/amqp"
)

//TournamentContext holds access to session variables and
//is a receiver on various HTTP handler functions
type TournamentContext struct {
	UserStore     models.Store
	RabbitChannel *amqp.Channel
	QueueName     string
}

// STANDINGS DATA STRUCTURE NOTES
//
// Standing for one player should include:
// current placing (num value)
// current standing (dqed, out, in)
// list of games played and the brackets they were in
// Next game and what bracket its in
//
// Overall Standings should include:
// List of player standings in ascending order

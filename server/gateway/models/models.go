package models

import "github.com/info441/info441finalproject/server/gateway/models/users"

//Tournament represents a tournament
type Tournament struct {
	ID        int64       `json:"id,omitempty"`
	URL       string      `json:"url,omitempty"`
	Location  string      `json:"location"`
	Organizer *users.User `json:"organizer"`
	PhotoURL  string      `json:"photoURL"`
}

type Standing struct {
	ID        int64   `json:"id,omitempty"`
	Placing   int     `json:"placing"`
	Standing  string  `json:"standing"`
	PastGames *[]Game `json:"pastGames"`
	NextGame  *Game   `json:"NextGame,omitempty"`
}

type StandingUpdate struct {
	Placing   int     `json:"placing"`
	Standing  string  `json:"standing"`
	PastGames *[]Game `json:"pastGames"`
	NextGame  *Game   `json:"NextGame,omitempty"`
}

type Game struct {
	ID        int64       `json:"id,omitempty"`
	URL       string      `json:"url,omitempty"`
	Location  string      `json:"location"`
	Organizer *users.User `json:"organizer"`
	PhotoURL  string      `json:"photoURL"`
}

type GameUpdate struct {
	ID        int64       `json:"id,omitempty"`
	URL       string      `json:"url,omitempty"`
	Location  string      `json:"location"`
	Organizer *users.User `json:"organizer"`
	PhotoURL  string      `json:"photoURL"`
}

// *TODO*
// Create standings struct
// Create standings update struct
// Create games struct
// Create games update struct

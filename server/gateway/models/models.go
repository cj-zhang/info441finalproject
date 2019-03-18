package models

import "time"

//Tournament represents a tournament
type Tournament struct {
	ID        int64  `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	Location  string `json:"location"`
	Organizer *User  `json:"organizer"`
	PhotoURL  string `json:"photoURL"`
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
	ID                    int64      `json:"id,omitempty"`
	TournamentID          int64      `json:"tournamentId,omitempty"`
	PlayerOne             int64      `json:"playerOne,omitempty"`
	PlayerTwo             int64      `json:"playerTwo,omitempty"`
	Victor                int64      `json:"Victor,omitempty"`
	DateTime              *time.Time `json:"datetime,omitempty"`
	BracketID             int64      `json:"bracketId,omitempty"`
	TournamentOrganizerID int64      `json:"tournamentOrganizerId,omitempty"`
	InProgress            bool       `json:"inProgress,omitempty"`
	Completed             bool       `json:"completed,omitempty"`
	Result                string     `json:"result,omitempty"`
}

type GameUpdate struct {
	ID         int64      `json:"id,omitempty"`
	PlayerOne  int64      `json:"playerOne,omitempty"`
	PlayerTwo  int64      `json:"playerTwo,omitempty"`
	Victor     int64      `json:"Victor,omitempty"`
	DateTime   *time.Time `json:"datetime,omitempty"`
	InProgress bool       `json:"inProgress,omitempty"`
	Completed  bool       `json:"completed,omitempty"`
	Result     string     `json:"result,omitempty"`
}

// *TODO*
// Create standings struct
// Create standings update struct

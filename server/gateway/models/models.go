package models

import "time"

//Tournament represents a tournament
type Tournament struct {
	ID        int64  `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	Location  string `json:"location"`
	Organizer int64  `json:"organizer"`
	PhotoURL  string `json:"photoURL"`
}

//TournamentUpdate represents a tournament
type TournamentUpdate struct {
	URL       string `json:"url,omitempty"`
	Location  string `json:"location,omitempty"`
	Organizer int64  `json:"organizer,omitempty"`
	PhotoURL  string `json:"photoURL,omitempty"`
	Open      bool   `json:"open,omitempty"`
}

// Standing represents player standing at a tournament
type Standing struct {
	UserID       int64   `json:"userID"`
	TournamentID int64   `json:"tournamentID"`
	Placing      int     `json:"placing"`
	Standing     string  `json:"standing"`
	PastGames    *[]Game `json:"pastGames,omitempty"`
	NextGame     int64   `json:"nextGame,omitempty"`
}

// Game represents a game between two players at a tournament
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

// GameUpdate represents an update to a game
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

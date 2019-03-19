package models

//Tournament represents a tournament
type Tournament struct {
	ID        int64  `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	Location  string `json:"location"`
	Organizer int64  `json:"organizer"`
	PhotoURL  string `json:"photoURL"`
	Open      bool   `json:"open"`
}

//TournamentUpdate represents a tournament
type TournamentUpdate struct {
	URL       string `json:"url,omitempty"`
	Location  string `json:"location,omitempty"`
	Organizer int64  `json:"organizer,omitempty"`
	Open      bool   `json:"open,omitempty"`
	PhotoURL  string `json:"photoURL,omitempty"`
}

// Standing represents player standing at a tournament
type Standing struct {
	UserID       int64   `json:"userID"`
	TournamentID int64   `json:"tournamentID"`
	Placing      int     `json:"placing"`
	Standing     string  `json:"standing"`
	PastGames    *[]Game `json:"pastGames,omitempty"`
}

// Game represents a game between two players at a tournament
type Game struct {
	ID                    int64  `json:"id,omitempty"`
	TournamentID          int64  `json:"tournamentId,omitempty"`
	PlayerOne             int64  `json:"playerOne,omitempty"`
	PlayerTwo             int64  `json:"playerTwo,omitempty"`
	Victor                int64  `json:"victor,omitempty"`
	TournamentOrganizerID int64  `json:"tournamentOrganizerId,omitempty"`
	InProgress            bool   `json:"inProgress"`
	Completed             bool   `json:"completed"`
	Result                string `json:"result,omitempty"`
	NextGame              int64  `json:"nextGame,omitempty"`
}

// GameUpdate represents an update to a game
type GameUpdate struct {
	ID         int64  `json:"id"`
	PlayerOne  int64  `json:"playerOne"`
	PlayerTwo  int64  `json:"playerTwo"`
	Victor     int64  `json:"victor"`
	InProgress bool   `json:"inProgress"`
	Completed  bool   `json:"completed"`
	Result     string `json:"result"`
}

// *TODO*
// Create standings struct
// Create standings update struct

package models

import (
	"errors"
	"time"
)

//ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

//ErrTournamentNotFound is returned when the tournament can't be found
var ErrTournamentNotFound = errors.New("tournament not found")

//Store represents a store for Users
type Store interface {
	LogUserSignIn(id int64, datetime time.Time, ip string) error

	//GetByID returns the User with the given ID
	GetByID(id int64) (*User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*User, error)

	//GetByUserName returns the User with the given Username
	GetByUserName(username string) (*User, error)

	//Insert inserts the user into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(user *User) (*User, error)

	//Update applies UserUpdates to the given user ID
	//and returns the newly-updated user
	Update(id int64, updates *Updates) (*User, error)

	//Delete deletes the user with the given ID
	Delete(id int64) error

	// GetAllTournaments gets all of the tournaments
	GetAllTournaments() ([]*Tournament, error)

	// GetTournament gets the information for one tournament
	GetTournament(id int64) (*Tournament, error)

	// DeleteTournament deletes the tournament with the given ID
	DeleteTournament(id int64) error

	// CreateTournament inserts a new tournament into the database
	CreateTournament(t *Tournament, creator int64) (*Tournament, error)

	// UpdateTournament updates a tournament with the given updates
	UpdateTournament(tID int64, tu *TournamentUpdate) (*Tournament, error)

	// GetAllPlayers gets all players from a given tournament
	GetAllPlayers(tID int64) ([]*User, error)

	// GetPlayers gets the information for a given amount of players from users
	GetPlayers(q int, tID int64) ([]*User, error)

	// RegisterPlayer inserts a new user into the players table
	RegisterPlayer(id int64, tID int64) error

	// RemovePlayer deletes a user from the players table
	RemovePlayer(id int64, tID int64) error

	// RegisterTO inserts a new TO into the TO table
	RegisterTO(id int64, tID int64) error

	// RemoveTO deletes a TO from the TO table
	RemoveTO(id int64, tID int64) error

	// GetTO gets the information for a given game from the tournament_organizers table
	GetTO(id int64, tID int64) (*User, error)

	// GetTOs gets the information for a given amount of tournament organizers from users
	GetTOs(q int, tID int64) ([]*User, error)

	// GetLeastBusyTO gets the TO with the least amount of brackets overseen
	GetLeastBusyTO(tID int64) (*User, error)

	// CreateGame creates and inserts a new game into the games table
	CreateGame(tID int64, g *Game) (*Game, error)

	// GetGame gets the information for a given game from the games table
	GetGame(gID int64) (*Game, error)

	// GetGames gets the information for a given amount of games from the games table
	GetGames(q int, tID int64) ([]*Game, error)

	// ReportGame applies given updates to a game
	ReportGame(updates *GameUpdate) (*Game, error)

	// UserIsTO checks if a given user is a tournament organizer for the given tournament
	UserIsTO(id int64, tID int64) bool

	// GetStanding gets a single standing for the given user at a given tournament
	GetStanding(id int64, tID int64) (*Standing, error)

	// GetStandings gets the standings associated with a given tournament
	GetStandings(q int, tID int64) ([]*Standing, error)
}

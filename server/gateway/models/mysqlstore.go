package models

import (
	"database/sql"
	"time"
)

const selAllSQL = "select id, username, first_name, last_name from users"
const insertSignIn = "insert into signins(id, login_time, ip_addr) values (?,?,?)"
const idSelect = "Select * From users Where id=?"
const emailSelect = "Select * From users Where email=?"
const usernameSelect = "Select * From users Where username=?"
const insertUser = "insert into users(email, pass_hash, username, first_name, last_name, photo_url) values (?,?,?,?,?,?)"
const updateUser = "update users set first_name=?, last_name=? where id=?"
const deleteUser = "delete from users where id=?"
const getTournament = "Select * From tournaments Where id=?"
const deleteTournament = "delete From tournaments Where id=?"
const insertTournament = "insert into tournaments(website, tournament_location, tournament_organizer_id, photo_url, open) values (?,?,?,?,?)"
const updateTournament = "update tournaments set website=?, tournament_location=?, tournament_organizer_id=?, registration_open=?, photo_url=? where id=?"
const insertPlayer = "insert into players(u_id, tournament_id) values (?,?)"
const deletePlayer = "delete From tournaments Where u_id=? and tournament_id=?"
const getPlayers = "Select id, email, username, pass_hash, first_name, last_name, photo_url From users u join players p on u.id = p.u_id where p.tournament_id=? limit ?"
const getTO = "Select id, email, username, pass_hash, first_name, last_name, photo_url From users u join tournament_organizers to on u.id = to.u_id where to.u_id=? and to.tournament_id=?"
const getTOs = "Select id, email, username, pass_hash, first_name, last_name, photo_url From users u join tournament_organizers to on u.id = to.u_id where to.tournament_id=? limit ?"
const insertTO = "insert into tournament_organizers(u_id, tournament_id) values (?,?)"
const deleteTO = "delete From tournaments_organizers Where u_id=? and tournament_id=?"
const getLeastBusyTO = "select top(1) u_id from tournament_organizers where tournament_id=? order by brackets_overseen asc"
const getGame = "Select * From games where id=?"
const getGames = "Select * From games where tournament_id=? limit ?"
const createGame = "insert into games(tournament_id, player_one, player_two, victor, date_time, tournament_organizer_id, in_progress, completed, result) values (?,?,?,?,?,?,?,?,?)"
const updateGame = "update games set player_one=?, player_two=?, victor=?, date_time=?, in_progress=?, completed=?, result=? where id=?"
const checkIfTO = "Select brackets_overseen from tournament_organizers where u_id=? and tournament_id=?"
const getStanding = "Select * from standings where u_id=? and tournament_id=?"
const getStandings = "Select * from standings where tournament_id=? limit ?"

// MySQLStore implements the Store interface and holds a pointer to a db
type MySQLStore struct {
	Client *sql.DB
}

// NewMySQLStore constructs and returns a pointer to a MySQLStore struct
func NewMySQLStore(db *sql.DB) *MySQLStore {
	if db != nil {
		return &MySQLStore{
			Client: db,
		}
	}
	return nil
}

// LogUserSignIn logs a user signin
func (store *MySQLStore) LogUserSignIn(id int64, datetime time.Time, ip string) error {
	_, err := store.Client.Exec(insertSignIn, id, datetime, ip)
	if err != nil {
		return err
	}
	return nil
}

// GetByID returns a user struct populated by a database row with a matching id
func (store *MySQLStore) GetByID(id int64) (*User, error) {
	return store.GetByParam(idSelect, id)

}

// GetByEmail returns a user struct populated by a database row with a matching email
func (store *MySQLStore) GetByEmail(email string) (*User, error) {
	return store.GetByParam(emailSelect, email)
}

// GetByUserName returns a user struct populated by a database row with a matching username
func (store *MySQLStore) GetByUserName(username string) (*User, error) {
	return store.GetByParam(usernameSelect, username)
}

// GetByParam returns a user struct populated by a database row with matching column values
func (store *MySQLStore) GetByParam(selectSQL string, paramVal interface{}) (*User, error) {
	u := &User{}
	row := store.Client.QueryRow(selectSQL, paramVal)
	if err := row.Scan(&u.ID, &u.Email, &u.PassHash, &u.UserName, &u.FirstName, &u.LastName, &u.PhotoURL); err != nil {
		return nil, ErrUserNotFound
	}

	return u, nil
}

// Insert inserts the user into the database, and returns
// the newly-inserted User, complete with the DBMS-assigned ID
func (store *MySQLStore) Insert(user *User) (*User, error) {
	res, err := store.Client.Exec(insertUser, user.Email, user.PassHash, user.UserName, user.FirstName, user.LastName, user.PhotoURL)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = id
	return user, nil
}

// Update applies Updates to the given user ID and returns the newly-updated user
func (store *MySQLStore) Update(id int64, updates *Updates) (*User, error) {
	_, err := store.Client.Exec(updateUser, updates.FirstName, updates.LastName, id)
	if err != nil {
		return nil, err
	}

	return store.GetByID(id)
}

// Delete deletes the user with the given ID
func (store *MySQLStore) Delete(id int64) error {
	_, err := store.Client.Exec(deleteUser, id)
	if err != nil {
		return err
	}
	return nil
}

// GetTournament gets the information for one tournament
func (store *MySQLStore) GetTournament(id int64) (*Tournament, error) {
	t := &Tournament{}
	row := store.Client.QueryRow(getTournament, id)
	if err := row.Scan(&t.ID, &t.URL, &t.Location, &t.Organizer, &t.PhotoURL, &t.Open); err != nil {
		return nil, ErrTournamentNotFound
	}

	return t, nil
}

// DeleteTournament deletes the tournament with the given ID
func (store *MySQLStore) DeleteTournament(id int64) error {
	_, err := store.Client.Exec(deleteTournament, id)
	if err != nil {
		return err
	}
	return nil
}

// CreateTournament inserts a new tournament into the database
func (store *MySQLStore) CreateTournament(t *Tournament, creator int64) (*Tournament, error) {
	res, err := store.Client.Exec(insertTournament, t.URL, t.Location, creator, t.PhotoURL, t.Open)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

//

// UpdateTournament updates a tournament with the given updates
func (store *MySQLStore) UpdateTournament(tID int64, tu *TournamentUpdate) (*Tournament, error) {
	_, err := store.Client.Exec(updateTournament, tu.URL, tu.Location, tu.Organizer, tu.Open, tu.PhotoURL, tID)
	if err != nil {
		return nil, err
	}
	return store.GetTournament(tID)
}

// GetAllPlayers gets all players from a given tournament
func (store *MySQLStore) GetAllPlayers(tID int64) ([]*User, error) {
	return store.GetPlayers(0, tID)
}

// GetPlayers gets the information for a given amount of players from users
func (store *MySQLStore) GetPlayers(q int, tID int64) ([]*User, error) {
	var result []*User
	rows, err := store.Client.Query(getPlayers, tID, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := &User{}
		err = rows.Scan(&u.ID, &u.Email, &u.UserName, &u.PassHash, &u.FirstName, &u.LastName, &u.PhotoURL)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

// RegisterPlayer inserts a new user into the players table
func (store *MySQLStore) RegisterPlayer(id int64, tID int64) error {
	_, err := store.Client.Exec(insertPlayer, id, tID)
	if err != nil {
		return err
	}
	return nil
}

// RemovePlayer deletes a user from the players table
func (store *MySQLStore) RemovePlayer(id int64, tID int64) error {
	_, err := store.Client.Exec(deletePlayer, id, tID)
	if err != nil {
		return err
	}
	return nil
}

// RegisterTO inserts a new TO into the TO table
func (store *MySQLStore) RegisterTO(id int64, tID int64) error {
	_, err := store.Client.Exec(insertTO, id, tID)
	if err != nil {
		return err
	}
	return nil
}

// RemoveTO deletes a TO from the TO table
func (store *MySQLStore) RemoveTO(id int64, tID int64) error {
	_, err := store.Client.Exec(deleteTO, id, tID)
	if err != nil {
		return err
	}
	return nil
}

// GetTO gets the information for a given user from the tournament_organizers table
func (store *MySQLStore) GetTO(id int64, tID int64) (*User, error) {
	u := &User{}
	row := store.Client.QueryRow(getTO, id, tID)
	if err := row.Scan(&u.ID, &u.Email, &u.UserName, &u.PassHash, &u.FirstName, &u.LastName, &u.PhotoURL); err != nil {
		return nil, err
	}

	return u, nil
}

// GetTOs gets the information for a given amount of tournament organizers from users
func (store *MySQLStore) GetTOs(q int, tID int64) ([]*User, error) {
	var result []*User
	rows, err := store.Client.Query(getTOs, tID, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		u := &User{}
		err = rows.Scan(&u.ID, &u.Email, &u.UserName, &u.PassHash, &u.FirstName, &u.LastName, &u.PhotoURL)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

// GetLeastBusyTO gets the TO with the least amount of brackets overseen
func (store *MySQLStore) GetLeastBusyTO(tID int64) (*User, error) {
	var userID int64
	row := store.Client.QueryRow(getLeastBusyTO, tID)
	if err := row.Scan(userID); err != nil {
		return nil, err
	}
	return store.GetTO(userID, tID)
}

// CreateGame creates and inserts a new game into the games table
func (store *MySQLStore) CreateGame(tID int64, g *Game) (*Game, error) {
	res, err := store.Client.Exec(createGame, tID, g.PlayerOne, g.PlayerTwo, g.Victor, g.DateTime, g.TournamentOrganizerID, g.InProgress, g.Completed, g.Result)
	if err != nil {
		return nil, err
	}
	gameID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	g.ID = gameID
	return g, nil
}

// GetGame gets the information for a given game from the games table
func (store *MySQLStore) GetGame(gID int64) (*Game, error) {
	g := &Game{}
	row := store.Client.QueryRow(getGame, gID)
	if err := row.Scan(&g.ID, &g.TournamentID, &g.PlayerOne, &g.PlayerTwo, &g.Victor, &g.DateTime, &g.TournamentOrganizerID, &g.InProgress, &g.Completed, &g.Result); err != nil {
		return nil, err
	}
	return g, nil
}

// GetGames gets the information for a given amount of games from the games table
func (store *MySQLStore) GetGames(q int, tID int64) ([]*Game, error) {
	var result []*Game
	rows, err := store.Client.Query(getGames, tID, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		g := &Game{}
		err = rows.Scan(&g.ID, &g.TournamentID, &g.PlayerOne, &g.PlayerTwo, &g.Victor, &g.DateTime, &g.TournamentOrganizerID, &g.InProgress, &g.Completed, &g.Result)
		if err != nil {
			return nil, err
		}
		result = append(result, g)
	}
	return result, nil
}

// ReportGame applies given updates to a game
func (store *MySQLStore) ReportGame(updates *GameUpdate) (*Game, error) {
	_, err := store.Client.Exec(updateGame, updates.PlayerOne, updates.PlayerTwo, updates.Victor, updates.DateTime, updates.InProgress, updates.Completed, updates.Result, updates.ID)
	if err != nil {
		return nil, err
	}

	return store.GetGame(updates.ID)
}

// UserIsTO checks if a given user is a tournament organizer for the given tournament
func (store *MySQLStore) UserIsTO(id int64, tID int64) bool {
	var bracketsOverseen int
	row := store.Client.QueryRow(checkIfTO, id, tID)
	err := row.Scan(&bracketsOverseen)
	return (err == nil && err != sql.ErrNoRows)
}

// GetStanding gets a single standing for the given user at a given tournament
func (store *MySQLStore) GetStanding(id int64, tID int64) (*Standing, error) {
	s := &Standing{}
	row := store.Client.QueryRow(getStanding, id, tID)
	if err := row.Scan(&s.UserID, &s.TournamentID, &s.Placing, &s.Standing); err != nil {
		return nil, err
	}
	return s, nil

}

// GetStandings gets the standings associated with a given tournament
func (store *MySQLStore) GetStandings(q int, tID int64) ([]*Standing, error) {
	var result []*Standing
	rows, err := store.Client.Query(getStandings, tID, q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		s := &Standing{}
		if err = rows.Scan(&s.UserID, &s.TournamentID, &s.Placing, &s.Standing); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

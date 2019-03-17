package users

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

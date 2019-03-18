package models

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const GETSTATEMENT = "Select * From users Where id=?"
const GETEMAILSTATEMENT = "Select * From users Where email=?"
const GETUSERNAMESTATEMENT = "Select * From users Where username=?"
const INSERTSTATEMENT = "insert into users(email, pass_hash, username, first_name, last_name, photo_url) values (?,?,?,?,?,?)"
const UPDATESTATEMENT = "update users set first_name=?, last_name=? where id=?"
const DELSTATEMENT = "delete from users where id=?"

func TestMySQLStore_GetByID(t *testing.T) {
	//create a new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error creating sql mock: %v", err)
	}

	//ensure it's closed at the end of the test
	defer db.Close()

	// Initialize a user struct we will use as a test variable
	expectedUser := createTestUser()

	// Initialize a MySQLStore struct to allow us to interface with the SQL client
	store := NewMySQLStore(db)

	// Create a row with the appropriate fields in your SQL database
	// Add the actual values to the row
	row := sqlmock.NewRows([]string{"id", "email", "pass_hash", "username", "firstname", "lastname", "photo_url"})
	row.AddRow(expectedUser.ID, expectedUser.Email, expectedUser.PassHash, expectedUser.UserName, expectedUser.FirstName, expectedUser.LastName, expectedUser.PhotoURL)

	// Expecting a successful "query"
	// This tells our db to expect this query (id) as well as supply a certain response (row)
	// REMINDER: Since sqlmock requires a regex string, in order for `?` to be interpreted, you'll
	// have to wrap it within a `regexp.QuoteMeta`. Be mindful that you will need to do this EVERY TIME you're
	// using any reserved metacharacters in regex.
	mock.ExpectQuery(regexp.QuoteMeta(GETSTATEMENT)).
		WithArgs(expectedUser.ID).WillReturnRows(row)

	// Since we know our query is successful, we want to test whether there happens to be
	// any expected error that may occur.
	user, err := store.GetByID(expectedUser.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Again, since we are assuming that our query is successful, we can test for when our
	// function doesn't work as expected.
	if err == nil && !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("User queried does not match expected user")
	}

	// Expecting a unsuccessful "query"
	// Attempting to search by an id that doesn't exist. This would result in a
	// sql.ErrNoRows error
	// REMINDER: Using a constant makes your code much clear, and is highly recommended.
	mock.ExpectQuery(regexp.QuoteMeta(GETSTATEMENT)).
		WithArgs(-1).WillReturnError(sql.ErrNoRows)

	// Since we are expecting an error here, we create a condition opposing that to see
	// if our GetById is working as expected
	if _, err = store.GetByID(-1); err == nil {
		t.Errorf("Expected error: %v, but recieved nil", sql.ErrNoRows)
	}

	// This attempts to check if there are any expectations that we haven't met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet sqlmock expectations: %v", err)
	}

}

func TestMySQLStore_GetByEmail(t *testing.T) {
	//create a new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error creating sql mock: %v", err)
	}

	//ensure it's closed at the end of the test
	defer db.Close()

	// Initialize a user struct we will use as a test variable
	expectedUser := createTestUser()

	// Initialize a MySQLStore struct to allow us to interface with the SQL client
	store := NewMySQLStore(db)

	// Create a row with the appropriate fields in your SQL database
	// Add the actual values to the row
	row := sqlmock.NewRows([]string{"id", "email", "pass_hash", "username", "firstname", "lastname", "photo_url"})
	row.AddRow(expectedUser.ID, expectedUser.Email, expectedUser.PassHash, expectedUser.UserName, expectedUser.FirstName, expectedUser.LastName, expectedUser.PhotoURL)

	// Expecting a successful "query"
	// This tells our db to expect this query (id) as well as supply a certain response (row)
	// REMINDER: Since sqlmock requires a regex string, in order for `?` to be interpreted, you'll
	// have to wrap it within a `regexp.QuoteMeta`. Be mindful that you will need to do this EVERY TIME you're
	// using any reserved metacharacters in regex.
	mock.ExpectQuery(regexp.QuoteMeta(GETEMAILSTATEMENT)).
		WithArgs(expectedUser.Email).WillReturnRows(row)

	// Since we know our query is successful, we want to test whether there happens to be
	// any expected error that may occur.
	user, err := store.GetByEmail(expectedUser.Email)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Again, since we are assuming that our query is successful, we can test for when our
	// function doesn't work as expected.
	if err == nil && !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("User queried does not match expected user")
	}

	// Expecting a unsuccessful "query"
	// Attempting to search by an id that doesn't exist. This would result in a
	// sql.ErrNoRows error
	// REMINDER: Using a constant makes your code much clear, and is highly recommended.
	mock.ExpectQuery(regexp.QuoteMeta(GETEMAILSTATEMENT)).
		WithArgs("").WillReturnError(sql.ErrNoRows)

	// Since we are expecting an error here, we create a condition opposing that to see
	// if our GetById is working as expected
	if _, err = store.GetByEmail(""); err == nil {
		t.Errorf("Expected error: %v, but recieved nil", sql.ErrNoRows)
	}

	// This attempts to check if there are any expectations that we haven't met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet sqlmock expectations: %v", err)
	}
}

func TestMySQLStore_GetByUserName(t *testing.T) {
	//create a new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error creating sql mock: %v", err)
	}

	//ensure it's closed at the end of the test
	defer db.Close()

	// Initialize a user struct we will use as a test variable
	expectedUser := createTestUser()

	// Initialize a MySQLStore struct to allow us to interface with the SQL client
	store := NewMySQLStore(db)

	// Create a row with the appropriate fields in your SQL database
	// Add the actual values to the row
	row := sqlmock.NewRows([]string{"id", "email", "pass_hash", "username", "firstname", "lastname", "photo_url"})
	row.AddRow(expectedUser.ID, expectedUser.Email, expectedUser.PassHash, expectedUser.UserName, expectedUser.FirstName, expectedUser.LastName, expectedUser.PhotoURL)

	// Expecting a successful "query"
	// This tells our db to expect this query (id) as well as supply a certain response (row)
	// REMINDER: Since sqlmock requires a regex string, in order for `?` to be interpreted, you'll
	// have to wrap it within a `regexp.QuoteMeta`. Be mindful that you will need to do this EVERY TIME you're
	// using any reserved metacharacters in regex.
	mock.ExpectQuery(regexp.QuoteMeta(GETUSERNAMESTATEMENT)).
		WithArgs(expectedUser.UserName).WillReturnRows(row)

	// Since we know our query is successful, we want to test whether there happens to be
	// any expected error that may occur.
	user, err := store.GetByUserName(expectedUser.UserName)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Again, since we are assuming that our query is successful, we can test for when our
	// function doesn't work as expected.
	if err == nil && !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("User queried does not match expected user")
	}

	// Expecting a unsuccessful "query"
	// Attempting to search by an id that doesn't exist. This would result in a
	// sql.ErrNoRows error
	// REMINDER: Using a constant makes your code much clear, and is highly recommended.
	mock.ExpectQuery(regexp.QuoteMeta(GETUSERNAMESTATEMENT)).
		WithArgs("").WillReturnError(sql.ErrNoRows)

	// Since we are expecting an error here, we create a condition opposing that to see
	// if our GetById is working as expected
	if _, err = store.GetByUserName(""); err == nil {
		t.Errorf("Expected error: %v, but recieved nil", sql.ErrNoRows)
	}

	// This attempts to check if there are any expectations that we haven't met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet sqlmock expectations: %v", err)
	}
}

func TestMySQLStore_Insert(t *testing.T) {
	//create a new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error creating sql mock: %v", err)
	}
	//ensure it's closed at the end of the test
	defer db.Close()

	// Initialize a user struct we will use as a test variable
	inputUser := createTestUser()

	// Initialize a MySQLStore struct to allow us to interface with the SQL client
	store := NewMySQLStore(db)

	// This tells our db to expect an insert query with certain arguments with a certain
	// return result
	mock.ExpectExec(regexp.QuoteMeta(INSERTSTATEMENT)).
		WithArgs(inputUser.Email, inputUser.PassHash, inputUser.UserName, inputUser.FirstName, inputUser.LastName, inputUser.PhotoURL).
		WillReturnResult(sqlmock.NewResult(2, 1))

	user, err := store.Insert(inputUser)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err == nil && !reflect.DeepEqual(user, inputUser) {
		t.Errorf("User returned does not match input user")
	}

	// Inserting an invalid user
	invalidUser := &User{
		ID:        -1,
		Email:     "test@test.com",
		PassHash:  []byte("testhash"),
		UserName:  "testusername",
		FirstName: "",
		LastName:  "Doe",
		PhotoURL:  "https://www.gravatar.com/avatar/",
	}
	insertErr := fmt.Errorf("Error executing INSERT operation")
	mock.ExpectExec(regexp.QuoteMeta(INSERTSTATEMENT)).
		WithArgs(invalidUser.Email, invalidUser.PassHash, invalidUser.UserName, invalidUser.FirstName, invalidUser.LastName, invalidUser.PhotoURL).
		WillReturnError(insertErr)

	if _, err = store.Insert(invalidUser); err == nil {
		t.Errorf("Expected error: %v but recieved nil", insertErr)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unmet sqlmock expectations: %v", err)
	}

}

func TestMySQLStore_Update(t *testing.T) {
	//create a new sql mock
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error creating sql mock: %v", err)
	}
	//ensure it's closed at the end of the test
	defer db.Close()

	// Initialize a user and update struct we will use as a test variable
	inputUser := createTestUser()

	updates := &Updates{
		FirstName: "updatedFirstName",
		LastName:  "updatedLastName",
	}

	// Initialize a MySQLStore struct to allow us to interface with the SQL client
	store := NewMySQLStore(db)
	row := sqlmock.NewRows([]string{"id", "email", "pass_hash", "username", "firstname", "lastname", "photo_url"})

	// This tells our db to expect an insert query with certain arguments with a certain
	// return result
	row.AddRow(inputUser.ID, inputUser.Email, inputUser.PassHash, inputUser.UserName, inputUser.FirstName, inputUser.LastName, inputUser.PhotoURL)

	mock.ExpectExec(regexp.QuoteMeta(UPDATESTATEMENT)).
		WithArgs(inputUser.ID).
		WillReturnResult(sqlmock.NewResult(inputUser.ID, 1))

	mock.ExpectQuery(regexp.QuoteMeta(GETSTATEMENT)).
		WithArgs(inputUser.ID).WillReturnRows(row)

	u, err := store.Update(inputUser.ID, updates)
	if err == nil && !reflect.DeepEqual(inputUser, u) {
		t.Errorf("User returned does not match model output user")
	}
}

func TestMySQLStore_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error creating sql mock: %v", err)
	}
	//ensure it's closed at the end of the test
	defer db.Close()

	inputUser := createTestUser()

	store := NewMySQLStore(db)
	row := sqlmock.NewRows([]string{"id", "email", "pass_hash", "username", "firstname", "lastname", "photo_url"})
	row.AddRow(inputUser.ID, inputUser.Email, inputUser.PassHash, inputUser.UserName, inputUser.FirstName, inputUser.LastName, inputUser.PhotoURL)

	mock.ExpectExec(regexp.QuoteMeta(DELSTATEMENT)).
		WithArgs(inputUser.ID).
		WillReturnResult(sqlmock.NewResult(inputUser.ID, 1))

	if err = store.Delete(-1); err == nil {
		t.Errorf("Expected error: %v but recieved nil", err)
	}
}

func createTestUser() *User {
	return &User{
		ID:        1,
		Email:     "test@test.com",
		PassHash:  []byte("testhash"),
		UserName:  "testusername",
		FirstName: "John",
		LastName:  "Doe",
		PhotoURL:  "https://www.gravatar.com/avatar/",
	}
}

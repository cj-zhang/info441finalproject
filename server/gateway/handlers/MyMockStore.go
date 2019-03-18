package handlers

import (
	"fmt"
	"info441finalproject/server/gateway/models"
	"time"
)

// MyMockStore mocks a store
type MyMockStore struct {
}

// LogUserSignIn logs a user sign in
func (ms *MyMockStore) LogUserSignIn(id int64, datetime time.Time, ip string) error {
	return nil
}

// GetByID mocks the behavior of GetByID
func (ms *MyMockStore) GetByID(id int64) (*models.User, error) {
	// We can trigger an error by passing in an id of 2
	if id == 2 {
		return nil, fmt.Errorf("Error getting user with id: %d", id)
	}
	user := &models.User{
		ID:        1,
		Email:     "test@test.com",
		PassHash:  []byte("passwordHash"),
		UserName:  "usernameTest",
		FirstName: "testFirst",
		LastName:  "testLast",
		PhotoURL:  "testphotourl.com",
	}

	return user, nil
}

// GetByEmail mocks the behavior of GetByEmail
func (ms *MyMockStore) GetByEmail(email string) (*models.User, error) {
	// We can trigger an error by passing in an email of "two"
	if email == "two" {
		return nil, fmt.Errorf("Error getting user with email: %s", email)
	}
	user := &models.User{
		ID:        1,
		Email:     "test@test.com",
		PassHash:  []byte("passwordHash"),
		UserName:  "usernameTest",
		FirstName: "testFirst",
		LastName:  "testLast",
		PhotoURL:  "testphotourl.com",
	}

	return user, nil
}

// GetByUserName mocks the behavior of GetByUserName
func (ms *MyMockStore) GetByUserName(username string) (*models.User, error) {
	// We can trigger an error by passing in an email of "two"
	if username == "two" {
		return nil, fmt.Errorf("Error getting user with username: %s", username)
	}
	user := &models.User{
		ID:        1,
		Email:     "test@test.com",
		PassHash:  []byte("passwordHash"),
		UserName:  "usernameTest",
		FirstName: "testFirst",
		LastName:  "testLast",
		PhotoURL:  "testphotourl.com",
	}

	return user, nil
}

// Insert mocks the behavior of Insert
func (ms *MyMockStore) Insert(user *models.User) (*models.User, error) {
	if user.FirstName == "Error" {
		return nil, fmt.Errorf("Error Inserting New User")
	}

	// Assumes that if user without the FirstName field being equal to "Error" will
	// always result in a successful insert
	return user, nil
}

// Update mocks the behavior of Update
func (ms *MyMockStore) Update(id int64, updates *models.Updates) (*models.User, error) {
	// We can trigger an error by passing in an id of 2
	if id == 2 {
		return nil, fmt.Errorf("Error updating user with id: %d", id)
	}
	user := &models.User{
		ID:        1,
		Email:     "test@test.com",
		PassHash:  []byte("passwordHash"),
		UserName:  "usernameTest",
		FirstName: "testFirst",
		LastName:  "testLast",
		PhotoURL:  "testphotourl.com",
	}

	user.FirstName = updates.FirstName
	user.LastName = updates.LastName
	// Assumes that if user without the FirstName field being equal to "Error" will
	// always result in a successful insert
	return user, nil
}

// Delete mocks the behavior of Delete
func (ms *MyMockStore) Delete(id int64) error {
	if id == 2 {
		return fmt.Errorf("Error deleting user")
	}

	// Assumes that if user without the FirstName field being equal to "Error" will
	// always result in a successful insert
	return nil
}

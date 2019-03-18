package models

import (
	"strings"
	"testing"
)

func makeNewUser() *NewUser {
	return &NewUser{
		Email:        "test@test.com",
		Password:     "password",
		PasswordConf: "password",
		UserName:     "testAcc",
		FirstName:    "testFirst",
		LastName:     "testLast",
	}
}

//TODO: add tests for the various functions in user.go, as described in the assignment.
//use `go test -cover` to ensure that you are covering all or nearly all of your code paths.

func TestValidate(t *testing.T) {
	newUsr := makeNewUser()
	newUsr.Email = "invalid email"
	if err := newUsr.Validate(); err == nil {
		t.Errorf("Expecting error about invalid email\n")
	}
	newUsr.Email = "valid@email.com"
	if err := newUsr.Validate(); err != nil {
		t.Errorf("Should not receive error about invalid email\n")
	}

	newUsr.Password = "five"
	if err := newUsr.Validate(); err == nil {
		t.Errorf("Expecting error about password length\n")
	}
	newUsr.Password = "password1"
	if err := newUsr.Validate(); err == nil {
		t.Errorf("Expecting error about password not matching confirmation\n")
	}
	newUsr.Password = "password"

	newUsr.UserName = ""
	if err := newUsr.Validate(); err == nil {
		t.Errorf("Expecting error about missing username\n")
	}
	newUsr.UserName = "t e s t "
	if err := newUsr.Validate(); err == nil {
		t.Errorf("Expecting error about username containing spaces\n")
	}
}
func TestNewUserToUser(t *testing.T) {
	nu := makeNewUser()
	u, err := nu.ToUser()
	if err != nil {
		t.Errorf("error converting NewUser to User: %s\n", err.Error())
	}
	if nil == u {
		t.Fatalf("ToUser() returned nil\n")
	}
	if nu.Email != u.Email {
		t.Errorf("User.Email !+ NewUser.Email: expected %s but got %s\n", nu.Email, u.Email)
	}
	if len(u.PassHash) == 0 {
		t.Errorf("User.PassHash is zero length, should be hashed password\n")
	}
	if len(u.PhotoURL) == 0 || !strings.HasPrefix(u.PhotoURL, gravatarBasePhotoURL) {
		t.Errorf("User.PhotoURL is zero length, should be gravatar profile image URL\n")
	}
}

func TestFullName(t *testing.T) {
	u, _ := makeNewUser().ToUser()
	u.FirstName = ""
	u.LastName = ""
	if result := u.FullName(); result != "" {
		t.Errorf("expecting empty string\n")
	}
	u.FirstName = "test"
	if result := u.FullName(); result != "test" {
		t.Errorf("expecting 'test'\n")
	}
	u.LastName = "test2"
	if result := u.FullName(); result != "test test2" {
		t.Errorf("expecting 'test test2'\n")
	}
}

func TestSetPassword(t *testing.T) {
	u := &User{}
	if err := u.SetPassword("password"); err != nil {
		t.Errorf("error setting password: %s\n", err.Error())
	}
	if len(u.PassHash) == 0 {
		t.Errorf("u.PassHash is zero length, should be hashed password\n")
	}
	if string(u.PassHash) == "password" {
		t.Errorf("plaintext password was stored in PassHash instead of hashed password\n")
	}
}

func TestAuthenticate(t *testing.T) {
	u := &User{}
	if err := u.SetPassword("password"); err != nil {
		t.Errorf("error setting password: %s\n", err.Error())
	}
	if err := u.Authenticate("password"); err != nil {
		t.Errorf("error authenticating valid password: %s\n", err.Error())
	}
	if err := u.Authenticate("incorrect"); err == nil {
		t.Errorf("no error authenticating incorrect password\n")
	}
}

func TestApplyUpdates(t *testing.T) {
	u := &User{}
	updates := &Updates{FirstName: "test", LastName: "test2"}
	if err := u.ApplyUpdates(nil); err == nil {
		t.Errorf("expecting error about no updates being applied\n")
	}

	if err := u.ApplyUpdates(updates); err != nil {
		t.Errorf("unexpected error occurred when applying valid updates: %s\n", err.Error())
	}
}

package handlers

import (
	"info441finalproject/server/gateway/models/users"
	"time"
)

//SessionState defines a session state struct for this web server
//see the assignment description for the fields you should include
//remember that other packages can only see exported fields!
type SessionState struct {
	StartTime time.Time
	User      *users.User
}

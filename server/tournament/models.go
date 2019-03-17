package main

import "github.com/info441/assignments-kxuojom97/servers/gateway/models/users"

//Tournament represents a tournament
type Tournament struct {
	ID        int64       `json:"id,omitempty"`
	URL       string      `json:"url,omitempty"`
	Location  string      `json:"location"`
	Organizer *users.User `json:"organizer"`
	PhotoURL  string      `json:"photoURL"`
}

package handlers

import "net/http"

//UsersHandler defines HTTP handler functions for users
func (ctx *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	} else if r.Method == "GET" {

	} else if r.Method == "PATCH" {

	} else {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}
}

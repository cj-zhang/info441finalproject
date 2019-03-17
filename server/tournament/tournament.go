package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// /smashqq/users
// /smashqq/tournaments
// /smashqq/tournaments?id={}/standings
// /smashqq/tournaments?id={}/players
// /smashqq/tournaments?id={}/brackets

// TourneyHandler handles requests for the '/smashqq/tournaments/' resource
func (ctx *TournamentContext) TourneyHandler(w http.ResponseWriter, r *http.Request) {
	// *TODO*
	// Check if authenticated
	// *TODO*
	if r.Method == http.MethodGet {
		tid := r.URL.Query().Get("id")
		tournament, err := ctx.UserStore.GetTournament(tid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(channelList); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		header := r.Header.Get("Content-Type")
		if !strings.HasPrefix(header, "application/json") {
			http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
			return
		}
		tournament := new(Tournament)
		if err := json.NewDecoder(r.Body).Decode(channel); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}
		returnTournament, err := ctx.UserStore.CreateTournament(tournament)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(returnTournament); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

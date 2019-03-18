package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/info441/info441finalproject/server/gateway/models"
)

//*TODO discussion questions*
// What exactly do we need to implement?
// Do we need to make the data structure that hold standings?
// Stretch goal for seeing what games are next?

// *TODO*
// Record all the mysql methods that need to be defined
// and record their parameter and return values
//
// Record all the methods that are needed to return
// standings from the data structure
//

// TourneyHandler handles requests for the '/smashqq/tournaments' resource
// and '/smashqq/tournaments/{tournamentID}' resource
// *TODO* maybe add update method?
func (ctx *TournamentContext) TourneyHandler(w http.ResponseWriter, r *http.Request) {
	// *TODO*
	// Check if authenticated
	//
	if r.Method != http.MethodPost {
		urlVar := path.Base(r.URL.String())
		tid, err := strconv.Atoi(urlVar)
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodGet {
			tournament, err := ctx.UserStore.GetTournament(tid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(tournament); err != nil {
				http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
					http.StatusInternalServerError)
				return
			}
		} else if r.Method == http.MethodDelete {
			err = ctx.UserStore.DeleteTournament(tid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("tournament deleted"))
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		header := r.Header.Get("Content-Type")
		if !strings.HasPrefix(header, "application/json") {
			http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
			return
		}
		tournament := new(models.Tournament)
		if err := json.NewDecoder(r.Body).Decode(tournament); err != nil {
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
	}
}

// PlayerHandler handles requests for the '/smashqq/tournaments/{tournamentID}/players' resource
// *TODO* re-evaluate if getting a single player is necessary or just return all players
func (ctx *TournamentContext) PlayerHandler(w http.ResponseWriter, r *http.Request) {
	// *TODO*
	// Check if authenticated
	//
	urlVar := path.Base(path.Dir(r.URL.String()))
	tid, err := strconv.Atoi(urlVar)
	if err != nil {
		http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
		return
	}
	queryID := r.URL.Query().Get("id")
	if r.Method == http.MethodGet {
		var players interface{}
		if queryID == "" {
			query := r.URL.Query().Get("q")
			if query == "" {
				http.Error(w, "Must supply query value", http.StatusBadRequest)
				return
			}
			q, err := strconv.Atoi(query)
			if err != nil {
				http.Error(w, "Must supply a valid query", http.StatusBadRequest)
				return
			}
			players, err = ctx.UserStore.GetPlayers(tid, q)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			uid, err := strconv.Atoi(queryID)
			if err != nil {
				http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
				return
			}
			players, err = ctx.UserStore.GetUser(uid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if players == nil {
				http.Error(w, "Found no user with given user id", http.StatusNotFound)
				return
			}
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(players); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		uid, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		err = ctx.UserStore.RegisterPlayer(uid, tid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Player registered"))
	} else if r.Method == http.MethodDelete {
		if queryID == "" {
			http.Error(w, "Must supply user ID", http.StatusBadRequest)
			return
		}
		uid, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		err = ctx.UserStore.RemovePlayer(uid, tid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Player removed"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// OrganizerHandler handles requests for the '/smashqq/tournaments/{tournamentID}/organizers' resource
func (ctx *TournamentContext) OrganizerHandler(w http.ResponseWriter, r *http.Request) {
	// *TODO*
	// Check if authenticated
	//
	urlVar := path.Base(path.Dir(r.URL.String()))
	tid, err := strconv.Atoi(urlVar)
	if err != nil {
		http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
		return
	}
	queryID := r.URL.Query().Get("id")
	if r.Method == http.MethodGet {
		var organizers interface{}
		if queryID == "" {
			query := r.URL.Query().Get("q")
			if query == "" {
				http.Error(w, "Must supply query value", http.StatusBadRequest)
				return
			}
			q, err := strconv.Atoi(query)
			if err != nil {
				http.Error(w, "Must supply a valid query", http.StatusBadRequest)
				return
			}
			organizers, err = ctx.UserStore.GetOrganizers(tid, q)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			oid, err := strconv.Atoi(queryID)
			if err != nil {
				http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
				return
			}
			organizers, err = ctx.UserStore.GetOrganizer(oid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if organizers == nil {
				http.Error(w, "Found no user with given user id", http.StatusNotFound)
				return
			}
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(organizers); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPost {
		// User must be a TO to access
		// *TODO*
		// Check is user is TO
		//
		oid, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		err = ctx.UserStore.RegisterOrganizer(oid, tid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("TO registered"))
	} else if r.Method == http.MethodDelete {
		// User must be a TO to access
		// *TODO*
		// Check is user is TO
		//
		if queryID == "" {
			http.Error(w, "Must supply user ID", http.StatusBadRequest)
			return
		}
		oid, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		err = ctx.UserStore.RemoveOrganizer(oid, tid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("TO removed"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// GamesHandler handles requests for the '/smashqq/tournaments/{tournamentID}/games' resource
func (ctx *TournamentContext) GamesHandler(w http.ResponseWriter, r *http.Request) {
	// *TODO*
	// Check if authenticated
	//
	urlVar := path.Base(path.Dir(r.URL.String()))
	tid, err := strconv.Atoi(urlVar)
	if err != nil {
		http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodGet {
		var games interface{}
		queryID := r.URL.Query().Get("id")
		if queryID != "" {
			gid, err := strconv.Atoi(queryID)
			if err != nil {
				http.Error(w, "Must supply a valid game ID", http.StatusBadRequest)
				return
			}
			games, err = ctx.UserStore.GetGame(gid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			query := r.URL.Query().Get("q")
			if query == "" {
				http.Error(w, "Must supply query value", http.StatusBadRequest)
				return
			}
			q, err := strconv.Atoi(query)
			if err != nil {
				http.Error(w, "Must supply a valid query", http.StatusBadRequest)
				return
			}
			games, err = ctx.Standings.GetGames(q, tid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(games); err != nil {
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
		game := new(models.Game)
		if err := json.NewDecoder(r.Body).Decode(game); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}
		err := ctx.UserStore.CreateGame(game, tid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Game Created"))
	} else if r.Method == http.MethodPatch {
		header := r.Header.Get("Content-Type")
		if !strings.HasPrefix(header, "application/json") {
			http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
			return
		}
		update := new(models.GamesUpdate)
		if err := json.NewDecoder(r.Body).Decode(update); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}
		game, err := ctx.UserStore.ReportGame(update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(games); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodDelete {
		queryID := r.URL.Query().Get("id")
		gid, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Must supply a valid game ID", http.StatusBadRequest)
			return
		}
		err = ctx.UserStore.DeleteGame(gid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Game deleted"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// StandingsHandler handles requests for the '/smashqq/tournaments/{tournamentID}/standings' resource
func (ctx *TournamentContext) StandingsHandler(w http.ResponseWriter, r *http.Request) {
	// *TODO*
	// Check if authenticated
	//
	if r.Method == http.MethodGet {
		var standings interface{}
		queryID := r.URL.Query().Get("id")
		if queryID != "" {
			uid, err := strconv.Atoi(queryID)
			if err != nil {
				http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
				return
			}
			standings, err = ctx.Standings.GetPlayer(uid)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			query := r.URL.Query().Get("q")
			if query == "" {
				http.Error(w, "Must supply query value", http.StatusBadRequest)
				return
			}
			q, err := strconv.Atoi(query)
			if err != nil {
				http.Error(w, "Must supply a valid query", http.StatusBadRequest)
				return
			}
			standings, err = ctx.Standings.Get(q)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(standings); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
	} else if r.Method == http.MethodPatch {
		header := r.Header.Get("Content-Type")
		if !strings.HasPrefix(header, "application/json") {
			http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
			return
		}
		update := new(models.StandingUpdate)
		if err := json.NewDecoder(r.Body).Decode(update); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}
		err := ctx.Standings.Update(update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Standings updated"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"info441finalproject/server/gateway/models"
)

//*TODO* change mysql game methods to include next game
// mysql get games should only return games with both players not nil

// GetTournamentIDFromURL retrieves the tournament id variable
// from the url. Variable must be at base of url
func GetTournamentIDFromURL(url string) (int, error) {
	urlVar := path.Base(url)
	tid, err := strconv.Atoi(urlVar)
	if err != nil {
		return 0, err
	}
	return tid, nil
}

// CreateGames creates all initial games for start of tournament once
// Registration is closed
func (ctx *TournamentContext) CreateGames(tid int64) error {
	// only id
	players, err := ctx.UserStore.GetAllPlayers(tid)
	if err != nil {
		return err
	}
	var game *models.Game
	gameStarted := false
	playerOne := true
	var roundOneGames []*models.Game
	for _, player := range players {
		if playerOne == true {
			game = new(models.Game)
			gameStarted = true
			game.TournamentID = tid
			game.PlayerOne = player.ID
			organizer, err := ctx.UserStore.GetLeastBusyTO(tid)
			if err != nil {
				return err
			}
			game.TournamentOrganizerID = organizer.ID
			game.InProgress = false
			game.Completed = false
			playerOne = false
		} else {
			game.PlayerTwo = player.ID
			_, err := ctx.UserStore.CreateGame(tid, game)
			if err != nil {
				return err
			}
			roundOneGames = append(roundOneGames, game)
			gameStarted = false
			playerOne = true
		}
	}
	if gameStarted == true {
		game.Victor = game.PlayerOne
		game.Completed = true
		game.Result = "Player one granted bye"
		_, err := ctx.UserStore.CreateGame(tid, game)
		if err != nil {
			return err
		}
		roundOneGames = append(roundOneGames, game)
	}

	err = ctx.CreateBracket(tid, roundOneGames)
	if err != nil {
		return err
	}
	return nil
}

// CreateBracket makes the rest of the bracket given an array of round1 games
func (ctx *TournamentContext) CreateBracket(tid int64, games []*models.Game) error {
	if len(games) > 1 {
		var nextGame *models.Game
		var nextRoundGames []*models.Game
		for i := range games {
			if i%2 == 0 {
				nextGame = new(models.Game)
				nextGame.TournamentID = tid
				organizer, err := ctx.UserStore.GetLeastBusyTO(tid)
				if err != nil {
					return err
				}
				nextGame.TournamentOrganizerID = organizer.ID
				nextGame.InProgress = false
				nextGame.Completed = false
				nextGame, err = ctx.UserStore.CreateGame(tid, nextGame)
				if err != nil {
					return err
				}
			} else {
				nextRoundGames = append(nextRoundGames, nextGame)
			}
			games[i].NextGame = nextGame.ID
		}

		// if odd number of gaes
		if len(games)%2 == 1 {
			nextRoundGames = append(nextRoundGames, nextGame)
		}
		return ctx.CreateBracket(tid, nextRoundGames)
	}
	return nil
}

//GetUserFromHeader returns the user object from "X-User" header
func GetUserFromHeader(r *http.Request) (*models.User, error) {
	xUser := r.Header.Get("X-User")
	if len(xUser) == 0 {
		return nil, fmt.Errorf("No User found")
	}
	user := new(models.User)
	err := json.Unmarshal([]byte(xUser), user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// TourneyHandler handles requests for the '/smashqq/tournaments' resource
// and '/smashqq/tournaments/{tournamentID}' resource
func (ctx *TournamentContext) TourneyHandler(w http.ResponseWriter, r *http.Request) {
	//Check if authenticated
	_, err := GetUserFromHeader(r)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(),
			http.StatusUnauthorized)
		return
	}
	if r.Method != http.MethodPost {
		if path.Base(r.URL.String()) == "tournaments" && r.Method == http.MethodGet {
			tournaments, err := ctx.UserStore.GetAllTournaments()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(tournaments); err != nil {
				http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
				return
			}
			return
		}
		tid, err := GetTournamentIDFromURL(r.URL.String())
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		if r.Method == http.MethodGet {
			tournament, err := ctx.UserStore.GetTournament(int64(tid))
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
			err = ctx.UserStore.DeleteTournament(int64(tid))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("tournament deleted"))
		} else if r.Method == http.MethodPatch {
			header := r.Header.Get("Content-Type")
			if !strings.HasPrefix(header, "application/json") {
				http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
				return
			}
			update := new(models.TournamentUpdate)
			if err := json.NewDecoder(r.Body).Decode(update); err != nil {
				http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
					http.StatusBadRequest)
				return
			}
			tournament, err := ctx.UserStore.UpdateTournament(int64(tid), update)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if tournament.Open == false {
				err := ctx.CreateGames(tournament.ID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(tournament); err != nil {
				http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
					http.StatusInternalServerError)
				return
			}
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
		returnTournament, err := ctx.UserStore.CreateTournament(tournament, int64(0)) // *TODO change to user id later
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
	//Check if authenticated
	_, err := GetUserFromHeader(r)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(),
			http.StatusUnauthorized)
		return
	}
	tid, err := GetTournamentIDFromURL(path.Dir(r.URL.String()))
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
			players, err = ctx.UserStore.GetPlayers(q, int64(tid))
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
			players, err = ctx.UserStore.GetByID(int64(uid))
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
		err = ctx.UserStore.RegisterPlayer(int64(uid), int64(tid))
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
		err = ctx.UserStore.RemovePlayer(int64(uid), int64(tid))
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
	//Check if authenticated
	xUser, err := GetUserFromHeader(r)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(),
			http.StatusUnauthorized)
		return
	}
	tid, err := GetTournamentIDFromURL(path.Dir(r.URL.String()))
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
			organizers, err = ctx.UserStore.GetTOs(q, int64(tid))
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
			organizers, err = ctx.UserStore.GetTO(int64(oid), int64(tid))
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
		if ctx.UserStore.UserIsTO(xUser.ID, int64(tid)) != true {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		oid, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		err = ctx.UserStore.RegisterTO(int64(oid), int64(tid))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("TO registered"))
	} else if r.Method == http.MethodDelete {
		if ctx.UserStore.UserIsTO(xUser.ID, int64(tid)) != true {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if queryID == "" {
			http.Error(w, "Must supply user ID", http.StatusBadRequest)
			return
		}
		oid, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
			return
		}
		err = ctx.UserStore.RemoveTO(int64(oid), int64(tid))
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
	//Check if authenticated
	_, err := GetUserFromHeader(r)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(),
			http.StatusUnauthorized)
		return
	}
	tid, err := GetTournamentIDFromURL(path.Dir(r.URL.String()))
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
			games, err = ctx.UserStore.GetGame(int64(gid))
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
			games, err = ctx.UserStore.GetGames(q, int64(tid))
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
		// } else if r.Method == http.MethodPost {
		// 	header := r.Header.Get("Content-Type")
		// 	if !strings.HasPrefix(header, "application/json") {
		// 		http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
		// 		return
		// 	}
		// 	game := new(models.Game)
		// 	if err := json.NewDecoder(r.Body).Decode(game); err != nil {
		// 		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
		// 			http.StatusBadRequest)
		// 		return
		// 	}
		// 	err := ctx.UserStore.CreateGame(game, tid)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusCreated)
		// 	w.Write([]byte("Game Created"))
	} else if r.Method == http.MethodPatch {
		header := r.Header.Get("Content-Type")
		if !strings.HasPrefix(header, "application/json") {
			http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
			return
		}
		update := new(models.GameUpdate)
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
		// If game is finished and game is not grands
		if update.Completed == true && game.NextGame != 0 {
			nextGameUpdate := new(models.GameUpdate)
			nextGameUpdate.ID = game.NextGame
			nextGame, err := ctx.UserStore.GetGame(game.NextGame)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if nextGame.PlayerOne == 0 {
				nextGameUpdate.PlayerOne = game.Victor
			} else if nextGame.PlayerTwo == 0 {
				nextGameUpdate.PlayerTwo = game.Victor
			}
			_, err = ctx.UserStore.ReportGame(update)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(game); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}
		// } else if r.Method == http.MethodDelete {
		// 	queryID := r.URL.Query().Get("id")
		// 	gid, err := strconv.Atoi(queryID)
		// 	if err != nil {
		// 		http.Error(w, "Must supply a valid game ID", http.StatusBadRequest)
		// 		return
		// 	}
		// 	err = ctx.UserStore.DeleteGame(gid)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte("Game deleted"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// StandingsHandler handles requests for the '/smashqq/tournaments/{tournamentID}/standings' resource
func (ctx *TournamentContext) StandingsHandler(w http.ResponseWriter, r *http.Request) {
	//Check if authenticated
	_, err := GetUserFromHeader(r)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(),
			http.StatusUnauthorized)
		return
	}
	tid, err := GetTournamentIDFromURL(path.Dir(r.URL.String()))
	if err != nil {
		http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodGet {
		var standings interface{}
		queryID := r.URL.Query().Get("id")
		if queryID != "" {
			uid, err := strconv.Atoi(queryID)
			if err != nil {
				http.Error(w, "Must supply a valid ID", http.StatusBadRequest)
				return
			}
			standings, err = ctx.UserStore.GetStanding(int64(uid), int64(tid))
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
			standings, err = ctx.UserStore.GetStandings(q, int64(tid))
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
		// } else if r.Method == http.MethodPatch {
		// 	header := r.Header.Get("Content-Type")
		// 	if !strings.HasPrefix(header, "application/json") {
		// 		http.Error(w, "Request body must in JSON", http.StatusUnsupportedMediaType)
		// 		return
		// 	}
		// 	update := new(models.StandingUpdate)
		// 	if err := json.NewDecoder(r.Body).Decode(update); err != nil {
		// 		http.Error(w, fmt.Sprintf("error decoding JSON: %v", err),
		// 			http.StatusBadRequest)
		// 		return
		// 	}
		// 	err := ctx.UserStore.Update(update)
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	w.Write([]byte("Standings updated"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

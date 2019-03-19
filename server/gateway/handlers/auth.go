package handlers

import (
	"encoding/json"
	"info441finalproject/server/gateway/indexes"
	"info441finalproject/server/gateway/models"
	"info441finalproject/server/gateway/sessions"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"
)

//UsersHandler defines HTTP handler functions as described in the
//assignment description. Remember to use your handler context
//struct as the receiver on these functions so that you have
//access to things like the session store and user store.
func (ctx *HandlerContext) UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.HasPrefix(r.Header.Get(contentType), applicationJSON) {
			http.Error(w, "request body must be json", http.StatusUnsupportedMediaType)
			return
		}
		// decode request body JSON to a new user struct
		// newUser --> user; store it into ctx.UserStore
		nu := new(models.NewUser)
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(nu)
		if err != nil {
			http.Error(w, "unable to decode response body to json", http.StatusBadRequest)
			return
		}
		u, err := nu.ToUser()
		if err != nil {
			http.Error(w, "cannot convert new user to normal user", http.StatusBadRequest)
			return
		}
		_, err = ctx.UserStore.GetByEmail(u.Email)
		if err == nil {
			http.Error(w, "user already exists", http.StatusBadRequest)
			return
		}
		u, err = ctx.UserStore.Insert(u)
		if err != nil {
			http.Error(w, "cannot insert new user", http.StatusBadRequest)
			return
		}

		addToTrie(ctx.SearchTrie, u.UserName, u.ID)
		addToTrie(ctx.SearchTrie, u.FirstName, u.ID)
		addToTrie(ctx.SearchTrie, u.LastName, u.ID)

		// create a new sessionState and begin a new session
		state := SessionState{
			StartTime: time.Now(),
			User:      u,
		}
		_, err = sessions.BeginSession(ctx.SigningKey, ctx.SessStore, state, w)
		if err != nil {
			http.Error(w, "session did not begin: "+err.Error(), http.StatusBadRequest)
			return
		}

		// set successful status code and content-type response header
		// encode new user profile to the response body
		encode(w, http.StatusCreated, u)
	} else if r.Method == "GET" {
		state := &SessionState{}
		sessionID, err := sessions.GetState(r, ctx.SigningKey, ctx.SessStore, state)
		if sessionID == sessions.InvalidSessionID || err != nil {
			http.Error(w, "current user is not authenticated: "+err.Error(), http.StatusUnauthorized)
			return
		}
		qValue := r.URL.Query().Get("q")
		if len(qValue) == 0 {
			http.Error(w, "no query string parameter given for q", http.StatusBadRequest)
			return
		}
		ids := ctx.SearchTrie.ReturnPrefixMatches(20, qValue)
		users := []models.User{}
		for _, id := range ids {
			u, err := ctx.UserStore.GetByID(id)
			if err != nil {
				http.Error(w, "error getting user from store", http.StatusBadRequest)
				return
			}
			users = append(users, *u)
		}
		sort.Slice(users, func(i, j int) bool {
			return strings.Compare(users[i].UserName, users[j].UserName) > 0
		})

		w.Header().Add(contentType, applicationJSON)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, "error encoding users to json", http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}
}

func addToTrie(t *indexes.Trie, param string, id int64) {
	splitParam := strings.Split(param, " ")
	for _, part := range splitParam {
		t.Add(strings.ToLower(part), int64(id))
	}
}

func deleteFromTrie(t *indexes.Trie, param string, id int64) {
	splitParam := strings.Split(param, " ")
	for _, part := range splitParam {
		t.Delete(strings.ToLower(part), int64(id))
	}
}

// SpecificUserHandler defines HTTP handler functions for one specific user given by
// a userID or "me" referring to the currently authenticated user
func (ctx *HandlerContext) SpecificUserHandler(w http.ResponseWriter, r *http.Request) {
	// how to check if user is authenticated?
	userID := path.Base(r.URL.Path)

	// authenticate the current user
	state := &SessionState{}
	sessionID, err := sessions.GetState(r, ctx.SigningKey, ctx.SessStore, state)
	if sessionID == sessions.InvalidSessionID || err != nil {
		http.Error(w, "current user is not authenticated: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// get user from the URL: "me" or a given ID
	u := getUserFromURL(userID, w, r, ctx)

	if r.Method == "GET" {
		encode(w, http.StatusOK, u)
	} else if r.Method == "PATCH" {
		// convert URL id to int64 and compare with the user id
		i64id, _ := strconv.ParseInt(userID, 10, 64)
		if u.ID != i64id {
			http.Error(w, "invalid user", http.StatusForbidden)
			return
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "invalid content type: request body must be in JSON", http.StatusUnsupportedMediaType)
			return
		}
		oldFirstname := u.FirstName
		oldLastname := u.LastName
		updates := new(models.Updates)
		err := json.NewDecoder(r.Body).Decode(updates)
		if err != nil {
			http.Error(w, "error decoding JSON from the response body", http.StatusBadRequest)
			return
		}
		err = u.ApplyUpdates(updates)
		if err != nil {
			http.Error(w, "error applying updates to user", http.StatusBadRequest)
			return
		}
		encode(w, http.StatusOK, u)

		deleteFromTrie(ctx.SearchTrie, oldFirstname, u.ID)
		deleteFromTrie(ctx.SearchTrie, oldLastname, u.ID)
		addToTrie(ctx.SearchTrie, u.FirstName, u.ID)
		addToTrie(ctx.SearchTrie, u.LastName, u.ID)
	} else {
		http.Error(w, "unsupported HTTP method type", http.StatusMethodNotAllowed)
		return
	}
}

func getUserFromURL(userID string, w http.ResponseWriter, r *http.Request, ctx *HandlerContext) *models.User {
	var result *models.User
	if userID == "me" {
		state := &SessionState{}
		sessions.GetState(r, ctx.SigningKey, ctx.SessStore, state)
		result = state.User
	} else {
		i64id, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			http.Error(w, "error converting userID to int64", http.StatusBadRequest)
			return nil
		}
		u, err := ctx.UserStore.GetByID(i64id)
		if err != nil {
			http.Error(w, "no user found with given id", http.StatusNotFound)
			return nil
		}
		_, err = ctx.UserStore.Insert(u)
		if err != nil {
			http.Error(w, "cannot insert new user", http.StatusBadRequest)
			return nil
		}

		result = u
	}
	return result
}

func encode(w http.ResponseWriter, statusCode int, u *models.User) {
	w.Header().Set(contentType, applicationJSON)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(u)
}

// SessionsHandler handles requests for the "sessions" resource, and allows clients to begin a new
// session using an existing user's credentials
func (ctx *HandlerContext) SessionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if !strings.HasPrefix(r.Header.Get(contentType), applicationJSON) {
			http.Error(w, "request body must be in json", http.StatusUnsupportedMediaType)
			return
		}
		creds := &models.Credentials{}
		err := json.NewDecoder(r.Body).Decode(creds)
		if err != nil {
			http.Error(w, "error decoding json", http.StatusBadRequest)
			return
		}
		u, err := ctx.UserStore.GetByEmail(creds.Email)
		if err != nil {
			time.Sleep(10 * time.Second)
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		} else if err = u.Authenticate(creds.Password); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		sessState := &SessionState{
			StartTime: time.Now(),
			User:      u,
		}
		ctx.UserStore.LogUserSignIn(u.ID, time.Now(), r.RemoteAddr)
		sessions.BeginSession(ctx.SigningKey, ctx.SessStore, sessState, w)
		encode(w, http.StatusCreated, u)
	} else {
		http.Error(w, "unsupported method type", http.StatusMethodNotAllowed)
		return
	}

}

// SpecificSessionHandler handles requests related to a specific authenticated session
func (ctx *HandlerContext) SpecificSessionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		if path.Base(r.URL.Path) != "mine" {
			http.Error(w, "invalid url, must end with \"mine\"", http.StatusForbidden)
			return
		}
		_, err := sessions.EndSession(r, ctx.SigningKey, ctx.SessStore)
		if err != nil {
			http.Error(w, "unexpected error ending current session", http.StatusBadRequest)
			return
		}
		w.Write([]byte("signed out"))
	} else {
		http.Error(w, "unsupported HTTP method type", http.StatusMethodNotAllowed)
		return
	}
}

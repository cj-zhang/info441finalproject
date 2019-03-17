package handlers

import (
	"encoding/json"
	"info441finalproject/server/gateway/sessions"
	"math/rand"
	"net/http"
	"net/url"
)

// Director redirects a request
type Director func(r *http.Request)

// CustomDirector redirects traffic
func (ctx *HandlerContext) CustomDirector(targets []*url.URL) Director {
	return func(r *http.Request) {
		targ := targets[rand.Int()%len(targets)]

		state := &SessionState{}
		sessionID, err := sessions.GetState(r, ctx.SigningKey, ctx.SessStore, state)
		if sessionID == sessions.InvalidSessionID || err != nil {
			r.Header.Del("X-User")
			return
		}
		u := state.User
		userJSON, err := json.Marshal(u)
		if err != nil {
			return
		}

		r.Header.Del("X-User")
		r.Header.Add("X-User", string(userJSON))
		r.Host = targ.Host
		r.URL.Host = targ.Host
		r.URL.Scheme = targ.Scheme
	}
}

package handlers

import (
	"info441finalproject/server/gateway/indexes"
	"info441finalproject/server/gateway/models"
	"info441finalproject/server/gateway/sessions"
)

//HandlerContext defines a handler context struct that
//will be a receiver on any of your HTTP
//handler functions that need access to
//globals, such as the key used for signing
//and verifying SessionIDs, the session store
//and the user store
type HandlerContext struct {
	SigningKey  string         `json:"signingKey,omitempty"`
	SessStore   sessions.Store `json:"sessStore,omitempty"`
	UserStore   models.Store   `json:"userStore,omitempty"`
	SearchTrie  *indexes.Trie
	SocketStore *SocketStore
}

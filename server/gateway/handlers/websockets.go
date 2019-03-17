package handlers

import (
	"fmt"
	"info441finalproject/server/gateway/sessions"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// SocketStore simple store to store all the connections
type SocketStore struct {
	Connections map[int64]*websocket.Conn
	lock        sync.Mutex
}

// Control messages for websocket
const (
	TextMessage   = 1
	BinaryMessage = 2
	CloseMessage  = 8
	PingMessage   = 9
	PongMessage   = 10
)

// InsertConnection is a thread-safe method for inserting a connection
func (s *SocketStore) InsertConnection(conn *websocket.Conn, userID int64) int {
	s.lock.Lock()
	connID := len(s.Connections)
	// insert socket connection
	s.Connections[userID] = conn
	s.lock.Unlock()
	return connID
}

// RemoveConnection is a thread-safe method for deleting a connection
func (s *SocketStore) RemoveConnection(userID int64) {
	s.lock.Lock()
	// delete socket connection
	delete(s.Connections, userID)
	s.lock.Unlock()
}

// WriteToAllConnections is a simple method for writing a message to all live connections.
// In your homework, you will be writing a message to a subset of connections
// (if the message is intended for a private channel), or to all of them (if the message
// is posted on a public channel
func (s *SocketStore) WriteToAllConnections(messageType int, data []byte) error {
	var writeError error
	for _, conn := range s.Connections {
		writeError = conn.WriteMessage(messageType, data)
		if writeError != nil {
			return writeError
		}
	}
	return nil
}

// WriteToConnection writes to connections for the given userIDs
func (s *SocketStore) WriteToConnection(userIDs []int64, messageType int, data []byte) error {
	var writeError error
	for _, id := range userIDs {
		if conn, ok := s.Connections[id]; ok {
			writeError = conn.WriteMessage(messageType, data)
			if writeError != nil {
				return writeError
			}
		}
	}
	return nil
}

// This is a struct to read our message into
type msg struct {
	Message string `json:"message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// This function's purpose is to reject websocket upgrade requests if the
		// origin of the websockete handshake request is coming from unknown domains.
		// This prevents some random domain from opening up a socket with your server.
		// TODO: make sure you modify this for your HW to check if r.Origin is your host
		return r.Header.Get("Origin") == r.Host
	},
}

// WebSocketConnectionHandler is a new handler thing
func (ctx *HandlerContext) WebSocketConnectionHandler(w http.ResponseWriter, r *http.Request) {
	// check if user is authenticated
	ss := &SessionState{}
	_, err := sessions.GetState(r, ctx.SigningKey, ctx.SessStore, ss)
	if err != nil {
		http.Error(w, "User is not authenticated", http.StatusUnauthorized)
		return
	}
	// handle the websocket handshake
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to open websocket connection", 401)
		return
	}

	// Insert our connection onto our datastructure for ongoing usage
	connID := ctx.SocketStore.InsertConnection(conn, ss.User.ID)

	// Invoke a goroutine for handling control messages from this connection
	go (func(conn *websocket.Conn, connId int) {
		defer conn.Close()
		defer ctx.SocketStore.RemoveConnection(ss.User.ID)

		for {
			messageType, p, err := conn.ReadMessage()

			if messageType == TextMessage || messageType == BinaryMessage {
				fmt.Printf("Client says %v\n", p)
				fmt.Printf("Writing %s to all sockets\n", string(p))
				ctx.SocketStore.WriteToAllConnections(TextMessage, append([]byte("Hello from server: "), p...))
			} else if messageType == CloseMessage {
				fmt.Println("Close message received.")
				break
			} else if err != nil {
				fmt.Println("Error reading message.")
				break
			}
			// ignore ping and pong messages
		}

	})(conn, connID)
}

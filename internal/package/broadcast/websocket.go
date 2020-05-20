package broadcast

import (
	"net/http"
    _ "fmt"
	"github.com/MuhtasimTanmoy/messaging_server/internal/package/logger"
	"github.com/gorilla/websocket"
)

// Message Object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Channel  string `json:"channel"`
}

// Websocket Object
type Websocket struct {
	Clients   map[string]map[*websocket.Conn]bool
	Broadcast chan Message
	Upgrader  websocket.Upgrader
}

// Websocket Init
func (e *Websocket) Init() {
	e.Clients = make(map[string]map[*websocket.Conn]bool)
	e.Broadcast = make(chan Message)
	e.Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

// Websocket HandleMessages
func (e *Websocket) HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-e.Broadcast
		// Send it out to every client that is currently connected
		for client := range e.Clients[msg.Channel] {
			err := client.WriteJSON(msg)
			if err != nil {
				logger.Infof("error: %v", err)
				client.Close()
				delete(e.Clients[msg.Channel], client)
			}
		}
	}
}

// Websocket HandleConnections
func (e *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request, channel string) {
	// Upgrade initial GET request to a websocket
	ws, err := e.Upgrader.Upgrade(w, r, nil)

	if err != nil {
		logger.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	e.Clients[channel][ws] = true

	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)

		if err != nil {
			logger.Infof("error: %v", err)
			delete(e.Clients[channel], ws)
			break
		}

		// Send the newly received message to the broadcast channel
		e.Broadcast <- msg
	}
}
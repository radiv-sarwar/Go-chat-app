package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Setting a variable to use to upgrade a http connection to a websocket conn
var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// This is the manager struct. Everytime a client connects, it gets assigned to a manager.
// This manager saves the client in its own Clients map.
type Manager struct {
	clients ClientList
	sync.RWMutex
}

// Method to return an empty manager struct (struct starts with a zero values) when initializing.
func NewManager() *Manager {
	return &Manager{
		clients: make(ClientList),
	}
}

// This is the main function. This will start everything by establishsing a websocket connection
// with the client.
func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {

	// Declaring the start of a new websocket connection.
	log.Println("Starting a new websocket connection")

	//upgrade to websocket
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// Adding the client to the manager's client list.
	client := NewClient(conn, m)
	m.addClient(client)

	// start two go routines per client
	// one is to read messages and one is to write messages
	go client.readMessages()

}

// Function to add clients to the clients list in our manager
func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	// Adding the client to the ClientList
	m.clients[client] = true
}

// Function to remove the client from the client list in our manager struct
func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
	}
}

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Setting a variable to use to upgrade a http
// connection to a websocket conn
var (
	websocketUpgrader = websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// This is the manager struct. Everytime a
// client connects, it gets assigned to a manager.
// This manager saves the client in its own Clients map.

// need a better name for this:
type handlers map[string]EventHandler

type Manager struct {
	clients  ClientList
	handlers handlers
	otps     RetentionMap
	sync.RWMutex
}

// Method to return an empty manager struct
// (struct starts with a zero values) when initializing.
func NewManager(ctx context.Context) *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
		otps:     NewRetentionMap(ctx, 5*time.Second),
	}
	m.setupEventHandlers()
	return m
}

func (m *Manager) setupEventHandlers() {
	m.handlers[EventSendMessage] = SendMessage
}

func SendMessage(event Event, c *Client) error {
	var chatevent SendMessageEvent

	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in the request: %v", err)
	}

	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	outgouingEvent := Event{
		Payload: data,
		Type:    EventNewMessage,
	}
	for client := range c.manager.clients {
		client.egress <- outgouingEvent
	}
	return nil
}

func (m *Manager) routeEvent(event Event, c *Client) error {
	if handler, ok := m.handlers[event.Type]; ok {
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
	}

}

// This is the main function. This will start everything by establishsing a websocket connection
// with the client.
func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {

	otp := r.URL.Query().Get("otp")
	if otp == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !m.otps.VerifyOTP(otp) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
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
	go client.writeMessages()

}

func (m *Manager) loginHandler(w http.ResponseWriter, r *http.Request) {
	type userLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req userLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// log.Println("here is the error.") // using this for debugging. will delete later
		return
	}

	// hardcoding my login because I won't be implementing
	// full on auth for now. Maybe later. Not the focus.
	if req.Username == "aaa" && req.Password == "111" {
		type response struct {
			OTP string `json:"otp"`
		}

		otp := m.otps.NewOTP()

		resp := response{
			OTP: otp.Key,
		}

		data, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return

	}

	w.WriteHeader(http.StatusUnauthorized)
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

func checkOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")

	switch origin {
	case "http://localhost:8080":
		return true
	default:
		return false
	}
}

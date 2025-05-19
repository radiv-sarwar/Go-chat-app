package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// The Client structure will take care of the client connecting to the server
type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

// A map of Clients that the manager is handling.
type ClientList map[*Client]bool

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}

func (c *Client) readMessages() {
	defer func() {
		// cleanup the connection
		c.manager.removeClient(c)
	}()

	for {
		messageType, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))

	}
}

package controllers

import (
	"chatsapi/internal/domain"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var connections []*websocket.Conn

func WebsocketControllers(router fiber.Router) {
	router.Get("/", initWebsocket)
}

func initWebsocket(c *fiber.Ctx) error {
	client := websocket.New(func(c *websocket.Conn) {
		// Store the new connection
		connections = append(connections, c)

		// Listen for new messages
		for {
			// Read message from the client
			_, msg, err := c.ReadMessage()
			if err != nil {
				// Remove the connection if an error occurs
				removeConnection(c)
				log.Printf("websocket error: %v", err)
				break
			}

			// Parse the message
			var message domain.LiveMessage
			err = json.Unmarshal(msg, &message)
			if err != nil {
				log.Printf("json parse error: %v", err)
				continue
			}

			// Broadcast the message to all connected clients
			broadcast(message)
		}
	})

	return client(c)
}

func removeConnection(conn *websocket.Conn) {
	for i, c := range connections {
		if c == conn {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func broadcast(message domain.LiveMessage) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("json encode error: %v", err)
		return
	}

	for _, conn := range connections {
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			// Remove the connection if an error occurs
			removeConnection(conn)
			log.Printf("websocket error: %v", err)
		}
	}
}

package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var connections []*websocket.Conn

func main() {
	app := fiber.New()

	app.Get("/messages/:id", getMessage)
	app.Post("/messages", createMessage)
	app.Delete("/messages/:id", deleteMessage)

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
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
			var message Message
			err = json.Unmarshal(msg, &message)
			if err != nil {
				log.Printf("json parse error: %v", err)
				continue
			}

			// Broadcast the message to all connected clients
			broadcast(message)
		}
	}))

	log.Fatal(app.Listen(":3000"))
}

func removeConnection(conn *websocket.Conn) {
	for i, c := range connections {
		if c == conn {
			connections = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func broadcast(message Message) {
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

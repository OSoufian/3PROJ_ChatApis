package controllers

import (
	"chatsapi/internal/domain"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type PartialLiveMessage struct {
	Message   string   `json:"message"`
	VideoId   uint     `json:"videoId"`
	Username  string   `gorm:"type:varchar(255);not null"`
}

func (p *PartialLiveMessage) Unmarshal(body []byte) error {
	return json.Unmarshal(body, &p)
}

var connections []*websocket.Conn

func WebsocketControllers(router fiber.Router) {
	router.Get("/", initWebsocket)
	router.Get("/connections", getNumConnections)
}

// Get All WebSockets
// @Summary LiveChat
// @Description init websocket broadcasr
// @Tags LiveChat
// @Success 200 {Message} LiveChat
// @Failure 404
// @Router /ws [get]
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

	// Return the number of connections as part of the response
	numConnections := len(connections)
	return c.JSON(fiber.Map{
		"connections": numConnections,
		"client": client(c),
	})
}

func getNumConnections(c *fiber.Ctx) error {
	numConnections := len(connections)
	return c.JSON(fiber.Map{"numConnections": numConnections})
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
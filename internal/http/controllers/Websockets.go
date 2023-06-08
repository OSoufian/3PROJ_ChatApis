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

			log.Println(message)
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

// package controllers

// import (
// 	"chatsapi/internal/domain"
// 	"encoding/json"
// 	"log"
// 	"strconv"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/websocket/v2"
// )

// type PartialLiveMessage struct {
// 	Message string `json:"message"`
// 	VideoId uint   `json:"videoId"`
// }

// var connections map[uint][]*websocket.Conn

// func (p *PartialLiveMessage) Unmarshal(body []byte) error {
// 	return json.Unmarshal(body, &p)
// }

// func WebsocketControllers(router fiber.Router) {
// 	router.Get("/", initWebsocket)
// }

// func initWebsocket(c *fiber.Ctx) error {
// 	client := websocket.New(func(conn *websocket.Conn) {
// 		// Store the new connection
// 		videoID := c.Query("videoId")
// 		videoIDUint, err := strconv.ParseUint(videoID, 10, 32)
// 		if err != nil {
// 			log.Printf("invalid video ID: %v", err)
// 			return
// 		}
// 		videoId := uint(videoIDUint)
// 		connections[videoId] = append(connections[videoId], conn)

// 		// Listen for new messages
// 		for {
// 			// Read message from the client
// 			_, msg, err := conn.ReadMessage()
// 			if err != nil {
// 				// Remove the connection if an error occurs
// 				removeConnection(conn, videoId)
// 				log.Printf("websocket error: %v", err)
// 				break
// 			}

// 			// Parse the message
// 			var message domain.LiveMessage
// 			err = json.Unmarshal(msg, &message)
// 			if err != nil {
// 				log.Printf("json parse error: %v", err)
// 				continue
// 			}

// 			message.VideoId = videoId

// 			log.Println(message)
// 			// Broadcast the message to all connected clients of the same video ID
// 			broadcast(message)
// 		}
// 	})

// 	return client(c)
// }

// func removeConnection(conn *websocket.Conn, videoID uint) {
// 	connList, exists := connections[videoID]
// 	if exists {
// 		for i, c := range connList {
// 			if c == conn {
// 				connections[videoID] = append(connList[:i], connList[i+1:]...)
// 				break
// 			}
// 		}
// 	}
// }

// func broadcast(message domain.LiveMessage) {
// 	data, err := json.Marshal(message)
// 	if err != nil {
// 		log.Printf("json encode error: %v", err)
// 		return
// 	}

// 	connList, exists := connections[message.VideoId]
// 	if exists {
// 		for _, conn := range connList {
// 			err = conn.WriteMessage(websocket.TextMessage, data)
// 			if err != nil {
// 				// Remove the connection if an error occurs
// 				removeConnection(conn, message.VideoId)
// 				log.Printf("websocket error: %v", err)
// 			}
// 		}
// 	}
// }
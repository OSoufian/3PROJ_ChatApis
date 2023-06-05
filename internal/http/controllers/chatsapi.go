package controllers

import (
	"encoding/json"
	"chatsapi/internal/domain"
	"time"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PartialMessage struct {
	Id      uint      `gorm:"primarykey;autoIncrement;not null"`
	Content string    `json:"Content"`
	VideoId uint      `gorm:"foreignKey:id"`
	UserId  uint      `gorm:"foreignKey:id"`
	Created string    `json:"created"`
}

func (p *PartialMessage) Unmarshal(body []byte) error {
	return json.Unmarshal(body, &p)
}

func ChatsApi(router fiber.Router) {
	router.Get("/", getAllMessages)
	router.Post("/", createMessage)
	router.Delete("/:MessageId", deleteMessage)
}

// Get All chats
// @Summary Chat
// @Description get all chats
// @Tags Chat
// @Success 200 {Message} List Chat
// @Failure 404
// @Router /chats/messages [get]
func getAllMessages(c *fiber.Ctx) error {

	videoParam := c.Query("q")
	videoID, err := strconv.Atoi(videoParam)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	messageModels := domain.Message{}
	message, err := messageModels.GetAll(videoID)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	return c.Status(200).JSON(message)
}

// Get All chats
// @Summary Chat
// @Description get all chats
// @Tags Chat
// @Success 200 {Message} List Chat
// @Failure 404
// @Router /chats/messages [post]
func createMessage(c *fiber.Ctx) error {
	videoObj := new(domain.Videos)
	userObj := new(domain.UserModel)
	
	message := new(domain.Message)
	
	partial := new(PartialMessage)

	if err := partial.Unmarshal(c.Body()); err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	videoObj.Id = uint(partial.VideoId)
	// video, err := videoObj.Get()
	// if err != nil {
	// 	return c.SendStatus(fiber.ErrBadGateway.Code)
	// }

	userObj.Id = uint(partial.UserId)
	user, err := userObj.Get()
	if err != nil {
		return c.SendStatus(fiber.ErrBadGateway.Code)
	}

	if partial.Content != "" {
		message.Content = partial.Content
	}

	message.UserId = user.Id
	// log.Println(user)
	if user.Icon != "" {
		message.User.Icon = user.Icon
	}
	if user.Username != "" {
		message.User.Username = user.Username
	}
	if user.Email != "" {
		message.User.Email = user.Email
	}
	message.User.Id = user.Id
	
	message.VideoId = videoObj.Id
	// message.Video = *video
	
	message.Created = time.Now().Format("2006-01-02 15:04:05")
	log.Println(*message)
	message.Create()

	// Return the new message
	return c.JSON(message)
}

// Get All chats
// @Summary Chat
// @Description get all chats
// @Tags Chat
// @Success 201
// @Failure 404
// @Router /chats/messages/:id [delete]
func deleteMessage(c *fiber.Ctx) error {

	messageParam := c.Params("MessageId")
	MessageID, err := strconv.Atoi(messageParam)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}
	message := domain.Message{}
	message.Id = MessageId

	if message.Find() {
		message.Delete()
	}

	return c.SendStatus(201)
}

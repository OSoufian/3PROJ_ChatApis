package controllers

import (
	"chatsapi/internal/domain"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

var messages = []domain.Message{
	{Id: 1, Content: "Hello, World!", Created: time.Now()},
	{Id: 2, Content: "Goodbye, World!", Created: time.Now()},
}

func ChatsApi(router fiber.Router) {
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"msg": "Hello World"})
	})
	router.Get("/:id", getMessage)
	router.Post("/", createMessage)
	router.Delete("/:id", deleteMessage)
}

// Get All chats
// @Summary Chat
// @Description get all chats
// @Tags Chat
// @Success 200 {Message} List Chat
// @Failure 404
// @Router /chats [get]
func getMessage(c *fiber.Ctx) error {
	// Get the ID from the URL parameters
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid message ID",
		})
	}

	// Find the message with the given ID
	var message *domain.Message
	for _, m := range messages {
		if m.Id == uint(id) {
			message = &m
			break
		}
	}

	// Return an error if the message was not found
	if message == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Message not found",
		})
	}

	// Return the message
	return c.JSON(message)
}

// Get All chats
// @Summary Chat
// @Description get all chats
// @Tags Chat
// @Success 200 {Message} List Chat
// @Failure 404
// @Router /chats [post]
func createMessage(c *fiber.Ctx) error {
	// Parse the request body into a new message
	var message domain.Message
	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Assign a new ID and creation time to the message
	message.Id = uint(len(messages) + 1)
	message.Created = time.Now()

	// Add the new message to the messages list
	messages = append(messages, message)

	// Return the new message
	return c.JSON(message)
}

// Get All chats
// @Summary Chat
// @Description get all chats
// @Tags Chat
// @Success 201
// @Failure 404
// @Router /chats/:id [delete]
func deleteMessage(c *fiber.Ctx) error {
	// Get the ID from the URL parameters
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid message ID",
		})
	}

	// Find the index of the message with the given ID
	index := -1
	for i, m := range messages {
		if m.Id == uint(id) {
			index = i
			break
		}
	}

	// Return an error if the message was not found
	if index == -1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Message not found",
		})
	}

	// Remove the message from the messages list
	messages = append(messages[:index], messages[index+1:]...)

	// Return a success message
	return c.JSON(fiber.Map{
		"message": "Message deleted",
	})
}

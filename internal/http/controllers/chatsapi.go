package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

var messages = []Message{
	{ID: 1, Content: "Hello, World!", Created: time.Now()},
	{ID: 2, Content: "Goodbye, World!", Created: time.Now()},
}

func ChatsApi(router fiber.Router) {
	router.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"msg": "Hello World"})
	})
	router.Get("/:id", getMessage)
	router.Post("/", createMessage)
	router.Delete("/:id", deleteMessage)
}

func getMessage(c *fiber.Ctx) error {
	// Get the ID from the URL parameters
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid message ID",
		})
	}

	// Find the message with the given ID
	var message *Message
	for _, m := range messages {
		if m.ID == id {
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

func createMessage(c *fiber.Ctx) error {
	// Parse the request body into a new message
	var message Message
	if err := c.BodyParser(&message); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Assign a new ID and creation time to the message
	message.ID = len(messages) + 1
	message.Created = time.Now()

	// Add the new message to the messages list
	messages = append(messages, message)

	// Return the new message
	return c.JSON(message)
}

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
		if m.ID == id {
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

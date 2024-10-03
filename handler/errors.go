package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (eh *ErrorHandler) HandleError(c *fiber.Ctx, err error) error {
	log.Println("Error occurred:", err)

	if fiberErr, ok := err.(*fiber.Error); ok {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"error": fiberErr.Message,
		})
	}

	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal Server Error",
	})
}

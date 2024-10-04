package main

import (
	"demo-go/errors"
	"demo-go/handler"
	"demo-go/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	bettingService := service.NewBettingService()
	errorHandler := errors.NewErrorHandler()
	bettingController := handler.NewBettingController(bettingService, errorHandler)

	api := app.Group("/api/betting")
	api.Get("/total-amount", bettingController.GetTotalAmount)
	api.Post("/place-bet", bettingController.PlaceBet)
	app.Listen(":3000")

}

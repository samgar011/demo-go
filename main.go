package main

import (
	"demo-go/controller"
	"demo-go/handler"
	"demo-go/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	bettingService := service.NewBettingService()
	errorHandler := handler.NewErrorHandler()
	bettingController := controller.NewBettingController(bettingService, errorHandler)

	api := app.Group("/api/betting")
	api.Get("/total-amount", bettingController.GetTotalAmount)
	api.Post("/place-bet", bettingController.PlaceBet)
	app.Listen(":3000")

}

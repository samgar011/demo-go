package handler

import (
	"demo-go/errors"
	"demo-go/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BettingController struct {
	bettingService *service.BettingService
	errorHandler   *errors.ErrorHandler
}

func NewBettingController(service *service.BettingService, errorHandler *errors.ErrorHandler) *BettingController {
	return &BettingController{
		bettingService: service,
		errorHandler:   errorHandler,
	}
}

func (bc *BettingController) PlaceBet(c *fiber.Ctx) error {
	var betRequest struct {
		GuessedNumber int64   `json:"guessedNumber"`
		BetAmount     float64 `json:"betAmount"`
	}

	if err := c.BodyParser(&betRequest); err != nil {
		return bc.errorHandler.HandleError(c, fiber.NewError(http.StatusBadRequest, "Invalid request body"))
	}

	if betRequest.GuessedNumber < 2 || betRequest.GuessedNumber > 98 {
		return bc.errorHandler.HandleError(c, fiber.NewError(http.StatusBadRequest, "Guessed number must be between 2 and 98"))
	}

	if betRequest.BetAmount < 1 || betRequest.BetAmount > 100 {
		return bc.errorHandler.HandleError(c, fiber.NewError(http.StatusBadRequest, "Bet amount must be between 1 and 100"))
	}

	exactNumber, err := bc.bettingService.GenerateExactNumber()
	if err != nil {
		return bc.errorHandler.HandleError(c, err)
	}

	winningPercentage := bc.bettingService.CalculateWinningPercentage(betRequest.GuessedNumber)
	payOut := bc.bettingService.CalculatePayOut(winningPercentage)

	resultMessage := "You lost! Better luck next time."
	won := false

	if exactNumber >= betRequest.GuessedNumber && exactNumber <= 98 {
		resultMessage = "You won!"
		won = true
	}

	bc.bettingService.UpdateTotalAmount(betRequest.BetAmount, payOut, won)

	currentRTP := bc.bettingService.CalculateAndSaveRTP()

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":     resultMessage,
		"exactNumber": exactNumber,
		"payOut":      payOut,
		"currentRTP":  currentRTP,
		"totalAmount": bc.bettingService.GetTotalAmount(),
		"rtpHistory":  bc.bettingService.GetRtpHistory(),
	})
}

func (bc *BettingController) GetTotalAmount(c *fiber.Ctx) error {
	return bc.errorHandler.HandleError(c, fiber.NewError(http.StatusNotImplemented, "Not implemented yet"))
}

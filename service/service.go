package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type BettingService struct {
	rtpHistory  []float64
	totalAmount float64
}

func NewBettingService() *BettingService {
	return &BettingService{
		rtpHistory:  make([]float64, 0),
		totalAmount: 5000,
	}
}

func (bs *BettingService) GenerateExactNumber() (int64, error) {
	return randomInt(2, 98)
}
func (bs *BettingService) CalculateWinningPercentage(guessedNumber int64) float64 {
	lowerBound := 2
	upperBound := 98

	if guessedNumber >= int64(lowerBound) && guessedNumber <= int64(upperBound) {
		return (float64(int64(upperBound)-guessedNumber) / float64(int64(upperBound)-int64(lowerBound))) * 100
	}

	return 0.0
}
func (bs *BettingService) CalculatePayOut(winningPercentage float64) float64 {
	if winningPercentage <= 0 {
		return 50.0
	}
	return 98.0 / winningPercentage
}
func (bs *BettingService) CalculateAndSaveRTP() float64 {
	var bet float64 = 0.0
	var win float64 = 0.0
	iterations := 1000000

	for i := 1; i <= iterations; i++ {
		bet += 1
		result, _ := randomInt(1, 100)
		choice, _ := randomInt(2, 98)

		if result > choice {
			win += multiplier(float64(choice))
		}
	}

	currentRTP := win / bet * 100
	bs.rtpHistory = append(bs.rtpHistory, currentRTP)

	return currentRTP
}

func (bs *BettingService) GetRtpHistory() []float64 {
	return bs.rtpHistory
}

func (bs *BettingService) UpdateTotalAmount(betAmount, payOut float64, won bool) {
	if won {
		bs.totalAmount += betAmount * payOut
	} else {
		bs.totalAmount -= betAmount
	}
}

func (bs *BettingService) GetTotalAmount() float64 {
	return bs.totalAmount
}

func randomInt(min, max int64) (int64, error) {

	if min > max {
		return 0, fmt.Errorf("invalid range [%d, %d]", min, max)
	}

	rangeSize := big.NewInt(max - min + 1)
	nBig, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return 0, err
	}

	return nBig.Int64() + min, nil
}

func multiplier(x float64) float64 {
	return 98 / (100 - x)
}

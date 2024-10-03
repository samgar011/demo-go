package model

type Bet struct {
	GuessedNumber int     `json:"guessedNumber"`
	BetAmount     float64 `json:"betAmount"`
	TotalAmount   float64 `json:"totalAmount"`
	PayOut        float64 `json:"payOut"`
	MaxBetAmount  float64 `json:"maxBetAmount"`
}

func NewBet() *Bet {
	return &Bet{
		TotalAmount:  5000,
		MaxBetAmount: 100,
	}
}

package gambler

import (
	"math"
)

//KellyInfo x
type KellyInfo struct {
	BetAmount          int     `json:"bet_amount,omitempty"`
	GrowthRate         float64 `json:"growth_rate,omitempty"`
	BankRollPercentage float64 `json:"bank_roll_percentage,omitempty"`
	IsMaxBet           bool    `json:"is_max_bet,omitempty"`
}

//KellyCriterion a payout of 1:1 would be 1, 3:2 would be 1.5.
//Returns the bet amount in wager units and expected growth rate
func KellyCriterion(bankRoll, minBet, betMultiple int, winProbability, payoutRatio, maximumWagerRatio, kellyRatio float64) KellyInfo {
	p := winProbability
	q := 1 - p
	b := payoutRatio
	bankRollPercentage := (b*p - q) / b
	if bankRollPercentage < 0 || math.IsInf(bankRollPercentage, 0) {
		return KellyInfo{}
	}
	bankRollPercentage /= kellyRatio

	betF := float64(bankRoll) * bankRollPercentage

	maxWager := int(math.Round(float64(bankRoll) * maximumWagerRatio))
	if maxWager < minBet && bankRollPercentage > 0 {
		betF = float64(minBet)
		bankRollPercentage = float64(minBet) / float64(bankRoll)
	} else if betF > float64(maxWager) {
		betF = float64(maxWager)
		bankRollPercentage = maximumWagerRatio
	}

	bm := float64(betMultiple)
	interval := int(betF / bm)
	bet := interval * betMultiple

	if bet < minBet {
		return KellyInfo{}
	}

	bankLeft := 1 - bankRollPercentage
	l1 := (1 - p)
	l2 := -math.Log(bankLeft)
	loss := l1 * l2

	g1 := p
	g2 := math.Log(1 + bankRollPercentage)
	gain := g1 + g2
	exp := math.Exp(gain - loss)
	growth := exp - 1

	ki := KellyInfo{
		BetAmount:          bet,
		GrowthRate:         growth,
		BankRollPercentage: bankRollPercentage,
		IsMaxBet:           bet == maxWager,
	}
	return ki
}

//CompoundGrowthRate x
func CompoundGrowthRate(growthRate float64, bankRoll, iterations int) int {
	c := float64(bankRoll) * math.Exp(growthRate*float64(iterations))
	return int(math.Round(c))
}

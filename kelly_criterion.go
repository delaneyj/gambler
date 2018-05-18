package gambler

import (
	"math"
)

//KellyCriterion a payout of 1:1 would be 1, 3:2 would be 1.5.
//Returns the bet amount in wager units and expected growth rate
func KellyCriterion(bankRoll, minBet, betMultiple int, winProbability, payoutRatio, maximumWagerRatio float64) (int, float64) {
	p := winProbability
	b := payoutRatio
	bankRollPercentage := (p*b + p - 1) / b
	if bankRollPercentage < 0 {
		return 0, 0
	}

	betF := float64(bankRoll) * bankRollPercentage

	bm := float64(betMultiple)
	interval := int(betF / bm)
	bet := interval * betMultiple

	if bet < minBet {
		return 0, 0
	}

	maxWager := int(math.Round(float64(bankRoll) * maximumWagerRatio))
	if bet > maxWager {
		bet = maxWager
		bankRollPercentage = maximumWagerRatio
	}

	bankLeft := 1 - bankRollPercentage
	l1 := (1 - p)
	l2 := -math.Log(bankLeft)
	loss := l1 * l2

	g1 := p
	g2 := math.Log(1 + bankRollPercentage)
	gain := g1 + g2
	growth := math.Exp(gain-loss) - 1

	return bet, growth / 100
}

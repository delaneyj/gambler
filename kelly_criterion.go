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

//KellyArgs x
type KellyArgs struct {
	Bankroll          int
	MinWagerAllowed   int
	MaxWagerAllowed   int
	WinProbability    float64
	PayoutRatio       float64
	MaximumWagerRatio float64
	KellyRatio        float64
}

//KellyCriterion a payout of 1:1 would be 1, 3:2 would be 1.5.
//Returns the bet amount in wager units and expected growth rate
func KellyCriterion(args KellyArgs) KellyInfo {
	p := args.WinProbability
	q := 1 - p
	b := args.PayoutRatio
	bankRollPercentage := (b*p - q) / b
	if bankRollPercentage < 0 || math.IsInf(bankRollPercentage, 0) {
		return KellyInfo{}
	}
	bankRollPercentage /= args.KellyRatio

	bankrollF := float64(args.Bankroll)
	betF := float64(args.Bankroll) * bankRollPercentage
	bet := int(math.Floor(betF))

	maxWager := int(math.Round(bankrollF * args.MaximumWagerRatio))
	if maxWager < args.MinWagerAllowed || bet < args.MinWagerAllowed {
		return KellyInfo{}
	}

	betF = float64(bet)
	bet = int(
		math.Min(
			math.Min(
				betF, float64(maxWager),
			),
			float64(args.MaxWagerAllowed),
		),
	)
	bankRollPercentage = betF / bankrollF

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
		BetAmount:          int(betF),
		GrowthRate:         growth,
		BankRollPercentage: bankRollPercentage,
		IsMaxBet:           bet == maxWager || bet == args.MaxWagerAllowed,
	}
	return ki
}

//CompoundGrowthRate x
func CompoundGrowthRate(growthRate float64, bankRoll, iterations int) int {
	c := float64(bankRoll) * math.Exp(growthRate*float64(iterations))
	return int(math.Round(c))
}

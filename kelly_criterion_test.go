package gambler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKellyCriterion(t *testing.T) {
	t.Parallel()

	type args struct {
		bankRoll       int
		minBet         int
		winProbability float64
		payoutRatio    float64
		maxWagerRatio  float64
	}
	tests := []struct {
		name               string
		args               args
		expected           int
		growthRate         float64
		bankRollPercentage float64
		isMaxBet           bool
	}{
		{
			name: "growthExample",
			args: args{
				payoutRatio:    2,
				bankRoll:       1000,
				winProbability: 0.6,
				minBet:         1,
				maxWagerRatio:  0.5,
			},
			expected:           399,
			growthRate:         1.0794294597813834,
			bankRollPercentage: 0.399,
			isMaxBet:           true,
		},
		{
			name: "basic",
			args: args{
				payoutRatio:    7.0 / 4.0,
				bankRoll:       100000,
				winProbability: 0.4,
				minBet:         1000,
				maxWagerRatio:  0.025,
			},
			expected:           5714,
			growthRate:         0.5223644213400322,
			bankRollPercentage: 0.05714,
			isMaxBet:           true,
		},
		{
			name: "bad",
			args: args{
				payoutRatio:    4.0 / 7.0,
				bankRoll:       100000,
				winProbability: 0.4,
				minBet:         1000,
				maxWagerRatio:  0.025,
			},
			expected:           0,
			growthRate:         0,
			bankRollPercentage: 0,
			isMaxBet:           false,
		},
		{
			name: "close",
			args: args{
				payoutRatio:    1.55,
				bankRoll:       123400,
				winProbability: 0.4,
				minBet:         1000,
				maxWagerRatio:  0.025,
			},
			expected:           1592,
			growthRate:         0.49934384784920227,
			bankRollPercentage: 0.012901134521880065,
			isMaxBet:           true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ki := KellyCriterion(KellyArgs{
				Bankroll:        tt.args.bankRoll,
				MinWagerAllowed: tt.args.minBet,
				WinProbability:  tt.args.winProbability,
				PayoutRatio:     tt.args.payoutRatio,
				MaxWagerRatio:   tt.args.maxWagerRatio,
				KellyRatio:      1,
			})
			assert.Equal(t, tt.expected, ki.BetAmount, tt.name)
			assert.Equal(t, tt.growthRate, ki.GrowthRate, tt.name)
			assert.Equal(t, tt.bankRollPercentage, ki.BankRollPercentage, tt.name)
			assert.Equal(t, tt.isMaxBet, ki.IsMaxBet, tt.name)
		})
	}
}

func TestCompoundGrowthRate(t *testing.T) {
	t.Parallel()

	//http://www.meta-financial.com/lessons/compound-interest/continuously-compounded-interest.php
	type args struct {
		growthRate float64
		bankRoll   int
		iterations int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "If you invest $1,000 at an annual interest rate of 5% compounded continuously, calculate the final amount you will have in the account after five years.",
			args: args{
				bankRoll:   100000,
				growthRate: 0.05,
				iterations: 5,
			},
			want: 128403,
		},
		{
			name: "If you invest $500 at an annual interest rate of 10% compounded continuously, calculate the final amount you will have in the account after five years.",
			args: args{
				bankRoll:   50000,
				growthRate: 0.1,
				iterations: 5,
			},
			want: 82436,
		},
		{
			name: "If you invest $2,000 at an annual interest rate of 13% compounded continuously, calculate the final amount you will have in the account after 20 years.",
			args: args{
				bankRoll:   200000,
				growthRate: 0.13,
				iterations: 20,
			},
			want: 2692748,
		},
		{
			name: "If you invest $20,000 at an annual interest rate of 1% compounded continuously, calculate the final amount you will have in the account after 20 years.",
			args: args{
				bankRoll:   2000000,
				growthRate: 0.01,
				iterations: 20,
			},
			want: 2442806,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CompoundGrowthRate(tt.args.growthRate, tt.args.bankRoll, tt.args.iterations)
			assert.Equal(t, tt.want, actual, tt.name)
		})
	}
}

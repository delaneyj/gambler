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
		betMultiple    int
		winProbability float64
		payoutRatio    float64
		maxWagerRatio  float64
	}
	tests := []struct {
		name       string
		args       args
		expected   int
		growthRate float64
	}{
		{
			name: "growthExample",
			args: args{
				payoutRatio:    2,
				bankRoll:       1000,
				winProbability: 0.6,
				betMultiple:    1,
				minBet:         1,
				maxWagerRatio:  0.5,
			},
			expected:   399,
			growthRate: 1.07953016734645361,
		},
		{
			name: "basic",
			args: args{
				payoutRatio:    7.0 / 4.0,
				bankRoll:       100000,
				winProbability: 0.4,
				betMultiple:    100,
				minBet:         1000,
				maxWagerRatio:  0.025,
			},
			expected:   2500,
			growthRate: 0.5060674679348933,
		},
		{
			name: "bad",
			args: args{
				payoutRatio:    4.0 / 7.0,
				bankRoll:       100000,
				winProbability: 0.4,
				betMultiple:    100,
				minBet:         1000,
				maxWagerRatio:  0.025,
			},
			expected:   0,
			growthRate: 0,
		},
		{
			name: "close",
			args: args{
				payoutRatio:    1.55,
				bankRoll:       123400,
				winProbability: 0.4,
				betMultiple:    100,
				minBet:         1000,
				maxWagerRatio:  0.025,
			},
			expected:   1500,
			growthRate: 0.4993450375407944,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, growthRate := KellyCriterion(
				tt.args.bankRoll, tt.args.minBet,
				tt.args.betMultiple, tt.args.winProbability,
				tt.args.payoutRatio, tt.args.maxWagerRatio)
			assert.Equal(t, tt.expected, actual, tt.name)
			assert.Equal(t, tt.growthRate, growthRate, tt.name)
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

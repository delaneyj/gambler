package gambler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKellyCriterion(t *testing.T) {
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
			growthRate: 0.0107953016734645361,
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
			growthRate: 0.005060674679348933,
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
			growthRate: 0.004993450375407944,
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

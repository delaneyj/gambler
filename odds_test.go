package gambler

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmerica_Odds(t *testing.T) {
	espilon := 0.001

	a := -110
	d := 1.91
	f := *big.NewRat(10, 11)
	ff, _ := f.Float64()

	assert.InDelta(t, d, AmericanToDecimal(a), espilon, "a2d")
	assert.Equal(t, f, AmericanToFractional(a), "a2f")

	assert.Equal(t, a, DecimalToAmerican(d), "d2a")
	d2f := DecimalToFractional(d)
	d2ff, _ := d2f.Float64()
	assert.InDelta(t, ff, d2ff, espilon, "d2f")

	assert.Equal(t, a, FractionalToAmerican(f), "f2a")
	assert.InDelta(t, d, FractionalToDecimal(f), espilon, "f2d")
}

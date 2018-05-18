package gambler

import (
	"math"
	"math/big"
)

//AmericanToDecimal x
func AmericanToDecimal(a int) float64 {
	fa := float64(a)
	if a > 0 {
		return (fa / 100) + 1
	}

	return (100 / -fa) + 1
}

//AmericanToFractional x
func AmericanToFractional(a int) big.Rat {
	f := big.Rat{}
	ai := big.NewInt(int64(a))
	i := big.NewInt(100)
	if a > 0 {
		f.SetFrac(ai, i)
	} else {
		f.SetFrac(i, ai.Neg(ai))
	}

	return f
}

//DecimalToAmerican x
func DecimalToAmerican(d float64) int {
	if d >= 2 {
		return int(math.Round((d - 1) * 100))
	}
	return int(math.Round(-100 / (d - 1)))
}

//DecimalToFractional x
func DecimalToFractional(d float64) big.Rat {
	r := big.Rat{}
	r.SetFloat64(float64(d) - 1)
	return r
}

//FractionalToDecimal x
func FractionalToDecimal(f big.Rat) float64 {
	x, _ := f.Float64()
	return x + 1
}

//FractionalToAmerican x
func FractionalToAmerican(f big.Rat) int {
	x, _ := f.Float64()
	if x > 1 {
		return int(x * 100)
	}
	return int(-100 / x)
}

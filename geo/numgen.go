package geo

import (
	"math"
	"math/rand"
)

// NumGen (Number Generator) is a function that returns a number.
type NumGen func() float64

// ConstNum returns a NumGen that always returns n.
func ConstNum(n float64) NumGen {
	return func() float64 {
		return n
	}
}

// RandNum returns a NumGen that returns a uniform random number between min and max.
func RandNum(min, max float64) NumGen {
	width := max - min
	return func() float64 {
		return rand.Float64()*width + min
	}
}

// RandRadius returns a NumGen that returns a uniform circle radius between minR and maxR.
func RandRadius(minR, maxR float64) NumGen {
	if maxR == 0 || maxR == minR {
		return func() float64 {
			return maxR
		}
	}
	unitMin := minR / maxR
	unitMin *= unitMin
	return func() float64 {
		return math.Sqrt(rand.Float64()*(1-unitMin)+unitMin) * maxR
	}
}

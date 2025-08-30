package main

import (
	"math"
	"math/rand"
)

func calculatePi() float64 {
	x := rand.Float64()
	y := rand.Float64()
	x = x/math.MaxFloat64*2.0 + 1.0
	y = y/math.MaxFloat64*2.0 + 1.0

	return 0.0
}

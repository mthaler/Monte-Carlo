package main

import (
	"math/rand"
)

func calculatePi(n int) float64 {
	circle := 0.0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y < 1.0 {
			circle += 1.0
		}
	}

	return 4.0 * circle / float64(n)
}

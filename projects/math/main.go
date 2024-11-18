package main

import (
	"math"
)

var k, p, v float64 = 1296, 6, 6

func M() float64 {
	return p * v
}

func W() float64 {
	m := M()
	return math.Sqrt(k / m)
}

func T() float64 {
	w := W()
	return 6 / w
}

func main() {
	T()
}

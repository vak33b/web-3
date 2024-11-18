package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a int
		b int
	)
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(math.Sqrt(float64(a*a + b*b)))
}

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	// implementation for custom square root
	var z float64 = 1
	var i int = 0
	for i < 10 {
		z = z - (z*z-x)/(2*z)
		if math.Abs(z*z-x) < 1e-9 {
			return z
		}
		fmt.Printf("z = %v\n", z)
		i += 1
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
}

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z float64 = 1
	var i int = 1
	for i < 20 {
		z = z - (z*z-x)/(2*z)
		if (z*z - x) < 1e-15 {
			return z
		}
		fmt.Printf("%v\n", z)
		i += 1
	}
	return z
}

func main() {
	var x int
	for {
		fmt.Printf("> ")
		fmt.Scan(&x)
		fmt.Println(Sqrt(float64(x)))
	}
}

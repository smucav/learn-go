package main

import (
	"fmt"
	"math"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

const (
	big   = 1 << 100
	small = big >> 99
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

var i, j int = 1, 2

const PI = 3.14

func main() {
	for {
		var input int
		fmt.Printf("Enter a number: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Enter valid number")
			continue
		}
		fmt.Println(split(input))
	}
	x := "hello"
	y := "world"
	fmt.Printf("first x = %s and y = %s\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("after swap x = %s and y = %s\n", x, y)
	fmt.Println(i, j)
	x := 2 << 2
	fmt.Println(x)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	v := 42
	fmt.Printf("the type of v is %T\n", v)
	var testme bool
	fmt.Printf("the type of testme variable is %T\n", testme)
	fmt.Printf("and it's value is %v\n", testme)
	v = 10
	fmt.Println(v)
	fmt.Println("Hello, world!")
	fmt.Println("Happy", PI, "day")
	const really = true

	fmt.Println("Enjoing Go?", really)
	fmt.Println(needFloat(big))
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))

}

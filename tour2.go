package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Println(v)
		return lim
	}
}

func main() {
	// var sum int
	// for true {
	// 	sum++
	// 	if sum == 1000 {
	// 		break
	// 	}
	// }
	// fmt.Println(sum)
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	//	today := time.Now().Weekday()
	//t := time.Now()

	// switch {
	// case today + 0:
	// 	fmt.Printf("Today\n")
	// case today + 1:
	// 	fmt.Printf("Tomorrow\n")
	// case today - 1:
	// 	fmt.Printf("Yesterday\n")
	// default:
	// 	fmt.Printf("Too far away\n")
	// }
	// fmt.Printf("%v \n", t.Hour() )
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Printf("Good Morning")
	// case t.Hour() < 17:
	// 	fmt.Printf("Good afternoon")
	// default:
	// 	fmt.Printf("Good evening")
	// }
	// defer fmt.Printf("world")
	// fmt.Printf("hello")
	fmt.Printf("Counting\n")
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d\n", i)
	}
	fmt.Printf("Done\n")
}

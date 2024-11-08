package main

import (
	"fmt"
	"errors"

)

func add(x, y int) (int, error) {
	return x + y, nil
}

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

func sub(x, y int) (int, error) {
	if y > x {
		return 0, errors.New("The result will be negative")
	}
	return x - y, nil
}

type User struct {
	name string
	email string
	age int
}
func main() {

	var number int

	for {
		fmt.Printf("Enter an integer: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Ivalid input, Please enter a valid integer")
			continue
		}
		if number % 2 == 0 {
			fmt.Printf("%d is even.\n", number)
		} else {
			fmt.Printf("%d is odd.\n", number)
		}
	}
	ans, err := divide(1, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("the result 1, 2 is %f: ", ans)
	}

	var x int
	var y int


	fmt.Println("Enter the first number: ")
	_, firsterr := fmt.Scan(&x)

	fmt.Println("Enter the second number: ")
	_, seconderr := fmt.Scan(&y)

	if firsterr == nil && seconderr == nil {
		result, err := sub(x, y)

		if err != nil {
			fmt.Printf("Error %v\n", err)
		} else {
			fmt.Printf("%d - %d = %d\n", x, y, result)
		}
	}
	a := 1
	b := 2
	result, _ := add(a, b)
	fmt.Printf("%d + %d = %d", a, b, result)

	user1 := User{name: "daniel", email: "danielsevani12@gmail.com", age: 24}
	fmt.Printf("name: %s\nage: %d\nemail: %s\n", user1.name, user1.age, user1.email)


}

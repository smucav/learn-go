package main

import (
	"fmt"
	"strings"
)

// type Person struct {
// 	Name string
// 	Age  int
// }

// var (
// 	p1 = &Person{Name: "Daniel", Age: 22}
// 	p2 = &Person{Name: "Bob", Age: 20}
// 	p3 = &Person{Name: "Tujuma", Age: 40}
// )

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // X = 1, Y = 2
	v2 = Vertex{X: 1}  // X = 1, Y = 0
	p  = &Vertex{1, 2} // pointer to struct value
	v3 = Vertex{}      // X = 0, Y = 0
)

func main() {
	// my first pointer in go
	// var p *int
	// i := 10
	// p = &i
	// fmt.Printf("before\n")
	// fmt.Printf("i = %v\n", i)
	// *p = 11
	// fmt.Printf("after\n")
	// fmt.Printf("i = %v\n", i)
	// v := Vertex{1, 2}
	// v.X = 10
	// fmt.Println(v.X)
	// pointer to struct
	// v := Vertex{1, 2}
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)
	// var x [2]string
	// x[0] = "hello"
	// x[1] = "world"
	// or
	x := [2]string{"hello", "world"}
	var my_var [4]string

	my_var[0] = "first element"
	my_var[1] = "hello"

	primes := [6]int{2, 3, 5, 7, 11, 13}

	first_3 := primes[0:3]
	first_3[2] = 20

	fmt.Println(x)
	fmt.Println(primes)

	// var my_ptr = Person{Name: "FirstName", Age: 24}

	// ptr := &my_ptr

	// fmt.Println((*ptr).Name)

	fmt.Println(v1, v2, v3, *p)

	var name [3]string

	name[0] = "Abebe"
	name[1] = "Bekele"
	name[2] = "Chala"

	my_name := name[:2]

	// change the sliced value
	my_name[0] = "changed"
	fmt.Println(name)

	pass_student := []struct {
		name   string
		status bool
	}{
		{"daniel", true},
		{"yohannes", true},
		{"abebe", false},
	}
	fmt.Printf("%v\n", pass_student)

	// my_list := []string{"tomato", "potato", "garlic", "chills", "other"}

	// printSlice(my_list)

	// my_list = my_list[:0]

	// printSlice(my_list)

	// if len(my_list) == nil {
	// 	fmt.Printf("nill!\n")
	// }

	// my_list = my_list[:cap(my_list)]

	// printSlice(my_list)

	// my_list = my_list[2:]

	// printSlice(my_list)

	// s := make([]int, 3, 3)

	// printSlice("s", s)

	// a := make([]int, 5)

	// printSlice("a", a)

	// b := a[:2]

	// printSlice("b", b)

	//Making board

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][2] = "x"
	board[1][1] = "x"
	board[2][0] = "x"

	fmt.Printf("the length of board = %d\n", len(board))

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	for index, value := range board {
		fmt.Printf("index = %d, value = %v\n", index, strings.Join(value, " "))
	}

	powers := make([]int, 10)

	for i := range powers {
		powers[i] = 1 << i
	}
	for _, value := range powers {
		fmt.Printf("%d ", value)
	}
	fmt.Printf("\n")

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}

package main

import (
	"fmt"
)

type Vertex struct {
	x int
	y int
}

func Position(pos int) (int, int) {
	m := make(map[int]Vertex)
	m[1] = Vertex{0, 0}
	m[2] = Vertex{0, 1}
	m[3] = Vertex{0, 2}
	m[4] = Vertex{1, 0}
	m[5] = Vertex{1, 1}
	m[6] = Vertex{1, 2}
	m[7] = Vertex{2, 0}
	m[8] = Vertex{2, 1}
	m[9] = Vertex{2, 2}

	fmt.Println(len(m))

	return m[pos].x, m[pos].y
}

func main() {

	// v1 := Vertex{1, 2}
	// v2 := Vertex{1, 3}

	// if v1 == v2 {
	// 	fmt.Printf("equal\n")
	// } else {
	// 	fmt.Printf("not equal\n")
	// }
	int(x, y, z)
	fmt.Println(x, y, z)

	first, second := Position(7)
	fmt.Println(first, second)

}

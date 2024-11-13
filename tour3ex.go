package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for i := 0; i < dy; i++{
		img[i] = make([]uint8, dx)
		for j := 0; j < dx; j++{
			value := i * j
			img[i][j] = uint8(value)
		}
	}

	return img

	// for i := 0; i < dy; i++{
	// 	for j := 0; j < dx; j++{
	// 		img[i][j] = (img[i][j] * 255) / max_val
	// 	}
	// }
	// return img

}
// func Pic(dx, dy int) [][]uint8 {
//     p := make([][]uint8, dy)
//     for y := range p {
//         p[y] = make([]uint8, dx)
//         for x := range p[y] {
//         p[y][x] = uint8(x^y)
//         // p[y][x] = uint8(x*y)
//         // p[y][x] = uint8((x+y)/2)
//         }
//     }
//     return p
// }

func main() {

	pic.Show(Pic)
}

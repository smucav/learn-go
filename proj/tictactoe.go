// simple tic tac toe game
package main

import (
	"fmt"
	"strings"
)

func Win(board [][]string, sign string) int {
	// generate win board that define possible combination
	// of a winning positions and based on that return 1 or 0
	var i int = 0
	win_board := make([][]Vertex, 8)

	win_board[0] = []Vertex{{0, 0}, {0, 1}, {0, 2}}
	win_board[1] = []Vertex{{1, 0}, {1, 1}, {1, 2}}
	win_board[2] = []Vertex{{2, 0}, {2, 1}, {2, 2}}
	win_board[3] = []Vertex{{0, 0}, {1, 1}, {2, 2}}
	win_board[4] = []Vertex{{0, 2}, {1, 1}, {2, 0}}
	win_board[5] = []Vertex{{0, 0}, {1, 0}, {2, 0}}
	win_board[6] = []Vertex{{0, 1}, {1, 1}, {2, 1}}
	win_board[7] = []Vertex{{0, 2}, {1, 2}, {2, 2}}

	for i < len(win_board) {
		c1, c2, c3 := win_board[i][0], win_board[i][1], win_board[i][2]
		cell1 := board[c1.x][c1.y]
		cell2 := board[c2.x][c2.y]
		cell3 := board[c3.x][c3.y]
		if cell1 == sign && cell2 == sign && cell3 == sign {
			return 1
		}
		i++
	}

	return 0
}

func Position(pos int) (int, int) {
	// used to get the X and Y vertex of enterd position
	// map 1-10 to vertex on the tic tac toe board
	// return x and y vertex of the board
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

	return m[pos].x, m[pos].y
}

type Player struct {
	// player info
	name string
	sign string
	win  int
	lost int
	draw int
}

type Vertex struct {
	// vertex of the board
	x int
	y int
}

func PrintBoard(board [][]string) {
	// print the current status of the board
	fmt.Println("====================")
	for i := range board {
		for j := range board[i] {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println("====================")
	fmt.Println()

}

func Again(board [][]string, Player1 Player, Player2 Player) int {
	// ask the winner to play again
	var again int
	fmt.Printf("%s Win\n", Player1.name)
	fmt.Printf("======================\n")
	fmt.Printf("------------------------------------------------\n")
	fmt.Printf("| Name        | Win      | Lost      | Draw    |\n")
	fmt.Printf("------------------------------------------------\n")
	fmt.Printf("| %s           | %d        | %d         | %d       |\n", Player1.name, Player1.win, Player1.lost, Player1.draw)
	fmt.Printf("------------------------------------------------\n")
	fmt.Printf("| %s           | %d        | %d         | %d       |\n", Player2.name, Player2.win, Player2.lost, Player2.draw)
	fmt.Printf("------------------------------------------------\n")

	fmt.Printf("======================\n")
	fmt.Println()

	fmt.Printf("Want to play again \n")
	fmt.Printf("(1)Yes (0)No\n")

	fmt.Scanf("%d", &again)
	if again == 1 {
		return 1
	}
	return 0

}

func ResetBoard(board [][]string) [][]string {
	// Reset the board for new play
	board = [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	return board
}

func Play(board [][]string, Player1 Player, Player2 Player) {
	// game logic where determine which one won or lost
	// by calling win function
	var pos int
	var status int
	var count int = 1
	var full int = 0
	var again int
	for {
		PrintBoard(board)

		// accept player1 position
		fmt.Printf("%s: ", Player1.name)
		fmt.Scanf("%d", &pos)

		first, second := Position(pos)
		if board[first][second] != "_" {
			fmt.Printf("choose another position!")
			continue
		}
		board[first][second] = Player1.sign
		if count >= 3 {
			status = Win(board, Player1.sign)
			if status == 1 {
				PrintBoard(board)
				Player1.win++
				Player2.lost++
				again = Again(board, Player1, Player2)
				if again == 1 {
					board = ResetBoard(board)
					continue
				} else {
					break
					return
				}

			}
		}
		PrintBoard(board)

		// check if the board is full or not
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == "_" {
					full = 0
					break
				} else {
					full++
				}
			}
		}
		if full >= 9 {
			fmt.Printf("Draw!\n")
			Player1.draw++
			Player2.draw++
			again = Again(board, Player1, Player2)
			if again == 1 {
				board = ResetBoard(board)
				continue
			} else {
				return
			}
		}

		// accept player2 position on board
		for {
			fmt.Printf("%s: ", Player2.name)
			fmt.Scanf("%d\n", &pos)
			first, second = Position(pos)

			if board[first][second] != "_" {
				fmt.Printf("choose another position!")
				continue
			} else {
				break
			}
		}

		board[first][second] = Player2.sign

		// logic starts after 3 attempts to check whether won or lost
		if count >= 3 {
			status = Win(board, Player2.sign)
			if status == 1 {
				fmt.Printf("%s Win\n", Player2.name)
				Player2.win++
				Player1.lost++
				PrintBoard(board)
				again = Again(board, Player1, Player2)
				if again == 1 {
					board = ResetBoard(board)
					continue
				} else {
					break
					return
				}
			}
		}
		count++
	}
}

func main() {

	// tic tac toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// define player1 and player2
	player1 := Player{}
	player2 := Player{}
	var choose string

	fmt.Printf("player 1 enter your name: ")
	fmt.Scanf("%s", &player1.name)
	fmt.Printf("Choose X or O: ")
	fmt.Scanf("%s", &choose)
	if strings.ToUpper(choose) == "X" || strings.ToUpper(choose) == "O" {
		player1.sign = strings.ToUpper(choose)
		if player1.sign == "X" {
			player2.sign = "O"
		} else {
			player2.sign = "X"
		}
		fmt.Printf("player 2 enter your name: ")
		fmt.Scanf("%s", &player2.name)
		Play(board, player1, player2)
	}

}

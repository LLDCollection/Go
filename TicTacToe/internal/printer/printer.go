package printer

import "fmt"

func PrintBoard(board [][]rune) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf(" %v ", string(board[i][j]))
			if j < len(board[i]) - 1 {
				fmt.Printf("|")
			}
		}
		fmt.Println()
	}
}
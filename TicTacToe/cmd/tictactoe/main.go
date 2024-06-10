package main

import (
	"fmt"
	"strconv"
	"strings"
	"tictactoe/internal/board"
)

func main() {
	game := board.New()

	for game.Winner() == -1 {
		fmt.Printf("Player %v make a move: (0-2),(0-2)", game.CurrentPlayer() + 1)
		
		var input string
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Printf("err: %v\n", err.Error())
		}

		moves := strings.Split(input, ",")
		move1, err := strconv.Atoi(moves[0])
		if err != nil {
			fmt.Printf("err: %v\n", err.Error())
		}

		move2, err := strconv.Atoi(moves[1])
		if err != nil {
			fmt.Printf("err: %v\n", err.Error())
		}

		err = game.Move(move1, move2)
		if err != nil {
			fmt.Printf("illegal move, play again\n")
		}
	}

	fmt.Printf("Winner is Player %d\n", game.Winner() + 1)
}

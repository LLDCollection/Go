package board

import (
	"fmt"
	"tictactoe/internal/player"
)

type Board struct {
	grid          [][]rune
	size          int
	currentPlayer int
	players       []*player.Player
	winner        int
}

func New() *Board {
	grid := make([][]rune, 3)
	for i := range grid {
		grid[i] = make([]rune, 3)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	p1 := player.New('X')
	p2 := player.New('O')

	return &Board{
		grid:          grid,
		size:          3,
		currentPlayer: 0,
		players:       []*player.Player{p1, p2},
		winner:        -1,
	}
}

func (b *Board) Move(row, col int) error {
	if !b.setMove(row, col) {
		return fmt.Errorf("unable to set move")
	}

	b.checkWinner(row, col)

	b.currentPlayer = (b.currentPlayer + 1) % len(b.players)
	fmt.Println(b.currentPlayer)

	return nil
}

func (b *Board) CurrentPlayer() int {
	return b.currentPlayer
}

func (b *Board) Winner() int {
	return b.winner
}

func (b *Board) checkWinner(row, col int) {
	player := b.players[0].Symbol()

	vertical := true
	for i := 0; i < len(b.grid); i++ {
		if b.grid[i][col] != player {
			vertical = false
			break
		}
	}
	if vertical {
		b.winner = b.currentPlayer
		return
	}

	horizontal := true
	for i := 0; i < len(b.grid[row]); i++ {
		if b.grid[row][i] != player {
			horizontal = false
			break
		}
	}
	if horizontal {
		b.winner = b.currentPlayer
		return
	}
}

func (b *Board) setMove(row, col int) bool {
	if b.grid[row][col] != ' ' {
		return false
	}

	b.grid[row][col] = b.players[0].Symbol()

	return true
}

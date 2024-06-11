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
	totalMoves    int
	draw          bool
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
		totalMoves:    0,
		draw:          false,
	}
}

func (b *Board) Move(row, col int) error {
	if !b.setMove(row, col) {
		return fmt.Errorf("unable to set move")
	}

	b.totalMoves++
	b.checkWinner(row, col)
	b.checkDraw()

	b.currentPlayer = (b.currentPlayer + 1) % len(b.players)

	return nil
}

func (b *Board) CurrentPlayer() int {
	return b.currentPlayer
}

func (b *Board) Winner() int {
	return b.winner
}

func (b *Board) IsDraw() bool {
	return b.draw
}

func (b *Board) Grid() [][]rune {
	return b.grid
}

func (b *Board) checkWinner(row, col int) {
	player := b.players[b.currentPlayer].Symbol()

	// Check for vertical win
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

	// Check for horizontal win
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

	// Check for diagonal win
	if row == col {
		diagonal := true
		for i := 0; i < len(b.grid); i++ {
			if b.grid[i][i] != player {
				diagonal = false
				break
			}
		}
		if diagonal {
			b.winner = b.currentPlayer
			return
		}
	}

	// Check for anti-diagonal win
	if row+col == len(b.grid)-1 {
		antiDiagonal := true
		for i := 0; i < len(b.grid); i++ {
			if b.grid[i][len(b.grid)-1-i] != player {
				antiDiagonal = false
				break
			}
		}
		if antiDiagonal {
			b.winner = b.currentPlayer
			return
		}
	}
}

func (b *Board) checkDraw() {
	if b.totalMoves == b.size*b.size {
		b.draw = true
	}
}

func (b *Board) setMove(row, col int) bool {
	const emptyCell rune = ' '
	if b.grid[row][col] != emptyCell {
		return false
	}

	b.grid[row][col] = b.players[b.currentPlayer].Symbol()

	return true
}

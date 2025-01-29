package game

import (
	"fmt"
)

type board struct {
	size int
	grid [][]int
}

func newBoard(n int) *board {
	createPieces()
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	return &board{
		size: n,
		grid: grid,
	}
}

func (b *board) fill(playerId, r, c int) bool {
	if b.grid[r][c] == 0 {
		b.grid[r][c] = playerId
		return true
	}
	return false
}

func (b *board) isWinner(p, row, col int) bool {
	rowcheck := true
	colcheck := true
	diacheck := true
	anticheck := true

	for i := 0; i < b.size; i++ {
		if p != b.grid[row][i] {
			rowcheck = false
			break
		}
	}
	for i := 0; i < b.size; i++ {
		if p != b.grid[i][col] {
			colcheck = false
			break
		}
	}
	for i := 0; i < b.size; i++ {
		if p != b.grid[i][i] {
			diacheck = false
			break
		}
	}
	for i := 0; i < b.size; i++ {
		if p != b.grid[i][b.size-i-1] {
			anticheck = false
			break
		}
	}
	return rowcheck || colcheck || anticheck || diacheck
}

func (b *board) printBoard() {
	for i := 0; i < b.size; i++ {
		extra := "_"
		if i == b.size-1 {
			extra = " "
		}
		s := ""
		for j := 0; j < b.size; j++ {
			if b.grid[i][j] == 0 {
				s += extra + extra + extra + "|"
			} else {
				s += extra + pieces[b.grid[i][j]].sign + extra + "|"
			}
		}
		fmt.Println(s[:len(s)-1])
	}
}

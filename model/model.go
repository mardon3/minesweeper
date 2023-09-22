package model

import (
	"math/rand"
)

type Cell struct {
	value int // 0-8 for number of mines in the adjacent cells, -1 for mine
	isRevealed bool
	isFlagged bool
}

type Board struct {
	cells [][]Cell
	height int
	width int
	mines int
	flags int
	hasStarted bool // Not sure if this needed
}

func NewBoard(width, height, mines int) *Board {
	board := &Board{
		cells: make([][]Cell, height),
		height: height,
		width: width,
		mines: mines,
	}

	for i := range board.cells {
		board.cells[i] = make([]Cell, width)
	}

	placedMines := 0
	for placedMines < mines {
		row := rand.Intn(height)
		col := rand.Intn(width)

		if !IsMine(board.cells, row, col) {
			board.cells[row][col].value = -1
			placedMines++
		}
	}

	for row := range board.cells {
		for col := range board.cells[row] {
			if !IsMine(board.cells, row, col) {
				board.cells[row][col].value = countAdjacentMines(board, row, col)
			}
		}
	}

	return board
}

func countAdjacentMines(board *Board, row, col int) int {
	adjacentMines := 0

	directions := []struct{ dirRow, dirCol int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	for _, direction := range directions {
		r, c := row + direction.dirRow, col + direction.dirCol

		// Check if the adjacent cell is within the board boundaries
		if r >= 0 && r < board.height && c >= 0 && c < board.width {
			if IsMine(board.cells, r, c) {
				adjacentMines++
			}
		}
	}

	return adjacentMines
}

func IsMine(cells [][]Cell, row, col int) bool {
	return cells[row][col].value == -1
}

func (board *Board) IsRevealed(row, col int) bool {
	return board.cells[row][col].isRevealed
}

func (board *Board) IsFlagged(row, col int) bool {
	return board.cells[row][col].isFlagged
}

func (board *Board) FlagCell(row, col int) {
	board.cells[row][col].isFlagged = !board.cells[row][col].isFlagged
}

func (board *Board) RevealCell(row, col int) {
	board.cells[row][col].isRevealed = true
}

// TODO: Reveal or GetCell function
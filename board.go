package main

import "math/rand"

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

		if board.cells[row][col].value != -1 {
			board.cells[row][col].value = -1
			placedMines++
		}
	}

	for row := range board.cells {
		for col := range board.cells[row] {
			if board.cells[row][col].value != -1 {
				board.cells[row][col].value = board.countAdjacentMines(row, col)
			}
		}
	}

	return board
}

func (board *Board) countAdjacentMines(row, col int) int {
	count := 0

	// TODO
}
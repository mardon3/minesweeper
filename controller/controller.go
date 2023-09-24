package controller

import (
	"musmanov/minesweeper/model"
)

// Reveal a cell
func LeftClickCell(row, col int) {
	if model.IsFlagged(row, col) || model.IsRevealed(row, col) {
		return
	}

	// if empty, reveal adjacent empty cells and their adjacent cells until all adjacent empty cells are revealed
	if model.GetCellType(row, col) == model.EmptyCell {

	} else {
		model.RevealCell(row, col)
	}
}

// Flags a cell
func RightClickCell(row, col int) {
	model.FlagCell(row, col)
}

func NewBoard(BoardDifficulty model.Difficulty) {
	model.NewBoard(BoardDifficulty)
}
package controller

import (
	"musmanov/minesweeper/model"
)

// Reveal a cell
func LeftClickCell(row, col int) {
	if model.IsFlagged(row, col) || model.IsRevealed(row, col) {
		return
	}

	// Ensures that the first click is always safe
	model.SafeStart(row, col)

	// if empty, reveal adjacent empty cells and their adjacent cells until all adjacent empty cells are revealed
	if model.GetCellType(row, col) == model.EmptyCell {
		model.RevealEmptyCell(row, col)
	} else { 
		model.RevealCell(row, col)
	}
}

// Flags a cell
func RightClickCell(row, col int) {
	model.FlagCell(row, col)
}

func NewBoard(boardDifficulty model.Difficulty) {
	model.NewBoard(boardDifficulty)
}
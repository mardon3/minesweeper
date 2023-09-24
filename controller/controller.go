package controller

import (
	"musmanov/minesweeper/model"
)

// Reveal a cell
func LeftClickCell(row, col int) {
	if model.IsFlagged(row, col) || model.IsRevealed(row, col) {
		return
	}

	if model.GetCellType(row, col) == model.ValueCell {
		model.RevealCell(row, col)
	} else {
		// if empty, reveal adjacent empty cells and their adjacent cells until all adjacent empty cells are revealed
	}
}

// Flags a cell
func RightClickCell(row, col int) {
	model.FlagCell(row, col)
}

func SelectDifficulty() {
}

func ClickResetGame() {
}

func NewBoard(row, col, mines int) {
	model.NewBoard(row, col, mines)
}
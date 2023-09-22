package controller

import (
	"musmanov/minesweeper/model"
)

// Reveal a cell
func LeftClickCell(row, col int) {
	model.NewBoard(row, col, 3)
}

// Flags a cell
func RightClickCell(row, col int) {
}

func SelectDifficulty() {
}

func ClickResetGame() {
}

func NewBoard(row, col, mines int) {
	model.NewBoard(row, col, mines)
}
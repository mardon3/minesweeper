package controller

import (
	"image/color"
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

// New/reset game
func NewBoard(boardDifficulty ...model.Difficulty) {
	if len(boardDifficulty) == 1 {
		model.NewBoard(boardDifficulty[0])
	} else if len(boardDifficulty) == 0 {
		model.NewBoard()
	} else {
		panic("Too many arguments passed to NewBoard")
	}
}

func GetBoard() [][]model.Cell {
	return model.GetBoard()
}

func GetScreenSize() (int, int) {
	return model.GetScreenSize()
}

func GetColor(row, col int) color.Color {
	return model.GetColor(row, col)
}

func GetDifficultyString() string {
	return model.GetDifficultyString()
}

func IsLost() bool {
	return model.IsLost()
}

func IsSolved() bool {
	return model.IsSolved()
}

func GetMineCount() int {
	return model.GetMineCount()
}

func GetFlagsString() string {
	return model.GetFlagsString()
}

func GetBoardHeight() int {
	return model.GetBoardHeight()
}

func GetBoardWidth() int {
	return model.GetBoardWidth()
}

func IsRevealed(row, col int) bool {
	return model.IsRevealed(row, col)
}

func IsFlagged(row, col int) bool {
	return model.IsFlagged(row, col)
}
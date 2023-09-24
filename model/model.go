package model

import (
	"math/rand"
)

type Cell struct {
	value int // 0-8 for number of mines in the adjacent cells, -1 for mine
	isRevealed bool
	isFlagged bool
}

var board [][]Cell
var flagsLeft int
var minesCount int
var unRevealedCells int

func NewBoard(BoardDifficulty Difficulty) {
	flagsLeft, minesCount = BoardDifficulty.Mines, BoardDifficulty.Mines
	unRevealedCells = BoardDifficulty.Height * BoardDifficulty.Width
	board = make([][]Cell, BoardDifficulty.Height)

	for i := range board {
		board[i] = make([]Cell, BoardDifficulty.Width)
	}

	placedMines := 0
	for placedMines < BoardDifficulty.Mines {
		row := rand.Intn(BoardDifficulty.Height)
		col := rand.Intn(BoardDifficulty.Width)

		if GetCellType(row, col) != MineCell {
			board[row][col].value = -1
			placedMines++
		}
	}

	for row := range board {
		for col := range board[row] {
			if GetCellType(row, col) != MineCell {
				board[row][col].value = countAdjacentMines(row, col)
			}
		}
	}
}

func countAdjacentMines(row, col int) int {
	adjacentMines := 0

	directions := []struct{dirRow, dirCol int}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	for _, direction := range directions {
		r, c := row + direction.dirRow, col + direction.dirCol

		// Check if the adjacent cell is within the board boundaries
		if r >= 0 && r < getBoardHeight() && c >= 0 && c < getBoardWidth() {
			if GetCellType(r, c) == MineCell {
				adjacentMines++
			}
		}
	}

	return adjacentMines
}

func getBoardHeight() int {
	return len(board)
}

func getBoardWidth() int {
	return len(board[0])
}

func IsRevealed(row, col int) bool {
	return board[row][col].isRevealed
}

func IsFlagged(row, col int) bool {
	return board[row][col].isFlagged
}

func FlagCell(row, col int) {
	if flagsLeft == 0 && !IsFlagged(row, col) {
		return
	}
	board[row][col].isFlagged = !board[row][col].isFlagged
}

func RevealCell(row, col int) {
	if board[row][col].isRevealed == false {
		board[row][col].isRevealed = true
		unRevealedCells--
	}
}

func RevealEmptyCells(row, col int) {

}

func GetCellType(row, col int) CellType {
	if board[row][col].value == -1 {
		return MineCell
	} else if board[row][col].value == 0 {
		return EmptyCell
	}
	return ValueCell
}

func IsLost() bool {
	for row := range board {
		for col := range board[row] {
			if GetCellType(row, col) == MineCell && IsRevealed(row, col) {
				return true
			}
		}
	}

	return false
}

func IsSolved() bool {
	if unRevealedCells == minesCount && !IsLost() {
		return true
	}

	return false
}
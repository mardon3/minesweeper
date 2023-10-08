package model

import (
	"image/color"
	"math/rand"
	"strconv"
)

type Cell struct {
	value int // 0-8 for number of mines in the adjacent cells, -1 for mine
	isRevealed bool
	isFlagged bool
	color color.Color
}

var ( 
	board [][]Cell
	flagsLeft int
	minesCount int
	unRevealedCells int
	currDifficulty Difficulty = Beginner
	firstClick bool
)


func NewBoard(boardDifficulty ...Difficulty) {
	if len(boardDifficulty) == 1 {
		currDifficulty = boardDifficulty[0]
	} else if len(boardDifficulty) > 1 {
		panic("Too many arguments passed to NewBoard")
	}
	firstClick = true
	flagsLeft, minesCount = currDifficulty.Mines, currDifficulty.Mines
	unRevealedCells = currDifficulty.Height * currDifficulty.Width
	board = make([][]Cell, currDifficulty.Height)

	for i := range board {
		board[i] = make([]Cell, currDifficulty.Width)
	}

	placedMines := 0
	for placedMines < currDifficulty.Mines {
		row := rand.Intn(currDifficulty.Height)
		col := rand.Intn(currDifficulty.Width)

		if GetCellType(row, col) != MineCell {
			board[row][col].value = -1
			placedMines++
		}
	}

	for row := range board {
		for col := range board[row] {
			board[row][col].color = color.RGBA{114, 199, 255, 255}
			if GetCellType(row, col) != MineCell {
				board[row][col].value = countAdjacentMines(row, col)
			}
		}
	}
}

func countAdjacentMines(currRow, currCol int) int {
	adjacentMines := 0

	directions := []struct{dirRow, dirCol int}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	for _, direction := range directions {
		r, c := currRow + direction.dirRow, currCol + direction.dirCol

		// Check if the adjacent cell is within the board boundaries
		if r >= 0 && r < GetBoardHeight() && c >= 0 && c < GetBoardWidth() {
			if GetCellType(r, c) == MineCell {
				adjacentMines++
			}
		}
	}

	return adjacentMines
}

func GetBoardHeight() int {
	return len(board)
}

func GetBoardWidth() int {
	return len(board[0])
}

func GetBoard() [][]Cell {
	return board
}

func GetScreenSize() (int, int) {
	cellMargins := 2
	if currDifficulty == Beginner {
		return (currDifficulty.Width * 50) + ((currDifficulty.Width + 1) * cellMargins),((currDifficulty.Height * 50) + ((currDifficulty.Height + 1) * cellMargins)) + 60
	} else if currDifficulty == Intermediate {
		return (currDifficulty.Width * 36) + ((currDifficulty.Width + 1) * cellMargins), ((currDifficulty.Height * 36) + ((currDifficulty.Height + 1) * cellMargins)) + 60
	} else {
		return (currDifficulty.Width * 30) + ((currDifficulty.Width + 1) * cellMargins), ((currDifficulty.Height * 30) + ((currDifficulty.Height + 1) * cellMargins)) + 60
	}
}

func GetDifficultyString() string {
	if currDifficulty == Beginner {
		return "Beginner"
	} else if currDifficulty == Intermediate {
		return "Intermediate"
	} else {
		return "Expert"
	}
}

func GetDifficulty() Difficulty {
	return currDifficulty
}

func GetMineCount() int {
	return minesCount
}

func GetCellType(row, col int) CellType {
	if board[row][col].value == -1 {
		return MineCell
	} else if board[row][col].value == 0 {
		return EmptyCell
	}

	return ValueCell
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
	} else if IsFlagged(row, col) {
		flagsLeft++
	} else {
		flagsLeft--
	}
	
	board[row][col].isFlagged = !board[row][col].isFlagged
}

func RevealCell(row, col int) {
	// When an entire region is being cleared, and there's an incorrectly flagged cell
	if IsFlagged(row, col) {
		flagsLeft++
	}

	board[row][col].color = color.RGBA{150, 200, 255, 255}
	board[row][col].isFlagged = false
	board[row][col].isRevealed = true
	unRevealedCells--
}

func RevealEmptyCell(row, col int) {
	// When an entire region is being cleared, and there's an incorrectly flagged cell
	if IsFlagged(row, col) {
		flagsLeft++
	}

	board[row][col].color = color.RGBA{150, 200, 255, 255}
	board[row][col].isFlagged = false
	board[row][col].isRevealed = true
	unRevealedCells--

	directions := []struct{dirRow, dirCol int}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}
	
	for _, direction := range directions {
		r, c := row + direction.dirRow, col + direction.dirCol

		// Check if the adjacent cell is within the board boundaries
		if r >= 0 && r < GetBoardHeight() && c >= 0 && c < GetBoardWidth() {
			if GetCellType(r, c) == ValueCell {
				RevealCell(r, c)
			} else if GetCellType(r, c) == EmptyCell {
				RevealEmptyCell(r, c)
			}
		}
	} 
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

func SafeStart(row, col int) {
	if !firstClick {
		return
	}

	if GetCellType(row, col) == MineCell {
		// Iterate from the bottom right of board to the top left
		for r := GetBoardHeight() - 1; r >= 0; r-- {
			for c := GetBoardWidth() - 1; c >= 0; c-- {
				if GetCellType(r, c) != MineCell {
					// Swap the cell with the given cell as param and adjust adjacent cells accordingly
					swapCells(row, col, r, c)
				}
			}
		}	
	}

	firstClick = false
}

func swapCells(row1, col1, row2, col2 int) {
	board[row1][col1], board[row2][col2] = board[row2][col2], board[row1][col1]

	board[row1][col1].value = countAdjacentMines(row1, col1)

	updateAdjacentCells(row1, col1)
	updateAdjacentCells(row2, col2)
}

// Update surrounding cells in-case of a cell swap for a safe start
func updateAdjacentCells(currRow, currCol int) {
	directions := []struct{dirRow, dirCol int}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}

	for _, direction := range directions {
		r, c := currRow + direction.dirRow, currCol + direction.dirCol

		// Check if the adjacent cell is within the board boundaries
		if r >= 0 && r < GetBoardHeight() && c >= 0 && c < GetBoardWidth() {
			if GetCellType(currRow, currCol) == MineCell && GetCellType(r, c) != MineCell {
				board[r][c].value++
			} else if GetCellType(currRow, currCol) != MineCell && GetCellType(r, c) != MineCell {
				board[r][c].value--
			}
		}
	}
}

func GetColor(row, col int) color.Color {
	return board[row][col].color
}

func GetFlagsString() string {
	return strconv.Itoa(flagsLeft)
}
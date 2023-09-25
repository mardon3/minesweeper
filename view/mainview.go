package view

import (
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	board [][]model.Cell
}

// Generates a new board
func NewGame(boardDifficulty ...model.Difficulty) *Game {
	controller.NewBoard(boardDifficulty...)
	ebiten.SetWindowSize(controller.GetScreenSize())
	ebiten.SetWindowTitle("Minesweeper")
	game := &Game{
		board: controller.GetBoard(),
	}

	return game
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int)  {
	return controller.GetScreenSize()
}
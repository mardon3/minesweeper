package view

import (
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game object used by ebiten
type Game struct {
	ui *ebitenui.UI
}

func NewRootContainer() *widget.Container {
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.BoardBackgroundColor)),
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(2),
			//Define how much padding to inset the child content
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(4)),
			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(0, 0),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{false, false, false, false, false, false, false, false, false}, []bool{false, false, false, false, false, false, false, false, false}),
		)),
	)

	return rootContainer
}
// Generates a new board
func NewGame(boardDifficulty ...model.Difficulty) *Game {
	controller.NewBoard(boardDifficulty...)
	ebiten.SetWindowSize(controller.GetScreenSize())
	ebiten.SetWindowTitle("Minesweeper")
	ui := ebitenui.UI{
		Container: NewRootContainer(),
	}

	game := &Game{
		ui: &ui,
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
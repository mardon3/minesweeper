package view

import (
	"image/color"
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewRootContainer() *widget.Container {
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.White)),
		// the container will use a single column layout, header on top, board on bottom
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(0)),
			widget.GridLayoutOpts.Spacing(0, 0),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true}),
		)),
	)

	RenderHeader()
	rootContainer.AddChild(HeaderContainer)
	RenderBoard()
	rootContainer.AddChild(BoardContainer)

	return rootContainer
}

// Generates a new board
func NewUI(boardDifficulty ...model.Difficulty) *ebitenui.UI {
	controller.NewBoard(boardDifficulty...)
	// To keep board square size, since header is set to 72 pixels height
	ebiten.SetWindowSize(boardDifficulty[0].ScreenWidth, boardDifficulty[0].ScreenHeight)
	ebiten.SetWindowTitle("Minesweeper")
	ui := ebitenui.UI{
		Container: NewRootContainer(),
	}

	return &ui
}
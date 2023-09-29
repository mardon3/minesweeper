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
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(1),
			//Define how much padding to inset the child content
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(0)),
			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(0, 0),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true}),
		)),
	)

	rootContainer.AddChild(NewHeaderContainer())
	rootContainer.AddChild(NewBoardContainer())

	return rootContainer
}
// Generates a new board
func NewUI(boardDifficulty ...model.Difficulty) *ebitenui.UI {
	controller.NewBoard(boardDifficulty...)
	NewRootContainer()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Minesweeper")
	ui := ebitenui.UI{
		Container: NewRootContainer(),
	}

	return &ui
}
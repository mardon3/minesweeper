package view

import (
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

/*
The Grid Layout is built to position children in a rows and columns.
*/
func NewBoardContainer() *widget.Container {
	boardContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.BoardBackgroundColor)),
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(1),
			//Define how much padding to inset the child content
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(4)),
			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(0, 0),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true}),
		)),
	)

	return boardContainer
}
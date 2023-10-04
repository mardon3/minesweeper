package view

import (
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

var (
	BoardContainer *widget.Container = newBoardContainer();
)

func newBoardContainer() *widget.Container {
	screenWidth := controller.GetBoardWidth()
	screenHeight := controller.GetBoardHeight()

	columnStretchSlice := make([]bool, screenWidth)
	for i := range columnStretchSlice {
		columnStretchSlice[i] = true
	}

	rowStretchSlice := make([]bool, screenHeight)
	for i := range rowStretchSlice {
		rowStretchSlice[i] = true
	}

	boardContainer := widget.NewContainer(
		// the container will use assigned background color
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.BoardBackgroundColor)),
		// container will generate as many columns as the screen width requires
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(screenWidth),
			//Define how much padding to inset the child content
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(4)),
			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(0, 0),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch(columnStretchSlice, rowStretchSlice),
		)),
	)

	return boardContainer
}

func RenderBoard() *widget.Container {	
	return BoardContainer
}
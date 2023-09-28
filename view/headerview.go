package view

import (
	"image/color"
	"log"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
The Grid Layout is built to position children in a rows and columns.
*/
func main() {

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.BoardBackgroundColor)),
		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			//Define number of columns in the grid
			widget.GridLayoutOpts.Columns(9),
			//Define how much padding to inset the child content
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(4)),
			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(4, 4),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{false, false, false, false, false, false, false, false, false}, []bool{false, false, false, false, false, false, false, false, false}),
		)),
	)

	headerCountainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.HeaderBackgroundColor)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 100),
		),
	)
	rootContainer.AddChild(headerCountainer)

	innerContainer2 := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 255, 0, 255})),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 100),
		),
	)
	rootContainer.AddChild(innerContainer2)

	innerContainer3 := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 0, 255, 255})),
		widget.ContainerOpts.WidgetOpts(
			//The widget in this cell has a MaxHeight and MaxWidth less than the
			//Size of the grid cell so it will use the Position fields below to
			//Determine where the widget should be displayed within that grid cell.
			widget.WidgetOpts.LayoutData(widget.GridLayoutData{
				HorizontalPosition: widget.GridLayoutPositionCenter,
				VerticalPosition:   widget.GridLayoutPositionCenter,
				MaxWidth:           100,
				MaxHeight:          100,
			}),
			widget.WidgetOpts.MinSize(100, 100),
		),
	)
	rootContainer.AddChild(innerContainer3)

	// innerContainer4 := widget.NewContainer(
	// 	widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0, 255, 255, 255})),
	// 	widget.ContainerOpts.WidgetOpts(
	// 		widget.WidgetOpts.MinSize(100, 100),
	// 	),
	// )
	// rootContainer.AddChild(innerContainer4)
	// construct the UI
	ui := ebitenui.UI{
		Container: rootContainer,
	}

	// // Ebiten setup
    // ebiten.SetWindowSize(400, 400)
	// ebiten.SetWindowTitle("Ebiten UI - Grid Layout")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := Game{
		ui: &ui,
	}

	// run Ebiten main loop
	err := ebiten.RunGame(&game)
	if err != nil {
		log.Println(err)
	}
}
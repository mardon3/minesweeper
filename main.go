package main

import (
	"log"
	"musmanov/minesweeper/model"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game object used by ebiten
type Game struct {
	ui *ebitenui.UI
}

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
			widget.GridLayoutOpts.Columns(2),
			//Define how much padding to inset the child content
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(0)),
			//Define how far apart the rows and columns should be
			widget.GridLayoutOpts.Spacing(0, 0),
			//Define how to stretch the rows and columns. Note it is required to
			//specify the Stretch for each row and column.
			widget.GridLayoutOpts.Stretch([]bool{true, false}, []bool{false, false}),
		)),
	)

	headerCountainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(model.HeaderBackgroundColor)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(10, 100),
		),
	)
	rootContainer.AddChild(headerCountainer)

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

// Layout implements Game.
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

// Update implements Game.
func (g *Game) Update() error {
	// update the UI
	g.ui.Update()
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *Game) Draw(screen *ebiten.Image) {
	// draw the UI onto the screen
	g.ui.Draw(screen)
}
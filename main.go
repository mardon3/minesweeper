package main

import (
	"log"
	"musmanov/minesweeper/controller"
	"musmanov/minesweeper/model"
	"musmanov/minesweeper/view"
	"strconv"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game object used by ebiten
type Game struct {
	ui *ebitenui.UI
}

func main() {
	game := Game{
		ui: view.NewUI(model.Beginner),
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
	if controller.IsSolved() {
		view.ResetEmoteGraphic.Image = view.HypersIcon
	} else if controller.IsLost() {
		view.ResetEmoteGraphic.Image = view.PepeHandsIcon
	} else {
		view.ResetEmoteGraphic.Image = view.PepoThinkIcon
	}
	
	for r := 0; r < controller.GetBoardHeight(); r++ {
		for c := 0; c < controller.GetBoardWidth(); c++ {
			view.BoardCells[r][c].Image, _ = view.LoadCellImage(r, c)
			if controller.IsRevealed(r, c) && controller.GetCellValue(r, c) == 0 {
				view.BoardCells[r][c].Text().Label = ""
			} else if controller.IsRevealed(r, c) {
				view.BoardCells[r][c].Text().Label = strconv.Itoa(controller.GetCellValue(r, c))
			}
		}
	}

	view.FlagsCounterText.Label = controller.GetFlagsString()

	g.ui.Update()
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *Game) Draw(screen *ebiten.Image) {
	// draw the UI onto the screen
	g.ui.Draw(screen)
}
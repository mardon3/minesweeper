package main

import (
	"image/color"
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
			if controller.IsRevealed(r, c) {
				cellValue := controller.GetCellValue(r, c)
				switch cellValue {
					case 0:
						view.BoardCells[r][c].TextColor.Idle = color.White
					case 1:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{0, 0, 255, 255}
					case 2:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{0, 128, 0, 255}
					case 3:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{255, 0, 0, 255}
					case 4:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{0, 0, 128, 255}
					case 5:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{128, 0, 0, 255}
					case 6:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{0, 128, 128, 255}
					case 7:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{0, 0, 0, 255}
					case 8:
						view.BoardCells[r][c].TextColor.Idle = color.NRGBA{128, 128, 128, 255}
				}
				view.BoardCells[r][c].Text().Label = strconv.Itoa(cellValue)
			}
		}
	}

	view.TimerText.Label = controller.GetTimerString()
	view.FlagsCounterText.Label = controller.GetFlagsString()

	g.ui.Update()
	return nil
}

// Draw implements Ebiten's Draw method.
func (g *Game) Draw(screen *ebiten.Image) {
	// draw the UI onto the screen
	g.ui.Draw(screen)
}
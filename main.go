package main

import (
	"log"

	"musmanov/minesweeper/view"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
    game := view.NewGame()
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
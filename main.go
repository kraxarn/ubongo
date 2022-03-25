package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/app"
	"github.com/kraxarn/ubongo/game"
	"log"
)

func main() {
	ebiten.SetWindowSize(540, 960)
	ebiten.SetWindowTitle(fmt.Sprintf("%s v%s", app.Name, app.Version()))
	ebiten.SetWindowResizable(true)

	err := ebiten.RunGame(game.NewGame())
	if err != nil {
		log.Fatal(err)
	}
}

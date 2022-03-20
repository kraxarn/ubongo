package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game"
	"log"
	"runtime"
)

func main() {
	ebiten.SetWindowTitle("Ubongo")
	ebiten.SetWindowResizable(true)

	// Fullscreen on mobile, 540x960 on desktop
	if runtime.GOOS == "android" {
		ebiten.SetFullscreen(true)
	} else {
		ebiten.SetWindowSize(540, 960)
	}

	err := ebiten.RunGame(game.NewGame())
	if err != nil {
		log.Fatal(err)
	}
}

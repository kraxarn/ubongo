package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(540, 960)
	ebiten.SetWindowTitle("Ubongo")

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}

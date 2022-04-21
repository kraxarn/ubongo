package main

import (
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game"
	"github.com/kraxarn/ubongo/game/app"
	"log"
)

func main() {
	ebiten.SetWindowSize(540, 960)
	ebiten.SetWindowTitle(fmt.Sprintf("%s v%s", app.Name, app.Version()))
	ebiten.SetWindowResizable(true)

	opt := game.Options{}
	flag.BoolVar(&opt.DebugMode, "debug", false, "enable debug drawing")
	flag.Parse()

	err := ebiten.RunGame(game.NewGame(&opt))
	if err != nil {
		log.Fatal(err)
	}
}

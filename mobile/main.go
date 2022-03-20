package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/kraxarn/ubongo/game"
)

func init() {
	mobile.SetGame(game.NewGame())
}

// Dummy is a dummy exported function
func Dummy() {
}

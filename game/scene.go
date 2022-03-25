package game

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Update(game *Game) error
	Draw(screen *ebiten.Image)
}

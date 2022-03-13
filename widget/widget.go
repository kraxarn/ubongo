package widget

import "github.com/hajimehoshi/ebiten/v2"

type Widget interface {
	Update()
	Draw(dst *ebiten.Image)
}

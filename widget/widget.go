package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Widget interface {
	Update(ui *Ui)
	Draw(dst *ebiten.Image)
}

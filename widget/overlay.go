package widget

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"golang.org/x/image/font"
)

type Overlay struct {
	label *Label
}

func NewOverlay(font font.Face) *Overlay {
	return &Overlay{
		label: NewLabel(font, 16, 24, "", colors.ForegroundAlt),
	}
}

func (o *Overlay) Update(*Ui) {
	o.label.text = fmt.Sprintf("%.0f\n%.0f",
		ebiten.CurrentFPS(), ebiten.CurrentTPS())
}

func (o *Overlay) Draw(dst *ebiten.Image) {
	o.label.Draw(dst)
}

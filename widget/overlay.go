package widget

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/res"
)

type Overlay struct {
	label *Label
}

func NewOverlay() (*Overlay, error) {
	font, err := res.Font(res.FontDebug, 18)
	if err != nil {
		return nil, err
	}

	return &Overlay{
		label: NewLabel(font, 16, 24, "", colors.ForegroundAlt),
	}, nil
}

func (o *Overlay) Update(*Ui) {
	o.label.text = fmt.Sprintf("FPS: %.0f\nTPS: %.0f",
		ebiten.CurrentFPS(), ebiten.CurrentTPS())
}

func (o *Overlay) Draw(dst *ebiten.Image) {
	o.label.Draw(dst)
}

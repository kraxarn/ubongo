package widget

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/colors"
	"github.com/kraxarn/ubongo/resources"
	"runtime"
)

type Overlay struct {
	fps   float64
	os    string
	label *Label
}

func NewOverlay() (*Overlay, error) {
	font, err := resources.Font(resources.FontDebug, 18)
	if err != nil {
		return nil, err
	}

	return &Overlay{
		fps:   ebiten.CurrentFPS(),
		os:    fmt.Sprintf("%s (%s)", runtime.GOOS, runtime.GOARCH),
		label: NewLabel(font, 16, 24, "", colors.ForegroundAlt),
	}, nil
}

func (o *Overlay) Update(_ *Ui) {
	o.label.text = fmt.Sprintf("%.1f\n%s",
		ebiten.CurrentFPS(), o.os)
}

func (o *Overlay) Draw(dst *ebiten.Image) {
	o.label.Draw(dst)
}

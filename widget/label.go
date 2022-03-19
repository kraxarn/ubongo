package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kraxarn/ubongo/color"
	"github.com/kraxarn/ubongo/util"
	"golang.org/x/image/font"
)

type Label struct {
	fontFace font.Face
	position util.Vector2[int]
	text     string
}

func NewLabel(fontFace font.Face, x, y int, text string) *Label {
	return &Label{
		fontFace: fontFace,
		position: util.Vec2(x, y),
		text:     text,
	}
}

func (l *Label) Update(_ *Ui) {
	// Static text, nothing to update
}

func (l *Label) Draw(dst *ebiten.Image) {
	text.Draw(dst, l.text, l.fontFace, l.position.X, l.position.Y, color.Foreground)
}

func (l *Label) SetPosition(x, y int) {
	l.position = util.Vec2(x, y)
}

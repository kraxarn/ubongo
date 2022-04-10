package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kraxarn/ubongo/util/vec2"
	"golang.org/x/image/font"
	"image/color"
)

type Label struct {
	fontFace font.Face
	position vec2.Vector2[int]
	text     string
	color    color.Color
}

func NewLabel(fontFace font.Face, x, y int, text string, color color.Color) *Label {
	return &Label{
		fontFace: fontFace,
		position: vec2.New(x, y),
		text:     text,
		color:    color,
	}
}

func (l *Label) Update(_ *Ui) {
	// Static text, nothing to update
}

func (l *Label) Draw(dst *ebiten.Image) {
	text.Draw(dst, l.text, l.fontFace, l.position.X, l.position.Y, l.color)
}

func (l *Label) SetPosition(x, y int) {
	l.position = vec2.New(x, y)
}

func (l *Label) Size() vec2.Vector2[int] {
	bounds, _ := font.BoundString(l.fontFace, l.text)
	x := (bounds.Max.X - bounds.Min.X).Ceil()
	y := (bounds.Max.Y - bounds.Min.Y).Ceil()
	return vec2.New(x, y)
}

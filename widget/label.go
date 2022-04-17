package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image"
	"image/color"
)

type Label struct {
	fontFace font.Face
	position image.Point
	text     string
	color    color.Color
}

func NewLabel(fontFace font.Face, x, y int, text string, color color.Color) *Label {
	return &Label{
		fontFace: fontFace,
		position: image.Pt(x, y),
		text:     text,
		color:    color,
	}
}

func (l *Label) Update(*Ui) {
}

func (l *Label) Draw(dst *ebiten.Image) {
	text.Draw(dst, l.text, l.fontFace, l.position.X, l.position.Y, l.color)
}

func (l *Label) SetPosition(x, y int) {
	l.position = image.Pt(x, y)
}

func (l *Label) Size() image.Point {
	bounds, _ := font.BoundString(l.fontFace, l.text)
	x := (bounds.Max.X - bounds.Min.X).Ceil()
	y := (bounds.Max.Y - bounds.Min.Y).Ceil()
	return image.Pt(x, y)
}

func (l *Label) SetText(text string) {
	l.text = text
}

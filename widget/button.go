package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kraxarn/ubongo/assets"
	"github.com/kraxarn/ubongo/color"
	"golang.org/x/image/font"
	"image"
)

type Button struct {
	background *ebiten.Image
	font       font.Face
	fontHeight int
	rect       image.Rectangle
	Text       string
	mouseDown  bool
	isPressed  bool
	onPressed  func(b *Button)
}

func NewButton(background *ebiten.Image, font font.Face, x, y, w, h int, text string) *Button {
	bounds, _, _ := font.GlyphBounds('M')

	return &Button{
		Text:       text,
		background: background,
		font:       font,
		fontHeight: (bounds.Max.Y - bounds.Min.Y).Ceil(),
		rect:       image.Rect(x, y, x+w, y+h),
		mouseDown:  false,
		isPressed:  false,
		onPressed:  nil,
	}
}

func (b *Button) SetOnPressed(pressed func(b *Button)) {
	b.onPressed = pressed
}

func (b *Button) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if b.rect.Min.X <= x && x < b.rect.Max.X && b.rect.Min.Y <= y && y < b.rect.Max.Y {
			b.mouseDown = true
		} else {
			b.mouseDown = false
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
}

func (b *Button) Draw(dst *ebiten.Image) {
	t := assets.UiButton
	if b.mouseDown {
		t = assets.UiButtonPressed
	}

	drawNinePatch(b.background, dst, assets.UiImageRects[t], b.rect)

	bounds, _ := font.BoundString(b.font, b.Text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := b.rect.Min.X + (b.rect.Dx()-w)/2
	y := b.rect.Max.Y - (b.rect.Dy()-b.fontHeight)/2
	text.Draw(dst, b.Text, b.font, x, y, color.ForegroundAlt)
}

package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kraxarn/ubongo/colors"
	"github.com/kraxarn/ubongo/resources"
	"github.com/kraxarn/ubongo/util/vec2"
	"golang.org/x/image/font"
	"image"
)

// Button is a default button that can be pressed
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

func (b *Button) Update(_ *Ui) {
	var positions []vec2.Vector2[int]
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		positions = append(positions, vec2.New(ebiten.CursorPosition()))
	} else {
		for _, id := range ebiten.AppendTouchIDs([]ebiten.TouchID{}) {
			positions = append(positions, vec2.New(ebiten.TouchPosition(id)))
		}
	}

	if len(positions) > 0 {
		b.mouseDown = false
		for _, pos := range positions {
			if b.rect.Min.X <= pos.X && pos.X < b.rect.Max.X && b.rect.Min.Y <= pos.Y && pos.Y < b.rect.Max.Y {
				b.mouseDown = true
				break
			}
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
	t := resources.UiButton
	if b.mouseDown {
		t = resources.UiButtonPressed
	}

	drawNinePatch(b.background, dst, resources.UiImageRects[t], b.rect)

	bounds, _ := font.BoundString(b.font, b.Text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := b.rect.Min.X + (b.rect.Dx()-w)/2
	y := b.rect.Max.Y - (b.rect.Dy()-b.fontHeight)/2
	text.Draw(dst, b.Text, b.font, x, y, colors.Foreground)
}

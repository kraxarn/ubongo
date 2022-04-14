package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/game/resources"
	"golang.org/x/image/font"
)

// TextButton is a button with text
type TextButton struct {
	*Button
	background *NinePatch
	font       font.Face
	fontHeight int
	Text       string
}

func NewTextButton(background *ebiten.Image, font font.Face, x, y, w, h int, text string) *TextButton {
	bounds, _, _ := font.GlyphBounds('M')
	return &TextButton{
		Button:     NewButton(x, y, w, h),
		background: NewNinePatch(background, x, y, w, h),
		font:       font,
		fontHeight: (bounds.Max.Y - bounds.Min.Y).Ceil(),
		Text:       text,
	}
}

func (t *TextButton) Update(ui *Ui) {
	t.Button.Update(ui)
}

func (t *TextButton) Draw(dst *ebiten.Image) {
	t.Button.Draw(dst)

	var imageType resources.UiImageType
	if t.Button.mouseDown {
		imageType = resources.UiButtonPressed
	} else {
		imageType = resources.UiButton
	}

	t.background.SetSourceRect(resources.UiImageRects[imageType])
	t.background.Draw(dst)

	bounds, _ := font.BoundString(t.font, t.Text)
	w := (bounds.Max.X - bounds.Min.X).Ceil()
	x := t.rect.Min.X + (t.rect.Dx()-w)/2
	y := t.rect.Max.Y - (t.rect.Dy()-t.fontHeight)/2

	text.Draw(dst, t.Text, t.font, x, y, colors.Foreground)
}

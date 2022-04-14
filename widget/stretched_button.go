package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// StretchedButton is a button that stretches horizontally across the screen
type StretchedButton struct {
	*TextButton
	x     int
	y     int
	align Alignment
}

func NewStretchedButton(background *ebiten.Image, fontFace font.Face, x, y, w, h int,
	align Alignment, text string) *StretchedButton {
	return &StretchedButton{
		NewTextButton(background, fontFace, x, y, w, h, text),
		x,
		y,
		align,
	}
}

func (s *StretchedButton) Update(ui *Ui) {
	// Because the button stretches, alignment on x doesn't matter
	if s.x >= 0 {
		s.TextButton.rect.Max.X = ui.screenSize.X - ScreenPadding
	}

	if s.y >= 0 {
		ui.setAbsoluteY(&s.TextButton.rect, s.y, s.align)
	}

	s.TextButton.Update(ui)
}

func (s *StretchedButton) Draw(dst *ebiten.Image) {
	s.TextButton.Draw(dst)
}

package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

// StretchedButton is a button that stretches horizontally across the screen
type StretchedButton struct {
	*Button
	x     int
	y     int
	align Alignment
}

func NewStretchedButton(background *ebiten.Image, fontFace font.Face, x, y, w, h int,
	align Alignment, text string) *StretchedButton {
	return &StretchedButton{
		NewButton(background, fontFace, x, y, w, h, text),
		x,
		y,
		align,
	}
}

func (s *StretchedButton) Update(ui *Ui) {
	// Because the button stretches, alignment on x doesn't matter
	if s.x >= 0 {
		s.Button.rect.Max.X = ui.screenSize.X - ScreenPadding
	}

	if s.y >= 0 {
		ui.setAbsoluteY(&s.Button.rect, s.y, s.align)
	}

	s.Button.Update(ui)
}

func (s *StretchedButton) Draw(dst *ebiten.Image) {
	s.Button.Draw(dst)
}

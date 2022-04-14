package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	resources2 "github.com/kraxarn/ubongo/game/resources"
	"github.com/kraxarn/ubongo/util/vec2"
	"golang.org/x/image/font"
	"image"
)

type Ui struct {
	image      *ebiten.Image
	fontButton font.Face
	fontTitle  font.Face
	widgets    []Widget
	screenSize vec2.Vector2[int]
}

func NewUi() (*Ui, error) {
	uiImage, err := resources2.Image("ui")
	if err != nil {
		return nil, err
	}

	fontButton, err := resources2.Font(resources2.FontRegular, ButtonFontSize)
	if err != nil {
		return nil, err
	}

	fontTitle, err := resources2.Font(resources2.FontTitle, TitleFontSize)
	if err != nil {
		return nil, err
	}

	return &Ui{
		image:      uiImage,
		fontButton: fontButton,
		fontTitle:  fontTitle,
	}, nil
}

func (u *Ui) addWidgets(widgets ...Widget) {
	u.widgets = append(u.widgets, widgets...)
}

func addWidget[W Widget](ui *Ui, widget W) W {
	ui.addWidgets(widget)
	return widget
}

func (u *Ui) setAbsoluteY(rect *image.Rectangle, y int, align Alignment) {
	if (align & AlignTop) > 0 {
		rect.Min.Y = y
		rect.Max.Y = rect.Min.Y + ButtonHeight
		return
	}

	if (align & AlignBottom) > 0 {
		rect.Max.Y = u.screenSize.Y - y
		rect.Min.Y = rect.Max.Y - ButtonHeight
		return
	}
}

func (u *Ui) Update(screenSize vec2.Vector2[int]) {
	u.screenSize = screenSize

	for _, widget := range u.widgets {
		widget.Update(u)
	}
}

func (u *Ui) Draw(screen *ebiten.Image) {
	for _, widget := range u.widgets {
		widget.Draw(screen)
	}
}

func (u *Ui) AddTextButton(x, y, w, h int, text string) *TextButton {
	return addWidget(u, NewTextButton(u.image, u.fontButton, x, y, w, h, text))
}

func (u *Ui) AddStretchedButton(y int, align Alignment, text string) *StretchedButton {
	return addWidget(u, NewStretchedButton(u.image, u.fontButton,
		ScreenPadding*2, y,
		u.screenSize.X-ScreenPadding, ButtonHeight,
		align, text))
}

func (u *Ui) AddImage(src *ebiten.Image, x, y, w, h int) *Image {
	return addWidget(u, NewImage(src, x, y, w, h))
}

func (u *Ui) AddRepeatImage(src *ebiten.Image, x, y, w, h int) *RepeatImage {
	return addWidget(u, NewRepeatImage(src, x, y, w, h))
}

func (u *Ui) AddTitle(x, y int, text string) *Label {
	return addWidget(u, NewLabel(u.fontTitle, x, y, text, colors.ForegroundAlt))
}

func (u *Ui) AddLabel(x, y int, text string) *Label {
	return addWidget(u, NewLabel(u.fontButton, x, y, text, colors.ForegroundAlt))
}

func (u *Ui) AddDebugOverlay() *Overlay {
	widget, err := NewOverlay()
	if err != nil {
		return nil
	}
	return addWidget(u, widget)
}

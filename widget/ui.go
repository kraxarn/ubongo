package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/assets"
	"github.com/kraxarn/ubongo/enum"
	"github.com/kraxarn/ubongo/util"
	"golang.org/x/image/font"
)

type Ui struct {
	image      *ebiten.Image
	fontButton font.Face
	widgets    []Widget
	screenSize util.Vector2[int]
}

func NewUi() (*Ui, error) {
	uiImage, err := assets.Image("ui")
	if err != nil {
		return nil, err
	}

	fontButton, err := assets.Font(assets.FontRegular, ButtonFontSize)
	if err != nil {
		return nil, err
	}

	return &Ui{
		image:      uiImage,
		fontButton: fontButton,
	}, nil
}

func (u *Ui) addWidgets(widgets ...Widget) {
	u.widgets = append(u.widgets, widgets...)
}

func (u *Ui) setAbsoluteY(rect *image.Rectangle, y int, align enum.Alignment) {
	if (align & enum.AlignTop) > 0 {
		rect.Min.Y = y
		rect.Max.Y = rect.Min.Y + ButtonHeight
		return
	}

	if (align & enum.AlignBottom) > 0 {
		rect.Max.Y = u.screenSize.Y - y
		rect.Min.Y = rect.Max.Y - ButtonHeight
		return
	}
}

func (u *Ui) Update(screenSize util.Vector2[int]) {
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

func (u *Ui) AddButton(x, y, w, h int, text string) *Button {
	button := NewButton(u.image, u.fontButton, x, y, w, h, text)
	u.addWidgets(button)
	return button
}

func (u *Ui) AddStretchedButton(y int, align enum.Alignment, text string) *StretchedButton {
	button := NewStretchedButton(u.image, u.fontButton,
		ScreenPadding*2, y,
		u.screenSize.X-ScreenPadding, ButtonHeight,
		align, text)

	u.addWidgets(button)
	return button
}

func (u *Ui) AddImage(src *ebiten.Image, x, y int) *Image {
	img := NewImage(src, x, y)
	u.addWidgets(img)
	return img
}

package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/color"
	"github.com/kraxarn/ubongo/enum"
	"github.com/kraxarn/ubongo/resources"
	"github.com/kraxarn/ubongo/util"
	"golang.org/x/image/font"
	"image"
)

type Ui struct {
	image      *ebiten.Image
	fontButton font.Face
	fontTitle  font.Face
	widgets    []Widget
	screenSize util.Vector2[int]
}

func NewUi() (*Ui, error) {
	uiImage, err := resources.Image("ui")
	if err != nil {
		return nil, err
	}

	fontButton, err := resources.Font(resources.FontRegular, ButtonFontSize)
	if err != nil {
		return nil, err
	}

	fontTitle, err := resources.Font(resources.FontTitle, TitleFontSize)
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
	return addWidget(u, NewButton(u.image, u.fontButton, x, y, w, h, text))
}

func (u *Ui) AddStretchedButton(y int, align enum.Alignment, text string) *StretchedButton {
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
	return addWidget(u, NewLabel(u.fontTitle, x, y, text, color.ForegroundAlt))
}

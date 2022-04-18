package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/res"
	"golang.org/x/image/font"
	"image"
)

type Ui struct {
	image      *ebiten.Image
	fontButton font.Face
	fontTitle  font.Face
	fontDebug  font.Face
	widgets    []Widget
	screenSize image.Point
}

func NewUi() (*Ui, error) {
	uiImage, err := res.Image("ui")
	if err != nil {
		return nil, err
	}

	fontButton, err := res.Font(res.FontRegular, ButtonFontSize)
	if err != nil {
		return nil, err
	}

	fontTitle, err := res.Font(res.FontTitle, TitleFontSize)
	if err != nil {
		return nil, err
	}

	fontDebug, err := res.Font(res.FontDebug, DebugFontSize)
	if err != nil {
		return nil, err
	}

	return &Ui{
		image:      uiImage,
		fontButton: fontButton,
		fontTitle:  fontTitle,
		fontDebug:  fontDebug,
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

func (u *Ui) Update(screenSize image.Point) {
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

func (u *Ui) AddImageButton(image res.UiImageType, x, y, w, h int) *ImageButton {
	button := NewImageButton(u.image, x, y, w, h)
	button.SetImageType(image)
	return addWidget(u, button)
}

func (u *Ui) AddImage(src *ebiten.Image, x, y, w, h int) *Image {
	return addWidget(u, NewImage(src, x, y, w, h))
}

func (u *Ui) AddNinePatch(image res.UiImageType, x, y, w, h int) *NinePatch {
	ninePatch := NewNinePatch(u.image, x, y, w, h)
	ninePatch.SetSourceRect(res.UiImageRects[image])
	return addWidget(u, ninePatch)
}

func (u *Ui) AddRepeatImage(image res.UiImageType, x, y, w, h int) *RepeatImage {
	img := NewRepeatImage(u.image, x, y, w, h)
	img.SetImageType(image)
	return addWidget(u, img)
}

func (u *Ui) AddTitle(x, y int, text string) *Label {
	return addWidget(u, NewLabel(u.fontTitle, x, y, text, colors.ForegroundAlt))
}

func (u *Ui) AddLabel(x, y int, text string) *Label {
	return addWidget(u, NewLabel(u.fontButton, x, y, text, colors.ForegroundAlt))
}

func (u *Ui) AddDebugLabel(x, y int, text string) *Label {
	return addWidget(u, NewLabel(u.fontDebug, x, y, text, colors.ForegroundAlt))
}

func (u *Ui) AddDebugOverlay() *Overlay {
	return addWidget(u, NewOverlay(u.fontDebug))
}

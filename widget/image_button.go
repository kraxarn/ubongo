package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ImageButton is a button with an image
type ImageButton struct {
	*Button
	*Image
}

func NewImageButton(image *ebiten.Image, x, y, w, h int) *ImageButton {
	imgWidth, imgHeight := image.Size()
	return &ImageButton{
		Button: NewButton(x, y, w, h),
		Image:  NewImage(image, x, y, imgWidth, imgHeight),
	}
}

func (i *ImageButton) Update(ui *Ui) {
	i.Button.Update(ui)
}

func (i *ImageButton) Draw(dst *ebiten.Image) {
	i.Button.Draw(dst)
	i.Image.Draw(dst)
}

func (i *ImageButton) SetPosition(x, y int) {
	i.Button.SetPosition(x, y)
	i.Image.SetPosition(x, y)
}

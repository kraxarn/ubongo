package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// ImageButton is a button with an image
type ImageButton struct {
	*Button
	image *Image
}

func NewImageButton(image *ebiten.Image, x, y, w, h int) *ImageButton {
	return &ImageButton{
		Button: NewButton(x, y, w, h),
		image:  NewImage(image, x, y, w, h),
	}
}

func (i *ImageButton) Update(ui *Ui) {
	i.Button.Update(ui)
}

func (i *ImageButton) Draw(dst *ebiten.Image) {
	i.Button.Draw(dst)
	i.image.Draw(dst)
}

func (i *ImageButton) SetPosition(x, y int) {
	i.Button.SetPosition(x, y)
	i.image.SetPosition(x, y)
}

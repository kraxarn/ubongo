package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Image struct {
	image  *ebiten.Image
	rect   image.Rectangle
	scaleX float64
	scaleY float64
}

func NewImage(src *ebiten.Image, x, y, w, h int) *Image {
	imgWidth, imgHeight := src.Size()
	return &Image{
		image:  src,
		rect:   image.Rect(x, y, x+w, y+h),
		scaleX: float64(w) / float64(imgWidth),
		scaleY: float64(h) / float64(imgHeight),
	}
}

func (i *Image) Update(_ *Ui) {
	// Static image, no need to update
}

func (i *Image) Draw(dst *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(i.scaleX, i.scaleY)
	dst.DrawImage(i.image, opt)
}

package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Image struct {
	image *ebiten.Image
	rect  image.Rectangle
}

func NewImage(src *ebiten.Image, x, y int) *Image {
	width, height := src.Size()
	return &Image{
		image: src,
		rect:  image.Rect(x, y, x+width, y+height),
	}
}

func (i *Image) Update(_ *Ui) {
	// Static image, no need to update
}

func (i *Image) Draw(dst *ebiten.Image) {
	dst.DrawImage(i.image, &ebiten.DrawImageOptions{})
}

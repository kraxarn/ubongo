package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type RepeatImage struct {
	image *ebiten.Image
	rect  image.Rectangle
}

func NewRepeatImage(src *ebiten.Image, x, y, w, h int) *RepeatImage {
	return &RepeatImage{
		image: src,
		rect:  Rect(x, y, w, h),
	}
}

func (r *RepeatImage) Update(*Ui) {
}

func (r *RepeatImage) Draw(dst *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	imgWidth, imgHeight := r.image.Size()

	for x := 0; x < r.rect.Dx(); x += imgWidth {
		for y := 0; y < r.rect.Dy(); y += imgHeight {
			opt.GeoM.Reset()
			opt.GeoM.Translate(float64(x), float64(y))
			dst.DrawImage(r.image, opt)
		}
	}
}

func (r *RepeatImage) SetSize(size image.Point) {
	pos := r.rect.Min
	r.rect = Rect(pos.X, pos.Y, size.X, size.Y)
}

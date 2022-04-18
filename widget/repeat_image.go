package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type RepeatImage struct {
	*Image
	rect image.Rectangle
}

func NewRepeatImage(src *ebiten.Image, x, y, w, h int) *RepeatImage {
	// TODO: Maybe we don't want 2x scale on all images?
	img := NewImage(src, x, y, 0, 0)
	img.Rescale(2.0, 2.0)

	return &RepeatImage{
		Image: img,
		rect:  Rect(x, y, w, h),
	}
}

func (r *RepeatImage) Update(*Ui) {
}

func (r *RepeatImage) Draw(dst *ebiten.Image) {
	size := r.Image.ScaledSize()
	for x := 0; x < r.rect.Dx(); x += size.X {
		for y := 0; y < r.rect.Dy(); y += size.Y {
			r.Image.SetPosition(x, y)
			r.Image.Draw(dst)
		}
	}
}

func (r *RepeatImage) SetSize(size image.Point) {
	pos := r.rect.Min
	r.rect = Rect(pos.X, pos.Y, size.X, size.Y)
}

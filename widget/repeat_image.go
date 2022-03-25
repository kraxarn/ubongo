package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/util/vec2"
)

type RepeatImage struct {
	image    *ebiten.Image
	position vec2.Vector2[int]
	size     vec2.Vector2[int]
}

func NewRepeatImage(src *ebiten.Image, x, y, w, h int) *RepeatImage {
	return &RepeatImage{
		image:    src,
		position: vec2.New(x, y),
		size:     vec2.New(w, h),
	}
}

func (r *RepeatImage) Update(_ *Ui) {
	// Static images, no need to update
}

func (r *RepeatImage) Draw(dst *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	imgWidth, imgHeight := r.image.Size()

	for x := 0; x < r.size.X; x += imgWidth {
		for y := 0; y < r.size.Y; y += imgHeight {
			opt.GeoM.Reset()
			opt.GeoM.Translate(float64(x), float64(y))
			dst.DrawImage(r.image, opt)
		}
	}
}

func (r *RepeatImage) SetSize(size vec2.Vector2[int]) {
	r.size = size
}

package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/util"
)

type RepeatImage struct {
	image    *ebiten.Image
	position util.Vector2[int]
	size     util.Vector2[int]
}

func NewRepeatImage(src *ebiten.Image, x, y, w, h int) *RepeatImage {
	return &RepeatImage{
		image:    src,
		position: util.Vec2(x, y),
		size:     util.Vec2(w, h),
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

func (r *RepeatImage) SetSize(size util.Vector2[int]) {
	r.size = size
}

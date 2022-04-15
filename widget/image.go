package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/util/vec2"
	"image"
)

// Image is an image, or a sheet of multiple images
type Image struct {
	image      *ebiten.Image
	scaleX     float64
	scaleY     float64
	sourceRect image.Rectangle
	position   vec2.Vector2[int]
}

func NewImage(src *ebiten.Image, x, y, w, h int) *Image {
	img := &Image{
		image:      src,
		sourceRect: image.Rectangle{},
	}

	img.SetPosition(x, y)
	img.Resize(w, h)
	return img
}

func (i *Image) Update(*Ui) {
}

func (i *Image) Draw(dst *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(i.scaleX, i.scaleY)
	opt.GeoM.Translate(float64(i.position.X), float64(i.position.Y))

	var img *ebiten.Image
	if i.sourceRect.Empty() {
		img = i.image
	} else {
		img = i.image.SubImage(i.sourceRect).(*ebiten.Image)
	}

	dst.DrawImage(img, opt)
}

func (i *Image) Size() vec2.Vector2[int] {
	if i.sourceRect.Empty() {
		return vec2.New(i.image.Size())
	}
	return vec2.NewFromPoint(i.sourceRect.Size())
}

func (i *Image) Resize(w, h int) {
	size := i.Size()
	x := float64(w) / float64(size.X)
	y := float64(h) / float64(size.Y)
	i.Rescale(x, y)
}

func (i *Image) Rescale(x, y float64) {
	i.scaleX = x
	i.scaleY = y
}

func (i *Image) GetWidth() int {
	size := i.Size()
	return int(float64(size.X) * i.scaleX)
}

func (i *Image) SetWidth(w int) {
	size := i.Size()
	i.scaleX = float64(w) / float64(size.X)
	i.scaleY = float64(w) / float64(size.Y)
}

func (i *Image) GetHeight() int {
	size := i.Size()
	return int(float64(size.Y) * i.scaleY)
}

func (i *Image) GetPosition() vec2.Vector2[int] {
	return i.position
}

func (i *Image) SetPosition(x, y int) {
	i.position = vec2.New(x, y)
}

func (i *Image) SetSourceRect(rect image.Rectangle) {
	i.sourceRect = rect
}

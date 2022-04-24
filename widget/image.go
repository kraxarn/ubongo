package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/res"
	"image"
	"math"
)

// Image is an image, or a sheet of multiple images
type Image struct {
	image      *ebiten.Image
	scaleX     float64
	scaleY     float64
	theta      float64
	deg        float64
	sourceRect image.Rectangle
	position   image.Point
	Origin     image.Point
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

	if i.theta != 0 {
		originX := float64(i.Origin.X)
		originY := float64(i.Origin.Y)

		opt.GeoM.Translate(-originX, -originY)
		opt.GeoM.Rotate(i.theta)
		opt.GeoM.Translate(originX, originY)
	}

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

func (i *Image) Size() image.Point {
	if i.sourceRect.Empty() {
		return image.Pt(i.image.Size())
	}
	return i.sourceRect.Size()
}

func (i *Image) ScaledSize() image.Point {
	var w, h int

	if i.sourceRect.Empty() {
		w, h = i.image.Size()
	} else {
		size := i.sourceRect.Size()
		w = size.X
		h = size.Y
	}

	return image.Pt(int(float64(w)*i.scaleX), int(float64(h)*i.scaleY))
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

func (i *Image) GetPosition() image.Point {
	return i.position
}

func (i *Image) SetPosition(x, y int) {
	i.position = image.Pt(x, y)
}

func (i *Image) SetSourceRect(rect image.Rectangle) {
	i.sourceRect = rect
}

func (i *Image) SetImageType(image res.UiImageType) {
	i.SetSourceRect(res.UiImageRects[image])
}

func (i *Image) Rotate(deg float64) {
	i.deg += deg
	i.theta = i.deg * 2 * math.Pi / 360
}

func (i *Image) Rotation() float64 {
	return i.deg
}

package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/util"
)

type Image struct {
	image    *ebiten.Image
	scaleX   float64
	scaleY   float64
	position util.Vector2[int]
}

func NewImage(src *ebiten.Image, x, y, w, h int) *Image {
	img := &Image{
		image: src,
	}

	img.SetPosition(x, y)
	img.Resize(w, h)
	return img
}

func (i *Image) Update(_ *Ui) {
	// Static image, no need to update
}

func (i *Image) Draw(dst *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(i.scaleX, i.scaleY)
	opt.GeoM.Translate(float64(i.position.X), float64(i.position.Y))

	dst.DrawImage(i.image, opt)
}

func (i *Image) Resize(w, h int) {
	imgWidth, imgHeight := i.image.Size()
	i.scaleX = float64(w) / float64(imgWidth)
	i.scaleY = float64(h) / float64(imgHeight)
}

func (i *Image) GetWidth() int {
	w, _ := i.image.Size()
	return int(float64(w) * i.scaleX)
}

func (i *Image) SetWidth(w int) {
	imgWidth, imgHeight := i.image.Size()
	i.scaleX = float64(w) / float64(imgWidth)
	i.scaleY = float64(w) / float64(imgHeight)
}

func (i *Image) GetHeight() int {
	_, h := i.image.Size()
	return int(float64(h) * i.scaleY)
}

func (i *Image) SetPosition(x, y int) {
	i.position = util.Vec2(x, y)
}

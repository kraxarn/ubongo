package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"image"
)

type Piece struct {
	image      *widget.Image
	sourceRect image.Rectangle
	size       image.Point
}

func NewPiece(pieces *ebiten.Image, tileSize, index, x, y int) *Piece {
	sourceRect := res.PieceImageRects[index]
	img := widget.NewImage(pieces, x, y, 0, 0)
	img.SetSourceRect(sourceRect)

	scale := float64(tileSize) / res.PieceTileSize
	img.Rescale(scale, scale)

	size := sourceRect.Size()
	w := int(float64(size.X) * scale)
	h := int(float64(size.Y) * scale)

	return &Piece{
		image:      img,
		sourceRect: sourceRect,
		size:       image.Pt(w, h),
	}
}

func (p *Piece) Update() {
}

func (p *Piece) Draw(dst *ebiten.Image) {
	p.image.Draw(dst)
}

func (p *Piece) GetPosition() image.Point {
	return p.image.GetPosition()
}

func (p *Piece) SetPosition(x, y int) {
	p.image.SetPosition(x, y)
}

func (p *Piece) GetRect() image.Rectangle {
	pos := p.GetPosition()
	size := p.Size()
	return widget.Rect(pos.X, pos.Y, size.X, size.Y)
}

func (p *Piece) Size() image.Point {
	return p.size
}

package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/util/vec2"
	"github.com/kraxarn/ubongo/widget"
	"image"
)

type Piece struct {
	image      *widget.Image
	sourceRect image.Rectangle
	size       vec2.Vector2[int]
}

// PieceScale is how much to scale the image
const PieceScale = 0.4

func NewPiece(pieces *ebiten.Image, index, x, y int) *Piece {
	sourceRect := res.PieceImageRects[index]
	img := widget.NewImage(pieces, x, y, 0, 0)
	img.SetSourceRect(sourceRect)
	img.Rescale(PieceScale, PieceScale)

	size := sourceRect.Size()
	w := int(float64(size.X) * PieceScale)
	h := int(float64(size.Y) * PieceScale)

	return &Piece{
		image:      img,
		sourceRect: sourceRect,
		size:       vec2.New(w, h),
	}
}

func (p *Piece) Update() {
}

func (p *Piece) Draw(dst *ebiten.Image) {
	p.image.Draw(dst)
}

func (p *Piece) GetPosition() vec2.Vector2[int] {
	return p.image.GetPosition()
}

func (p *Piece) SetPosition(x, y int) {
	p.image.SetPosition(x, y)
}

func (p *Piece) GetRect() image.Rectangle {
	pos := p.GetPosition()
	size := p.image.Size()
	return widget.Rect(pos.X, pos.Y, size.X, size.Y)
}

func (p *Piece) Size() vec2.Vector2[int] {
	return p.size
}

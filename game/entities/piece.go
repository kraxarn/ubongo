package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/res"
	"github.com/kraxarn/ubongo/widget"
	"image"
	"math"
)

type Piece struct {
	image      *widget.Image
	sourceRect image.Rectangle
	size       image.Point
	origin     image.Point
	points     []image.Point
	rotation   float64
}

func NewPiece(pieces *ebiten.Image, tileSize, index, x, y int) *Piece {
	sourceRect := res.PieceImageRects[index]
	img := widget.NewImage(pieces, x, y, 0, 0)
	img.SetSourceRect(sourceRect)

	scale := float64(tileSize) / res.PieceTileSize
	img.Rescale(scale, scale)

	points := PieceTiles(index)
	origin := PieceOrigin(points)
	img.Origin = origin.Mul(tileSize)

	size := sourceRect.Size()
	w := int(float64(size.X) * scale)
	h := int(float64(size.Y) * scale)

	return &Piece{
		image:      img,
		sourceRect: sourceRect,
		size:       image.Pt(w, h),
		origin:     origin,
		points:     points,
	}
}

func (p *Piece) Update() {
	current := p.image.Rotation()
	if p.rotation == current {
		return
	}

	pos := math.Abs(p.rotation - current)
	speed := (90 - math.Abs(45-pos)) / 7.5
	step := math.Min(pos, speed)

	if p.rotation > current {
		p.image.Rotate(step)
	} else {
		p.image.Rotate(-step)
	}

	if p.image.Rotation() == p.rotation {
		p.rotate()
	}
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

func (p *Piece) Origin() image.Point {
	return p.image.Origin
}

func (p *Piece) Rect() image.Rectangle {
	pos := p.GetPosition()
	size := p.Size()
	return widget.Rect(pos.X, pos.Y, size.X, size.Y)
}

func (p *Piece) Size() image.Point {
	return p.size
}

func (p *Piece) Points() []image.Point {
	return p.points
}

func (p *Piece) Rotate(deg float64) {
	if p.rotation != p.image.Rotation() {
		return
	}

	p.rotation += deg
}

func (p *Piece) rotate() {
	for i, point := range p.points {
		p.points[i].X = -point.Y + p.origin.X
		p.points[i].Y = point.X + p.origin.Y
	}
}

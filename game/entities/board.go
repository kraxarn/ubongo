package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/widget"
	"image"
)

type TileType uint8

const (
	// None is a tile outside the board
	None TileType = iota

	// Empty is an empty tile inside the board
	Empty TileType = iota

	// Filled is a filled tile inside the board
	Filled TileType = iota
)

// tileCount is the number of tiles horizontally and vertically
const tileCount = 8

const (
	lineWidth   = 4
	linePadding = widget.ScreenPadding / 3
)

type Board struct {
	tiles      [tileCount][tileCount]TileType
	rect       image.Rectangle
	tileSize   int
	background *ebiten.Image
	line       *ebiten.Image
}

func NewBoard(x, y, w, h int) *Board {
	background := ebiten.NewImage(w, h)
	background.Fill(colors.BackgroundBoard)

	line := ebiten.NewImage(lineWidth, h-linePadding*2)
	line.Fill(colors.BorderBoard)

	return &Board{
		tiles:      [8][8]TileType{},
		rect:       widget.Rect(x, y, w, h),
		tileSize:   w / tileCount,
		background: background,
		line:       line,
	}
}

func (b *Board) Draw(dst *ebiten.Image) {
	// Background
	opt := &ebiten.DrawImageOptions{}
	pos := b.rect.Min
	opt.GeoM.Translate(float64(pos.X), float64(pos.Y))
	dst.DrawImage(b.background, opt)

	// Lines
	for i := 1; i < tileCount; i++ {
		// Vertical
		opt.GeoM.Reset()
		opt.GeoM.Translate(float64(pos.X+(i*b.tileSize)), float64(pos.Y+linePadding))
		dst.DrawImage(b.line, opt)

		// Horizontal
		opt.GeoM.Reset()
		opt.GeoM.Rotate(-1.57)
		opt.GeoM.Translate(float64(pos.X+linePadding), float64(pos.Y+(i*b.tileSize)))
		dst.DrawImage(b.line, opt)
	}
}

func (b *Board) TileSize() int {
	return b.tileSize
}

func (b *Board) Rect() image.Rectangle {
	return b.rect
}

func (b *Board) Position() image.Point {
	return b.rect.Min
}

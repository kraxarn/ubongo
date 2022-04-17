package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/widget"
	"image"
	"math/rand"
)

// PieceCount is the total amount of pieces
const PieceCount = 5

// tileCount is the number of tiles horizontally and vertically
const tileCount = 8

const (
	lineWidth   = 4
	linePadding = widget.ScreenPadding / 3
)

type Board struct {
	tiles      []image.Point
	rect       image.Rectangle
	tileSize   int
	background *ebiten.Image
	line       *ebiten.Image
}

func NewBoard(pieces [PieceCount]*Piece, x, y, w, h int) *Board {
	tileSize := TileSize(w)
	background := ebiten.NewImage(tileSize, tileSize)
	background.Fill(colors.BackgroundBoard)

	line := ebiten.NewImage(lineWidth, h-linePadding*2)
	line.Fill(colors.BorderBoard)

	return &Board{
		tiles:      generateBoard(pieces),
		rect:       widget.Rect(x, y, w, h),
		tileSize:   tileSize,
		background: background,
		line:       line,
	}
}

func (b *Board) Draw(dst *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	pos := b.Position()

	// Tiles
	for _, tile := range b.tiles {
		opt.GeoM.Reset()
		opt.GeoM.Translate(float64(pos.X+(tile.X*b.tileSize)), float64(pos.Y+(tile.Y*b.tileSize)))
		dst.DrawImage(b.background, opt)
	}

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

func TileSize(boardSize int) int {
	return boardSize / tileCount
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

func generateBoard(pieces [PieceCount]*Piece) []image.Point {
	var tiles []image.Point

	var shuffled [PieceCount]*Piece
	copy(shuffled[0:], pieces[0:])
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	for i, piece := range shuffled {
		tileData := PieceTiles(piece.Index())

		// Place first in center
		if i == 0 {
			tileY := PieceHeight(tileData)
			tileX := PieceWidth(tileData)
			centerX := tileCount/2 - tileX/2
			centerY := tileCount/2 - tileY/2
			for _, point := range tileData {
				tiles = append(tiles, image.Pt(centerX, centerY).Add(point))
			}
		}
	}

	return tiles
}

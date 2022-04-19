package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/game/colors"
	"github.com/kraxarn/ubongo/widget"
	"image"
	"math"
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
		opt.GeoM.Rotate(-90.0 * 2 * math.Pi / 360)
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

func (b *Board) AllTilesFilled(pieces [PieceCount]*Piece) bool {
	var pieceTiles []image.Point
	for _, piece := range pieces {
		offset := piece.GetPosition().Sub(b.Position()).Div(b.TileSize())
		for _, pieceTile := range PieceTiles(piece.Index()) {
			pieceTiles = append(pieceTiles, offset.Add(pieceTile))
		}
	}

	for _, tile := range b.tiles {
		if !containsPoint(pieceTiles, tile) {
			return false
		}
	}
	return true
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
			tileSize := PieceSize(tileData)
			centerX := tileCount/2 - tileSize.X/2
			centerY := tileCount/2 - tileSize.Y/2
			for _, point := range tileData {
				tiles = append(tiles, image.Pt(centerX, centerY).Add(point))
			}
			continue
		}

		// Find best position for the rest
		max := 0
		var results []image.Point

		for y := 0; y < tileCount; y++ {
			for x := 0; x < tileCount; x++ {
				offset := image.Pt(x, y)
				if anyOverflow(tileData, offset) || !allTilesFree(tiles, tileData, offset) {
					continue
				}
				count := adjacentTileCount(tiles, tileData, offset)
				// New max
				if count > max {
					max = count
					results = []image.Point{
						offset,
					}
					continue
				}
				// Same as current max
				if count == max {
					results = append(results, offset)
				}
			}
		}

		// Tile doesn't fit anywhere
		if len(results) <= 0 {
			fmt.Println("No results found for tile", i)
			continue
		}

		// Pick a random result and add
		result := results[rand.Intn(len(results))]
		for _, point := range tileData {
			tiles = append(tiles, result.Add(point))
		}
	}

	return tiles
}

func allTilesFree(tiles, piece []image.Point, offset image.Point) bool {
	for _, shapePoint := range piece {
		current := offset.Add(shapePoint)
		for _, tilePoint := range tiles {
			if tilePoint.Eq(current) {
				return false
			}
		}
	}
	return true
}

func anyOverflow(piece []image.Point, offset image.Point) bool {
	for _, point := range piece {
		current := offset.Add(point)
		if current.X < 0 || current.Y < 0 || current.X >= tileCount || current.Y >= tileCount {
			return true
		}
	}
	return false
}

func adjacentTileCount(tiles, piece []image.Point, offset image.Point) int {
	count := 0

	for _, point := range piece {
		// Left
		left := point.Add(image.Pt(-1, 0))
		if !containsPoint(piece, left) && containsPoint(tiles, offset.Add(left)) {
			count++
		}
		// Top
		top := point.Add(image.Pt(0, -1))
		if !containsPoint(piece, top) && containsPoint(tiles, offset.Add(top)) {
			count++
		}
		// Right
		right := point.Add(image.Pt(1, 0))
		if !containsPoint(piece, right) && containsPoint(tiles, offset.Add(right)) {
			count++
		}
		// Bottom
		bottom := point.Add(image.Pt(0, 1))
		if !containsPoint(piece, bottom) && containsPoint(tiles, offset.Add(bottom)) {
			count++
		}
	}

	return count
}

// containsPoint checks if point exists in tiles
func containsPoint(tiles []image.Point, point image.Point) bool {
	for _, tile := range tiles {
		if tile.Eq(point) {
			return true
		}
	}
	return false
}

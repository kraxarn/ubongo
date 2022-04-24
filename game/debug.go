package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

func debugImage(size int) *ebiten.Image {
	img := ebiten.NewImage(size, size)
	img.Fill(color.RGBA{
		R: 0x7f,
		G: 0x7f,
		B: 0x7f,
		A: 0x7f,
	})
	return img
}

func debugTile(tileSize int) *ebiten.Image {
	return debugImage(int(float64(tileSize) * 0.75))
}

func debugPoint() *ebiten.Image {
	return debugImage(3)
}

func (s *SceneGame) drawDebugTile(screen *ebiten.Image) {
	opt := ebiten.DrawImageOptions{}
	tileSize := s.board.TileSize()
	margin := image.Pt(1, 1).Mul(int(float64(tileSize) * 0.125))

	for _, piece := range s.pieces {
		offset := piece.GetPosition()
		for _, pieceTile := range piece.Points() {
			pos := offset.Add(pieceTile.Mul(tileSize)).Add(margin)
			opt.GeoM.Reset()
			opt.GeoM.Translate(float64(pos.X), float64(pos.Y))
			screen.DrawImage(s.debugTile, &opt)
		}
	}
}

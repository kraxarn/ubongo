package res

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:embed image/*.png
var images embed.FS

type UiImageType uint8

const (
	UiButton        UiImageType = iota
	UiButtonPressed UiImageType = iota
	PanelBackground UiImageType = iota
	MusicOff        UiImageType = iota
	MusicOn         UiImageType = iota
)

var UiImageRects = map[UiImageType]image.Rectangle{
	UiButton:        image.Rect(0, 0, 45, 49),
	UiButtonPressed: image.Rect(45, 0, 90, 45),
	PanelBackground: image.Rect(0, 49, 64, 113),
	MusicOff:        image.Rect(0, 113, 50, 163),
	MusicOn:         image.Rect(50, 113, 100, 163),
}

// PieceTileSize is the width/height of each tile in a piece
const PieceTileSize = 128

var PieceImageRects = []image.Rectangle{
	// Row 1
	image.Rect(PieceTileSize*0, PieceTileSize*0, PieceTileSize*3, PieceTileSize*3),
	image.Rect(PieceTileSize*3, PieceTileSize*0, PieceTileSize*5, PieceTileSize*3),
	image.Rect(PieceTileSize*5, PieceTileSize*0, PieceTileSize*7, PieceTileSize*2),
	image.Rect(PieceTileSize*7, PieceTileSize*0, PieceTileSize*8, PieceTileSize*2),
	// Row 2
	image.Rect(PieceTileSize*0, PieceTileSize*3, PieceTileSize*4, PieceTileSize*4),
	image.Rect(PieceTileSize*4, PieceTileSize*3, PieceTileSize*8, PieceTileSize*5),
	// Row 3
	image.Rect(PieceTileSize*0, PieceTileSize*5, PieceTileSize*3, PieceTileSize*7),
	image.Rect(PieceTileSize*3, PieceTileSize*5, PieceTileSize*6, PieceTileSize*7),
	image.Rect(PieceTileSize*6, PieceTileSize*5, PieceTileSize*9, PieceTileSize*6),
	// Row 4
	image.Rect(PieceTileSize*0, PieceTileSize*7, PieceTileSize*4, PieceTileSize*9),
	image.Rect(PieceTileSize*4, PieceTileSize*7, PieceTileSize*7, PieceTileSize*9),
	image.Rect(PieceTileSize*7, PieceTileSize*7, PieceTileSize*9, PieceTileSize*9),
}

func decodeImage(name string) (*ebiten.Image, error) {
	reader, err := images.Open(fmt.Sprintf("image/%s.png", name))
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func Image(name string) (*ebiten.Image, error) {
	return decodeImage(name)
}

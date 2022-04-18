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
	BackgroundButton UiImageType = iota
	BackgroundDialog UiImageType = iota
	BackgroundPanel  UiImageType = iota
	Forward          UiImageType = iota
	MusicOff         UiImageType = iota
	MusicOn          UiImageType = iota
	Pause            UiImageType = iota
	Return           UiImageType = iota
	Stripes          UiImageType = iota
)

// UiImageRects are coordinates for ui.png from ui.atlas
// TODO: Auto parse from atlas file
var UiImageRects = map[UiImageType]image.Rectangle{
	BackgroundButton: r(378, 44, 40, 20),
	BackgroundDialog: r(64, 0, 64, 64),
	BackgroundPanel:  r(0, 0, 64, 64),
	Forward:          r(278, 14, 50, 50),
	MusicOff:         r(178, 14, 50, 50),
	MusicOn:          r(328, 14, 50, 50),
	Pause:            r(228, 14, 50, 50),
	Return:           r(128, 14, 50, 50),
	Stripes:          r(378, 16, 28, 28),
}

// PieceTileSize is the width/height of each tile in a piece
const PieceTileSize = 128

// PieceCount is the total number of pieces
const PieceCount = 12

var PieceImageRects = [PieceCount]image.Rectangle{
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

func r(x, y, w, h int) image.Rectangle {
	return image.Rect(x, y, x+w, y+h)
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

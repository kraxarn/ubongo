package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

//go:embed image/pieces.png
var pieces []byte

//go:embed image/ui.png
var ui []byte

type UiImageType uint8

const (
	UiButton        UiImageType = iota
	UiButtonPressed UiImageType = iota
)

var UiImageRects = map[UiImageType]image.Rectangle{
	UiButton:        image.Rect(0, 0, 45, 49),
	UiButtonPressed: image.Rect(45, 0, 90, 45),
}

func decode(data []byte) (*ebiten.Image, error) {
	reader := bytes.NewReader(data)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func ImageUi() (*ebiten.Image, error) {
	return decode(ui)
}

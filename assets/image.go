package assets

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
)

var UiImageRects = map[UiImageType]image.Rectangle{
	UiButton:        image.Rect(0, 0, 45, 49),
	UiButtonPressed: image.Rect(45, 0, 90, 45),
}

func decode(name string) (*ebiten.Image, error) {
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

// Deprecated: Use Image with "ui" instead
func ImageUi() (*ebiten.Image, error) {
	return decode("ui")
}

func Image(name string) (*ebiten.Image, error) {
	return decode(name)
}

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
)

var UiImageRects = map[UiImageType]image.Rectangle{
	UiButton:        image.Rect(0, 0, 45, 49),
	UiButtonPressed: image.Rect(45, 0, 90, 45),
}

var PieceImageRects = []image.Rectangle{
	// Row 1
	image.Rect(0, 0, 384, 384),
	image.Rect(384, 0, 640, 384),
	image.Rect(640, 0, 896, 256),
	image.Rect(896, 0, 1024, 256),
	// Row 2
	image.Rect(0, 384, 512, 512),
	image.Rect(512, 384, 1024, 640),
	// Row 3
	image.Rect(0, 640, 384, 896),
	image.Rect(384, 640, 768, 896),
	image.Rect(768, 640, 1152, 768),
	// Row 4
	image.Rect(0, 896, 512, 1152),
	image.Rect(512, 896, 896, 1152),
	image.Rect(896, 896, 1152, 1152),
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

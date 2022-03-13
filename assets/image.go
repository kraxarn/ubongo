package assets

import (
	_ "embed"
	"image"
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
	UiButton:        image.Rect(0, 0, 16, 16),
	UiButtonPressed: image.Rect(16, 0, 32, 16),
}

package assets

import (
	"embed"
	"fmt"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed font/*.ttf
var fonts embed.FS

type FontType uint8

const (
	FontDebug   FontType = iota
	FontMenu    FontType = iota
	FontSubmenu FontType = iota
)

func fontPath(font FontType) (string, error) {
	switch font {
	case FontDebug:
		return "font/debug.ttf", nil

	case FontMenu:
		return "font/menu.ttf", nil

	case FontSubmenu:
		return "font/submenu.ttf", nil

	default:
		return "", fmt.Errorf("invalid font: %d", font)
	}
}

func fontData(font FontType) ([]byte, error) {
	path, err := fontPath(font)
	if err != nil {
		return nil, err
	}

	return fonts.ReadFile(path)
}

func Font(fontType FontType, fontSize float64) (font.Face, error) {
	data, err := fontData(fontType)
	if err != nil {
		return nil, err
	}

	font, err := opentype.Parse(data)
	if err != nil {
		return nil, err
	}

	return opentype.NewFace(font, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

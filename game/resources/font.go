package resources

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
	FontRegular FontType = iota
	FontTitle   FontType = iota
	FontDebug   FontType = iota
)

func fontPath(font FontType) (string, error) {
	switch font {
	case FontRegular:
		return "font/regular.ttf", nil

	case FontTitle:
		return "font/title.ttf", nil

	case FontDebug:
		return "font/debug.ttf", nil

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

	ttf, err := opentype.Parse(data)
	if err != nil {
		return nil, err
	}

	return opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

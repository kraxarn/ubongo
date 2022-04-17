package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func Rect(x, y, w, h int) image.Rectangle {
	return image.Rect(x, y, x+w, y+h)
}

// TouchPositions gets all currently active positions (mouse or touch)
func TouchPositions() []image.Point {
	var positions []image.Point

	// Mouse
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		positions = append(positions, image.Pt(ebiten.CursorPosition()))
	}

	// Touch
	for _, id := range ebiten.AppendTouchIDs([]ebiten.TouchID{}) {
		positions = append(positions, image.Pt(ebiten.TouchPosition(id)))
	}

	return positions
}

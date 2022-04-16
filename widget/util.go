package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/util/vec2"
	"image"
)

func Rect(x, y, w, h int) image.Rectangle {
	return image.Rect(x, y, x+w, y+h)
}

// TouchPositions gets all currently active positions (mouse or touch)
func TouchPositions() []vec2.Vector2[int] {
	var positions []vec2.Vector2[int]

	// Mouse
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		positions = append(positions, vec2.New(ebiten.CursorPosition()))
	}

	// Touch
	for _, id := range ebiten.AppendTouchIDs([]ebiten.TouchID{}) {
		positions = append(positions, vec2.New(ebiten.TouchPosition(id)))
	}

	return positions
}

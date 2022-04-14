package widget

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kraxarn/ubongo/util/vec2"
	"image"
)

// Button is an abstract widget that can be pressed
type Button struct {
	rect      image.Rectangle
	mouseDown bool
	isPressed bool
	onPressed func(b *Button)
}

func NewButton(x, y, w, h int) *Button {
	return &Button{
		rect:      Rect(x, y, w, h),
		mouseDown: false,
		isPressed: false,
		onPressed: nil,
	}
}

func (b *Button) SetOnPressed(pressed func(b *Button)) {
	b.onPressed = pressed
}

func (b *Button) Update(*Ui) {
	var positions []vec2.Vector2[int]
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		positions = append(positions, vec2.New(ebiten.CursorPosition()))
	} else {
		for _, id := range ebiten.AppendTouchIDs([]ebiten.TouchID{}) {
			positions = append(positions, vec2.New(ebiten.TouchPosition(id)))
		}
	}

	if len(positions) > 0 {
		b.mouseDown = false
		for _, pos := range positions {
			if b.rect.Min.X <= pos.X && pos.X < b.rect.Max.X && b.rect.Min.Y <= pos.Y && pos.Y < b.rect.Max.Y {
				b.mouseDown = true
				break
			}
		}
	} else {
		if b.mouseDown {
			if b.onPressed != nil {
				b.onPressed(b)
			}
		}
		b.mouseDown = false
	}
}

func (b *Button) Draw(*ebiten.Image) {
}

func (b *Button) SetPosition(x, y int) {
	size := b.rect.Size()
	b.rect = Rect(x, y, size.X, size.Y)
}

func (b *Button) Size() vec2.Vector2[int] {
	return vec2.NewFromPoint(b.rect.Size())
}

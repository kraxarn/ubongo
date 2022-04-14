package vec2

import "image"

type Vector2[V any] struct {
	X V
	Y V
}

func New[V any](x, y V) Vector2[V] {
	return Vector2[V]{
		X: x,
		Y: y,
	}
}

func NewFromPoint(point image.Point) Vector2[int] {
	return New(point.X, point.Y)
}

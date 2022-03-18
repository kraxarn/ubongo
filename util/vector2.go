package util

type Vector2[V any] struct {
	X V
	Y V
}

func Vec2[V any](x, y V) Vector2[V] {
	return Vector2[V]{
		X: x,
		Y: y,
	}
}

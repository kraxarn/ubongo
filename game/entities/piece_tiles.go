package entities

import (
	"image"
	"math"
)

func PieceTiles(index int) []image.Point {
	switch index {
	// 1 : 1
	case 0:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(1, 1),
			image.Pt(1, 2),
			image.Pt(2, 2),
		}

	// 1 : 2
	case 1:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(0, 1),
			image.Pt(0, 2),
		}

	// 1 : 3
	case 2:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(0, 1),
			image.Pt(1, 1),
		}

	// 1 : 4
	case 3:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(0, 1),
		}

	// 2 : 1
	case 4:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(2, 0),
			image.Pt(3, 0),
		}

	// 2 : 2
	case 5:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(2, 0),
			image.Pt(2, 1),
			image.Pt(3, 0),
		}

	// 3 : 1
	case 6:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(2, 0),
			image.Pt(1, 1),
			image.Pt(2, 1),
		}

	// 3 : 2
	case 7:
		return []image.Point{
			image.Pt(1, 0),
			image.Pt(2, 0),
			image.Pt(0, 1),
			image.Pt(1, 1),
		}

	// 3 : 3
	case 8:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(2, 0),
		}

	// 4 : 1
	case 9:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(2, 0),
			image.Pt(3, 0),
			image.Pt(3, 1),
		}

	// 4 : 2
	case 10:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(1, 1),
			image.Pt(2, 0),
		}

	// 4 : 3
	case 11:
		return []image.Point{
			image.Pt(0, 0),
			image.Pt(1, 0),
			image.Pt(0, 1),
		}

	default:
		return []image.Point{}
	}
}

func PieceSize(points []image.Point) image.Point {
	min := image.Point{}
	max := image.Point{}

	for _, point := range points {
		if point.X < min.X {
			min.X = point.X
		}
		if point.X > max.X {
			max.X = point.X
		}
		if point.Y < min.Y {
			min.Y = point.Y
		}
		if point.Y > max.Y {
			max.Y = point.Y
		}
	}

	max.Sub(min).Add(image.Pt(1, 1))
	return max
}

func PieceOrigin(points []image.Point) image.Point {
	size := PieceSize(points)
	x := math.Ceil(float64(size.X) / 2.0)
	y := math.Ceil(float64(size.Y) / 2.0)
	return image.Pt(int(x), int(y))
}

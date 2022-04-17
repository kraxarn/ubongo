package entities

import (
	"image"
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

func PieceWidth(points []image.Point) int {
	max := 0
	for _, point := range points {
		if point.X > max {
			max = point.X
		}
	}
	return max + 1
}

func PieceHeight(points []image.Point) int {
	max := 0
	for _, point := range points {
		if point.Y > max {
			max = point.Y
		}
	}
	return max + 1
}

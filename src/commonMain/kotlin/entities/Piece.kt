package entities

import com.soywiz.korma.geom.IPointInt
import com.soywiz.korma.geom.PointInt

enum class Piece(val points: Array<IPointInt>)
{
	ONE(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(1, 1),
			PointInt(1, 2),
			PointInt(2, 2),
		),
	),

	TWO(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(0, 1),
			PointInt(0, 2),
		),
	),

	THREE(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(0, 1),
			PointInt(1, 1),
		),
	),

	FOUR(
		arrayOf(
			PointInt(0, 0),
			PointInt(0, 1),
		),
	),

	FIVE(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(2, 0),
			PointInt(3, 0),
		),
	),

	SIX(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(2, 0),
			PointInt(2, 1),
			PointInt(3, 0),
		),
	),

	SEVEN(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(2, 0),
			PointInt(1, 1),
			PointInt(2, 1),
		),
	),

	EIGHT(
		arrayOf(
			PointInt(1, 0),
			PointInt(2, 0),
			PointInt(0, 1),
			PointInt(1, 1),
		),
	),

	NINE(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(2, 0),
		),
	),

	TEN(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(2, 0),
			PointInt(3, 0),
			PointInt(3, 1),
		),
	),

	ELEVEN(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(1, 1),
			PointInt(2, 0),
		),
	),

	TWELVE(
		arrayOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(0, 1),
		),
	),
}
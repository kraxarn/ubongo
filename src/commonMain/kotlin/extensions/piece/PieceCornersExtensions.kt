package extensions.piece

import com.soywiz.korma.geom.PointInt
import enums.PieceShape

private val PieceShape.cornerSize
	get() = when (this)
	{
		PieceShape.I1 -> 2
		PieceShape.I2 -> 3
		PieceShape.I3 -> 4

		PieceShape.L1 -> 1
		PieceShape.L2 -> 2
		PieceShape.L3 -> 3

		PieceShape.T1 -> 1
		PieceShape.T2 -> 2

		else -> 0
	}

val PieceShape.corners: List<PointInt>
	get() = when (this)
	{
		PieceShape.I1, PieceShape.I2, PieceShape.I3 -> listOf(
			PointInt(0, 0),
			PointInt(this.cornerSize, 0),
			PointInt(this.cornerSize, 1),
			PointInt(0, 1),
			PointInt(0, 0),
		)

		PieceShape.L1, PieceShape.L2, PieceShape.L3 ->
			listOf(
				PointInt(0, 0),
				PointInt(1, 0),
				PointInt(1, this.cornerSize),
				PointInt(2, this.cornerSize),
				PointInt(2, this.cornerSize + 1),
				PointInt(0, this.cornerSize + 1),
				PointInt(0, 0),
			)

		PieceShape.O1 -> listOf(
			PointInt(0, 0),
			PointInt(2, 0),
			PointInt(2, 2),
			PointInt(0, 2),
			PointInt(0, 0),
		)

		PieceShape.P1 -> listOf(
			PointInt(0, 0),
			PointInt(2, 0),
			PointInt(2, 2),
			PointInt(1, 2),
			PointInt(1, 3),
			PointInt(0, 3),
			PointInt(0, 0),
		)

		PieceShape.T1, PieceShape.T2 -> listOf(
			PointInt(0, 0),
			PointInt(this.cornerSize + 2, 0),
			PointInt(this.cornerSize + 2, 1),
			PointInt(this.cornerSize + 1, 1),
			PointInt(this.cornerSize + 1, 2),
			PointInt(this.cornerSize, 2),
			PointInt(this.cornerSize, 1),
			PointInt(0, 1),
			PointInt(0, 0),
		)

		PieceShape.Z1 -> listOf(
			PointInt(0, 0),
			PointInt(1, 0),
			PointInt(1, 1),
			PointInt(1, 1),
			PointInt(2, 1),
			PointInt(2, 3),
			PointInt(1, 3),
			PointInt(1, 2),
			PointInt(0, 2),
			PointInt(0, 0),
		)

		PieceShape.Z2 -> listOf(
			PointInt(0, 0),
			PointInt(2, 0),
			PointInt(2, 2),
			PointInt(3, 2),
			PointInt(3, 3),
			PointInt(1, 3),
			PointInt(1, 1),
			PointInt(0, 1),
			PointInt(0, 0),
		)
	}
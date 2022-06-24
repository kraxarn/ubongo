package extensions.piece

import com.soywiz.korma.geom.PointInt
import entities.Piece

private val Piece.cornerSize
	get() = when (this)
	{
		Piece.I1 -> 2
		Piece.I2 -> 3
		Piece.I3 -> 4

		Piece.L1 -> 1
		Piece.L2 -> 2
		Piece.L3 -> 3

		Piece.T1 -> 1
		Piece.T2 -> 2

		else -> 0
	}

val Piece.corners: List<PointInt>
	get() = when (this)
	{
		Piece.I1, Piece.I2, Piece.I3 -> listOf(
			PointInt(0, 0),
			PointInt(this.cornerSize, 0),
			PointInt(this.cornerSize, 1),
			PointInt(0, 1),
			PointInt(0, 0),
		)

		Piece.L1, Piece.L2, Piece.L3 ->
			listOf(
				PointInt(0, 0),
				PointInt(1, 0),
				PointInt(1, this.cornerSize),
				PointInt(2, this.cornerSize),
				PointInt(2, this.cornerSize + 1),
				PointInt(0, this.cornerSize + 1),
				PointInt(0, 0),
			)

		Piece.O1 -> listOf(
			PointInt(0, 0),
			PointInt(2, 0),
			PointInt(2, 2),
			PointInt(0, 2),
			PointInt(0, 0),
		)

		Piece.P1 -> listOf(
			PointInt(0, 0),
			PointInt(2, 0),
			PointInt(2, 2),
			PointInt(1, 2),
			PointInt(1, 3),
			PointInt(0, 3),
			PointInt(0, 0),
		)

		Piece.T1, Piece.T2 -> listOf(
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

		Piece.Z1 -> listOf(
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

		Piece.Z2 -> listOf(
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
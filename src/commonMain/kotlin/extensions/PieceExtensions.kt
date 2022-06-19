package extensions

import com.soywiz.korim.color.Colors
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.minus
import com.soywiz.korma.geom.plus
import entities.Piece

val Piece.size: PointInt
	get()
	{
		val min = PointInt()
		val max = PointInt()

		for (point in this.points)
		{
			if (point.x < min.x) min.x = point.x
			if (point.x > max.x) max.x = point.x
			if (point.y < min.y) min.y = point.y
			if (point.y > max.y) max.y = point.y
		}

		return max - min + PointInt(1, 1)
	}

val Piece.color
	get() = when (this)
	{
		Piece.I1 -> Colors["#9c27b0"] // Purple
		Piece.I2 -> Colors["#2196f3"] // Dark blue
		Piece.I3 -> Colors["#795548"] // Brown
		Piece.L1 -> Colors["#009688"] // Teal
		Piece.L2 -> Colors["#8bc34a"] // Light green
		Piece.L3 -> Colors["#e91e63"] // Pink
		Piece.O1 -> Colors["#f44336"] // Red
		Piece.P1 -> Colors["#4caf50"] // Green
		Piece.T1 -> Colors["#ffeb3b"] // Yellow
		Piece.T2 -> Colors["#a1887f"] // Light brown
		Piece.Z1 -> Colors["#3f51b5"] // Indigo
		Piece.Z2 -> Colors["#fbc02d"] // Dark yellow
	}
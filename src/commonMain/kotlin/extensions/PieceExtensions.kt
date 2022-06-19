package extensions

import com.soywiz.korim.color.Colors
import entities.Piece

val Piece.size get() = this.shape.size2

val Piece.color
	get() = when (this)
	{
		Piece.I1 -> Colors["#9c27b0"] // Purple
		Piece.I2 -> Colors["#00bcd4"] // Cyan (Dark blue)
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

val Piece.displayName
	get() = when (this)
	{
		Piece.I1 -> "I (small)"
		Piece.I2 -> "I (medium)"
		Piece.I3 -> "I (large)"
		Piece.L1 -> "L (small)"
		Piece.L2 -> "L (medium)"
		Piece.L3 -> "L (large)"
		Piece.O1 -> "O"
		Piece.P1 -> "P"
		Piece.T1 -> "T (small)"
		Piece.T2 -> "T (large"
		Piece.Z1 -> "Z (small)"
		Piece.Z2 -> "Z (large)"
	}
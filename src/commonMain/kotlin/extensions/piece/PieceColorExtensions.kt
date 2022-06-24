package extensions.piece

import com.soywiz.korim.color.Colors
import entities.Piece

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

val Piece.borderColor
	get() = when (this)
	{
		Piece.I1 -> Colors["#6a0080"]
		Piece.I2 -> Colors["#008ba3"]
		Piece.I3 -> Colors["#4b2c20"]
		Piece.L1 -> Colors["#00675b"]
		Piece.L2 -> Colors["#5a9216"]
		Piece.L3 -> Colors["#b0003a"]
		Piece.O1 -> Colors["#ba000d"]
		Piece.P1 -> Colors["#087f23"]
		Piece.T1 -> Colors["#c8b900"]
		Piece.T2 -> Colors["#725b53"]
		Piece.Z1 -> Colors["#002984"]
		Piece.Z2 -> Colors["#c49000"]
	}
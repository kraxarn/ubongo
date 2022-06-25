package extensions.piece

import com.soywiz.korim.color.Colors
import enums.PieceShape

val PieceShape.color
	get() = when (this)
	{
		PieceShape.I1 -> Colors["#9c27b0"] // Purple
		PieceShape.I2 -> Colors["#00bcd4"] // Cyan (Dark blue)
		PieceShape.I3 -> Colors["#795548"] // Brown
		PieceShape.L1 -> Colors["#009688"] // Teal
		PieceShape.L2 -> Colors["#8bc34a"] // Light green
		PieceShape.L3 -> Colors["#e91e63"] // Pink
		PieceShape.O1 -> Colors["#f44336"] // Red
		PieceShape.P1 -> Colors["#4caf50"] // Green
		PieceShape.T1 -> Colors["#ffeb3b"] // Yellow
		PieceShape.T2 -> Colors["#a1887f"] // Light brown
		PieceShape.Z1 -> Colors["#3f51b5"] // Indigo
		PieceShape.Z2 -> Colors["#fbc02d"] // Dark yellow
	}

val PieceShape.borderColor
	get() = when (this)
	{
		PieceShape.I1 -> Colors["#6a0080"]
		PieceShape.I2 -> Colors["#008ba3"]
		PieceShape.I3 -> Colors["#4b2c20"]
		PieceShape.L1 -> Colors["#00675b"]
		PieceShape.L2 -> Colors["#5a9216"]
		PieceShape.L3 -> Colors["#b0003a"]
		PieceShape.O1 -> Colors["#ba000d"]
		PieceShape.P1 -> Colors["#087f23"]
		PieceShape.T1 -> Colors["#c8b900"]
		PieceShape.T2 -> Colors["#725b53"]
		PieceShape.Z1 -> Colors["#002984"]
		PieceShape.Z2 -> Colors["#c49000"]
	}
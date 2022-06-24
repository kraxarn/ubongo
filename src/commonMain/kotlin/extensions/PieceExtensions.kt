package extensions

import com.soywiz.korge.input.Gestures
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.color.Colors
import com.soywiz.korim.vector.StrokeInfo
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.vector.LineCap
import com.soywiz.korma.geom.vector.lineTo
import com.soywiz.korma.geom.vector.moveTo
import com.soywiz.korma.geom.vector.rect
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

val Piece.points: Sequence<PointInt>
	get() = sequence {
		for (x in 0 until this@points.shape.width)
		{
			for (y in 0 until this@points.shape.height)
			{
				if (this@points.shape[x, y]) yield(PointInt(x, y))
			}
		}
	}

val Piece.corners: Sequence<PointInt>
	get() = sequence {
		val shape = this@corners.shape

		// Start at 0,0 and find next empty right tile
		yield(PointInt(0, 0))

		Gestures.Direction

		// Right-most
		for (x in shape.width - 1 downTo 0)
		{
			if ((shape[x, 0] && (x == shape.width - 1 || !shape[x + 1, 0])))
			{
				yield(PointInt(x + 1, 0))
				break
			}
		}
	}

val Piece.bitmap: Bitmap
	get()
	{
		val scale = 64
		val borderSize = 6
		val borderCap = LineCap.ROUND
		val size = this.shape.size2 * scale

		return NativeImage(size.x + borderSize * 2, size.y + borderSize * 2).context2d {
			fill(this@bitmap.color) {
				for (point in this@bitmap.points)
				{
					rect(borderSize + point.x * scale, borderSize + point.y * scale, scale, scale)
				}
			}
			stroke(
				this@bitmap.borderColor,
				StrokeInfo(borderSize.toDouble(), startCap = borderCap, endCap = borderCap)
			) {
				moveTo(borderSize, borderSize)

				for (corner in this@bitmap.corners)
				{
					lineTo(corner.x * scale + borderSize, corner.y * scale + borderSize)
				}
			}
		}
	}
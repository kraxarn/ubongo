package extensions.piece

import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.vector.StrokeInfo
import com.soywiz.korma.geom.vector.LineCap
import com.soywiz.korma.geom.vector.lineTo
import com.soywiz.korma.geom.vector.moveTo
import com.soywiz.korma.geom.vector.rect
import entities.Piece
import extensions.size2
import extensions.times

val Piece.bitmap: Bitmap
	get()
	{
		val scale = 64
		val borderSize = 4
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

				for (corner in this@bitmap.corners.drop(1))
				{
					lineTo(corner.x * scale + borderSize, corner.y * scale + borderSize)
				}
			}
		}
	}
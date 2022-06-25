package entities

import com.soywiz.klock.milliseconds
import com.soywiz.korge.tween.get
import com.soywiz.korge.tween.tween
import com.soywiz.korge.view.Container
import com.soywiz.korge.view.graphics
import com.soywiz.korim.vector.StrokeInfo
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.degrees
import com.soywiz.korma.geom.plus
import com.soywiz.korma.geom.vector.LineCap
import com.soywiz.korma.geom.vector.rect
import com.soywiz.korma.interpolation.Easing
import enums.PieceShape
import extensions.piece.borderColor
import extensions.piece.color
import extensions.piece.corners

class Piece(private val pieceShape: PieceShape, private val tileSize: Int) : Container()
{
	private val borderSize get() = tileSize / 16.0

	init
	{
		val borderCap = LineCap.ROUND

		graphics {
			fill(pieceShape.color) {
				for (point in points)
				{
					rect(
						borderSize.toInt() + point.x * tileSize,
						borderSize.toInt() + point.y * tileSize,
						tileSize, tileSize,
					)
				}
			}

			stroke(
				pieceShape.borderColor,
				StrokeInfo(borderSize, startCap = borderCap, endCap = borderCap)
			) {
				moveTo(borderSize, borderSize)

				for (corner in pieceShape.corners.drop(1))
				{
					lineTo(corner.x * tileSize + borderSize, corner.y * tileSize + borderSize)
				}
			}
		}
	}

	suspend fun rotate()
	{
		// Don't allow rotation while already rotating
		if (rotation.degrees.toInt() % 90 != 0) return

		tween(
			::rotation[rotation + 90.degrees],
			time = 200.milliseconds,
			easing = Easing.EASE_OUT,
		)
	}

	val points: Sequence<PointInt>
		get() = sequence {
			for (x in 0 until pieceShape.shape.width)
			{
				for (y in 0 until pieceShape.shape.height)
				{
					if (pieceShape.shape[x, y]) yield(PointInt(x, y))
				}
			}
		}
}
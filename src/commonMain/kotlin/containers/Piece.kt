package containers

import com.soywiz.korge.view.Container
import com.soywiz.korge.view.View
import com.soywiz.korge.view.graphics
import com.soywiz.korge.view.hitShape
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.vector.LineCap
import com.soywiz.korma.geom.vector.StrokeInfo
import com.soywiz.korma.geom.vector.VectorBuilder
import com.soywiz.korma.geom.vector.rect
import enums.PieceShape
import extensions.piece.*
import extensions.size2

class Piece(private val pieceShape: PieceShape, private val tileSize: Double) : Container()
{
	private val borderSize get() = tileSize / 16.0
	private val borderOffset = Point()

	private val shape: View
	val shapePos get() = shape.pos
	val shapeSize get() = shape.size2

	init
	{
		name = pieceShape.name
		val borderCap = LineCap.ROUND

		shape = graphics {
			fill(pieceShape.color) {
				for (point in pieceShape.points)
				{
					rect(
						borderSize + point.x * tileSize,
						borderSize + point.y * tileSize,
						tileSize, tileSize,
					)
				}
			}

			stroke(
				pieceShape.borderColor,
				StrokeInfo(borderSize, startCap = borderCap, endCap = borderCap)
			) { drawBorder(this) }

			x = -(pieceShape.center.x * tileSize + borderSize)
			y = -(pieceShape.center.y * tileSize + borderSize)
			borderOffset.setTo(x, y)
		}

		hitShape { drawBorder(this) }
	}

	private fun drawBorder(builder: VectorBuilder)
	{
		builder.moveTo(borderOffset.x + borderSize, borderOffset.y + borderSize)

		for (corner in pieceShape.corners.drop(1))
		{
			builder.lineTo(
				borderOffset.x + corner.x * tileSize + borderSize,
				borderOffset.y + corner.y * tileSize + borderSize,
			)
		}
	}
}
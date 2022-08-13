package extensions

import com.soywiz.korma.geom.IPoint
import com.soywiz.korma.geom.Point
import enums.PieceShape
import kotlin.random.Random

fun Random.nextPiece() = PieceShape.values().random(this)

fun Random.pieceShapes(): Sequence<PieceShape> = sequence {
	val shapes = PieceShape.values().toMutableList()
	while (shapes.any()) yield(shapes.removeAt(this@pieceShapes.nextInt(shapes.size)))
}

fun Random.nextPoint(from: IPoint, until: IPoint): Point =
	Point(this.nextDouble(from.x, until.x), this.nextDouble(from.y, until.y))
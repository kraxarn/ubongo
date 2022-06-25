package extensions

import enums.PieceShape
import kotlin.random.Random

fun Random.nextPiece() = PieceShape.values().random(this)

fun Random.pieceShapes(): Sequence<PieceShape> = sequence {
	val shapes = PieceShape.values().toMutableList()
	while (shapes.any()) yield(shapes.removeAt(this@pieceShapes.nextInt(shapes.size)))
}
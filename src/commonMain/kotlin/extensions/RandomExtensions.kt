package extensions

import enums.PieceShape
import kotlin.random.Random

fun Random.nextPiece() = PieceShape.values().random(this)
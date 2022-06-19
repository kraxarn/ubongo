package extensions

import entities.Piece
import kotlin.random.Random

fun Random.nextPiece() = Piece.values().random(this)
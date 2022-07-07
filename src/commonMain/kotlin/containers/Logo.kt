package containers

import com.soywiz.korge.view.*
import com.soywiz.korma.geom.degrees
import enums.PieceShape

fun Container.logo(
	size: Double = 128.0,
	callback: @ViewDslMarker Logo.() -> Unit = {},
) = Logo(size).addTo(this, callback)

class Logo(size: Double) : Container()
{
	init
	{
		val leftPiece = Piece(PieceShape.L2, size)
		leftPiece.position(0, 0)

		val rightPiece = Piece(PieceShape.T2, size)
		rightPiece.alignLeftToRightOf(leftPiece, -size)
		rightPiece.alignBottomToBottomOf(leftPiece)
		rightPiece.rotation = 90.degrees

		addChild(leftPiece)
		addChild(rightPiece)
	}
}
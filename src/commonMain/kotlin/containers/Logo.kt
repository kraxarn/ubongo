package containers

import com.soywiz.korge.view.Container
import com.soywiz.korge.view.alignBottomToBottomOf
import com.soywiz.korge.view.alignLeftToRightOf
import com.soywiz.korge.view.position
import com.soywiz.korma.geom.degrees
import enums.PieceShape

class Logo(size: Double = 128.0) : Container()
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
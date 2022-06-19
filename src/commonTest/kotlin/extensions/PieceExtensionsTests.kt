package extensions

import com.soywiz.korma.geom.PointInt
import entities.Piece
import kotlin.test.Test
import kotlin.test.assertEquals

class PieceExtensionsTests
{
	@Test
	fun testSize()
	{
		assertEquals(PointInt(3, 3), Piece.ONE.size)
		assertEquals(PointInt(2, 3), Piece.TWO.size)
	}
}
package extensions

import kotlin.random.Random
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotEquals

class RandomExtensionsTests
{
	@Test
	fun testSameSeedSamePiece()
	{
		val piece1 = Random(1).nextPiece()
		val piece2 = Random(1).nextPiece()
		assertEquals(piece1, piece2)
	}

	@Test
	fun testTwoPiecesDifferent()
	{
		val random = Random(1)
		val piece1 = random.nextPiece()
		val piece2 = random.nextPiece()
		assertNotEquals(piece1, piece2)
	}
}
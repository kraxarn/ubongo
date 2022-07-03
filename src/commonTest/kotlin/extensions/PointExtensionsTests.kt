package extensions

import com.soywiz.korma.geom.Point
import kotlin.test.Test
import kotlin.test.assertEquals

class PointExtensionsTests
{
	@Test
	fun testSum()
	{
		val sum = sequenceOf(Point(0, 0), Point(-1, -2), Point(4, 3)).sum()
		assertEquals(Point(3, 1), sum)
	}
}
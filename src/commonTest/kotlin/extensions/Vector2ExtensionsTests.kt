package extensions

import com.soywiz.kds.Array2
import com.soywiz.korma.geom.PointInt
import kotlin.test.Test
import kotlin.test.assertEquals

class Vector2ExtensionsTests
{
	@Test
	fun testSize2()
	{
		val arr2x2 = Array2(2, 2, 0)
		assertEquals(PointInt(2, 2), arr2x2.size2)

		val arr2x3 = Array2(2, 3, 0)
		assertEquals(PointInt(2, 3), arr2x3.size2)
	}
}
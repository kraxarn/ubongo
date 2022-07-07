package extensions

import com.soywiz.klock.TimeSpan
import kotlin.test.Test
import kotlin.test.assertFalse
import kotlin.test.assertTrue

class TimeSpanExtensionsTests
{
	@Test
	fun testIsNil()
	{
		assertFalse(TimeSpan.ZERO.isNil())
		assertTrue(TimeSpan.NIL.isNil())
	}
}
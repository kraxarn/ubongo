package utils

import com.soywiz.korio.async.suspendTest
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotEquals

class RandomWordTests
{
	@Test
	fun sameGeneratedWithSameSeed() = suspendTest {
		val word1 = randomWord(1)
		val word2 = randomWord(1)
		assertEquals(word1, word2)
	}

	@Test
	fun differentGeneratedWithDifferentSeed() = suspendTest {
		val word1 = randomWord(1)
		val word2 = randomWord(2)
		assertNotEquals(word1, word2)
	}
}
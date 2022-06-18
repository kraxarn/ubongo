import com.soywiz.klock.DateTime
import com.soywiz.klock.blockingSleep
import com.soywiz.klock.seconds
import kotlin.test.Test
import kotlin.test.assertEquals

class GameStateTests
{
	@Test
	fun testGenerateSeedAccuracy()
	{
		// Where we should get another seed 1 second later
		if (DateTime.now().seconds == 59)
		{
			blockingSleep(1.seconds)
		}

		// A new seed one second later should be the same
		val seed1 = GameState.generateSeed()
		blockingSleep(1.seconds)
		val seed2 = GameState.generateSeed()

		assertEquals(seed1, seed2)
	}
}
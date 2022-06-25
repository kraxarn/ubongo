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
		val gameState = GameState()
		val seed1 = gameState.seed
		blockingSleep(1.seconds)
		gameState.regenerate()
		val seed2 = gameState.seed

		assertEquals(seed1, seed2)
	}
}
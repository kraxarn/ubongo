import com.soywiz.kds.random.FastRandom
import com.soywiz.klock.DateTime
import enums.Difficulty
import kotlin.random.Random

class GameState
{
	private var currentLevel = 1
	val level get() = currentLevel

	private var currentSeed = generateSeed()
	val seed get() = currentSeed + currentLevel

	private var currentRandom: Random? = null
	val random: Random
		get()
		{
			if (currentRandom == null) currentRandom = FastRandom(seed)
			return currentRandom!!
		}

	val res = Resources()

	var difficulty = Difficulty.MEDIUM

	/**
	 * Generate a new seed and reset random number generator
	 */
	fun regenerate()
	{
		currentLevel = 1
		currentSeed = generateSeed()
		currentRandom = null
	}

	/**
	 * Go to the next level, generating a new seed
	 */
	fun nextLevel()
	{
		currentLevel++
		currentRandom = null
	}

	companion object
	{
		private fun generateSeed() = DateTime.now().startOfMinute.unixMillisLong
	}
}
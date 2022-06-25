import com.soywiz.kds.random.FastRandom
import com.soywiz.klock.DateTime
import kotlin.random.Random

class GameState
{
	var currentLevel = 1

	private var currentSeed = generateSeed()
	val seed get() = currentSeed

	private var currentRandom: Random? = null
	val random: Random
		get()
		{
			if (currentRandom == null) currentRandom = Random(seed)
			return currentRandom!!
		}

	/**
	 * Generate a new seed and reset random number generator
	 */
	fun regenerate()
	{
		currentSeed = generateSeed()
		currentRandom = null
	}

	companion object
	{
		private fun generateSeed() = DateTime.now().startOfMinute.unixMillisLong
	}
}
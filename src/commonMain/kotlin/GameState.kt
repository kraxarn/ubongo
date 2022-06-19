import com.soywiz.klock.DateTime

class GameState
{
	var seed = generateSeed()
	var currentLevel = 1

	companion object
	{
		fun generateSeed() = DateTime.now().startOfMinute.unixMillisLong
	}
}
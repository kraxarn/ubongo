import com.soywiz.klock.DateTime

class GameState
{
	var seed = generateSeed()

	companion object
	{
		fun generateSeed() = DateTime.now().startOfMinute.unixMillisLong
	}
}
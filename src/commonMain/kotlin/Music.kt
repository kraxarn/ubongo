import com.soywiz.klock.milliseconds
import com.soywiz.korau.sound.*
import com.soywiz.korge.sound.fadeOutPause
import com.soywiz.korge.sound.fadeTo
import com.soywiz.korio.file.std.resourcesVfs
import enums.ResMusic

class Music
{
	private var sound: Sound? = null
	private var soundChannel: SoundChannel? = null

	suspend fun play()
	{
		if (soundChannel?.playing == true) return

		if (soundChannel != null)
		{
			soundChannel?.resume()
			soundChannel?.fadeTo(VOLUME, FADE_TIME)
			return
		}

		if (sound == null)
		{
			sound = resourcesVfs[ResMusic.MUSIC.path].readMusic()
		}

		val params = PlaybackParameters(PlaybackTimes.INFINITE, volume = 0.0)
		soundChannel = sound?.play(params)
		soundChannel?.fadeTo(VOLUME, FADE_TIME)
	}

	suspend fun pause()
	{
		soundChannel?.fadeOutPause(FADE_TIME)
	}

	companion object
	{
		private const val VOLUME = 0.5
		private val FADE_TIME = 200.milliseconds
	}
}
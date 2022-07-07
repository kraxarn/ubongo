package helpers

import com.soywiz.klogger.Logger

class MemoryLogOutput : Logger.Output
{
	private val messages = mutableMapOf<Logger.Level, MutableList<Any?>>()

	override fun output(logger: Logger, level: Logger.Level, msg: Any?)
	{
		messages[level]?.add(msg) ?: messages.put(level, mutableListOf(msg))
	}

	operator fun get(warn: Logger.Level): Iterable<Any?>? = messages[warn]
}
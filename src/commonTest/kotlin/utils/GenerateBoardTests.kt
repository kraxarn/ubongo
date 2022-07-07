package utils

import com.soywiz.kds.random.FastRandom
import com.soywiz.klogger.Logger
import enums.Difficulty
import enums.PieceShape
import helpers.MemoryLogOutput
import kotlin.test.Test
import kotlin.test.assertEquals

class GenerateBoardTests
{
	@Test
	fun canGenerateBoardWithoutWarnings()
	{
		for (difficulty in Difficulty.values())
		{
			for (i in 0L..100L)
			{
				val warnings = getWarningsOnGeneration(i, difficulty)
				assertEquals(0, warnings, "$difficulty ($i)")
			}
		}
	}

	private fun getWarningsOnGeneration(seed: Long, difficulty: Difficulty): Int
	{
		val log = Logger("generateBoard")
		val logOutput = MemoryLogOutput()
		log.output = logOutput

		generateBoard(
			FastRandom(seed), PieceShape.values().toList(),
			difficulty.pieceCount, difficulty.rotation
		)

		return logOutput[Logger.Level.WARN]?.count() ?: 0
	}
}
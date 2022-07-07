package containers

import com.soywiz.korge.view.Container
import com.soywiz.korge.view.graphics
import com.soywiz.korge.view.position
import com.soywiz.korge.view.roundRect
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.vector.roundRect
import constants.GameColors
import enums.Difficulty
import enums.PieceShape
import extensions.containsPoint
import utils.generateBoard
import kotlin.random.Random

class Board(
	random: Random,
	pieces: Iterable<PieceShape>, difficulty: Difficulty,
	width: Double, height: Double = width,
) : Container()
{
	/**
	 * Size of each tile, including margin
	 */
	val tileSize: Double

	/**
	 * Top-left position of all tiles that should be filled
	 */
	private val tilePositions: Iterable<Point>

	init
	{
		val board = roundRect(width, height, 16.0, fill = GameColors.boardBackground)
		val tileCount = difficulty.boardSize
		tileSize = (board.width - TILE_SPACING * 2) / tileCount
		val rectSize = tileSize - TILE_SPACING
		val tiles = generateBoard(random, pieces, difficulty).toHashSet()
		val positions = mutableListOf<Point>()

		graphics {
			position(board.pos)
			fill(GameColors.cellBackground) {
				for (x in 0 until tileCount)
				{
					for (y in 0 until tileCount)
					{
						if (PointInt(x, y) !in tiles) continue

						val xPos = (TILE_SPACING * 1.5) + (tileSize * x)
						val yPos = (TILE_SPACING * 1.5) + (tileSize * y)
						roundRect(xPos, yPos, rectSize, rectSize, 8.0)
						positions.add(Point(xPos, yPos))
					}
				}
			}
		}

		tilePositions = positions
	}

	fun allTilesFilled(pieces: Iterable<Piece>): Boolean
	{
		val tileCenterOffset = tileSize / 2.0
		val tileCenter = Point(tileCenterOffset, tileCenterOffset)

		return tilePositions
			.map { localToGlobal(it + tileCenter) }
			.all { tile -> pieces.any { it.containsPoint(it.globalToLocal(tile)) } }
	}

	companion object
	{
		/**
		 * Spacing between each cell
		 */
		const val TILE_SPACING = 12.0
	}
}
package containers

import com.soywiz.korge.view.Container
import com.soywiz.korge.view.graphics
import com.soywiz.korge.view.position
import com.soywiz.korge.view.roundRect
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.vector.roundRect
import constants.GameColors
import enums.PieceShape
import utils.generateBoard
import kotlin.random.Random

class Board(random: Random, pieces: Iterable<PieceShape>, width: Double, height: Double = width) : Container()
{
	val tileSize: Double

	init
	{
		val board = roundRect(width, height, 16.0, fill = GameColors.boardBackground)
		tileSize = (board.width - TILE_SPACING * (TILE_COUNT + 1)) / TILE_COUNT
		val tiles = generateBoard(random, pieces).toHashSet()

		graphics {
			position(board.pos)
			fill(GameColors.cellBackground) {
				for (x in 0 until TILE_COUNT)
				{
					for (y in 0 until TILE_COUNT)
					{
						if (PointInt(x, y) !in tiles)
						{
							continue
						}

						val xPos = TILE_SPACING + (TILE_SPACING + tileSize) * x
						val yPos = TILE_SPACING + (TILE_SPACING + tileSize) * y
						roundRect(xPos, yPos, tileSize, tileSize, 8.0)
					}
				}
			}
		}
	}

	companion object
	{
		/**
		 * Total amount of pieces
		 */
		const val PIECE_COUNT = 5

		/**
		 * Number of tiles horizontally and vertically
		 */
		const val TILE_COUNT = 8

		/**
		 * Spacing between each cell
		 */
		const val TILE_SPACING = 12.0
	}
}
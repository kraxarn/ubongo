package entities

import com.soywiz.korge.view.Container
import com.soywiz.korge.view.graphics
import com.soywiz.korge.view.position
import com.soywiz.korge.view.roundRect
import com.soywiz.korma.geom.vector.roundRect
import constants.GameColors

class Board(width: Double, height: Double = width) : Container()
{
	init
	{
		val board = roundRect(width, height, 16.0, fill = GameColors.boardBackground)
		val cellSize = (board.width - CELL_SPACING * (TILE_COUNT + 1)) / TILE_COUNT

		graphics {
			position(board.pos)
			fill(GameColors.cellBackground) {
				for (x in 0 until TILE_COUNT)
				{
					for (y in 0 until TILE_COUNT)
					{
						val xPos = CELL_SPACING + (CELL_SPACING + cellSize) * x
						val yPos = CELL_SPACING + (CELL_SPACING + cellSize) * y
						roundRect(xPos, yPos, cellSize, cellSize, 8.0)
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
		const val CELL_SPACING = 12.0
	}
}
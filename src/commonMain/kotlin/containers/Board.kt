package containers

import com.soywiz.korge.view.Container
import com.soywiz.korge.view.graphics
import com.soywiz.korge.view.position
import com.soywiz.korge.view.roundRect
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.vector.roundRect
import constants.GameColors
import enums.PieceShape
import extensions.containsPoint
import extensions.localToGlobalXY
import utils.generateBoard
import kotlin.random.Random

class Board(random: Random, pieces: Iterable<PieceShape>, width: Double, height: Double = width) : Container()
{
	/**
	 * Size of each tile, including margin
	 */
	val tileSize: Double

	/**
	 * Top-left position of all tiles that should be filled
	 */
	private val tiles: Iterable<PointInt>

	init
	{
		val board = roundRect(width, height, 16.0, fill = GameColors.boardBackground)
		tileSize = (board.width - TILE_SPACING * 2) / TILE_COUNT
		val rectSize = tileSize - TILE_SPACING
		tiles = generateBoard(random, pieces).toHashSet()

		graphics {
			position(board.pos)
			fill(GameColors.cellBackground) {
				for (x in 0 until TILE_COUNT)
				{
					for (y in 0 until TILE_COUNT)
					{
						val point = PointInt(x, y)
						if (point !in tiles) continue

						val pos = getTilePosition(point)
						roundRect(pos.x, pos.y, rectSize, rectSize, 8.0)
					}
				}
			}
		}
	}

	private fun getTilePosition(pos: PointInt): Point
	{
		val xPos = (TILE_SPACING * 1.5) + (tileSize * pos.x)
		val yPos = (TILE_SPACING * 1.5) + (tileSize * pos.y)
		return Point(xPos, yPos)
	}

	fun allTilesFilled(pieces: Iterable<Piece>): Boolean
	{
		val tileCenterOffset = tileSize / 2.0
		val tileCenter = Point(tileCenterOffset, tileCenterOffset)

		return tiles
			.map { localToGlobalXY(getTilePosition(it) + tileCenter) }
			.all { tile -> pieces.any { it.hitShape2d.containsPoint(it.globalToLocal(tile)) } }
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
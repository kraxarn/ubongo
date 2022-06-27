package utils

import com.soywiz.klogger.Logger
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.plus
import containers.Board
import enums.PieceShape
import extensions.piece.points
import extensions.piece.size
import kotlin.random.Random

fun generateBoard(random: Random, pieces: Iterable<PieceShape>): Iterable<PointInt>
{
	val log = Logger("generateBoard")
	val tiles = mutableListOf<PointInt>()
	val shuffled = pieces.shuffled(random)

	for (i in shuffled.indices)
	{
		val piece = shuffled[i]
		val points = piece.points.toList()

		// First is placed in center
		if (i == 0)
		{
			val pieceSize = piece.size
			val centerX = Board.TILE_COUNT / 2 - pieceSize.x / 2
			val centerY = Board.TILE_COUNT / 2 - pieceSize.y / 2
			val center = PointInt(centerX, centerY)

			for (point in points)
			{
				tiles += center + point
			}
			continue
		}

		// Find best position for the rest
		var max = 0
		var results = mutableListOf<PointInt>()

		for (y in 0 until Board.TILE_COUNT)
		{
			for (x in 0 until Board.TILE_COUNT)
			{
				val offset = PointInt(x, y)
				if (anyOverflow(points, offset) || !allTilesFree(tiles, points, offset)) continue

				val count = adjacentTileCount(tiles, points, offset)
				if (count > max)
				{
					max = count
					results = mutableListOf(offset)
					continue
				}
				if (count == max)
				{
					results += offset
				}
			}
		}

		// Tile doesn't fit anywhere
		if (results.isEmpty())
		{
			log.warn { "No results found for tile $i" }
			continue
		}

		// Pick a random result
		val result = results.random(random)
		points.mapTo(tiles) { result + it }
	}

	return tiles
}

private fun allTilesFree(tiles: Iterable<PointInt>, piece: Iterable<PointInt>, offset: PointInt): Boolean
{
	for (shapePoint in piece)
	{
		val current = offset + shapePoint
		for (tilePoint in tiles)
		{
			if (tilePoint == current)
			{
				return false
			}
		}
	}

	return true
}

private fun anyOverflow(piece: Iterable<PointInt>, offset: PointInt): Boolean
{
	for (point in piece)
	{
		val current = offset + point
		if (current.x < 0 || current.y < 0
			|| current.x >= Board.TILE_COUNT
			|| current.y >= Board.TILE_COUNT)
		{
			return true
		}
	}

	return false
}

private fun adjacentTileCount(tiles: Iterable<PointInt>, piece: Iterable<PointInt>, offset: PointInt): Int
{
	var count = 0
	val isInTiles = { point: PointInt -> point !in piece && offset + point in tiles }

	for (point in piece)
	{
		if (isInTiles(point + PointInt(-1, 0))) count++ // Left
		if (isInTiles(point + PointInt(0, -1))) count++ // Top
		if (isInTiles(point + PointInt(1, 0))) count++ // Right
		if (isInTiles(point + PointInt(0, 1))) count++ // Bottom
	}

	return count
}
package scenes

import GameState
import com.soywiz.klock.TimeSpan
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.draggable
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korio.file.std.resourcesVfs
import com.soywiz.korio.util.toStringDecimal
import com.soywiz.korma.geom.Point
import constants.GameColors
import constants.TextSize
import containers.Board
import containers.Piece
import extensions.*

@KorgeExperimental
class GameScene(private val gameState: GameState) : Scene()
{
	private lateinit var titleSkin: UISkin
	private lateinit var textSkin: UISkin

	private var startTime = TimeSpan.NIL

	private lateinit var pieces: List<Piece>

	override suspend fun Container.sceneInit()
	{
		val regularFont = resourcesVfs["fonts/regular.ttf"].readTtfFont()
		val lightFont = resourcesVfs["fonts/light.ttf"].readTtfFont()

		titleSkin = UISkin {
			textFont = lightFont
			textColor = GameColors.foregroundAlt.withAd(0.75)
			textSize = TextSize.button * 0.5
		}

		textSkin = UISkin {
			textFont = regularFont
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button
		}
	}

	override suspend fun Container.sceneMain()
	{
		val size = views.virtualWidth - PADDING * 2

		val hud = uiGridFill(size * 0.8, 70.0, 2, 2) {
			position(PADDING, PADDING)
			uiText("Time") {
				uiSkin = titleSkin
			}
			uiText("Level") {
				uiSkin = titleSkin
			}
			val time = uiText("0.0") {
				uiSkin = textSkin
			}
			uiText(gameState.currentLevel.toString()) {
				uiSkin = textSkin
			}

			addUpdater {
				val diff = (TimeSpan.now() - startTime).seconds
				time.text = diff.toStringDecimal(1)
			}
		}

		val pieceShapes = gameState.random.pieceShapes()
			.take(Board.PIECE_COUNT)
			.toList()

		val board = Board(gameState.random, pieceShapes, size).addTo(this) {
			position(PADDING, 0.0)
			alignTopToBottomOf(hud, PADDING)
		}

		pieces = pieceShapes
			.map { Piece(it, board.tileSize) }
			.toList()

		val pieceContainer = roundRect(size, 630.0, 16.0, fill = GameColors.boardBackground) {
			position(PADDING, 0.0)
			alignTopToBottomOf(board, PADDING)
		}

		for (piece in pieces)
		{
			val padding = Point(PADDING)
			val topLeft = pieceContainer.pos + piece.size2 * 0.5 + padding
			val bottomRight = topLeft + pieceContainer.size2 - piece.size2 * 0.5 - padding

			piece.position(gameState.random.nextPoint(topLeft, bottomRight))
			addChild(piece)

			piece.draggable(autoMove = false) {
				if (piece.collidesWith(board))
				{
					// Snap to grid
					val tileSize = board.tileSize.toInt()
					val piecePos = it.viewNextXY.toInt()
					val boardPos = (board.pos + Point(Board.TILE_SPACING)).toInt()
					val x = ((piecePos.x - boardPos.x) / tileSize * tileSize) + boardPos.x
					val y = ((piecePos.y - boardPos.y) / tileSize * tileSize) + boardPos.y
					piece.position(x, y)
				}
				else piece.position(it.viewNextXY)
			}
		}
	}

	override suspend fun sceneAfterInit()
	{
		startTime = TimeSpan.now()
	}

	companion object
	{
		const val PADDING = 64.0
	}
}
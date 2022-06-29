package scenes

import GameState
import Resources
import com.soywiz.klock.TimeSpan
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.draggable
import com.soywiz.korge.input.onClick
import com.soywiz.korge.input.onMouseDrag
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import com.soywiz.korio.util.toStringDecimal
import com.soywiz.korma.geom.Point
import constants.GameColors
import constants.TextSize
import containers.*
import enums.ResFont
import enums.ResImage
import extensions.*

@KorgeExperimental
class GameScene(private val gameState: GameState) : Scene()
{
	private lateinit var titleSkin: UISkin
	private lateinit var textSkin: UISkin

	private lateinit var pauseIcon: Bitmap
	private lateinit var pauseButton: Image

	private var pauseDialog: PauseDialog? = null
	private var winDialog: WinDialog? = null
	private var dialogBackdrop: View? = null

	private var duration = TimeSpan.NIL

	private lateinit var pieces: List<Piece>

	private var paused
		get() = dialogBackdrop?.visible == true
		set(value)
		{
			dialogBackdrop?.bringToTop()
			dialogBackdrop?.visible = value

			if (winDialog == null)
			{
				pauseDialog?.bringToTop()
				pauseDialog?.visible = value
			}
			else
			{
				winDialog?.bringToTop()
			}
		}

	override suspend fun Container.sceneInit()
	{
		Resources.loadAll()

		titleSkin = UISkin {
			textFont = Resources[ResFont.LIGHT]
			textColor = GameColors.foregroundAlt.withAd(0.75)
			textSize = TextSize.button * 0.5
		}

		textSkin = UISkin {
			textFont = Resources[ResFont.REGULAR]
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button
		}

		pauseIcon = SVG(Resources[ResImage.UI_PAUSE]).render()
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
				if (paused || duration == TimeSpan.NIL) return@addUpdater
				duration += it
				time.text = duration.seconds.toStringDecimal(1)
			}
		}

		val pieceShapes = gameState.random.pieceShapes()
			.take(Board.PIECE_COUNT)
			.toList()

		val board = Board(gameState.random, pieceShapes, size).addTo(this) {
			position(PADDING, 0.0)
			alignTopToBottomOf(hud, PADDING)
		}

		pauseButton = image(pauseIcon) {
			maxWidthOrHeight(96.0)
			alignRightToRightOf(board, PADDING / 2)
			alignBottomToTopOf(board, PADDING / 4)
			onClick { paused = true }
		}

		pieces = pieceShapes
			.map { Piece(it, board.tileSize) }
			.toList()

		val pieceContainer = roundRect(size, 630.0, 16.0, fill = GameColors.boardBackground) {
			position(PADDING, 0.0)
			alignTopToBottomOf(board, PADDING)
		}

		val dialogSize = Point(views.virtualWidth * 0.75, views.virtualHeight * 0.25)

		for (piece in pieces)
		{
			val padding = Point(PADDING)
			val topLeft = pieceContainer.pos + piece.size2 * 0.5 + padding
			val bottomRight = topLeft + pieceContainer.size2 - piece.size2 * 0.5 - padding

			piece.position(gameState.random.nextPoint(topLeft, bottomRight))
			addChild(piece)

			piece.draggable(autoMove = false) {
				piece.bringToTop()
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

				if (it.end)
				{
					if (board.allTilesFilled(pieces))
					{
						winDialog = winDialog(duration, dialogSize.x, dialogSize.y) {
							position(views.virtualWidth * 0.125, views.virtualHeight / 2 - height / 2)
							onBack { sceneContainer.changeTo<MenuScene>() }
							onNext {
								gameState.currentLevel++
								sceneContainer.changeTo<GameScene>()
							}
						}
						paused = true
					}
				}
			}
		}

		dialogBackdrop = solidRect(views.actualVirtualWidth, views.actualVirtualHeight, GameColors.dialogBackdrop) {
			position(views.actualVirtualLeft, views.actualVirtualTop)
			visible = false
			// Prevent pieces from being moved
			onMouseDrag { }
		}

		pauseDialog = pauseDialog(dialogSize.x, dialogSize.y) {
			position(views.virtualWidth * 0.125, views.virtualHeight / 2 - height / 2)
			visible = false
			onBack { sceneContainer.changeTo<MenuScene>() }
			onResume { paused = false }
		}
	}

	override suspend fun sceneAfterInit()
	{
		duration = TimeSpan.ZERO
	}

	companion object
	{
		const val PADDING = 64.0
	}
}
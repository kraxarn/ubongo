package scenes

import GameState
import com.soywiz.klock.TimeSpan
import com.soywiz.klock.milliseconds
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.draggable
import com.soywiz.korge.input.onClick
import com.soywiz.korge.input.onMouseDrag
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.tween.V2
import com.soywiz.korge.tween.get
import com.soywiz.korge.tween.tween
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.color.Colors
import com.soywiz.korim.vector.StrokeInfo
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import com.soywiz.korio.async.launchImmediately
import com.soywiz.korio.util.OS
import com.soywiz.korio.util.toStringDecimal
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.plus
import com.soywiz.korma.geom.vector.rect
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
	private val dialogSize get() = Point(views.virtualWidth * 0.75, views.virtualHeight * 0.25)

	private var duration = TimeSpan.NIL

	private lateinit var board: Board
	private lateinit var pieces: List<Piece>

	private fun View.visible(visible: Boolean, vararg animations: V2<*>)
	{
		this.bringToTop()
		launchImmediately {
			if (visible) this.visible = true
			this.tween(*animations, time = 150.milliseconds)
			if (!visible) this.visible = false
		}
	}

	private var paused
		get() = dialogBackdrop?.visible == true
		set(value)
		{
			dialogBackdrop?.visible(value, dialogBackdrop!!::alpha[if (value) 0.6 else 0.0])

			val dialog = if (winDialog != null) winDialog else pauseDialog
			dialog?.visible(
				value,
				dialog::scale[if (value) 1.0 else 0.5],
				dialog::alpha[if (value) 1.0 else 0.0],
			)
		}

	override suspend fun Container.sceneInit()
	{
		titleSkin = UISkin {
			textFont = gameState.res[ResFont.LIGHT]
			textColor = GameColors.foregroundAlt.withAd(0.75)
			textSize = TextSize.button * 0.5
		}

		textSkin = UISkin {
			textFont = gameState.res[ResFont.REGULAR]
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button
		}

		pauseIcon = SVG(gameState.res[ResImage.UI_PAUSE]).render()
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
			uiText(gameState.level.toString()) {
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

		board = Board(gameState.random, pieceShapes, size).addTo(this) {
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
		addPieces(pieceContainer)

		dialogBackdrop = solidRect(views.actualVirtualWidth, views.actualVirtualHeight, GameColors.dialogBackdrop) {
			position(views.actualVirtualLeft, views.actualVirtualTop)
			alpha = 0.0
			visible = false
			// Prevent pieces from being moved
			onMouseDrag { }
		}

		pauseDialog = pauseDialog(gameState.res, dialogSize.x, dialogSize.y) {
			position(views.virtualWidth * 0.125 + width / 2.0, views.virtualHeight / 2.0)
			visible = false
			scale = 0.5
			alpha = 0.0
			onBack { sceneContainer.changeTo<MenuScene>() }
			onResume { paused = false }
		}
	}

	private fun addPieces(container: View)
	{
		for (piece in pieces)
		{
			val padding = Point(PADDING / 2.0)
			val topLeft = container.pos + (padding)
			val bottomRight = topLeft + container.size2 - (padding * 2)

			if (OS.isJvm)
			{
				sceneView.graphics {
					stroke(Colors.RED, StrokeInfo()) {
						rect(
							topLeft.x, topLeft.y,
							bottomRight.x - topLeft.x,
							bottomRight.y - topLeft.y,
						)
					}
				}
			}

			val pos = gameState.random.nextPoint(topLeft, bottomRight)

			val pieceTopLeft = pos + piece.shapePos
			val pieceBottomRight = pieceTopLeft + piece.shapeSize

			pos += sequenceOf<Pair<Double, (Double) -> Point>>(
				topLeft.x - pieceTopLeft.x to { Point(it, 0.0) },          // Left
				pieceBottomRight.x - bottomRight.x to { Point(-it, 0.0) }, // Right
				topLeft.y - pieceTopLeft.y to { Point(0.0, it) },          // Top
				pieceBottomRight.y - bottomRight.y to { Point(0.0, -it) }) // Bottom
				.filter { it.first > 0 }
				.fold(Point.Zero) { sum, pair -> sum + pair.second(pair.first) }
				as Point

			addPiece(piece) { position(pos) }
		}
	}

	private fun addPiece(piece: Piece, callback: @ViewDslMarker Piece.() -> Unit)
	{
		piece.addTo(sceneView, callback)

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
					winDialog = sceneView.winDialog(gameState.res, duration, dialogSize.x, dialogSize.y) {
						position(views.virtualWidth * 0.125 + width / 2.0, views.virtualHeight / 2.0)
						scale = 0.5
						alpha = 0.0
						onBack { sceneContainer.changeTo<MenuScene>() }
						onNext {
							gameState.nextLevel()
							sceneContainer.changeTo<GameScene>()
						}
					}
					paused = true
				}
			}
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
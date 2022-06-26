package scenes

import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.onClick
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korio.file.std.resourcesVfs
import containers.Piece
import enums.PieceShape
import extensions.piece.corners
import skins.ButtonSkin

@KorgeExperimental
class DebugPiecesScene : Scene()
{
	private var index = 0
	private lateinit var text: UIText
	private lateinit var debugText: UIText
	private lateinit var piece: Piece

	override suspend fun Container.sceneInit()
	{
		val font = resourcesVfs["fonts/debug.ttf"].readTtfFont()
		uiSkin = ButtonSkin(font)
	}

	override suspend fun Container.sceneMain()
	{
		uiButton("<- <- Menu", width = BUTTON_WIDTH, height = BUTTON_HEIGHT) {
			alignLeftToLeftOf(this@sceneMain, PADDING)
			alignTopToTopOf(this@sceneMain, PADDING)
			onClick {
				sceneContainer.changeTo<MenuScene>()
			}
		}

		text = uiText("none") {
			textAlignment = TextAlignment.MIDDLE_CENTER
			size(views.virtualWidth, 128)
			position(0, views.virtualHeight - 200)
		}

		debugText = uiText("") {
			textAlignment = TextAlignment.MIDDLE_RIGHT
			centerOn(this@sceneMain)
			alignRightToRightOf(this@sceneMain, PADDING)
		}

		val previous = uiButton("<- Previous", width = BUTTON_WIDTH, height = BUTTON_HEIGHT) {
			alignLeftToLeftOf(this@sceneMain, PADDING)
			alignBottomToBottomOf(this@sceneMain)
			onClick { navigate(-1) }
		}

		uiButton("Mirror", width = previous.width, height = previous.height) {
			disable()
			alignLeftToLeftOf(previous)
			alignBottomToTopOf(previous, PADDING)
		}

		val next = uiButton("Next ->", width = BUTTON_WIDTH, height = BUTTON_HEIGHT) {
			alignRightToRightOf(this@sceneMain, PADDING)
			alignBottomToBottomOf(this@sceneMain)
			onClick { navigate(1) }
		}

		uiButton("Rotate", width = next.width, height = next.height) {
			alignRightToRightOf(next)
			alignBottomToTopOf(next, PADDING)
			onClick { piece.rotate() }
		}

		piece = Piece(PieceShape.values()[index], PIECE_SIZE)
		this@DebugPiecesScene.addPiece(piece)
	}

	private fun addPiece(piece: Piece)
	{
		piece.addTo(root) {
			centerOn(root)
		}
	}

	private fun navigate(steps: Int)
	{
		val pieceShapes = PieceShape.values()
		index += steps
		if (index < 0) index = pieceShapes.size - 1
		else if (index >= pieceShapes.size) index = 0
		val pieceShape = pieceShapes[index]

		root.removeChild(piece)
		piece = Piece(pieceShape, PIECE_SIZE)
		addPiece(piece)

		text.text = "$index: ${pieceShapes[index].name}"

		debugText.text = pieceShapes[index].corners
			.joinToString("\n") { it.toString() }
	}

	companion object
	{
		private const val BUTTON_WIDTH = 320.0
		private const val BUTTON_HEIGHT = 128.0
		private const val PADDING = 42.0
		private const val PIECE_SIZE = 160.0
	}
}
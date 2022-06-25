package scenes

import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.onClick
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.slice
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korio.file.std.resourcesVfs
import entities.Piece
import extensions.piece.bitmap
import extensions.piece.corners
import extensions.piece.rotatePiece
import skins.ButtonSkin

@KorgeExperimental
class DebugPiecesScene : Scene()
{
	private var index = 0
	private lateinit var text: UIText
	private lateinit var debugText: UIText
	private lateinit var piece: Image

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
			onClick { piece.rotatePiece() }
		}

		piece = image(NativeImage(0, 0)) {
			center()
			position(views.virtualWidth / 2, views.virtualHeight / 2)
			size(3, 3)
		}

		navigate(0)
	}

	private fun navigate(steps: Int)
	{
		val pieces = Piece.values()
		index += steps
		if (index < 0) index = pieces.size - 1
		else if (index >= pieces.size) index = 0

		piece.bitmap = pieces[index].bitmap.slice()
		text.text = "$index: ${pieces[index].name}"

		debugText.text = pieces[index].corners
			.joinToString("\n") { it.toString() }
	}

	companion object
	{
		private const val BUTTON_WIDTH = 320.0
		private const val BUTTON_HEIGHT = 128.0
		private const val PADDING = 42.0
	}
}
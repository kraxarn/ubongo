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
import extensions.bitmap
import skins.ButtonSkin

@KorgeExperimental
class DebugPiecesScene : Scene()
{
	private var index = 0
	private lateinit var text: UIText
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

		uiButton("<- Previous", width = BUTTON_WIDTH, height = BUTTON_HEIGHT) {
			alignLeftToLeftOf(this@sceneMain, PADDING)
			alignBottomToBottomOf(this@sceneMain)
			onClick { navigate(-1) }
		}

		uiButton("Next ->", width = BUTTON_WIDTH, height = BUTTON_HEIGHT) {
			alignRightToRightOf(this@sceneMain, PADDING)
			alignBottomToBottomOf(this@sceneMain)
			onClick { navigate(1) }
		}

		piece = image(NativeImage(0, 0)) {
			smoothing = false
			center()
			position(views.virtualWidth / 2, views.virtualHeight / 2)
			size(128, 128)
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
	}

	companion object
	{
		private const val BUTTON_WIDTH = 320.0
		private const val BUTTON_HEIGHT = 128.0
		private const val PADDING = 42.0
	}
}
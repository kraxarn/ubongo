package scenes

import GameState
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.onClick
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korim.format.readBitmap
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korio.file.std.resourcesVfs
import constants.Application
import constants.GameColors
import constants.TextSize
import images.background
import skins.ButtonSkin
import utils.randomWord

@KorgeExperimental
class MenuScene(private val gameState: GameState) : Scene()
{
	private lateinit var logoBitmap: Bitmap

	private lateinit var titleSkin: UISkin
	private lateinit var title2Skin: UISkin
	private lateinit var buttonSkin: UISkin
	private lateinit var textSkin: UISkin

	override suspend fun Container.sceneInit()
	{
		logoBitmap = resourcesVfs["images/logo.png"].readBitmap()

		titleSkin = UISkin {
			textFont = resourcesVfs["fonts/title.ttf"].readTtfFont()
			textColor = GameColors.foregroundAlt
			textSize = TextSize.title
		}

		title2Skin = UISkin {
			textFont = resourcesVfs["fonts/debug.ttf"].readTtfFont()
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button
		}

		val regularFont = resourcesVfs["fonts/regular.ttf"].readTtfFont()
		buttonSkin = ButtonSkin(regularFont)

		textSkin = UISkin {
			textFont = regularFont
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button
		}
	}

	override suspend fun Container.sceneMain()
	{
		addChild(background(views.virtualWidth, views.virtualHeight))

		container {
			centerXOn(this@sceneMain)
			alignTopToTopOf(this@sceneMain, views.virtualHeight * 0.2)

			image(logoBitmap) {
				anchor(0.5, 0.0)
			}

			val title = uiText("Ubongo") {
				textAlignment = TextAlignment.MIDDLE_RIGHT
				uiSkin = titleSkin
				size(340, 100)
				alignBottomToBottomOf(this@container)
				alignLeftToLeftOf(this@container)
			}

			uiText(Application.VERSION) {
				uiSkin = title2Skin
				textAlignment = TextAlignment.MIDDLE_RIGHT
				alignTopToBottomOf(title)
				alignRightToRightOf(title)
			}
		}

		val generateSeed = uiButton("Generate Seed") {
			uiSkin = buttonSkin
			x = PADDING
			size(views.virtualWidth - PADDING * 2, BUTTON_HEIGHT)
			alignBottomToBottomOf(this@sceneMain, PADDING * 2)
		}

		val startGame = uiButton("Start Game") {
			uiSkin = buttonSkin
			size(views.virtualWidth - PADDING * 2, BUTTON_HEIGHT)
			alignBottomToTopOf(generateSeed, PADDING)
			alignLeftToLeftOf(generateSeed)
			onClick {
				sceneContainer.changeTo<GameScene>()
			}
		}

		val seedName = uiText(randomWord(gameState.seed)) {
			uiSkin = textSkin
			textAlignment = TextAlignment.BOTTOM_CENTER
			size(startGame.width, 0.0)
			alignBottomToTopOf(startGame, PADDING)
			alignLeftToLeftOf(startGame)
		}

		generateSeed.onClick {
			gameState.regenerate()
			seedName.text = randomWord(gameState.seed)
		}
	}

	companion object
	{
		private const val PADDING = 64.0
		private const val BUTTON_HEIGHT = 180.0
	}
}
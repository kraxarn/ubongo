package scenes

import GameState
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.component.length.*
import com.soywiz.korge.input.onClick
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korim.format.readBitmap
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korio.file.std.resourcesVfs
import components.Background
import constants.GameColors
import constants.TextSize
import skins.ButtonSkin
import utils.randomWord

@KorgeExperimental
class MenuScene(private val gameState: GameState) : Scene()
{
	private lateinit var logoBitmap: Bitmap

	private lateinit var titleSkin: UISkin
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
		addComponent(Background(views))

		container {
			image(logoBitmap).centered
			uiText("Ubongo") {
				textAlignment = TextAlignment.MIDDLE_RIGHT
				uiSkin = titleSkin
				size(340, 100)
				alignBottomToBottomOf(this@container)
				alignLeftToLeftOf(this@container)
			}
		}.lengths {
			x = 50.vw
			y = 30.vh
			width = min(50.vw, 40.vh)
			height = width
		}

		val seedName = uiText(randomWord(gameState.seed)) {
			uiSkin = textSkin
			textAlignment = TextAlignment.BOTTOM_CENTER
		}

		val startGame = uiButton("Start Game") {
			uiSkin = buttonSkin
		}

		val generateSeed = uiButton("Generate Seed") {
			uiSkin = buttonSkin
			onClick {
				gameState.seed = GameState.generateSeed()
				seedName.text = randomWord(gameState.seed)
			}
		}

		generateSeed.lengths {
			width = 100.vw - 120.pt
			height = 160.pt
			x = 60.pt
			y = 100.vh - (height ?: 0.pt) - 15.percent
		}

		startGame.lengths {
			width = generateSeed.lengths.width
			height = generateSeed.lengths.height
			x = generateSeed.lengths.x
			y = (generateSeed.lengths.y ?: 0.pt) - (height ?: 0.pt) - 60.pt
		}

		seedName.lengths {
			y = (startGame.lengths.y ?: 0.pt) - 120.pt
			width = 100.vw
		}
	}
}



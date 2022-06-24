package scenes

import GameState
import com.soywiz.klock.TimeSpan
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korio.file.std.resourcesVfs
import com.soywiz.korio.util.toStringDecimal
import constants.GameColors
import constants.TextSize
import entities.Board
import extensions.now
import images.background

@KorgeExperimental
class GameScene(private val gameState: GameState) : Scene()
{
	private lateinit var titleSkin: UISkin
	private lateinit var textSkin: UISkin

	private var startTime = TimeSpan.NIL

	override suspend fun Container.sceneInit()
	{
		val regularFont = resourcesVfs["fonts/regular.ttf"].readTtfFont()

		titleSkin = UISkin {
			textFont = regularFont
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
		addChild(background(views.virtualWidth, views.virtualHeight))

		val size = views.virtualWidth - PADDING * 2

		val hud = uiGridFill(size * 0.8, 100.0, 2, 2) {
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

		val board = Board(size).addTo(this) {
			position(PADDING, 0.0)
			alignTopToBottomOf(hud, PADDING / 2)
		}

		roundRect(size, 630.0, 16.0, fill = GameColors.boardBackground) {
			position(PADDING, 0.0)
			alignTopToBottomOf(board, PADDING)
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
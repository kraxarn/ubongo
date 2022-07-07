package scenes

import GameState
import com.soywiz.klock.DateTime
import com.soywiz.klock.seconds
import com.soywiz.korge.animate.launchAnimate
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.onClick
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.tween.get
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korge.view.filter.TransitionFilter
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korio.async.launch
import com.soywiz.korma.interpolation.Easing
import constants.Application
import constants.GameColors
import constants.TextSize
import containers.Logo
import enums.Difficulty
import enums.ResFont
import skins.ButtonSkin
import utils.randomWord

@KorgeExperimental
class MenuScene(private val gameState: GameState) : Scene()
{
	private lateinit var titleSkin: UISkin
	private lateinit var title2Skin: UISkin
	private lateinit var buttonSkin: UISkin
	private lateinit var textSkin: UISkin

	override suspend fun Container.sceneInit()
	{
		titleSkin = UISkin {
			textFont = gameState.res[ResFont.BOLD]
			textColor = GameColors.foregroundAlt
			textSize = TextSize.title
		}

		title2Skin = UISkin {
			textFont = gameState.res[ResFont.LIGHT]
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button * 0.75
		}

		val regularFont = gameState.res[ResFont.REGULAR]
		buttonSkin = ButtonSkin(regularFont)

		textSkin = UISkin {
			textFont = regularFont
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button
		}
	}

	override suspend fun Container.sceneMain()
	{
		gameState.regenerate()

		container {
			val logo = Logo()
			addChild(logo)

			val title = uiText("Ubongo") {
				textAlignment = TextAlignment.MIDDLE_RIGHT
				uiSkin = titleSkin
				size(340, 100)
				alignTopToBottomOf(logo, -96)
				alignRightToRightOf(logo, 160)
			}

			uiText(Application.VERSION) {
				uiSkin = title2Skin
				textAlignment = TextAlignment.MIDDLE_RIGHT
				alignTopToBottomOf(title, 32.0)
				alignRightToRightOf(title)
			}

			position(views.virtualWidth / 2.0 - width / 4.0, views.virtualHeight * 0.15)
		}

		val back = uiButton("<")
		val difficulties = uiHorizontalFill()

		val startGame = uiButton("Start Game") {
			uiSkin = buttonSkin
			position(PADDING, views.virtualHeight * 0.8)
			size(views.virtualWidth - PADDING * 2, BUTTON_HEIGHT)
			onClick {
				visible(false)
				back.visible(true)
				difficulties.visible(true)
			}
		}

		val seedName = uiText(randomWord(gameState.seed)) {
			uiSkin = textSkin
			textAlignment = TextAlignment.MIDDLE_LEFT
			size(startGame.width, 68.0)
			alignLeftToLeftOf(startGame, PADDING / 2)
			alignBottomToTopOf(startGame, PADDING)
		}

		circle(seedName.height / 2.0, GameColors.foregroundAlt) {
			alignRightToRightOf(startGame, PADDING / 2)
			alignBottomToTopOf(startGame, PADDING)
			val seedFilter = TransitionFilter(TransitionFilter.Transition.SWEEP, smooth = false)
			filter = seedFilter
			launchAnimate {
				val nextMinute = 60 - DateTime.now().seconds
				block { seedFilter.ratio = 0.5 + (nextMinute / 60.0 / 2.0) }
				tween(seedFilter::ratio[0.5], time = nextMinute.seconds, easing = Easing.LINEAR)
				sequence(looped = true) {
					block {
						launch {
							gameState.regenerate()
							seedName.text = randomWord(gameState.seed)
						}
						seedFilter.ratio = 1.0
					}
					tween(seedFilter::ratio[0.5], time = 60.seconds, easing = Easing.LINEAR)
				}
			}
		}

		back.apply {
			uiSkin = buttonSkin
			size(startGame.width / 6.0, startGame.height)
			alignLeftToLeftOf(startGame)
			alignTopToTopOf(startGame)
			visible(false)
			onClick {
				back.visible(false)
				difficulties.visible(false)
				startGame.visible(true)
			}
		}

		difficulties.apply {
			size(startGame.width - back.width, startGame.height)
			alignLeftToRightOf(back)
			alignTopToTopOf(back)
			visible(false)

			val buttonSkin = this@MenuScene.buttonSkin.copy()
			buttonSkin.textSize *= 0.7

			for (difficulty in Difficulty.values())
			{
				uiButton(difficulty.name.lowercase().replaceFirstChar { it.uppercase() }) {
					uiSkin = buttonSkin
					onClick {
						gameState.difficulty = difficulty
						sceneContainer.changeTo<GameScene>()
					}
				}
			}
		}
	}

	companion object
	{
		private const val PADDING = 64.0
		private const val BUTTON_HEIGHT = 180.0
	}
}
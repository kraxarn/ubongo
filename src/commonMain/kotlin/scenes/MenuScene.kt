package scenes

import GameState
import com.soywiz.klock.DateTime
import com.soywiz.klock.milliseconds
import com.soywiz.klock.seconds
import com.soywiz.klogger.Logger
import com.soywiz.korge.animate.animate
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.input.onClick
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.tween.get
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korge.view.filter.TransitionFilter
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korim.format.readBitmap
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import com.soywiz.korio.async.launch
import com.soywiz.korio.file.std.resourcesVfs
import com.soywiz.korma.interpolation.Easing
import constants.Application
import constants.GameColors
import constants.TextSize
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

	private lateinit var refreshIcon: Bitmap

	override suspend fun Container.sceneInit()
	{
		logoBitmap = resourcesVfs["images/logo.png"].readBitmap()

		titleSkin = UISkin {
			textFont = resourcesVfs["fonts/bold.ttf"].readTtfFont()
			textColor = GameColors.foregroundAlt
			textSize = TextSize.title
		}

		title2Skin = UISkin {
			textFont = resourcesVfs["fonts/light.ttf"].readTtfFont()
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button * 0.75
		}

		val regularFont = resourcesVfs["fonts/regular.ttf"].readTtfFont()
		buttonSkin = ButtonSkin(regularFont)

		textSkin = UISkin {
			textFont = regularFont
			textColor = GameColors.foregroundAlt
			textSize = TextSize.button
		}

		refreshIcon = SVG(resourcesVfs["images/ui/refresh.svg"].readString()).render()
	}

	override suspend fun Container.sceneMain()
	{
		container {
			position(views.virtualWidth / 2.0, views.virtualHeight * 0.15)

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
				alignTopToBottomOf(title, 32.0)
				alignRightToRightOf(title)
			}
		}

		val startGame = uiButton("Start Game") {
			uiSkin = buttonSkin
			position(PADDING, views.virtualHeight * 0.8)
			size(views.virtualWidth - PADDING * 2, BUTTON_HEIGHT)
			onClick {
				sceneContainer.changeTo<GameScene>()
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
			animate {
				sequence(looped = true) {
					val nextMinute = 60.seconds - DateTime.now().milliseconds.milliseconds
					tween(seedFilter::ratio[0.5], time = nextMinute, easing = Easing.LINEAR)
					block {
						launch {
							gameState.regenerate()
							seedName.text = randomWord(gameState.seed)
						}
						seedFilter.ratio = 1.0
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
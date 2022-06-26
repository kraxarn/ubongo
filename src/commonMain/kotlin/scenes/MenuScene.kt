package scenes

import GameState
import com.soywiz.klock.DateTime
import com.soywiz.klock.seconds
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
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import com.soywiz.korio.async.launch
import com.soywiz.korio.file.std.resourcesVfs
import com.soywiz.korma.geom.degrees
import com.soywiz.korma.interpolation.Easing
import constants.Application
import constants.GameColors
import constants.TextSize
import containers.Piece
import enums.PieceShape
import skins.ButtonSkin
import utils.randomWord

@KorgeExperimental
class MenuScene(private val gameState: GameState) : Scene()
{
	private lateinit var titleSkin: UISkin
	private lateinit var title2Skin: UISkin
	private lateinit var buttonSkin: UISkin
	private lateinit var textSkin: UISkin

	private lateinit var refreshIcon: Bitmap

	override suspend fun Container.sceneInit()
	{
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
			val leftPiece = Piece(PieceShape.L2, 128.0)
			leftPiece.position(0, 0)

			val rightPiece = Piece(PieceShape.T2, 128.0)
			rightPiece.alignLeftToRightOf(leftPiece, -128)
			rightPiece.alignBottomToBottomOf(leftPiece)
			rightPiece.rotation = 90.degrees

			addChild(leftPiece)
			addChild(rightPiece)

			val title = uiText("Ubongo") {
				textAlignment = TextAlignment.MIDDLE_RIGHT
				uiSkin = titleSkin
				size(340, 100)
				alignTopToBottomOf(leftPiece, 32)
				alignRightToLeftOf(rightPiece, -96)
			}

			uiText(Application.VERSION) {
				uiSkin = title2Skin
				textAlignment = TextAlignment.MIDDLE_RIGHT
				alignTopToBottomOf(title, 32.0)
				alignRightToRightOf(title)
			}

			position(views.virtualWidth / 2.0 - width / 4.0, views.virtualHeight * 0.15)
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
				tween(seedFilter::ratio[0.5], time = (60 - DateTime.now().seconds).seconds, easing = Easing.LINEAR)
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
	}

	companion object
	{
		private const val PADDING = 64.0
		private const val BUTTON_HEIGHT = 180.0
	}
}
package scenes

import com.soywiz.korge.component.length.*
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.atlas.Atlas
import com.soywiz.korim.atlas.readAtlas
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.bitmap.Bitmap32
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korim.format.readBitmap
import com.soywiz.korim.paint.LinearGradientPaint
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korio.file.std.resourcesVfs
import com.soywiz.korma.geom.vector.rect
import constants.GameColors
import constants.TextSize
import skins.ButtonSkin

class MenuScene : Scene()
{
	private lateinit var logoBitmap: Bitmap
	private lateinit var uiAtlas: Atlas

	private lateinit var titleSkin: UISkin
	private lateinit var buttonSkin: UISkin

	override suspend fun Container.sceneInit()
	{
		logoBitmap = resourcesVfs["images/logo.png"].readBitmap()
		uiAtlas = resourcesVfs["images/ui.atlas.json"].readAtlas()

		val width = views.virtualWidth
		val height = views.virtualHeight

		val background = Bitmap32(width, height).context2d {
			fill(
				LinearGradientPaint(0, 0, 0, height)
					.add(0.0, GameColors.backgroundStart)
					.add(1.0, GameColors.backgroundEnd)
			) {
				rect(0, 0, width, height)
			}
		}

		image(background)

		titleSkin = UISkin {
			textFont = resourcesVfs["fonts/title.ttf"].readTtfFont()
			textColor = GameColors.foregroundAlt
			textSize = TextSize.title
		}

		val buttonFont = resourcesVfs["fonts/regular.ttf"].readTtfFont()
		buttonSkin = ButtonSkin(buttonFont)
	}

	override suspend fun Container.sceneMain()
	{
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

		val startGame = uiButton("Start Game") {
			uiSkin = buttonSkin
		}

		val generateSeed = uiButton("Generate Seed") {
			uiSkin = buttonSkin
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
	}
}



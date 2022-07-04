package containers

import GameSettings
import Resources
import com.soywiz.korge.input.onClick
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.Container
import com.soywiz.korge.view.ViewDslMarker
import com.soywiz.korge.view.addTo
import com.soywiz.korim.font.Font
import com.soywiz.korim.text.TextAlignment
import enums.ResFont

fun Container.quickSettings(
	res: Resources,
	settings: GameSettings,
	width: Double, height: Double,
	callback: @ViewDslMarker QuickSettings.() -> Unit = {},
) = QuickSettings(res, settings, width, height).addTo(this, callback)

class QuickSettings(
	private val res: Resources,
	private val settings: GameSettings,
	width: Double, height: Double,
) : UIGridFill(width, height, 3, 2)
{
	private val defaultSettings = GameSettings()

	private val titleSkin = UISkin {
		textFont = res[ResFont.BOLD]
		textSize = 64.0
		textAlignment = TextAlignment.BOTTOM_CENTER
	}

	private val infoSkin = UISkin {
		textFont = res[ResFont.REGULAR]
		textSize = 32.0
		textAlignment = TextAlignment.TOP_CENTER
	}

	init
	{
		uiText(settings.rotation.displayName) {
			uiSkin = titleSkin
			onClick {
				text = (++settings.rotation).displayName
				textFont = getFont { it.rotation }
			}
		}

		uiText(settings.pieceCount.displayName) {
			uiSkin = titleSkin
			onClick {
				text = (++settings.pieceCount).displayName
				textFont = getFont { it.pieceCount }
			}
		}

		uiText(settings.boardSize.displayName) {
			uiSkin = titleSkin
			onClick {
				text = (++settings.boardSize).displayName
				textFont = getFont { it.boardSize }
			}
		}

		uiText("Rotation") { uiSkin = infoSkin }
		uiText("Pieces") { uiSkin = infoSkin }
		uiText("Board size") { uiSkin = infoSkin }
	}

	private fun <T> getFont(option: (GameSettings) -> T): Font
	{
		val isDefault = option(settings) == option(defaultSettings)
		return res[if (isDefault) ResFont.BOLD else ResFont.BOLD_ITALIC]
	}
}

private inline operator fun <reified T : Enum<T>> Enum<T>.inc(): T
{
	val values = enumValues<T>()
	return values[(this.ordinal + 1) % values.size]
}
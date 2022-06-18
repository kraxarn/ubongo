package skins

import com.soywiz.korge.ui.*
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.NinePatchBmpSlice
import com.soywiz.korim.bitmap.asNinePatchSimple
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.color.Colors
import com.soywiz.korim.color.RGBA
import com.soywiz.korim.format.writeTo
import com.soywiz.korio.async.runBlockingNoJs
import com.soywiz.korio.file.std.applicationVfs
import com.soywiz.korio.lang.InvalidArgumentException
import extensions.roundRect

class ButtonSkin : UISkin()
{
	init
	{
		buttonNormal = buildButton(UiSkinType.NORMAL)
		buttonOver = buildButton(UiSkinType.OVER)
		buttonDown = buildButton(UiSkinType.DOWN)
		buttonDisabled = buildButton(UiSkinType.DISABLED)
	}

	private fun buildButton(type: UiSkinType): NinePatchBmpSlice
	{
		val shadow = type == UiSkinType.NORMAL
			|| type == UiSkinType.OVER

		val img = NativeImage(BUTTON_SIZE, BUTTON_SIZE).context2d {
			val fillColor = when (type)
			{
				UiSkinType.NORMAL -> Colors["#3f51b5"]   // Indigo 500
				UiSkinType.OVER -> Colors["#757de8"]     // Indigo 500 (Light)
				UiSkinType.DOWN -> Colors["#002984"]     // Indigo 500 (Dark)
				UiSkinType.DISABLED -> Colors["#607d8b"] // Blue grey 500
				else -> throw InvalidArgumentException(type.toString())
			}

			if (shadow)
			{
				val shadowAlpha = (when (type)
				{
					UiSkinType.OVER -> 0.3
					UiSkinType.DISABLED -> 0.7
					else -> 0.5
				} * 255).toInt()

				fill(RGBA(0, 0, 0, shadowAlpha)) {
					roundRect(
						PADDING + SHADOW_WIDTH,
						BUTTON_SIZE - PADDING * 2 - SHADOW_WIDTH,
						BORDER_RADIUS
					)
				}
			}

			fill(fillColor) {
				roundRect(
					PADDING
						+ if (type == UiSkinType.DOWN) SHADOW_WIDTH else 0,
					BUTTON_SIZE - PADDING * 2
						- (if (shadow || type == UiSkinType.DISABLED || type == UiSkinType.DOWN) SHADOW_WIDTH else 0),
					BORDER_RADIUS
				)
			}
		}

		runBlockingNoJs {
			img.writeTo(applicationVfs["button/${type.name}.png"])
		}

		return img.asNinePatchSimple(INSET, INSET, INSET, INSET)
	}

	companion object
	{
		const val BUTTON_SIZE = 64
		const val PADDING = 4
		const val BORDER_RADIUS = 12
		const val SHADOW_WIDTH = 10
		const val INSET = 22
	}
}

private val UiSkinType.name
	get() = when (this)
	{
		UiSkinType.NORMAL -> "normal"
		UiSkinType.OVER -> "over"
		UiSkinType.DOWN -> "down"
		UiSkinType.DISABLED -> "disabled"
		else -> null
	}
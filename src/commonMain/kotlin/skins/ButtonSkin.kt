package skins

import com.soywiz.korge.ui.*
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.asNinePatch
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.color.Colors
import com.soywiz.korma.geom.vector.roundRect

class ButtonSkin : UISkin()
{

	init
	{
		buttonNormal = buildButton(UiSkinType.NORMAL).asNinePatch()
		buttonOver = buildButton(UiSkinType.OVER).asNinePatch()
		buttonDown = buildButton(UiSkinType.DOWN).asNinePatch()
		buttonDisabled = buildButton(UiSkinType.DISABLED).asNinePatch()
	}

	private fun buildButton(type: UiSkinType): Bitmap
	{
		return NativeImage(BUTTON_SIZE, BUTTON_SIZE).context2d {
			val fillColor = when (type)
			{
				UiSkinType.NORMAL -> Colors["#3f51b5"]   // Indigo 500
				UiSkinType.OVER -> Colors["#757de8"]     // Indigo 500 (Light)
				UiSkinType.DOWN -> Colors["#002984"]     // Indigo 500 (Dark)
				UiSkinType.DISABLED -> Colors["#607d8b"] // Blue grey 500
				else -> TODO()
			}

			fill(fillColor) {
				roundRect(
					PADDING, PADDING,
					BUTTON_SIZE - PADDING * 2, BUTTON_SIZE - PADDING * 2,
					BORDER_RADIUS, BORDER_RADIUS
				)
			}
		}
	}

	companion object
	{
		const val BUTTON_SIZE = 64
		const val PADDING = 4
		const val BORDER_RADIUS = 6
	}
}
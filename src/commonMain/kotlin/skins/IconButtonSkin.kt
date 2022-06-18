package skins

import com.soywiz.korge.ui.*
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.asNinePatch
import com.soywiz.korim.font.Font
import constants.GameColors
import constants.TextSize

class IconButtonSkin(font: Font) : UISkin()
{
	init
	{
		textFont = font
		textColor = GameColors.foregroundAlt
		textSize = TextSize.button

		val empty = NativeImage(0, 0).asNinePatch()

		buttonNormal = empty
		buttonOver = empty
		buttonDown = empty
		buttonDisabled = empty
	}
}
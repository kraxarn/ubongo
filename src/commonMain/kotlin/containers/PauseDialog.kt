package containers

import Resources
import com.soywiz.korge.input.onClick
import com.soywiz.korge.view.Container
import com.soywiz.korge.view.ViewDslMarker
import com.soywiz.korge.view.addTo
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.font.Font
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import enums.ResFont
import enums.ResImage

fun Container.pauseDialog(
	width: Double, height: Double,
	callback: @ViewDslMarker PauseDialog.() -> Unit = {},
): PauseDialog
{
	val backIcon = SVG(Resources[ResImage.UI_ARROW_LEFT]).render()
	val resumeIcon = SVG(Resources[ResImage.UI_PLAY]).render()

	val titleFont = Resources[ResFont.BOLD]
	val bodyFont = Resources[ResFont.REGULAR]

	return PauseDialog(width, height, titleFont, bodyFont, backIcon, resumeIcon)
		.addTo(this, callback)
}

class PauseDialog(
	width: Double, height: Double,
	titleFont: Font, bodyFont: Font,
	backIcon: Bitmap, resumeIcon: Bitmap,
) : Dialog(width, height, titleFont, bodyFont)
{
	val onBack = onLeftAction
	val onResume = onRightAction

	init
	{
		title("Paused")
		body("Game is paused")

		leftAction("Back to menu", backIcon) {
			onClick { onBack(it) }
		}

		rightAction("Resume", resumeIcon) {
			onClick { onResume(it) }
		}
	}
}
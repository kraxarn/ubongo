package containers

import com.soywiz.korge.input.onClick
import com.soywiz.korge.view.Container
import com.soywiz.korge.view.ViewDslMarker
import com.soywiz.korge.view.addTo
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.font.Font
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import com.soywiz.korio.file.std.resourcesVfs

suspend fun Container.pauseDialog(
	width: Double, height: Double,
	callback: @ViewDslMarker Dialog.() -> Unit = {},
): PauseDialog
{
	val backIcon = SVG(resourcesVfs["images/ui/pause.svg"].readString()).render()
	val resumeIcon = SVG(resourcesVfs["images/ui/play.svg"].readString()).render()

	val titleFont = resourcesVfs["fonts/bold.ttf"].readTtfFont()
	val bodyFont = resourcesVfs["fonts/regular.ttf"].readTtfFont()

	return PauseDialog(width, height, titleFont, bodyFont, backIcon, resumeIcon)
		.addTo(this, callback)
}

class PauseDialog(
	width: Double, height: Double,
	titleFont: Font, bodyFont: Font,
	backIcon: Bitmap, resumeIcon: Bitmap,
) : Dialog(width, height, titleFont, bodyFont)
{
	init
	{
		title("Paused")
		body("Game is paused")

		leftAction("Back to menu", backIcon) {
			onClick { TODO() }
		}

		rightAction("Resume", resumeIcon) {
			onClick { TODO() }
		}
	}
}
package containers

import Resources
import com.soywiz.klock.TimeSpan
import com.soywiz.klock.minutes
import com.soywiz.klock.seconds
import com.soywiz.korge.input.onClick
import com.soywiz.korge.view.Container
import com.soywiz.korge.view.ViewDslMarker
import com.soywiz.korge.view.addTo
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import com.soywiz.korio.util.toStringDecimal
import enums.ResFont
import enums.ResImage

fun Container.winDialog(
	total: TimeSpan,
	width: Double, height: Double,
	callback: @ViewDslMarker WinDialog.() -> Unit = {},
) = WinDialog(total, width, height).addTo(this, callback)

class WinDialog(
	total: TimeSpan,
	width: Double, height: Double,
) : Dialog(width, height, Resources[ResFont.BOLD], Resources[ResFont.REGULAR])
{
	val onBack = onLeftAction
	val onNext = onRightAction

	init
	{
		title(randomMessage(total))
		body("You completed in ${total.seconds.toStringDecimal(3)} seconds")

		val backIcon = SVG(Resources[ResImage.UI_ARROW_LEFT]).render()
		val nextIcon = SVG(Resources[ResImage.UI_PLAY]).render()

		leftAction("Back to menu", backIcon) {
			onClick { onBack(it) }
		}

		rightAction("Next level", nextIcon) {
			onClick { onNext(it) }
		}
	}

	private fun randomMessage(total: TimeSpan) = (when
	{
		total < 10.seconds -> arrayOf("Amazing!", "Incredible!", "That was quick!")
		total > 1.minutes -> arrayOf("neat", "Puzzle Defeated", "Just in time!", "Phew!")
		else -> arrayOf("You won!", "You did it!", "Cool!", "Congratulations!")
	}).random()
}
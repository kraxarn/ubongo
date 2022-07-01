package containers

import Resources
import com.soywiz.korge.input.onClick
import com.soywiz.korge.view.Container
import com.soywiz.korge.view.ViewDslMarker
import com.soywiz.korge.view.addTo
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import enums.ResImage

fun Container.pauseDialog(
	res: Resources,
	width: Double, height: Double,
	callback: @ViewDslMarker PauseDialog.() -> Unit = {},
) = PauseDialog(width, height, res).addTo(this, callback)

class PauseDialog(
	width: Double, height: Double,
	res: Resources,
) : Dialog(width, height, res)
{
	val onBack = onLeftAction
	val onResume = onRightAction

	init
	{
		title("Paused")
		body("Game is paused")

		leftAction("Back to menu", SVG(res[ResImage.UI_ARROW_LEFT]).render()) {
			onClick { onBack(it) }
		}

		rightAction("Resume", SVG(res[ResImage.UI_PLAY]).render()) {
			onClick { onResume(it) }
		}
	}
}
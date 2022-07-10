package containers

import com.soywiz.klock.TimeSpan
import com.soywiz.klock.milliseconds
import com.soywiz.korge.tween.get
import com.soywiz.korge.tween.tween
import com.soywiz.korge.view.*
import constants.GameColors

fun Container.backdrop(
	width: Int, height: Int,
	callback: @ViewDslMarker Backdrop.() -> Unit = {},
) = Backdrop(width, height).addTo(this, callback)

class Backdrop(width: Int, height: Int) : Container()
{
	init
	{
		graphics {
			solidRect(width, height, GameColors.dialogBackdrop)
		}

		alpha = 0.0
		visible = false
	}

	suspend fun show(time: TimeSpan = 300.milliseconds)
	{
		visible = true
		tween(::alpha[0.6], time = time)
	}

	suspend fun hide(time: TimeSpan = 300.milliseconds)
	{
		tween(::alpha[0.0], time = time)
		visible = false
	}
}
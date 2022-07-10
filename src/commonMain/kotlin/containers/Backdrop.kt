package containers

import com.soywiz.klock.TimeSpan
import com.soywiz.klock.milliseconds
import com.soywiz.korge.tween.get
import com.soywiz.korge.tween.tween
import com.soywiz.korge.view.Container
import com.soywiz.korge.view.graphics
import com.soywiz.korge.view.solidRect
import constants.GameColors

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
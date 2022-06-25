package extensions.piece

import com.soywiz.klock.milliseconds
import com.soywiz.korge.tween.get
import com.soywiz.korge.tween.tween
import com.soywiz.korge.view.Image
import com.soywiz.korma.geom.degrees
import com.soywiz.korma.geom.plus
import com.soywiz.korma.interpolation.Easing

/**
 * Rotate 90 degrees
 */
suspend fun Image.rotatePiece()
{
	// Don't allow rotation while already rotating
	if (this.rotation.degrees.toInt() % 90 != 0)
	{
		return
	}

	this.tween(
		this::rotation[this.rotation + 90.degrees],
		time = 200.milliseconds,
		easing = Easing.EASE_OUT,
	)
}
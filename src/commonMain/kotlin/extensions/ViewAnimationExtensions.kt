package extensions

import com.soywiz.klock.milliseconds
import com.soywiz.korge.tween.get
import com.soywiz.korge.tween.tween
import com.soywiz.korge.view.View
import com.soywiz.korma.geom.degrees
import com.soywiz.korma.geom.plus
import com.soywiz.korma.interpolation.Easing
import kotlin.math.abs

private val DEFAULT_TIME = 200.milliseconds
private val DEFAULT_EASING = Easing.EASE_OUT

/**
 * Move view by 90 degrees.
 *
 * **Note:** Currently doesn't rotate if already rotating.
 */
suspend fun View.rotate()
{
	// Don't allow rotation while already rotating
	if (this.rotation.degrees.toInt() % 90 != 0) return

	// Always rotate clockwise
	val degrees = (if (this.isMirrored) -90 else 90).degrees

	this.tween(
		this::rotation[this.rotation + degrees],
		time = DEFAULT_TIME,
		easing = DEFAULT_EASING,
	)
}

/**
 * Mirror view, or return if already mirrored.
 *
 * **Note:** Currently doesn't mirror if already mirroring.
 */
suspend fun View.mirror()
{
	// Don't allow scaling while already mirroring
	if (this.isMirroring) return

	this.tween(
		this::scaleX[-this.scaleX],
		time = DEFAULT_TIME,
		easing = DEFAULT_EASING,
	)
}

/**
 * View is currently animating being mirrored
 */
val View.isMirroring get() = abs(this.scaleX) != 1.0

/**
 * View is mirrored
 */
val View.isMirrored get() = this.scaleX < 0
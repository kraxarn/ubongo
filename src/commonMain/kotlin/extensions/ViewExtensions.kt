package extensions

import com.soywiz.klock.milliseconds
import com.soywiz.korge.tween.get
import com.soywiz.korge.tween.tween
import com.soywiz.korge.view.View
import com.soywiz.korge.view.scale
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.degrees
import com.soywiz.korma.geom.plus
import com.soywiz.korma.interpolation.Easing
import kotlin.math.abs
import kotlin.math.max

val View.size2 get() = Point(this.width, this.height)

fun <T : View> T.maxWidthOrHeight(scale: Double) =
	this.scale(scale / max(this.width, this.height))

fun View.containsPoint(point: Point) =
	this.hitShape2d.containsPoint(point.x, point.y)

val DEFAULT_TIME = 200.milliseconds
val DEFAULT_EASING = Easing.EASE_OUT

/**
 * Move view by 90 degrees.
 *
 * **Note:** Currently doesn't rotate if already rotating.
 */
suspend fun View.rotate()
{
	// Don't allow rotation while already rotating
	if (this.rotation.degrees.toInt() % 90 != 0) return

	this.tween(
		this::rotation[this.rotation + 90.degrees],
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
	if (abs(this.scaleX) != 1.0) return

	this.tween(
		this::scaleX[-this.scaleX],
		time = DEFAULT_TIME,
		easing = DEFAULT_EASING,
	)
}
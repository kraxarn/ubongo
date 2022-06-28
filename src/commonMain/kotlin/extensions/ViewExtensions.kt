package extensions

import com.soywiz.korge.view.View
import com.soywiz.korge.view.scale
import com.soywiz.korma.geom.Point
import kotlin.math.max

val View.size2 get() = Point(this.width, this.height)

fun <T : View> T.maxWidthOrHeight(scale: Double) =
	this.scale(scale / max(this.width, this.height))
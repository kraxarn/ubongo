package extensions

import com.soywiz.korge.view.View
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.PointInt
import kotlin.math.round

val View.size2 get() = Point(this.width, this.height)

val View.intPos
	get() = PointInt(round(this.x).toInt(), round(this.y).toInt())
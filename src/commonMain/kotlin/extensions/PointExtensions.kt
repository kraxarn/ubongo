package extensions

import com.soywiz.korma.geom.IPoint
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.PointInt
import kotlin.math.round

operator fun IPoint.rem(value: Double) =
	Point(this.x % value, this.y % value)

fun IPoint.toInt() =
	PointInt(round(this.x).toInt(), round(this.y).toInt())
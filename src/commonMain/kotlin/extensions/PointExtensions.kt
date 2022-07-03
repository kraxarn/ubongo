package extensions

import com.soywiz.korma.geom.IPoint
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.PointInt
import kotlin.math.round

fun IPoint.toInt() =
	PointInt(round(this.x).toInt(), round(this.y).toInt())

fun Sequence<Point>.sum() =
	this.fold(Point(0.0)) { sum, point -> sum + point }
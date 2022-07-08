package extensions

import com.soywiz.korma.geom.IPoint
import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.PointInt
import kotlin.math.max
import kotlin.math.min
import kotlin.math.round

fun IPoint.toInt() =
	PointInt(round(this.x).toInt(), round(this.y).toInt())

fun Sequence<Point>.sum() =
	this.fold(Point(0.0)) { sum, point -> sum + point }

fun min(a: Point, b: Point) = Point(min(a.x, b.x), min(a.y, b.y))

fun max(a: Point, b: Point) = Point(max(a.x, b.x), max(a.y, b.y))

fun Iterable<Point>.topLeft() =
	this.drop(1).fold(this.first()) { a, b -> min(a, b) }

fun Iterable<Point>.bottomRight() =
	this.drop(1).fold(this.first()) { a, b -> max(a, b) }
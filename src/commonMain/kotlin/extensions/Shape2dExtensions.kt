package extensions

import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.shape.Shape2d

fun Shape2d.containsPoint(point: Point) =
	this.containsPoint(point.x, point.y)
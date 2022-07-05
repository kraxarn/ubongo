package extensions

import com.soywiz.korma.geom.IPointInt
import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.times

operator fun IPointInt.times(scale: Int) =
	this * PointInt(scale, scale)

fun Sequence<PointInt>.rotated(origin: PointInt) =
	this.map { PointInt(-it.y + origin.x, it.x - origin.y) }
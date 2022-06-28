package extensions

import com.soywiz.korma.geom.IPoint
import com.soywiz.korma.geom.PointInt
import kotlin.math.round

fun IPoint.toInt() =
	PointInt(round(this.x).toInt(), round(this.y).toInt())
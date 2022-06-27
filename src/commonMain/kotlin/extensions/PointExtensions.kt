package extensions

import com.soywiz.korma.geom.IPoint
import com.soywiz.korma.geom.Point

operator fun IPoint.rem(value: Double) =
	Point(this.x % value, this.y % value)
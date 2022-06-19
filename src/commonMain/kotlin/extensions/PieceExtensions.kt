package extensions

import com.soywiz.korma.geom.PointInt
import com.soywiz.korma.geom.minus
import com.soywiz.korma.geom.plus
import entities.Piece

val Piece.size: PointInt
	get()
	{
		val min = PointInt()
		val max = PointInt()

		for (point in this.points)
		{
			if (point.x < min.x) min.x = point.x
			if (point.x > max.x) max.x = point.x
			if (point.y < min.y) min.y = point.y
			if (point.y > max.y) max.y = point.y
		}

		return max - min + PointInt(1, 1)
	}
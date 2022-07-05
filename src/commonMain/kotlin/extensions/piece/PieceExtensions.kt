package extensions.piece

import com.soywiz.korma.geom.Point
import com.soywiz.korma.geom.PointInt
import enums.PieceShape
import extensions.size2
import kotlin.math.floor

val PieceShape.size get() = this.shape.size2

val PieceShape.points: Sequence<PointInt>
	get() = sequence {
		for (x in 0 until this@points.shape.width)
		{
			for (y in 0 until this@points.shape.height)
			{
				if (this@points.shape[x, y]) yield(PointInt(x, y))
			}
		}
	}

val PieceShape.center
	get() = Point(floor(this.shape.width / 2.0), floor(this.shape.height / 2.0))
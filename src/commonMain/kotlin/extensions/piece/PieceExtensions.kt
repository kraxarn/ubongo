package extensions.piece

import com.soywiz.korma.geom.PointInt
import enums.PieceShape
import extensions.size2

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
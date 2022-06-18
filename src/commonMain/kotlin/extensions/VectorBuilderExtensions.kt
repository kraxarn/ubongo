package extensions

import com.soywiz.korma.geom.vector.VectorBuilder
import com.soywiz.korma.geom.vector.roundRect

public fun VectorBuilder.roundRect(position: Int, size: Int, radius: Int) =
	this.roundRect(position, position, size, size, radius, radius)
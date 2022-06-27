package extensions

import com.soywiz.korge.view.View
import com.soywiz.korma.geom.Point

val View.size2 get() = Point(this.width, this.height)
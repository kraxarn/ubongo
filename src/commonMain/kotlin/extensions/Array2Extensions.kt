package extensions

import com.soywiz.kds.Array2
import com.soywiz.korma.geom.PointInt

val <T> Array2<T>.size2 get() = PointInt(this.width, this.height)
package images

import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.paint.LinearGradientPaint
import com.soywiz.korma.geom.vector.rect
import constants.GameColors

fun background(width: Int, height: Int): Bitmap
{
	return NativeImage(width, height).context2d {
		fill(
			LinearGradientPaint(0, 0, 0, height)
				.add(0.0, GameColors.backgroundStart)
				.add(1.0, GameColors.backgroundEnd)
		) {
			rect(0, 0, width, height)
		}
	}
}
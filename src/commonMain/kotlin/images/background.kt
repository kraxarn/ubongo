package images

import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.component.length.height
import com.soywiz.korge.component.length.lengths
import com.soywiz.korge.component.length.width
import com.soywiz.korge.view.Image
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.paint.LinearGradientPaint
import com.soywiz.korma.geom.vector.rect
import constants.GameColors

@KorgeExperimental
fun background(width: Int, height: Int): Image
{
	val bitmap = NativeImage(width, height).context2d {
		fill(
			LinearGradientPaint(0, 0, 0, height)
				.add(0.0, GameColors.backgroundStart)
				.add(1.0, GameColors.backgroundEnd)
		) {
			rect(0, 0, width, height)
		}
	}

	val image = Image(bitmap)
	image.lengths {
		this.width = 100.vw
		this.height = 100.vh
	}
	return image
}
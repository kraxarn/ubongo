package containers

import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.component.onStageResized
import com.soywiz.korge.scene.Scene
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.NativeImage
import com.soywiz.korim.bitmap.context2d
import com.soywiz.korim.bitmap.slice
import com.soywiz.korim.paint.LinearGradientPaint
import com.soywiz.korma.geom.vector.rect
import constants.GameColors

@KorgeExperimental
fun Scene.background(callback: @ViewDslMarker Background.() -> Unit = {}) =
	Background(this.views).addTo(this.sceneView, callback)

@KorgeExperimental
class Background(private val views: Views) : Container()
{
	private val actualWidth get() = views.actualVirtualWidth
	private val actualHeight get() = views.actualVirtualHeight

	private val actualX get() = views.actualVirtualLeft
	private val actualY get() = views.actualVirtualTop

	private val image = image(bitmap)

	private val bitmap
		get() = NativeImage(actualWidth, actualHeight).context2d {
			fill(
				LinearGradientPaint(0, 0, 0, height)
					.add(0.0, GameColors.backgroundStart)
					.add(1.0, GameColors.backgroundEnd)
			) {
				rect(0, 0, width, height)
			}
		}

	init
	{
		onStageResized { _, _ ->
			// TODO: Stage shouldn't be resized often, so generating a new image should be fine
			image.bitmap = bitmap.slice()
			position(actualX, actualY)
		}
	}
}
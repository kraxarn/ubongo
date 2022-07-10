package containers

import Resources
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.component.onStageResized
import com.soywiz.korge.input.onClick
import com.soywiz.korge.service.storage.storage
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.slice
import com.soywiz.korim.color.Colors
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korim.vector.format.SVG
import com.soywiz.korim.vector.render
import com.soywiz.korio.async.AsyncSignal
import com.soywiz.korma.geom.vector.roundRect
import enums.ResFont
import enums.ResImage
import extensions.maxWidthOrHeight
import extensions.music
import extensions.snapWhileDragging
import extensions.sound

@KorgeExperimental
fun Container.settings(res: Resources, views: Views, callback: @ViewDslMarker Settings.() -> Unit = {}) =
	Settings(res, views).addTo(this, callback)

@KorgeExperimental
class Settings(res: Resources, private val views: Views) : Container()
{
	val onToggleMusic = AsyncSignal<Boolean>()

	private val actualWidth get() = views.virtualWidth.toDouble()
	val actualHeight get() = 580.0

	private val actualX get() = 0.0
	private val actualY get() = views.actualVirtualTop + views.actualVirtualHeight - actualHeight + 16.0

	private val disabledIcon = SVG(res[ResImage.UI_SQUARE]).render()
	private val enabledIcon = SVG(res[ResImage.UI_SQUARE_CHECK]).render()

	private val textSkin: UISkin

	init
	{
		textSkin = UISkin {
			textFont = res[ResFont.REGULAR]
			textSize = 42.0
		}

		graphics {
			fill(Colors["#3f51b5"]) {
				roundRect(0.0, 0.0, actualWidth, actualHeight, 32.0)
			}
		}

		val music = option("Music", views.storage.music) {
			position(PADDING, PADDING)
			onClick {
				views.storage.music = !views.storage.music
				onToggleMusic(views.storage.music)
			}
		}

		val sound = option("Sound effects", views.storage.sound) {
			alignLeftToLeftOf(music)
			alignTopToBottomOf(music, PADDING / 2.0)
			onClick {
				views.storage.sound = !views.storage.sound
			}
		}

		option("Snap to board while dragging", views.storage.snapWhileDragging) {
			alignLeftToLeftOf(sound)
			alignTopToBottomOf(sound, PADDING / 2.0)
			onClick {
				views.storage.snapWhileDragging = !views.storage.snapWhileDragging
			}
		}

		onStageResized { _, _ -> position(actualX, actualY) }
	}

	private fun Container.option(
		text: String, value: Boolean,
		callback: @ViewDslMarker Image.() -> Unit = {},
	): View
	{
		val image = this.image(if (value) enabledIcon else disabledIcon) {
			maxWidthOrHeight(64.0)
			onClick {
				bitmap = (if (bitmap.base == disabledIcon) enabledIcon else disabledIcon).slice()
			}
			callback(this)
		}

		this.uiText(text) {
			uiSkin = textSkin
			textAlignment = TextAlignment.MIDDLE_LEFT
			size(200.0, image.scaledHeight)
			alignTopToTopOf(image)
			alignLeftToRightOf(image, PADDING / 2.0)
		}

		return image
	}

	companion object
	{
		const val PADDING = 96.0
	}
}
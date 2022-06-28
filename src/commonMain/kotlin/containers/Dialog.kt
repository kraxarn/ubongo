package containers

import com.soywiz.korge.input.MouseEvents
import com.soywiz.korge.ui.*
import com.soywiz.korge.view.*
import com.soywiz.korim.bitmap.Bitmap
import com.soywiz.korim.color.Colors
import com.soywiz.korim.font.Font
import com.soywiz.korim.text.TextAlignment
import com.soywiz.korio.async.AsyncSignal
import com.soywiz.korma.geom.vector.roundRect

fun Container.dialog(
	width: Double, height: Double,
	titleFont: Font, bodyFont: Font,
	callback: @ViewDslMarker Dialog.() -> Unit = {},
) = Dialog(width, height, titleFont, bodyFont).addTo(this, callback)

open class Dialog(
	width: Double, height: Double,
	private val titleFont: Font, private val bodyFont: Font,
) : UIContainer(width, height)
{
	val onLeftAction = AsyncSignal<MouseEvents>()
	val onRightAction = AsyncSignal<MouseEvents>()

	private var title: UIText? = null
	private var body: UIText? = null

	init
	{
		graphics {
			fill(Colors["#3f51b5"]) {
				this.roundRect(0.0, 0.0, width, height, BORDER_RADIUS)
			}
		}
	}

	fun title(text: String)
	{
		if (title != null)
		{
			title?.text = text
			return
		}

		title = uiText(text) {
			textFont = titleFont
			textAlignment = TextAlignment.MIDDLE_CENTER
			textSize = TITLE_TEXT_SIZE
			position(PADDING, PADDING)
			size(this@Dialog.width - PADDING * 2, 90.0)
		}
	}

	fun body(text: String)
	{
		if (body != null)
		{
			body?.text = text
			return
		}

		body = uiText(text) {
			textFont = bodyFont
			textAlignment = TextAlignment.TOP_CENTER
			textSize = BODY_TEXT_SIZE
			size(this@Dialog.width - PADDING * 4, 140.0)

			if (title != null)
			{
				alignTopToBottomOf(title!!, PADDING / 2)
				alignLeftToLeftOf(title!!, PADDING)
			}
		}
	}

	private fun actionText(text: String, block: @ViewDslMarker UIText.() -> Unit): UIText
	{
		return uiText(text) {
			textFont = bodyFont
			textAlignment = TextAlignment.TOP_CENTER
			textSize = ACTION_TEXT_SIZE
			size(this@Dialog.width / 2 - PADDING, 52.0)
			alignBottomToBottomOf(this@Dialog, PADDING)
			block(this)
		}
	}

	private fun actionIcon(icon: Bitmap, callback: @ViewDslMarker Image.() -> Unit): Image
	{
		return image(icon, 0.0, -0.25) {
			size(96, 96)
			callback(this)
		}
	}

	fun leftAction(text: String, icon: Bitmap, callback: @ViewDslMarker Image.() -> Unit = {})
	{
		val actionText = actionText(text) { alignLeftToLeftOf(this@Dialog, PADDING) }
		actionIcon(icon) {
			centerXOn(actionText)
			alignBottomToTopOf(actionText, PADDING)
			callback(this)
		}
	}

	fun rightAction(text: String, icon: Bitmap, callback: @ViewDslMarker Image.() -> Unit = {})
	{
		val actionText = actionText(text) { alignRightToRightOf(this@Dialog, PADDING) }
		actionIcon(icon) {
			centerXOn(actionText)
			alignBottomToTopOf(actionText, PADDING)
			callback(this)
		}
	}

	companion object
	{
		const val PADDING = 32.0
		const val BORDER_RADIUS = 12.0
		const val TITLE_TEXT_SIZE = 64.0
		const val BODY_TEXT_SIZE = 42.0
		const val ACTION_TEXT_SIZE = 36.0
	}
}
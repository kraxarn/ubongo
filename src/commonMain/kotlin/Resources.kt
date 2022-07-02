import com.soywiz.klock.measureTime
import com.soywiz.korim.font.Font
import com.soywiz.korim.font.readTtfFont
import com.soywiz.korio.file.std.resourcesVfs
import enums.ResFont
import enums.ResImage
import extensions.logger

class Resources
{
	private val images = mutableMapOf<ResImage, String>()
	private val fonts = mutableMapOf<ResFont, Font>()

	suspend fun loadAll()
	{
		val ms = measureTime {
			for (image in ResImage.values())
			{
				if (images.containsKey(image)) continue
				images[image] = resourcesVfs[image.path].readString()
			}

			for (font in ResFont.values())
			{
				if (fonts.containsKey(font)) continue
				fonts[font] = resourcesVfs[font.path].readTtfFont()
			}
		}.millisecondsInt
		logger.debug { "Loaded in $ms ms" }
	}

	operator fun get(image: ResImage) = images.getValue(image)

	operator fun get(font: ResFont) = fonts.getValue(font)
}
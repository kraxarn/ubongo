import com.soywiz.korio.async.suspendTest
import com.soywiz.korio.lang.assert
import enums.ResFont
import enums.ResImage
import kotlin.test.Test
import kotlin.test.assertFails

class ResourcesTests
{
	@Test
	fun testLoadingResources() = suspendTest {
		val res = Resources()
		assertFails { res[ResImage.UI_ARROW_LEFT] }

		res.loadAll()

		for (image in ResImage.values())
		{
			assert(res[image].isNotBlank()) { image.name }
		}

		for (font in ResFont.values())
		{
			assert(res[font].name.isNotBlank()) { font.name }
		}
	}
}
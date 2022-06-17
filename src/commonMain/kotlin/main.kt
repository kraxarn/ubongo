import com.soywiz.korge.Korge
import com.soywiz.korge.scene.Module
import com.soywiz.korinject.AsyncInjector
import com.soywiz.korma.geom.Anchor
import com.soywiz.korma.geom.ScaleMode
import com.soywiz.korma.geom.SizeInt
import constants.GameColors
import scenes.MenuScene

suspend fun main() = Korge(Korge.Config(Game))

object Game : Module()
{
	override val mainScene get() = MenuScene::class
	override val windowSize get() = SizeInt(540, 960)
	override val size get() = SizeInt(1080, 1920)
	override val bgcolor get() = GameColors.backgroundStart
	override val title get() = "Ubongo"
	override val icon get() = "images/logo.png"
	override val scaleMode get() = ScaleMode.SHOW_ALL
	override val clipBorders get() = false
	override val scaleAnchor get() = Anchor.TOP_LEFT

	override suspend fun AsyncInjector.configure()
	{
		mapPrototype { MenuScene() }
	}
}
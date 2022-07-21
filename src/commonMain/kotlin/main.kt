import com.soywiz.klogger.Logger
import com.soywiz.korge.Korge
import com.soywiz.korge.annotations.KorgeExperimental
import com.soywiz.korge.scene.Module
import com.soywiz.korim.format.PNG
import com.soywiz.korinject.AsyncInjector
import com.soywiz.korio.util.OS
import com.soywiz.korma.geom.Anchor
import com.soywiz.korma.geom.ScaleMode
import com.soywiz.korma.geom.SizeInt
import constants.GameColors
import scenes.GameScene
import scenes.MenuScene

@KorgeExperimental
@ExperimentalUnsignedTypes
suspend fun main() = Korge(Korge.Config(Game))

@KorgeExperimental
@ExperimentalUnsignedTypes
object Game : Module()
{
	override val mainScene get() = MenuScene::class
	override val windowSize get() = SizeInt(540, 1080)
	override val size get() = SizeInt(1080, 1920)
	override val bgcolor get() = GameColors.backgroundStart
	override val title get() = "Ubongo"
	override val icon get() = "images/logo.png"
	override val scaleMode get() = ScaleMode.SHOW_ALL
	override val clipBorders get() = false
	override val scaleAnchor get() = Anchor.CENTER
	override val imageFormats get() = listOf(PNG)

	init
	{
		if (OS.isJvm)
		{
			Logger.defaultLevel = Logger.Level.DEBUG
		}
	}

	override suspend fun AsyncInjector.configure()
	{
		val gameState = GameState()
		gameState.res.loadAll()

		mapInstance(gameState)
		mapPrototype { MenuScene(get()) }
		mapPrototype { GameScene(get()) }
	}
}
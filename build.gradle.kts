import com.soywiz.korge.gradle.KorgeGradlePlugin
import com.soywiz.korge.gradle.Orientation
import com.soywiz.korge.gradle.korge
import java.io.FileInputStream
import java.time.Year.now
import java.util.*

buildscript {
	val korgePluginVersion: String by project

	repositories {
		mavenLocal()
		mavenCentral()
		google()
		maven { url = uri("https://plugins.gradle.org/m2/") }
	}
	dependencies {
		classpath("com.soywiz.korlibs.korge.plugins:korge-gradle-plugin:$korgePluginVersion")
	}
}

apply<KorgeGradlePlugin>()

val projectVersion: String by project

korge {
	id = "com.kraxarn.ubongo"
	name = "Ubongo"
	copyright = "Copyright (c) ${now().value} kraxarn"
	icon = File("src/commonMain/resources/images/logo.png")

	version = projectVersion
	versionCode = projectVersion
		.split('.')
		.joinToString("") { it.padStart(2, '0') }
		.toInt()

	orientation = Orientation.PORTRAIT

	androidMinSdk = 19
	androidCompileSdk = 30
	androidTargetSdk = 30

	val keystorePropertiesFile = rootProject.file("key.properties")
	if (keystorePropertiesFile.exists())
	{
		val properties = Properties()
		properties.load(FileInputStream(keystorePropertiesFile))

		androidReleaseSignStoreFile = properties.getProperty("storeFile")
		androidReleaseSignStorePassword = properties.getProperty("storePassword")
		androidReleaseSignKeyAlias = properties.getProperty("keyAlias")
		androidReleaseSignKeyPassword = properties.getProperty("keyPassword")
	}

	targetJvm()
	targetJs()
	targetDesktop()
	targetIos()
	targetAndroidDirect()
}

val versionFile = File("src/commonMain/kotlin/constants/Application.kt")
versionFile.writer().use {
	it.write("package constants\nobject Application { const val VERSION = \"$projectVersion\" }")
}
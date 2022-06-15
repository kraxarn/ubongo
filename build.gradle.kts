import com.soywiz.korge.gradle.KorgeGradlePlugin
import com.soywiz.korge.gradle.Orientation
import com.soywiz.korge.gradle.korge
import com.soywiz.korge.gradle.util.get
import java.time.Year.now

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

korge {
	id = "com.kraxarn.ubongo2"
	name = "Ubongo"
	copyright = "Copyright (c) ${now().value} kraxarn"

	val version: String by project
	this.version = version

	versionCode = version
		.split('.')
		.joinToString("") { it.padStart(2, '0') }
		.toInt()

	orientation = Orientation.PORTRAIT

	androidMinSdk = 19
	androidCompileSdk = 30
	androidTargetSdk = 30

	targetJvm()
	targetJs()
	targetDesktop()
	targetIos()

	// AndroidDirect is required for version code to be set properly
	// AndroidIndirect is required for atlases to work properly
	targetAndroidIndirect()
}
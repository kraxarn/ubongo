package extensions

import com.soywiz.korge.service.storage.IStorage
import kotlin.properties.ReadWriteProperty
import kotlin.reflect.KProperty

/**
 * If music should play on the main menu
 */
var IStorage.music by property("music", false)

/**
 * If sound effects should play
 */
var IStorage.sound by property("sound", true)

/**
 * Snap to board grid while dragging instead of only when dropped
 */
var IStorage.snapWhileDragging by property("snapWhileDragging", true)

private fun property(key: String, defaultValue: Boolean) = object : ReadWriteProperty<IStorage, Boolean>
{
	override fun getValue(thisRef: IStorage, property: KProperty<*>): Boolean
	{
		return thisRef.getOrNull(key)?.toBoolean() ?: defaultValue
	}

	override fun setValue(thisRef: IStorage, property: KProperty<*>, value: Boolean)
	{
		thisRef[key] = value.toString()
	}
}

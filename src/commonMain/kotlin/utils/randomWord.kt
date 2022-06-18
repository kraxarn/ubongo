package utils

import com.soywiz.korio.file.std.resourcesVfs
import kotlin.random.Random

private suspend fun read(name: String): List<String>
{
	return resourcesVfs["text/english-$name.txt"]
		.readString()
		.split('\n')
}

suspend fun randomWord(seed: Long): String
{
	val random = Random(seed)
	val adjective = read("adjectives").random(random)
	val noun = read("nouns").random(random)
	return "$adjective $noun"
}
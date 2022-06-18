package utils

import com.soywiz.korio.file.std.resourcesVfs
import kotlin.random.Random

private suspend fun read(name: String): List<String>
{
	return resourcesVfs["text/english-$name.txt"]
		.readString()
		.split('\n')
}

private suspend fun readLine(name: String, index: Int): String
{
	val words = read(name)
	return if (words.isEmpty()) "" else words[index % words.size]
}

private suspend fun adjective(index: Int): String
{
	return readLine("adjectives", index)
}

private suspend fun noun(index: Int): String
{
	return readLine("nouns", index)
}

suspend fun randomWord(seed: Long): String
{
	val random = Random(seed)
	return "${adjective(random.nextInt())} ${noun(random.nextInt())}"
}
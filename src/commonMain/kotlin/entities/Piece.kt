package entities

import com.soywiz.kds.Array2

enum class Piece(val shape: Array2<Boolean>)
{
	/**
	 * Large Z
	 */
	Z2(
		shape(
			"""
			##.
			.#.
			.##
		""".trimIndent(),
		),
	),

	/**
	 * Medium L
	 */
	L2(
		shape(
			"""
			##
			#.
			#.
		""".trimIndent(),
		),
	),

	/**
	 * O (square)
	 */
	O(
		shape(
			"""
			##
			##
		""".trimIndent(),
		),
	),

	/**
	 * Small I
	 */
	I1(
		shape(
			"""
			#
			#
		""".trimIndent(),
		),
	),

	/**
	 * Large I
	 */
	I3(
		shape(
			"""
			####
		""".trimIndent(),
		),
	),

	/**
	 * Large T
	 */
	T2(
		shape(
			"""
			####
			..#.
		""".trimIndent(),
		),
	),

	/**
	 * P
	 */
	P(
		shape(
			"""
			###
			.##
		""".trimIndent(),
		),
	),

	/**
	 * Small Z
	 */
	Z1(
		shape(
			"""
			.##
			##.
		""".trimIndent(),
		),
	),

	/**
	 * Medium I
	 */
	I2(
		shape(
			"""
			###
		""".trimIndent(),
		),
	),

	/**
	 * Large L
	 */
	L3(
		shape(
			"""
		####
		...#
	""".trimIndent(),
		),
	),

	/**
	 * Small L
	 */
	L1(
		shape(
			"""
		##
		#.
	""".trimIndent(),
		),
	),
}

private fun shape(map: String): Array2<Boolean> =
	Array2(map) { char, _, _ -> char == '#' }
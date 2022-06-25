package enums

import com.soywiz.kds.Array2

enum class PieceShape(val shape: Array2<Boolean>)
{
	//region I

	/**
	 * Small I
	 */
	I1(
		shape(
			"""
			##
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
	 * Large I
	 */
	I3(
		shape(
			"""
			####
		""".trimIndent(),
		),
	),

	//endregion

	//region L

	/**
	 * Small L
	 */
	L1(
		shape(
			"""
		#.
		##
	""".trimIndent(),
		),
	),

	/**
	 * Medium L
	 */
	L2(
		shape(
			"""
			#-
			#.
			##
		""".trimIndent(),
		),
	),

	/**
	 * Large L
	 */
	L3(
		shape(
			"""
			#.
			#.
			#.
			##
	""".trimIndent(),
		),
	),

	//endregion

	//region O/P

	/**
	 * O (square)
	 */
	O1(
		shape(
			"""
			##
			##
		""".trimIndent(),
		),
	),

	/**
	 * P
	 */
	P1(
		shape(
			"""
			##
			##
			#.
		""".trimIndent(),
		),
	),

	//endregion

	//region T

	T1(
		shape(
			"""
			###
			.#.
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

	//endregion

	//region Z

	/**
	 * Small Z
	 */
	Z1(
		shape(
			"""
			#.
			##
			.#
		""".trimIndent(),
		),
	),

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

	//endregion
}

private fun shape(map: String): Array2<Boolean> =
	Array2(map) { char, _, _ -> char == '#' }
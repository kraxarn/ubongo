package enums

enum class Difficulty(
	/**
	 * If boards should be generated with random rotation
	 * (pieces can still be manually rotated)
	 */
	val rotation: Boolean,

	/**
	 * Number of pieces generated on board, and allowed to be placed
	 */
	val pieceCount: Int,

	/**
	 * Number of tiles in board vertically and horizontally
	 */
	val boardSize: Int,
)
{
	EASY(false, 3, 7),
	MEDIUM(true, 5, 8),
	HARD(true, 8, 9),
}
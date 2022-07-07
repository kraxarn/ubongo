package enums

enum class Difficulty(
	/**
	 * If boards should be generated with random rotation
	 * (pieces can still be manually rotated)
	 */
	val rotation: Boolean,

	/**
	 * If boards should be generated with random mirrored pieces
	 * (pieces can still be manually mirrored)
	 */
	val mirroring: Boolean,

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
	EASY(false, false, 3, 7),
	MEDIUM(true, false, 5, 8),
	HARD(true, true, 8, 9),
}
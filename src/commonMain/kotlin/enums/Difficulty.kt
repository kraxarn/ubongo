package enums

enum class Difficulty(val rotation: Boolean, val pieceCount: Int, val boardSize: Int)
{
	EASY(false, 3, 6),
	MEDIUM(true, 5, 8),
	HARD(true, 8, 10),
}
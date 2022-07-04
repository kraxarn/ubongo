package enums

enum class PieceCountOption(val displayName: String, val value: Int)
{
	EASY("3", 3),
	MEDIUM("5", 5),
	HARD("8", 8),
	ALL("12", 12),
}
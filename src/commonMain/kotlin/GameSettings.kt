import enums.BoardSizeOption
import enums.PieceCountOption
import enums.RotationOption

class GameSettings
{
	/**
	 * If pieces should be allowed to be rotated
	 */
	var rotation = RotationOption.ON

	/**
	 * Number of pieces generated on board, and allowed to be placed
	 */
	var pieceCount = PieceCountOption.MEDIUM

	/**
	 * Number of tiles in board vertically and horizontally
	 */
	var boardSize = BoardSizeOption.MEDIUM
}
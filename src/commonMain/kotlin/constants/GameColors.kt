package constants

import com.soywiz.korim.color.Colors

/**
 * Predefined colors
 */
object GameColors
{
	/**
	 * Background gradient start/top
	 */
	val backgroundStart = Colors["#2196f3"] // Blue 500

	/**
	 * Background gradient end/bottom
	 */
	val backgroundEnd = Colors["#1e88e5"] // Blue 600

	/**
	 * Background color for buttons
	 */
	val buttonBackground = Colors["#3f51b5"] // Indigo 500

	/**
	 * Primary color for labels etc.
	 */
	val foreground = Colors["#263238"]

	/**
	 * Alternative color for buttons etc.
	 */
	val foregroundAlt = Colors["#eceff1"]

	/**
	 * Foreground color for errors
	 */
	val error = Colors["#e86a17"]

	/**
	 * Background for board of pieces
	 */
	val boardBackground = Colors["#03a9f4"] // Light blue 500

	/**
	 * Background for cells in board of pieces
	 */
	val cellBackground = Colors["#29b6f6"] // Light blue 400

	/**
	 * Background for pieces panel
	 */
	val panelBackground = Colors["#d7ccc8"] // Brown 100

	/**
	 * Shadow behind buttons
	 */
	val shadow = Colors["#0000007f"]

	/**
	 * Background color behind dialogs
	 */
	val dialogBackdrop = Colors["#000000"]
}
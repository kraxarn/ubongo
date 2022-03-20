package colors

import "image/color"

var (
	// Background is the main background color
	Background = color.RGBA{
		R: 0x33,
		G: 0x99,
		B: 0xda,
		A: 0xff,
	}

	// Foreground is the primary color for labels etc.
	Foreground = color.RGBA{
		R: 0x26,
		G: 0x32,
		B: 0x38,
		A: 0xff,
	}

	// ForegroundAlt is the alternative color for buttons etc.
	ForegroundAlt = color.RGBA{
		R: 0xec,
		G: 0xef,
		B: 0xf1,
		A: 0xff,
	}

	// ForegroundError is the foreground color for errors
	ForegroundError = color.RGBA{
		R: 0xe8,
		G: 0x6a,
		B: 0x17,
		A: 0xff,
	}
)

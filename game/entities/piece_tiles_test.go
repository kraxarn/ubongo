package entities

import (
	"github.com/kraxarn/ubongo/res"
	"testing"
)

func TestPieceTiles(t *testing.T) {
	// Valid tiles should return some data
	for i := 0; i < res.PieceCount; i++ {
		tiles := PieceTiles(i)
		if len(tiles) <= 0 || len(tiles[0]) <= 0 {
			t.Errorf("tile %d returned empty data", i)
		}
	}

	// Invalid tiles should return empty data
	tiles := PieceTiles(res.PieceCount + 1)
	if len(tiles) > 0 && len(tiles[0]) > 0 {
		t.Error("invalid tile returned non-empty data")
	}
}

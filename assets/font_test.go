package assets

import (
	"testing"
)

func TestLoadAllFonts(t *testing.T) {
	for i := FontRegular; i <= FontDebug; i++ {
		_, err := Font(i, 12)
		if err != nil {
			t.Error(err)
		}
	}
}

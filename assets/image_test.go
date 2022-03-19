package assets

import (
	"os"
	"strings"
	"testing"
)

func TestLoadAllImages(t *testing.T) {
	entries, err := os.ReadDir("./image")
	if err != nil {
		t.Error(err)
	}

	if len(entries) == 0 {
		t.Errorf("no files")
	}

	for _, entry := range entries {
		filename := entry.Name()
		imageName := filename[:strings.Index(filename, ".")]

		img, err := Image(imageName)
		if err != nil {
			t.Error(err)
		}

		w, h := img.Size()
		if w <= 0 || h <= 0 {
			t.Errorf("no image loaded: %s", imageName)
		}
	}
}

package res

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

//go:embed music/music.ogg
var music []byte

const AudioSampleRate = 44_100

func Music() (*vorbis.Stream, error) {
	return vorbis.DecodeWithSampleRate(AudioSampleRate, bytes.NewReader(music))
}

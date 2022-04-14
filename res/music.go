package res

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

//go:embed music/*a.ogg
var music embed.FS

const AudioSampleRate = 44_100

func Music(index int) (*vorbis.Stream, error) {
	reader, err := music.Open(fmt.Sprintf("music/0%da.ogg", index))
	if err != nil {
		return nil, err
	}
	return vorbis.DecodeWithSampleRate(AudioSampleRate, reader)
}

package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/kraxarn/ubongo/resources"
)

type MusicManager struct {
	context *audio.Context
	player  *audio.Player
	streams []*vorbis.Stream
	current int
	volume  float64
}

const min = 1
const max = 3

func NewMusicManager() (*MusicManager, error) {
	var streams []*vorbis.Stream

	for i := min; i <= max; i++ {
		stream, err := resources.Music(i)
		if err != nil {
			return nil, err
		}
		streams = append(streams, stream)
	}

	return &MusicManager{
		context: audio.NewContext(resources.AudioSampleRate),
		streams: streams,
		current: 0,
		volume:  0.25,
	}, nil
}

func (m *MusicManager) Play() error {
	player, err := m.context.NewPlayer(m.streams[m.current])
	if err != nil {
		return err
	}

	m.player = player
	player.SetVolume(m.volume)
	player.Play()
	return nil
}

func (m *MusicManager) Previous() error {
	if m.player == nil {
		return fmt.Errorf("no player")
	}

	if m.current <= min {
		return nil
	}

	m.current--
	return m.Play()
}

func (m *MusicManager) Next() error {
	if m.player == nil {
		return fmt.Errorf("no player")
	}

	if m.current >= max {
		return nil
	}

	m.current++
	return m.Play()
}

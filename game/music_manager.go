package game

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/kraxarn/ubongo/res"
)

type MusicManager struct {
	context *audio.Context
	player  *audio.Player
	stream  *vorbis.Stream
	current int
	volume  float64
}

func NewMusicManager() (*MusicManager, error) {
	stream, err := res.Music()
	if err != nil {
		return nil, err
	}

	return &MusicManager{
		context: audio.NewContext(res.AudioSampleRate),
		stream:  stream,
		current: 0,
		volume:  0.25,
	}, nil
}

func (m *MusicManager) Update() {

}

func (m *MusicManager) Play() error {
	player, err := m.context.NewPlayer(m.stream)
	if err != nil {
		return err
	}

	m.player = player
	player.SetVolume(m.volume)
	player.Play()
	return nil
}

func (m *MusicManager) Stop() error {
	if m.player == nil {
		return nil
	}

	return m.player.Close()
}

func (m *MusicManager) IsPlaying() bool {
	return m.player != nil && m.player.IsPlaying()
}

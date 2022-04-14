package settings

import (
	"strconv"
)

type Settings struct {
	MusicVolume float64
}

func Load() Settings {
	return Settings{
		MusicVolume: getFloat64("musicVolume", 0),
	}
}

func (s *Settings) Save() {
	set("musicVolume", strconv.FormatFloat(s.MusicVolume, 'f', 2, 64))
}

func getFloat64(key string, fallback float64) float64 {
	value, err := strconv.ParseFloat(get(key), 64)
	if err == nil {
		return value
	}
	return fallback
}

func (s *Settings) ToggleMusic(enabled bool) {
	if enabled {
		s.MusicVolume = 0.50
	} else {
		s.MusicVolume = 0
	}
}

func (s *Settings) IsMusicEnabled() bool {
	return s.MusicVolume > 0
}

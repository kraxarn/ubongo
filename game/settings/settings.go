package settings

import (
	"strconv"
)

type Settings struct {
	MusicVolume float32
}

func Load() Settings {
	return Settings{
		MusicVolume: getFloat32("musicVolume", 0),
	}
}

func (s *Settings) Save() {
	set("musicVolume", strconv.FormatFloat(float64(s.MusicVolume), 'f', 2, 32))
}

func getFloat32(key string, fallback float32) float32 {
	value, err := strconv.ParseFloat(get(key), 32)
	if err == nil {
		return float32(value)
	}
	return fallback
}

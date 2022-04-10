package settings

import (
	"fmt"
	"strconv"
	"time"
)

const key = "timestamp"

type Settings struct {
	timestamp int64
}

func Load() Settings {
	data := get(key)
	timestamp, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return Settings{
			timestamp: time.Now().Unix(),
		}
	}
	return Settings{
		timestamp: timestamp,
	}
}

func (s *Settings) Save() {
	set(key, fmt.Sprintf("%d", s.timestamp))
}

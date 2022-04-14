//go:build !(js && wasm)

package settings

import (
	"encoding/json"
	"github.com/kraxarn/ubongo/game/app"
	"os"
	"path"
)

func configDir() string {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "."
	}
	appPath := path.Join(dir, app.Name)
	if err = os.MkdirAll(appPath, perm()); err != nil {
		return "."
	}
	return appPath
}

func configFile() string {
	return path.Join(configDir(), "settings.json")
}

func configs() map[string]string {
	data, err := os.ReadFile(configFile())
	if err != nil {
		return map[string]string{}
	}

	var cfg map[string]string
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return map[string]string{}
	}
	return cfg
}

func perm() os.FileMode {
	dir, err := os.UserConfigDir()
	if err != nil {
		return 0755
	}
	stat, err := os.Stat(dir)
	if err != nil {
		return 0755
	}
	return stat.Mode()
}

func get(key string) string {
	return configs()[key]
}

func set(key, value string) {
	cfg := configs()
	cfg[key] = value

	data, err := json.Marshal(cfg)
	if err == nil {
		_ = os.WriteFile(configFile(), data, perm())
	}
}

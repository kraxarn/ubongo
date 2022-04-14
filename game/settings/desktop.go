//go:build !(js && wasm)

package settings

import (
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
	data, err := os.ReadFile(path.Join(configDir(), key))
	if err != nil {
		return ""
	}
	return string(data)
}

func set(key, value string) {
	_ = os.WriteFile(path.Join(configDir(), key), []byte(value), perm())
}

package app

import "fmt"

const (
	Name  = "Ubongo"
	major = 0
	minor = 1
	patch = 0
)

func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

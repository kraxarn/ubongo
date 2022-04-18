package app

import "fmt"

const (
	Name  = "Ubongo"
	major = 0
	minor = 4
	patch = 1
)

func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

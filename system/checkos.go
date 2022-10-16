package system

import (
	"runtime"
)

// CheckOS checks the operating system and returns the name of the operating system.
func CheckOS() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "MacOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return "Unknown"
	}
}

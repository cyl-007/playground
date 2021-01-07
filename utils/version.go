package utils

import "os"

var version = "unknown"

func init() {
	resolveVersion()
}

func GetVersion() string {
	return version
}

func resolveVersion() {
	if v := os.Getenv("VERSION"); v != "" {
		version = v
	}
}

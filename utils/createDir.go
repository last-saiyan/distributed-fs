package utils

import (
	"os"
)

// CreateDir creates a dir
func CreateDir(path string) error {
	return os.Mkdir(path, 0640)
}

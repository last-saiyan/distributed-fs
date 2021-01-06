package utils

import (
	"fmt"
	"os"
)

// CreateDir creates a dir
func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0640)
		if err != nil {
			fmt.Println(err, "err creating file ", path)
		}
	}
}

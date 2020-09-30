package client

import (
	"fmt"
)

// Read a file
// returns bytes of the file
func Read(fileName string) []byte {
	fmt.Print("read file")
	return nil
}

// Write a new file
// returns true if successful
// returns false if error
func Write(fileName string, data []byte) bool {
	fmt.Print("read file")
	return false
}

// Append to a existing file
// returns true if successful
// returns false if error
func Append(fileName string) bool {
	return false
}

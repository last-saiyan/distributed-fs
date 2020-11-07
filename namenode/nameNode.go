package namenode

import (
	"fmt"
)

type fileName = string
type chunkName string
type ipAddr string

var filenameTochunk map[fileName][]chunkName
var chunkNameToServer map[chunkName][]ipAddr

// Init the namenode datastructures,
// todo recover namenode from crash
func Init() {
	filenameTochunk = make(map[fileName][]chunkName)
	chunkNameToServer = make(map[chunkName][]ipAddr)
}

// AppendToFile appends to existing file or creates a new file
func AppendToFile(name string) string {
	gfc := fileName(name)
	fmt.Print("asdf", gfc)
	return "asdf"
}

// GetFileLocation asdf
func GetFileLocation(fileName string) map[string][]string {
	return nil
}

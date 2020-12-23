package namenode

import (
	"fmt"
)

// Block uhb
type Block struct {
	name blockName
	gs   int
}

// BlockMeta wewer
type BlockMeta struct {
	name     blockName
	fileSize int
	addr     ipAddr
}

type blockName string
type fileName string
type ipAddr string

// this data needs to be persisted in disk
// for recovery of namenode
var fileToBlock map[fileName][]Block

// this is not necessary to be in disk
// this can be obtained from datanode blockreport()
var blockToLocation map[blockName][]BlockMeta

// Init the namenode datastructures,
// todo recover namenode from crash
func Init() {
	fileToBlock = make(map[fileName][]Block)
	blockToLocation = make(map[blockName][]BlockMeta)
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

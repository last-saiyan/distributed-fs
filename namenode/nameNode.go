package namenode

import (
	"errors"
	"time"
)

// Block uhb
type Block struct {
	name blockName
	gs   int64
}

var BlockSize = 100

// BlockMeta wewer
type BlockMeta struct {
	name     blockName
	fileSize int
	addr     ipAddr
	state    string
}

type blockName string
type fileName string
type ipAddr string

// DatanodeMeta metadata of datanode
type DatanodeMeta struct {
	addr   ipAddr
	memory int
}

// this data needs to be persisted in disk
// for recovery of namenode
var fileToBlock map[fileName][]Block

// this is not necessary to be in disk
// this can be obtained from datanode blockreport()
var blockToLocation map[blockName][]BlockMeta

var datanodeList []DatanodeMeta

// ErrFileNotFound file is not found
var ErrFileNotFound = errors.New("file not found")

// ErrFileExists file with this name is already present
var ErrFileExists = errors.New("file with this name is present")

// ErrFileOpen asd
var ErrFileOpen = errors.New("file is not yet closed")

// Init the namenode datastructures,
// todo recover namenode from crash
func Init() {
	fileToBlock = make(map[fileName][]Block)
	blockToLocation = make(map[blockName][]BlockMeta)
}

// AddDataNode adds ip to list of datanodeList
func AddDataNode(meta DatanodeMeta) {
	datanodeList = append(datanodeList, meta)
}

// pick datanode with most memory
func pickDataNodeToAddNewBlock(count int) []DatanodeMeta {
	return nil
}

// AppendBlock appends to existing file or creates a new file
func AppendBlock(name string) ([]BlockMeta, error) {
	file := fileName(name)
	blockArr, found := fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}
	blkName := blockName(name + "-" + string(len(blockArr)))
	blk := Block{name: blkName, gs: time.Now().Unix()}
	blockArr = append(blockArr, blk)
	fileToBlock[file] = blockArr

	datanodes := pickDataNodeToAddNewBlock(5)
	blockMetaArr := make([]BlockMeta, 0)
	for _, datanode := range datanodes {
		blkMeta := BlockMeta{name: blkName, addr: datanode.addr, fileSize: 0}
		blockMetaArr = append(blockMetaArr, blkMeta)
	}
	blockToLocation[blkName] = blockMetaArr
	return blockMetaArr, nil
}

// GetFileLocation asdf
func GetFileLocation(fileName string) map[string][]string {
	return nil
}

// CreateFile todo ! research about this
func CreateFile(name string) error {
	file := fileName(name)
	_, found := fileToBlock[file]
	if found {
		return ErrFileExists
	}
	blckArr := make([]Block, 0)
	fileToBlock[file] = blckArr
	return nil
}

// WriteToFile returns datanode and block name
func WriteToFile(name string) ([]BlockMeta, error) {
	file := fileName(name)
	blockArr, found := fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}
	lastBlock := blockArr[len(blockArr)-1]
	// perform checks on blockMetaArr if new block has to be created
	blockMetaArr, found := blockToLocation[lastBlock.name]
	for _, blockMeta := range blockMetaArr {
		if blockMeta.state != "committed" {
			return nil, ErrFileOpen
		}
	}
	if blockMetaArr[0].fileSize == BlockSize {
		return AppendBlock(name)
	}
	return blockMetaArr, nil
}

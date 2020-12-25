package namenode

import (
	"sort"
	"strconv"
	"time"
)

// Block uhb
type Block struct {
	name blockName
	gs   int64
}

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
	Addr      ipAddr
	DiskUsage int
}

// NameNode struct to interact with the namenode
type NameNode struct {
	// fileToBlock data needs to be persisted in disk
	// for recovery of namenode
	fileToBlock map[fileName][]Block

	// blockToLocation is not necessary to be in disk
	// blockToLocation can be obtained from datanode blockreport()
	blockToLocation map[blockName][]BlockMeta

	// datanodeList contains list of datanode ipAddr
	datanodeList []DatanodeMeta

	blockSize         int
	replicationFactor int
}

// Init the namenode datastructures,
// todo recover namenode from crash
func (nn *NameNode) Init(blockSize int, replicationFactor int) {
	nn.fileToBlock = make(map[fileName][]Block)
	nn.blockToLocation = make(map[blockName][]BlockMeta)
	nn.blockSize = blockSize
	nn.replicationFactor = replicationFactor
}

// RegisterDataNode adds ip to list of datanodeList
func (nn *NameNode) RegisterDataNode(meta DatanodeMeta) {
	nn.datanodeList = append(nn.datanodeList, meta)
}

// pick datanode with most memory
func (nn *NameNode) pickDataNodeToAddNewBlock(count int) ([]DatanodeMeta, error) {
	if len(nn.datanodeList) < nn.replicationFactor {
		return nil, ErrReplicaCount
	}
	sort.SliceStable(nn.datanodeList, func(i, j int) bool {
		return nn.datanodeList[i].DiskUsage < nn.datanodeList[j].DiskUsage
	})
	return nn.datanodeList[0:count], nil
}

func createBlockMeta(name blockName, addr ipAddr, fileSize int, state string) BlockMeta {
	blk := BlockMeta{}
	blk.name = name
	blk.addr = addr
	blk.fileSize = fileSize
	blk.state = state
	return blk
}

// appendBlock appends new block to the file returns BlockMeta array
func (nn *NameNode) appendBlock(name string) ([]BlockMeta, error) {
	file := fileName(name)
	blockArr, found := nn.fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}
	blkName := blockName(name + "-" + strconv.Itoa(len(blockArr)))
	blk := Block{name: blkName, gs: time.Now().Unix()}
	blockArr = append(blockArr, blk)
	nn.fileToBlock[file] = blockArr

	datanodes, err := nn.pickDataNodeToAddNewBlock(nn.replicationFactor)
	if err != nil {
		return nil, err
	}
	blockMetaArr := make([]BlockMeta, 0)
	for _, datanode := range datanodes {
		blkMeta := createBlockMeta(blkName, datanode.Addr, 0, "pending")
		blockMetaArr = append(blockMetaArr, blkMeta)
	}
	nn.blockToLocation[blkName] = blockMetaArr
	return blockMetaArr, nil
}

// GetFileLocation asdf
func GetFileLocation(fileName string) map[string][]string {
	return nil
}

// CreateFile todo ! research about this
func (nn *NameNode) CreateFile(name string) ([]BlockMeta, error) {
	file := fileName(name)
	_, found := nn.fileToBlock[file]
	if found {
		return nil, ErrFileExists
	}
	blckArr := make([]Block, 0)
	nn.fileToBlock[file] = blckArr

	blockMetaArr, err := nn.appendBlock(name)
	if err != nil {
		delete(nn.fileToBlock, file)
		return nil, err
	}
	return blockMetaArr, nil
}

// WriteToFile returns datanode and block name
func (nn *NameNode) WriteToFile(name string) ([]BlockMeta, error) {
	file := fileName(name)
	blockArr, found := nn.fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}
	lastBlock := blockArr[len(blockArr)-1]
	// perform checks on blockMetaArr if new block has to be created
	blockMetaArr, found := nn.blockToLocation[lastBlock.name]
	for _, blockMeta := range blockMetaArr {
		if blockMeta.state != "committed" {
			return nil, ErrFileOpen
		}
	}
	if blockMetaArr[0].fileSize == nn.blockSize {
		return nn.appendBlock(name)
	}
	return blockMetaArr, nil
}

// Complete is called when datanode completes
func (nn *NameNode) Complete(blkName string, dataNodeAddr string, memory int) error {
	blockArr, found := nn.blockToLocation[blockName(blkName)]
	if !found {
		return ErrFileNotFound
	}
	for i := 0; i < len(blockArr); i++ {
		if blockArr[i].addr == ipAddr(dataNodeAddr) {
			blockArr[i].state = "committed"
		}
	}
	// fmt.Println(blockArr, blkName)
	return nil
}

// BlockReport is status that datanode sends periodically
// based on this data namenode asks the datanode to
// replicate, delete blocks
func (nn *NameNode) BlockReport() {

}

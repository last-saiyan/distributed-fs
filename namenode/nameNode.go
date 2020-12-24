package namenode

import (
	"sort"
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
	addr      ipAddr
	diskUsage int
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

// AddDataNode adds ip to list of datanodeList
func (nn *NameNode) AddDataNode(meta DatanodeMeta) {
	nn.datanodeList = append(nn.datanodeList, meta)
}

// pick datanode with most memory
func (nn *NameNode) pickDataNodeToAddNewBlock(count int) ([]DatanodeMeta, error) {
	if len(nn.datanodeList) < nn.replicationFactor {
		return nil, ErrReplicaCount
	}
	sort.SliceStable(nn.datanodeList, func(i, j int) bool {
		return nn.datanodeList[i].diskUsage < nn.datanodeList[j].diskUsage
	})
	return nn.datanodeList[0:count], nil
}

// AppendBlock appends to existing file or creates a new file
func (nn *NameNode) AppendBlock(name string) ([]BlockMeta, error) {
	file := fileName(name)
	blockArr, found := nn.fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}
	blkName := blockName(name + "-" + string(len(blockArr)))
	blk := Block{name: blkName, gs: time.Now().Unix()}
	blockArr = append(blockArr, blk)
	nn.fileToBlock[file] = blockArr

	datanodes, err := nn.pickDataNodeToAddNewBlock(nn.replicationFactor)
	if err != nil {
		return nil, err
	}
	blockMetaArr := make([]BlockMeta, 0)
	for _, datanode := range datanodes {
		blkMeta := BlockMeta{name: blkName, addr: datanode.addr, fileSize: 0}
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
func (nn *NameNode) CreateFile(name string) error {
	file := fileName(name)
	_, found := nn.fileToBlock[file]
	if found {
		return ErrFileExists
	}
	blckArr := make([]Block, 0)
	nn.fileToBlock[file] = blckArr
	return nil
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
		return nn.AppendBlock(name)
	}
	return blockMetaArr, nil
}

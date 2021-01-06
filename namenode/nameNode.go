package namenode

import (
	"dfs/proto"
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
	ipAddr   string
	state    string
}

type blockName string
type fileName string

// type IpAddr string

// DatanodeMeta metadata of datanode
type DatanodeMeta struct {
	IPAddr    string
	DiskUsage uint64
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

func createBlockMeta(name blockName, ipAddr string, fileSize int, state string) BlockMeta {
	blk := BlockMeta{}
	blk.name = name
	blk.ipAddr = ipAddr
	blk.fileSize = fileSize
	blk.state = state
	return blk
}

// appendBlock appends new block to the file returns BlockMeta array
func (nn *NameNode) appendBlock(name string) (*proto.FileLocationArr, error) {
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
		blkMeta := createBlockMeta(blkName, datanode.IPAddr, 0, "pending")
		blockMetaArr = append(blockMetaArr, blkMeta)
	}
	nn.blockToLocation[blkName] = blockMetaArr
	return convertBlockMetaToProtoFileLocation(blockArr, nn.blockToLocation), nil
}

// GetFileLocation returns block names that consist a file
func (nn *NameNode) GetFileLocation(file string) (*proto.FileLocationArr, error) {
	blockArr, found := nn.fileToBlock[fileName(file)]
	if !found {
		return nil, ErrFileNotFound
	}

	fileBlocksList := make([]*proto.BlockReplicaList, 0)
	for _, blk := range blockArr {
		blockMetaArr := nn.blockToLocation[blk.name]
		blockReplicasArr := make([]*proto.BlockLocation, 0)
		for _, blockMeta := range blockMetaArr {
			// todo add check to see if datanode is up
			blockLocation := &proto.BlockLocation{BlockName: string(blockMeta.name), IpAddr: string(blockMeta.ipAddr)}
			blockReplicasArr = append(blockReplicasArr, blockLocation)
		}
		blockReplicas := &proto.BlockReplicaList{BlockReplicaList: blockReplicasArr}
		fileBlocksList = append(fileBlocksList, blockReplicas)
	}
	return &proto.FileLocationArr{FileBlocksList: fileBlocksList}, nil
}

// CreateFile todo ! research about this
func (nn *NameNode) CreateFile(name string) (*proto.FileLocationArr, error) {
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
func (nn *NameNode) WriteToFile(name string) (*proto.FileLocationArr, error) {
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

	for i := 0; i < len(blockMetaArr); i++ {
		blockMetaArr[i].state = "pending"
	}
	return convertBlockMetaToProtoFileLocation(blockArr, nn.blockToLocation), nil
}

// blockToLocation => blockToReplicaLocation
func convertBlockMetaToProtoFileLocation(blockArr []Block, blockToLocation map[blockName][]BlockMeta) *proto.FileLocationArr {
	lastBlock := blockArr[len(blockArr)-1]
	blockMetaArr := blockToLocation[lastBlock.name]
	fileBlocksList := make([]*proto.BlockReplicaList, 0)

	blockLocationReplicas := make([]*proto.BlockLocation, 0)
	for _, blockMeta := range blockMetaArr {
		// todo add check to see if datanode is up
		blockLocation := &proto.BlockLocation{BlockName: string(blockMeta.name), IpAddr: string(blockMeta.ipAddr)}
		blockLocationReplicas = append(blockLocationReplicas, blockLocation)
	}

	blockReplicas := &proto.BlockReplicaList{BlockReplicaList: blockLocationReplicas}
	fileBlocksList = append(fileBlocksList, blockReplicas)

	return &proto.FileLocationArr{FileBlocksList: fileBlocksList}

}

// Complete is called when datanode completes
func (nn *NameNode) Complete(blkName string, dataNodeAddr string, fileSize int) error {
	blockArr, found := nn.blockToLocation[blockName(blkName)]
	if !found {
		return ErrFileNotFound
	}
	for i := 0; i < len(blockArr); i++ {
		if blockArr[i].ipAddr == dataNodeAddr {
			blockArr[i].state = "committed"
			blockArr[i].fileSize = fileSize
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

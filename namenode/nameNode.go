package namenode

import (
	"dfs/proto"
	"dfs/utils"
	"sort"
	"strconv"
	"time"
)

// Block uhb
type blockMeta struct {
	blockName string
	gs        int64
	blockID   int
}

// BlockMeta wewer
type replicaMeta struct {
	blockName string
	fileSize  int
	ipAddr    string
	state     string
	replicaID int
}

// DatanodeMeta metadata of datanode
type DatanodeMeta struct {
	IPAddr             string
	DiskUsage          uint64
	heartbeatTimeStamp int64
	status             datanodeStatus
}

type datanodeStatus = string

const (
	datanodeDown = datanodeStatus("datanodeDown")
	datanodeUp   = datanodeStatus("datanodeUp")
)

// NameNode struct to interact with the namenode
type NameNode struct {
	// fileToBlock data needs to be persisted in disk
	// for recovery of namenode
	fileToBlock map[string][]blockMeta

	// blockToLocation is not necessary to be in disk
	// blockToLocation can be obtained from datanode blockreport()
	blockToLocation map[string][]replicaMeta

	// datanodeList contains list of datanode ipAddr
	datanodeList []DatanodeMeta

	blockSize         int
	replicationFactor int
}

// Init the namenode datastructures,
// todo recover namenode from crash
func (nn *NameNode) Init(blockSize int, replicationFactor int) {
	nn.fileToBlock = make(map[string][]blockMeta)
	nn.blockToLocation = make(map[string][]replicaMeta)
	nn.blockSize = blockSize
	nn.replicationFactor = replicationFactor
	nn.heartbeatMonitor()
}

// RegisterDataNode adds ip to list of datanodeList
func (nn *NameNode) RegisterDataNode(datanodeIPAddr string, diskUsage uint64) {
	meta := DatanodeMeta{IPAddr: datanodeIPAddr, DiskUsage: diskUsage, heartbeatTimeStamp: time.Now().Unix(), status: datanodeUp}
	// meta.heartbeatTimeStamp = time.Now().Unix()
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

func createBlockMeta(blockName string, ipAddr string, fileSize int, state string) replicaMeta {
	blk := replicaMeta{}
	blk.blockName = blockName
	blk.ipAddr = ipAddr
	blk.fileSize = fileSize
	blk.state = state
	return blk
}

// appendBlock appends new block to the file returns BlockMeta array
func (nn *NameNode) appendBlock(file string) (*proto.FileLocationArr, error) {
	blockArr, found := nn.fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}
	blkName := (file + "-" + strconv.Itoa(len(blockArr)))
	blk := blockMeta{blockName: blkName, gs: time.Now().Unix(), blockID: len(blockArr)}
	blockArr = append(blockArr, blk)
	nn.fileToBlock[file] = blockArr

	datanodes, err := nn.pickDataNodeToAddNewBlock(nn.replicationFactor)
	if err != nil {
		return nil, err
	}
	blockMetaArr := make([]replicaMeta, 0)
	for _, datanode := range datanodes {
		blkMeta := createBlockMeta(blkName, datanode.IPAddr, 0, "pending")
		blockMetaArr = append(blockMetaArr, blkMeta)
	}
	nn.blockToLocation[blkName] = blockMetaArr
	return convertBlockMetaToProtoFileLocation(blockArr, nn.blockToLocation), nil
}

// GetFileLocation returns block names that consist a file
func (nn *NameNode) GetFileLocation(file string) (*proto.FileLocationArr, error) {
	blockArr, found := nn.fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}

	fileBlocksList := make([]*proto.BlockReplicaList, 0)
	for _, blk := range blockArr {
		blockMetaArr := nn.blockToLocation[blk.blockName]
		blockReplicasArr := make([]*proto.BlockLocation, 0)
		for _, blockMeta := range blockMetaArr {
			// todo add check to see if datanode is up
			blockLocation := &proto.BlockLocation{BlockName: blockMeta.blockName, IpAddr: blockMeta.ipAddr}
			blockReplicasArr = append(blockReplicasArr, blockLocation)
		}
		blockReplicas := &proto.BlockReplicaList{BlockReplicaList: blockReplicasArr}
		fileBlocksList = append(fileBlocksList, blockReplicas)
	}
	return &proto.FileLocationArr{FileBlocksList: fileBlocksList}, nil
}

// CreateFile todo ! research about this
func (nn *NameNode) CreateFile(name string) (*proto.FileLocationArr, error) {
	file := name
	_, found := nn.fileToBlock[file]
	if found {
		return nil, ErrFileExists
	}
	blckArr := make([]blockMeta, 0)
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
	file := name
	blockArr, found := nn.fileToBlock[file]
	if !found {
		return nil, ErrFileNotFound
	}
	lastBlock := blockArr[len(blockArr)-1]
	// perform checks on blockMetaArr if new block has to be created
	blockMetaArr, found := nn.blockToLocation[lastBlock.blockName]
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
func convertBlockMetaToProtoFileLocation(blockArr []blockMeta, blockToLocation map[string][]replicaMeta) *proto.FileLocationArr {
	lastBlock := blockArr[len(blockArr)-1]
	blockMetaArr := blockToLocation[lastBlock.blockName]
	fileBlocksList := make([]*proto.BlockReplicaList, 0)

	blockLocationReplicas := make([]*proto.BlockLocation, 0)
	for _, blockMeta := range blockMetaArr {
		// todo add check to see if datanode is up
		blockLocation := &proto.BlockLocation{BlockName: blockMeta.blockName, IpAddr: blockMeta.ipAddr}
		blockLocationReplicas = append(blockLocationReplicas, blockLocation)
	}

	blockReplicas := &proto.BlockReplicaList{BlockReplicaList: blockLocationReplicas}
	fileBlocksList = append(fileBlocksList, blockReplicas)

	return &proto.FileLocationArr{FileBlocksList: fileBlocksList}

}

// Complete is called when datanode completes
func (nn *NameNode) Complete(blkName string, dataNodeAddr string, fileSize int) error {
	blockArr, found := nn.blockToLocation[blkName]
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

// Heartbeat is called when namenode receives heartbeat
func (nn *NameNode) Heartbeat(datanodeIPAddr string, diskUsage uint64) {
	for id, datanode := range nn.datanodeList {
		if datanode.IPAddr == datanodeIPAddr {
			nn.datanodeList[id].heartbeatTimeStamp = time.Now().Unix()
		}
	}
}

// heartbeatMonitor is a recursive function to check if datanode is available
func (nn *NameNode) heartbeatMonitor() {
	heartbeatTimeout := utils.GetConfig().HeartbeatTimeout
	heartbeatTimeoutDuration := time.Second * time.Duration(heartbeatTimeout)
	time.Sleep(heartbeatTimeoutDuration)

	for id, datanode := range nn.datanodeList {
		if time.Since(time.Unix(datanode.heartbeatTimeStamp, 0)) > heartbeatTimeoutDuration {
			nn.datanodeList[id].status = datanodeDown
		}
	}
	nn.heartbeatMonitor()
}

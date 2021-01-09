package namenode

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	nn := NameNode{}
	nn.Init(1000, 3) // without datanode
	fileName := "testFile"
	fileLocationArr, err := nn.CreateFile(fileName)
	if err == ErrReplicaCount {
		t.Logf("throws ErrReplicaCount count err if datanode count is less than replicas")
	} else {
		t.Errorf("does not throw ErrReplicaCount err")
	}

	// adding datanodes
	dn := DatanodeMeta{IPAddr: "123", DiskUsage: 23}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "234", DiskUsage: 3}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "122", DiskUsage: 53}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "121", DiskUsage: 28}
	nn.RegisterDataNode(dn)
	fileLocationArr, err = nn.CreateFile(fileName)
	if err != nil {
		t.Errorf("creating file has err")
		t.Log(err)
	}
	allBlocksPending := true
	allBlocksZero := true
	BlocksList := fileLocationArr.FileBlocksList
	for _, block := range BlocksList {
		for _, replica := range block.BlockReplicaList {
			if replica.BlockSize != 0 {
				t.Errorf("created block size is not zero")
				allBlocksZero = false
			}
		}
	}

	blkArr := nn.fileToBlock[fileName]
	for _, blk := range blkArr {
		replicas := nn.blockToLocation[blk.blockName]
		for _, replica := range replicas {
			if replica.state != "pending" {
				t.Errorf("created block is not in pending state")
				allBlocksPending = false
			}
		}
	}
	if allBlocksPending {
		t.Logf("created blocks are in pending state")
	}
	if allBlocksZero {
		t.Logf("created blocks have size zero")
	}

	fileLocationArr, err = nn.CreateFile(fileName)
	if err == ErrFileExists {
		t.Logf("duplicate files are not created")
	} else {
		t.Errorf("error duplicate files are created")
	}
}

func TestWriteFile(t *testing.T) {
	fileName := "testFile"
	nn := NameNode{}
	nn.Init(1000, 3)
	// CANNOT WRITE IF FILE IS NOT PRESENT
	dn := DatanodeMeta{IPAddr: "123", DiskUsage: 23}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "234", DiskUsage: 3}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "122", DiskUsage: 53}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "121", DiskUsage: 28}
	nn.RegisterDataNode(dn)
	fileLocationArr, err := nn.WriteToFile(fileName)
	if err == ErrFileNotFound {
		t.Logf("cant write to files that is not yet created")
	} else {
		t.Errorf("error writing to file that is not yet created")
	}

	// CANNOT WRITE TO OPEN FILE
	fileLocationArr, err = nn.CreateFile(fileName)
	_, err = nn.WriteToFile(fileName)
	if err == ErrFileOpen {
		t.Logf("cant write to pending file")
	} else {
		t.Errorf("error able to write to a pending file")
	}

	// COMPLETE SIGNAL FOR ALL REPLICAS
	BlocksList := fileLocationArr.FileBlocksList
	for _, block := range BlocksList {
		for _, replica := range block.BlockReplicaList {
			err := nn.Complete(replica.BlockName, replica.IpAddr, 0)
			if err != nil {
				t.Errorf("error when Complete() the file")
				t.Log(err)
			}
		}
	}

	fileLocationArr, err = nn.WriteToFile(fileName)
	if err == ErrFileOpen {
		t.Errorf("error cant write to pending file")
	} else if err != nil {
		t.Errorf("error when writing to a file")
	}

	BlocksList = fileLocationArr.FileBlocksList
	for _, block := range BlocksList {
		for _, replica := range block.BlockReplicaList {
			err := nn.Complete(replica.BlockName, replica.IpAddr, 10)
			if err != nil {
				t.Errorf("error when Complete() the file")
				t.Log(err)
			}
		}
	}
}

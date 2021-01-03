package namenode

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	nn := NameNode{}
	nn.Init(1000, 3) // without datanode
	fileName := "testFile"
	blkArr, err := nn.CreateFile(fileName)
	if err == ErrReplicaCount {
		t.Logf("throws ErrReplicaCount count err if datanode count is less")
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
	blkArr, err = nn.CreateFile(fileName)
	if err != nil {
		t.Errorf("creating file has err")
		t.Log(err)
	}
	allBlocksPending := true
	allBlocksZero := true

	for _, block := range blkArr {
		if block.state != "pending" {
			t.Errorf("created block is not in pending state")
			allBlocksPending = false
		}
		if block.fileSize != 0 {
			t.Errorf("created block size is not zero")
			allBlocksZero = false
		}
	}
	if allBlocksPending {
		t.Logf("created blocks are in pending state")
	}
	if allBlocksZero {
		t.Logf("created blocks have size zero")
	}

	blkArr, err = nn.CreateFile(fileName)

	if err == ErrFileExists {
		t.Logf("duplicate files are not created")
	} else {
		t.Errorf("error duplicate files are created")
	}
}

func TestWriteFile(t *testing.T) {
	nn := NameNode{}
	nn.Init(1000, 3)
	dn := DatanodeMeta{IPAddr: "123", DiskUsage: 23}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "234", DiskUsage: 3}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "122", DiskUsage: 53}
	nn.RegisterDataNode(dn)
	dn = DatanodeMeta{IPAddr: "121", DiskUsage: 28}
	nn.RegisterDataNode(dn)
	fileName := "testFile"

	blkArr, err := nn.WriteToFile(fileName)

	if err == ErrFileNotFound {
		t.Logf("cant write to files that is not yet created")
	} else {
		t.Errorf("error writing to file that is not yet created")
	}

	blkArr, err = nn.CreateFile(fileName)
	tempBlkArr := blkArr
	blkArr, err = nn.WriteToFile(fileName)
	if err == ErrFileOpen {
		t.Logf("cant write to pending file")
	} else {
		t.Errorf("error able to write to a pending file")
	}

	for _, block := range tempBlkArr.FileBlocksList {
		err := nn.Complete(string(block.name), string(block.addr), 10)
		if err != nil {
			t.Errorf("error when Complete the file")
			t.Log(err)
		}
	}

	blkArr, err = nn.WriteToFile(fileName)
	if err != nil {
		t.Errorf("error when writing to file")
		t.Log(err)
	}
	allBlocksPending := true
	for _, block := range blkArr.FileBlocksList {
		if block.state != "pending" {
			t.Errorf("error created block is not in pending state")
			allBlocksPending = false
		}
	}
	if allBlocksPending {
		t.Logf("created blocks are in pending state")
	}
}

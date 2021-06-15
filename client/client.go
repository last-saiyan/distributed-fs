package client

import (
	"bytes"
	"context"
	"dfs/proto"
	"dfs/utils"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	config       = utils.GetConfig()
	address      = config.NameNodeHost + config.NameNodePort
	datanodePort = config.DataNodePort
	blockSize    = config.BlockSize
)

func getGrpcClientConn(address string) (*grpc.ClientConn, *proto.DfsClient, *context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock())
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect to %v error %v", address, err)
	}
	client := proto.NewDfsClient(conn)
	return conn, &client, &cancel, err
}

// Read a file
// returns bytes of the file
// todo add offset
func Read(fileName string) []byte {
	fileLocationArr := getFileLocation(fileName, proto.FileName_READ)
	blockList := fileLocationArr.FileBlocksList

	file := make([]byte, 0)
	for _, blockReplicas := range blockList {
		replicaList := blockReplicas.BlockReplicaList
		for _, block := range replicaList {
			// add retry here with arr
			tempBlock := readBlock(block.BlockName, block.IpAddr)
			file = append(file, tempBlock...)
			break
		}
	}
	return file
}

func readBlock(chunkName string, ipAddr string) []byte {
	conn, client, cancel1, _ := getGrpcClientConn(ipAddr + datanodePort)
	defer (*cancel1)()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fileStream, err := (*client).GetBlock(ctx, &proto.FileName{FileName: chunkName})
	if err != nil {
		log.Fatalf("error getting block %v", err)
	}
	chunkData := bytes.Buffer{}
	for {
		res, err := fileStream.Recv()
		if err == io.EOF {
			return chunkData.Bytes()
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		chunkData.Write(res.GetContent())
	}
}

func getFileLocation(fileName string, mode proto.FileName_Mode) *proto.FileLocationArr {
	conn, client, _, _ := getGrpcClientConn(address)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fileLocationArr, err := (*client).GetFileLocation(ctx, &proto.FileName{FileName: fileName, Mode: mode})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return fileLocationArr
}

func writeBlock(chunkName string, ipAddr string, data []byte, blockReplicaList *proto.BlockReplicaList) error {
	conn, client, _, _ := getGrpcClientConn(ipAddr + datanodePort)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	writeBlockClient, err := (*client).WriteBlock(ctx)
	if err != nil {
		return err
	}
	sentDataLength := 0
	chunkSize := 50
	err = writeBlockClient.Send(&proto.FileWriteStream{BlockReplicaList: blockReplicaList})
	if err != nil {
		return err
	}
	for sentDataLength < len(data) {
		max := (sentDataLength + chunkSize)
		if max > len(data) {
			max = len(data)
		}
		chunk := data[sentDataLength:max]
		_ = writeBlockClient.Send(&proto.FileWriteStream{File: &proto.File{Content: chunk}})
		sentDataLength = chunkSize + sentDataLength
	}
	blockStatus, error := writeBlockClient.CloseAndRecv()
	fmt.Println(blockStatus)
	if error != nil {
		return error
	}
	return nil
}

// Write a new file
// returns true if successful
// returns false if error
func Write(fileName string, data []byte) bool {
	for len(data) > 0 {
		fileLocation := getFileLocation(fileName, proto.FileName_WRITE)
		blockReplicas := fileLocation.FileBlocksList[0]
		blockCapacity := blockSize - blockReplicas.BlockReplicaList[0].BlockSize
		limit := int64(len(data))
		if blockCapacity > int64(len(data)) {
			limit = blockCapacity
		}
		writeBlock(fileName, blockReplicas.BlockReplicaList[0].IpAddr, data[0:limit], blockReplicas)
		data = data[limit:len(data)]
	}
	return false
}

// Append to a existing file
// returns true if successful
// returns false if error
func Append(fileName string) bool {
	return false
}

// sends heartbeat to namenode to check if its alive and
// renews the file lease
func renewLock(file string) {
	conn, dfsClient, cancel1, _ := getGrpcClientConn(address + datanodePort)
	defer (*cancel1)()
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := (*dfsClient).RenewLock(ctx, &proto.FileName{FileName: file})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	if res.GetSuccess() {
		log.Printf("renewed lease")
	} else {
		log.Printf("was not able to renew lease")
	}
}

func createFileNameNode(fileName string) *proto.FileLocationArr {
	conn, dfsClient, cancel1, _ := getGrpcClientConn(address + datanodePort)
	defer conn.Close()
	defer (*cancel1)()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fileLocationArr, err := (*dfsClient).GetFileLocation(ctx, &proto.FileName{FileName: fileName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return fileLocationArr
}

// CreateFile creates a file
func CreateFile(file string) error {
	fileLocation := createFileNameNode(file)
	fileBlocks := fileLocation.FileBlocksList
	blockReplicas := fileBlocks[0]
	writeBlock(file, blockReplicas.BlockReplicaList[0].IpAddr, make([]byte, 0), blockReplicas)
	for _, replica := range blockReplicas.BlockReplicaList {
		fmt.Println(replica.IpAddr, "IpAddr")
		fmt.Println(replica.BlockName, "BlockName")
	}
	return nil
}

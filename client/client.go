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
	// address = "localhost:8001"
	// address = "192.168.99.101:30007"
	config  = utils.GetConfig()
	address = config.NameNodeHost + ":" + config.NameNodePort
)

// Read a file
// returns bytes of the file
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
	conn, err := grpc.Dial(ipAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewDfsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fileStream, err := c.GetBlock(ctx, &proto.FileName{FileName: chunkName})
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewDfsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetFileLocation(ctx, &proto.FileName{FileName: fileName, Mode: mode})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r
}

func writeBlock(chunkName string, ipAddr string, data []byte, blockReplicaList *proto.BlockReplicaList) error {
	conn, err := grpc.Dial(ipAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewDfsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// client := proto.Dfs_WriteBlockClient
	writeBlockClient, err := c.WriteBlock(ctx)
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
		err = writeBlockClient.Send(&proto.FileWriteStream{File: &proto.File{Content: chunk}})
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
	fileLocation := getFileLocation(fileName, proto.FileName_WRITE)
	fileBlocks := fileLocation.FileBlocksList
	blockReplicas := fileBlocks[0]
	for _, replica := range blockReplicas.BlockReplicaList {
		fmt.Println(replica.IpAddr, "IpAddr")
		fmt.Println(replica.BlockName, "BlockName")
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	dfsClient := proto.NewDfsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := dfsClient.RenewLock(ctx, &proto.FileName{FileName: file})
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
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewDfsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetFileLocation(ctx, &proto.FileName{FileName: fileName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r
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

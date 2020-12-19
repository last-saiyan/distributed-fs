package client

import (
	"bytes"
	"context"
	"dfs/proto"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// Read a file
// returns bytes of the file
func Read(fileName string) []byte {
	chunkName, ipAddr := getChunkName(fileName, proto.FileName_READ)
	// assigning for testing purpose
	chunkName = "temp/temp.txt"
	ipAddr = address
	return readBlock(chunkName, ipAddr)
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

func getChunkName(fileName string, mode proto.FileName_Mode) (string, string) {
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
	return r.GetChunkName(), r.GetIpAddr()
}

// Write a new file
// returns true if successful
// returns false if error
func Write(fileName string, data []byte) bool {
	fmt.Print("read file")
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

package client

import (
	"context"
	"dfs/proto"
	"fmt"
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
	log.Printf("Greeting: %s", r.GetChunkName())

	fmt.Print("read file")
	return nil
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

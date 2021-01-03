package main

import (
	"context"
	datanode "dfs/dataNode"
	"dfs/proto"
	"dfs/utils"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	config          = utils.GetConfig()
	port            = ":" + config.DataNodePort
	nameNodeHostURL = "http://" + config.NameNodeHost + ":" + config.NameNodePort
)

type server struct {
	proto.UnimplementedDfsServer
}

func (s *server) GetBlock(in *proto.FileName, stream proto.Dfs_GetBlockServer) error {
	b := datanode.Block{}
	b.InitBlock(in.FileName, "r")
	for b.HasNextChunk() {
		chunk, n, err := b.GetNextChunk()
		if err != nil {
			return err
		}
		stream.Send(&proto.File{Content: (*chunk)[:n]})
	}
	b.Close()
	return nil
}

func (s *server) WriteBlock(stream proto.Dfs_WriteBlockServer) error {
	b := datanode.Block{}
	fileWriteStream, err := stream.Recv()
	if err == io.EOF {
		blockStatus := proto.BlockStatus{Success: false}
		stream.SendAndClose(&blockStatus)
	}
	// use this for writepipeline
	// replicaList := fileWriteStream.BlockReplicaList.BlockReplicaList
	// fmt.Println(replicaList, "replicaList")
	fileName := fileWriteStream.BlockReplicaList.BlockReplicaList[0].BlockName
	b.InitBlock(fileName, "w")
	fmt.Println(fileWriteStream, "fileWriteStream")
	file := make([]byte, 0)
	for {
		fileWriteStream, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("file", string(file))
			b.Close()
			blockStatus := proto.BlockStatus{Success: true}
			stream.SendAndClose(&blockStatus)
			break
		}
		content := fileWriteStream.File.Content
		err = b.WriteChunk(content)
		if err != nil {
			blockStatus := proto.BlockStatus{Success: false}
			stream.SendAndClose(&blockStatus)
		}
		file = append(file, content...)
	}
	return nil
}

func registerDataNode() error {
	fmt.Println("hwllo world")
	conn, err := grpc.Dial(nameNodeHostURL, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	defer conn.Close()
	c := proto.NewDfsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	registerStatus, err := c.RegisterDataNode(ctx, &proto.RegisterDataNodeReq{New: true})
	if err != nil {
		log.Fatalf("did not register: %v", err)
		return err
	}
	fmt.Println(registerStatus, "registerStatus")
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	err = registerDataNode()
	if err != nil {
		log.Fatalf("failed to regester to namenode: %v", err)
	}
	proto.RegisterDfsServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	datanode "dfs/dataNode"
	"dfs/proto"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8010"
)

type server struct {
	proto.UnimplementedDfsServer
}

func (s *server) GetBlock(in *proto.FileName, stream proto.Dfs_GetBlockServer) error {
	b := datanode.Block{}
	b.InitBlock(in.FileName, "r", 500, 20)
	for b.HasNextChunk() {
		chunk, n, err := b.GetNextChunk()
		if err != nil {
			return err
		}
		stream.Send(&proto.File{Content: (*chunk)[:n]})
	}
	return nil
}

func (s *server) WriteBlock(stream proto.Dfs_WriteBlockServer) error {

	fileWriteStream, err := stream.Recv()
	if err == io.EOF {
		blockStatus := proto.BlockStatus{Success: false}
		stream.SendAndClose(&blockStatus)
	}
	// use this for writepipeline
	replicaList := fileWriteStream.BlockReplicaList.BlockReplicaList
	fmt.Println(replicaList, "replicaList")
	file := make([]byte, 0)
	for {
		fileWriteStream, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("file", file)
			blockStatus := proto.BlockStatus{Success: true}
			stream.SendAndClose(&blockStatus)
			break
		}
		content := fileWriteStream.File.Content
		file = append(file, content...)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterDfsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

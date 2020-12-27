package main

import (
	datanode "dfs/dataNode"
	"dfs/namenode"
	"dfs/proto"
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
	namenode.GetFileLocation(in.FileName)
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

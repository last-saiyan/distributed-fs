package main

import (
	"dfs/namenode"
	"dfs/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	proto.UnimplementedDfsServer
}

// func GetBlock() ()  {

// }

func (s *server) GetBlock(in *proto.FileName, stream proto.Dfs_GetBlockServer) error {
	println(in.FileName)
	namenode.GetFileLocation(in.FileName)
	return stream.Send(&proto.File{Content: make([]byte, 10)})
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

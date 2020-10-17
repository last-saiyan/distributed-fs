package main

import (
	"context"
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

func (s *server) GetFileLocation(ctx context.Context, in *proto.FileLocReq) (*proto.FileLocation, error) {
	println(in.FileName)
	return &proto.FileLocation{ChunkName: "chunkNNamee"}, nil
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

package main

import (
	"context"
	namenode "dfs/namenode"
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

func (s *server) GetFileLocation(ctx context.Context, in *proto.FileName) (*proto.FileLocation, error) {
	if proto.FileName_READ == in.Mode {
		namenode.GetFileLocation(in.FileName)
		return &proto.FileLocation{ChunkName: "chunkNNamee"}, nil
	} else if proto.FileName_WRITE == in.Mode {

	}
	return nil, nil
}

func (s *server) RenewLock(ctx context.Context, in *proto.FileName) (*proto.RenewalStatus, error) {
	println(in.FileName)
	namenode.GetFileLocation(in.FileName)
	return &proto.RenewalStatus{Success: true}, nil
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

	namenode.Init()
}

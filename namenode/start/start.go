package main

import (
	"context"
	namenode "dfs/namenode"
	"dfs/proto"
	"dfs/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	dataNodePort = utils.GetConfig().DataNodePort
	blockSize    = utils.GetConfig().BlockSize
	replicaCount = utils.GetConfig().Replica
	nn           = namenode.NameNode{}
)

type server struct {
	proto.UnimplementedDfsServer
}

func (s *server) GetFileLocation(ctx context.Context, in *proto.FileName) (*proto.FileLocationArr, error) {
	if proto.FileName_READ == in.Mode {
		return nn.GetFileLocation(in.FileName)
	} else if proto.FileName_WRITE == in.Mode {

	}
	return nil, nil
}

func (s *server) RenewLock(ctx context.Context, in *proto.FileName) (*proto.RenewalStatus, error) {
	println(in.FileName)
	nn.GetFileLocation(in.FileName)
	return &proto.RenewalStatus{Success: true}, nil
}

func main() {
	utils.GetConfig()
	lis, err := net.Listen("tcp", dataNodePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterDfsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	nn.Init(blockSize, replicaCount)
}

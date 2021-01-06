package main

import (
	"context"
	namenode "dfs/namenode"
	"dfs/proto"
	"dfs/utils"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/peer"

	"google.golang.org/grpc"
)

var (
	nameNodePort = utils.GetConfig().NameNodePort
	blockSize    = utils.GetConfig().BlockSize
	replicaCount = utils.GetConfig().Replica
	nn           = namenode.NameNode{}
)

type server struct {
	proto.UnimplementedDfsServer
}

func (s *server) GetFileLocation(ctx context.Context, in *proto.FileName) (*proto.FileLocationArr, error) {
	fmt.Println("getfile")

	if proto.FileName_READ == in.Mode {
		fmt.Println("getfile")

		return nn.GetFileLocation(in.FileName)
	} else if proto.FileName_WRITE == in.Mode {
		return nn.WriteToFile(in.FileName)
	}
	return nil, nil
}

func (s *server) RegisterDataNode(ctx context.Context, datanode *proto.RegisterDataNodeReq) (*proto.RegisterStatus, error) {
	dataNodePeer, _ := peer.FromContext(ctx)
	nn.RegisterDataNode(namenode.DatanodeMeta{IPAddr: dataNodePeer.Addr.String(), DiskUsage: datanode.DiskUsage})
	fmt.Println(dataNodePeer.Addr.String(), "ipAdddr")
	return &proto.RegisterStatus{Status: true}, nil
}

func (s *server) CreateFile(ctx context.Context, in *proto.FileName) (*proto.FileLocationArr, error) {
	return nn.CreateFile(in.FileName)
}

func (s *server) RenewLock(ctx context.Context, in *proto.FileName) (*proto.RenewalStatus, error) {
	println(in.FileName)
	nn.GetFileLocation(in.FileName)
	return &proto.RenewalStatus{Success: true}, nil
}

func main() {
	nn.Init(blockSize, replicaCount)
	utils.GetConfig()
	lis, err := net.Listen("tcp", nameNodePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterDfsServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DfsClient is the client API for Dfs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DfsClient interface {
	GetFileLocation(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileLocationArr, error)
	RenewLock(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*RenewalStatus, error)
	CheckDataNode(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	CreateFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileLocationArr, error)
	GetBlock(ctx context.Context, in *FileName, opts ...grpc.CallOption) (Dfs_GetBlockClient, error)
	WriteBlock(ctx context.Context, opts ...grpc.CallOption) (Dfs_WriteBlockClient, error)
	RegisterDataNode(ctx context.Context, in *RegisterDataNodeReq, opts ...grpc.CallOption) (*RegisterStatus, error)
}

type dfsClient struct {
	cc grpc.ClientConnInterface
}

func NewDfsClient(cc grpc.ClientConnInterface) DfsClient {
	return &dfsClient{cc}
}

func (c *dfsClient) GetFileLocation(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileLocationArr, error) {
	out := new(FileLocationArr)
	err := c.cc.Invoke(ctx, "/proto.dfs/GetFileLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfsClient) RenewLock(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*RenewalStatus, error) {
	out := new(RenewalStatus)
	err := c.cc.Invoke(ctx, "/proto.dfs/RenewLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfsClient) CheckDataNode(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/proto.dfs/CheckDataNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfsClient) CreateFile(ctx context.Context, in *FileName, opts ...grpc.CallOption) (*FileLocationArr, error) {
	out := new(FileLocationArr)
	err := c.cc.Invoke(ctx, "/proto.dfs/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dfsClient) GetBlock(ctx context.Context, in *FileName, opts ...grpc.CallOption) (Dfs_GetBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dfs_serviceDesc.Streams[0], "/proto.dfs/GetBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfsGetBlockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dfs_GetBlockClient interface {
	Recv() (*File, error)
	grpc.ClientStream
}

type dfsGetBlockClient struct {
	grpc.ClientStream
}

func (x *dfsGetBlockClient) Recv() (*File, error) {
	m := new(File)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfsClient) WriteBlock(ctx context.Context, opts ...grpc.CallOption) (Dfs_WriteBlockClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Dfs_serviceDesc.Streams[1], "/proto.dfs/WriteBlock", opts...)
	if err != nil {
		return nil, err
	}
	x := &dfsWriteBlockClient{stream}
	return x, nil
}

type Dfs_WriteBlockClient interface {
	Send(*FileWriteStream) error
	CloseAndRecv() (*BlockStatus, error)
	grpc.ClientStream
}

type dfsWriteBlockClient struct {
	grpc.ClientStream
}

func (x *dfsWriteBlockClient) Send(m *FileWriteStream) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dfsWriteBlockClient) CloseAndRecv() (*BlockStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BlockStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dfsClient) RegisterDataNode(ctx context.Context, in *RegisterDataNodeReq, opts ...grpc.CallOption) (*RegisterStatus, error) {
	out := new(RegisterStatus)
	err := c.cc.Invoke(ctx, "/proto.dfs/RegisterDataNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DfsServer is the server API for Dfs service.
// All implementations must embed UnimplementedDfsServer
// for forward compatibility
type DfsServer interface {
	GetFileLocation(context.Context, *FileName) (*FileLocationArr, error)
	RenewLock(context.Context, *FileName) (*RenewalStatus, error)
	CheckDataNode(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	CreateFile(context.Context, *FileName) (*FileLocationArr, error)
	GetBlock(*FileName, Dfs_GetBlockServer) error
	WriteBlock(Dfs_WriteBlockServer) error
	RegisterDataNode(context.Context, *RegisterDataNodeReq) (*RegisterStatus, error)
	mustEmbedUnimplementedDfsServer()
}

// UnimplementedDfsServer must be embedded to have forward compatible implementations.
type UnimplementedDfsServer struct {
}

func (UnimplementedDfsServer) GetFileLocation(context.Context, *FileName) (*FileLocationArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileLocation not implemented")
}
func (UnimplementedDfsServer) RenewLock(context.Context, *FileName) (*RenewalStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewLock not implemented")
}
func (UnimplementedDfsServer) CheckDataNode(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckDataNode not implemented")
}
func (UnimplementedDfsServer) CreateFile(context.Context, *FileName) (*FileLocationArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedDfsServer) GetBlock(*FileName, Dfs_GetBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedDfsServer) WriteBlock(Dfs_WriteBlockServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteBlock not implemented")
}
func (UnimplementedDfsServer) RegisterDataNode(context.Context, *RegisterDataNodeReq) (*RegisterStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDataNode not implemented")
}
func (UnimplementedDfsServer) mustEmbedUnimplementedDfsServer() {}

// UnsafeDfsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DfsServer will
// result in compilation errors.
type UnsafeDfsServer interface {
	mustEmbedUnimplementedDfsServer()
}

func RegisterDfsServer(s *grpc.Server, srv DfsServer) {
	s.RegisterService(&_Dfs_serviceDesc, srv)
}

func _Dfs_GetFileLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfsServer).GetFileLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dfs/GetFileLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfsServer).GetFileLocation(ctx, req.(*FileName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dfs_RenewLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfsServer).RenewLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dfs/RenewLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfsServer).RenewLock(ctx, req.(*FileName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dfs_CheckDataNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfsServer).CheckDataNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dfs/CheckDataNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfsServer).CheckDataNode(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dfs_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfsServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dfs/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfsServer).CreateFile(ctx, req.(*FileName))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dfs_GetBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DfsServer).GetBlock(m, &dfsGetBlockServer{stream})
}

type Dfs_GetBlockServer interface {
	Send(*File) error
	grpc.ServerStream
}

type dfsGetBlockServer struct {
	grpc.ServerStream
}

func (x *dfsGetBlockServer) Send(m *File) error {
	return x.ServerStream.SendMsg(m)
}

func _Dfs_WriteBlock_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DfsServer).WriteBlock(&dfsWriteBlockServer{stream})
}

type Dfs_WriteBlockServer interface {
	SendAndClose(*BlockStatus) error
	Recv() (*FileWriteStream, error)
	grpc.ServerStream
}

type dfsWriteBlockServer struct {
	grpc.ServerStream
}

func (x *dfsWriteBlockServer) SendAndClose(m *BlockStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dfsWriteBlockServer) Recv() (*FileWriteStream, error) {
	m := new(FileWriteStream)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dfs_RegisterDataNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDataNodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DfsServer).RegisterDataNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.dfs/RegisterDataNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DfsServer).RegisterDataNode(ctx, req.(*RegisterDataNodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dfs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.dfs",
	HandlerType: (*DfsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileLocation",
			Handler:    _Dfs_GetFileLocation_Handler,
		},
		{
			MethodName: "RenewLock",
			Handler:    _Dfs_RenewLock_Handler,
		},
		{
			MethodName: "CheckDataNode",
			Handler:    _Dfs_CheckDataNode_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _Dfs_CreateFile_Handler,
		},
		{
			MethodName: "RegisterDataNode",
			Handler:    _Dfs_RegisterDataNode_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlock",
			Handler:       _Dfs_GetBlock_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WriteBlock",
			Handler:       _Dfs_WriteBlock_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/dfs.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: proto/dfs.proto

// option go_package = "google.golang.org/grpc/examples/helloworld/helloworld";

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FileLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr    string `protobuf:"bytes,1,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	ChunkName string `protobuf:"bytes,2,opt,name=chunkName,proto3" json:"chunkName,omitempty"`
}

func (x *FileLocation) Reset() {
	*x = FileLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileLocation) ProtoMessage() {}

func (x *FileLocation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileLocation.ProtoReflect.Descriptor instead.
func (*FileLocation) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{0}
}

func (x *FileLocation) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *FileLocation) GetChunkName() string {
	if x != nil {
		return x.ChunkName
	}
	return ""
}

type FileLocReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
}

func (x *FileLocReq) Reset() {
	*x = FileLocReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileLocReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileLocReq) ProtoMessage() {}

func (x *FileLocReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileLocReq.ProtoReflect.Descriptor instead.
func (*FileLocReq) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{1}
}

func (x *FileLocReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

var File_proto_dfs_proto protoreflect.FileDescriptor

var file_proto_dfs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x28,
	0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x40, 0x0a, 0x03, 0x64, 0x66, 0x73, 0x12,
	0x39, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c,
	0x6f, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_dfs_proto_rawDescOnce sync.Once
	file_proto_dfs_proto_rawDescData = file_proto_dfs_proto_rawDesc
)

func file_proto_dfs_proto_rawDescGZIP() []byte {
	file_proto_dfs_proto_rawDescOnce.Do(func() {
		file_proto_dfs_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dfs_proto_rawDescData)
	})
	return file_proto_dfs_proto_rawDescData
}

var file_proto_dfs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_dfs_proto_goTypes = []interface{}{
	(*FileLocation)(nil), // 0: proto.FileLocation
	(*FileLocReq)(nil),   // 1: proto.FileLocReq
}
var file_proto_dfs_proto_depIdxs = []int32{
	1, // 0: proto.dfs.GetFileLocation:input_type -> proto.FileLocReq
	0, // 1: proto.dfs.GetFileLocation:output_type -> proto.FileLocation
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_dfs_proto_init() }
func file_proto_dfs_proto_init() {
	if File_proto_dfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_dfs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileLocReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_dfs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dfs_proto_goTypes,
		DependencyIndexes: file_proto_dfs_proto_depIdxs,
		MessageInfos:      file_proto_dfs_proto_msgTypes,
	}.Build()
	File_proto_dfs_proto = out.File
	file_proto_dfs_proto_rawDesc = nil
	file_proto_dfs_proto_goTypes = nil
	file_proto_dfs_proto_depIdxs = nil
}

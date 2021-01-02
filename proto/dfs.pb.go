// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: proto/dfs.proto

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

type FileName_Mode int32

const (
	FileName_READ  FileName_Mode = 0
	FileName_WRITE FileName_Mode = 1
)

// Enum value maps for FileName_Mode.
var (
	FileName_Mode_name = map[int32]string{
		0: "READ",
		1: "WRITE",
	}
	FileName_Mode_value = map[string]int32{
		"READ":  0,
		"WRITE": 1,
	}
)

func (x FileName_Mode) Enum() *FileName_Mode {
	p := new(FileName_Mode)
	*p = x
	return p
}

func (x FileName_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileName_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_dfs_proto_enumTypes[0].Descriptor()
}

func (FileName_Mode) Type() protoreflect.EnumType {
	return &file_proto_dfs_proto_enumTypes[0]
}

func (x FileName_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileName_Mode.Descriptor instead.
func (FileName_Mode) EnumDescriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{3, 0}
}

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3 // Used only by the Watch method.
)

// Enum value maps for HealthCheckResponse_ServingStatus.
var (
	HealthCheckResponse_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
		3: "SERVICE_UNKNOWN",
	}
	HealthCheckResponse_ServingStatus_value = map[string]int32{
		"UNKNOWN":         0,
		"SERVING":         1,
		"NOT_SERVING":     2,
		"SERVICE_UNKNOWN": 3,
	}
)

func (x HealthCheckResponse_ServingStatus) Enum() *HealthCheckResponse_ServingStatus {
	p := new(HealthCheckResponse_ServingStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_dfs_proto_enumTypes[1].Descriptor()
}

func (HealthCheckResponse_ServingStatus) Type() protoreflect.EnumType {
	return &file_proto_dfs_proto_enumTypes[1]
}

func (x HealthCheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServingStatus.Descriptor instead.
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{6, 0}
}

type BlockLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr    string `protobuf:"bytes,1,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	BlockName string `protobuf:"bytes,2,opt,name=blockName,proto3" json:"blockName,omitempty"`
	BlockSize int32  `protobuf:"varint,3,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
}

func (x *BlockLocation) Reset() {
	*x = BlockLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockLocation) ProtoMessage() {}

func (x *BlockLocation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BlockLocation.ProtoReflect.Descriptor instead.
func (*BlockLocation) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{0}
}

func (x *BlockLocation) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *BlockLocation) GetBlockName() string {
	if x != nil {
		return x.BlockName
	}
	return ""
}

func (x *BlockLocation) GetBlockSize() int32 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

type FileLocationArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileReplicasList []*BlockReplicaList `protobuf:"bytes,1,rep,name=FileReplicasList,proto3" json:"FileReplicasList,omitempty"`
}

func (x *FileLocationArr) Reset() {
	*x = FileLocationArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileLocationArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileLocationArr) ProtoMessage() {}

func (x *FileLocationArr) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FileLocationArr.ProtoReflect.Descriptor instead.
func (*FileLocationArr) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{1}
}

func (x *FileLocationArr) GetFileReplicasList() []*BlockReplicaList {
	if x != nil {
		return x.FileReplicasList
	}
	return nil
}

type BlockReplicaList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockReplicaList []*BlockLocation `protobuf:"bytes,1,rep,name=BlockReplicaList,proto3" json:"BlockReplicaList,omitempty"`
}

func (x *BlockReplicaList) Reset() {
	*x = BlockReplicaList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockReplicaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockReplicaList) ProtoMessage() {}

func (x *BlockReplicaList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockReplicaList.ProtoReflect.Descriptor instead.
func (*BlockReplicaList) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{2}
}

func (x *BlockReplicaList) GetBlockReplicaList() []*BlockLocation {
	if x != nil {
		return x.BlockReplicaList
	}
	return nil
}

type FileName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string        `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	Mode     FileName_Mode `protobuf:"varint,2,opt,name=mode,proto3,enum=proto.FileName_Mode" json:"mode,omitempty"`
}

func (x *FileName) Reset() {
	*x = FileName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileName) ProtoMessage() {}

func (x *FileName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileName.ProtoReflect.Descriptor instead.
func (*FileName) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{3}
}

func (x *FileName) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *FileName) GetMode() FileName_Mode {
	if x != nil {
		return x.Mode
	}
	return FileName_READ
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{4}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type RenewalStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RenewalStatus) Reset() {
	*x = RenewalStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewalStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewalStatus) ProtoMessage() {}

func (x *RenewalStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewalStatus.ProtoReflect.Descriptor instead.
func (*RenewalStatus) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{5}
}

func (x *RenewalStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{6}
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheckResponse_UNKNOWN
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{7}
}

func (x *File) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_proto_dfs_proto protoreflect.FileDescriptor

var file_proto_dfs_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x66, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x56, 0x0a,
	0x0f, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72,
	0x12, 0x43, 0x0a, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x10, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x10, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x08, 0x46,
	0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x1b, 0x0a,
	0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x01, 0x22, 0x2e, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f,
	0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03,
	0x22, 0x20, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x32, 0xeb, 0x01, 0x0a, 0x03, 0x64, 0x66, 0x73, 0x12, 0x3a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x12, 0x32, 0x0a, 0x09, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x4c,
	0x6f, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x0b, 0x5a, 0x09, 0x64, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_proto_dfs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_dfs_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_dfs_proto_goTypes = []interface{}{
	(FileName_Mode)(0),                     // 0: proto.FileName.Mode
	(HealthCheckResponse_ServingStatus)(0), // 1: proto.HealthCheckResponse.ServingStatus
	(*BlockLocation)(nil),                  // 2: proto.BlockLocation
	(*FileLocationArr)(nil),                // 3: proto.FileLocationArr
	(*BlockReplicaList)(nil),               // 4: proto.BlockReplicaList
	(*FileName)(nil),                       // 5: proto.FileName
	(*HealthCheckRequest)(nil),             // 6: proto.HealthCheckRequest
	(*RenewalStatus)(nil),                  // 7: proto.RenewalStatus
	(*HealthCheckResponse)(nil),            // 8: proto.HealthCheckResponse
	(*File)(nil),                           // 9: proto.File
}
var file_proto_dfs_proto_depIdxs = []int32{
	4, // 0: proto.FileLocationArr.FileReplicasList:type_name -> proto.BlockReplicaList
	2, // 1: proto.BlockReplicaList.BlockReplicaList:type_name -> proto.BlockLocation
	0, // 2: proto.FileName.mode:type_name -> proto.FileName.Mode
	1, // 3: proto.HealthCheckResponse.status:type_name -> proto.HealthCheckResponse.ServingStatus
	5, // 4: proto.dfs.GetFileLocation:input_type -> proto.FileName
	5, // 5: proto.dfs.RenewLock:input_type -> proto.FileName
	6, // 6: proto.dfs.CheckDataNode:input_type -> proto.HealthCheckRequest
	5, // 7: proto.dfs.GetBlock:input_type -> proto.FileName
	3, // 8: proto.dfs.GetFileLocation:output_type -> proto.FileLocationArr
	7, // 9: proto.dfs.RenewLock:output_type -> proto.RenewalStatus
	8, // 10: proto.dfs.CheckDataNode:output_type -> proto.HealthCheckResponse
	9, // 11: proto.dfs.GetBlock:output_type -> proto.File
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_dfs_proto_init() }
func file_proto_dfs_proto_init() {
	if File_proto_dfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockLocation); i {
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
			switch v := v.(*FileLocationArr); i {
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
		file_proto_dfs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockReplicaList); i {
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
		file_proto_dfs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileName); i {
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
		file_proto_dfs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_proto_dfs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewalStatus); i {
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
		file_proto_dfs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_proto_dfs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dfs_proto_goTypes,
		DependencyIndexes: file_proto_dfs_proto_depIdxs,
		EnumInfos:         file_proto_dfs_proto_enumTypes,
		MessageInfos:      file_proto_dfs_proto_msgTypes,
	}.Build()
	File_proto_dfs_proto = out.File
	file_proto_dfs_proto_rawDesc = nil
	file_proto_dfs_proto_goTypes = nil
	file_proto_dfs_proto_depIdxs = nil
}

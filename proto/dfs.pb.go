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

type DatanodeOperation_Operation int32

const (
	DatanodeOperation_DELETE    DatanodeOperation_Operation = 0
	DatanodeOperation_REPLICATE DatanodeOperation_Operation = 1
)

// Enum value maps for DatanodeOperation_Operation.
var (
	DatanodeOperation_Operation_name = map[int32]string{
		0: "DELETE",
		1: "REPLICATE",
	}
	DatanodeOperation_Operation_value = map[string]int32{
		"DELETE":    0,
		"REPLICATE": 1,
	}
)

func (x DatanodeOperation_Operation) Enum() *DatanodeOperation_Operation {
	p := new(DatanodeOperation_Operation)
	*p = x
	return p
}

func (x DatanodeOperation_Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatanodeOperation_Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_dfs_proto_enumTypes[0].Descriptor()
}

func (DatanodeOperation_Operation) Type() protoreflect.EnumType {
	return &file_proto_dfs_proto_enumTypes[0]
}

func (x DatanodeOperation_Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatanodeOperation_Operation.Descriptor instead.
func (DatanodeOperation_Operation) EnumDescriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{1, 0}
}

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
	return file_proto_dfs_proto_enumTypes[1].Descriptor()
}

func (FileName_Mode) Type() protoreflect.EnumType {
	return &file_proto_dfs_proto_enumTypes[1]
}

func (x FileName_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileName_Mode.Descriptor instead.
func (FileName_Mode) EnumDescriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{8, 0}
}

type Heartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiskUsage uint64 `protobuf:"varint,2,opt,name=DiskUsage,proto3" json:"DiskUsage,omitempty"`
}

func (x *Heartbeat) Reset() {
	*x = Heartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heartbeat) ProtoMessage() {}

func (x *Heartbeat) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Heartbeat.ProtoReflect.Descriptor instead.
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{0}
}

func (x *Heartbeat) GetDiskUsage() uint64 {
	if x != nil {
		return x.DiskUsage
	}
	return 0
}

type DatanodeOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation DatanodeOperation_Operation `protobuf:"varint,1,opt,name=operation,proto3,enum=proto.DatanodeOperation_Operation" json:"operation,omitempty"`
	IpAddr    string                      `protobuf:"bytes,2,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
}

func (x *DatanodeOperation) Reset() {
	*x = DatanodeOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatanodeOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatanodeOperation) ProtoMessage() {}

func (x *DatanodeOperation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DatanodeOperation.ProtoReflect.Descriptor instead.
func (*DatanodeOperation) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{1}
}

func (x *DatanodeOperation) GetOperation() DatanodeOperation_Operation {
	if x != nil {
		return x.Operation
	}
	return DatanodeOperation_DELETE
}

func (x *DatanodeOperation) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

type RegisterDataNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	New       bool   `protobuf:"varint,1,opt,name=new,proto3" json:"new,omitempty"`
	DiskUsage uint64 `protobuf:"varint,2,opt,name=DiskUsage,proto3" json:"DiskUsage,omitempty"`
}

func (x *RegisterDataNodeReq) Reset() {
	*x = RegisterDataNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterDataNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterDataNodeReq) ProtoMessage() {}

func (x *RegisterDataNodeReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterDataNodeReq.ProtoReflect.Descriptor instead.
func (*RegisterDataNodeReq) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterDataNodeReq) GetNew() bool {
	if x != nil {
		return x.New
	}
	return false
}

func (x *RegisterDataNodeReq) GetDiskUsage() uint64 {
	if x != nil {
		return x.DiskUsage
	}
	return 0
}

type RegisterStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RegisterStatus) Reset() {
	*x = RegisterStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterStatus) ProtoMessage() {}

func (x *RegisterStatus) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RegisterStatus.ProtoReflect.Descriptor instead.
func (*RegisterStatus) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterStatus) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type FileWriteStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File             *File             `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	BlockReplicaList *BlockReplicaList `protobuf:"bytes,2,opt,name=BlockReplicaList,proto3" json:"BlockReplicaList,omitempty"`
}

func (x *FileWriteStream) Reset() {
	*x = FileWriteStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileWriteStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileWriteStream) ProtoMessage() {}

func (x *FileWriteStream) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FileWriteStream.ProtoReflect.Descriptor instead.
func (*FileWriteStream) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{4}
}

func (x *FileWriteStream) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FileWriteStream) GetBlockReplicaList() *BlockReplicaList {
	if x != nil {
		return x.BlockReplicaList
	}
	return nil
}

type BlockLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr    string `protobuf:"bytes,1,opt,name=ipAddr,proto3" json:"ipAddr,omitempty"`
	BlockName string `protobuf:"bytes,2,opt,name=blockName,proto3" json:"blockName,omitempty"`
	BlockSize int64  `protobuf:"varint,3,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
}

func (x *BlockLocation) Reset() {
	*x = BlockLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockLocation) ProtoMessage() {}

func (x *BlockLocation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BlockLocation.ProtoReflect.Descriptor instead.
func (*BlockLocation) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{5}
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

func (x *BlockLocation) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

type FileLocationArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileBlocksList []*BlockReplicaList `protobuf:"bytes,1,rep,name=FileBlocksList,proto3" json:"FileBlocksList,omitempty"`
}

func (x *FileLocationArr) Reset() {
	*x = FileLocationArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileLocationArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileLocationArr) ProtoMessage() {}

func (x *FileLocationArr) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FileLocationArr.ProtoReflect.Descriptor instead.
func (*FileLocationArr) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{6}
}

func (x *FileLocationArr) GetFileBlocksList() []*BlockReplicaList {
	if x != nil {
		return x.FileBlocksList
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
		mi := &file_proto_dfs_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockReplicaList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockReplicaList) ProtoMessage() {}

func (x *BlockReplicaList) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use BlockReplicaList.ProtoReflect.Descriptor instead.
func (*BlockReplicaList) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{7}
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
		mi := &file_proto_dfs_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileName) ProtoMessage() {}

func (x *FileName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[8]
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
	return file_proto_dfs_proto_rawDescGZIP(), []int{8}
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

type RenewalStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RenewalStatus) Reset() {
	*x = RenewalStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewalStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewalStatus) ProtoMessage() {}

func (x *RenewalStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[9]
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
	return file_proto_dfs_proto_rawDescGZIP(), []int{9}
}

func (x *RenewalStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BlockStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BlockStatus) Reset() {
	*x = BlockStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dfs_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockStatus) ProtoMessage() {}

func (x *BlockStatus) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockStatus.ProtoReflect.Descriptor instead.
func (*BlockStatus) Descriptor() ([]byte, []int) {
	return file_proto_dfs_proto_rawDescGZIP(), []int{10}
}

func (x *BlockStatus) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
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
		mi := &file_proto_dfs_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dfs_proto_msgTypes[11]
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
	return file_proto_dfs_proto_rawDescGZIP(), []int{11}
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
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x09, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x22, 0x26, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x52, 0x45, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x01, 0x22, 0x45, 0x0a, 0x13, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x6e, 0x65, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x77, 0x0a, 0x0f,
	0x46, 0x69, 0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12,
	0x1f, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x43, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x52, 0x0a, 0x0f, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x12, 0x3f, 0x0a,
	0x0e, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0e,
	0x46, 0x69, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x54,
	0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x40, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x04,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x1b, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x52, 0x49, 0x54,
	0x45, 0x10, 0x01, 0x22, 0x29, 0x0a, 0x0d, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x27,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x20, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xa0, 0x03, 0x0a, 0x03, 0x64, 0x66,
	0x73, 0x12, 0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x12, 0x32, 0x0a,
	0x09, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3f, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x6e, 0x6f, 0x64, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x72, 0x12, 0x2c, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x45, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0x5a, 0x09,
	0x64, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
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
var file_proto_dfs_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_dfs_proto_goTypes = []interface{}{
	(DatanodeOperation_Operation)(0), // 0: proto.DatanodeOperation.Operation
	(FileName_Mode)(0),               // 1: proto.FileName.Mode
	(*Heartbeat)(nil),                // 2: proto.Heartbeat
	(*DatanodeOperation)(nil),        // 3: proto.DatanodeOperation
	(*RegisterDataNodeReq)(nil),      // 4: proto.RegisterDataNodeReq
	(*RegisterStatus)(nil),           // 5: proto.RegisterStatus
	(*FileWriteStream)(nil),          // 6: proto.FileWriteStream
	(*BlockLocation)(nil),            // 7: proto.BlockLocation
	(*FileLocationArr)(nil),          // 8: proto.FileLocationArr
	(*BlockReplicaList)(nil),         // 9: proto.BlockReplicaList
	(*FileName)(nil),                 // 10: proto.FileName
	(*RenewalStatus)(nil),            // 11: proto.RenewalStatus
	(*BlockStatus)(nil),              // 12: proto.BlockStatus
	(*File)(nil),                     // 13: proto.File
}
var file_proto_dfs_proto_depIdxs = []int32{
	0,  // 0: proto.DatanodeOperation.operation:type_name -> proto.DatanodeOperation.Operation
	13, // 1: proto.FileWriteStream.file:type_name -> proto.File
	9,  // 2: proto.FileWriteStream.BlockReplicaList:type_name -> proto.BlockReplicaList
	9,  // 3: proto.FileLocationArr.FileBlocksList:type_name -> proto.BlockReplicaList
	7,  // 4: proto.BlockReplicaList.BlockReplicaList:type_name -> proto.BlockLocation
	1,  // 5: proto.FileName.mode:type_name -> proto.FileName.Mode
	10, // 6: proto.dfs.GetFileLocation:input_type -> proto.FileName
	10, // 7: proto.dfs.RenewLock:input_type -> proto.FileName
	2,  // 8: proto.dfs.DatanodeHeartbeat:input_type -> proto.Heartbeat
	10, // 9: proto.dfs.CreateFile:input_type -> proto.FileName
	10, // 10: proto.dfs.GetBlock:input_type -> proto.FileName
	6,  // 11: proto.dfs.WriteBlock:input_type -> proto.FileWriteStream
	4,  // 12: proto.dfs.RegisterDataNode:input_type -> proto.RegisterDataNodeReq
	8,  // 13: proto.dfs.GetFileLocation:output_type -> proto.FileLocationArr
	11, // 14: proto.dfs.RenewLock:output_type -> proto.RenewalStatus
	3,  // 15: proto.dfs.DatanodeHeartbeat:output_type -> proto.DatanodeOperation
	8,  // 16: proto.dfs.CreateFile:output_type -> proto.FileLocationArr
	13, // 17: proto.dfs.GetBlock:output_type -> proto.File
	12, // 18: proto.dfs.WriteBlock:output_type -> proto.BlockStatus
	5,  // 19: proto.dfs.RegisterDataNode:output_type -> proto.RegisterStatus
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_dfs_proto_init() }
func file_proto_dfs_proto_init() {
	if File_proto_dfs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dfs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heartbeat); i {
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
			switch v := v.(*DatanodeOperation); i {
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
			switch v := v.(*RegisterDataNodeReq); i {
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
			switch v := v.(*RegisterStatus); i {
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
			switch v := v.(*FileWriteStream); i {
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
		file_proto_dfs_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_dfs_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_dfs_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_dfs_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_dfs_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockStatus); i {
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
		file_proto_dfs_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   12,
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: transport.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Range struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileOffset     int64 `protobuf:"varint,1,opt,name=file_offset,json=fileOffset,proto3" json:"file_offset,omitempty"`
	OriginalOffset int64 `protobuf:"varint,2,opt,name=original_offset,json=originalOffset,proto3" json:"original_offset,omitempty"`
	FileLength     int64 `protobuf:"varint,3,opt,name=file_length,json=fileLength,proto3" json:"file_length,omitempty"`
	Length         int64 `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *Range) Reset() {
	*x = Range{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Range) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Range) ProtoMessage() {}

func (x *Range) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Range.ProtoReflect.Descriptor instead.
func (*Range) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0}
}

func (x *Range) GetFileOffset() int64 {
	if x != nil {
		return x.FileOffset
	}
	return 0
}

func (x *Range) GetOriginalOffset() int64 {
	if x != nil {
		return x.OriginalOffset
	}
	return 0
}

func (x *Range) GetFileLength() int64 {
	if x != nil {
		return x.FileLength
	}
	return 0
}

func (x *Range) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type Index struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges []*Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *Index) Reset() {
	*x = Index{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Index) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Index) ProtoMessage() {}

func (x *Index) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Index.ProtoReflect.Descriptor instead.
func (*Index) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1}
}

func (x *Index) GetRanges() []*Range {
	if x != nil {
		return x.Ranges
	}
	return nil
}

// A message to encode a filesystem path (maybe for raw access)
// Next field: 15
type PathSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Accessor string `protobuf:"bytes,3,opt,name=accessor,proto3" json:"accessor,omitempty"`
}

func (x *PathSpec) Reset() {
	*x = PathSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathSpec) ProtoMessage() {}

func (x *PathSpec) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathSpec.ProtoReflect.Descriptor instead.
func (*PathSpec) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2}
}

func (x *PathSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PathSpec) GetAccessor() string {
	if x != nil {
		return x.Accessor
	}
	return ""
}

// The Velociraptor client sends back the buffer and the filename and
// the server saves the entire file directly in the file storage
// filesystem. This allows easy recovery as well as data expiration
// policies (since the filestore is just a directory on disk with
// regular files and timestamps).
type FileBuffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pathspec *PathSpec `protobuf:"bytes,1,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	Offset   uint64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Expected size of file.
	Size       uint64 `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	StoredSize uint64 `protobuf:"varint,8,opt,name=stored_size,json=storedSize,proto3" json:"stored_size,omitempty"`
	IsSparse   bool   `protobuf:"varint,9,opt,name=is_sparse,json=isSparse,proto3" json:"is_sparse,omitempty"`
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	FlowId     string `protobuf:"bytes,4,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Eof        bool   `protobuf:"varint,5,opt,name=eof,proto3" json:"eof,omitempty"`
	// Set when the file is sparse.
	Index *Index `protobuf:"bytes,6,opt,name=index,proto3" json:"index,omitempty"`
	Mtime int64  `protobuf:"varint,10,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Atime int64  `protobuf:"varint,11,opt,name=atime,proto3" json:"atime,omitempty"`
	Ctime int64  `protobuf:"varint,12,opt,name=ctime,proto3" json:"ctime,omitempty"`
	Btime int64  `protobuf:"varint,13,opt,name=btime,proto3" json:"btime,omitempty"`
}

func (x *FileBuffer) Reset() {
	*x = FileBuffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileBuffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileBuffer) ProtoMessage() {}

func (x *FileBuffer) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileBuffer.ProtoReflect.Descriptor instead.
func (*FileBuffer) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{3}
}

func (x *FileBuffer) GetPathspec() *PathSpec {
	if x != nil {
		return x.Pathspec
	}
	return nil
}

func (x *FileBuffer) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FileBuffer) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileBuffer) GetStoredSize() uint64 {
	if x != nil {
		return x.StoredSize
	}
	return 0
}

func (x *FileBuffer) GetIsSparse() bool {
	if x != nil {
		return x.IsSparse
	}
	return false
}

func (x *FileBuffer) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *FileBuffer) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

func (x *FileBuffer) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

func (x *FileBuffer) GetIndex() *Index {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *FileBuffer) GetMtime() int64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

func (x *FileBuffer) GetAtime() int64 {
	if x != nil {
		return x.Atime
	}
	return 0
}

func (x *FileBuffer) GetCtime() int64 {
	if x != nil {
		return x.Ctime
	}
	return 0
}

func (x *FileBuffer) GetBtime() int64 {
	if x != nil {
		return x.Btime
	}
	return 0
}

type ForemanCheckin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastHuntTimestamp     uint64 `protobuf:"varint,1,opt,name=last_hunt_timestamp,json=lastHuntTimestamp,proto3" json:"last_hunt_timestamp,omitempty"`
	LastEventTableVersion uint64 `protobuf:"varint,2,opt,name=last_event_table_version,json=lastEventTableVersion,proto3" json:"last_event_table_version,omitempty"`
}

func (x *ForemanCheckin) Reset() {
	*x = ForemanCheckin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForemanCheckin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForemanCheckin) ProtoMessage() {}

func (x *ForemanCheckin) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForemanCheckin.ProtoReflect.Descriptor instead.
func (*ForemanCheckin) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{4}
}

func (x *ForemanCheckin) GetLastHuntTimestamp() uint64 {
	if x != nil {
		return x.LastHuntTimestamp
	}
	return 0
}

func (x *ForemanCheckin) GetLastEventTableVersion() uint64 {
	if x != nil {
		return x.LastEventTableVersion
	}
	return 0
}

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a,
	0x01, 0x0a, 0x05, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66,
	0x69, 0x6c, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x2d, 0x0a, 0x05, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0xdb, 0x01, 0x0a, 0x08, 0x50,
	0x61, 0x74, 0x68, 0x53, 0x70, 0x65, 0x63, 0x12, 0x81, 0x01, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x6d, 0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x67, 0x12, 0x65,
	0x54, 0x68, 0x65, 0x20, 0x70, 0x61, 0x74, 0x68, 0x20, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x20,
	0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x2e, 0x20, 0x54, 0x68, 0x69, 0x73,
	0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x74, 0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x69, 0x74, 0x73, 0x20, 0x6f, 0x77, 0x6e,
	0x20, 0x77, 0x61, 0x79, 0x2e, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x4b, 0x0a, 0x08, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe2,
	0xfc, 0xe3, 0xc4, 0x01, 0x29, 0x12, 0x27, 0x54, 0x68, 0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x20, 0x75, 0x73, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x65, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x08,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x22, 0x89, 0x03, 0x0a, 0x0a, 0x46, 0x69, 0x6c,
	0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x68, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x53, 0x70, 0x65, 0x63, 0x52, 0x08, 0x70, 0x61, 0x74, 0x68,
	0x73, 0x70, 0x65, 0x63, 0x12, 0x41, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x29, 0xe2, 0xfc, 0xe3, 0xc4, 0x01, 0x23, 0x12, 0x21, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x73, 0x70, 0x61, 0x72, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a,
	0x07, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x66, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x61, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x79, 0x0a, 0x0e, 0x46, 0x6f, 0x72, 0x65, 0x6d, 0x61, 0x6e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x68,
	0x75, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x48, 0x75, 0x6e, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x37, 0x0a, 0x18, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x35, 0x5a, 0x33, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_rawDescOnce sync.Once
	file_transport_proto_rawDescData = file_transport_proto_rawDesc
)

func file_transport_proto_rawDescGZIP() []byte {
	file_transport_proto_rawDescOnce.Do(func() {
		file_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_rawDescData)
	})
	return file_transport_proto_rawDescData
}

var file_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_transport_proto_goTypes = []interface{}{
	(*Range)(nil),          // 0: proto.Range
	(*Index)(nil),          // 1: proto.Index
	(*PathSpec)(nil),       // 2: proto.PathSpec
	(*FileBuffer)(nil),     // 3: proto.FileBuffer
	(*ForemanCheckin)(nil), // 4: proto.ForemanCheckin
}
var file_transport_proto_depIdxs = []int32{
	0, // 0: proto.Index.ranges:type_name -> proto.Range
	2, // 1: proto.FileBuffer.pathspec:type_name -> proto.PathSpec
	1, // 2: proto.FileBuffer.index:type_name -> proto.Index
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Range); i {
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
		file_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Index); i {
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
		file_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathSpec); i {
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
		file_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileBuffer); i {
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
		file_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForemanCheckin); i {
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
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
		MessageInfos:      file_transport_proto_msgTypes,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}

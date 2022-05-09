// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: objects.proto

package proto

import (
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

type ServerState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Currently running server state
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ServerState) Reset() {
	*x = ServerState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_objects_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerState) ProtoMessage() {}

func (x *ServerState) ProtoReflect() protoreflect.Message {
	mi := &file_objects_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerState.ProtoReflect.Descriptor instead.
func (*ServerState) Descriptor() ([]byte, []int) {
	return file_objects_proto_rawDescGZIP(), []int{0}
}

func (x *ServerState) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_objects_proto protoreflect.FileDescriptor

var file_objects_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42,
	0x31, 0x5a, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f,
	0x63, 0x69, 0x72, 0x61, 0x70, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_objects_proto_rawDescOnce sync.Once
	file_objects_proto_rawDescData = file_objects_proto_rawDesc
)

func file_objects_proto_rawDescGZIP() []byte {
	file_objects_proto_rawDescOnce.Do(func() {
		file_objects_proto_rawDescData = protoimpl.X.CompressGZIP(file_objects_proto_rawDescData)
	})
	return file_objects_proto_rawDescData
}

var file_objects_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_objects_proto_goTypes = []interface{}{
	(*ServerState)(nil), // 0: proto.ServerState
}
var file_objects_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_objects_proto_init() }
func file_objects_proto_init() {
	if File_objects_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_objects_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerState); i {
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
			RawDescriptor: file_objects_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_objects_proto_goTypes,
		DependencyIndexes: file_objects_proto_depIdxs,
		MessageInfos:      file_objects_proto_msgTypes,
	}.Build()
	File_objects_proto = out.File
	file_objects_proto_rawDesc = nil
	file_objects_proto_goTypes = nil
	file_objects_proto_depIdxs = nil
}

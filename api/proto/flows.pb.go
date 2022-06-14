// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: flows.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	proto1 "www.velocidex.com/golang/velociraptor/crypto/proto"
	proto "www.velocidex.com/golang/velociraptor/flows/proto"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AvailableDownloadFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Path     string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	Complete bool   `protobuf:"varint,2,opt,name=complete,proto3" json:"complete,omitempty"`
	Size     uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Date     string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *AvailableDownloadFile) Reset() {
	*x = AvailableDownloadFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableDownloadFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableDownloadFile) ProtoMessage() {}

func (x *AvailableDownloadFile) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableDownloadFile.ProtoReflect.Descriptor instead.
func (*AvailableDownloadFile) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{0}
}

func (x *AvailableDownloadFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AvailableDownloadFile) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AvailableDownloadFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AvailableDownloadFile) GetComplete() bool {
	if x != nil {
		return x.Complete
	}
	return false
}

func (x *AvailableDownloadFile) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *AvailableDownloadFile) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type AvailableDownloads struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*AvailableDownloadFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *AvailableDownloads) Reset() {
	*x = AvailableDownloads{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableDownloads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableDownloads) ProtoMessage() {}

func (x *AvailableDownloads) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableDownloads.ProtoReflect.Descriptor instead.
func (*AvailableDownloads) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{1}
}

func (x *AvailableDownloads) GetFiles() []*AvailableDownloadFile {
	if x != nil {
		return x.Files
	}
	return nil
}

type FlowDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Context            *proto.ArtifactCollectorContext `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	AvailableDownloads *AvailableDownloads             `protobuf:"bytes,16,opt,name=available_downloads,json=availableDownloads,proto3" json:"available_downloads,omitempty"`
}

func (x *FlowDetails) Reset() {
	*x = FlowDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowDetails) ProtoMessage() {}

func (x *FlowDetails) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowDetails.ProtoReflect.Descriptor instead.
func (*FlowDetails) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{2}
}

func (x *FlowDetails) GetContext() *proto.ArtifactCollectorContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *FlowDetails) GetAvailableDownloads() *AvailableDownloads {
	if x != nil {
		return x.AvailableDownloads
	}
	return nil
}

// This shows the requests that were actually sent to the client. When
// the user selects artifacts to send they are compiled into raw VQL
// for sending to the client. NOTE: Clients do not know anything about
// artifacts - they only interprect raw VQL as compiled by the server.
type ApiFlowRequestDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*proto1.VeloMessage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ApiFlowRequestDetails) Reset() {
	*x = ApiFlowRequestDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiFlowRequestDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiFlowRequestDetails) ProtoMessage() {}

func (x *ApiFlowRequestDetails) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiFlowRequestDetails.ProtoReflect.Descriptor instead.
func (*ApiFlowRequestDetails) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{3}
}

func (x *ApiFlowRequestDetails) GetItems() []*proto1.VeloMessage {
	if x != nil {
		return x.Items
	}
	return nil
}

type ApiFlowResultDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*proto1.VeloMessage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ApiFlowResultDetails) Reset() {
	*x = ApiFlowResultDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiFlowResultDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiFlowResultDetails) ProtoMessage() {}

func (x *ApiFlowResultDetails) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiFlowResultDetails.ProtoReflect.Descriptor instead.
func (*ApiFlowResultDetails) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{4}
}

func (x *ApiFlowResultDetails) GetItems() []*proto1.VeloMessage {
	if x != nil {
		return x.Items
	}
	return nil
}

type ApiFlowLogDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*proto1.LogMessage `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ApiFlowLogDetails) Reset() {
	*x = ApiFlowLogDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiFlowLogDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiFlowLogDetails) ProtoMessage() {}

func (x *ApiFlowLogDetails) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiFlowLogDetails.ProtoReflect.Descriptor instead.
func (*ApiFlowLogDetails) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{5}
}

func (x *ApiFlowLogDetails) GetItems() []*proto1.LogMessage {
	if x != nil {
		return x.Items
	}
	return nil
}

type ApiFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId        string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	FlowId          string `protobuf:"bytes,2,opt,name=flow_id,json=flowId,proto3" json:"flow_id,omitempty"`
	Offset          uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Count           uint64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	IncludeArchived bool   `protobuf:"varint,5,opt,name=include_archived,json=includeArchived,proto3" json:"include_archived,omitempty"`
	// If specified we only return flows that collected this artifact.
	Artifact string `protobuf:"bytes,6,opt,name=artifact,proto3" json:"artifact,omitempty"`
}

func (x *ApiFlowRequest) Reset() {
	*x = ApiFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiFlowRequest) ProtoMessage() {}

func (x *ApiFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiFlowRequest.ProtoReflect.Descriptor instead.
func (*ApiFlowRequest) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{6}
}

func (x *ApiFlowRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ApiFlowRequest) GetFlowId() string {
	if x != nil {
		return x.FlowId
	}
	return ""
}

func (x *ApiFlowRequest) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ApiFlowRequest) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ApiFlowRequest) GetIncludeArchived() bool {
	if x != nil {
		return x.IncludeArchived
	}
	return false
}

func (x *ApiFlowRequest) GetArtifact() string {
	if x != nil {
		return x.Artifact
	}
	return ""
}

type ApiFlowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*proto.ArtifactCollectorContext `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ApiFlowResponse) Reset() {
	*x = ApiFlowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flows_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiFlowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiFlowResponse) ProtoMessage() {}

func (x *ApiFlowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flows_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiFlowResponse.ProtoReflect.Descriptor instead.
func (*ApiFlowResponse) Descriptor() ([]byte, []int) {
	return file_flows_proto_rawDescGZIP(), []int{7}
}

func (x *ApiFlowResponse) GetItems() []*proto.ArtifactCollectorContext {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_flows_proto protoreflect.FileDescriptor

var file_flows_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a, 0x15, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x48, 0x0a,
	0x12, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x94, 0x01, 0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x4a, 0x0a, 0x13, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x22, 0x41,
	0x0a, 0x15, 0x41, 0x70, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56,
	0x65, 0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x40, 0x0a, 0x14, 0x41, 0x70, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x65, 0x6c, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x3c, 0x0a, 0x11, 0x41, 0x70, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x4c, 0x6f,
	0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x22,
	0x48, 0x0a, 0x0f, 0x41, 0x70, 0x69, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x31, 0x5a, 0x2f, 0x77, 0x77, 0x77,
	0x2e, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x72, 0x61, 0x70, 0x74,
	0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flows_proto_rawDescOnce sync.Once
	file_flows_proto_rawDescData = file_flows_proto_rawDesc
)

func file_flows_proto_rawDescGZIP() []byte {
	file_flows_proto_rawDescOnce.Do(func() {
		file_flows_proto_rawDescData = protoimpl.X.CompressGZIP(file_flows_proto_rawDescData)
	})
	return file_flows_proto_rawDescData
}

var file_flows_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_flows_proto_goTypes = []interface{}{
	(*AvailableDownloadFile)(nil),          // 0: proto.AvailableDownloadFile
	(*AvailableDownloads)(nil),             // 1: proto.AvailableDownloads
	(*FlowDetails)(nil),                    // 2: proto.FlowDetails
	(*ApiFlowRequestDetails)(nil),          // 3: proto.ApiFlowRequestDetails
	(*ApiFlowResultDetails)(nil),           // 4: proto.ApiFlowResultDetails
	(*ApiFlowLogDetails)(nil),              // 5: proto.ApiFlowLogDetails
	(*ApiFlowRequest)(nil),                 // 6: proto.ApiFlowRequest
	(*ApiFlowResponse)(nil),                // 7: proto.ApiFlowResponse
	(*proto.ArtifactCollectorContext)(nil), // 8: proto.ArtifactCollectorContext
	(*proto1.VeloMessage)(nil),             // 9: proto.VeloMessage
	(*proto1.LogMessage)(nil),              // 10: proto.LogMessage
}
var file_flows_proto_depIdxs = []int32{
	0,  // 0: proto.AvailableDownloads.files:type_name -> proto.AvailableDownloadFile
	8,  // 1: proto.FlowDetails.context:type_name -> proto.ArtifactCollectorContext
	1,  // 2: proto.FlowDetails.available_downloads:type_name -> proto.AvailableDownloads
	9,  // 3: proto.ApiFlowRequestDetails.items:type_name -> proto.VeloMessage
	9,  // 4: proto.ApiFlowResultDetails.items:type_name -> proto.VeloMessage
	10, // 5: proto.ApiFlowLogDetails.items:type_name -> proto.LogMessage
	8,  // 6: proto.ApiFlowResponse.items:type_name -> proto.ArtifactCollectorContext
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_flows_proto_init() }
func file_flows_proto_init() {
	if File_flows_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flows_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableDownloadFile); i {
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
		file_flows_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableDownloads); i {
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
		file_flows_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowDetails); i {
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
		file_flows_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiFlowRequestDetails); i {
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
		file_flows_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiFlowResultDetails); i {
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
		file_flows_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiFlowLogDetails); i {
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
		file_flows_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiFlowRequest); i {
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
		file_flows_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiFlowResponse); i {
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
			RawDescriptor: file_flows_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flows_proto_goTypes,
		DependencyIndexes: file_flows_proto_depIdxs,
		MessageInfos:      file_flows_proto_msgTypes,
	}.Build()
	File_flows_proto = out.File
	file_flows_proto_rawDesc = nil
	file_flows_proto_goTypes = nil
	file_flows_proto_depIdxs = nil
}

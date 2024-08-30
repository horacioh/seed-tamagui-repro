// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: p2p/v1alpha/syncing.proto

package p2p

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

type SetReconciliationRange_Mode int32

const (
	SetReconciliationRange_SKIP        SetReconciliationRange_Mode = 0
	SetReconciliationRange_FINGERPRINT SetReconciliationRange_Mode = 1
	SetReconciliationRange_LIST        SetReconciliationRange_Mode = 2
)

// Enum value maps for SetReconciliationRange_Mode.
var (
	SetReconciliationRange_Mode_name = map[int32]string{
		0: "SKIP",
		1: "FINGERPRINT",
		2: "LIST",
	}
	SetReconciliationRange_Mode_value = map[string]int32{
		"SKIP":        0,
		"FINGERPRINT": 1,
		"LIST":        2,
	}
)

func (x SetReconciliationRange_Mode) Enum() *SetReconciliationRange_Mode {
	p := new(SetReconciliationRange_Mode)
	*p = x
	return p
}

func (x SetReconciliationRange_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetReconciliationRange_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_p2p_v1alpha_syncing_proto_enumTypes[0].Descriptor()
}

func (SetReconciliationRange_Mode) Type() protoreflect.EnumType {
	return &file_p2p_v1alpha_syncing_proto_enumTypes[0]
}

func (x SetReconciliationRange_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetReconciliationRange_Mode.Descriptor instead.
func (SetReconciliationRange_Mode) EnumDescriptor() ([]byte, []int) {
	return file_p2p_v1alpha_syncing_proto_rawDescGZIP(), []int{3, 0}
}

type ReconcileBlobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Filters to narrow down the blobs to reconcile.
	// If not set, all public blobs are reconciled.
	Filters []*Filter `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
	// Optional. The ranges for the sender's part of the set.
	Ranges []*SetReconciliationRange `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *ReconcileBlobsRequest) Reset() {
	*x = ReconcileBlobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_v1alpha_syncing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcileBlobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcileBlobsRequest) ProtoMessage() {}

func (x *ReconcileBlobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_v1alpha_syncing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcileBlobsRequest.ProtoReflect.Descriptor instead.
func (*ReconcileBlobsRequest) Descriptor() ([]byte, []int) {
	return file_p2p_v1alpha_syncing_proto_rawDescGZIP(), []int{0}
}

func (x *ReconcileBlobsRequest) GetFilters() []*Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ReconcileBlobsRequest) GetRanges() []*SetReconciliationRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

type ReconcileBlobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ranges []*SetReconciliationRange `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
}

func (x *ReconcileBlobsResponse) Reset() {
	*x = ReconcileBlobsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_v1alpha_syncing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcileBlobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcileBlobsResponse) ProtoMessage() {}

func (x *ReconcileBlobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_v1alpha_syncing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcileBlobsResponse.ProtoReflect.Descriptor instead.
func (*ReconcileBlobsResponse) Descriptor() ([]byte, []int) {
	return file_p2p_v1alpha_syncing_proto_rawDescGZIP(), []int{1}
}

func (x *ReconcileBlobsResponse) GetRanges() []*SetReconciliationRange {
	if x != nil {
		return x.Ranges
	}
	return nil
}

// Filter describes which blobs to select for reconciliation.
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects only blobs related to the given resource.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// If its recursive, then all the documents below the path are
	// will also pass the filter.
	Recursive bool `protobuf:"varint,2,opt,name=recursive,proto3" json:"recursive,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_v1alpha_syncing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_v1alpha_syncing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_p2p_v1alpha_syncing_proto_rawDescGZIP(), []int{2}
}

func (x *Filter) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Filter) GetRecursive() bool {
	if x != nil {
		return x.Recursive
	}
	return false
}

type SetReconciliationRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mode for the range.
	Mode SetReconciliationRange_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=com.seed.p2p.v1alpha.SetReconciliationRange_Mode" json:"mode,omitempty"`
	// Timestamp of the upper bound of the range.
	BoundTimestamp int64 `protobuf:"varint,2,opt,name=bound_timestamp,json=boundTimestamp,proto3" json:"bound_timestamp,omitempty"`
	// Value of the upper bound of the range.
	BoundValue []byte `protobuf:"bytes,3,opt,name=bound_value,json=boundValue,proto3" json:"bound_value,omitempty"`
	// Only for LIST mode. List of values in the range.
	Values [][]byte `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty"`
	// Only for the FINGERPRINT mode. Fingerprint of the range.
	Fingerprint []byte `protobuf:"bytes,5,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
}

func (x *SetReconciliationRange) Reset() {
	*x = SetReconciliationRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_v1alpha_syncing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetReconciliationRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetReconciliationRange) ProtoMessage() {}

func (x *SetReconciliationRange) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_v1alpha_syncing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetReconciliationRange.ProtoReflect.Descriptor instead.
func (*SetReconciliationRange) Descriptor() ([]byte, []int) {
	return file_p2p_v1alpha_syncing_proto_rawDescGZIP(), []int{3}
}

func (x *SetReconciliationRange) GetMode() SetReconciliationRange_Mode {
	if x != nil {
		return x.Mode
	}
	return SetReconciliationRange_SKIP
}

func (x *SetReconciliationRange) GetBoundTimestamp() int64 {
	if x != nil {
		return x.BoundTimestamp
	}
	return 0
}

func (x *SetReconciliationRange) GetBoundValue() []byte {
	if x != nil {
		return x.BoundValue
	}
	return nil
}

func (x *SetReconciliationRange) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *SetReconciliationRange) GetFingerprint() []byte {
	if x != nil {
		return x.Fingerprint
	}
	return nil
}

var File_p2p_v1alpha_syncing_proto protoreflect.FileDescriptor

var file_p2p_v1alpha_syncing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x32, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x79,
	0x6e, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x42,
	0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x70,
	0x32, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x5e, 0x0a, 0x16, 0x52, 0x65, 0x63,
	0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x70,
	0x32, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x06, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x22, 0x42, 0x0a, 0x06, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0x90, 0x02,
	0x0a, 0x16, 0x53, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65,
	0x64, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53,
	0x4b, 0x49, 0x50, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x4e, 0x47, 0x45, 0x52, 0x50,
	0x52, 0x49, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x02,
	0x32, 0x76, 0x0a, 0x07, 0x53, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x6b, 0x0a, 0x0e, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x2b, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x73, 0x65, 0x65, 0x64,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x70, 0x32,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_v1alpha_syncing_proto_rawDescOnce sync.Once
	file_p2p_v1alpha_syncing_proto_rawDescData = file_p2p_v1alpha_syncing_proto_rawDesc
)

func file_p2p_v1alpha_syncing_proto_rawDescGZIP() []byte {
	file_p2p_v1alpha_syncing_proto_rawDescOnce.Do(func() {
		file_p2p_v1alpha_syncing_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_v1alpha_syncing_proto_rawDescData)
	})
	return file_p2p_v1alpha_syncing_proto_rawDescData
}

var file_p2p_v1alpha_syncing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_p2p_v1alpha_syncing_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_p2p_v1alpha_syncing_proto_goTypes = []any{
	(SetReconciliationRange_Mode)(0), // 0: com.seed.p2p.v1alpha.SetReconciliationRange.Mode
	(*ReconcileBlobsRequest)(nil),    // 1: com.seed.p2p.v1alpha.ReconcileBlobsRequest
	(*ReconcileBlobsResponse)(nil),   // 2: com.seed.p2p.v1alpha.ReconcileBlobsResponse
	(*Filter)(nil),                   // 3: com.seed.p2p.v1alpha.Filter
	(*SetReconciliationRange)(nil),   // 4: com.seed.p2p.v1alpha.SetReconciliationRange
}
var file_p2p_v1alpha_syncing_proto_depIdxs = []int32{
	3, // 0: com.seed.p2p.v1alpha.ReconcileBlobsRequest.filters:type_name -> com.seed.p2p.v1alpha.Filter
	4, // 1: com.seed.p2p.v1alpha.ReconcileBlobsRequest.ranges:type_name -> com.seed.p2p.v1alpha.SetReconciliationRange
	4, // 2: com.seed.p2p.v1alpha.ReconcileBlobsResponse.ranges:type_name -> com.seed.p2p.v1alpha.SetReconciliationRange
	0, // 3: com.seed.p2p.v1alpha.SetReconciliationRange.mode:type_name -> com.seed.p2p.v1alpha.SetReconciliationRange.Mode
	1, // 4: com.seed.p2p.v1alpha.Syncing.ReconcileBlobs:input_type -> com.seed.p2p.v1alpha.ReconcileBlobsRequest
	2, // 5: com.seed.p2p.v1alpha.Syncing.ReconcileBlobs:output_type -> com.seed.p2p.v1alpha.ReconcileBlobsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_p2p_v1alpha_syncing_proto_init() }
func file_p2p_v1alpha_syncing_proto_init() {
	if File_p2p_v1alpha_syncing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_v1alpha_syncing_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ReconcileBlobsRequest); i {
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
		file_p2p_v1alpha_syncing_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ReconcileBlobsResponse); i {
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
		file_p2p_v1alpha_syncing_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Filter); i {
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
		file_p2p_v1alpha_syncing_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SetReconciliationRange); i {
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
			RawDescriptor: file_p2p_v1alpha_syncing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_p2p_v1alpha_syncing_proto_goTypes,
		DependencyIndexes: file_p2p_v1alpha_syncing_proto_depIdxs,
		EnumInfos:         file_p2p_v1alpha_syncing_proto_enumTypes,
		MessageInfos:      file_p2p_v1alpha_syncing_proto_msgTypes,
	}.Build()
	File_p2p_v1alpha_syncing_proto = out.File
	file_p2p_v1alpha_syncing_proto_rawDesc = nil
	file_p2p_v1alpha_syncing_proto_goTypes = nil
	file_p2p_v1alpha_syncing_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.24.4
// source: documents/v3alpha/access_control.proto

package documents

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	// Invalid default value.
	Role_ROLE_UNSPECIFIED Role = 0
	// Has write access to the document
	Role_WRITER Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		2: "WRITER",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"WRITER":           2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_documents_v3alpha_access_control_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_documents_v3alpha_access_control_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_documents_v3alpha_access_control_proto_rawDescGZIP(), []int{0}
}

// Request to list capabilities.
type ListCapabilitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Account for which to list the capabilities.
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// Required. Path within the account to list the capabilities for.
	// Empty string means root document.
	// String "*" means all documents.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Optional. By default all capabilities that match the path are returned,
	// even if they were issued for some parent path.
	// If this field is true, only capabilities that match the path exactly are returned.
	IgnoreInherited bool `protobuf:"varint,3,opt,name=ignore_inherited,json=ignoreInherited,proto3" json:"ignore_inherited,omitempty"`
	// Optional. Number of capabilities to return in the response.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. Page token to continue listing capabilities.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListCapabilitiesRequest) Reset() {
	*x = ListCapabilitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_v3alpha_access_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCapabilitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapabilitiesRequest) ProtoMessage() {}

func (x *ListCapabilitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_documents_v3alpha_access_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapabilitiesRequest.ProtoReflect.Descriptor instead.
func (*ListCapabilitiesRequest) Descriptor() ([]byte, []int) {
	return file_documents_v3alpha_access_control_proto_rawDescGZIP(), []int{0}
}

func (x *ListCapabilitiesRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *ListCapabilitiesRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListCapabilitiesRequest) GetIgnoreInherited() bool {
	if x != nil {
		return x.IgnoreInherited
	}
	return false
}

func (x *ListCapabilitiesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCapabilitiesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response to list capabilities.
type ListCapabilitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of capabilities.
	Capabilities []*Capability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	// Token for the next page, if any.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCapabilitiesResponse) Reset() {
	*x = ListCapabilitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_v3alpha_access_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCapabilitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapabilitiesResponse) ProtoMessage() {}

func (x *ListCapabilitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_documents_v3alpha_access_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapabilitiesResponse.ProtoReflect.Descriptor instead.
func (*ListCapabilitiesResponse) Descriptor() ([]byte, []int) {
	return file_documents_v3alpha_access_control_proto_rawDescGZIP(), []int{1}
}

func (x *ListCapabilitiesResponse) GetCapabilities() []*Capability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *ListCapabilitiesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request to create a new capability.
type CreateCapabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Name of the key to use for signing the capability.
	SigningKeyName string `protobuf:"bytes,1,opt,name=signing_key_name,json=signingKeyName,proto3" json:"signing_key_name,omitempty"`
	// Required. Account ID to which this capability is delegated.
	Delegate string `protobuf:"bytes,2,opt,name=delegate,proto3" json:"delegate,omitempty"`
	// Required. Account ID to which this capability gives access.
	Account string `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	// Required. Path within the account that this capability grants access to.
	// Empty string means root document.
	Path string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// Optional. By default capabilities give access to the path recursively.
	// This flag can be used to restrict the capability only to specific path.
	NoRecursive bool `protobuf:"varint,5,opt,name=no_recursive,json=noRecursive,proto3" json:"no_recursive,omitempty"`
}

func (x *CreateCapabilityRequest) Reset() {
	*x = CreateCapabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_v3alpha_access_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCapabilityRequest) ProtoMessage() {}

func (x *CreateCapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_documents_v3alpha_access_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCapabilityRequest.ProtoReflect.Descriptor instead.
func (*CreateCapabilityRequest) Descriptor() ([]byte, []int) {
	return file_documents_v3alpha_access_control_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCapabilityRequest) GetSigningKeyName() string {
	if x != nil {
		return x.SigningKeyName
	}
	return ""
}

func (x *CreateCapabilityRequest) GetDelegate() string {
	if x != nil {
		return x.Delegate
	}
	return ""
}

func (x *CreateCapabilityRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CreateCapabilityRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateCapabilityRequest) GetNoRecursive() bool {
	if x != nil {
		return x.NoRecursive
	}
	return false
}

// Request to get a single capability.
type GetCapabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. ID of the capability to get.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCapabilityRequest) Reset() {
	*x = GetCapabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_v3alpha_access_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCapabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCapabilityRequest) ProtoMessage() {}

func (x *GetCapabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_documents_v3alpha_access_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCapabilityRequest.ProtoReflect.Descriptor instead.
func (*GetCapabilityRequest) Descriptor() ([]byte, []int) {
	return file_documents_v3alpha_access_control_proto_rawDescGZIP(), []int{3}
}

func (x *GetCapabilityRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Capability is an unforgeable token that grants access to a specific path within an account.
type Capability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of this capability.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the account that issued the capability.
	Issuer string `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// ID of the account that the capability is delegated to.
	Delegate string `protobuf:"bytes,3,opt,name=delegate,proto3" json:"delegate,omitempty"`
	// Account ID that capability grants access to.
	// This is the same as issuer when it's a first-grade capability,
	// but issuer can be different if the capability is delegated further down.
	Account string `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	// Path within the account which the capability grants access to.
	Path string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	// Role that the capability grants to the delegate.
	Role Role `protobuf:"varint,6,opt,name=role,proto3,enum=com.seed.documents.v3alpha.Role" json:"role,omitempty"`
	// Normally capabilities are applied recursively (i.e. path + all the subpaths),
	// but it can be limited to only to the exact path match.
	IsExact bool `protobuf:"varint,7,opt,name=is_exact,json=isExact,proto3" json:"is_exact,omitempty"`
	// IDs of parent capabilities, unless it's a first-grade capability,
	// in which case it has no parents.
	Parents []string `protobuf:"bytes,8,rep,name=parents,proto3" json:"parents,omitempty"`
	// IDs of children capabilities
	// if this capability was ever delegated further down.
	Children []string `protobuf:"bytes,9,rep,name=children,proto3" json:"children,omitempty"`
	// Timestamp when this capability was issued.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Local timestamp when we have discovered this capability.
	LocalTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=local_time,json=localTime,proto3" json:"local_time,omitempty"`
}

func (x *Capability) Reset() {
	*x = Capability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_documents_v3alpha_access_control_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capability) ProtoMessage() {}

func (x *Capability) ProtoReflect() protoreflect.Message {
	mi := &file_documents_v3alpha_access_control_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capability.ProtoReflect.Descriptor instead.
func (*Capability) Descriptor() ([]byte, []int) {
	return file_documents_v3alpha_access_control_proto_rawDescGZIP(), []int{4}
}

func (x *Capability) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Capability) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Capability) GetDelegate() string {
	if x != nil {
		return x.Delegate
	}
	return ""
}

func (x *Capability) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *Capability) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Capability) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *Capability) GetIsExact() bool {
	if x != nil {
		return x.IsExact
	}
	return false
}

func (x *Capability) GetParents() []string {
	if x != nil {
		return x.Parents
	}
	return nil
}

func (x *Capability) GetChildren() []string {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Capability) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Capability) GetLocalTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LocalTime
	}
	return nil
}

var File_documents_v3alpha_access_control_proto protoreflect.FileDescriptor

var file_documents_v3alpha_access_control_proto_rawDesc = []byte{
	0x0a, 0x26, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65,
	0x65, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x29, 0x0a, 0x10, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x6e, 0x68, 0x65, 0x72, 0x69,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x49, 0x6e, 0x68, 0x65, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8e, 0x01, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb0, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x72, 0x65,
	0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x6e,
	0x6f, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0xfd, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x6c,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x34, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45,
	0x78, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x69,
	0x6d, 0x65, 0x2a, 0x28, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f,
	0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x10, 0x02, 0x32, 0xea, 0x02, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x7d,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65,
	0x65, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x33, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65,
	0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x69,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x30, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x73, 0x65, 0x65,
	0x64, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x3b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_documents_v3alpha_access_control_proto_rawDescOnce sync.Once
	file_documents_v3alpha_access_control_proto_rawDescData = file_documents_v3alpha_access_control_proto_rawDesc
)

func file_documents_v3alpha_access_control_proto_rawDescGZIP() []byte {
	file_documents_v3alpha_access_control_proto_rawDescOnce.Do(func() {
		file_documents_v3alpha_access_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_documents_v3alpha_access_control_proto_rawDescData)
	})
	return file_documents_v3alpha_access_control_proto_rawDescData
}

var file_documents_v3alpha_access_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_documents_v3alpha_access_control_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_documents_v3alpha_access_control_proto_goTypes = []interface{}{
	(Role)(0),                        // 0: com.seed.documents.v3alpha.Role
	(*ListCapabilitiesRequest)(nil),  // 1: com.seed.documents.v3alpha.ListCapabilitiesRequest
	(*ListCapabilitiesResponse)(nil), // 2: com.seed.documents.v3alpha.ListCapabilitiesResponse
	(*CreateCapabilityRequest)(nil),  // 3: com.seed.documents.v3alpha.CreateCapabilityRequest
	(*GetCapabilityRequest)(nil),     // 4: com.seed.documents.v3alpha.GetCapabilityRequest
	(*Capability)(nil),               // 5: com.seed.documents.v3alpha.Capability
	(*timestamppb.Timestamp)(nil),    // 6: google.protobuf.Timestamp
}
var file_documents_v3alpha_access_control_proto_depIdxs = []int32{
	5, // 0: com.seed.documents.v3alpha.ListCapabilitiesResponse.capabilities:type_name -> com.seed.documents.v3alpha.Capability
	0, // 1: com.seed.documents.v3alpha.Capability.role:type_name -> com.seed.documents.v3alpha.Role
	6, // 2: com.seed.documents.v3alpha.Capability.create_time:type_name -> google.protobuf.Timestamp
	6, // 3: com.seed.documents.v3alpha.Capability.local_time:type_name -> google.protobuf.Timestamp
	1, // 4: com.seed.documents.v3alpha.AccessControl.ListCapabilities:input_type -> com.seed.documents.v3alpha.ListCapabilitiesRequest
	3, // 5: com.seed.documents.v3alpha.AccessControl.CreateCapability:input_type -> com.seed.documents.v3alpha.CreateCapabilityRequest
	4, // 6: com.seed.documents.v3alpha.AccessControl.GetCapability:input_type -> com.seed.documents.v3alpha.GetCapabilityRequest
	2, // 7: com.seed.documents.v3alpha.AccessControl.ListCapabilities:output_type -> com.seed.documents.v3alpha.ListCapabilitiesResponse
	5, // 8: com.seed.documents.v3alpha.AccessControl.CreateCapability:output_type -> com.seed.documents.v3alpha.Capability
	5, // 9: com.seed.documents.v3alpha.AccessControl.GetCapability:output_type -> com.seed.documents.v3alpha.Capability
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_documents_v3alpha_access_control_proto_init() }
func file_documents_v3alpha_access_control_proto_init() {
	if File_documents_v3alpha_access_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_documents_v3alpha_access_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCapabilitiesRequest); i {
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
		file_documents_v3alpha_access_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCapabilitiesResponse); i {
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
		file_documents_v3alpha_access_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCapabilityRequest); i {
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
		file_documents_v3alpha_access_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCapabilityRequest); i {
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
		file_documents_v3alpha_access_control_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capability); i {
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
			RawDescriptor: file_documents_v3alpha_access_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_documents_v3alpha_access_control_proto_goTypes,
		DependencyIndexes: file_documents_v3alpha_access_control_proto_depIdxs,
		EnumInfos:         file_documents_v3alpha_access_control_proto_enumTypes,
		MessageInfos:      file_documents_v3alpha_access_control_proto_msgTypes,
	}.Build()
	File_documents_v3alpha_access_control_proto = out.File
	file_documents_v3alpha_access_control_proto_rawDesc = nil
	file_documents_v3alpha_access_control_proto_goTypes = nil
	file_documents_v3alpha_access_control_proto_depIdxs = nil
}

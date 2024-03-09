// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: gw.proto

package gwconn

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

type UserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Act   uint32 `protobuf:"varint,1,opt,name=act,proto3" json:"act,omitempty"`
	Uname string `protobuf:"bytes,2,opt,name=uname,proto3" json:"uname,omitempty"`
	Pass  string `protobuf:"bytes,3,opt,name=pass,proto3" json:"pass,omitempty"`
	Role  string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserReq) Reset() {
	*x = UserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReq) ProtoMessage() {}

func (x *UserReq) ProtoReflect() protoreflect.Message {
	mi := &file_gw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReq.ProtoReflect.Descriptor instead.
func (*UserReq) Descriptor() ([]byte, []int) {
	return file_gw_proto_rawDescGZIP(), []int{0}
}

func (x *UserReq) GetAct() uint32 {
	if x != nil {
		return x.Act
	}
	return 0
}

func (x *UserReq) GetUname() string {
	if x != nil {
		return x.Uname
	}
	return ""
}

func (x *UserReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *UserReq) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type BridgeListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bridges []*BridgeInfo `protobuf:"bytes,1,rep,name=bridges,proto3" json:"bridges,omitempty"`
}

func (x *BridgeListResp) Reset() {
	*x = BridgeListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeListResp) ProtoMessage() {}

func (x *BridgeListResp) ProtoReflect() protoreflect.Message {
	mi := &file_gw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeListResp.ProtoReflect.Descriptor instead.
func (*BridgeListResp) Descriptor() ([]byte, []int) {
	return file_gw_proto_rawDescGZIP(), []int{1}
}

func (x *BridgeListResp) GetBridges() []*BridgeInfo {
	if x != nil {
		return x.Bridges
	}
	return nil
}

type SetBridgeStateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State uint32 `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *SetBridgeStateReq) Reset() {
	*x = SetBridgeStateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBridgeStateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBridgeStateReq) ProtoMessage() {}

func (x *SetBridgeStateReq) ProtoReflect() protoreflect.Message {
	mi := &file_gw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBridgeStateReq.ProtoReflect.Descriptor instead.
func (*SetBridgeStateReq) Descriptor() ([]byte, []int) {
	return file_gw_proto_rawDescGZIP(), []int{2}
}

func (x *SetBridgeStateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetBridgeStateReq) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type DomainReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fqdn string `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
}

func (x *DomainReq) Reset() {
	*x = DomainReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainReq) ProtoMessage() {}

func (x *DomainReq) ProtoReflect() protoreflect.Message {
	mi := &file_gw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainReq.ProtoReflect.Descriptor instead.
func (*DomainReq) Descriptor() ([]byte, []int) {
	return file_gw_proto_rawDescGZIP(), []int{3}
}

func (x *DomainReq) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

type ManagedDomain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fqdn  string `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	State uint32 `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *ManagedDomain) Reset() {
	*x = ManagedDomain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagedDomain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagedDomain) ProtoMessage() {}

func (x *ManagedDomain) ProtoReflect() protoreflect.Message {
	mi := &file_gw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagedDomain.ProtoReflect.Descriptor instead.
func (*ManagedDomain) Descriptor() ([]byte, []int) {
	return file_gw_proto_rawDescGZIP(), []int{4}
}

func (x *ManagedDomain) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *ManagedDomain) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

type ManagedDomains struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ManagedDomain `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ManagedDomains) Reset() {
	*x = ManagedDomains{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagedDomains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagedDomains) ProtoMessage() {}

func (x *ManagedDomains) ProtoReflect() protoreflect.Message {
	mi := &file_gw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagedDomains.ProtoReflect.Descriptor instead.
func (*ManagedDomains) Descriptor() ([]byte, []int) {
	return file_gw_proto_rawDescGZIP(), []int{5}
}

func (x *ManagedDomains) GetList() []*ManagedDomain {
	if x != nil {
		return x.List
	}
	return nil
}

var File_gw_proto protoreflect.FileDescriptor

var file_gw_proto_rawDesc = []byte{
	0x0a, 0x08, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x77, 0x63, 0x6f,
	0x6e, 0x6e, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x59, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x75, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x3e, 0x0a, 0x0e, 0x42,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a,
	0x07, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x07, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x11, 0x53,
	0x65, 0x74, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x1f, 0x0a, 0x09, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x22, 0x39, 0x0a, 0x0d, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x3b, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32,
	0xf0, 0x05, 0x0a, 0x02, 0x47, 0x57, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x10, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x12, 0x0f, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x65,
	0x74, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42,
	0x72, 0x69, 0x64, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x3e, 0x0a, 0x0f, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x41, 0x64, 0x64, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x12, 0x14, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x0a, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x0f, 0x2e,
	0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0f, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x10, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x12, 0x0f, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x42, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x15, 0x2e,
	0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x6f,
	0x67, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x09, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0f, 0x2e,
	0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x11, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x77,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x67, 0x77, 0x63,
	0x6f, 0x6e, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x77,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x2e, 0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x66, 0x74, 0x2e, 0x67, 0x77,
	0x63, 0x6f, 0x6e, 0x6e, 0x42, 0x00, 0x50, 0x01, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x66, 0x75, 0x72, 0x61, 0x2f, 0x67, 0x77, 0x63, 0x6f,
	0x6e, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gw_proto_rawDescOnce sync.Once
	file_gw_proto_rawDescData = file_gw_proto_rawDesc
)

func file_gw_proto_rawDescGZIP() []byte {
	file_gw_proto_rawDescOnce.Do(func() {
		file_gw_proto_rawDescData = protoimpl.X.CompressGZIP(file_gw_proto_rawDescData)
	})
	return file_gw_proto_rawDescData
}

var file_gw_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gw_proto_goTypes = []interface{}{
	(*UserReq)(nil),           // 0: gwconn.UserReq
	(*BridgeListResp)(nil),    // 1: gwconn.BridgeListResp
	(*SetBridgeStateReq)(nil), // 2: gwconn.SetBridgeStateReq
	(*DomainReq)(nil),         // 3: gwconn.DomainReq
	(*ManagedDomain)(nil),     // 4: gwconn.ManagedDomain
	(*ManagedDomains)(nil),    // 5: gwconn.ManagedDomains
	(*BridgeInfo)(nil),        // 6: gwconn.BridgeInfo
	(*LoginReq)(nil),          // 7: gwconn.LoginReq
	(*ListReq)(nil),           // 8: gwconn.ListReq
	(*BridgeTarget)(nil),      // 9: gwconn.BridgeTarget
	(*JoinStreamReq)(nil),     // 10: gwconn.JoinStreamReq
	(*LoginResp)(nil),         // 11: gwconn.LoginResp
	(*GeneralResp)(nil),       // 12: gwconn.GeneralResp
	(*StringList)(nil),        // 13: gwconn.StringList
	(*LogMsg)(nil),            // 14: gwconn.LogMsg
}
var file_gw_proto_depIdxs = []int32{
	6,  // 0: gwconn.BridgeListResp.bridges:type_name -> gwconn.BridgeInfo
	4,  // 1: gwconn.ManagedDomains.list:type_name -> gwconn.ManagedDomain
	7,  // 2: gwconn.GW.Login:input_type -> gwconn.LoginReq
	8,  // 3: gwconn.GW.BridgeListNew:input_type -> gwconn.ListReq
	2,  // 4: gwconn.GW.BridgeSetState:input_type -> gwconn.SetBridgeStateReq
	8,  // 5: gwconn.GW.BridgeList:input_type -> gwconn.ListReq
	9,  // 6: gwconn.GW.BridgeAddTarget:input_type -> gwconn.BridgeTarget
	8,  // 7: gwconn.GW.BridgeLogs:input_type -> gwconn.ListReq
	8,  // 8: gwconn.GW.BridgeRouteInfo:input_type -> gwconn.ListReq
	8,  // 9: gwconn.GW.BridgeSystemStat:input_type -> gwconn.ListReq
	10, // 10: gwconn.GW.BridgeStreamLogs:input_type -> gwconn.JoinStreamReq
	3,  // 11: gwconn.GW.DomainAdd:input_type -> gwconn.DomainReq
	8,  // 12: gwconn.GW.DomainList:input_type -> gwconn.ListReq
	3,  // 13: gwconn.GW.DomainVerify:input_type -> gwconn.DomainReq
	0,  // 14: gwconn.GW.User:input_type -> gwconn.UserReq
	11, // 15: gwconn.GW.Login:output_type -> gwconn.LoginResp
	1,  // 16: gwconn.GW.BridgeListNew:output_type -> gwconn.BridgeListResp
	12, // 17: gwconn.GW.BridgeSetState:output_type -> gwconn.GeneralResp
	1,  // 18: gwconn.GW.BridgeList:output_type -> gwconn.BridgeListResp
	12, // 19: gwconn.GW.BridgeAddTarget:output_type -> gwconn.GeneralResp
	13, // 20: gwconn.GW.BridgeLogs:output_type -> gwconn.StringList
	13, // 21: gwconn.GW.BridgeRouteInfo:output_type -> gwconn.StringList
	13, // 22: gwconn.GW.BridgeSystemStat:output_type -> gwconn.StringList
	14, // 23: gwconn.GW.BridgeStreamLogs:output_type -> gwconn.LogMsg
	12, // 24: gwconn.GW.DomainAdd:output_type -> gwconn.GeneralResp
	5,  // 25: gwconn.GW.DomainList:output_type -> gwconn.ManagedDomains
	12, // 26: gwconn.GW.DomainVerify:output_type -> gwconn.GeneralResp
	12, // 27: gwconn.GW.User:output_type -> gwconn.GeneralResp
	15, // [15:28] is the sub-list for method output_type
	2,  // [2:15] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_gw_proto_init() }
func file_gw_proto_init() {
	if File_gw_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReq); i {
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
		file_gw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeListResp); i {
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
		file_gw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBridgeStateReq); i {
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
		file_gw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainReq); i {
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
		file_gw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagedDomain); i {
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
		file_gw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagedDomains); i {
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
			RawDescriptor: file_gw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gw_proto_goTypes,
		DependencyIndexes: file_gw_proto_depIdxs,
		MessageInfos:      file_gw_proto_msgTypes,
	}.Build()
	File_gw_proto = out.File
	file_gw_proto_rawDesc = nil
	file_gw_proto_goTypes = nil
	file_gw_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: bridge.proto

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

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type JoinStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinStreamReq) Reset() {
	*x = JoinStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinStreamReq) ProtoMessage() {}

func (x *JoinStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinStreamReq.ProtoReflect.Descriptor instead.
func (*JoinStreamReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{1}
}

type BridgeTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Act     uint32 `protobuf:"varint,1,opt,name=act,proto3" json:"act,omitempty"`
	TargeId string `protobuf:"bytes,2,opt,name=targeId,proto3" json:"targeId,omitempty"`
	Fqdn    string `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Addr    string `protobuf:"bytes,4,opt,name=addr,proto3" json:"addr,omitempty"`
	Typ     uint32 `protobuf:"varint,5,opt,name=typ,proto3" json:"typ,omitempty"`
}

func (x *BridgeTarget) Reset() {
	*x = BridgeTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeTarget) ProtoMessage() {}

func (x *BridgeTarget) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeTarget.ProtoReflect.Descriptor instead.
func (*BridgeTarget) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{2}
}

func (x *BridgeTarget) GetAct() uint32 {
	if x != nil {
		return x.Act
	}
	return 0
}

func (x *BridgeTarget) GetTargeId() string {
	if x != nil {
		return x.TargeId
	}
	return ""
}

func (x *BridgeTarget) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *BridgeTarget) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *BridgeTarget) GetTyp() uint32 {
	if x != nil {
		return x.Typ
	}
	return 0
}

type DnsChangerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSet bool   `protobuf:"varint,1,opt,name=isSet,proto3" json:"isSet,omitempty"`
	Fqdn  string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Val   string `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DnsChangerReq) Reset() {
	*x = DnsChangerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsChangerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsChangerReq) ProtoMessage() {}

func (x *DnsChangerReq) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsChangerReq.ProtoReflect.Descriptor instead.
func (*DnsChangerReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{3}
}

func (x *DnsChangerReq) GetIsSet() bool {
	if x != nil {
		return x.IsSet
	}
	return false
}

func (x *DnsChangerReq) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *DnsChangerReq) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

type SystemStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpustat  []uint32 `protobuf:"varint,1,rep,packed,name=cpustat,proto3" json:"cpustat,omitempty"`
	MaxRam   uint64   `protobuf:"varint,2,opt,name=maxRam,proto3" json:"maxRam,omitempty"`
	RamBuf   uint64   `protobuf:"varint,3,opt,name=ramBuf,proto3" json:"ramBuf,omitempty"`
	RamUsed  uint64   `protobuf:"varint,4,opt,name=ramUsed,proto3" json:"ramUsed,omitempty"`
	MaxDisk  uint64   `protobuf:"varint,5,opt,name=maxDisk,proto3" json:"maxDisk,omitempty"`
	DiskUsed uint64   `protobuf:"varint,6,opt,name=diskUsed,proto3" json:"diskUsed,omitempty"`
	Conns    uint32   `protobuf:"varint,7,opt,name=conns,proto3" json:"conns,omitempty"`
}

func (x *SystemStat) Reset() {
	*x = SystemStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemStat) ProtoMessage() {}

func (x *SystemStat) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemStat.ProtoReflect.Descriptor instead.
func (*SystemStat) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{4}
}

func (x *SystemStat) GetCpustat() []uint32 {
	if x != nil {
		return x.Cpustat
	}
	return nil
}

func (x *SystemStat) GetMaxRam() uint64 {
	if x != nil {
		return x.MaxRam
	}
	return 0
}

func (x *SystemStat) GetRamBuf() uint64 {
	if x != nil {
		return x.RamBuf
	}
	return 0
}

func (x *SystemStat) GetRamUsed() uint64 {
	if x != nil {
		return x.RamUsed
	}
	return 0
}

func (x *SystemStat) GetMaxDisk() uint64 {
	if x != nil {
		return x.MaxDisk
	}
	return 0
}

func (x *SystemStat) GetDiskUsed() uint64 {
	if x != nil {
		return x.DiskUsed
	}
	return 0
}

func (x *SystemStat) GetConns() uint32 {
	if x != nil {
		return x.Conns
	}
	return 0
}

type RouteStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteId      string `protobuf:"bytes,1,opt,name=routeId,proto3" json:"routeId,omitempty"`
	Conns        uint64 `protobuf:"varint,2,opt,name=conns,proto3" json:"conns,omitempty"`
	TrafOut      uint64 `protobuf:"varint,3,opt,name=trafOut,proto3" json:"trafOut,omitempty"`
	TrafOutDelta uint64 `protobuf:"varint,4,opt,name=trafOutDelta,proto3" json:"trafOutDelta,omitempty"`
	TrafIn       uint64 `protobuf:"varint,5,opt,name=trafIn,proto3" json:"trafIn,omitempty"`
	TrafInDelta  uint64 `protobuf:"varint,6,opt,name=trafInDelta,proto3" json:"trafInDelta,omitempty"`
}

func (x *RouteStat) Reset() {
	*x = RouteStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteStat) ProtoMessage() {}

func (x *RouteStat) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteStat.ProtoReflect.Descriptor instead.
func (*RouteStat) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{5}
}

func (x *RouteStat) GetRouteId() string {
	if x != nil {
		return x.RouteId
	}
	return ""
}

func (x *RouteStat) GetConns() uint64 {
	if x != nil {
		return x.Conns
	}
	return 0
}

func (x *RouteStat) GetTrafOut() uint64 {
	if x != nil {
		return x.TrafOut
	}
	return 0
}

func (x *RouteStat) GetTrafOutDelta() uint64 {
	if x != nil {
		return x.TrafOutDelta
	}
	return 0
}

func (x *RouteStat) GetTrafIn() uint64 {
	if x != nil {
		return x.TrafIn
	}
	return 0
}

func (x *RouteStat) GetTrafInDelta() uint64 {
	if x != nil {
		return x.TrafInDelta
	}
	return 0
}

type LogLineMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Line   string `protobuf:"bytes,2,opt,name=line,proto3" json:"line,omitempty"`
}

func (x *LogLineMsg) Reset() {
	*x = LogLineMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLineMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLineMsg) ProtoMessage() {}

func (x *LogLineMsg) ProtoReflect() protoreflect.Message {
	mi := &file_bridge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLineMsg.ProtoReflect.Descriptor instead.
func (*LogLineMsg) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{6}
}

func (x *LogLineMsg) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *LogLineMsg) GetLine() string {
	if x != nil {
		return x.Line
	}
	return ""
}

var File_bridge_proto protoreflect.FileDescriptor

var file_bridge_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x0f, 0x0a, 0x0d, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x22, 0x74, 0x0a, 0x0c, 0x42, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74,
	0x79, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x79, 0x70, 0x22, 0x4b, 0x0a,
	0x0d, 0x44, 0x6e, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x73, 0x53, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69,
	0x73, 0x53, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0xbc, 0x01, 0x0a, 0x0a, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x70, 0x75,
	0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x70, 0x75, 0x73,
	0x74, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x6d, 0x61, 0x78, 0x52, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x61, 0x6d, 0x42, 0x75, 0x66, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x61, 0x6d,
	0x42, 0x75, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x78, 0x44, 0x69, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x6d, 0x61, 0x78, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x55,
	0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x55,
	0x73, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x09, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x66, 0x4f,
	0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x74, 0x72, 0x61, 0x66, 0x4f, 0x75,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x4f, 0x75, 0x74, 0x44, 0x65, 0x6c, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x66, 0x4f, 0x75, 0x74,
	0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x66, 0x49, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x72, 0x61, 0x66, 0x49, 0x6e, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x66, 0x49, 0x6e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x66, 0x49, 0x6e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x22,
	0x38, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x32, 0xae, 0x03, 0x0a, 0x06, 0x42, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x42, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x12, 0x10, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x67,
	0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x14, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x67, 0x77,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3e, 0x0a, 0x0e,
	0x44, 0x6e, 0x73, 0x41, 0x63, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x15,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x44, 0x6e, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74,
	0x61, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x2e, 0x67, 0x77, 0x63,
	0x6f, 0x6e, 0x6e, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x1a, 0x13, 0x2e,
	0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x12,
	0x12, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x6e, 0x65,
	0x4d, 0x73, 0x67, 0x1a, 0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x2e, 0x0a, 0x0e, 0x69, 0x6f,
	0x2e, 0x74, 0x79, 0x66, 0x74, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x42, 0x00, 0x50, 0x01,
	0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x79, 0x66,
	0x75, 0x72, 0x61, 0x2f, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_bridge_proto_rawDescOnce sync.Once
	file_bridge_proto_rawDescData = file_bridge_proto_rawDesc
)

func file_bridge_proto_rawDescGZIP() []byte {
	file_bridge_proto_rawDescOnce.Do(func() {
		file_bridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_bridge_proto_rawDescData)
	})
	return file_bridge_proto_rawDescData
}

var file_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bridge_proto_goTypes = []interface{}{
	(*RegisterResp)(nil),  // 0: gwconn.RegisterResp
	(*JoinStreamReq)(nil), // 1: gwconn.JoinStreamReq
	(*BridgeTarget)(nil),  // 2: gwconn.BridgeTarget
	(*DnsChangerReq)(nil), // 3: gwconn.DnsChangerReq
	(*SystemStat)(nil),    // 4: gwconn.SystemStat
	(*RouteStat)(nil),     // 5: gwconn.RouteStat
	(*LogLineMsg)(nil),    // 6: gwconn.LogLineMsg
	(*LoginReq)(nil),      // 7: gwconn.LoginReq
	(*BridgeInfo)(nil),    // 8: gwconn.BridgeInfo
	(*LoginResp)(nil),     // 9: gwconn.LoginResp
	(*GeneralResp)(nil),   // 10: gwconn.GeneralResp
}
var file_bridge_proto_depIdxs = []int32{
	7,  // 0: gwconn.Bridge.LoginBridge:input_type -> gwconn.LoginReq
	8,  // 1: gwconn.Bridge.RegisterBridge:input_type -> gwconn.BridgeInfo
	1,  // 2: gwconn.Bridge.GetTargetStream:input_type -> gwconn.JoinStreamReq
	3,  // 3: gwconn.Bridge.DnsAcmeChanger:input_type -> gwconn.DnsChangerReq
	4,  // 4: gwconn.Bridge.SendSystemStat:input_type -> gwconn.SystemStat
	5,  // 5: gwconn.Bridge.SendRouteInfo:input_type -> gwconn.RouteStat
	6,  // 6: gwconn.Bridge.SendLog:input_type -> gwconn.LogLineMsg
	9,  // 7: gwconn.Bridge.LoginBridge:output_type -> gwconn.LoginResp
	0,  // 8: gwconn.Bridge.RegisterBridge:output_type -> gwconn.RegisterResp
	2,  // 9: gwconn.Bridge.GetTargetStream:output_type -> gwconn.BridgeTarget
	10, // 10: gwconn.Bridge.DnsAcmeChanger:output_type -> gwconn.GeneralResp
	10, // 11: gwconn.Bridge.SendSystemStat:output_type -> gwconn.GeneralResp
	10, // 12: gwconn.Bridge.SendRouteInfo:output_type -> gwconn.GeneralResp
	10, // 13: gwconn.Bridge.SendLog:output_type -> gwconn.GeneralResp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_bridge_proto_init() }
func file_bridge_proto_init() {
	if File_bridge_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_bridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinStreamReq); i {
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
		file_bridge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeTarget); i {
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
		file_bridge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsChangerReq); i {
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
		file_bridge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemStat); i {
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
		file_bridge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteStat); i {
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
		file_bridge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLineMsg); i {
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
			RawDescriptor: file_bridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bridge_proto_goTypes,
		DependencyIndexes: file_bridge_proto_depIdxs,
		MessageInfos:      file_bridge_proto_msgTypes,
	}.Build()
	File_bridge_proto = out.File
	file_bridge_proto_rawDesc = nil
	file_bridge_proto_goTypes = nil
	file_bridge_proto_depIdxs = nil
}

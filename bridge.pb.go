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

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
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

type DnsChallReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSet bool   `protobuf:"varint,1,opt,name=isSet,proto3" json:"isSet,omitempty"`
	Fqdn  string `protobuf:"bytes,2,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Val   string `protobuf:"bytes,3,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *DnsChallReq) Reset() {
	*x = DnsChallReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsChallReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsChallReq) ProtoMessage() {}

func (x *DnsChallReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DnsChallReq.ProtoReflect.Descriptor instead.
func (*DnsChallReq) Descriptor() ([]byte, []int) {
	return file_bridge_proto_rawDescGZIP(), []int{1}
}

func (x *DnsChallReq) GetIsSet() bool {
	if x != nil {
		return x.IsSet
	}
	return false
}

func (x *DnsChallReq) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *DnsChallReq) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_bridge_proto protoreflect.FileDescriptor

var file_bridge_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x49, 0x0a, 0x0b, 0x44, 0x6e, 0x73, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x53, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x53, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x71, 0x64, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x32, 0xfa, 0x02, 0x0a, 0x06, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x06,
	0x42, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e,
	0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x09, 0x42, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x77, 0x63,
	0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x14,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x42, 0x56, 0x70, 0x6e, 0x12, 0x0e,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x56, 0x70, 0x6e, 0x43, 0x66, 0x67, 0x22, 0x00,
	0x12, 0x32, 0x0a, 0x05, 0x42, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x67, 0x77, 0x63, 0x6f,
	0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x1a, 0x13, 0x2e,
	0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x42, 0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x2e, 0x67,
	0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x73, 0x67, 0x1a, 0x13, 0x2e, 0x67,
	0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x42, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x4a, 0x6f,
	0x69, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x67, 0x77,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x09, 0x42, 0x44, 0x6e, 0x73, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x44, 0x6e, 0x73, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e,
	0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x2e,
	0x0a, 0x0e, 0x69, 0x6f, 0x2e, 0x74, 0x79, 0x66, 0x74, 0x2e, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e,
	0x42, 0x00, 0x50, 0x01, 0x5a, 0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x79, 0x66, 0x75, 0x72, 0x61, 0x2f, 0x67, 0x77, 0x63, 0x6f, 0x6e, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_bridge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bridge_proto_goTypes = []interface{}{
	(*RegisterResp)(nil),  // 0: gwconn.RegisterResp
	(*DnsChallReq)(nil),   // 1: gwconn.DnsChallReq
	(*LoginReq)(nil),      // 2: gwconn.LoginReq
	(*BridgeInfo)(nil),    // 3: gwconn.BridgeInfo
	(*VpnReq)(nil),        // 4: gwconn.VpnReq
	(*BridgeStat)(nil),    // 5: gwconn.BridgeStat
	(*LogMsg)(nil),        // 6: gwconn.LogMsg
	(*JoinStreamReq)(nil), // 7: gwconn.JoinStreamReq
	(*LoginResp)(nil),     // 8: gwconn.LoginResp
	(*VpnCfg)(nil),        // 9: gwconn.VpnCfg
	(*GeneralResp)(nil),   // 10: gwconn.GeneralResp
	(*BridgeTarget)(nil),  // 11: gwconn.BridgeTarget
}
var file_bridge_proto_depIdxs = []int32{
	2,  // 0: gwconn.Bridge.BLogin:input_type -> gwconn.LoginReq
	3,  // 1: gwconn.Bridge.BRegister:input_type -> gwconn.BridgeInfo
	4,  // 2: gwconn.Bridge.BVpn:input_type -> gwconn.VpnReq
	5,  // 3: gwconn.Bridge.BStat:input_type -> gwconn.BridgeStat
	6,  // 4: gwconn.Bridge.BLog:input_type -> gwconn.LogMsg
	7,  // 5: gwconn.Bridge.BTargetStream:input_type -> gwconn.JoinStreamReq
	1,  // 6: gwconn.Bridge.BDnsChall:input_type -> gwconn.DnsChallReq
	8,  // 7: gwconn.Bridge.BLogin:output_type -> gwconn.LoginResp
	0,  // 8: gwconn.Bridge.BRegister:output_type -> gwconn.RegisterResp
	9,  // 9: gwconn.Bridge.BVpn:output_type -> gwconn.VpnCfg
	10, // 10: gwconn.Bridge.BStat:output_type -> gwconn.GeneralResp
	10, // 11: gwconn.Bridge.BLog:output_type -> gwconn.GeneralResp
	11, // 12: gwconn.Bridge.BTargetStream:output_type -> gwconn.BridgeTarget
	10, // 13: gwconn.Bridge.BDnsChall:output_type -> gwconn.GeneralResp
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
			switch v := v.(*DnsChallReq); i {
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
			NumMessages:   2,
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

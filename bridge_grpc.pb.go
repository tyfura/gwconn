// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: bridge.proto

package gwconn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Bridge_BLogin_FullMethodName     = "/gwconn.Bridge/BLogin"
	Bridge_BRegister_FullMethodName  = "/gwconn.Bridge/BRegister"
	Bridge_BVpn_FullMethodName       = "/gwconn.Bridge/BVpn"
	Bridge_BStat_FullMethodName      = "/gwconn.Bridge/BStat"
	Bridge_BLog_FullMethodName       = "/gwconn.Bridge/BLog"
	Bridge_BCmd_FullMethodName       = "/gwconn.Bridge/BCmd"
	Bridge_BAcmeChall_FullMethodName = "/gwconn.Bridge/BAcmeChall"
)

// BridgeClient is the client API for Bridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeClient interface {
	BLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	BRegister(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	BVpn(ctx context.Context, in *VpnReq, opts ...grpc.CallOption) (*VpnCfg, error)
	BStat(ctx context.Context, in *BridgeStat, opts ...grpc.CallOption) (*GeneralResp, error)
	BLog(ctx context.Context, in *LogMsg, opts ...grpc.CallOption) (*GeneralResp, error)
	BCmd(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BridgeCmd, error)
	BAcmeChall(ctx context.Context, in *AcmeChallReq, opts ...grpc.CallOption) (*GeneralResp, error)
}

type bridgeClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeClient(cc grpc.ClientConnInterface) BridgeClient {
	return &bridgeClient{cc}
}

func (c *bridgeClient) BLogin(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Bridge_BLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) BRegister(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, Bridge_BRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) BVpn(ctx context.Context, in *VpnReq, opts ...grpc.CallOption) (*VpnCfg, error) {
	out := new(VpnCfg)
	err := c.cc.Invoke(ctx, Bridge_BVpn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) BStat(ctx context.Context, in *BridgeStat, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, Bridge_BStat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) BLog(ctx context.Context, in *LogMsg, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, Bridge_BLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) BCmd(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*BridgeCmd, error) {
	out := new(BridgeCmd)
	err := c.cc.Invoke(ctx, Bridge_BCmd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeClient) BAcmeChall(ctx context.Context, in *AcmeChallReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, Bridge_BAcmeChall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeServer is the server API for Bridge service.
// All implementations must embed UnimplementedBridgeServer
// for forward compatibility
type BridgeServer interface {
	BLogin(context.Context, *LoginReq) (*LoginResp, error)
	BRegister(context.Context, *RegisterReq) (*RegisterResp, error)
	BVpn(context.Context, *VpnReq) (*VpnCfg, error)
	BStat(context.Context, *BridgeStat) (*GeneralResp, error)
	BLog(context.Context, *LogMsg) (*GeneralResp, error)
	BCmd(context.Context, *EmptyReq) (*BridgeCmd, error)
	BAcmeChall(context.Context, *AcmeChallReq) (*GeneralResp, error)
	mustEmbedUnimplementedBridgeServer()
}

// UnimplementedBridgeServer must be embedded to have forward compatible implementations.
type UnimplementedBridgeServer struct {
}

func (UnimplementedBridgeServer) BLogin(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BLogin not implemented")
}
func (UnimplementedBridgeServer) BRegister(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BRegister not implemented")
}
func (UnimplementedBridgeServer) BVpn(context.Context, *VpnReq) (*VpnCfg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BVpn not implemented")
}
func (UnimplementedBridgeServer) BStat(context.Context, *BridgeStat) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BStat not implemented")
}
func (UnimplementedBridgeServer) BLog(context.Context, *LogMsg) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BLog not implemented")
}
func (UnimplementedBridgeServer) BCmd(context.Context, *EmptyReq) (*BridgeCmd, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BCmd not implemented")
}
func (UnimplementedBridgeServer) BAcmeChall(context.Context, *AcmeChallReq) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BAcmeChall not implemented")
}
func (UnimplementedBridgeServer) mustEmbedUnimplementedBridgeServer() {}

// UnsafeBridgeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeServer will
// result in compilation errors.
type UnsafeBridgeServer interface {
	mustEmbedUnimplementedBridgeServer()
}

func RegisterBridgeServer(s grpc.ServiceRegistrar, srv BridgeServer) {
	s.RegisterService(&Bridge_ServiceDesc, srv)
}

func _Bridge_BLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bridge_BLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BLogin(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_BRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bridge_BRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BRegister(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_BVpn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VpnReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BVpn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bridge_BVpn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BVpn(ctx, req.(*VpnReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_BStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BridgeStat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bridge_BStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BStat(ctx, req.(*BridgeStat))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_BLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bridge_BLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BLog(ctx, req.(*LogMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_BCmd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BCmd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bridge_BCmd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BCmd(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bridge_BAcmeChall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcmeChallReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeServer).BAcmeChall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Bridge_BAcmeChall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeServer).BAcmeChall(ctx, req.(*AcmeChallReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Bridge_ServiceDesc is the grpc.ServiceDesc for Bridge service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Bridge_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gwconn.Bridge",
	HandlerType: (*BridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BLogin",
			Handler:    _Bridge_BLogin_Handler,
		},
		{
			MethodName: "BRegister",
			Handler:    _Bridge_BRegister_Handler,
		},
		{
			MethodName: "BVpn",
			Handler:    _Bridge_BVpn_Handler,
		},
		{
			MethodName: "BStat",
			Handler:    _Bridge_BStat_Handler,
		},
		{
			MethodName: "BLog",
			Handler:    _Bridge_BLog_Handler,
		},
		{
			MethodName: "BCmd",
			Handler:    _Bridge_BCmd_Handler,
		},
		{
			MethodName: "BAcmeChall",
			Handler:    _Bridge_BAcmeChall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bridge.proto",
}

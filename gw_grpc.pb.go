// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: gw.proto

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
	GW_Login_FullMethodName            = "/gwconn.GW/Login"
	GW_BridgeListNew_FullMethodName    = "/gwconn.GW/BridgeListNew"
	GW_BridgeAck_FullMethodName        = "/gwconn.GW/BridgeAck"
	GW_BridgeList_FullMethodName       = "/gwconn.GW/BridgeList"
	GW_BridgeAddTarget_FullMethodName  = "/gwconn.GW/BridgeAddTarget"
	GW_BridgeLogs_FullMethodName       = "/gwconn.GW/BridgeLogs"
	GW_BridgeRouteInfo_FullMethodName  = "/gwconn.GW/BridgeRouteInfo"
	GW_BridgeSystemStat_FullMethodName = "/gwconn.GW/BridgeSystemStat"
	GW_BridgeStreamLogs_FullMethodName = "/gwconn.GW/BridgeStreamLogs"
	GW_UserAdd_FullMethodName          = "/gwconn.GW/UserAdd"
	GW_UserChange_FullMethodName       = "/gwconn.GW/UserChange"
	GW_UserDel_FullMethodName          = "/gwconn.GW/UserDel"
	GW_UserList_FullMethodName         = "/gwconn.GW/UserList"
	GW_PolicyManage_FullMethodName     = "/gwconn.GW/PolicyManage"
	GW_PolicyList_FullMethodName       = "/gwconn.GW/PolicyList"
)

// GWClient is the client API for GW service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GWClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	// bridges
	BridgeListNew(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error)
	BridgeAck(ctx context.Context, in *BridgeAckReq, opts ...grpc.CallOption) (*GeneralResp, error)
	BridgeList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error)
	BridgeAddTarget(ctx context.Context, in *BridgeTarget, opts ...grpc.CallOption) (*GeneralResp, error)
	BridgeLogs(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error)
	BridgeRouteInfo(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error)
	BridgeSystemStat(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error)
	BridgeStreamLogs(ctx context.Context, in *JoinStreamReq, opts ...grpc.CallOption) (GW_BridgeStreamLogsClient, error)
	// users
	UserAdd(ctx context.Context, in *User, opts ...grpc.CallOption) (*GeneralResp, error)
	UserChange(ctx context.Context, in *User, opts ...grpc.CallOption) (*GeneralResp, error)
	UserDel(ctx context.Context, in *UserDelReq, opts ...grpc.CallOption) (*GeneralResp, error)
	UserList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error)
	// IP policies
	PolicyManage(ctx context.Context, in *PolicyManageReq, opts ...grpc.CallOption) (*GeneralResp, error)
	PolicyList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error)
}

type gWClient struct {
	cc grpc.ClientConnInterface
}

func NewGWClient(cc grpc.ClientConnInterface) GWClient {
	return &gWClient{cc}
}

func (c *gWClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, GW_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeListNew(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, GW_BridgeListNew_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeAck(ctx context.Context, in *BridgeAckReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, GW_BridgeAck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, GW_BridgeList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeAddTarget(ctx context.Context, in *BridgeTarget, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, GW_BridgeAddTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeLogs(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, GW_BridgeLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeRouteInfo(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, GW_BridgeRouteInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeSystemStat(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, GW_BridgeSystemStat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) BridgeStreamLogs(ctx context.Context, in *JoinStreamReq, opts ...grpc.CallOption) (GW_BridgeStreamLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &GW_ServiceDesc.Streams[0], GW_BridgeStreamLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gWBridgeStreamLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GW_BridgeStreamLogsClient interface {
	Recv() (*LogMsg, error)
	grpc.ClientStream
}

type gWBridgeStreamLogsClient struct {
	grpc.ClientStream
}

func (x *gWBridgeStreamLogsClient) Recv() (*LogMsg, error) {
	m := new(LogMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gWClient) UserAdd(ctx context.Context, in *User, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, GW_UserAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) UserChange(ctx context.Context, in *User, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, GW_UserChange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) UserDel(ctx context.Context, in *UserDelReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, GW_UserDel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) UserList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, GW_UserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) PolicyManage(ctx context.Context, in *PolicyManageReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	out := new(GeneralResp)
	err := c.cc.Invoke(ctx, GW_PolicyManage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) PolicyList(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, GW_PolicyList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GWServer is the server API for GW service.
// All implementations must embed UnimplementedGWServer
// for forward compatibility
type GWServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	// bridges
	BridgeListNew(context.Context, *ListReq) (*ListResponse, error)
	BridgeAck(context.Context, *BridgeAckReq) (*GeneralResp, error)
	BridgeList(context.Context, *ListReq) (*ListResponse, error)
	BridgeAddTarget(context.Context, *BridgeTarget) (*GeneralResp, error)
	BridgeLogs(context.Context, *ListReq) (*ListResponse, error)
	BridgeRouteInfo(context.Context, *ListReq) (*ListResponse, error)
	BridgeSystemStat(context.Context, *ListReq) (*ListResponse, error)
	BridgeStreamLogs(*JoinStreamReq, GW_BridgeStreamLogsServer) error
	// users
	UserAdd(context.Context, *User) (*GeneralResp, error)
	UserChange(context.Context, *User) (*GeneralResp, error)
	UserDel(context.Context, *UserDelReq) (*GeneralResp, error)
	UserList(context.Context, *ListReq) (*ListResponse, error)
	// IP policies
	PolicyManage(context.Context, *PolicyManageReq) (*GeneralResp, error)
	PolicyList(context.Context, *ListReq) (*ListResponse, error)
	mustEmbedUnimplementedGWServer()
}

// UnimplementedGWServer must be embedded to have forward compatible implementations.
type UnimplementedGWServer struct {
}

func (UnimplementedGWServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGWServer) BridgeListNew(context.Context, *ListReq) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeListNew not implemented")
}
func (UnimplementedGWServer) BridgeAck(context.Context, *BridgeAckReq) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeAck not implemented")
}
func (UnimplementedGWServer) BridgeList(context.Context, *ListReq) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeList not implemented")
}
func (UnimplementedGWServer) BridgeAddTarget(context.Context, *BridgeTarget) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeAddTarget not implemented")
}
func (UnimplementedGWServer) BridgeLogs(context.Context, *ListReq) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeLogs not implemented")
}
func (UnimplementedGWServer) BridgeRouteInfo(context.Context, *ListReq) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeRouteInfo not implemented")
}
func (UnimplementedGWServer) BridgeSystemStat(context.Context, *ListReq) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeSystemStat not implemented")
}
func (UnimplementedGWServer) BridgeStreamLogs(*JoinStreamReq, GW_BridgeStreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method BridgeStreamLogs not implemented")
}
func (UnimplementedGWServer) UserAdd(context.Context, *User) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (UnimplementedGWServer) UserChange(context.Context, *User) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserChange not implemented")
}
func (UnimplementedGWServer) UserDel(context.Context, *UserDelReq) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDel not implemented")
}
func (UnimplementedGWServer) UserList(context.Context, *ListReq) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedGWServer) PolicyManage(context.Context, *PolicyManageReq) (*GeneralResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyManage not implemented")
}
func (UnimplementedGWServer) PolicyList(context.Context, *ListReq) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PolicyList not implemented")
}
func (UnimplementedGWServer) mustEmbedUnimplementedGWServer() {}

// UnsafeGWServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GWServer will
// result in compilation errors.
type UnsafeGWServer interface {
	mustEmbedUnimplementedGWServer()
}

func RegisterGWServer(s grpc.ServiceRegistrar, srv GWServer) {
	s.RegisterService(&GW_ServiceDesc, srv)
}

func _GW_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeListNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).BridgeListNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_BridgeListNew_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).BridgeListNew(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeAck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BridgeAckReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).BridgeAck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_BridgeAck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).BridgeAck(ctx, req.(*BridgeAckReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).BridgeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_BridgeList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).BridgeList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeAddTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BridgeTarget)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).BridgeAddTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_BridgeAddTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).BridgeAddTarget(ctx, req.(*BridgeTarget))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).BridgeLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_BridgeLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).BridgeLogs(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeRouteInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).BridgeRouteInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_BridgeRouteInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).BridgeRouteInfo(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeSystemStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).BridgeSystemStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_BridgeSystemStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).BridgeSystemStat(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_BridgeStreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(JoinStreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GWServer).BridgeStreamLogs(m, &gWBridgeStreamLogsServer{stream})
}

type GW_BridgeStreamLogsServer interface {
	Send(*LogMsg) error
	grpc.ServerStream
}

type gWBridgeStreamLogsServer struct {
	grpc.ServerStream
}

func (x *gWBridgeStreamLogsServer) Send(m *LogMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _GW_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_UserAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).UserAdd(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_UserChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).UserChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_UserChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).UserChange(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_UserDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).UserDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_UserDel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).UserDel(ctx, req.(*UserDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).UserList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_PolicyManage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyManageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).PolicyManage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_PolicyManage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).PolicyManage(ctx, req.(*PolicyManageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_PolicyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).PolicyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_PolicyList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).PolicyList(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GW_ServiceDesc is the grpc.ServiceDesc for GW service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GW_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gwconn.GW",
	HandlerType: (*GWServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _GW_Login_Handler,
		},
		{
			MethodName: "BridgeListNew",
			Handler:    _GW_BridgeListNew_Handler,
		},
		{
			MethodName: "BridgeAck",
			Handler:    _GW_BridgeAck_Handler,
		},
		{
			MethodName: "BridgeList",
			Handler:    _GW_BridgeList_Handler,
		},
		{
			MethodName: "BridgeAddTarget",
			Handler:    _GW_BridgeAddTarget_Handler,
		},
		{
			MethodName: "BridgeLogs",
			Handler:    _GW_BridgeLogs_Handler,
		},
		{
			MethodName: "BridgeRouteInfo",
			Handler:    _GW_BridgeRouteInfo_Handler,
		},
		{
			MethodName: "BridgeSystemStat",
			Handler:    _GW_BridgeSystemStat_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _GW_UserAdd_Handler,
		},
		{
			MethodName: "UserChange",
			Handler:    _GW_UserChange_Handler,
		},
		{
			MethodName: "UserDel",
			Handler:    _GW_UserDel_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _GW_UserList_Handler,
		},
		{
			MethodName: "PolicyManage",
			Handler:    _GW_PolicyManage_Handler,
		},
		{
			MethodName: "PolicyList",
			Handler:    _GW_PolicyList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BridgeStreamLogs",
			Handler:       _GW_BridgeStreamLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gw.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: gw.proto

package gwconn

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GW_GetNewBridges_FullMethodName = "/gwconn.GW/GetNewBridges"
	GW_AckBridge_FullMethodName     = "/gwconn.GW/AckBridge"
)

// GWClient is the client API for GW service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GWClient interface {
	GetNewBridges(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NewBridgeList, error)
	AckBridge(ctx context.Context, in *AcceptBridge, opts ...grpc.CallOption) (*empty.Empty, error)
}

type gWClient struct {
	cc grpc.ClientConnInterface
}

func NewGWClient(cc grpc.ClientConnInterface) GWClient {
	return &gWClient{cc}
}

func (c *gWClient) GetNewBridges(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*NewBridgeList, error) {
	out := new(NewBridgeList)
	err := c.cc.Invoke(ctx, GW_GetNewBridges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gWClient) AckBridge(ctx context.Context, in *AcceptBridge, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, GW_AckBridge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GWServer is the server API for GW service.
// All implementations must embed UnimplementedGWServer
// for forward compatibility
type GWServer interface {
	GetNewBridges(context.Context, *empty.Empty) (*NewBridgeList, error)
	AckBridge(context.Context, *AcceptBridge) (*empty.Empty, error)
	mustEmbedUnimplementedGWServer()
}

// UnimplementedGWServer must be embedded to have forward compatible implementations.
type UnimplementedGWServer struct {
}

func (UnimplementedGWServer) GetNewBridges(context.Context, *empty.Empty) (*NewBridgeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewBridges not implemented")
}
func (UnimplementedGWServer) AckBridge(context.Context, *AcceptBridge) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AckBridge not implemented")
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

func _GW_GetNewBridges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).GetNewBridges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_GetNewBridges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).GetNewBridges(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GW_AckBridge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptBridge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GWServer).AckBridge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GW_AckBridge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GWServer).AckBridge(ctx, req.(*AcceptBridge))
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
			MethodName: "GetNewBridges",
			Handler:    _GW_GetNewBridges_Handler,
		},
		{
			MethodName: "AckBridge",
			Handler:    _GW_AckBridge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gw.proto",
}

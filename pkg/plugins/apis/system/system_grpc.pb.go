// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragù               v0.2.3
// source: pkg/plugins/apis/system/system.proto

package system

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SystemClient is the client API for System service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemClient interface {
	UseManagementAPI(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type systemClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemClient(cc grpc.ClientConnInterface) SystemClient {
	return &systemClient{cc}
}

func (c *systemClient) UseManagementAPI(ctx context.Context, in *BrokerID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/system.System/UseManagementAPI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServer is the server API for System service.
// All implementations must embed UnimplementedSystemServer
// for forward compatibility
type SystemServer interface {
	UseManagementAPI(context.Context, *BrokerID) (*emptypb.Empty, error)
	mustEmbedUnimplementedSystemServer()
}

// UnimplementedSystemServer must be embedded to have forward compatible implementations.
type UnimplementedSystemServer struct {
}

func (UnimplementedSystemServer) UseManagementAPI(context.Context, *BrokerID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseManagementAPI not implemented")
}
func (UnimplementedSystemServer) mustEmbedUnimplementedSystemServer() {}

// UnsafeSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemServer will
// result in compilation errors.
type UnsafeSystemServer interface {
	mustEmbedUnimplementedSystemServer()
}

func RegisterSystemServer(s grpc.ServiceRegistrar, srv SystemServer) {
	s.RegisterService(&System_ServiceDesc, srv)
}

func _System_UseManagementAPI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServer).UseManagementAPI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/system.System/UseManagementAPI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServer).UseManagementAPI(ctx, req.(*BrokerID))
	}
	return interceptor(ctx, in, info, handler)
}

// System_ServiceDesc is the grpc.ServiceDesc for System service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var System_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.System",
	HandlerType: (*SystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UseManagementAPI",
			Handler:    _System_UseManagementAPI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/plugins/apis/system/system.proto",
}

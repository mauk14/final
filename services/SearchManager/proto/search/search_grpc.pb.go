// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: search.proto

package search

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

// SearchManagerClient is the client API for SearchManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchManagerClient interface {
	Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ResponseSM, error)
}

type searchManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchManagerClient(cc grpc.ClientConnInterface) SearchManagerClient {
	return &searchManagerClient{cc}
}

func (c *searchManagerClient) Get(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*ResponseSM, error) {
	out := new(ResponseSM)
	err := c.cc.Invoke(ctx, "/SearchManager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchManagerServer is the server API for SearchManager service.
// All implementations must embed UnimplementedSearchManagerServer
// for forward compatibility
type SearchManagerServer interface {
	Get(context.Context, *Filter) (*ResponseSM, error)
}

// UnimplementedSearchManagerServer must be embedded to have forward compatible implementations.
type UnimplementedSearchManagerServer struct {
}

func (UnimplementedSearchManagerServer) Get(context.Context, *Filter) (*ResponseSM, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSearchManagerServer) mustEmbedUnimplementedSearchManagerServer() {}

// UnsafeSearchManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchManagerServer will
// result in compilation errors.
type UnsafeSearchManagerServer interface {
	mustEmbedUnimplementedSearchManagerServer()
}

func RegisterSearchManagerServer(s grpc.ServiceRegistrar, srv SearchManagerServer) {
	s.RegisterService(&SearchManager_ServiceDesc, srv)
}

func _SearchManager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SearchManager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchManagerServer).Get(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchManager_ServiceDesc is the grpc.ServiceDesc for SearchManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SearchManager",
	HandlerType: (*SearchManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SearchManager_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.3
// source: parser.proto

package parser

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

// ParserManagerClient is the client API for ParserManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserManagerClient interface {
	GetNewBooks(ctx context.Context, in *RequestRM, opts ...grpc.CallOption) (*ResponseRM, error)
}

type parserManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewParserManagerClient(cc grpc.ClientConnInterface) ParserManagerClient {
	return &parserManagerClient{cc}
}

func (c *parserManagerClient) GetNewBooks(ctx context.Context, in *RequestRM, opts ...grpc.CallOption) (*ResponseRM, error) {
	out := new(ResponseRM)
	err := c.cc.Invoke(ctx, "/ParserManager/GetNewBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserManagerServer is the server API for ParserManager service.
// All implementations must embed UnimplementedParserManagerServer
// for forward compatibility
type ParserManagerServer interface {
	GetNewBooks(context.Context, *RequestRM) (*ResponseRM, error)
}

// UnimplementedParserManagerServer must be embedded to have forward compatible implementations.
type UnimplementedParserManagerServer struct {
}

func (UnimplementedParserManagerServer) GetNewBooks(context.Context, *RequestRM) (*ResponseRM, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewBooks not implemented")
}
func (UnimplementedParserManagerServer) mustEmbedUnimplementedParserManagerServer() {}

// UnsafeParserManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserManagerServer will
// result in compilation errors.
type UnsafeParserManagerServer interface {
	mustEmbedUnimplementedParserManagerServer()
}

func RegisterParserManagerServer(s grpc.ServiceRegistrar, srv ParserManagerServer) {
	s.RegisterService(&ParserManager_ServiceDesc, srv)
}

func _ParserManager_GetNewBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRM)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserManagerServer).GetNewBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ParserManager/GetNewBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserManagerServer).GetNewBooks(ctx, req.(*RequestRM))
	}
	return interceptor(ctx, in, info, handler)
}

// ParserManager_ServiceDesc is the grpc.ServiceDesc for ParserManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParserManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ParserManager",
	HandlerType: (*ParserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewBooks",
			Handler:    _ParserManager_GetNewBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "parser.proto",
}

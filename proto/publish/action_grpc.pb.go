// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: proto/publish/action.proto

package publish

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

// PublishServiceClient is the client API for PublishService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishServiceClient interface {
	// option (gorm.server).autogen = true;
	PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error)
}

type publishServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishServiceClient(cc grpc.ClientConnInterface) PublishServiceClient {
	return &publishServiceClient{cc}
}

func (c *publishServiceClient) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error) {
	out := new(DouyinPublishActionResponse)
	err := c.cc.Invoke(ctx, "/douyin.core.PublishService/PublishAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishServiceServer is the server API for PublishService service.
// All implementations must embed UnimplementedPublishServiceServer
// for forward compatibility
type PublishServiceServer interface {
	// option (gorm.server).autogen = true;
	PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error)
	mustEmbedUnimplementedPublishServiceServer()
}

// UnimplementedPublishServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublishServiceServer struct {
}

func (UnimplementedPublishServiceServer) PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAction not implemented")
}
func (UnimplementedPublishServiceServer) mustEmbedUnimplementedPublishServiceServer() {}

// UnsafePublishServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishServiceServer will
// result in compilation errors.
type UnsafePublishServiceServer interface {
	mustEmbedUnimplementedPublishServiceServer()
}

func RegisterPublishServiceServer(s grpc.ServiceRegistrar, srv PublishServiceServer) {
	s.RegisterService(&PublishService_ServiceDesc, srv)
}

func _PublishService_PublishAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServiceServer).PublishAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/douyin.core.PublishService/PublishAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServiceServer).PublishAction(ctx, req.(*DouyinPublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublishService_ServiceDesc is the grpc.ServiceDesc for PublishService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublishService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "douyin.core.PublishService",
	HandlerType: (*PublishServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishAction",
			Handler:    _PublishService_PublishAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/publish/action.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: comment.proto

package comment

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

// CommentActionClient is the client API for CommentAction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentActionClient interface {
	CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error)
	CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
}

type commentActionClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentActionClient(cc grpc.ClientConnInterface) CommentActionClient {
	return &commentActionClient{cc}
}

func (c *commentActionClient) CommentAction(ctx context.Context, in *CommentActionRequest, opts ...grpc.CallOption) (*CommentActionResponse, error) {
	out := new(CommentActionResponse)
	err := c.cc.Invoke(ctx, "/mini_tiktok.proto.comment.CommentAction/CommentAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentActionClient) CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	out := new(CommentListResponse)
	err := c.cc.Invoke(ctx, "/mini_tiktok.proto.comment.CommentAction/CommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentActionServer is the server API for CommentAction service.
// All implementations must embed UnimplementedCommentActionServer
// for forward compatibility
type CommentActionServer interface {
	CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error)
	CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error)
	mustEmbedUnimplementedCommentActionServer()
}

// UnimplementedCommentActionServer must be embedded to have forward compatible implementations.
type UnimplementedCommentActionServer struct {
}

func (UnimplementedCommentActionServer) CommentAction(context.Context, *CommentActionRequest) (*CommentActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentAction not implemented")
}
func (UnimplementedCommentActionServer) CommentList(context.Context, *CommentListRequest) (*CommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentList not implemented")
}
func (UnimplementedCommentActionServer) mustEmbedUnimplementedCommentActionServer() {}

// UnsafeCommentActionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentActionServer will
// result in compilation errors.
type UnsafeCommentActionServer interface {
	mustEmbedUnimplementedCommentActionServer()
}

func RegisterCommentActionServer(s grpc.ServiceRegistrar, srv CommentActionServer) {
	s.RegisterService(&CommentAction_ServiceDesc, srv)
}

func _CommentAction_CommentAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentActionServer).CommentAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mini_tiktok.proto.comment.CommentAction/CommentAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentActionServer).CommentAction(ctx, req.(*CommentActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentAction_CommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentActionServer).CommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mini_tiktok.proto.comment.CommentAction/CommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentActionServer).CommentList(ctx, req.(*CommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentAction_ServiceDesc is the grpc.ServiceDesc for CommentAction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentAction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mini_tiktok.proto.comment.CommentAction",
	HandlerType: (*CommentActionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentAction",
			Handler:    _CommentAction_CommentAction_Handler,
		},
		{
			MethodName: "CommentList",
			Handler:    _CommentAction_CommentList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
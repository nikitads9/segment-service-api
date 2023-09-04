// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.6
// source: segment.proto

package segment_v1

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

const (
	SegmentV1Service_AddSegment_FullMethodName    = "/segment.service.api.SegmentV1Service/AddSegment"
	SegmentV1Service_RemoveSegment_FullMethodName = "/segment.service.api.SegmentV1Service/RemoveSegment"
)

// SegmentV1ServiceClient is the client API for SegmentV1Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SegmentV1ServiceClient interface {
	AddSegment(ctx context.Context, in *AddSegmentRequest, opts ...grpc.CallOption) (*AddSegmentResponse, error)
	RemoveSegment(ctx context.Context, in *RemoveSegmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type segmentV1ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSegmentV1ServiceClient(cc grpc.ClientConnInterface) SegmentV1ServiceClient {
	return &segmentV1ServiceClient{cc}
}

func (c *segmentV1ServiceClient) AddSegment(ctx context.Context, in *AddSegmentRequest, opts ...grpc.CallOption) (*AddSegmentResponse, error) {
	out := new(AddSegmentResponse)
	err := c.cc.Invoke(ctx, SegmentV1Service_AddSegment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *segmentV1ServiceClient) RemoveSegment(ctx context.Context, in *RemoveSegmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SegmentV1Service_RemoveSegment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SegmentV1ServiceServer is the server API for SegmentV1Service service.
// All implementations must embed UnimplementedSegmentV1ServiceServer
// for forward compatibility
type SegmentV1ServiceServer interface {
	AddSegment(context.Context, *AddSegmentRequest) (*AddSegmentResponse, error)
	RemoveSegment(context.Context, *RemoveSegmentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSegmentV1ServiceServer()
}

// UnimplementedSegmentV1ServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSegmentV1ServiceServer struct {
}

func (UnimplementedSegmentV1ServiceServer) AddSegment(context.Context, *AddSegmentRequest) (*AddSegmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSegment not implemented")
}
func (UnimplementedSegmentV1ServiceServer) RemoveSegment(context.Context, *RemoveSegmentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSegment not implemented")
}
func (UnimplementedSegmentV1ServiceServer) mustEmbedUnimplementedSegmentV1ServiceServer() {}

// UnsafeSegmentV1ServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SegmentV1ServiceServer will
// result in compilation errors.
type UnsafeSegmentV1ServiceServer interface {
	mustEmbedUnimplementedSegmentV1ServiceServer()
}

func RegisterSegmentV1ServiceServer(s grpc.ServiceRegistrar, srv SegmentV1ServiceServer) {
	s.RegisterService(&SegmentV1Service_ServiceDesc, srv)
}

func _SegmentV1Service_AddSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentV1ServiceServer).AddSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SegmentV1Service_AddSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentV1ServiceServer).AddSegment(ctx, req.(*AddSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SegmentV1Service_RemoveSegment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SegmentV1ServiceServer).RemoveSegment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SegmentV1Service_RemoveSegment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SegmentV1ServiceServer).RemoveSegment(ctx, req.(*RemoveSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SegmentV1Service_ServiceDesc is the grpc.ServiceDesc for SegmentV1Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SegmentV1Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "segment.service.api.SegmentV1Service",
	HandlerType: (*SegmentV1ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSegment",
			Handler:    _SegmentV1Service_AddSegment_Handler,
		},
		{
			MethodName: "RemoveSegment",
			Handler:    _SegmentV1Service_RemoveSegment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "segment.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package producers

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

// ProducerServiceClient is the client API for ProducerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProducerServiceClient interface {
	// Return the list of new producers of an Agency. A producer is considered
	// new if:
	// 1. It is not in the list of processed producers.
	// 2. Since its creation, it has not been updated.
	ListNewProducers(ctx context.Context, in *ListNewProducersRequest, opts ...grpc.CallOption) (*ListNewProducersResponse, error)
	// Return the list of updated producers of an Agency. A producer is considered
	// updated if:
	// 1. It is not in the list of processed producers.
	// 2. Since its creation, it has been updated.
	ListUpdatedProducers(ctx context.Context, in *ListUpdatedProducersRequest, opts ...grpc.CallOption) (*ListUpdatedProducersResponse, error)
	// Mark a list of producers as processed, so they will not be returned
	// in future calls to ListNewProducers or ListUpdatedProducers.
	MarkAsProcessed(ctx context.Context, in *MarkAsProcessedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// NewProducer starts the onboarding process for the given producer.
	NewProducer(ctx context.Context, in *NewProducerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// NewProducers does the same as NewProducer but accepts a list of producers
	// to onboard.
	NewProducers(ctx context.Context, in *NewProducersRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type producerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProducerServiceClient(cc grpc.ClientConnInterface) ProducerServiceClient {
	return &producerServiceClient{cc}
}

func (c *producerServiceClient) ListNewProducers(ctx context.Context, in *ListNewProducersRequest, opts ...grpc.CallOption) (*ListNewProducersResponse, error) {
	out := new(ListNewProducersResponse)
	err := c.cc.Invoke(ctx, "/producer.ProducerService/ListNewProducers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *producerServiceClient) ListUpdatedProducers(ctx context.Context, in *ListUpdatedProducersRequest, opts ...grpc.CallOption) (*ListUpdatedProducersResponse, error) {
	out := new(ListUpdatedProducersResponse)
	err := c.cc.Invoke(ctx, "/producer.ProducerService/ListUpdatedProducers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *producerServiceClient) MarkAsProcessed(ctx context.Context, in *MarkAsProcessedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/producer.ProducerService/MarkAsProcessed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *producerServiceClient) NewProducer(ctx context.Context, in *NewProducerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/producer.ProducerService/NewProducer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *producerServiceClient) NewProducers(ctx context.Context, in *NewProducersRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/producer.ProducerService/NewProducers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProducerServiceServer is the server API for ProducerService service.
// All implementations must embed UnimplementedProducerServiceServer
// for forward compatibility
type ProducerServiceServer interface {
	// Return the list of new producers of an Agency. A producer is considered
	// new if:
	// 1. It is not in the list of processed producers.
	// 2. Since its creation, it has not been updated.
	ListNewProducers(context.Context, *ListNewProducersRequest) (*ListNewProducersResponse, error)
	// Return the list of updated producers of an Agency. A producer is considered
	// updated if:
	// 1. It is not in the list of processed producers.
	// 2. Since its creation, it has been updated.
	ListUpdatedProducers(context.Context, *ListUpdatedProducersRequest) (*ListUpdatedProducersResponse, error)
	// Mark a list of producers as processed, so they will not be returned
	// in future calls to ListNewProducers or ListUpdatedProducers.
	MarkAsProcessed(context.Context, *MarkAsProcessedRequest) (*emptypb.Empty, error)
	// NewProducer starts the onboarding process for the given producer.
	NewProducer(context.Context, *NewProducerRequest) (*emptypb.Empty, error)
	// NewProducers does the same as NewProducer but accepts a list of producers
	// to onboard.
	NewProducers(context.Context, *NewProducersRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedProducerServiceServer()
}

// UnimplementedProducerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProducerServiceServer struct {
}

func (UnimplementedProducerServiceServer) ListNewProducers(context.Context, *ListNewProducersRequest) (*ListNewProducersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewProducers not implemented")
}
func (UnimplementedProducerServiceServer) ListUpdatedProducers(context.Context, *ListUpdatedProducersRequest) (*ListUpdatedProducersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUpdatedProducers not implemented")
}
func (UnimplementedProducerServiceServer) MarkAsProcessed(context.Context, *MarkAsProcessedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsProcessed not implemented")
}
func (UnimplementedProducerServiceServer) NewProducer(context.Context, *NewProducerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewProducer not implemented")
}
func (UnimplementedProducerServiceServer) NewProducers(context.Context, *NewProducersRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewProducers not implemented")
}
func (UnimplementedProducerServiceServer) mustEmbedUnimplementedProducerServiceServer() {}

// UnsafeProducerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProducerServiceServer will
// result in compilation errors.
type UnsafeProducerServiceServer interface {
	mustEmbedUnimplementedProducerServiceServer()
}

func RegisterProducerServiceServer(s grpc.ServiceRegistrar, srv ProducerServiceServer) {
	s.RegisterService(&ProducerService_ServiceDesc, srv)
}

func _ProducerService_ListNewProducers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewProducersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).ListNewProducers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/producer.ProducerService/ListNewProducers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).ListNewProducers(ctx, req.(*ListNewProducersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProducerService_ListUpdatedProducers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUpdatedProducersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).ListUpdatedProducers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/producer.ProducerService/ListUpdatedProducers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).ListUpdatedProducers(ctx, req.(*ListUpdatedProducersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProducerService_MarkAsProcessed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkAsProcessedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).MarkAsProcessed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/producer.ProducerService/MarkAsProcessed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).MarkAsProcessed(ctx, req.(*MarkAsProcessedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProducerService_NewProducer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewProducerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).NewProducer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/producer.ProducerService/NewProducer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).NewProducer(ctx, req.(*NewProducerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProducerService_NewProducers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewProducersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).NewProducers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/producer.ProducerService/NewProducers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).NewProducers(ctx, req.(*NewProducersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProducerService_ServiceDesc is the grpc.ServiceDesc for ProducerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProducerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "producer.ProducerService",
	HandlerType: (*ProducerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNewProducers",
			Handler:    _ProducerService_ListNewProducers_Handler,
		},
		{
			MethodName: "ListUpdatedProducers",
			Handler:    _ProducerService_ListUpdatedProducers_Handler,
		},
		{
			MethodName: "MarkAsProcessed",
			Handler:    _ProducerService_MarkAsProcessed_Handler,
		},
		{
			MethodName: "NewProducer",
			Handler:    _ProducerService_NewProducer_Handler,
		},
		{
			MethodName: "NewProducers",
			Handler:    _ProducerService_NewProducers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "producer.proto",
}
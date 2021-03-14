// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package clientproto

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

// CentrifugeUniClient is the client API for CentrifugeUni service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CentrifugeUniClient interface {
	Consume(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (CentrifugeUni_ConsumeClient, error)
}

type centrifugeUniClient struct {
	cc grpc.ClientConnInterface
}

func NewCentrifugeUniClient(cc grpc.ClientConnInterface) CentrifugeUniClient {
	return &centrifugeUniClient{cc}
}

func (c *centrifugeUniClient) Consume(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (CentrifugeUni_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &CentrifugeUni_ServiceDesc.Streams[0], "/protocol.CentrifugeUni/Consume", opts...)
	if err != nil {
		return nil, err
	}
	x := &centrifugeUniConsumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CentrifugeUni_ConsumeClient interface {
	Recv() (*StreamData, error)
	grpc.ClientStream
}

type centrifugeUniConsumeClient struct {
	grpc.ClientStream
}

func (x *centrifugeUniConsumeClient) Recv() (*StreamData, error) {
	m := new(StreamData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CentrifugeUniServer is the server API for CentrifugeUni service.
// All implementations should embed UnimplementedCentrifugeUniServer
// for forward compatibility
type CentrifugeUniServer interface {
	Consume(*ConnectRequest, CentrifugeUni_ConsumeServer) error
}

// UnimplementedCentrifugeUniServer should be embedded to have forward compatible implementations.
type UnimplementedCentrifugeUniServer struct {
}

func (UnimplementedCentrifugeUniServer) Consume(*ConnectRequest, CentrifugeUni_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}

// UnsafeCentrifugeUniServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CentrifugeUniServer will
// result in compilation errors.
type UnsafeCentrifugeUniServer interface {
	mustEmbedUnimplementedCentrifugeUniServer()
}

func RegisterCentrifugeUniServer(s grpc.ServiceRegistrar, srv CentrifugeUniServer) {
	s.RegisterService(&CentrifugeUni_ServiceDesc, srv)
}

func _CentrifugeUni_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CentrifugeUniServer).Consume(m, &centrifugeUniConsumeServer{stream})
}

type CentrifugeUni_ConsumeServer interface {
	Send(*StreamData) error
	grpc.ServerStream
}

type centrifugeUniConsumeServer struct {
	grpc.ServerStream
}

func (x *centrifugeUniConsumeServer) Send(m *StreamData) error {
	return x.ServerStream.SendMsg(m)
}

// CentrifugeUni_ServiceDesc is the grpc.ServiceDesc for CentrifugeUni service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CentrifugeUni_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.CentrifugeUni",
	HandlerType: (*CentrifugeUniServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Consume",
			Handler:       _CentrifugeUni_Consume_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "client.proto",
}

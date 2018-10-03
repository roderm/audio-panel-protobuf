// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/BoxControll.proto

package box_controll // import "github.com/roderm/audio-panel-protobuf/go/box_controll"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import msg "github.com/roderm/audio-panel-protobuf/go/msg"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BoxControll service

type BoxControllClient interface {
	GetBoxes(ctx context.Context, in *msg.BoxFilter, opts ...grpc.CallOption) (BoxControll_GetBoxesClient, error)
	GetDevices(ctx context.Context, in *msg.DeviceFilter, opts ...grpc.CallOption) (BoxControll_GetDevicesClient, error)
	OnUpdate(ctx context.Context, in *msg.UpdateFilter, opts ...grpc.CallOption) (BoxControll_OnUpdateClient, error)
	StateUpdate(ctx context.Context, in *msg.UpdateReq, opts ...grpc.CallOption) (*msg.UpdateRes, error)
}

type boxControllClient struct {
	cc *grpc.ClientConn
}

func NewBoxControllClient(cc *grpc.ClientConn) BoxControllClient {
	return &boxControllClient{cc}
}

func (c *boxControllClient) GetBoxes(ctx context.Context, in *msg.BoxFilter, opts ...grpc.CallOption) (BoxControll_GetBoxesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BoxControll_serviceDesc.Streams[0], c.cc, "/BoxControll/GetBoxes", opts...)
	if err != nil {
		return nil, err
	}
	x := &boxControllGetBoxesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BoxControll_GetBoxesClient interface {
	Recv() (*msg.Box, error)
	grpc.ClientStream
}

type boxControllGetBoxesClient struct {
	grpc.ClientStream
}

func (x *boxControllGetBoxesClient) Recv() (*msg.Box, error) {
	m := new(msg.Box)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *boxControllClient) GetDevices(ctx context.Context, in *msg.DeviceFilter, opts ...grpc.CallOption) (BoxControll_GetDevicesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BoxControll_serviceDesc.Streams[1], c.cc, "/BoxControll/GetDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &boxControllGetDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BoxControll_GetDevicesClient interface {
	Recv() (*msg.Device, error)
	grpc.ClientStream
}

type boxControllGetDevicesClient struct {
	grpc.ClientStream
}

func (x *boxControllGetDevicesClient) Recv() (*msg.Device, error) {
	m := new(msg.Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *boxControllClient) OnUpdate(ctx context.Context, in *msg.UpdateFilter, opts ...grpc.CallOption) (BoxControll_OnUpdateClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BoxControll_serviceDesc.Streams[2], c.cc, "/BoxControll/OnUpdate", opts...)
	if err != nil {
		return nil, err
	}
	x := &boxControllOnUpdateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BoxControll_OnUpdateClient interface {
	Recv() (*msg.Update, error)
	grpc.ClientStream
}

type boxControllOnUpdateClient struct {
	grpc.ClientStream
}

func (x *boxControllOnUpdateClient) Recv() (*msg.Update, error) {
	m := new(msg.Update)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *boxControllClient) StateUpdate(ctx context.Context, in *msg.UpdateReq, opts ...grpc.CallOption) (*msg.UpdateRes, error) {
	out := new(msg.UpdateRes)
	err := grpc.Invoke(ctx, "/BoxControll/StateUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BoxControll service

type BoxControllServer interface {
	GetBoxes(*msg.BoxFilter, BoxControll_GetBoxesServer) error
	GetDevices(*msg.DeviceFilter, BoxControll_GetDevicesServer) error
	OnUpdate(*msg.UpdateFilter, BoxControll_OnUpdateServer) error
	StateUpdate(context.Context, *msg.UpdateReq) (*msg.UpdateRes, error)
}

func RegisterBoxControllServer(s *grpc.Server, srv BoxControllServer) {
	s.RegisterService(&_BoxControll_serviceDesc, srv)
}

func _BoxControll_GetBoxes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(msg.BoxFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BoxControllServer).GetBoxes(m, &boxControllGetBoxesServer{stream})
}

type BoxControll_GetBoxesServer interface {
	Send(*msg.Box) error
	grpc.ServerStream
}

type boxControllGetBoxesServer struct {
	grpc.ServerStream
}

func (x *boxControllGetBoxesServer) Send(m *msg.Box) error {
	return x.ServerStream.SendMsg(m)
}

func _BoxControll_GetDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(msg.DeviceFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BoxControllServer).GetDevices(m, &boxControllGetDevicesServer{stream})
}

type BoxControll_GetDevicesServer interface {
	Send(*msg.Device) error
	grpc.ServerStream
}

type boxControllGetDevicesServer struct {
	grpc.ServerStream
}

func (x *boxControllGetDevicesServer) Send(m *msg.Device) error {
	return x.ServerStream.SendMsg(m)
}

func _BoxControll_OnUpdate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(msg.UpdateFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BoxControllServer).OnUpdate(m, &boxControllOnUpdateServer{stream})
}

type BoxControll_OnUpdateServer interface {
	Send(*msg.Update) error
	grpc.ServerStream
}

type boxControllOnUpdateServer struct {
	grpc.ServerStream
}

func (x *boxControllOnUpdateServer) Send(m *msg.Update) error {
	return x.ServerStream.SendMsg(m)
}

func _BoxControll_StateUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(msg.UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoxControllServer).StateUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BoxControll/StateUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoxControllServer).StateUpdate(ctx, req.(*msg.UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BoxControll_serviceDesc = grpc.ServiceDesc{
	ServiceName: "BoxControll",
	HandlerType: (*BoxControllServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StateUpdate",
			Handler:    _BoxControll_StateUpdate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBoxes",
			Handler:       _BoxControll_GetBoxes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetDevices",
			Handler:       _BoxControll_GetDevices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "OnUpdate",
			Handler:       _BoxControll_OnUpdate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/BoxControll.proto",
}

func init() {
	proto.RegisterFile("services/BoxControll.proto", fileDescriptor_BoxControll_facce1a9d2178300)
}

var fileDescriptor_BoxControll_facce1a9d2178300 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x09, 0x88, 0x1e, 0x73, 0x5c, 0xb3, 0xa0, 0x45, 0x0a, 0x2b, 0xb5, 0xbb, 0xdd, 0xa0,
	0x20, 0xd6, 0x51, 0xbc, 0x52, 0x50, 0x6c, 0x6c, 0x64, 0x93, 0xcc, 0xc5, 0x40, 0x92, 0x89, 0x3b,
	0x13, 0xd9, 0xdf, 0xe3, 0x2f, 0x95, 0xdb, 0x5d, 0xe2, 0x75, 0xdf, 0xbc, 0xf9, 0x78, 0xf0, 0x20,
	0x67, 0x74, 0x3f, 0x5d, 0x8d, 0x6c, 0x4a, 0xf2, 0x8f, 0x34, 0x8a, 0xa3, 0xbe, 0xd7, 0x93, 0x23,
	0xa1, 0xfc, 0x7c, 0x40, 0x66, 0xdb, 0x22, 0x9b, 0x06, 0x0f, 0x4e, 0x8a, 0x2f, 0x96, 0x78, 0xdf,
	0xf5, 0x82, 0x8e, 0x63, 0x7e, 0xfb, 0x9b, 0xc1, 0xfa, 0xa8, 0x44, 0x5d, 0xc2, 0x6a, 0x87, 0x52,
	0x92, 0x47, 0x56, 0xa0, 0x4b, 0xf2, 0xcf, 0x41, 0xcf, 0x4f, 0x0e, 0x5c, 0x64, 0xea, 0x06, 0x60,
	0x87, 0xf2, 0x14, 0xaa, 0x59, 0x6d, 0x74, 0xa4, 0x24, 0x9d, 0xa5, 0xb3, 0xc8, 0xd4, 0x15, 0xac,
	0x5e, 0xc6, 0xf7, 0xa9, 0xb1, 0x82, 0x6a, 0xa3, 0x23, 0x2c, 0x56, 0x3c, 0x8b, 0x4c, 0x5d, 0xc3,
	0xfa, 0x4d, 0xac, 0x60, 0x12, 0x21, 0x7d, 0x5e, 0xf1, 0x3b, 0xff, 0x67, 0x2e, 0x1f, 0x3e, 0xee,
	0xdb, 0x4e, 0xbe, 0xe6, 0x4a, 0xd7, 0x34, 0x18, 0x47, 0x0d, 0xba, 0xc1, 0xd8, 0xb9, 0xe9, 0x68,
	0x3b, 0xd9, 0x11, 0xfb, 0x6d, 0xd8, 0x52, 0xcd, 0x7b, 0xd3, 0x92, 0xa9, 0xc8, 0x7f, 0xd6, 0x69,
	0x4e, 0x75, 0x1a, 0x3e, 0x77, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x33, 0x77, 0xbb, 0x32,
	0x01, 0x00, 0x00,
}
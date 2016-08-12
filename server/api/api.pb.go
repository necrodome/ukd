// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	VersionRequest
	VersionReply
	StartRequest
	StartReply
	StopRequest
	StopReply
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type VersionRequest struct {
}

func (m *VersionRequest) Reset()                    { *m = VersionRequest{} }
func (m *VersionRequest) String() string            { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()               {}
func (*VersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Ukd server version.
type VersionReply struct {
	Major int32 `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor int32 `protobuf:"varint,2,opt,name=minor" json:"minor,omitempty"`
}

func (m *VersionReply) Reset()                    { *m = VersionReply{} }
func (m *VersionReply) String() string            { return proto.CompactTextString(m) }
func (*VersionReply) ProtoMessage()               {}
func (*VersionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request message containing image name and image location.
type StartRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Location string `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message signalling result of start attempt.
type StartReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Ip      string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Reason  string `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
}

func (m *StartReply) Reset()                    { *m = StartReply{} }
func (m *StartReply) String() string            { return proto.CompactTextString(m) }
func (*StartReply) ProtoMessage()               {}
func (*StartReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Request message containing the image name.
type StopRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Response message signalling result of stop attempt.
type StopReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Reason  string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *StopReply) Reset()                    { *m = StopReply{} }
func (m *StopReply) String() string            { return proto.CompactTextString(m) }
func (*StopReply) ProtoMessage()               {}
func (*StopReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*VersionRequest)(nil), "VersionRequest")
	proto.RegisterType((*VersionReply)(nil), "VersionReply")
	proto.RegisterType((*StartRequest)(nil), "StartRequest")
	proto.RegisterType((*StartReply)(nil), "StartReply")
	proto.RegisterType((*StopRequest)(nil), "StopRequest")
	proto.RegisterType((*StopReply)(nil), "StopReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Ukd service

type UkdClient interface {
	// Get Server Version.
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error)
	// Start a Unikernel.
	StartUK(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error)
	// Stop a Unikernel.
	StopUK(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error)
}

type ukdClient struct {
	cc *grpc.ClientConn
}

func NewUkdClient(cc *grpc.ClientConn) UkdClient {
	return &ukdClient{cc}
}

func (c *ukdClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionReply, error) {
	out := new(VersionReply)
	err := grpc.Invoke(ctx, "/Ukd/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) StartUK(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error) {
	out := new(StartReply)
	err := grpc.Invoke(ctx, "/Ukd/StartUK", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) StopUK(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := grpc.Invoke(ctx, "/Ukd/StopUK", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ukd service

type UkdServer interface {
	// Get Server Version.
	GetVersion(context.Context, *VersionRequest) (*VersionReply, error)
	// Start a Unikernel.
	StartUK(context.Context, *StartRequest) (*StartReply, error)
	// Stop a Unikernel.
	StopUK(context.Context, *StopRequest) (*StopReply, error)
}

func RegisterUkdServer(s *grpc.Server, srv UkdServer) {
	s.RegisterService(&_Ukd_serviceDesc, srv)
}

func _Ukd_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_StartUK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).StartUK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/StartUK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).StartUK(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_StopUK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).StopUK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/StopUK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).StopUK(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ukd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Ukd",
	HandlerType: (*UkdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Ukd_GetVersion_Handler,
		},
		{
			MethodName: "StartUK",
			Handler:    _Ukd_StartUK_Handler,
		},
		{
			MethodName: "StopUK",
			Handler:    _Ukd_StopUK_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x6c, 0x52, 0x9a, 0x36, 0xd7, 0xb4, 0xa0, 0x27, 0x84, 0xaa, 0x4c, 0x60, 0x21, 0xc1, 0xe4,
	0x01, 0x36, 0x24, 0x58, 0x19, 0x90, 0x18, 0x52, 0x95, 0xdd, 0x04, 0x0f, 0x81, 0x36, 0x0e, 0xb6,
	0x3b, 0xb0, 0xf3, 0xe1, 0x38, 0xae, 0x49, 0xd3, 0xa5, 0x9b, 0xef, 0xf4, 0xee, 0xdd, 0xbd, 0x33,
	0x52, 0xd1, 0x54, 0xbc, 0xd1, 0xca, 0x2a, 0x76, 0x86, 0xf9, 0x9b, 0xd4, 0xa6, 0x52, 0x75, 0x21,
	0xbf, 0xb7, 0xd2, 0x58, 0xf6, 0x80, 0xac, 0x63, 0x9a, 0xf5, 0x0f, 0x9d, 0x63, 0xb4, 0x11, 0x9f,
	0x4a, 0x2f, 0xa2, 0xcb, 0xe8, 0x76, 0x54, 0xec, 0x80, 0x67, 0xab, 0xda, 0xb1, 0x71, 0x60, 0x5b,
	0xc0, 0x9e, 0x90, 0x2d, 0xad, 0xd0, 0x36, 0xec, 0x22, 0xc2, 0x49, 0x2d, 0x36, 0xd2, 0x4b, 0xd3,
	0xc2, 0xbf, 0x29, 0xc7, 0x64, 0xad, 0x4a, 0x61, 0x9d, 0x81, 0x17, 0xa7, 0x45, 0x87, 0xd9, 0x2b,
	0x10, 0xf4, 0xad, 0xf3, 0x02, 0x63, 0xb3, 0x2d, 0x4b, 0x69, 0x8c, 0x5f, 0x30, 0x29, 0xfe, 0x21,
	0xcd, 0x11, 0x57, 0x4d, 0x50, 0xbb, 0x17, 0x5d, 0x20, 0xd1, 0x52, 0x18, 0xb7, 0x71, 0xe8, 0xb9,
	0x80, 0xd8, 0x15, 0xa6, 0x4b, 0xab, 0x9a, 0x23, 0x71, 0xd8, 0x23, 0xd2, 0xdd, 0xc8, 0x71, 0xc7,
	0xbd, 0x43, 0xdc, 0x77, 0xb8, 0xfb, 0x8d, 0x30, 0x5c, 0x7d, 0x7d, 0x10, 0x07, 0x9e, 0xa5, 0x0d,
	0xc5, 0xd1, 0x29, 0x3f, 0x2c, 0x35, 0x9f, 0xf1, 0x7e, 0xa7, 0x6c, 0x40, 0x37, 0x18, 0xfb, 0x4b,
	0x57, 0x2f, 0x34, 0xe3, 0xfd, 0xce, 0xf2, 0x29, 0xdf, 0x57, 0xe0, 0x06, 0xaf, 0x91, 0xb4, 0xf9,
	0xdc, 0x5c, 0xc6, 0x7b, 0xb7, 0xe4, 0xe0, 0x5d, 0x6c, 0x36, 0x78, 0x4f, 0xfc, 0x6f, 0xde, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x83, 0x61, 0xb7, 0xda, 0x01, 0x00, 0x00,
}

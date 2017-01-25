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
	ImageSignatureRequest
	ImageSignatureReply
	UpdateImageRequest
	UpdateImageReply
	StatusRequest
	StatusReply
	LogRequest
	LogReply
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
	Major int32  `protobuf:"varint,1,opt,name=major" json:"major,omitempty"`
	Minor string `protobuf:"bytes,2,opt,name=minor" json:"minor,omitempty"`
}

func (m *VersionReply) Reset()                    { *m = VersionReply{} }
func (m *VersionReply) String() string            { return proto.CompactTextString(m) }
func (*VersionReply) ProtoMessage()               {}
func (*VersionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request message containing image name and image location.
type StartRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Visor    string `protobuf:"bytes,2,opt,name=visor" json:"visor,omitempty"`
	Location string `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Response message signalling result of start attempt.
type StartReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Ip      string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Info    string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
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
	Info    string `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *StopReply) Reset()                    { *m = StopReply{} }
func (m *StopReply) String() string            { return proto.CompactTextString(m) }
func (*StopReply) ProtoMessage()               {}
func (*StopReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ImageSignatureRequest struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *ImageSignatureRequest) Reset()                    { *m = ImageSignatureRequest{} }
func (m *ImageSignatureRequest) String() string            { return proto.CompactTextString(m) }
func (*ImageSignatureRequest) ProtoMessage()               {}
func (*ImageSignatureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ImageSignatureReply struct {
	Success   bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Info      string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *ImageSignatureReply) Reset()                    { *m = ImageSignatureReply{} }
func (m *ImageSignatureReply) String() string            { return proto.CompactTextString(m) }
func (*ImageSignatureReply) ProtoMessage()               {}
func (*ImageSignatureReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type UpdateImageRequest struct {
	Base    string `protobuf:"bytes,1,opt,name=base" json:"base,omitempty"`
	Basesig []byte `protobuf:"bytes,2,opt,name=basesig,proto3" json:"basesig,omitempty"`
	Newsig  []byte `protobuf:"bytes,3,opt,name=newsig,proto3" json:"newsig,omitempty"`
	Diff    []byte `protobuf:"bytes,4,opt,name=diff,proto3" json:"diff,omitempty"`
}

func (m *UpdateImageRequest) Reset()                    { *m = UpdateImageRequest{} }
func (m *UpdateImageRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateImageRequest) ProtoMessage()               {}
func (*UpdateImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type UpdateImageReply struct {
	Success         bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	StagedImagePath string `protobuf:"bytes,2,opt,name=stagedImagePath" json:"stagedImagePath,omitempty"`
	Info            string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *UpdateImageReply) Reset()                    { *m = UpdateImageReply{} }
func (m *UpdateImageReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateImageReply) ProtoMessage()               {}
func (*UpdateImageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type StatusRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// Response message with status of an application.
type StatusReply struct {
	Success bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Info    string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *StatusReply) Reset()                    { *m = StatusReply{} }
func (m *StatusReply) String() string            { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()               {}
func (*StatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type LogRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *LogRequest) Reset()                    { *m = LogRequest{} }
func (m *LogRequest) String() string            { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()               {}
func (*LogRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type LogReply struct {
	Success    bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	LogContent []byte `protobuf:"bytes,2,opt,name=logContent,proto3" json:"logContent,omitempty"`
	Info       string `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
}

func (m *LogReply) Reset()                    { *m = LogReply{} }
func (m *LogReply) String() string            { return proto.CompactTextString(m) }
func (*LogReply) ProtoMessage()               {}
func (*LogReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*VersionRequest)(nil), "VersionRequest")
	proto.RegisterType((*VersionReply)(nil), "VersionReply")
	proto.RegisterType((*StartRequest)(nil), "StartRequest")
	proto.RegisterType((*StartReply)(nil), "StartReply")
	proto.RegisterType((*StopRequest)(nil), "StopRequest")
	proto.RegisterType((*StopReply)(nil), "StopReply")
	proto.RegisterType((*ImageSignatureRequest)(nil), "ImageSignatureRequest")
	proto.RegisterType((*ImageSignatureReply)(nil), "ImageSignatureReply")
	proto.RegisterType((*UpdateImageRequest)(nil), "UpdateImageRequest")
	proto.RegisterType((*UpdateImageReply)(nil), "UpdateImageReply")
	proto.RegisterType((*StatusRequest)(nil), "StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "StatusReply")
	proto.RegisterType((*LogRequest)(nil), "LogRequest")
	proto.RegisterType((*LogReply)(nil), "LogReply")
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
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error)
	// Stop a Unikernel.
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error)
	// Get image signature.
	GetImageSignature(ctx context.Context, in *ImageSignatureRequest, opts ...grpc.CallOption) (*ImageSignatureReply, error)
	// Update a Unikernel on-disk image.
	UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*UpdateImageReply, error)
	// Status of an app.
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	// Get log of an app.
	GetLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error)
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

func (c *ukdClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartReply, error) {
	out := new(StartReply)
	err := grpc.Invoke(ctx, "/Ukd/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopReply, error) {
	out := new(StopReply)
	err := grpc.Invoke(ctx, "/Ukd/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) GetImageSignature(ctx context.Context, in *ImageSignatureRequest, opts ...grpc.CallOption) (*ImageSignatureReply, error) {
	out := new(ImageSignatureReply)
	err := grpc.Invoke(ctx, "/Ukd/GetImageSignature", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) UpdateImage(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*UpdateImageReply, error) {
	out := new(UpdateImageReply)
	err := grpc.Invoke(ctx, "/Ukd/UpdateImage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := grpc.Invoke(ctx, "/Ukd/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ukdClient) GetLog(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*LogReply, error) {
	out := new(LogReply)
	err := grpc.Invoke(ctx, "/Ukd/GetLog", in, out, c.cc, opts...)
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
	Start(context.Context, *StartRequest) (*StartReply, error)
	// Stop a Unikernel.
	Stop(context.Context, *StopRequest) (*StopReply, error)
	// Get image signature.
	GetImageSignature(context.Context, *ImageSignatureRequest) (*ImageSignatureReply, error)
	// Update a Unikernel on-disk image.
	UpdateImage(context.Context, *UpdateImageRequest) (*UpdateImageReply, error)
	// Status of an app.
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	// Get log of an app.
	GetLog(context.Context, *LogRequest) (*LogReply, error)
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

func _Ukd_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_GetImageSignature_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageSignatureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).GetImageSignature(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/GetImageSignature",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).GetImageSignature(ctx, req.(*ImageSignatureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_UpdateImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).UpdateImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/UpdateImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).UpdateImage(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ukd_GetLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UkdServer).GetLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Ukd/GetLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UkdServer).GetLog(ctx, req.(*LogRequest))
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
			MethodName: "Start",
			Handler:    _Ukd_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Ukd_Stop_Handler,
		},
		{
			MethodName: "GetImageSignature",
			Handler:    _Ukd_GetImageSignature_Handler,
		},
		{
			MethodName: "UpdateImage",
			Handler:    _Ukd_UpdateImage_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Ukd_Status_Handler,
		},
		{
			MethodName: "GetLog",
			Handler:    _Ukd_GetLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xde, 0xfa, 0x8b, 0xe6, 0x9a, 0x76, 0x9b, 0x37, 0xa6, 0x2a, 0x42, 0x68, 0x18, 0x21, 0x55,
	0x42, 0xf2, 0x03, 0x3c, 0x20, 0x78, 0xdd, 0x03, 0x02, 0xf1, 0x80, 0x52, 0x86, 0x78, 0xf5, 0x5a,
	0xb7, 0x64, 0xb4, 0x76, 0x88, 0x5d, 0x10, 0xff, 0x23, 0x7f, 0x14, 0xf6, 0xc5, 0x69, 0x9d, 0x12,
	0xe5, 0x29, 0x77, 0xe7, 0xcf, 0xdf, 0x77, 0xe7, 0xbb, 0x0b, 0x44, 0x3c, 0xcf, 0x58, 0x5e, 0x28,
	0xa3, 0xe8, 0x39, 0x4c, 0xbe, 0x8a, 0x42, 0x67, 0x4a, 0xa6, 0xe2, 0xe7, 0x4e, 0x68, 0x43, 0xdf,
	0x41, 0xbc, 0x8f, 0xe4, 0x9b, 0x3f, 0xe4, 0x0a, 0xfa, 0x5b, 0xfe, 0xa0, 0x8a, 0xe9, 0xe9, 0xcd,
	0xe9, 0xac, 0x9f, 0x96, 0x0e, 0x46, 0x33, 0x69, 0xa3, 0x1d, 0x1b, 0x8d, 0xd2, 0xd2, 0xa1, 0x5f,
	0x20, 0x9e, 0x1b, 0x5e, 0x18, 0xcf, 0x45, 0x08, 0xf4, 0x24, 0xdf, 0x0a, 0xbc, 0x1a, 0xa5, 0x68,
	0xbb, 0x9b, 0xbf, 0x32, 0x7d, 0xb8, 0x89, 0x0e, 0x49, 0x60, 0xb8, 0x51, 0x0b, 0x6e, 0xac, 0xec,
	0xb4, 0x8b, 0x07, 0x7b, 0x9f, 0x7e, 0x04, 0xf0, 0xac, 0x2e, 0x9f, 0x29, 0x3c, 0xd2, 0xbb, 0xc5,
	0x42, 0x68, 0x8d, 0xb4, 0xc3, 0xb4, 0x72, 0xc9, 0x04, 0x3a, 0x59, 0xee, 0x69, 0xad, 0xe5, 0xd4,
	0x33, 0xb9, 0x52, 0x9e, 0x0f, 0x6d, 0xfa, 0x0c, 0x46, 0x73, 0xa3, 0xf2, 0x96, 0x04, 0xe9, 0x5b,
	0x88, 0x4a, 0x48, 0xbb, 0x5a, 0xc5, 0xde, 0x09, 0xd8, 0x5f, 0xc2, 0xe3, 0x0f, 0x5b, 0xbe, 0x16,
	0xf3, 0x6c, 0x2d, 0xb9, 0xd9, 0x15, 0x22, 0xd0, 0xc9, 0xb9, 0xf9, 0x5e, 0xe9, 0x38, 0x9b, 0x72,
	0xb8, 0x3c, 0x06, 0xb7, 0x2b, 0x3e, 0x81, 0x48, 0x57, 0x58, 0x94, 0x8d, 0xd3, 0x43, 0xa0, 0xb1,
	0x5a, 0x09, 0xe4, 0x2e, 0x5f, 0x72, 0x23, 0x50, 0x28, 0x48, 0xe6, 0x9e, 0xeb, 0x7d, 0xd1, 0xce,
	0x76, 0xaa, 0xee, 0x6b, 0xe9, 0x3c, 0x73, 0xe5, 0x92, 0x6b, 0x18, 0x48, 0xf1, 0xdb, 0x1d, 0x74,
	0xf1, 0xc0, 0x7b, 0x8e, 0x65, 0x99, 0xad, 0x56, 0xd3, 0x1e, 0x46, 0xd1, 0xa6, 0x0f, 0x70, 0x5e,
	0xd3, 0x6b, 0xaf, 0x67, 0x06, 0x67, 0xda, 0x58, 0xdc, 0x12, 0xd1, 0x9f, 0xdd, 0xfb, 0x94, 0x8f,
	0x79, 0x1c, 0x6e, 0xac, 0xed, 0x39, 0x8c, 0xed, 0x54, 0x98, 0x9d, 0x6e, 0xeb, 0xe5, 0xdc, 0xb5,
	0xbb, 0x04, 0xb5, 0xe7, 0x62, 0xab, 0xd4, 0x08, 0xf4, 0x29, 0x78, 0xaf, 0x51, 0xf9, 0x06, 0xe0,
	0x93, 0x5a, 0xb7, 0xc9, 0x7e, 0x83, 0x21, 0x22, 0xda, 0x35, 0x9f, 0x02, 0x6c, 0xd4, 0xfa, 0x56,
	0x49, 0x23, 0xa4, 0xf1, 0xcf, 0x1e, 0x44, 0x9a, 0xb4, 0x5f, 0xfd, 0xed, 0x40, 0xf7, 0xee, 0xc7,
	0x92, 0x30, 0x80, 0xf7, 0xc2, 0xf8, 0x45, 0x25, 0x67, 0xac, 0xbe, 0xc4, 0xc9, 0x98, 0x85, 0x3b,
	0x4c, 0x4f, 0xc8, 0x0b, 0xe8, 0xe3, 0x0e, 0x91, 0x31, 0x0b, 0x37, 0x34, 0x19, 0xb1, 0xc3, 0x6a,
	0x59, 0x18, 0x85, 0x9e, 0x9b, 0x7d, 0x12, 0xb3, 0x60, 0x4b, 0x12, 0x60, 0xfb, 0x85, 0xb0, 0x98,
	0x5b, 0xb8, 0xb0, 0xd2, 0xf5, 0xd1, 0x25, 0xd7, 0xac, 0x71, 0xf0, 0x93, 0x2b, 0xd6, 0x30, 0xe3,
	0x96, 0xe4, 0x0d, 0x8c, 0x82, 0x49, 0x21, 0x97, 0xec, 0xff, 0x39, 0x4d, 0x2e, 0xd8, 0xf1, 0x30,
	0xd9, 0x8b, 0x33, 0x18, 0x94, 0x1d, 0x25, 0x13, 0x56, 0xeb, 0x7f, 0x12, 0xb3, 0xa0, 0xd5, 0x58,
	0xcb, 0xc0, 0xe6, 0x69, 0xfb, 0x40, 0x46, 0xec, 0xd0, 0xaf, 0x24, 0x62, 0x55, 0x6b, 0xe8, 0xc9,
	0xfd, 0x00, 0xff, 0x82, 0xaf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x16, 0x07, 0x4d, 0x3d, 0x12,
	0x05, 0x00, 0x00,
}

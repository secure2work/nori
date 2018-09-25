// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package commands

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandsClient is the client API for Commands service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandsClient interface {
	PluginListCommand(ctx context.Context, in *PluginListRequest, opts ...grpc.CallOption) (*PluginListReply, error)
	PluginGetCommand(ctx context.Context, in *PluginGetRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginRemoveCommand(ctx context.Context, in *PluginRemoveRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginInstallCommand(ctx context.Context, in *PluginInstallRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginUninstallCommand(ctx context.Context, in *PluginUninstallRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	PluginUploadCommand(ctx context.Context, in *PluginUploadRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	CertsUploadCommand(ctx context.Context, in *CertsUploadRequest, opts ...grpc.CallOption) (*ErrorReply, error)
}

type commandsClient struct {
	cc *grpc.ClientConn
}

func NewCommandsClient(cc *grpc.ClientConn) CommandsClient {
	return &commandsClient{cc}
}

func (c *commandsClient) PluginListCommand(ctx context.Context, in *PluginListRequest, opts ...grpc.CallOption) (*PluginListReply, error) {
	out := new(PluginListReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/PluginListCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) PluginGetCommand(ctx context.Context, in *PluginGetRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/PluginGetCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) PluginRemoveCommand(ctx context.Context, in *PluginRemoveRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/PluginRemoveCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) PluginInstallCommand(ctx context.Context, in *PluginInstallRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/PluginInstallCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) PluginUninstallCommand(ctx context.Context, in *PluginUninstallRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/PluginUninstallCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) PluginUploadCommand(ctx context.Context, in *PluginUploadRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/PluginUploadCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) CertsUploadCommand(ctx context.Context, in *CertsUploadRequest, opts ...grpc.CallOption) (*ErrorReply, error) {
	out := new(ErrorReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/CertsUploadCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandsServer is the server API for Commands service.
type CommandsServer interface {
	PluginListCommand(context.Context, *PluginListRequest) (*PluginListReply, error)
	PluginGetCommand(context.Context, *PluginGetRequest) (*ErrorReply, error)
	PluginRemoveCommand(context.Context, *PluginRemoveRequest) (*ErrorReply, error)
	PluginInstallCommand(context.Context, *PluginInstallRequest) (*ErrorReply, error)
	PluginUninstallCommand(context.Context, *PluginUninstallRequest) (*ErrorReply, error)
	PluginUploadCommand(context.Context, *PluginUploadRequest) (*ErrorReply, error)
	CertsUploadCommand(context.Context, *CertsUploadRequest) (*ErrorReply, error)
}

func RegisterCommandsServer(s *grpc.Server, srv CommandsServer) {
	s.RegisterService(&_Commands_serviceDesc, srv)
}

func _Commands_PluginListCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).PluginListCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/PluginListCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).PluginListCommand(ctx, req.(*PluginListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_PluginGetCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).PluginGetCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/PluginGetCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).PluginGetCommand(ctx, req.(*PluginGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_PluginRemoveCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).PluginRemoveCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/PluginRemoveCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).PluginRemoveCommand(ctx, req.(*PluginRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_PluginInstallCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginInstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).PluginInstallCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/PluginInstallCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).PluginInstallCommand(ctx, req.(*PluginInstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_PluginUninstallCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginUninstallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).PluginUninstallCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/PluginUninstallCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).PluginUninstallCommand(ctx, req.(*PluginUninstallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_PluginUploadCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).PluginUploadCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/PluginUploadCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).PluginUploadCommand(ctx, req.(*PluginUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_CertsUploadCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertsUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).CertsUploadCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/CertsUploadCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).CertsUploadCommand(ctx, req.(*CertsUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Commands_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commands.Commands",
	HandlerType: (*CommandsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginListCommand",
			Handler:    _Commands_PluginListCommand_Handler,
		},
		{
			MethodName: "PluginGetCommand",
			Handler:    _Commands_PluginGetCommand_Handler,
		},
		{
			MethodName: "PluginRemoveCommand",
			Handler:    _Commands_PluginRemoveCommand_Handler,
		},
		{
			MethodName: "PluginInstallCommand",
			Handler:    _Commands_PluginInstallCommand_Handler,
		},
		{
			MethodName: "PluginUninstallCommand",
			Handler:    _Commands_PluginUninstallCommand_Handler,
		},
		{
			MethodName: "PluginUploadCommand",
			Handler:    _Commands_PluginUploadCommand_Handler,
		},
		{
			MethodName: "CertsUploadCommand",
			Handler:    _Commands_CertsUploadCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_5ea2e3fe5ae8b5fe) }

var fileDescriptor_service_5ea2e3fe5ae8b5fe = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xce, 0xcf, 0xcd, 0x4d,
	0xcc, 0x4b, 0x29, 0x96, 0x12, 0x2d, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x8b, 0x87, 0x09, 0x40, 0x14,
	0x48, 0x89, 0x24, 0xa7, 0x16, 0x95, 0x14, 0xa3, 0x89, 0x1a, 0xad, 0x65, 0xe1, 0xe2, 0x70, 0x86,
	0x0a, 0x09, 0xf9, 0x72, 0x09, 0x06, 0x80, 0xf5, 0xfa, 0x64, 0x16, 0x97, 0x40, 0x45, 0x85, 0xa4,
	0xf5, 0xe0, 0x5a, 0x10, 0x92, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x52, 0x92, 0xd8, 0x25,
	0x0b, 0x72, 0x2a, 0x95, 0x18, 0x84, 0xdc, 0xb8, 0x04, 0x20, 0x82, 0xee, 0xa9, 0x70, 0xd3, 0xa4,
	0xd0, 0x35, 0xb8, 0xa7, 0xc2, 0x0d, 0x13, 0x41, 0xc8, 0xb9, 0x16, 0x15, 0xe5, 0x17, 0xc1, 0xcc,
	0xf1, 0xe1, 0x12, 0x86, 0xa8, 0x0d, 0x4a, 0xcd, 0xcd, 0x2f, 0x4b, 0x85, 0x19, 0x25, 0x8b, 0x6e,
	0x14, 0x44, 0x9a, 0x90, 0x69, 0x7e, 0x5c, 0x22, 0x10, 0xe5, 0x9e, 0x79, 0xc5, 0x25, 0x89, 0x39,
	0x39, 0x30, 0xe3, 0xe4, 0xd0, 0x8d, 0x83, 0xca, 0x13, 0x32, 0x2f, 0x88, 0x4b, 0x0c, 0xa2, 0x3e,
	0x34, 0x2f, 0x13, 0xd5, 0x44, 0x05, 0x74, 0x13, 0xe1, 0x2a, 0x88, 0xf6, 0x71, 0x68, 0x41, 0x4e,
	0x7e, 0x62, 0x0a, 0x4e, 0x1f, 0x43, 0xa4, 0x09, 0x99, 0xe6, 0xc5, 0x25, 0xe4, 0x0c, 0x8a, 0x7b,
	0x54, 0xc3, 0x64, 0x10, 0xaa, 0x91, 0x64, 0x09, 0x98, 0x95, 0xc4, 0x06, 0x4e, 0x36, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x55, 0xfb, 0x3e, 0xcd, 0x7e, 0x02, 0x00, 0x00,
}

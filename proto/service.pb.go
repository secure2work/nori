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
	PluginMetaCommand(ctx context.Context, in *PluginMetaRequest, opts ...grpc.CallOption) (*PluginMetaReply, error)
	CertsUploadCommand(ctx context.Context, in *CertsUploadRequest, opts ...grpc.CallOption) (*ErrorReply, error)
	SendPingCommand(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error)
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

func (c *commandsClient) PluginMetaCommand(ctx context.Context, in *PluginMetaRequest, opts ...grpc.CallOption) (*PluginMetaReply, error) {
	out := new(PluginMetaReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/PluginMetaCommand", in, out, opts...)
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

func (c *commandsClient) SendPingCommand(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error) {
	out := new(PongReply)
	err := c.cc.Invoke(ctx, "/commands.Commands/SendPingCommand", in, out, opts...)
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
	PluginMetaCommand(context.Context, *PluginMetaRequest) (*PluginMetaReply, error)
	CertsUploadCommand(context.Context, *CertsUploadRequest) (*ErrorReply, error)
	SendPingCommand(context.Context, *PingRequest) (*PongReply, error)
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

func _Commands_PluginMetaCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginMetaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).PluginMetaCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/PluginMetaCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).PluginMetaCommand(ctx, req.(*PluginMetaRequest))
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

func _Commands_SendPingCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).SendPingCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/SendPingCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).SendPingCommand(ctx, req.(*PingRequest))
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
			MethodName: "PluginMetaCommand",
			Handler:    _Commands_PluginMetaCommand_Handler,
		},
		{
			MethodName: "CertsUploadCommand",
			Handler:    _Commands_CertsUploadCommand_Handler,
		},
		{
			MethodName: "SendPingCommand",
			Handler:    _Commands_SendPingCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_d62566e47f22f704) }

var fileDescriptor_service_d62566e47f22f704 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xbd, 0x51, 0x46, 0x40, 0xd4, 0xae, 0x13, 0xad, 0x7f, 0x10, 0x1f, 0x60, 0x17, 0xfa,
	0x00, 0x5e, 0x0c, 0x1d, 0xca, 0x26, 0xa3, 0xb2, 0xeb, 0x51, 0xdb, 0x43, 0x08, 0xa4, 0x49, 0x4c,
	0xb2, 0x81, 0x0f, 0xe7, 0xbb, 0x49, 0x7b, 0x9a, 0xb4, 0x0b, 0x96, 0x7a, 0xfb, 0xfd, 0xbe, 0xf3,
	0x4b, 0xc2, 0x09, 0x39, 0x36, 0xa0, 0x77, 0x2c, 0x87, 0xa9, 0xd2, 0xd2, 0xca, 0x68, 0x94, 0xcb,
	0xb2, 0xcc, 0x44, 0x61, 0x92, 0x89, 0xe2, 0x5b, 0xca, 0xc4, 0xc6, 0x05, 0x58, 0x48, 0xe2, 0x1c,
	0xb4, 0x35, 0x61, 0x7a, 0xa1, 0x98, 0xa0, 0x1b, 0x25, 0x05, 0x0d, 0xc8, 0xc3, 0xcf, 0x21, 0x19,
	0xcd, 0x9a, 0x28, 0x5a, 0x92, 0xb3, 0x55, 0x6d, 0x5d, 0x30, 0x63, 0x9b, 0x34, 0xba, 0x9a, 0xfa,
	0x91, 0x16, 0xa6, 0xf0, 0xb5, 0x05, 0x63, 0x93, 0xcb, 0xbf, 0xa1, 0xe2, 0xdf, 0xf7, 0x07, 0xd1,
	0x0b, 0x39, 0xc5, 0x70, 0x0e, 0xde, 0x96, 0x84, 0x03, 0x73, 0xf0, 0xb2, 0xb8, 0x65, 0xcf, 0x5a,
	0x4b, 0xed, 0x3c, 0x0b, 0x32, 0xc6, 0x6e, 0x0a, 0xa5, 0xdc, 0x81, 0x53, 0xdd, 0x84, 0x2a, 0xc4,
	0x43, 0xb6, 0x77, 0x12, 0x63, 0xfd, 0x55, 0x18, 0x9b, 0x71, 0xee, 0x74, 0xb7, 0xa1, 0xae, 0xe1,
	0x43, 0xbe, 0x94, 0x9c, 0x63, 0x7f, 0x2d, 0xd8, 0xbe, 0xf1, 0x2e, 0x34, 0xfa, 0xc6, 0xbf, 0x5f,
	0xbc, 0x56, 0x5c, 0x66, 0x45, 0xef, 0x8b, 0x11, 0x0f, 0xd9, 0xfc, 0x5a, 0x97, 0x60, 0xb3, 0xde,
	0xb5, 0x56, 0xb0, 0x77, 0xad, 0x08, 0x51, 0xf7, 0x46, 0xa2, 0x59, 0xf5, 0xc9, 0xf6, 0xef, 0x76,
	0xdd, 0x8e, 0x74, 0xe8, 0xd0, 0xd5, 0x9e, 0xc8, 0xc9, 0x07, 0x88, 0x62, 0xc5, 0x04, 0x75, 0xa2,
	0x49, 0xe7, 0x6c, 0x26, 0xa8, 0x33, 0x8c, 0x3b, 0xb1, 0xac, 0xe2, 0x5a, 0xf0, 0x79, 0x54, 0x7f,
	0xe3, 0xc7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xec, 0x81, 0x13, 0x28, 0x03, 0x00, 0x00,
}

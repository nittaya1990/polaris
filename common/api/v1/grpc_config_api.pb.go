// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_config_api.proto

package v1

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

// PolarisConfigGRPCClient is the client API for PolarisConfigGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisConfigGRPCClient interface {
	// 拉取配置
	GetConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error)
	// 订阅配置变更
	WatchConfigFiles(ctx context.Context, in *ClientWatchConfigFileRequest, opts ...grpc.CallOption) (*ConfigClientResponse, error)
}

type polarisConfigGRPCClient struct {
	cc *grpc.ClientConn
}

func NewPolarisConfigGRPCClient(cc *grpc.ClientConn) PolarisConfigGRPCClient {
	return &polarisConfigGRPCClient{cc}
}

func (c *polarisConfigGRPCClient) GetConfigFile(ctx context.Context, in *ClientConfigFileInfo, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/GetConfigFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polarisConfigGRPCClient) WatchConfigFiles(ctx context.Context, in *ClientWatchConfigFileRequest, opts ...grpc.CallOption) (*ConfigClientResponse, error) {
	out := new(ConfigClientResponse)
	err := c.cc.Invoke(ctx, "/v1.PolarisConfigGRPC/WatchConfigFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolarisConfigGRPCServer is the server API for PolarisConfigGRPC service.
type PolarisConfigGRPCServer interface {
	// 拉取配置
	GetConfigFile(context.Context, *ClientConfigFileInfo) (*ConfigClientResponse, error)
	// 订阅配置变更
	WatchConfigFiles(context.Context, *ClientWatchConfigFileRequest) (*ConfigClientResponse, error)
}

func RegisterPolarisConfigGRPCServer(s *grpc.Server, srv PolarisConfigGRPCServer) {
	s.RegisterService(&_PolarisConfigGRPC_serviceDesc, srv)
}

func _PolarisConfigGRPC_GetConfigFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigFileInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).GetConfigFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/GetConfigFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).GetConfigFile(ctx, req.(*ClientConfigFileInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolarisConfigGRPC_WatchConfigFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientWatchConfigFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolarisConfigGRPCServer).WatchConfigFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PolarisConfigGRPC/WatchConfigFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolarisConfigGRPCServer).WatchConfigFiles(ctx, req.(*ClientWatchConfigFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolarisConfigGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PolarisConfigGRPC",
	HandlerType: (*PolarisConfigGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigFile",
			Handler:    _PolarisConfigGRPC_GetConfigFile_Handler,
		},
		{
			MethodName: "WatchConfigFiles",
			Handler:    _PolarisConfigGRPC_WatchConfigFiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_config_api.proto",
}

func init() {
	proto.RegisterFile("grpc_config_api.proto", fileDescriptor_grpc_config_api_57c0c78d97511b1f)
}

var fileDescriptor_grpc_config_api_57c0c78d97511b1f = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0x2f, 0x2a, 0x48,
	0x8e, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x8f, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x12, 0x84, 0x8a, 0xa6, 0x65, 0xe6, 0xa4, 0x42, 0x84, 0xa5,
	0xa4, 0x90, 0x84, 0xe2, 0x8b, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0xa1, 0x72, 0x46, 0x6b, 0x18,
	0xb9, 0x04, 0x03, 0xf2, 0x73, 0x12, 0x8b, 0x32, 0x8b, 0x9d, 0xc1, 0xaa, 0xdc, 0x83, 0x02, 0x9c,
	0x85, 0x5c, 0xb9, 0x78, 0xdd, 0x53, 0x4b, 0x20, 0x02, 0x6e, 0x99, 0x39, 0xa9, 0x42, 0x12, 0x7a,
	0x65, 0x86, 0x7a, 0xce, 0x39, 0x99, 0xa9, 0x79, 0x48, 0xa2, 0x9e, 0x79, 0x69, 0xf9, 0x52, 0x10,
	0x19, 0xb0, 0x18, 0x44, 0x3e, 0x08, 0x6a, 0x81, 0x12, 0x83, 0x50, 0x00, 0x97, 0x40, 0x78, 0x62,
	0x49, 0x72, 0x06, 0x42, 0x4b, 0xb1, 0x90, 0x02, 0xc2, 0x24, 0x34, 0xb9, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x7c, 0x26, 0x3a, 0xb1, 0x44, 0x31, 0x95, 0x19, 0x26, 0xb1, 0x81, 0xdd, 0x6e,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe6, 0xb8, 0xa9, 0xf9, 0x07, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: collect.stock.proto

package api

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("collect.stock.proto", fileDescriptor_12a3b0403e171ae9) }

var fileDescriptor_12a3b0403e171ae9 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4a, 0xc4, 0x40,
	0x10, 0x86, 0x2f, 0x2a, 0x16, 0xc1, 0x2a, 0x07, 0x82, 0x51, 0x52, 0x5c, 0x29, 0xba, 0xcb, 0x45,
	0xb0, 0x57, 0xd1, 0x46, 0x04, 0x21, 0x56, 0x36, 0xc7, 0x66, 0xdd, 0xdb, 0x1b, 0x4c, 0x76, 0x96,
	0xdb, 0x49, 0x20, 0xad, 0xaf, 0x60, 0xe3, 0xeb, 0xd8, 0x5d, 0x29, 0xf8, 0x02, 0x1a, 0x7d, 0x10,
	0xc9, 0x26, 0x72, 0x95, 0x85, 0xdd, 0xce, 0xfc, 0x3b, 0xdf, 0xff, 0xf3, 0x87, 0x63, 0x89, 0x45,
	0xa1, 0x24, 0x31, 0x47, 0x28, 0x1f, 0x99, 0x5d, 0x22, 0x61, 0x94, 0x6a, 0xa0, 0x45, 0x95, 0x33,
	0x89, 0x25, 0xcb, 0xaa, 0x6b, 0x01, 0x57, 0x0a, 0xd8, 0x5c, 0x18, 0x12, 0x8e, 0x40, 0xce, 0x16,
	0xc2, 0x5a, 0x30, 0xca, 0x39, 0xe6, 0xd4, 0xb2, 0x06, 0xa9, 0x58, 0x3d, 0x8d, 0x8f, 0xd7, 0x37,
	0x5c, 0xa3, 0x46, 0xee, 0x51, 0x79, 0x35, 0xf7, 0x93, 0x1f, 0xfc, 0xab, 0xb7, 0x88, 0xf7, 0x35,
	0xa2, 0x2e, 0xd4, 0xfa, 0x97, 0x2a, 0x2d, 0x35, 0x83, 0x78, 0x30, 0x88, 0xc2, 0x02, 0x17, 0xc6,
	0x20, 0x09, 0x02, 0x34, 0xae, 0x57, 0xd3, 0xd7, 0x20, 0xdc, 0xc9, 0xba, 0xb4, 0x17, 0x7d, 0xf4,
	0xe8, 0x34, 0xdc, 0xba, 0x05, 0xa3, 0xa3, 0x5d, 0xd6, 0xdf, 0xb1, 0x5f, 0x28, 0xbb, 0xec, 0xa0,
	0xf1, 0x1f, 0xfb, 0xa8, 0x09, 0xc7, 0x59, 0x63, 0xa4, 0x67, 0xdd, 0x80, 0x39, 0xa3, 0x3b, 0x7c,
	0x10, 0xcd, 0x7f, 0x31, 0x93, 0xf4, 0xe9, 0xfd, 0xfb, 0x79, 0xe3, 0x68, 0x72, 0xe8, 0xf3, 0xd6,
	0x53, 0xee, 0xab, 0x9c, 0x0d, 0xc5, 0x72, 0xd7, 0x18, 0x39, 0xac, 0x4a, 0x30, 0x9c, 0x3a, 0x8f,
	0xf3, 0xbd, 0xd5, 0x67, 0x32, 0x5a, 0xb5, 0x49, 0xf0, 0xd6, 0x26, 0xc1, 0x47, 0x9b, 0x04, 0x2f,
	0x5f, 0xc9, 0xe8, 0x7e, 0x53, 0x58, 0xc8, 0xb7, 0x3d, 0xfe, 0xe4, 0x27, 0x00, 0x00, 0xff, 0xff,
	0xa7, 0xcf, 0x7f, 0xf9, 0x9a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StockCollectClient is the client API for StockCollect service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockCollectClient interface {
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 作业同步今天深沪证券交易所股票-分钟数据
	SyncStockMinAtToday(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type stockCollectClient struct {
	cc *grpc.ClientConn
}

func NewStockCollectClient(cc *grpc.ClientConn) StockCollectClient {
	return &stockCollectClient{cc}
}

func (c *stockCollectClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.com.SuKaiFei.fantastic_happiness.service.v1.StockCollect/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockCollectClient) SyncStockMinAtToday(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/github.com.SuKaiFei.fantastic_happiness.service.v1.StockCollect/SyncStockMinAtToday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockCollectServer is the server API for StockCollect service.
type StockCollectServer interface {
	Ping(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// 作业同步今天深沪证券交易所股票-分钟数据
	SyncStockMinAtToday(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
}

// UnimplementedStockCollectServer can be embedded to have forward compatible implementations.
type UnimplementedStockCollectServer struct {
}

func (*UnimplementedStockCollectServer) Ping(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedStockCollectServer) SyncStockMinAtToday(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncStockMinAtToday not implemented")
}

func RegisterStockCollectServer(s *grpc.Server, srv StockCollectServer) {
	s.RegisterService(&_StockCollect_serviceDesc, srv)
}

func _StockCollect_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockCollectServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.SuKaiFei.fantastic_happiness.service.v1.StockCollect/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockCollectServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockCollect_SyncStockMinAtToday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockCollectServer).SyncStockMinAtToday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.SuKaiFei.fantastic_happiness.service.v1.StockCollect/SyncStockMinAtToday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockCollectServer).SyncStockMinAtToday(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _StockCollect_serviceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.SuKaiFei.fantastic_happiness.service.v1.StockCollect",
	HandlerType: (*StockCollectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _StockCollect_Ping_Handler,
		},
		{
			MethodName: "SyncStockMinAtToday",
			Handler:    _StockCollect_SyncStockMinAtToday_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collect.stock.proto",
}

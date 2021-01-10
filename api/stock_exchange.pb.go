// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stock_exchange.proto

package api

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type StockExchange int32

const (
	StockExchange__ StockExchange = 0
	// 上海证券交易所
	StockExchange_SSE StockExchange = 1
	// 深圳证券交易所
	StockExchange_SZSE StockExchange = 2
)

var StockExchange_name = map[int32]string{
	0: "_",
	1: "SSE",
	2: "SZSE",
}

var StockExchange_value = map[string]int32{
	"_":    0,
	"SSE":  1,
	"SZSE": 2,
}

func (x StockExchange) String() string {
	return proto.EnumName(StockExchange_name, int32(x))
}

func (StockExchange) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_23483f1ca87257ff, []int{0}
}

func init() {
	proto.RegisterEnum("github.com.SuKaiFei.fantastic_happiness.service.v1.StockExchange", StockExchange_name, StockExchange_value)
}

func init() { proto.RegisterFile("stock_exchange.proto", fileDescriptor_23483f1ca87257ff) }

var fileDescriptor_23483f1ca87257ff = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0xc9, 0x4f,
	0xce, 0x8e, 0x4f, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x32, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0b, 0x2e, 0xf5,
	0x4e, 0xcc, 0x74, 0x4b, 0xcd, 0xd4, 0x4b, 0x4b, 0xcc, 0x2b, 0x49, 0x2c, 0x2e, 0xc9, 0x4c, 0x8e,
	0xcf, 0x48, 0x2c, 0x28, 0xc8, 0xcc, 0x4b, 0x2d, 0x2e, 0xd6, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0xd5, 0x2b, 0x33, 0x94, 0xd2, 0x45, 0xe8, 0xd1, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x1b,
	0x95, 0x54, 0x9a, 0x06, 0xe6, 0x81, 0x39, 0x60, 0x16, 0xc4, 0x0a, 0x2d, 0x4d, 0x2e, 0xde, 0x60,
	0x90, 0xd5, 0xae, 0x50, 0x9b, 0x85, 0x58, 0xb9, 0x18, 0xe3, 0x05, 0x18, 0x84, 0xd8, 0xb9, 0x98,
	0x83, 0x83, 0x5d, 0x05, 0x18, 0x85, 0x38, 0xb8, 0x58, 0x82, 0xa3, 0x82, 0x5d, 0x05, 0x98, 0x9c,
	0x24, 0x4f, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0x62, 0x4e, 0x2c, 0xc8, 0x4c, 0x62, 0x03, 0x1b, 0x66,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x13, 0x77, 0x46, 0x3f, 0xc7, 0x00, 0x00, 0x00,
}
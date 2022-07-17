// Code generated by protoc-gen-go. DO NOT EDIT.
// source: location.proto

package blockstore

import (
	fmt "fmt"
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

type BlockLocation struct {
	FileChunkNum         uint64   `protobuf:"varint,1,opt,name=file_chunk_num,json=fileChunkNum,proto3" json:"file_chunk_num,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Length               int64    `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockLocation) Reset()         { *m = BlockLocation{} }
func (m *BlockLocation) String() string { return proto.CompactTextString(m) }
func (*BlockLocation) ProtoMessage()    {}
func (*BlockLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f0f35158dcf9f2c, []int{0}
}

func (m *BlockLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockLocation.Unmarshal(m, b)
}
func (m *BlockLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockLocation.Marshal(b, m, deterministic)
}
func (m *BlockLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockLocation.Merge(m, src)
}
func (m *BlockLocation) XXX_Size() int {
	return xxx_messageInfo_BlockLocation.Size(m)
}
func (m *BlockLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockLocation.DiscardUnknown(m)
}

var xxx_messageInfo_BlockLocation proto.InternalMessageInfo

func (m *BlockLocation) GetFileChunkNum() uint64 {
	if m != nil {
		return m.FileChunkNum
	}
	return 0
}

func (m *BlockLocation) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *BlockLocation) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockLocation)(nil), "blockstore.BlockLocation")
}

func init() { proto.RegisterFile("location.proto", fileDescriptor_4f0f35158dcf9f2c) }

var fileDescriptor_4f0f35158dcf9f2c = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xce, 0xb1, 0x6b, 0x84, 0x30,
	0x14, 0xc7, 0x71, 0xac, 0xc5, 0x21, 0xb4, 0x0e, 0x0e, 0xc5, 0x51, 0x4a, 0x07, 0x17, 0xcd, 0xd0,
	0xb5, 0x74, 0xb0, 0x6b, 0xb9, 0xc1, 0xf1, 0x16, 0x31, 0xb9, 0xa7, 0x09, 0xc6, 0x3c, 0x79, 0x49,
	0x0e, 0xee, 0xbf, 0x3f, 0x72, 0x0a, 0x37, 0x7e, 0x3f, 0xfc, 0x86, 0x1f, 0xcb, 0x0d, 0xca, 0xd1,
	0x6b, 0xb4, 0xed, 0x46, 0xe8, 0xb1, 0x60, 0xc2, 0xa0, 0x5c, 0x9c, 0x47, 0x82, 0x4f, 0x60, 0xef,
	0x5d, 0xac, 0xff, 0x63, 0x52, 0x7c, 0xb1, 0x7c, 0xd2, 0x06, 0x06, 0xa9, 0x82, 0x5d, 0x06, 0x1b,
	0xd6, 0x32, 0xa9, 0x92, 0xfa, 0xb5, 0x7f, 0x8b, 0xfa, 0x17, 0xf1, 0x14, 0xd6, 0xe2, 0x83, 0x65,
	0x38, 0x4d, 0x0e, 0x7c, 0xf9, 0x52, 0x25, 0x75, 0xda, 0x1f, 0x15, 0xdd, 0x80, 0x9d, 0xbd, 0x2a,
	0xd3, 0xdd, 0xf7, 0xea, 0x7e, 0xcf, 0x3f, 0xb3, 0xf6, 0x2a, 0x88, 0x56, 0xe2, 0xca, 0xd5, 0x6d,
	0x03, 0x32, 0x70, 0x99, 0x81, 0x1a, 0x33, 0x0a, 0xc7, 0x91, 0x34, 0xda, 0xc6, 0x01, 0x5d, 0x81,
	0xb8, 0xb6, 0x1e, 0xc8, 0x8e, 0x86, 0x3f, 0x6f, 0x8a, 0xec, 0xf1, 0xfc, 0xfb, 0x1e, 0x00, 0x00,
	0xff, 0xff, 0x7b, 0xfb, 0x11, 0xba, 0xcb, 0x00, 0x00, 0x00,
}

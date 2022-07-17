// Code generated by protoc-gen-go. DO NOT EDIT.
// source: block_txs_id.proto

package blockstore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	types "github.com/hyperledger-labs/orion-server/pkg/types"
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

type BlockTxIDs struct {
	TxIds                []string `protobuf:"bytes,1,rep,name=tx_ids,json=txIds,proto3" json:"tx_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockTxIDs) Reset()         { *m = BlockTxIDs{} }
func (m *BlockTxIDs) String() string { return proto.CompactTextString(m) }
func (*BlockTxIDs) ProtoMessage()    {}
func (*BlockTxIDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ac3dcb8cc34951, []int{0}
}

func (m *BlockTxIDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTxIDs.Unmarshal(m, b)
}
func (m *BlockTxIDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTxIDs.Marshal(b, m, deterministic)
}
func (m *BlockTxIDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTxIDs.Merge(m, src)
}
func (m *BlockTxIDs) XXX_Size() int {
	return xxx_messageInfo_BlockTxIDs.Size(m)
}
func (m *BlockTxIDs) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTxIDs.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTxIDs proto.InternalMessageInfo

func (m *BlockTxIDs) GetTxIds() []string {
	if m != nil {
		return m.TxIds
	}
	return nil
}

type TxInfo struct {
	BlockNumber          uint64                `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	TxIndex              uint64                `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	Validation           *types.ValidationInfo `protobuf:"bytes,3,opt,name=validation,proto3" json:"validation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *TxInfo) Reset()         { *m = TxInfo{} }
func (m *TxInfo) String() string { return proto.CompactTextString(m) }
func (*TxInfo) ProtoMessage()    {}
func (*TxInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1ac3dcb8cc34951, []int{1}
}

func (m *TxInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxInfo.Unmarshal(m, b)
}
func (m *TxInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxInfo.Marshal(b, m, deterministic)
}
func (m *TxInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxInfo.Merge(m, src)
}
func (m *TxInfo) XXX_Size() int {
	return xxx_messageInfo_TxInfo.Size(m)
}
func (m *TxInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TxInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TxInfo proto.InternalMessageInfo

func (m *TxInfo) GetBlockNumber() uint64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *TxInfo) GetTxIndex() uint64 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *TxInfo) GetValidation() *types.ValidationInfo {
	if m != nil {
		return m.Validation
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockTxIDs)(nil), "blockstore.BlockTxIDs")
	proto.RegisterType((*TxInfo)(nil), "blockstore.TxInfo")
}

func init() { proto.RegisterFile("block_txs_id.proto", fileDescriptor_c1ac3dcb8cc34951) }

var fileDescriptor_c1ac3dcb8cc34951 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x0a, 0x05, 0xae, 0x4c, 0x96, 0x2a, 0x05, 0x58, 0x42, 0x59, 0xb2, 0x34, 0x91,
	0x40, 0x6c, 0x88, 0xa1, 0x62, 0xc9, 0xc2, 0x10, 0x55, 0x0c, 0x2c, 0x91, 0x13, 0x1f, 0xad, 0x45,
	0x6a, 0x47, 0xe7, 0x6b, 0xe5, 0x0e, 0xfc, 0x77, 0x64, 0x47, 0x02, 0x46, 0x7f, 0xdf, 0xd3, 0xf3,
	0xdd, 0x81, 0x68, 0x7b, 0xdb, 0x7d, 0x35, 0xec, 0x5d, 0xa3, 0x55, 0x31, 0x90, 0x65, 0x2b, 0x20,
	0x32, 0xc7, 0x96, 0xf0, 0xe6, 0x76, 0xf4, 0xd2, 0xa8, 0x86, 0x49, 0x1a, 0x27, 0x3b, 0xd6, 0xd6,
	0x8c, 0xc1, 0xc5, 0x3d, 0xc0, 0x2a, 0xe8, 0xb5, 0xaf, 0x5e, 0x9d, 0x98, 0xc3, 0x94, 0x7d, 0xa3,
	0x95, 0x4b, 0x93, 0x6c, 0x92, 0x5f, 0xd6, 0x67, 0xec, 0x2b, 0xe5, 0x16, 0xdf, 0x30, 0x5d, 0xfb,
	0xca, 0x7c, 0x5a, 0x71, 0x07, 0x57, 0x63, 0x9b, 0xd9, 0xef, 0x5a, 0xa4, 0x34, 0xc9, 0x92, 0xfc,
	0xb4, 0x9e, 0x45, 0xf6, 0x16, 0x91, 0xb8, 0x86, 0x8b, 0xd0, 0x61, 0x14, 0xfa, 0xf4, 0x24, 0xea,
	0x73, 0xf6, 0x55, 0x78, 0x8a, 0x27, 0x80, 0x83, 0xec, 0xb5, 0x92, 0x61, 0x80, 0x74, 0x92, 0x25,
	0xf9, 0xec, 0x61, 0x5e, 0xf0, 0x71, 0x40, 0x57, 0xbc, 0xff, 0x8a, 0xf0, 0x51, 0xfd, 0x2f, 0xb8,
	0x7a, 0xf9, 0x78, 0xde, 0x68, 0xde, 0xee, 0xdb, 0xa2, 0xb3, 0xbb, 0x72, 0x7b, 0x1c, 0x90, 0x7a,
	0x54, 0x1b, 0xa4, 0x65, 0x2f, 0x5b, 0x57, 0x5a, 0xd2, 0xd6, 0x2c, 0x1d, 0xd2, 0x01, 0xa9, 0xd4,
	0x86, 0x91, 0x8c, 0xec, 0xcb, 0xbf, 0x03, 0xb4, 0xd3, 0xb8, 0xea, 0xe3, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0x84, 0xc1, 0xb6, 0x29, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package go_raft_chunking

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

type ChunkInfo struct {
	// OpNum is the ID of the op, used to ensure values are applied to the
	// right operation
	OpNum uint64 `protobuf:"varint,1,opt,name=op_num,json=opNum,proto3" json:"op_num,omitempty"`
	// SequenceNum is the current number of the ops; when applying we should
	// see this start at zero and increment by one without skips
	SequenceNum uint32 `protobuf:"varint,2,opt,name=sequence_num,json=sequenceNum,proto3" json:"sequence_num,omitempty"`
	// NumChunks is used to check whether all chunks have been received and
	// reconstruction should be attempted
	NumChunks            uint32   `protobuf:"varint,3,opt,name=num_chunks,json=numChunks,proto3" json:"num_chunks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkInfo) Reset()         { *m = ChunkInfo{} }
func (m *ChunkInfo) String() string { return proto.CompactTextString(m) }
func (*ChunkInfo) ProtoMessage()    {}
func (*ChunkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *ChunkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkInfo.Unmarshal(m, b)
}
func (m *ChunkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkInfo.Marshal(b, m, deterministic)
}
func (m *ChunkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkInfo.Merge(m, src)
}
func (m *ChunkInfo) XXX_Size() int {
	return xxx_messageInfo_ChunkInfo.Size(m)
}
func (m *ChunkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkInfo proto.InternalMessageInfo

func (m *ChunkInfo) GetOpNum() uint64 {
	if m != nil {
		return m.OpNum
	}
	return 0
}

func (m *ChunkInfo) GetSequenceNum() uint32 {
	if m != nil {
		return m.SequenceNum
	}
	return 0
}

func (m *ChunkInfo) GetNumChunks() uint32 {
	if m != nil {
		return m.NumChunks
	}
	return 0
}

func init() {
	proto.RegisterType((*ChunkInfo)(nil), "chunking.ChunkInfo")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355) }

var fileDescriptor_d938547f84707355 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0xce, 0x28, 0xcd, 0xcb, 0xce, 0xcc,
	0x4b, 0x57, 0x4a, 0xe1, 0xe2, 0x74, 0x06, 0xb1, 0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0x44, 0xb9, 0xd8,
	0xf2, 0x0b, 0xe2, 0xf3, 0x4a, 0x73, 0x25, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x58, 0xf3, 0x0b,
	0xfc, 0x4a, 0x73, 0x85, 0x14, 0xb9, 0x78, 0x8a, 0x53, 0x0b, 0x4b, 0x53, 0xf3, 0x92, 0x53, 0xc1,
	0x92, 0x4c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0xdc, 0x30, 0x31, 0x90, 0x12, 0x59, 0x2e, 0xae, 0xbc,
	0xd2, 0xdc, 0x78, 0xb0, 0xb1, 0xc5, 0x12, 0xcc, 0x60, 0x05, 0x9c, 0x79, 0xa5, 0xb9, 0x60, 0xb3,
	0x8b, 0x9d, 0xd4, 0xa3, 0x54, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x33, 0x12, 0x8b, 0x33, 0x32, 0x93, 0xf3, 0x8b, 0x0a, 0xf4, 0xd3, 0xf3, 0x75, 0x8b, 0x12, 0xd3,
	0x4a, 0x74, 0x61, 0xce, 0x49, 0x62, 0x03, 0xbb, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0xe0, 0xba, 0xb4, 0xae, 0x00, 0x00, 0x00,
}
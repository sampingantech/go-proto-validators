// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/uuid.proto

package validator_examples

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/sampingantech/go-proto-validators"
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

type UUIDMsg struct {
	// user_id must be a valid version 4 UUID.
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUIDMsg) Reset()         { *m = UUIDMsg{} }
func (m *UUIDMsg) String() string { return proto.CompactTextString(m) }
func (*UUIDMsg) ProtoMessage()    {}
func (*UUIDMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_0029f00507e892b3, []int{0}
}

func (m *UUIDMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUIDMsg.Unmarshal(m, b)
}
func (m *UUIDMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUIDMsg.Marshal(b, m, deterministic)
}
func (m *UUIDMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUIDMsg.Merge(m, src)
}
func (m *UUIDMsg) XXX_Size() int {
	return xxx_messageInfo_UUIDMsg.Size(m)
}
func (m *UUIDMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_UUIDMsg.DiscardUnknown(m)
}

var xxx_messageInfo_UUIDMsg proto.InternalMessageInfo

func (m *UUIDMsg) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDMsg)(nil), "validator.examples.UUIDMsg")
}

func init() { proto.RegisterFile("examples/uuid.proto", fileDescriptor_0029f00507e892b3) }

var fileDescriptor_0029f00507e892b3 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x2d, 0xcd, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x2a, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x83, 0x49, 0x4b, 0xd9, 0xa4,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x17, 0x27, 0xe6, 0x16, 0x64, 0xe6,
	0xa5, 0x27, 0xe6, 0x95, 0xa4, 0x26, 0x67, 0xe8, 0xa7, 0xe7, 0xeb, 0x82, 0xb5, 0xe9, 0xc2, 0x75,
	0x15, 0xeb, 0x23, 0x0c, 0x00, 0x4b, 0x29, 0xe9, 0x72, 0xb1, 0x87, 0x86, 0x7a, 0xba, 0xf8, 0x16,
	0xa7, 0x0b, 0x29, 0x71, 0xb1, 0x97, 0x16, 0xa7, 0x16, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x3a, 0x71, 0x3e, 0xba, 0x2f, 0xcf, 0x1a, 0xc1, 0x38, 0x81, 0x91, 0x25, 0x88, 0x0d,
	0x24, 0xe3, 0x99, 0x92, 0xc4, 0x06, 0xd6, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xd8,
	0x81, 0x76, 0x9e, 0x00, 0x00, 0x00,
}

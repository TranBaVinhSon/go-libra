// +build !js
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language_storage.proto

package pbtypes

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

/// The unique identifier for a module on the chain.
type ModuleId struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModuleId) Reset()         { *m = ModuleId{} }
func (m *ModuleId) String() string { return proto.CompactTextString(m) }
func (*ModuleId) ProtoMessage()    {}
func (*ModuleId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cb47e965437a843, []int{0}
}

func (m *ModuleId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModuleId.Unmarshal(m, b)
}
func (m *ModuleId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModuleId.Marshal(b, m, deterministic)
}
func (m *ModuleId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleId.Merge(m, src)
}
func (m *ModuleId) XXX_Size() int {
	return xxx_messageInfo_ModuleId.Size(m)
}
func (m *ModuleId) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleId.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleId proto.InternalMessageInfo

func (m *ModuleId) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ModuleId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ModuleId)(nil), "types.ModuleId")
}

func init() { proto.RegisterFile("language_storage.proto", fileDescriptor_2cb47e965437a843) }

var fileDescriptor_2cb47e965437a843 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x49, 0xcc, 0x4b,
	0x2f, 0x4d, 0x4c, 0x4f, 0x8d, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0xb2, 0xe0, 0xe2, 0xf0, 0xcd, 0x4f,
	0x29, 0xcd, 0x49, 0xf5, 0x4c, 0x11, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e,
	0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x82, 0x71, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73,
	0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x27, 0xbd, 0x28, 0x9d, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x92, 0x8c, 0x54, 0x73, 0x23, 0x4b, 0xfd,
	0xf4, 0x7c, 0xdd, 0x9c, 0xcc, 0xa4, 0xa2, 0x44, 0xfd, 0xf4, 0xd4, 0xbc, 0xd4, 0xa2, 0xc4, 0x92,
	0xd4, 0x14, 0xfd, 0x82, 0x24, 0xb0, 0x4d, 0x49, 0x6c, 0x60, 0x7b, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0x81, 0x58, 0x63, 0x91, 0x00, 0x00, 0x00,
}

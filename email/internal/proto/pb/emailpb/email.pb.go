// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email.proto

package emailpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type TwoFAPayload struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Role                 string   `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TwoFAPayload) Reset()         { *m = TwoFAPayload{} }
func (m *TwoFAPayload) String() string { return proto.CompactTextString(m) }
func (*TwoFAPayload) ProtoMessage()    {}
func (*TwoFAPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_6175298cb4ed6faa, []int{0}
}

func (m *TwoFAPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TwoFAPayload.Unmarshal(m, b)
}
func (m *TwoFAPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TwoFAPayload.Marshal(b, m, deterministic)
}
func (m *TwoFAPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwoFAPayload.Merge(m, src)
}
func (m *TwoFAPayload) XXX_Size() int {
	return xxx_messageInfo_TwoFAPayload.Size(m)
}
func (m *TwoFAPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_TwoFAPayload.DiscardUnknown(m)
}

var xxx_messageInfo_TwoFAPayload proto.InternalMessageInfo

func (m *TwoFAPayload) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TwoFAPayload) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *TwoFAPayload) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *TwoFAPayload) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*TwoFAPayload)(nil), "email.TwoFAPayload")
}

func init() {
	proto.RegisterFile("email.proto", fileDescriptor_6175298cb4ed6faa)
}

var fileDescriptor_6175298cb4ed6faa = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcd, 0x4d, 0xcc,
	0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0xa4, 0xd3, 0xf3, 0xf3,
	0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x82, 0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x10,
	0x35, 0x4a, 0x99, 0x5c, 0x3c, 0x21, 0xe5, 0xf9, 0x6e, 0x8e, 0x01, 0x89, 0x95, 0x39, 0xf9, 0x89,
	0x29, 0x42, 0xe2, 0x5c, 0xec, 0xa5, 0xc5, 0xa9, 0x45, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x6c, 0x20, 0xae, 0x67, 0x8a, 0x90, 0x08, 0x17, 0xc4, 0x38, 0x09, 0x26, 0xb0,
	0x30, 0x84, 0x23, 0x24, 0xc4, 0xc5, 0x52, 0x94, 0x9f, 0x93, 0x2a, 0xc1, 0x0c, 0x16, 0x04, 0xb3,
	0x85, 0xc4, 0xb8, 0xd8, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x25, 0x58, 0x20, 0x26, 0x40, 0x78,
	0x46, 0x3e, 0x5c, 0x3c, 0xae, 0x20, 0x4d, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x36,
	0x5c, 0xbc, 0xc1, 0xa9, 0x79, 0x29, 0x60, 0xeb, 0x9d, 0xf3, 0x53, 0x52, 0x85, 0x84, 0xf5, 0x20,
	0xae, 0x47, 0x76, 0x90, 0x94, 0x98, 0x1e, 0xc4, 0xf9, 0x7a, 0x30, 0xe7, 0xeb, 0xb9, 0x82, 0x9c,
	0xef, 0xc4, 0x17, 0xc5, 0xa3, 0xa7, 0x5f, 0x90, 0xa4, 0x0f, 0xd6, 0x52, 0x90, 0x94, 0xc4, 0x06,
	0x96, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x63, 0x65, 0xd2, 0xc3, 0x02, 0x01, 0x00, 0x00,
}
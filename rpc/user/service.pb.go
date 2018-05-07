// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/user/service.proto

package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_a26f4ce6fbf8bd3f, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Users struct {
	User                 []*User  `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_a26f4ce6fbf8bd3f, []int{1}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_a26f4ce6fbf8bd3f, []int{2}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "resonate.toyapi.user.User")
	proto.RegisterType((*Users)(nil), "resonate.toyapi.user.Users")
	proto.RegisterType((*Empty)(nil), "resonate.toyapi.user.Empty")
}

func init() { proto.RegisterFile("rpc/user/service.proto", fileDescriptor_service_a26f4ce6fbf8bd3f) }

var fileDescriptor_service_a26f4ce6fbf8bd3f = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0x03, 0x41,
	0x10, 0x85, 0xb9, 0xcb, 0x5d, 0x2e, 0x8e, 0x60, 0xb1, 0x04, 0x59, 0x36, 0x4d, 0xb8, 0x2a, 0xd5,
	0x06, 0x62, 0x61, 0xaf, 0x11, 0xfb, 0xa0, 0x8d, 0xdd, 0x26, 0x3b, 0xc5, 0x82, 0x97, 0x5d, 0x66,
	0x56, 0xe1, 0x7e, 0x88, 0xff, 0x57, 0x76, 0x82, 0x56, 0x49, 0xca, 0x37, 0xef, 0xf1, 0xde, 0xc7,
	0xc0, 0x3d, 0xa5, 0xc3, 0xfa, 0x8b, 0x91, 0xd6, 0x8c, 0xf4, 0x1d, 0x0e, 0x68, 0x13, 0xc5, 0x1c,
	0xd5, 0x9c, 0x90, 0xe3, 0xd1, 0x65, 0xb4, 0x39, 0x8e, 0x2e, 0x05, 0x5b, 0x32, 0xfd, 0x1e, 0x9a,
	0x77, 0x46, 0x52, 0x77, 0x50, 0x07, 0xaf, 0xab, 0x65, 0xb5, 0xba, 0xd9, 0xd5, 0xc1, 0x2b, 0x03,
	0xb3, 0xe2, 0x1f, 0xdd, 0x80, 0xba, 0x96, 0xeb, 0xbf, 0x56, 0x73, 0x68, 0x71, 0x70, 0xe1, 0x53,
	0x4f, 0xc4, 0x38, 0x09, 0xa5, 0xa1, 0x73, 0xde, 0x13, 0x32, 0xeb, 0x46, 0xee, 0x7f, 0xb2, 0x7f,
	0x84, 0xb6, 0x6c, 0xb0, 0xb2, 0xd0, 0x94, 0x12, 0x5d, 0x2d, 0x27, 0xab, 0xdb, 0x8d, 0xb1, 0xe7,
	0x88, 0x6c, 0x89, 0xee, 0x24, 0xd7, 0x77, 0xd0, 0xbe, 0x0c, 0x29, 0x8f, 0x9b, 0x9f, 0x0a, 0xba,
	0xb7, 0x38, 0x0a, 0xe9, 0x16, 0xe0, 0x99, 0xd0, 0x65, 0x14, 0x75, 0xa5, 0xc4, 0x5c, 0xf1, 0xd4,
	0x16, 0x66, 0xaf, 0x98, 0x4f, 0x58, 0x8b, 0xf3, 0x39, 0x99, 0x36, 0x8b, 0xcb, 0x25, 0xfc, 0x34,
	0xfd, 0x10, 0xd0, 0xfd, 0x54, 0x5e, 0xfc, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xbc, 0x2a,
	0xfa, 0x7c, 0x01, 0x00, 0x00,
}

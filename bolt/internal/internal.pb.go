// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Dial
	User
*/
package internal

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Dial struct {
	ID               *int64   `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	UserID           *int64   `protobuf:"varint,2,opt,name=UserID" json:"UserID,omitempty"`
	Name             *string  `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Level            *float64 `protobuf:"fixed64,4,opt,name=Level" json:"Level,omitempty"`
	ModTime          *int64   `protobuf:"varint,5,opt,name=ModTime" json:"ModTime,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Dial) Reset()                    { *m = Dial{} }
func (m *Dial) String() string            { return proto.CompactTextString(m) }
func (*Dial) ProtoMessage()               {}
func (*Dial) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *Dial) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *Dial) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Dial) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Dial) GetLevel() float64 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *Dial) GetModTime() int64 {
	if m != nil && m.ModTime != nil {
		return *m.ModTime
	}
	return 0
}

type User struct {
	ID               *int64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Username         *string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *User) GetID() int64 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Dial)(nil), "internal.Dial")
	proto.RegisterType((*User)(nil), "internal.User")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x02,
	0xb8, 0x58, 0x5c, 0x32, 0x13, 0x73, 0x84, 0xb8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x98, 0x85, 0xf8, 0xb8, 0xd8, 0x42, 0x8b, 0x53, 0x8b, 0x3c, 0x5d, 0x24, 0x98, 0xc0, 0x7c,
	0x1e, 0x2e, 0x16, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x4e, 0x21, 0x5e, 0x2e,
	0x56, 0x9f, 0xd4, 0xb2, 0xd4, 0x1c, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x46, 0x21, 0x7e, 0x2e, 0x76,
	0xdf, 0xfc, 0x94, 0x90, 0xcc, 0xdc, 0x54, 0x09, 0x56, 0x90, 0x6a, 0x25, 0x15, 0x2e, 0x16, 0x90,
	0x6e, 0x14, 0x13, 0x05, 0xb8, 0x38, 0x40, 0x62, 0x79, 0x20, 0x53, 0x40, 0x66, 0x72, 0x02, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x58, 0x82, 0x8c, 0x1c, 0x92, 0x00, 0x00, 0x00,
}
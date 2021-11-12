// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rmapi.proto

/*
Package rmapi is a generated protocol buffer package.

It is generated from these files:
	rmapi.proto

It has these top-level messages:
	RandomParams
	RandomInt
	DateTime
	RequestDateTime
	RequestPass
	RandomPass
*/
package rmapi

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

// For random number
type RandomParams struct {
	Seed  int64 `protobuf:"varint,1,opt,name=Seed,json=seed" json:"Seed,omitempty"`
	Place int64 `protobuf:"varint,2,opt,name=Place,json=place" json:"Place,omitempty"`
}

func (m *RandomParams) Reset()                    { *m = RandomParams{} }
func (m *RandomParams) String() string            { return proto.CompactTextString(m) }
func (*RandomParams) ProtoMessage()               {}
func (*RandomParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RandomParams) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *RandomParams) GetPlace() int64 {
	if m != nil {
		return m.Place
	}
	return 0
}

type RandomInt struct {
	Value int64 `protobuf:"varint,1,opt,name=Value,json=value" json:"Value,omitempty"`
}

func (m *RandomInt) Reset()                    { *m = RandomInt{} }
func (m *RandomInt) String() string            { return proto.CompactTextString(m) }
func (*RandomInt) ProtoMessage()               {}
func (*RandomInt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RandomInt) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// For date time
type DateTime struct {
	Value string `protobuf:"bytes,1,opt,name=Value,json=value" json:"Value,omitempty"`
}

func (m *DateTime) Reset()                    { *m = DateTime{} }
func (m *DateTime) String() string            { return proto.CompactTextString(m) }
func (*DateTime) ProtoMessage()               {}
func (*DateTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DateTime) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type RequestDateTime struct {
	Value string `protobuf:"bytes,2,opt,name=Value,json=value" json:"Value,omitempty"`
}

func (m *RequestDateTime) Reset()                    { *m = RequestDateTime{} }
func (m *RequestDateTime) String() string            { return proto.CompactTextString(m) }
func (*RequestDateTime) ProtoMessage()               {}
func (*RequestDateTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestDateTime) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// For random password
type RequestPass struct {
	Seed   int64 `protobuf:"varint,1,opt,name=Seed,json=seed" json:"Seed,omitempty"`
	Length int64 `protobuf:"varint,8,opt,name=Length,json=length" json:"Length,omitempty"`
}

func (m *RequestPass) Reset()                    { *m = RequestPass{} }
func (m *RequestPass) String() string            { return proto.CompactTextString(m) }
func (*RequestPass) ProtoMessage()               {}
func (*RequestPass) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestPass) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *RequestPass) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type RandomPass struct {
	Password string `protobuf:"bytes,1,opt,name=Password,json=password" json:"Password,omitempty"`
}

func (m *RandomPass) Reset()                    { *m = RandomPass{} }
func (m *RandomPass) String() string            { return proto.CompactTextString(m) }
func (*RandomPass) ProtoMessage()               {}
func (*RandomPass) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RandomPass) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*RandomParams)(nil), "RandomParams")
	proto.RegisterType((*RandomInt)(nil), "RandomInt")
	proto.RegisterType((*DateTime)(nil), "DateTime")
	proto.RegisterType((*RequestDateTime)(nil), "RequestDateTime")
	proto.RegisterType((*RequestPass)(nil), "RequestPass")
	proto.RegisterType((*RandomPass)(nil), "RandomPass")
}

func init() { proto.RegisterFile("rmapi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6d, 0xe2, 0xe6, 0x4d, 0x8b, 0xb2, 0x88, 0x84, 0x9c, 0x6a, 0x0e, 0xb5, 0x78,
	0x58, 0x41, 0x2f, 0x8a, 0x37, 0x11, 0x8a, 0xe0, 0x21, 0x44, 0xf1, 0xe0, 0x6d, 0x35, 0x83, 0x16,
	0xf2, 0x65, 0x76, 0xab, 0x7f, 0xc0, 0x1f, 0x2e, 0xbb, 0xdd, 0xb6, 0xf1, 0xe3, 0xb4, 0xbc, 0xb3,
	0xcf, 0xc0, 0x33, 0x33, 0x88, 0xba, 0x4a, 0xb6, 0x0b, 0xd1, 0x76, 0x8d, 0x6e, 0xd2, 0x0b, 0x8c,
	0x72, 0x59, 0x17, 0x4d, 0x95, 0xc9, 0x4e, 0x56, 0x8a, 0x73, 0x0c, 0xef, 0x89, 0x8a, 0xd8, 0x9b,
	0x78, 0xb3, 0x9d, 0x7c, 0xa8, 0x88, 0x0a, 0x7e, 0x00, 0x3f, 0x2b, 0xe5, 0x0b, 0xc5, 0x03, 0x5b,
	0xf4, 0x5b, 0x13, 0xd2, 0x23, 0x84, 0xab, 0xce, 0xdb, 0x5a, 0x1b, 0xe4, 0x51, 0x96, 0x4b, 0x72,
	0x7d, 0xfe, 0x87, 0x09, 0xe9, 0x04, 0xec, 0x46, 0x6a, 0x7a, 0x58, 0x54, 0xf4, 0x93, 0x08, 0xd7,
	0xc4, 0x31, 0xf6, 0x72, 0x7a, 0x5f, 0x92, 0xd2, 0x7f, 0xc1, 0x41, 0x1f, 0xbc, 0x44, 0xe4, 0xc0,
	0x4c, 0xaa, 0xff, 0x35, 0x0f, 0x11, 0xdc, 0x51, 0xfd, 0xaa, 0xdf, 0x62, 0x66, 0xab, 0x41, 0x69,
	0x53, 0x3a, 0x03, 0xd6, 0x23, 0x2a, 0xc5, 0x13, 0x30, 0xf3, 0x7e, 0x36, 0x5d, 0xe1, 0x54, 0x58,
	0xeb, 0xf2, 0xd9, 0x97, 0x87, 0x60, 0x85, 0xf2, 0x29, 0x76, 0xe7, 0x64, 0xa5, 0xf8, 0xbe, 0xf8,
	0xa5, 0x98, 0x84, 0x62, 0x63, 0x3b, 0x45, 0x38, 0x27, 0xed, 0x9a, 0xc6, 0xa2, 0xbf, 0xcb, 0x04,
	0x62, 0xbb, 0xa0, 0x13, 0x8c, 0x37, 0x9c, 0xf5, 0x18, 0x89, 0xde, 0x3c, 0x49, 0x24, 0xb6, 0x5f,
	0xd7, 0x78, 0x62, 0xe2, 0xf4, 0xca, 0x5e, 0xe9, 0x39, 0xb0, 0x67, 0x3a, 0xff, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x21, 0x23, 0x38, 0x20, 0xb5, 0x01, 0x00, 0x00,
}

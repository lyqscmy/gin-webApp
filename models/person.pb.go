// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models/person.proto

package models

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

type Person struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_3878c18f357c89b6, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type UsersResponse struct {
	Users                []*Person `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UsersResponse) Reset()         { *m = UsersResponse{} }
func (m *UsersResponse) String() string { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()    {}
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3878c18f357c89b6, []int{1}
}

func (m *UsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResponse.Unmarshal(m, b)
}
func (m *UsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResponse.Marshal(b, m, deterministic)
}
func (m *UsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResponse.Merge(m, src)
}
func (m *UsersResponse) XXX_Size() int {
	return xxx_messageInfo_UsersResponse.Size(m)
}
func (m *UsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResponse proto.InternalMessageInfo

func (m *UsersResponse) GetUsers() []*Person {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*UsersResponse)(nil), "UsersResponse")
}

func init() {
	proto.RegisterFile("models/person.proto", fileDescriptor_3878c18f357c89b6)
}

var fileDescriptor_3878c18f357c89b6 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0x8e, 0xbf, 0x0a, 0xc2, 0x30,
	0x10, 0x87, 0x49, 0xd3, 0x56, 0x3c, 0xff, 0x20, 0x71, 0xc9, 0x22, 0x84, 0x4e, 0x99, 0x5a, 0xd0,
	0x37, 0xf0, 0x09, 0xe4, 0xc0, 0xc5, 0xad, 0xda, 0x43, 0x02, 0xb6, 0x09, 0xbd, 0x3a, 0xf8, 0xf6,
	0x92, 0x64, 0xfb, 0xee, 0xc7, 0x07, 0xf7, 0xc1, 0x71, 0xf4, 0x03, 0x7d, 0xb8, 0x0b, 0x34, 0xb3,
	0x9f, 0xda, 0x30, 0xfb, 0xc5, 0x37, 0x08, 0xf5, 0x2d, 0xdd, 0x6a, 0x0f, 0x85, 0x1b, 0xb4, 0x30,
	0xc2, 0x4a, 0x2c, 0xdc, 0xa0, 0x14, 0x94, 0x53, 0x3f, 0x92, 0x2e, 0x8c, 0xb0, 0x6b, 0x4c, 0xac,
	0x0e, 0x20, 0xfb, 0x37, 0x69, 0x69, 0x84, 0xad, 0x30, 0x62, 0xb4, 0x5e, 0x6e, 0xf9, 0xe9, 0x32,
	0x5b, 0x91, 0x9b, 0x16, 0x76, 0x77, 0xa6, 0x99, 0x91, 0x38, 0xf8, 0x89, 0x49, 0x9d, 0xa0, 0xfa,
	0xc6, 0x41, 0x0b, 0x23, 0xed, 0xe6, 0xbc, 0x6a, 0xf3, 0x4b, 0xcc, 0xeb, 0x75, 0xfb, 0x80, 0x3e,
	0xb8, 0x2e, 0xe7, 0x3d, 0xeb, 0x14, 0x76, 0xf9, 0x07, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xfa, 0x97,
	0xc8, 0xaf, 0x00, 0x00, 0x00,
}

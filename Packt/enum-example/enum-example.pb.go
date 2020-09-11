// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enum-example.proto

package enumpb

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

type DayOfTheWeek int32

const (
	DayOfTheWeek_UNKNOWN   DayOfTheWeek = 0
	DayOfTheWeek_MONDAY    DayOfTheWeek = 1
	DayOfTheWeek_TUESDAY   DayOfTheWeek = 2
	DayOfTheWeek_WEDNESDAY DayOfTheWeek = 3
	DayOfTheWeek_THURSDAY  DayOfTheWeek = 4
	DayOfTheWeek_FRIDAY    DayOfTheWeek = 5
	DayOfTheWeek_SATURDAY  DayOfTheWeek = 6
	DayOfTheWeek_SUNDAY    DayOfTheWeek = 7
)

var DayOfTheWeek_name = map[int32]string{
	0: "UNKNOWN",
	1: "MONDAY",
	2: "TUESDAY",
	3: "WEDNESDAY",
	4: "THURSDAY",
	5: "FRIDAY",
	6: "SATURDAY",
	7: "SUNDAY",
}

var DayOfTheWeek_value = map[string]int32{
	"UNKNOWN":   0,
	"MONDAY":    1,
	"TUESDAY":   2,
	"WEDNESDAY": 3,
	"THURSDAY":  4,
	"FRIDAY":    5,
	"SATURDAY":  6,
	"SUNDAY":    7,
}

func (x DayOfTheWeek) String() string {
	return proto.EnumName(DayOfTheWeek_name, int32(x))
}

func (DayOfTheWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59f7b3214899c0cd, []int{0}
}

type EnumMessage struct {
	Id                   int32        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DayOfTheWeek         DayOfTheWeek `protobuf:"varint,2,opt,name=day_of_the_week,json=dayOfTheWeek,proto3,enum=example.enumerations.DayOfTheWeek" json:"day_of_the_week,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EnumMessage) Reset()         { *m = EnumMessage{} }
func (m *EnumMessage) String() string { return proto.CompactTextString(m) }
func (*EnumMessage) ProtoMessage()    {}
func (*EnumMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_59f7b3214899c0cd, []int{0}
}

func (m *EnumMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumMessage.Unmarshal(m, b)
}
func (m *EnumMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumMessage.Marshal(b, m, deterministic)
}
func (m *EnumMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumMessage.Merge(m, src)
}
func (m *EnumMessage) XXX_Size() int {
	return xxx_messageInfo_EnumMessage.Size(m)
}
func (m *EnumMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EnumMessage proto.InternalMessageInfo

func (m *EnumMessage) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EnumMessage) GetDayOfTheWeek() DayOfTheWeek {
	if m != nil {
		return m.DayOfTheWeek
	}
	return DayOfTheWeek_UNKNOWN
}

func init() {
	proto.RegisterEnum("example.enumerations.DayOfTheWeek", DayOfTheWeek_name, DayOfTheWeek_value)
	proto.RegisterType((*EnumMessage)(nil), "example.enumerations.EnumMessage")
}

func init() {
	proto.RegisterFile("enum-example.proto", fileDescriptor_59f7b3214899c0cd)
}

var fileDescriptor_59f7b3214899c0cd = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcd, 0x2b, 0xcd,
	0xd5, 0x4d, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x81, 0x71, 0x41, 0x72, 0xa9, 0x45, 0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x4a, 0x19, 0x5c, 0xdc,
	0xae, 0x79, 0xa5, 0xb9, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99,
	0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x9e, 0x5c, 0xfc, 0x29,
	0x89, 0x95, 0xf1, 0xf9, 0x69, 0xf1, 0x25, 0x19, 0xa9, 0xf1, 0xe5, 0xa9, 0xa9, 0xd9, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x7c, 0x46, 0x4a, 0x7a, 0xd8, 0x8c, 0xd3, 0x73, 0x49, 0xac, 0xf4, 0x4f, 0x0b,
	0xc9, 0x48, 0x0d, 0x4f, 0x4d, 0xcd, 0x0e, 0xe2, 0x49, 0x41, 0xe2, 0x69, 0x95, 0x73, 0xf1, 0x20,
	0xcb, 0x0a, 0x71, 0x73, 0xb1, 0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x08, 0x71,
	0x71, 0xb1, 0xf9, 0xfa, 0xfb, 0xb9, 0x38, 0x46, 0x0a, 0x30, 0x82, 0x24, 0x42, 0x42, 0x5d, 0x83,
	0x41, 0x1c, 0x26, 0x21, 0x5e, 0x2e, 0xce, 0x70, 0x57, 0x17, 0x3f, 0x08, 0x97, 0x59, 0x88, 0x87,
	0x8b, 0x23, 0xc4, 0x23, 0x34, 0x08, 0xcc, 0x63, 0x01, 0xe9, 0x72, 0x0b, 0xf2, 0x04, 0xb1, 0x59,
	0x41, 0x32, 0xc1, 0x8e, 0x21, 0xa1, 0x41, 0x20, 0x1e, 0x1b, 0x48, 0x26, 0x38, 0x14, 0x6c, 0x1e,
	0xbb, 0x13, 0x47, 0x14, 0x1b, 0xc8, 0x8d, 0x05, 0x49, 0x49, 0x6c, 0xe0, 0x90, 0x30, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x58, 0xbc, 0x3c, 0x08, 0x1f, 0x01, 0x00, 0x00,
}

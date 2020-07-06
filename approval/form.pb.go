// Code generated by protoc-gen-go. DO NOT EDIT.
// source: approval/form.proto

package approval

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

type ComponentType int32

const (
	ComponentType_UNKNOW             ComponentType = 0
	ComponentType_TEXT_FIELD         ComponentType = 1
	ComponentType_TEXTAREA_FIELD     ComponentType = 2
	ComponentType_NUMBER_FIELD       ComponentType = 3
	ComponentType_SELECT_FIELD       ComponentType = 4
	ComponentType_MULTI_SELECT_FIELD ComponentType = 5
	ComponentType_DATE_FIELD         ComponentType = 6
	ComponentType_DATE_RANGE_FIELD   ComponentType = 7
	ComponentType_MONEY_FIELD        ComponentType = 8
	ComponentType_DEPARTMENT_FIELD   ComponentType = 9
	ComponentType_CALCULATE_FIELD    ComponentType = 10
)

var ComponentType_name = map[int32]string{
	0:  "UNKNOW",
	1:  "TEXT_FIELD",
	2:  "TEXTAREA_FIELD",
	3:  "NUMBER_FIELD",
	4:  "SELECT_FIELD",
	5:  "MULTI_SELECT_FIELD",
	6:  "DATE_FIELD",
	7:  "DATE_RANGE_FIELD",
	8:  "MONEY_FIELD",
	9:  "DEPARTMENT_FIELD",
	10: "CALCULATE_FIELD",
}

var ComponentType_value = map[string]int32{
	"UNKNOW":             0,
	"TEXT_FIELD":         1,
	"TEXTAREA_FIELD":     2,
	"NUMBER_FIELD":       3,
	"SELECT_FIELD":       4,
	"MULTI_SELECT_FIELD": 5,
	"DATE_FIELD":         6,
	"DATE_RANGE_FIELD":   7,
	"MONEY_FIELD":        8,
	"DEPARTMENT_FIELD":   9,
	"CALCULATE_FIELD":    10,
}

func (x ComponentType) String() string {
	return proto.EnumName(ComponentType_name, int32(x))
}

func (ComponentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b259d78f125d2b4f, []int{0}
}

type TextField struct {
	ComponentType        ComponentType `protobuf:"varint,1,opt,name=component_type,json=componentType,proto3,enum=approval.ComponentType" json:"component_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *TextField) Reset()         { *m = TextField{} }
func (m *TextField) String() string { return proto.CompactTextString(m) }
func (*TextField) ProtoMessage()    {}
func (*TextField) Descriptor() ([]byte, []int) {
	return fileDescriptor_b259d78f125d2b4f, []int{0}
}

func (m *TextField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextField.Unmarshal(m, b)
}
func (m *TextField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextField.Marshal(b, m, deterministic)
}
func (m *TextField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextField.Merge(m, src)
}
func (m *TextField) XXX_Size() int {
	return xxx_messageInfo_TextField.Size(m)
}
func (m *TextField) XXX_DiscardUnknown() {
	xxx_messageInfo_TextField.DiscardUnknown(m)
}

var xxx_messageInfo_TextField proto.InternalMessageInfo

func (m *TextField) GetComponentType() ComponentType {
	if m != nil {
		return m.ComponentType
	}
	return ComponentType_UNKNOW
}

func init() {
	proto.RegisterEnum("approval.ComponentType", ComponentType_name, ComponentType_value)
	proto.RegisterType((*TextField)(nil), "approval.TextField")
}

func init() { proto.RegisterFile("approval/form.proto", fileDescriptor_b259d78f125d2b4f) }

var fileDescriptor_b259d78f125d2b4f = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4e, 0x83, 0x40,
	0x14, 0x85, 0x8b, 0x3f, 0x58, 0xae, 0x96, 0x92, 0x5b, 0xa3, 0x26, 0x6e, 0x8c, 0x2b, 0xe3, 0x02,
	0x13, 0x75, 0x6d, 0x32, 0xd0, 0xc1, 0x34, 0x85, 0x29, 0xc1, 0x21, 0xfe, 0x6c, 0x08, 0xd6, 0xd1,
	0x98, 0x94, 0xce, 0x04, 0x89, 0xca, 0xc3, 0xfa, 0x2e, 0x06, 0xc2, 0x58, 0xbb, 0x9b, 0xef, 0xbb,
	0x67, 0xce, 0xe2, 0xc0, 0x28, 0x57, 0xaa, 0x94, 0x9f, 0xf9, 0xe2, 0xe2, 0x55, 0x96, 0x85, 0xab,
	0x4a, 0x59, 0x49, 0xec, 0x6b, 0x79, 0x3a, 0x05, 0x8b, 0x8b, 0xef, 0x2a, 0x78, 0x17, 0x8b, 0x17,
	0xbc, 0x01, 0x7b, 0x2e, 0x0b, 0x25, 0x97, 0x62, 0x59, 0x65, 0x55, 0xad, 0xc4, 0x91, 0x71, 0x62,
	0x9c, 0xd9, 0x97, 0x87, 0xae, 0xce, 0xbb, 0xbe, 0xbe, 0xf3, 0x5a, 0x89, 0x64, 0x30, 0xff, 0x8f,
	0xe7, 0x3f, 0x06, 0x0c, 0xd6, 0x02, 0x08, 0x60, 0xa6, 0x6c, 0xca, 0x66, 0xf7, 0x4e, 0x0f, 0x6d,
	0x00, 0x4e, 0x1f, 0x78, 0x16, 0x4c, 0x68, 0x38, 0x76, 0x0c, 0x44, 0xb0, 0x1b, 0x26, 0x09, 0x25,
	0x9d, 0xdb, 0x40, 0x07, 0xf6, 0x58, 0x1a, 0x79, 0x34, 0xe9, 0xcc, 0x66, 0x63, 0xee, 0x68, 0x48,
	0x7d, 0xfd, 0x6f, 0x0b, 0x0f, 0x00, 0xa3, 0x34, 0xe4, 0x93, 0x6c, 0xcd, 0x6f, 0x37, 0xfd, 0x63,
	0xc2, 0x69, 0xc7, 0x26, 0xee, 0x83, 0xd3, 0x72, 0x42, 0xd8, 0xad, 0xb6, 0x3b, 0x38, 0x84, 0xdd,
	0x68, 0xc6, 0xe8, 0x63, 0x27, 0xfa, 0x6d, 0x8c, 0xc6, 0x24, 0xe1, 0x11, 0x65, 0xba, 0xcc, 0xc2,
	0x11, 0x0c, 0x7d, 0x12, 0xfa, 0x69, 0xb8, 0x6a, 0x04, 0xef, 0x1a, 0x8e, 0x45, 0xfe, 0x21, 0xdc,
	0xb7, 0xbc, 0x12, 0x5f, 0x79, 0xbd, 0x5a, 0xa5, 0x5d, 0xd5, 0xb3, 0x02, 0x59, 0x16, 0x71, 0xf3,
	0x8c, 0x7b, 0x4f, 0x7f, 0x13, 0x3f, 0x9b, 0xed, 0xf5, 0xea, 0x37, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x76, 0x2c, 0x13, 0x8a, 0x01, 0x00, 0x00,
}

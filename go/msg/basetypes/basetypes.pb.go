// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/basetypes.proto

package basetypes // import "github.com/roderm/audio-panel-protobuf/go/msg/basetypes"

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

type UpdateType int32

const (
	UpdateType_NEW          UpdateType = 0
	UpdateType_REMOVED      UpdateType = 1
	UpdateType_REACHABLE    UpdateType = 2
	UpdateType_UNREACHABLE  UpdateType = 3
	UpdateType_STATE_CHANGE UpdateType = 4
)

var UpdateType_name = map[int32]string{
	0: "NEW",
	1: "REMOVED",
	2: "REACHABLE",
	3: "UNREACHABLE",
	4: "STATE_CHANGE",
}
var UpdateType_value = map[string]int32{
	"NEW":          0,
	"REMOVED":      1,
	"REACHABLE":    2,
	"UNREACHABLE":  3,
	"STATE_CHANGE": 4,
}

func (x UpdateType) String() string {
	return proto.EnumName(UpdateType_name, int32(x))
}
func (UpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_basetypes_462ddd54e27a0d0d, []int{0}
}

type ValueList struct {
	KeyActive            string       `protobuf:"bytes,1,opt,name=KeyActive" json:"KeyActive"`
	Values               []*ListValue `protobuf:"bytes,2,rep,name=Values" json:"Values"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ValueList) Reset()         { *m = ValueList{} }
func (m *ValueList) String() string { return proto.CompactTextString(m) }
func (*ValueList) ProtoMessage()    {}
func (*ValueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_basetypes_462ddd54e27a0d0d, []int{0}
}
func (m *ValueList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueList.Unmarshal(m, b)
}
func (m *ValueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueList.Marshal(b, m, deterministic)
}
func (dst *ValueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueList.Merge(dst, src)
}
func (m *ValueList) XXX_Size() int {
	return xxx_messageInfo_ValueList.Size(m)
}
func (m *ValueList) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueList.DiscardUnknown(m)
}

var xxx_messageInfo_ValueList proto.InternalMessageInfo

func (m *ValueList) GetKeyActive() string {
	if m != nil {
		return m.KeyActive
	}
	return ""
}

func (m *ValueList) GetValues() []*ListValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type ListValue struct {
	Key    string `protobuf:"bytes,1,opt,name=Key" json:"Key"`
	Active bool   `protobuf:"varint,2,opt,name=Active" json:"Active"`
	// Types that are valid to be assigned to Value:
	//	*ListValue_Boolean
	//	*ListValue_Number
	//	*ListValue_Decimal
	//	*ListValue_Text
	Value                isListValue_Value `protobuf_oneof:"Value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListValue) Reset()         { *m = ListValue{} }
func (m *ListValue) String() string { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()    {}
func (*ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_basetypes_462ddd54e27a0d0d, []int{1}
}
func (m *ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListValue.Unmarshal(m, b)
}
func (m *ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListValue.Marshal(b, m, deterministic)
}
func (dst *ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListValue.Merge(dst, src)
}
func (m *ListValue) XXX_Size() int {
	return xxx_messageInfo_ListValue.Size(m)
}
func (m *ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListValue proto.InternalMessageInfo

type isListValue_Value interface {
	isListValue_Value()
}

type ListValue_Boolean struct {
	Boolean bool `protobuf:"varint,3,opt,name=Boolean,oneof"`
}
type ListValue_Number struct {
	Number int64 `protobuf:"varint,4,opt,name=Number,oneof"`
}
type ListValue_Decimal struct {
	Decimal float32 `protobuf:"fixed32,5,opt,name=Decimal,oneof"`
}
type ListValue_Text struct {
	Text string `protobuf:"bytes,6,opt,name=Text,oneof"`
}

func (*ListValue_Boolean) isListValue_Value() {}
func (*ListValue_Number) isListValue_Value()  {}
func (*ListValue_Decimal) isListValue_Value() {}
func (*ListValue_Text) isListValue_Value()    {}

func (m *ListValue) GetValue() isListValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *ListValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ListValue) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *ListValue) GetBoolean() bool {
	if x, ok := m.GetValue().(*ListValue_Boolean); ok {
		return x.Boolean
	}
	return false
}

func (m *ListValue) GetNumber() int64 {
	if x, ok := m.GetValue().(*ListValue_Number); ok {
		return x.Number
	}
	return 0
}

func (m *ListValue) GetDecimal() float32 {
	if x, ok := m.GetValue().(*ListValue_Decimal); ok {
		return x.Decimal
	}
	return 0
}

func (m *ListValue) GetText() string {
	if x, ok := m.GetValue().(*ListValue_Text); ok {
		return x.Text
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ListValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ListValue_OneofMarshaler, _ListValue_OneofUnmarshaler, _ListValue_OneofSizer, []interface{}{
		(*ListValue_Boolean)(nil),
		(*ListValue_Number)(nil),
		(*ListValue_Decimal)(nil),
		(*ListValue_Text)(nil),
	}
}

func _ListValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ListValue)
	// Value
	switch x := m.Value.(type) {
	case *ListValue_Boolean:
		t := uint64(0)
		if x.Boolean {
			t = 1
		}
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *ListValue_Number:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Number))
	case *ListValue_Decimal:
		b.EncodeVarint(5<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.Decimal)))
	case *ListValue_Text:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Text)
	case nil:
	default:
		return fmt.Errorf("ListValue.Value has unexpected type %T", x)
	}
	return nil
}

func _ListValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ListValue)
	switch tag {
	case 3: // Value.Boolean
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ListValue_Boolean{x != 0}
		return true, err
	case 4: // Value.Number
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &ListValue_Number{int64(x)}
		return true, err
	case 5: // Value.Decimal
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &ListValue_Decimal{math.Float32frombits(uint32(x))}
		return true, err
	case 6: // Value.Text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &ListValue_Text{x}
		return true, err
	default:
		return false, nil
	}
}

func _ListValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ListValue)
	// Value
	switch x := m.Value.(type) {
	case *ListValue_Boolean:
		n += 1 // tag and wire
		n += 1
	case *ListValue_Number:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Number))
	case *ListValue_Decimal:
		n += 1 // tag and wire
		n += 4
	case *ListValue_Text:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Text)))
		n += len(x.Text)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ValueList)(nil), "ValueList")
	proto.RegisterType((*ListValue)(nil), "ListValue")
	proto.RegisterEnum("UpdateType", UpdateType_name, UpdateType_value)
}

func init() { proto.RegisterFile("messages/basetypes.proto", fileDescriptor_basetypes_462ddd54e27a0d0d) }

var fileDescriptor_basetypes_462ddd54e27a0d0d = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0x5f, 0x6b, 0xa3, 0x40,
	0x14, 0xc5, 0xfd, 0x93, 0xe8, 0x7a, 0xb3, 0xcb, 0xca, 0xb0, 0x2c, 0xc3, 0xb2, 0x0f, 0x92, 0x27,
	0x59, 0x48, 0x84, 0xed, 0x43, 0xe9, 0xa3, 0x26, 0x52, 0x21, 0x89, 0x85, 0xa9, 0x49, 0xa1, 0x2f,
	0x65, 0x4c, 0x6e, 0xad, 0xa0, 0x19, 0x71, 0xb4, 0xd4, 0x2f, 0xd4, 0xcf, 0x59, 0x14, 0xdb, 0xbc,
	0xcd, 0xf9, 0x9d, 0x39, 0x67, 0x86, 0x7b, 0x81, 0x96, 0x28, 0x25, 0xcf, 0x50, 0x7a, 0x29, 0x97,
	0xd8, 0x74, 0x15, 0xca, 0x65, 0x55, 0x8b, 0x46, 0xcc, 0x77, 0x60, 0x1d, 0x78, 0xd1, 0xe2, 0x36,
	0x97, 0x0d, 0xf9, 0x0b, 0xd6, 0x06, 0x3b, 0xff, 0xd8, 0xe4, 0xaf, 0x48, 0x55, 0x47, 0x75, 0x2d,
	0x76, 0x01, 0x64, 0x0e, 0xc6, 0x70, 0x55, 0x52, 0xcd, 0xd1, 0xdd, 0xd9, 0x7f, 0x58, 0xf6, 0xa1,
	0x01, 0xb1, 0xd1, 0x99, 0xbf, 0xab, 0x60, 0x7d, 0x51, 0x62, 0x83, 0xbe, 0xc1, 0x6e, 0x6c, 0xea,
	0x8f, 0xe4, 0x37, 0x18, 0x63, 0xbd, 0xe6, 0xa8, 0xee, 0x37, 0x36, 0x2a, 0xf2, 0x07, 0xcc, 0x40,
	0x88, 0x02, 0xf9, 0x99, 0xea, 0xbd, 0x11, 0x29, 0xec, 0x13, 0x10, 0x0a, 0x46, 0xdc, 0x96, 0x29,
	0xd6, 0x74, 0xe2, 0xa8, 0xae, 0x1e, 0x29, 0x6c, 0xd4, 0x7d, 0x6a, 0x8d, 0xc7, 0xbc, 0xe4, 0x05,
	0x9d, 0x3a, 0xaa, 0xab, 0xf5, 0xa9, 0x11, 0x90, 0x5f, 0x30, 0x49, 0xf0, 0xad, 0xa1, 0x46, 0xff,
	0x78, 0xa4, 0xb0, 0x41, 0x05, 0x26, 0x4c, 0x87, 0xaf, 0xfd, 0x4b, 0x00, 0xf6, 0xd5, 0x89, 0x37,
	0x98, 0x74, 0x15, 0x12, 0x13, 0xf4, 0x38, 0x7c, 0xb0, 0x15, 0x32, 0x03, 0x93, 0x85, 0xbb, 0xbb,
	0x43, 0xb8, 0xb6, 0x55, 0xf2, 0x03, 0x2c, 0x16, 0xfa, 0xab, 0xc8, 0x0f, 0xb6, 0xa1, 0xad, 0x91,
	0x9f, 0x30, 0xdb, 0xc7, 0x17, 0xa0, 0x13, 0x1b, 0xbe, 0xdf, 0x27, 0x7e, 0x12, 0x3e, 0xad, 0x22,
	0x3f, 0xbe, 0x0d, 0xed, 0x49, 0x70, 0xf3, 0x78, 0x9d, 0xe5, 0xcd, 0x4b, 0x9b, 0x2e, 0x8f, 0xa2,
	0xf4, 0x6a, 0x71, 0xc2, 0xba, 0xf4, 0x78, 0x7b, 0xca, 0xc5, 0xa2, 0xe2, 0x67, 0x2c, 0x16, 0xc3,
	0xd0, 0xd3, 0xf6, 0xd9, 0xcb, 0x84, 0x57, 0xca, 0xec, 0xb2, 0x8e, 0xd4, 0x18, 0xac, 0xab, 0x8f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xfa, 0x4a, 0xea, 0xab, 0x01, 0x00, 0x00,
}
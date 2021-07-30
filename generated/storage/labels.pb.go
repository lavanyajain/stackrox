// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/labels.proto

package storage

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type LabelSelector_Operator int32

const (
	LabelSelector_UNKNOWN    LabelSelector_Operator = 0
	LabelSelector_IN         LabelSelector_Operator = 1
	LabelSelector_NOT_IN     LabelSelector_Operator = 2
	LabelSelector_EXISTS     LabelSelector_Operator = 3
	LabelSelector_NOT_EXISTS LabelSelector_Operator = 4
)

var LabelSelector_Operator_name = map[int32]string{
	0: "UNKNOWN",
	1: "IN",
	2: "NOT_IN",
	3: "EXISTS",
	4: "NOT_EXISTS",
}

var LabelSelector_Operator_value = map[string]int32{
	"UNKNOWN":    0,
	"IN":         1,
	"NOT_IN":     2,
	"EXISTS":     3,
	"NOT_EXISTS": 4,
}

func (x LabelSelector_Operator) String() string {
	return proto.EnumName(LabelSelector_Operator_name, int32(x))
}

func (LabelSelector_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a13142881b5e7a12, []int{0, 0}
}

type SetBasedLabelSelector_Operator int32

const (
	SetBasedLabelSelector_UNKNOWN    SetBasedLabelSelector_Operator = 0
	SetBasedLabelSelector_IN         SetBasedLabelSelector_Operator = 1
	SetBasedLabelSelector_NOT_IN     SetBasedLabelSelector_Operator = 2
	SetBasedLabelSelector_EXISTS     SetBasedLabelSelector_Operator = 3
	SetBasedLabelSelector_NOT_EXISTS SetBasedLabelSelector_Operator = 4
)

var SetBasedLabelSelector_Operator_name = map[int32]string{
	0: "UNKNOWN",
	1: "IN",
	2: "NOT_IN",
	3: "EXISTS",
	4: "NOT_EXISTS",
}

var SetBasedLabelSelector_Operator_value = map[string]int32{
	"UNKNOWN":    0,
	"IN":         1,
	"NOT_IN":     2,
	"EXISTS":     3,
	"NOT_EXISTS": 4,
}

func (x SetBasedLabelSelector_Operator) String() string {
	return proto.EnumName(SetBasedLabelSelector_Operator_name, int32(x))
}

func (SetBasedLabelSelector_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a13142881b5e7a12, []int{1, 0}
}

// Label selector components are joined with logical AND, see
//     https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
//
// Next available tag: 3
type LabelSelector struct {
	// This is actually a oneof, but we can't make it one due to backwards
	// compatibility constraints.
	MatchLabels          map[string]string            `protobuf:"bytes,1,rep,name=match_labels,json=matchLabels,proto3" json:"match_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Requirements         []*LabelSelector_Requirement `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *LabelSelector) Reset()         { *m = LabelSelector{} }
func (m *LabelSelector) String() string { return proto.CompactTextString(m) }
func (*LabelSelector) ProtoMessage()    {}
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_a13142881b5e7a12, []int{0}
}
func (m *LabelSelector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelSelector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelector.Merge(m, src)
}
func (m *LabelSelector) XXX_Size() int {
	return m.Size()
}
func (m *LabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelector proto.InternalMessageInfo

func (m *LabelSelector) GetMatchLabels() map[string]string {
	if m != nil {
		return m.MatchLabels
	}
	return nil
}

func (m *LabelSelector) GetRequirements() []*LabelSelector_Requirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *LabelSelector) MessageClone() proto.Message {
	return m.Clone()
}
func (m *LabelSelector) Clone() *LabelSelector {
	if m == nil {
		return nil
	}
	cloned := new(LabelSelector)
	*cloned = *m

	if m.MatchLabels != nil {
		cloned.MatchLabels = make(map[string]string, len(m.MatchLabels))
		for k, v := range m.MatchLabels {
			cloned.MatchLabels[k] = v
		}
	}
	if m.Requirements != nil {
		cloned.Requirements = make([]*LabelSelector_Requirement, len(m.Requirements))
		for idx, v := range m.Requirements {
			cloned.Requirements[idx] = v.Clone()
		}
	}
	return cloned
}

// Next available tag: 4
type LabelSelector_Requirement struct {
	Key                  string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Op                   LabelSelector_Operator `protobuf:"varint,2,opt,name=op,proto3,enum=storage.LabelSelector_Operator" json:"op,omitempty"`
	Values               []string               `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *LabelSelector_Requirement) Reset()         { *m = LabelSelector_Requirement{} }
func (m *LabelSelector_Requirement) String() string { return proto.CompactTextString(m) }
func (*LabelSelector_Requirement) ProtoMessage()    {}
func (*LabelSelector_Requirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_a13142881b5e7a12, []int{0, 0}
}
func (m *LabelSelector_Requirement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LabelSelector_Requirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LabelSelector_Requirement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LabelSelector_Requirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelSelector_Requirement.Merge(m, src)
}
func (m *LabelSelector_Requirement) XXX_Size() int {
	return m.Size()
}
func (m *LabelSelector_Requirement) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelSelector_Requirement.DiscardUnknown(m)
}

var xxx_messageInfo_LabelSelector_Requirement proto.InternalMessageInfo

func (m *LabelSelector_Requirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelSelector_Requirement) GetOp() LabelSelector_Operator {
	if m != nil {
		return m.Op
	}
	return LabelSelector_UNKNOWN
}

func (m *LabelSelector_Requirement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *LabelSelector_Requirement) MessageClone() proto.Message {
	return m.Clone()
}
func (m *LabelSelector_Requirement) Clone() *LabelSelector_Requirement {
	if m == nil {
		return nil
	}
	cloned := new(LabelSelector_Requirement)
	*cloned = *m

	if m.Values != nil {
		cloned.Values = make([]string, len(m.Values))
		copy(cloned.Values, m.Values)
	}
	return cloned
}

// SetBasedLabelSelector only allows set-based label requirements.
//
// Next available tag: 3
type SetBasedLabelSelector struct {
	Requirements         []*SetBasedLabelSelector_Requirement `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *SetBasedLabelSelector) Reset()         { *m = SetBasedLabelSelector{} }
func (m *SetBasedLabelSelector) String() string { return proto.CompactTextString(m) }
func (*SetBasedLabelSelector) ProtoMessage()    {}
func (*SetBasedLabelSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_a13142881b5e7a12, []int{1}
}
func (m *SetBasedLabelSelector) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetBasedLabelSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetBasedLabelSelector.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetBasedLabelSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBasedLabelSelector.Merge(m, src)
}
func (m *SetBasedLabelSelector) XXX_Size() int {
	return m.Size()
}
func (m *SetBasedLabelSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBasedLabelSelector.DiscardUnknown(m)
}

var xxx_messageInfo_SetBasedLabelSelector proto.InternalMessageInfo

func (m *SetBasedLabelSelector) GetRequirements() []*SetBasedLabelSelector_Requirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *SetBasedLabelSelector) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SetBasedLabelSelector) Clone() *SetBasedLabelSelector {
	if m == nil {
		return nil
	}
	cloned := new(SetBasedLabelSelector)
	*cloned = *m

	if m.Requirements != nil {
		cloned.Requirements = make([]*SetBasedLabelSelector_Requirement, len(m.Requirements))
		for idx, v := range m.Requirements {
			cloned.Requirements[idx] = v.Clone()
		}
	}
	return cloned
}

// Next available tag: 4
type SetBasedLabelSelector_Requirement struct {
	Key                  string                         `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Op                   SetBasedLabelSelector_Operator `protobuf:"varint,2,opt,name=op,proto3,enum=storage.SetBasedLabelSelector_Operator" json:"op,omitempty"`
	Values               []string                       `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SetBasedLabelSelector_Requirement) Reset()         { *m = SetBasedLabelSelector_Requirement{} }
func (m *SetBasedLabelSelector_Requirement) String() string { return proto.CompactTextString(m) }
func (*SetBasedLabelSelector_Requirement) ProtoMessage()    {}
func (*SetBasedLabelSelector_Requirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_a13142881b5e7a12, []int{1, 0}
}
func (m *SetBasedLabelSelector_Requirement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetBasedLabelSelector_Requirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetBasedLabelSelector_Requirement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetBasedLabelSelector_Requirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBasedLabelSelector_Requirement.Merge(m, src)
}
func (m *SetBasedLabelSelector_Requirement) XXX_Size() int {
	return m.Size()
}
func (m *SetBasedLabelSelector_Requirement) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBasedLabelSelector_Requirement.DiscardUnknown(m)
}

var xxx_messageInfo_SetBasedLabelSelector_Requirement proto.InternalMessageInfo

func (m *SetBasedLabelSelector_Requirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetBasedLabelSelector_Requirement) GetOp() SetBasedLabelSelector_Operator {
	if m != nil {
		return m.Op
	}
	return SetBasedLabelSelector_UNKNOWN
}

func (m *SetBasedLabelSelector_Requirement) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SetBasedLabelSelector_Requirement) MessageClone() proto.Message {
	return m.Clone()
}
func (m *SetBasedLabelSelector_Requirement) Clone() *SetBasedLabelSelector_Requirement {
	if m == nil {
		return nil
	}
	cloned := new(SetBasedLabelSelector_Requirement)
	*cloned = *m

	if m.Values != nil {
		cloned.Values = make([]string, len(m.Values))
		copy(cloned.Values, m.Values)
	}
	return cloned
}

func init() {
	proto.RegisterEnum("storage.LabelSelector_Operator", LabelSelector_Operator_name, LabelSelector_Operator_value)
	proto.RegisterEnum("storage.SetBasedLabelSelector_Operator", SetBasedLabelSelector_Operator_name, SetBasedLabelSelector_Operator_value)
	proto.RegisterType((*LabelSelector)(nil), "storage.LabelSelector")
	proto.RegisterMapType((map[string]string)(nil), "storage.LabelSelector.MatchLabelsEntry")
	proto.RegisterType((*LabelSelector_Requirement)(nil), "storage.LabelSelector.Requirement")
	proto.RegisterType((*SetBasedLabelSelector)(nil), "storage.SetBasedLabelSelector")
	proto.RegisterType((*SetBasedLabelSelector_Requirement)(nil), "storage.SetBasedLabelSelector.Requirement")
}

func init() { proto.RegisterFile("storage/labels.proto", fileDescriptor_a13142881b5e7a12) }

var fileDescriptor_a13142881b5e7a12 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0xcf, 0x49, 0x4c, 0x4a, 0xcd, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x87, 0x8a, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0xc5, 0xf4, 0x41, 0x2c, 0x88,
	0xb4, 0xd2, 0x34, 0x66, 0x2e, 0x5e, 0x1f, 0x90, 0xfa, 0xe0, 0xd4, 0x9c, 0xd4, 0xe4, 0x92, 0xfc,
	0x22, 0x21, 0x2f, 0x2e, 0x9e, 0xdc, 0xc4, 0x92, 0xe4, 0x8c, 0x78, 0x88, 0x31, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0xea, 0x7a, 0x50, 0x73, 0xf4, 0x50, 0x54, 0xeb, 0xf9, 0x82, 0x94, 0x82,
	0x85, 0x8a, 0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xb8, 0x73, 0x11, 0x22, 0x42, 0x6e, 0x5c, 0x3c,
	0x45, 0xa9, 0x85, 0xa5, 0x99, 0x45, 0xa9, 0xb9, 0xa9, 0x79, 0x25, 0xc5, 0x12, 0x4c, 0x60, 0xb3,
	0x94, 0x70, 0x98, 0x15, 0x84, 0x50, 0x1a, 0x84, 0xa2, 0x4f, 0x2a, 0x83, 0x8b, 0x1b, 0x49, 0x52,
	0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc4, 0x14,
	0xd2, 0xe7, 0x62, 0xca, 0x2f, 0x90, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x33, 0x92, 0xc7, 0x61, 0xbc,
	0x7f, 0x41, 0x6a, 0x51, 0x62, 0x49, 0x7e, 0x51, 0x10, 0x53, 0x7e, 0x81, 0x90, 0x18, 0x17, 0x5b,
	0x59, 0x62, 0x4e, 0x69, 0x6a, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x94, 0x27, 0x65,
	0xc7, 0x25, 0x80, 0xee, 0x25, 0x2c, 0xd6, 0x89, 0x70, 0xb1, 0x82, 0xd5, 0x83, 0x6d, 0xe4, 0x0c,
	0x82, 0x70, 0xac, 0x98, 0x2c, 0x18, 0x95, 0xdc, 0xb9, 0x38, 0x60, 0xf6, 0x08, 0x71, 0x73, 0xb1,
	0x87, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0x09, 0x30, 0x08, 0xb1, 0x71, 0x31, 0x79, 0xfa, 0x09,
	0x30, 0x0a, 0x71, 0x71, 0xb1, 0xf9, 0xf9, 0x87, 0xc4, 0x7b, 0xfa, 0x09, 0x30, 0x81, 0xd8, 0xae,
	0x11, 0x9e, 0xc1, 0x21, 0xc1, 0x02, 0xcc, 0x42, 0x7c, 0x5c, 0x5c, 0x20, 0x71, 0x28, 0x9f, 0x45,
	0x69, 0x05, 0x13, 0x97, 0x68, 0x70, 0x6a, 0x89, 0x53, 0x62, 0x71, 0x6a, 0x0a, 0x6a, 0x04, 0xf9,
	0x61, 0x0d, 0x54, 0x2d, 0xb8, 0xaf, 0xb1, 0xea, 0xc2, 0x13, 0xb8, 0x05, 0x84, 0x02, 0xd7, 0x1c,
	0x29, 0x70, 0xd5, 0x09, 0x58, 0x43, 0x4c, 0x20, 0x53, 0x2d, 0x90, 0xbc, 0x58, 0x38, 0x18, 0x05,
	0x98, 0x9c, 0x4c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x19, 0x8f, 0xe5, 0x18, 0xb8, 0x24, 0x33, 0xf3, 0xf5, 0x8a, 0x4b, 0x12, 0x93, 0xb3, 0x8b, 0xf2,
	0x2b, 0x20, 0x09, 0x1d, 0xe6, 0xec, 0x28, 0x58, 0x7e, 0x48, 0x62, 0x03, 0x8b, 0x1b, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe1, 0x7b, 0x5b, 0xcc, 0x37, 0x03, 0x00, 0x00,
}

func (m *LabelSelector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LabelSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Requirements) > 0 {
		for iNdEx := len(m.Requirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLabels(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.MatchLabels) > 0 {
		for k := range m.MatchLabels {
			v := m.MatchLabels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintLabels(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintLabels(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintLabels(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LabelSelector_Requirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelSelector_Requirement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LabelSelector_Requirement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarintLabels(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Op != 0 {
		i = encodeVarintLabels(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintLabels(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SetBasedLabelSelector) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetBasedLabelSelector) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetBasedLabelSelector) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Requirements) > 0 {
		for iNdEx := len(m.Requirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLabels(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	return len(dAtA) - i, nil
}

func (m *SetBasedLabelSelector_Requirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetBasedLabelSelector_Requirement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetBasedLabelSelector_Requirement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Values) > 0 {
		for iNdEx := len(m.Values) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Values[iNdEx])
			copy(dAtA[i:], m.Values[iNdEx])
			i = encodeVarintLabels(dAtA, i, uint64(len(m.Values[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Op != 0 {
		i = encodeVarintLabels(dAtA, i, uint64(m.Op))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintLabels(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLabels(dAtA []byte, offset int, v uint64) int {
	offset -= sovLabels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LabelSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MatchLabels) > 0 {
		for k, v := range m.MatchLabels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovLabels(uint64(len(k))) + 1 + len(v) + sovLabels(uint64(len(v)))
			n += mapEntrySize + 1 + sovLabels(uint64(mapEntrySize))
		}
	}
	if len(m.Requirements) > 0 {
		for _, e := range m.Requirements {
			l = e.Size()
			n += 1 + l + sovLabels(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LabelSelector_Requirement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovLabels(uint64(l))
	}
	if m.Op != 0 {
		n += 1 + sovLabels(uint64(m.Op))
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovLabels(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SetBasedLabelSelector) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requirements) > 0 {
		for _, e := range m.Requirements {
			l = e.Size()
			n += 1 + l + sovLabels(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SetBasedLabelSelector_Requirement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovLabels(uint64(l))
	}
	if m.Op != 0 {
		n += 1 + sovLabels(uint64(m.Op))
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovLabels(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLabels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLabels(x uint64) (n int) {
	return sovLabels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LabelSelector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLabels
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LabelSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LabelSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchLabels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLabels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLabels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MatchLabels == nil {
				m.MatchLabels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLabels
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLabels
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthLabels
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthLabels
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLabels
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthLabels
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthLabels
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipLabels(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthLabels
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MatchLabels[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLabels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLabels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requirements = append(m.Requirements, &LabelSelector_Requirement{})
			if err := m.Requirements[len(m.Requirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLabels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLabels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LabelSelector_Requirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLabels
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Requirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Requirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLabels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLabels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= LabelSelector_Operator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLabels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLabels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLabels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLabels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetBasedLabelSelector) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLabels
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetBasedLabelSelector: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetBasedLabelSelector: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLabels
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLabels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requirements = append(m.Requirements, &SetBasedLabelSelector_Requirement{})
			if err := m.Requirements[len(m.Requirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLabels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLabels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SetBasedLabelSelector_Requirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLabels
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Requirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Requirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLabels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLabels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= SetBasedLabelSelector_Operator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLabels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLabels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLabels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLabels
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLabels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLabels
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLabels
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLabels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLabels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLabels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLabels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLabels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLabels = fmt.Errorf("proto: unexpected end of group")
)
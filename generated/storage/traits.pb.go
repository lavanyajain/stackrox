// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/traits.proto

package storage

import (
	fmt "fmt"
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

// EXPERIMENTAL.
//
// MutabilityMode specifies whether and how an object can be modified. Default
// is ALLOW_MUTATE and means there are no modification restrictions; this is equivalent
// to the absence of MutabilityMode specification. ALLOW_MUTATE_FORCED forbids all
// modifying operations except object removal with force bit on.
//
// Be careful when changing the state of this field. For example, modifying an
// object from ALLOW_MUTATE to ALLOW_MUTATE_FORCED is allowed but will prohibit any further
// changes to it, including modifying it back to ALLOW_MUTATE.
type Traits_MutabilityMode int32

const (
	Traits_ALLOW_MUTATE        Traits_MutabilityMode = 0
	Traits_ALLOW_MUTATE_FORCED Traits_MutabilityMode = 1
)

var Traits_MutabilityMode_name = map[int32]string{
	0: "ALLOW_MUTATE",
	1: "ALLOW_MUTATE_FORCED",
}

var Traits_MutabilityMode_value = map[string]int32{
	"ALLOW_MUTATE":        0,
	"ALLOW_MUTATE_FORCED": 1,
}

func (x Traits_MutabilityMode) String() string {
	return proto.EnumName(Traits_MutabilityMode_name, int32(x))
}

func (Traits_MutabilityMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec31914177d462a1, []int{0, 0}
}

// EXPERIMENTAL.
// visibility allows to specify whether the object should be visible for certain APIs.
type Traits_Visibility int32

const (
	Traits_VISIBLE Traits_Visibility = 0
	Traits_HIDDEN  Traits_Visibility = 1
)

var Traits_Visibility_name = map[int32]string{
	0: "VISIBLE",
	1: "HIDDEN",
}

var Traits_Visibility_value = map[string]int32{
	"VISIBLE": 0,
	"HIDDEN":  1,
}

func (x Traits_Visibility) String() string {
	return proto.EnumName(Traits_Visibility_name, int32(x))
}

func (Traits_Visibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ec31914177d462a1, []int{0, 1}
}

// EXPERIMENTAL.
type Traits struct {
	MutabilityMode       Traits_MutabilityMode `protobuf:"varint,1,opt,name=mutability_mode,json=mutabilityMode,proto3,enum=storage.Traits_MutabilityMode" json:"mutability_mode,omitempty"`
	Visibility           Traits_Visibility     `protobuf:"varint,2,opt,name=visibility,proto3,enum=storage.Traits_Visibility" json:"visibility,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Traits) Reset()         { *m = Traits{} }
func (m *Traits) String() string { return proto.CompactTextString(m) }
func (*Traits) ProtoMessage()    {}
func (*Traits) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec31914177d462a1, []int{0}
}
func (m *Traits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Traits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Traits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Traits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Traits.Merge(m, src)
}
func (m *Traits) XXX_Size() int {
	return m.Size()
}
func (m *Traits) XXX_DiscardUnknown() {
	xxx_messageInfo_Traits.DiscardUnknown(m)
}

var xxx_messageInfo_Traits proto.InternalMessageInfo

func (m *Traits) GetMutabilityMode() Traits_MutabilityMode {
	if m != nil {
		return m.MutabilityMode
	}
	return Traits_ALLOW_MUTATE
}

func (m *Traits) GetVisibility() Traits_Visibility {
	if m != nil {
		return m.Visibility
	}
	return Traits_VISIBLE
}

func (m *Traits) MessageClone() proto.Message {
	return m.Clone()
}
func (m *Traits) Clone() *Traits {
	if m == nil {
		return nil
	}
	cloned := new(Traits)
	*cloned = *m

	return cloned
}

func init() {
	proto.RegisterEnum("storage.Traits_MutabilityMode", Traits_MutabilityMode_name, Traits_MutabilityMode_value)
	proto.RegisterEnum("storage.Traits_Visibility", Traits_Visibility_name, Traits_Visibility_value)
	proto.RegisterType((*Traits)(nil), "storage.Traits")
}

func init() { proto.RegisterFile("storage/traits.proto", fileDescriptor_ec31914177d462a1) }

var fileDescriptor_ec31914177d462a1 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x87, 0x8a, 0x2a, 0x7d, 0x64, 0xe4, 0x62, 0x0b, 0x01, 0xcb, 0x08, 0xb9, 0x73, 0xf1,
	0xe7, 0x96, 0x96, 0x24, 0x26, 0x65, 0xe6, 0x64, 0x96, 0x54, 0xc6, 0xe7, 0xe6, 0xa7, 0xa4, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xe9, 0x41, 0x55, 0xeb, 0x41, 0x54, 0xea, 0xf9, 0xc2,
	0x95, 0xf9, 0xe6, 0xa7, 0xa4, 0x06, 0xf1, 0xe5, 0xa2, 0xf0, 0x85, 0xac, 0xb8, 0xb8, 0xca, 0x32,
	0x8b, 0x33, 0x21, 0x22, 0x12, 0x4c, 0x60, 0x33, 0xa4, 0xd0, 0xcd, 0x08, 0x83, 0xab, 0x08, 0x42,
	0x52, 0xad, 0x64, 0xcd, 0xc5, 0x87, 0x6a, 0xba, 0x90, 0x00, 0x17, 0x8f, 0xa3, 0x8f, 0x8f, 0x7f,
	0x78, 0xbc, 0x6f, 0x68, 0x88, 0x63, 0x88, 0xab, 0x00, 0x83, 0x90, 0x38, 0x97, 0x30, 0xb2, 0x48,
	0xbc, 0x9b, 0x7f, 0x90, 0xb3, 0xab, 0x8b, 0x00, 0xa3, 0x92, 0x2a, 0x17, 0x17, 0xc2, 0x58, 0x21,
	0x6e, 0x2e, 0xf6, 0x30, 0xcf, 0x60, 0x4f, 0x27, 0x1f, 0x90, 0x1e, 0x2e, 0x2e, 0x36, 0x0f, 0x4f,
	0x17, 0x17, 0x57, 0x3f, 0x01, 0x46, 0x27, 0x93, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63,
	0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xc9, 0xcc, 0x7c, 0xbd, 0xe2, 0x92,
	0xc4, 0xe4, 0xec, 0xa2, 0xfc, 0x0a, 0x48, 0x08, 0xc1, 0x9c, 0x1b, 0x05, 0x0b, 0xa9, 0x24, 0x36,
	0xb0, 0xb8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x1b, 0x2f, 0xff, 0x51, 0x01, 0x00, 0x00,
}

func (m *Traits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Traits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Traits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Visibility != 0 {
		i = encodeVarintTraits(dAtA, i, uint64(m.Visibility))
		i--
		dAtA[i] = 0x10
	}
	if m.MutabilityMode != 0 {
		i = encodeVarintTraits(dAtA, i, uint64(m.MutabilityMode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTraits(dAtA []byte, offset int, v uint64) int {
	offset -= sovTraits(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Traits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MutabilityMode != 0 {
		n += 1 + sovTraits(uint64(m.MutabilityMode))
	}
	if m.Visibility != 0 {
		n += 1 + sovTraits(uint64(m.Visibility))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTraits(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTraits(x uint64) (n int) {
	return sovTraits(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Traits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTraits
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
			return fmt.Errorf("proto: Traits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Traits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MutabilityMode", wireType)
			}
			m.MutabilityMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraits
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MutabilityMode |= Traits_MutabilityMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Visibility", wireType)
			}
			m.Visibility = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTraits
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Visibility |= Traits_Visibility(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTraits(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTraits
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
func skipTraits(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTraits
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
					return 0, ErrIntOverflowTraits
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
					return 0, ErrIntOverflowTraits
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
				return 0, ErrInvalidLengthTraits
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTraits
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTraits
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTraits        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTraits          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTraits = fmt.Errorf("proto: unexpected end of group")
)
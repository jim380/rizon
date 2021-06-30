// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: treasury/treasury.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Currency defines a single currency info
type Currency struct {
	// denom is the name of the currency
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	// desc is a description of the currency
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" yaml:"desc"`
	// owner is who can mint this currency
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// mintable indicates whether this currency can be minted or not
	Mintable bool `protobuf:"varint,4,opt,name=mintable,proto3" json:"mintable,omitempty" yaml:"mintable"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcc352ab3574cb55, []int{0}
}
func (m *Currency) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return m.Size()
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Currency) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Currency) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Currency) GetMintable() bool {
	if m != nil {
		return m.Mintable
	}
	return false
}

// Currencies
type Currencies struct {
	// denoms is the denom list of all currencies
	Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty" yaml:"denoms"`
}

func (m *Currencies) Reset()         { *m = Currencies{} }
func (m *Currencies) String() string { return proto.CompactTextString(m) }
func (*Currencies) ProtoMessage()    {}
func (*Currencies) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcc352ab3574cb55, []int{1}
}
func (m *Currencies) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Currencies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Currencies.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Currencies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currencies.Merge(m, src)
}
func (m *Currencies) XXX_Size() int {
	return m.Size()
}
func (m *Currencies) XXX_DiscardUnknown() {
	xxx_messageInfo_Currencies.DiscardUnknown(m)
}

var xxx_messageInfo_Currencies proto.InternalMessageInfo

func (m *Currencies) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

// Sequence
type Sequence struct {
	// number is the currency sequence number of current state
	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty" yaml:"number"`
}

func (m *Sequence) Reset()         { *m = Sequence{} }
func (m *Sequence) String() string { return proto.CompactTextString(m) }
func (*Sequence) ProtoMessage()    {}
func (*Sequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_fcc352ab3574cb55, []int{2}
}
func (m *Sequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sequence.Merge(m, src)
}
func (m *Sequence) XXX_Size() int {
	return m.Size()
}
func (m *Sequence) XXX_DiscardUnknown() {
	xxx_messageInfo_Sequence.DiscardUnknown(m)
}

var xxx_messageInfo_Sequence proto.InternalMessageInfo

func (m *Sequence) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*Currency)(nil), "rizonworld.rizon.treasury.Currency")
	proto.RegisterType((*Currencies)(nil), "rizonworld.rizon.treasury.Currencies")
	proto.RegisterType((*Sequence)(nil), "rizonworld.rizon.treasury.Sequence")
}

func init() { proto.RegisterFile("treasury/treasury.proto", fileDescriptor_fcc352ab3574cb55) }

var fileDescriptor_fcc352ab3574cb55 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x6b, 0x5a, 0xaa, 0xd4, 0x80, 0x0a, 0x01, 0x89, 0xc0, 0xe0, 0x54, 0x46, 0x42, 0x65,
	0x68, 0x3c, 0x20, 0x84, 0xc4, 0x18, 0x06, 0x76, 0xb3, 0xb1, 0x35, 0xe9, 0x53, 0x89, 0xd4, 0xd8,
	0xc5, 0x4e, 0x54, 0xc2, 0x29, 0x38, 0x0a, 0xc7, 0x60, 0xec, 0xc8, 0x14, 0xa1, 0xf6, 0x06, 0x39,
	0x01, 0xaa, 0x9d, 0x40, 0xc5, 0xf6, 0xeb, 0xfd, 0xdf, 0x7b, 0xfa, 0xfd, 0x1b, 0x9f, 0x66, 0x0a,
	0xc6, 0x3a, 0x57, 0x05, 0x6b, 0x44, 0x30, 0x57, 0x32, 0x93, 0xee, 0x99, 0x4a, 0xde, 0xa4, 0x58,
	0x48, 0x35, 0x9b, 0x04, 0x46, 0x06, 0x0d, 0x70, 0x7e, 0x32, 0x95, 0x53, 0x69, 0x28, 0xb6, 0x51,
	0x76, 0x81, 0x7e, 0x20, 0xec, 0xdc, 0xe7, 0x4a, 0x81, 0x88, 0x0b, 0xf7, 0x12, 0xef, 0x4e, 0x40,
	0xc8, 0xd4, 0x43, 0x03, 0x34, 0xec, 0x85, 0x87, 0x55, 0xe9, 0xef, 0x17, 0xe3, 0x74, 0x76, 0x47,
	0xcd, 0x98, 0x72, 0x6b, 0xbb, 0x17, 0xb8, 0x33, 0x01, 0x1d, 0x7b, 0x3b, 0x06, 0xeb, 0x57, 0xa5,
	0xbf, 0xd7, 0x60, 0x3a, 0xa6, 0xdc, 0x98, 0x9b, 0x63, 0x72, 0x21, 0x40, 0x79, 0xed, 0xff, 0xc7,
	0xcc, 0x98, 0x72, 0x6b, 0xbb, 0x0c, 0x3b, 0x69, 0x22, 0xb2, 0x71, 0x34, 0x03, 0xaf, 0x33, 0x40,
	0x43, 0x27, 0x3c, 0xae, 0x4a, 0xbf, 0x6f, 0xd1, 0xc6, 0xa1, 0xfc, 0x17, 0xa2, 0xb7, 0x18, 0xd7,
	0x89, 0x13, 0xd0, 0xee, 0x15, 0xee, 0x9a, 0x50, 0xda, 0x43, 0x83, 0xf6, 0xb0, 0x17, 0x1e, 0x55,
	0xa5, 0x7f, 0xb0, 0x15, 0x5a, 0x53, 0x5e, 0x03, 0xf4, 0x06, 0x3b, 0x8f, 0xf0, 0x92, 0x83, 0x88,
	0x61, 0xb3, 0x26, 0xf2, 0x34, 0x02, 0x65, 0xde, 0xda, 0xde, 0x5e, 0xb3, 0x73, 0xca, 0x6b, 0x20,
	0x7c, 0xf8, 0x5c, 0x11, 0xb4, 0x5c, 0x11, 0xf4, 0xbd, 0x22, 0xe8, 0x7d, 0x4d, 0x5a, 0xcb, 0x35,
	0x69, 0x7d, 0xad, 0x49, 0xeb, 0x69, 0x34, 0x4d, 0xb2, 0xe7, 0x3c, 0x0a, 0x62, 0x99, 0x32, 0xd3,
	0xf6, 0xc8, 0x34, 0x6f, 0x35, 0x7b, 0x65, 0x7f, 0xbf, 0x54, 0xcc, 0x41, 0x47, 0x5d, 0x53, 0xf9,
	0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xb4, 0x14, 0x3f, 0xbe, 0x01, 0x00, 0x00,
}

func (m *Currency) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Currency) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Currency) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Mintable {
		i--
		if m.Mintable {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTreasury(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Currencies) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Currencies) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Currencies) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintTreasury(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Sequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		i = encodeVarintTreasury(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTreasury(dAtA []byte, offset int, v uint64) int {
	offset -= sovTreasury(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Currency) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTreasury(uint64(l))
	}
	if m.Mintable {
		n += 2
	}
	return n
}

func (m *Currencies) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovTreasury(uint64(l))
		}
	}
	return n
}

func (m *Sequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovTreasury(uint64(m.Number))
	}
	return n
}

func sovTreasury(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTreasury(x uint64) (n int) {
	return sovTreasury(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Currency) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: Currency: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Currency: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mintable", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Mintable = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Currencies) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: Currencies: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Currencies: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
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
				return ErrInvalidLengthTreasury
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTreasury
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Sequence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTreasury
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
			return fmt.Errorf("proto: Sequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTreasury
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTreasury(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTreasury
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTreasury(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
					return 0, ErrIntOverflowTreasury
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
				return 0, ErrInvalidLengthTreasury
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTreasury
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTreasury
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTreasury        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTreasury          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTreasury = fmt.Errorf("proto: unexpected end of group")
)
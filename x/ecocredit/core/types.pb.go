// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/v1/types.proto

package core

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the updatable global parameters of the ecocredit module for
// use with the x/params module.
type Params struct {
	// credit_class_fee is the fixed fee charged on creation of a new credit class
	CreditClassFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=credit_class_fee,json=creditClassFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"credit_class_fee"`
	// basket_fee is the fixed fee charged on creation of a new basket
	BasketFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=basket_fee,json=basketFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"basket_fee"`
	// allowed_class_creators is an allowlist defining the addresses with
	// the required permissions to create credit classes
	AllowedClassCreators []string `protobuf:"bytes,3,rep,name=allowed_class_creators,json=allowedClassCreators,proto3" json:"allowed_class_creators,omitempty"`
	// allowlist_enabled is a param that enables/disables the allowlist for credit
	// creation
	AllowlistEnabled bool `protobuf:"varint,4,opt,name=allowlist_enabled,json=allowlistEnabled,proto3" json:"allowlist_enabled,omitempty"`
	// credit_types is a list of definitions for credit types
	CreditTypes []*CreditType `protobuf:"bytes,5,rep,name=credit_types,json=creditTypes,proto3" json:"credit_types,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b044b6b740b984f, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetCreditClassFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CreditClassFee
	}
	return nil
}

func (m *Params) GetBasketFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.BasketFee
	}
	return nil
}

func (m *Params) GetAllowedClassCreators() []string {
	if m != nil {
		return m.AllowedClassCreators
	}
	return nil
}

func (m *Params) GetAllowlistEnabled() bool {
	if m != nil {
		return m.AllowlistEnabled
	}
	return false
}

func (m *Params) GetCreditTypes() []*CreditType {
	if m != nil {
		return m.CreditTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "regen.ecocredit.v1.Params")
}

func init() { proto.RegisterFile("regen/ecocredit/v1/types.proto", fileDescriptor_7b044b6b740b984f) }

var fileDescriptor_7b044b6b740b984f = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x13, 0x02, 0x15, 0x75, 0x11, 0x2a, 0x51, 0x85, 0x42, 0x0f, 0x6e, 0xc5, 0x29, 0x12,
	0xaa, 0x4d, 0x00, 0x71, 0xa7, 0x11, 0x9c, 0xab, 0x88, 0x13, 0x97, 0xca, 0x71, 0x86, 0x10, 0x9a,
	0xc6, 0x95, 0xed, 0xb6, 0xf4, 0x2d, 0xe0, 0x35, 0x78, 0x92, 0x1e, 0x7b, 0xe4, 0xc4, 0xae, 0xda,
	0x17, 0x59, 0xc5, 0xf6, 0x56, 0x95, 0x76, 0x8f, 0x7b, 0xca, 0xc4, 0x9f, 0x33, 0xdf, 0xfc, 0x19,
	0x84, 0x25, 0x94, 0xd0, 0x50, 0xe0, 0x82, 0x4b, 0x28, 0x2a, 0x4d, 0x37, 0x09, 0xd5, 0xbb, 0x15,
	0x28, 0xb2, 0x92, 0x42, 0x8b, 0x30, 0x34, 0x9c, 0x9c, 0x39, 0xd9, 0x24, 0xc3, 0x41, 0x29, 0x4a,
	0x61, 0x30, 0x6d, 0x2b, 0x7b, 0x73, 0x88, 0xb9, 0x50, 0x4b, 0xa1, 0x68, 0xce, 0x14, 0xd0, 0x4d,
	0x92, 0x83, 0x66, 0x09, 0xe5, 0xa2, 0x6a, 0x6e, 0xf9, 0x3d, 0x26, 0xa5, 0x99, 0x06, 0xcb, 0x5f,
	0xff, 0x09, 0x50, 0x67, 0xc6, 0x24, 0x5b, 0xaa, 0x70, 0x8d, 0xfa, 0xf6, 0xce, 0x9c, 0xd7, 0x4c,
	0xa9, 0xf9, 0x77, 0x80, 0xc8, 0x1f, 0x07, 0x71, 0xef, 0xdd, 0x2b, 0x62, 0x2d, 0xa4, 0xb5, 0x10,
	0x67, 0x21, 0xa9, 0xa8, 0x9a, 0xe9, 0xdb, 0xfd, 0xff, 0x91, 0xf7, 0xf7, 0x6a, 0x14, 0x97, 0x95,
	0xfe, 0xb1, 0xce, 0x09, 0x17, 0x4b, 0xea, 0x46, 0xb2, 0x8f, 0x89, 0x2a, 0x16, 0x2e, 0x5b, 0xfb,
	0x81, 0xca, 0x9e, 0x5b, 0x49, 0xda, 0x3a, 0xbe, 0x00, 0x84, 0x3f, 0x11, 0xca, 0x99, 0x5a, 0x80,
	0x36, 0xc2, 0x47, 0x0f, 0x2f, 0xec, 0xda, 0xf6, 0xad, 0xeb, 0x03, 0x7a, 0xc9, 0xea, 0x5a, 0x6c,
	0xa1, 0x70, 0x19, 0xb9, 0x04, 0xa6, 0x85, 0x54, 0x51, 0x30, 0x0e, 0xe2, 0x6e, 0x36, 0x70, 0xd4,
	0x0c, 0x97, 0x3a, 0x16, 0xbe, 0x41, 0x2f, 0xcc, 0x79, 0x5d, 0x29, 0x3d, 0x87, 0x86, 0xe5, 0x35,
	0x14, 0xd1, 0xe3, 0xb1, 0x1f, 0x3f, 0xcd, 0xfa, 0x67, 0xf0, 0xd9, 0x9e, 0x87, 0x9f, 0xd0, 0x33,
	0xf7, 0x17, 0xcd, 0x0c, 0xd1, 0x13, 0x13, 0x08, 0x93, 0xbb, 0x1b, 0x25, 0xa9, 0xa9, 0xbe, 0xee,
	0x56, 0x90, 0xf5, 0xf8, 0xb9, 0x56, 0xd3, 0xd9, 0xfe, 0x88, 0xfd, 0xc3, 0x11, 0xfb, 0xd7, 0x47,
	0xec, 0xff, 0x3e, 0x61, 0xef, 0x70, 0xc2, 0xde, 0xbf, 0x13, 0xf6, 0xbe, 0x7d, 0xbc, 0x08, 0x6d,
	0x1a, 0x4e, 0x1a, 0xd0, 0x5b, 0x21, 0x17, 0xee, 0xad, 0x86, 0xa2, 0x04, 0x49, 0x7f, 0x5d, 0xec,
	0x9b, 0x0b, 0x09, 0x79, 0xc7, 0x2c, 0xfb, 0xfd, 0x4d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x06, 0x14,
	0xeb, 0x25, 0x78, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CreditTypes) > 0 {
		for iNdEx := len(m.CreditTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.AllowlistEnabled {
		i--
		if m.AllowlistEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.AllowedClassCreators) > 0 {
		for iNdEx := len(m.AllowedClassCreators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedClassCreators[iNdEx])
			copy(dAtA[i:], m.AllowedClassCreators[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.AllowedClassCreators[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BasketFee) > 0 {
		for iNdEx := len(m.BasketFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BasketFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CreditClassFee) > 0 {
		for iNdEx := len(m.CreditClassFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditClassFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CreditClassFee) > 0 {
		for _, e := range m.CreditClassFee {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.BasketFee) > 0 {
		for _, e := range m.BasketFee {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.AllowedClassCreators) > 0 {
		for _, s := range m.AllowedClassCreators {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.AllowlistEnabled {
		n += 2
	}
	if len(m.CreditTypes) > 0 {
		for _, e := range m.CreditTypes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditClassFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditClassFee = append(m.CreditClassFee, types.Coin{})
			if err := m.CreditClassFee[len(m.CreditClassFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketFee = append(m.BasketFee, types.Coin{})
			if err := m.BasketFee[len(m.BasketFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedClassCreators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedClassCreators = append(m.AllowedClassCreators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowlistEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
			m.AllowlistEnabled = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditTypes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditTypes = append(m.CreditTypes, &CreditType{})
			if err := m.CreditTypes[len(m.CreditTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)

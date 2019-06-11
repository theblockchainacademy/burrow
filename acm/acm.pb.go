// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acm.proto

package acm

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	crypto "github.com/hyperledger/burrow/crypto"
	github_com_hyperledger_burrow_crypto "github.com/hyperledger/burrow/crypto"
	permission "github.com/hyperledger/burrow/permission"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Address              github_com_hyperledger_burrow_crypto.Address `protobuf:"bytes,1,opt,name=Address,proto3,customtype=github.com/hyperledger/burrow/crypto.Address" json:"Address"`
	PublicKey            crypto.PublicKey                             `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey"`
	Sequence             uint64                                       `protobuf:"varint,3,opt,name=Sequence,proto3" json:"Sequence,omitempty"`
	Balance              uint64                                       `protobuf:"varint,4,opt,name=Balance,proto3" json:"Balance,omitempty"`
	Code                 Bytecode                                     `protobuf:"bytes,5,opt,name=Code,proto3,customtype=Bytecode" json:"Code"`
	Permissions          permission.AccountPermissions                `protobuf:"bytes,6,opt,name=Permissions,proto3" json:"Permissions"`
	WASM                 Bytecode                                     `protobuf:"bytes,7,opt,name=WASM,proto3,customtype=Bytecode" json:"WASM"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *Account) Reset()      { *m = Account{} }
func (*Account) ProtoMessage() {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_49ed775bc0a6adf6, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetPublicKey() crypto.PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return crypto.PublicKey{}
}

func (m *Account) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetPermissions() permission.AccountPermissions {
	if m != nil {
		return m.Permissions
	}
	return permission.AccountPermissions{}
}

func (*Account) XXX_MessageName() string {
	return "acm.Account"
}
func init() {
	proto.RegisterType((*Account)(nil), "acm.Account")
	golang_proto.RegisterType((*Account)(nil), "acm.Account")
}

func init() { proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }
func init() { golang_proto.RegisterFile("acm.proto", fileDescriptor_49ed775bc0a6adf6) }

var fileDescriptor_49ed775bc0a6adf6 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xeb, 0x36, 0x5f, 0xd3, 0xba, 0x1d, 0xfa, 0x79, 0xb2, 0x3a, 0xa4, 0x05, 0x31, 0x74,
	0x80, 0x44, 0xe2, 0x8f, 0x90, 0xd8, 0x1a, 0x24, 0x16, 0x04, 0xaa, 0xd2, 0x01, 0x89, 0x2d, 0x71,
	0x2e, 0x69, 0xa4, 0xa6, 0x0e, 0x4e, 0x2c, 0x94, 0x37, 0x61, 0xe4, 0x51, 0x18, 0x3b, 0x32, 0x33,
	0x54, 0x55, 0xfb, 0x22, 0x28, 0xc6, 0x2d, 0x59, 0x60, 0xcb, 0xc9, 0xef, 0xde, 0x7b, 0x8e, 0x8e,
	0x71, 0xdb, 0x67, 0x89, 0x9d, 0x0a, 0x9e, 0x73, 0xd2, 0xf0, 0x59, 0xd2, 0x3f, 0x89, 0xe2, 0x7c,
	0x26, 0x03, 0x9b, 0xf1, 0xc4, 0x89, 0x78, 0xc4, 0x1d, 0xc5, 0x02, 0xf9, 0xa4, 0x94, 0x12, 0xea,
	0xeb, 0x7b, 0xa7, 0xdf, 0x4b, 0x41, 0x24, 0x71, 0x96, 0xc5, 0x7c, 0xa1, 0xff, 0x74, 0x99, 0x28,
	0xd2, 0x5c, 0xf3, 0xc3, 0x75, 0x1d, 0x9b, 0x63, 0xc6, 0xb8, 0x5c, 0xe4, 0xe4, 0x1e, 0x9b, 0xe3,
	0x30, 0x14, 0x90, 0x65, 0x14, 0x0d, 0xd1, 0xa8, 0xeb, 0x9e, 0x2f, 0x57, 0x83, 0xda, 0xe7, 0x6a,
	0x70, 0x5c, 0xf1, 0x9c, 0x15, 0x29, 0x88, 0x39, 0x84, 0x11, 0x08, 0x27, 0x90, 0x42, 0xf0, 0x17,
	0x47, 0x1f, 0xd4, 0xbb, 0xde, 0xee, 0x08, 0xb9, 0xc0, 0xed, 0x89, 0x0c, 0xe6, 0x31, 0xbb, 0x85,
	0x82, 0xd6, 0x87, 0x68, 0xd4, 0x39, 0xfd, 0x6f, 0xeb, 0xe1, 0x3d, 0x70, 0x8d, 0xd2, 0xc4, 0xfb,
	0x99, 0x24, 0x7d, 0xdc, 0x9a, 0xc2, 0xb3, 0x84, 0x05, 0x03, 0xda, 0x18, 0xa2, 0x91, 0xe1, 0xed,
	0x35, 0xa1, 0xd8, 0x74, 0xfd, 0xb9, 0x5f, 0x22, 0x43, 0xa1, 0x9d, 0x24, 0x47, 0xd8, 0xb8, 0xe6,
	0x21, 0xd0, 0x7f, 0x2a, 0x79, 0x4f, 0x27, 0x6f, 0xb9, 0x45, 0x0e, 0x8c, 0x87, 0xe0, 0x29, 0x4a,
	0x6e, 0x70, 0x67, 0xb2, 0x2f, 0x24, 0xa3, 0x4d, 0x15, 0xca, 0xb2, 0x2b, 0x25, 0xe9, 0x32, 0x2a,
	0x53, 0x3a, 0x61, 0x75, 0xb1, 0x74, 0x7b, 0x18, 0x4f, 0xef, 0xa8, 0xf9, 0x9b, 0x5b, 0x49, 0xaf,
	0x8c, 0xd7, 0xb7, 0x41, 0xcd, 0xbd, 0x5c, 0x6e, 0x2c, 0xf4, 0xb1, 0xb1, 0xd0, 0x7a, 0x63, 0xa1,
	0xf7, 0xad, 0x85, 0x96, 0x5b, 0x0b, 0x3d, 0x1e, 0xfc, 0xdd, 0xa9, 0xcf, 0x92, 0xa0, 0xa9, 0x9e,
	0xe8, 0xec, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x36, 0x89, 0x52, 0x03, 0x02, 0x00, 0x00,
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Address.Size()))
	n1, err := m.Address.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.PublicKey.Size()))
	n2, err := m.PublicKey.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Sequence != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Sequence))
	}
	if m.Balance != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintAcm(dAtA, i, uint64(m.Balance))
	}
	dAtA[i] = 0x2a
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Code.Size()))
	n3, err := m.Code.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x32
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.Permissions.Size()))
	n4, err := m.Permissions.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x3a
	i++
	i = encodeVarintAcm(dAtA, i, uint64(m.WASM.Size()))
	n5, err := m.WASM.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintAcm(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Address.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.PublicKey.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.Sequence != 0 {
		n += 1 + sovAcm(uint64(m.Sequence))
	}
	if m.Balance != 0 {
		n += 1 + sovAcm(uint64(m.Balance))
	}
	l = m.Code.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.Permissions.Size()
	n += 1 + l + sovAcm(uint64(l))
	l = m.WASM.Size()
	n += 1 + l + sovAcm(uint64(l))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAcm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAcm(x uint64) (n int) {
	return sovAcm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAcm
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			m.Balance = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Balance |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Code.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
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
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Permissions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WASM", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthAcm
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAcm
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WASM.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAcm(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAcm
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAcm
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
func skipAcm(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAcm
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
					return 0, ErrIntOverflowAcm
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAcm
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
				return 0, ErrInvalidLengthAcm
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthAcm
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAcm
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipAcm(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthAcm
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthAcm = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAcm   = fmt.Errorf("proto: integer overflow")
)

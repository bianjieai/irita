// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: irita/wservice/wservice.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type RequestSequence struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" yaml:"key"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
}

func (m *RequestSequence) Reset()         { *m = RequestSequence{} }
func (m *RequestSequence) String() string { return proto.CompactTextString(m) }
func (*RequestSequence) ProtoMessage()    {}
func (*RequestSequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f32f32012b5eee9, []int{0}
}
func (m *RequestSequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestSequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestSequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestSequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestSequence.Merge(m, src)
}
func (m *RequestSequence) XXX_Size() int {
	return m.Size()
}
func (m *RequestSequence) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestSequence.DiscardUnknown(m)
}

var xxx_messageInfo_RequestSequence proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RequestSequence)(nil), "irita.wservice.RequestSequence")
}

func init() { proto.RegisterFile("irita/wservice/wservice.proto", fileDescriptor_6f32f32012b5eee9) }

var fileDescriptor_6f32f32012b5eee9 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0x2c, 0xca, 0x2c,
	0x49, 0xd4, 0x2f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x85, 0x33, 0xf4, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xf8, 0xc0, 0xd2, 0x7a, 0x30, 0x51, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0,
	0x94, 0x3e, 0x88, 0x05, 0x51, 0xa5, 0x94, 0xc8, 0xc5, 0x1f, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c,
	0x12, 0x0c, 0xa2, 0xf2, 0x92, 0x53, 0x85, 0x14, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x9d, 0xf8, 0x3e, 0xdd, 0x93, 0xe7, 0xaa, 0x4c, 0xcc, 0xcd, 0xb1, 0x52, 0xca,
	0x4e, 0xad, 0x54, 0x0a, 0x02, 0x49, 0x09, 0xa9, 0x71, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a,
	0x30, 0x81, 0xd5, 0x08, 0x7c, 0xba, 0x27, 0xcf, 0x03, 0x51, 0x03, 0x16, 0x56, 0x0a, 0x82, 0x48,
	0x5b, 0xb1, 0xbc, 0x58, 0x20, 0xcf, 0xe8, 0xe4, 0x7f, 0xe2, 0xa1, 0x1c, 0xc3, 0x89, 0x47, 0x72,
	0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7,
	0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0x27, 0x65, 0x26, 0xe6, 0x65, 0x65, 0xa6, 0x26, 0x66, 0xea, 0x43, 0xbc, 0x96,
	0x9b, 0x9f, 0x52, 0x9a, 0x93, 0x5a, 0x8c, 0xf0, 0x62, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b,
	0xd8, 0xe9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xeb, 0x64, 0x6f, 0x01, 0x01, 0x00,
	0x00,
}

func (this *RequestSequence) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RequestSequence)
	if !ok {
		that2, ok := that.(RequestSequence)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (m *RequestSequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestSequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestSequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintWservice(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintWservice(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovWservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestSequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovWservice(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovWservice(uint64(l))
	}
	return n
}

func sovWservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWservice(x uint64) (n int) {
	return sovWservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestSequence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWservice
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
			return fmt.Errorf("proto: RequestSequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestSequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWservice
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
				return ErrInvalidLengthWservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWservice
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
				return ErrInvalidLengthWservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWservice
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
func skipWservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWservice
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
					return 0, ErrIntOverflowWservice
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
					return 0, ErrIntOverflowWservice
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
				return 0, ErrInvalidLengthWservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWservice = fmt.Errorf("proto: unexpected end of group")
)

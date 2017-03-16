// Code generated by protoc-gen-gogo.
// source: protobuf_test.proto
// DO NOT EDIT!

/*
	Package agent is a generated protocol buffer package.

	It is generated from these files:
		protobuf_test.proto

	It has these top-level messages:
		ProtoArgs
		ProtoReply
*/
package agent

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ProtoArgs struct {
	A string `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B int32  `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (m *ProtoArgs) Reset()                    { *m = ProtoArgs{} }
func (m *ProtoArgs) String() string            { return proto.CompactTextString(m) }
func (*ProtoArgs) ProtoMessage()               {}
func (*ProtoArgs) Descriptor() ([]byte, []int) { return fileDescriptorProtobufTest, []int{0} }

func (m *ProtoArgs) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *ProtoArgs) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type ProtoReply struct {
	C string `protobuf:"bytes,1,opt,name=C,proto3" json:"C,omitempty"`
	D int32  `protobuf:"varint,2,opt,name=D,proto3" json:"D,omitempty"`
}

func (m *ProtoReply) Reset()                    { *m = ProtoReply{} }
func (m *ProtoReply) String() string            { return proto.CompactTextString(m) }
func (*ProtoReply) ProtoMessage()               {}
func (*ProtoReply) Descriptor() ([]byte, []int) { return fileDescriptorProtobufTest, []int{1} }

func (m *ProtoReply) GetC() string {
	if m != nil {
		return m.C
	}
	return ""
}

func (m *ProtoReply) GetD() int32 {
	if m != nil {
		return m.D
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtoArgs)(nil), "agent.ProtoArgs")
	proto.RegisterType((*ProtoReply)(nil), "agent.ProtoReply")
}
func (m *ProtoArgs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtoArgs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.A) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtobufTest(dAtA, i, uint64(len(m.A)))
		i += copy(dAtA[i:], m.A)
	}
	if m.B != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtobufTest(dAtA, i, uint64(m.B))
	}
	return i, nil
}

func (m *ProtoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtoReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.C) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintProtobufTest(dAtA, i, uint64(len(m.C)))
		i += copy(dAtA[i:], m.C)
	}
	if m.D != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintProtobufTest(dAtA, i, uint64(m.D))
	}
	return i, nil
}

func encodeFixed64ProtobufTest(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ProtobufTest(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintProtobufTest(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ProtoArgs) Size() (n int) {
	var l int
	_ = l
	l = len(m.A)
	if l > 0 {
		n += 1 + l + sovProtobufTest(uint64(l))
	}
	if m.B != 0 {
		n += 1 + sovProtobufTest(uint64(m.B))
	}
	return n
}

func (m *ProtoReply) Size() (n int) {
	var l int
	_ = l
	l = len(m.C)
	if l > 0 {
		n += 1 + l + sovProtobufTest(uint64(l))
	}
	if m.D != 0 {
		n += 1 + sovProtobufTest(uint64(m.D))
	}
	return n
}

func sovProtobufTest(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProtobufTest(x uint64) (n int) {
	return sovProtobufTest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProtoArgs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtobufTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtoArgs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtoArgs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtobufTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.A = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			m.B = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.B |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtobufTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtobufTest
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
func (m *ProtoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProtobufTest
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field C", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProtobufTest
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.C = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field D", wireType)
			}
			m.D = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProtobufTest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.D |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProtobufTest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProtobufTest
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
func skipProtobufTest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProtobufTest
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
					return 0, ErrIntOverflowProtobufTest
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
					return 0, ErrIntOverflowProtobufTest
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthProtobufTest
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProtobufTest
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
				next, err := skipProtobufTest(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthProtobufTest = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProtobufTest   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("protobuf_test.proto", fileDescriptorProtobufTest) }

var fileDescriptorProtobufTest = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0x8b, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0xf3, 0x84, 0x58, 0x13, 0xd3,
	0x53, 0xf3, 0x4a, 0x94, 0xd4, 0xb9, 0x38, 0x03, 0x40, 0x7c, 0xc7, 0xa2, 0xf4, 0x62, 0x21, 0x1e,
	0x2e, 0x46, 0x47, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x46, 0x47, 0x10, 0xcf, 0x49, 0x82,
	0x49, 0x81, 0x51, 0x83, 0x35, 0x88, 0xd1, 0x49, 0x49, 0x83, 0x8b, 0x0b, 0xac, 0x30, 0x28, 0xb5,
	0x20, 0xa7, 0x12, 0x24, 0xe7, 0x0c, 0x53, 0xe9, 0x0c, 0xe2, 0xb9, 0xc0, 0x54, 0xba, 0x38, 0x09,
	0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x24, 0xb1, 0x81, 0xad, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xee, 0xcf, 0xc8,
	0x89, 0x00, 0x00, 0x00,
}

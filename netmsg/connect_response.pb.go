// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: connect_response.proto

/*
	Package netmsg is a generated protocol buffer package.

	It is generated from these files:
		connect_response.proto

	It has these top-level messages:
		ConnectResponse
*/
package netmsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConnectResponse struct {
	ClientSlot int32   `protobuf:"varint,1,opt,name=ClientSlot,proto3" json:"ClientSlot,omitempty"`
	X          float64 `protobuf:"fixed64,2,opt,name=X,proto3" json:"X,omitempty"`
	Y          float64 `protobuf:"fixed64,3,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *ConnectResponse) Reset()                    { *m = ConnectResponse{} }
func (m *ConnectResponse) String() string            { return proto.CompactTextString(m) }
func (*ConnectResponse) ProtoMessage()               {}
func (*ConnectResponse) Descriptor() ([]byte, []int) { return fileDescriptorConnectResponse, []int{0} }

func (m *ConnectResponse) GetClientSlot() int32 {
	if m != nil {
		return m.ClientSlot
	}
	return 0
}

func (m *ConnectResponse) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ConnectResponse) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*ConnectResponse)(nil), "netmsg.ConnectResponse")
}
func (m *ConnectResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClientSlot != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintConnectResponse(dAtA, i, uint64(m.ClientSlot))
	}
	if m.X != 0 {
		dAtA[i] = 0x11
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.X))))
		i += 8
	}
	if m.Y != 0 {
		dAtA[i] = 0x19
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Y))))
		i += 8
	}
	return i, nil
}

func encodeVarintConnectResponse(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConnectResponse) Size() (n int) {
	var l int
	_ = l
	if m.ClientSlot != 0 {
		n += 1 + sovConnectResponse(uint64(m.ClientSlot))
	}
	if m.X != 0 {
		n += 9
	}
	if m.Y != 0 {
		n += 9
	}
	return n
}

func sovConnectResponse(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConnectResponse(x uint64) (n int) {
	return sovConnectResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConnectResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnectResponse
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
			return fmt.Errorf("proto: ConnectResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientSlot", wireType)
			}
			m.ClientSlot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnectResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientSlot |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field X", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.X = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Y", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Y = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipConnectResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConnectResponse
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
func skipConnectResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConnectResponse
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
					return 0, ErrIntOverflowConnectResponse
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
					return 0, ErrIntOverflowConnectResponse
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
				return 0, ErrInvalidLengthConnectResponse
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConnectResponse
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
				next, err := skipConnectResponse(dAtA[start:])
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
	ErrInvalidLengthConnectResponse = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConnectResponse   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("connect_response.proto", fileDescriptorConnectResponse) }

var fileDescriptorConnectResponse = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xcf, 0xcb,
	0x4b, 0x4d, 0x2e, 0x89, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0xcb, 0x4b, 0x2d, 0xc9, 0x2d, 0x4e, 0x57, 0xf2, 0xe5, 0xe2, 0x77, 0x86,
	0xa8, 0x08, 0x82, 0x2a, 0x10, 0x92, 0xe3, 0xe2, 0x72, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x09, 0xce,
	0xc9, 0x2f, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x42, 0x12, 0x11, 0xe2, 0xe1, 0x62, 0x8c,
	0x90, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0c, 0x62, 0x8c, 0x00, 0xf1, 0x22, 0x25, 0x98, 0x21, 0xbc,
	0x48, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0x7d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03,
	0x94, 0x49, 0x55, 0x89, 0x00, 0x00, 0x00,
}

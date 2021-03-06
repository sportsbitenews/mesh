// Code generated by protoc-gen-gogo.
// source: github.com/evanphx/mesh/grpc/frame.proto
// DO NOT EDIT!

/*
	Package grpc is a generated protocol buffer package.

	It is generated from these files:
		github.com/evanphx/mesh/grpc/frame.proto

	It has these top-level messages:
		Frame
*/
package grpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import reflect "reflect"

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

type Frame_Type int32

const (
	ERROR      Frame_Type = 0
	REQUEST    Frame_Type = 1
	REPLY      Frame_Type = 2
	DATA       Frame_Type = 3
	CLOSE_SEND Frame_Type = 4
)

var Frame_Type_name = map[int32]string{
	0: "ERROR",
	1: "REQUEST",
	2: "REPLY",
	3: "DATA",
	4: "CLOSE_SEND",
}
var Frame_Type_value = map[string]int32{
	"ERROR":      0,
	"REQUEST":    1,
	"REPLY":      2,
	"DATA":       3,
	"CLOSE_SEND": 4,
}

func (Frame_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorFrame, []int{0, 0} }

type Frame struct {
	Type   Frame_Type `protobuf:"varint,1,opt,name=type,proto3,enum=grpc.Frame_Type" json:"type,omitempty"`
	Method string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Body   []byte     `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptorFrame, []int{0} }

func init() {
	proto.RegisterType((*Frame)(nil), "grpc.Frame")
	proto.RegisterEnum("grpc.Frame_Type", Frame_Type_name, Frame_Type_value)
}
func (x Frame_Type) String() string {
	s, ok := Frame_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Frame) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Frame)
	if !ok {
		that2, ok := that.(Frame)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Method != that1.Method {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	return true
}
func (this *Frame) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&grpc.Frame{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Method: "+fmt.Sprintf("%#v", this.Method)+",\n")
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFrame(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringFrame(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *Frame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Frame) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFrame(dAtA, i, uint64(m.Type))
	}
	if len(m.Method) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFrame(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFrame(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func encodeFixed64Frame(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Frame(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintFrame(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Frame) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovFrame(uint64(m.Type))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovFrame(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovFrame(uint64(l))
	}
	return n
}

func sovFrame(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFrame(x uint64) (n int) {
	return sovFrame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Frame) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Frame{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFrame(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Frame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrame
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
			return fmt.Errorf("proto: Frame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Frame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Frame_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrame
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
				return ErrInvalidLengthFrame
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFrame
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFrame
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
func skipFrame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFrame
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
					return 0, ErrIntOverflowFrame
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
					return 0, ErrIntOverflowFrame
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
				return 0, ErrInvalidLengthFrame
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFrame
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
				next, err := skipFrame(dAtA[start:])
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
	ErrInvalidLengthFrame = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFrame   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/evanphx/mesh/grpc/frame.proto", fileDescriptorFrame) }

var fileDescriptorFrame = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2d, 0x4b, 0xcc, 0x2b, 0xc8, 0xa8, 0xd0, 0xcf,
	0x4d, 0x2d, 0xce, 0xd0, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0x2b, 0x4a, 0xcc, 0x4d, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0x89, 0x28, 0xcd, 0x63, 0xe4, 0x62, 0x75, 0x03, 0x89,
	0x0a, 0xa9, 0x70, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x09,
	0xe8, 0x81, 0xa4, 0xf5, 0xc0, 0x52, 0x7a, 0x21, 0x95, 0x05, 0xa9, 0x41, 0x60, 0x59, 0x21, 0x31,
	0x2e, 0xb6, 0xdc, 0xd4, 0x92, 0x8c, 0xfc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28,
	0x4f, 0x48, 0x88, 0x8b, 0x25, 0x29, 0x3f, 0xa5, 0x52, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27, 0x08,
	0xcc, 0x56, 0x72, 0xe6, 0x62, 0x01, 0xe9, 0x14, 0xe2, 0xe4, 0x62, 0x75, 0x0d, 0x0a, 0xf2, 0x0f,
	0x12, 0x60, 0x10, 0xe2, 0xe6, 0x62, 0x0f, 0x72, 0x0d, 0x0c, 0x75, 0x0d, 0x0e, 0x11, 0x60, 0x04,
	0x89, 0x07, 0xb9, 0x06, 0xf8, 0x44, 0x0a, 0x30, 0x09, 0x71, 0x70, 0xb1, 0xb8, 0x38, 0x86, 0x38,
	0x0a, 0x30, 0x0b, 0xf1, 0x71, 0x71, 0x39, 0xfb, 0xf8, 0x07, 0xbb, 0xc6, 0x07, 0xbb, 0xfa, 0xb9,
	0x08, 0xb0, 0x38, 0xe9, 0x5c, 0x78, 0x28, 0xc7, 0x70, 0xe3, 0xa1, 0x1c, 0xc3, 0x87, 0x87, 0x72,
	0x8c, 0x0d, 0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x5f, 0x3c, 0x92, 0x63, 0xf8, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x24, 0x36, 0xb0, 0xdf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xc1,
	0xf5, 0xdc, 0x07, 0x01, 0x00, 0x00,
}

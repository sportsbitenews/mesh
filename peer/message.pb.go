// Code generated by protoc-gen-gogo.
// source: github.com/evanphx/mesh/peer/message.proto
// DO NOT EDIT!

package peer

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

type Header_Type int32

const (
	PIPE_OPEN    Header_Type = 0
	PIPE_CLOSE   Header_Type = 1
	PIPE_DATA    Header_Type = 2
	PIPE_UNKNOWN Header_Type = 3
	PIPE_OPENED  Header_Type = 4
	PING         Header_Type = 5
	PONG         Header_Type = 6
)

var Header_Type_name = map[int32]string{
	0: "PIPE_OPEN",
	1: "PIPE_CLOSE",
	2: "PIPE_DATA",
	3: "PIPE_UNKNOWN",
	4: "PIPE_OPENED",
	5: "PING",
	6: "PONG",
}
var Header_Type_value = map[string]int32{
	"PIPE_OPEN":    0,
	"PIPE_CLOSE":   1,
	"PIPE_DATA":    2,
	"PIPE_UNKNOWN": 3,
	"PIPE_OPENED":  4,
	"PING":         5,
	"PONG":         6,
}

func (Header_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0, 0} }

type Header struct {
	Sender      []byte      `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Destination []byte      `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Session     uint64      `protobuf:"varint,3,opt,name=session,proto3" json:"session,omitempty"`
	Type        Header_Type `protobuf:"varint,4,opt,name=type,proto3,enum=peer.Header_Type" json:"type,omitempty"`
	PipeName    string      `protobuf:"bytes,5,opt,name=pipe_name,json=pipeName,proto3" json:"pipe_name,omitempty"`
	Body        []byte      `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"`
	Encrypted   bool        `protobuf:"varint,7,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

func init() {
	proto.RegisterType((*Header)(nil), "peer.Header")
	proto.RegisterEnum("peer.Header_Type", Header_Type_name, Header_Type_value)
}
func (x Header_Type) String() string {
	s, ok := Header_Type_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Header) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Header)
	if !ok {
		that2, ok := that.(Header)
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
	if !bytes.Equal(this.Sender, that1.Sender) {
		return false
	}
	if !bytes.Equal(this.Destination, that1.Destination) {
		return false
	}
	if this.Session != that1.Session {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.PipeName != that1.PipeName {
		return false
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	if this.Encrypted != that1.Encrypted {
		return false
	}
	return true
}
func (this *Header) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&peer.Header{")
	s = append(s, "Sender: "+fmt.Sprintf("%#v", this.Sender)+",\n")
	s = append(s, "Destination: "+fmt.Sprintf("%#v", this.Destination)+",\n")
	s = append(s, "Session: "+fmt.Sprintf("%#v", this.Session)+",\n")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "PipeName: "+fmt.Sprintf("%#v", this.PipeName)+",\n")
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "Encrypted: "+fmt.Sprintf("%#v", this.Encrypted)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringMessage(m github_com_gogo_protobuf_proto.Message) string {
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
func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sender) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Sender)))
		i += copy(dAtA[i:], m.Sender)
	}
	if len(m.Destination) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Destination)))
		i += copy(dAtA[i:], m.Destination)
	}
	if m.Session != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Session))
	}
	if m.Type != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Type))
	}
	if len(m.PipeName) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.PipeName)))
		i += copy(dAtA[i:], m.PipeName)
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	if m.Encrypted {
		dAtA[i] = 0x38
		i++
		if m.Encrypted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Message(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Message(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Header) Size() (n int) {
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Destination)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Session != 0 {
		n += 1 + sovMessage(uint64(m.Session))
	}
	if m.Type != 0 {
		n += 1 + sovMessage(uint64(m.Type))
	}
	l = len(m.PipeName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Encrypted {
		n += 2
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Header) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Header{`,
		`Sender:` + fmt.Sprintf("%v", this.Sender) + `,`,
		`Destination:` + fmt.Sprintf("%v", this.Destination) + `,`,
		`Session:` + fmt.Sprintf("%v", this.Session) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`PipeName:` + fmt.Sprintf("%v", this.PipeName) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`Encrypted:` + fmt.Sprintf("%v", this.Encrypted) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Destination", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Destination = append(m.Destination[:0], dAtA[iNdEx:postIndex]...)
			if m.Destination == nil {
				m.Destination = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Session", wireType)
			}
			m.Session = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Session |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Header_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PipeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PipeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
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
				return ErrInvalidLengthMessage
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encrypted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Encrypted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
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
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
					return 0, ErrIntOverflowMessage
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
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
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
				next, err := skipMessage(dAtA[start:])
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
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("github.com/evanphx/mesh/peer/message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x33, 0xed, 0x34, 0x6d, 0xbe, 0xd6, 0x3a, 0xce, 0x41, 0x02, 0xca, 0x10, 0x0a, 0x42,
	0x10, 0x49, 0x41, 0x9f, 0xa0, 0xda, 0x50, 0x8b, 0x92, 0x84, 0x58, 0xf1, 0x58, 0xd2, 0xe6, 0xa3,
	0x0d, 0x92, 0x64, 0x48, 0xa2, 0x98, 0x9b, 0x8f, 0xe0, 0x63, 0xf8, 0x16, 0x5e, 0x3d, 0xf6, 0xb8,
	0xc7, 0x6d, 0xf6, 0xb2, 0xc7, 0x3e, 0xc2, 0x92, 0xd9, 0xdd, 0xee, 0xde, 0xfe, 0xff, 0xdf, 0x8f,
	0x99, 0x3f, 0x7c, 0xf0, 0x7a, 0x97, 0x54, 0xfb, 0x1f, 0x1b, 0x67, 0x9b, 0xa7, 0x53, 0xfc, 0x19,
	0x65, 0x72, 0xff, 0x6b, 0x9a, 0x62, 0xb9, 0x9f, 0x4a, 0xc4, 0xa2, 0x4d, 0x65, 0xb4, 0x43, 0x47,
	0x16, 0x79, 0x95, 0x73, 0xda, 0xb2, 0xc9, 0xbf, 0x0e, 0xe8, 0x1f, 0x31, 0x8a, 0xb1, 0xe0, 0xcf,
	0x41, 0x2f, 0x31, 0x8b, 0xb1, 0x30, 0x89, 0x45, 0xec, 0x51, 0x78, 0xd7, 0xb8, 0x05, 0xc3, 0x18,
	0xcb, 0x2a, 0xc9, 0xa2, 0x2a, 0xc9, 0x33, 0xb3, 0xa3, 0xe4, 0x63, 0xc4, 0x4d, 0xe8, 0x97, 0x58,
	0x96, 0xad, 0xed, 0x5a, 0xc4, 0xa6, 0xe1, 0x7d, 0xe5, 0xaf, 0x80, 0x56, 0xb5, 0x44, 0x93, 0x5a,
	0xc4, 0x1e, 0xbf, 0x7d, 0xe6, 0xb4, 0x9b, 0xce, 0xed, 0x9e, 0xb3, 0xaa, 0x25, 0x86, 0x4a, 0xf3,
	0x17, 0x60, 0xc8, 0x44, 0xe2, 0x3a, 0x8b, 0x52, 0x34, 0x7b, 0x16, 0xb1, 0x8d, 0x70, 0xd0, 0x02,
	0x2f, 0x4a, 0x91, 0x73, 0xa0, 0x9b, 0x3c, 0xae, 0x4d, 0x5d, 0x0d, 0xab, 0xcc, 0x5f, 0x82, 0x81,
	0xd9, 0xb6, 0xa8, 0x65, 0x85, 0xb1, 0xd9, 0xb7, 0x88, 0x3d, 0x08, 0x1f, 0xc0, 0xe4, 0x3b, 0xd0,
	0xf6, 0x73, 0xfe, 0x04, 0x8c, 0x60, 0x19, 0xb8, 0x6b, 0x3f, 0x70, 0x3d, 0xa6, 0xf1, 0x31, 0x80,
	0xaa, 0x1f, 0x3e, 0xfb, 0x5f, 0x5c, 0x46, 0xce, 0x7a, 0x3e, 0x5b, 0xcd, 0x58, 0x87, 0x33, 0x18,
	0xa9, 0xfa, 0xd5, 0xfb, 0xe4, 0xf9, 0xdf, 0x3c, 0xd6, 0xe5, 0x4f, 0x61, 0x78, 0x7e, 0xef, 0xce,
	0x19, 0xe5, 0x03, 0xa0, 0xc1, 0xd2, 0x5b, 0xb0, 0x9e, 0x4a, 0xbe, 0xb7, 0x60, 0xfa, 0xfb, 0x37,
	0x87, 0xa3, 0xd0, 0x2e, 0x8e, 0x42, 0x3b, 0x1d, 0x05, 0xf9, 0xdd, 0x08, 0xf2, 0xb7, 0x11, 0xe4,
	0x7f, 0x23, 0xc8, 0xa1, 0x11, 0xe4, 0xb2, 0x11, 0xe4, 0xba, 0x11, 0xda, 0xa9, 0x11, 0xe4, 0xcf,
	0x95, 0xd0, 0x36, 0xba, 0x3a, 0xfe, 0xbb, 0x9b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x99, 0x92,
	0x4e, 0xaa, 0x01, 0x00, 0x00,
}

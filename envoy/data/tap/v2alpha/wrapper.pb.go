// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/data/tap/v2alpha/wrapper.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

// Wrapper for all fully buffered tap traces that Envoy emits. This is required for sending traces
// over gRPC APIs or more easily persisting binary messages to files.
type BufferedTraceWrapper struct {
	// Types that are valid to be assigned to Trace:
	//	*BufferedTraceWrapper_HttpBufferedTrace
	Trace                isBufferedTraceWrapper_Trace `protobuf_oneof:"trace"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *BufferedTraceWrapper) Reset()         { *m = BufferedTraceWrapper{} }
func (m *BufferedTraceWrapper) String() string { return proto.CompactTextString(m) }
func (*BufferedTraceWrapper) ProtoMessage()    {}
func (*BufferedTraceWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_wrapper_ad6206363d8f9e37, []int{0}
}
func (m *BufferedTraceWrapper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BufferedTraceWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BufferedTraceWrapper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *BufferedTraceWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BufferedTraceWrapper.Merge(dst, src)
}
func (m *BufferedTraceWrapper) XXX_Size() int {
	return m.Size()
}
func (m *BufferedTraceWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_BufferedTraceWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_BufferedTraceWrapper proto.InternalMessageInfo

type isBufferedTraceWrapper_Trace interface {
	isBufferedTraceWrapper_Trace()
	MarshalTo([]byte) (int, error)
	Size() int
}

type BufferedTraceWrapper_HttpBufferedTrace struct {
	HttpBufferedTrace *HttpBufferedTrace `protobuf:"bytes,1,opt,name=http_buffered_trace,json=httpBufferedTrace,proto3,oneof"`
}

func (*BufferedTraceWrapper_HttpBufferedTrace) isBufferedTraceWrapper_Trace() {}

func (m *BufferedTraceWrapper) GetTrace() isBufferedTraceWrapper_Trace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func (m *BufferedTraceWrapper) GetHttpBufferedTrace() *HttpBufferedTrace {
	if x, ok := m.GetTrace().(*BufferedTraceWrapper_HttpBufferedTrace); ok {
		return x.HttpBufferedTrace
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BufferedTraceWrapper) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BufferedTraceWrapper_OneofMarshaler, _BufferedTraceWrapper_OneofUnmarshaler, _BufferedTraceWrapper_OneofSizer, []interface{}{
		(*BufferedTraceWrapper_HttpBufferedTrace)(nil),
	}
}

func _BufferedTraceWrapper_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BufferedTraceWrapper)
	// trace
	switch x := m.Trace.(type) {
	case *BufferedTraceWrapper_HttpBufferedTrace:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpBufferedTrace); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("BufferedTraceWrapper.Trace has unexpected type %T", x)
	}
	return nil
}

func _BufferedTraceWrapper_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BufferedTraceWrapper)
	switch tag {
	case 1: // trace.http_buffered_trace
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpBufferedTrace)
		err := b.DecodeMessage(msg)
		m.Trace = &BufferedTraceWrapper_HttpBufferedTrace{msg}
		return true, err
	default:
		return false, nil
	}
}

func _BufferedTraceWrapper_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BufferedTraceWrapper)
	// trace
	switch x := m.Trace.(type) {
	case *BufferedTraceWrapper_HttpBufferedTrace:
		s := proto.Size(x.HttpBufferedTrace)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*BufferedTraceWrapper)(nil), "envoy.data.tap.v2alpha.BufferedTraceWrapper")
}
func (m *BufferedTraceWrapper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BufferedTraceWrapper) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Trace != nil {
		nn1, err := m.Trace.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *BufferedTraceWrapper_HttpBufferedTrace) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.HttpBufferedTrace != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintWrapper(dAtA, i, uint64(m.HttpBufferedTrace.Size()))
		n2, err := m.HttpBufferedTrace.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func encodeVarintWrapper(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BufferedTraceWrapper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Trace != nil {
		n += m.Trace.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *BufferedTraceWrapper_HttpBufferedTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpBufferedTrace != nil {
		l = m.HttpBufferedTrace.Size()
		n += 1 + l + sovWrapper(uint64(l))
	}
	return n
}

func sovWrapper(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWrapper(x uint64) (n int) {
	return sovWrapper(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BufferedTraceWrapper) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWrapper
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
			return fmt.Errorf("proto: BufferedTraceWrapper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BufferedTraceWrapper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpBufferedTrace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWrapper
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWrapper
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &HttpBufferedTrace{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Trace = &BufferedTraceWrapper_HttpBufferedTrace{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWrapper(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWrapper
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
func skipWrapper(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWrapper
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
					return 0, ErrIntOverflowWrapper
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
					return 0, ErrIntOverflowWrapper
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
				return 0, ErrInvalidLengthWrapper
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWrapper
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
				next, err := skipWrapper(dAtA[start:])
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
	ErrInvalidLengthWrapper = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWrapper   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/data/tap/v2alpha/wrapper.proto", fileDescriptor_wrapper_ad6206363d8f9e37)
}

var fileDescriptor_wrapper_ad6206363d8f9e37 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0x49, 0x2c, 0x49, 0xd4, 0x2f, 0x49, 0x2c, 0xd0, 0x2f, 0x33, 0x4a, 0xcc, 0x29,
	0xc8, 0x48, 0xd4, 0x2f, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x03, 0xab, 0xd2, 0x03, 0xa9, 0xd2, 0x2b, 0x49, 0x2c, 0xd0, 0x83, 0xaa, 0x92, 0x52,
	0xc4, 0xa1, 0x3b, 0xa3, 0xa4, 0xa4, 0x00, 0xa2, 0x55, 0x4a, 0xbc, 0x2c, 0x31, 0x27, 0x33, 0x25,
	0xb1, 0x24, 0x55, 0x1f, 0xc6, 0x80, 0x48, 0x28, 0x35, 0x33, 0x72, 0x89, 0x38, 0x95, 0xa6, 0xa5,
	0xa5, 0x16, 0xa5, 0xa6, 0x84, 0x14, 0x25, 0x26, 0xa7, 0x86, 0x43, 0xac, 0x14, 0x8a, 0xe6, 0x12,
	0x06, 0xe9, 0x8f, 0x4f, 0x82, 0x4a, 0xc6, 0x97, 0x80, 0x64, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8,
	0x8d, 0x34, 0xf5, 0xb0, 0x3b, 0x45, 0xcf, 0xa3, 0xa4, 0xa4, 0x00, 0xc5, 0x38, 0x0f, 0x86, 0x20,
	0xc1, 0x0c, 0x74, 0x41, 0x27, 0x3e, 0x2e, 0x56, 0xb0, 0x71, 0x42, 0xac, 0x3b, 0x5e, 0x1e, 0x60,
	0x66, 0x74, 0x32, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0xb9, 0x54, 0x32, 0xf3, 0x21, 0xe6, 0x17, 0x14, 0xe5, 0x57, 0x54, 0xe2, 0xb0, 0x2a, 0x89, 0x0d,
	0xec, 0x01, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x33, 0x96, 0xf1, 0x3c, 0x01, 0x00,
	0x00,
}

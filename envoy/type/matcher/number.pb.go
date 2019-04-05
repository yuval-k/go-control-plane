// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/number.proto

package matcher

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Specifies the way to match a double value.
type DoubleMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*DoubleMatcher_Range
	//	*DoubleMatcher_Exact
	MatchPattern         isDoubleMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DoubleMatcher) Reset()         { *m = DoubleMatcher{} }
func (m *DoubleMatcher) String() string { return proto.CompactTextString(m) }
func (*DoubleMatcher) ProtoMessage()    {}
func (*DoubleMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_58ad3770a33c5fc2, []int{0}
}
func (m *DoubleMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DoubleMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DoubleMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DoubleMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleMatcher.Merge(m, src)
}
func (m *DoubleMatcher) XXX_Size() int {
	return m.Size()
}
func (m *DoubleMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleMatcher proto.InternalMessageInfo

type isDoubleMatcher_MatchPattern interface {
	isDoubleMatcher_MatchPattern()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DoubleMatcher_Range struct {
	Range *_type.DoubleRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}
type DoubleMatcher_Exact struct {
	Exact float64 `protobuf:"fixed64,2,opt,name=exact,proto3,oneof"`
}

func (*DoubleMatcher_Range) isDoubleMatcher_MatchPattern() {}
func (*DoubleMatcher_Exact) isDoubleMatcher_MatchPattern() {}

func (m *DoubleMatcher) GetMatchPattern() isDoubleMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *DoubleMatcher) GetRange() *_type.DoubleRange {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Range); ok {
		return x.Range
	}
	return nil
}

func (m *DoubleMatcher) GetExact() float64 {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Exact); ok {
		return x.Exact
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DoubleMatcher) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DoubleMatcher_OneofMarshaler, _DoubleMatcher_OneofUnmarshaler, _DoubleMatcher_OneofSizer, []interface{}{
		(*DoubleMatcher_Range)(nil),
		(*DoubleMatcher_Exact)(nil),
	}
}

func _DoubleMatcher_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DoubleMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *DoubleMatcher_Range:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Range); err != nil {
			return err
		}
	case *DoubleMatcher_Exact:
		_ = b.EncodeVarint(2<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.Exact))
	case nil:
	default:
		return fmt.Errorf("DoubleMatcher.MatchPattern has unexpected type %T", x)
	}
	return nil
}

func _DoubleMatcher_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DoubleMatcher)
	switch tag {
	case 1: // match_pattern.range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(_type.DoubleRange)
		err := b.DecodeMessage(msg)
		m.MatchPattern = &DoubleMatcher_Range{msg}
		return true, err
	case 2: // match_pattern.exact
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.MatchPattern = &DoubleMatcher_Exact{math.Float64frombits(x)}
		return true, err
	default:
		return false, nil
	}
}

func _DoubleMatcher_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DoubleMatcher)
	// match_pattern
	switch x := m.MatchPattern.(type) {
	case *DoubleMatcher_Range:
		s := proto.Size(x.Range)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DoubleMatcher_Exact:
		n += 1 // tag and wire
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*DoubleMatcher)(nil), "envoy.type.matcher.DoubleMatcher")
}

func init() { proto.RegisterFile("envoy/type/matcher/number.proto", fileDescriptor_58ad3770a33c5fc2) }

var fileDescriptor_58ad3770a33c5fc2 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0xcf,
	0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0x2b, 0xd0,
	0x03, 0x29, 0xd0, 0x83, 0x2a, 0x90, 0x12, 0x43, 0xd2, 0x54, 0x94, 0x98, 0x97, 0x9e, 0x0a, 0x51,
	0x2b, 0x25, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0xaa, 0x0f, 0x63, 0x40, 0x24, 0x94,
	0x0a, 0xb8, 0x78, 0x5d, 0xf2, 0x4b, 0x93, 0x72, 0x52, 0x7d, 0x21, 0x26, 0x08, 0xe9, 0x73, 0xb1,
	0x82, 0x35, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xeb, 0x21, 0xd9, 0x02, 0x51, 0x19,
	0x04, 0x92, 0xf6, 0x60, 0x08, 0x82, 0xa8, 0x13, 0x12, 0xe3, 0x62, 0x4d, 0xad, 0x48, 0x4c, 0x2e,
	0x91, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x04, 0x89, 0x83, 0xb9, 0x4e, 0x62, 0x5c, 0xbc, 0x60, 0x57,
	0xc5, 0x17, 0x24, 0x96, 0x94, 0xa4, 0x16, 0xe5, 0x09, 0xb1, 0xee, 0x78, 0x79, 0x80, 0x99, 0xd1,
	0xc9, 0xed, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0xe4, 0x52,
	0xc8, 0xcc, 0x87, 0xd8, 0x50, 0x50, 0x94, 0x5f, 0x51, 0xa9, 0x87, 0xe9, 0x25, 0x27, 0x6e, 0x3f,
	0xb0, 0xa7, 0x03, 0x40, 0xce, 0x0d, 0x60, 0x8c, 0x62, 0x87, 0x8a, 0x27, 0xb1, 0x81, 0x3d, 0x60,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x22, 0x18, 0x20, 0x39, 0x28, 0x01, 0x00, 0x00,
}

func (m *DoubleMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DoubleMatcher) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.MatchPattern != nil {
		nn1, err := m.MatchPattern.MarshalTo(dAtA[i:])
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

func (m *DoubleMatcher_Range) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Range != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNumber(dAtA, i, uint64(m.Range.Size()))
		n2, err := m.Range.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *DoubleMatcher_Exact) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x11
	i++
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Exact))))
	i += 8
	return i, nil
}
func encodeVarintNumber(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DoubleMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MatchPattern != nil {
		n += m.MatchPattern.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DoubleMatcher_Range) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Range != nil {
		l = m.Range.Size()
		n += 1 + l + sovNumber(uint64(l))
	}
	return n
}
func (m *DoubleMatcher_Exact) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	return n
}

func sovNumber(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNumber(x uint64) (n int) {
	return sovNumber(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DoubleMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNumber
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
			return fmt.Errorf("proto: DoubleMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DoubleMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Range", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNumber
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
				return ErrInvalidLengthNumber
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNumber
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &_type.DoubleRange{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MatchPattern = &DoubleMatcher_Range{v}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exact", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.MatchPattern = &DoubleMatcher_Exact{float64(math.Float64frombits(v))}
		default:
			iNdEx = preIndex
			skippy, err := skipNumber(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNumber
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNumber
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
func skipNumber(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNumber
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
					return 0, ErrIntOverflowNumber
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
					return 0, ErrIntOverflowNumber
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
				return 0, ErrInvalidLengthNumber
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNumber
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNumber
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
				next, err := skipNumber(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNumber
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
	ErrInvalidLengthNumber = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNumber   = fmt.Errorf("proto: integer overflow")
)

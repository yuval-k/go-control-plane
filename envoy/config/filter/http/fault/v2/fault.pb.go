// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v2/fault.proto

package v2

import (
	fmt "fmt"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v2"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage           *_type.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{0}
}
func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return m.Size()
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultAbort_OneofMarshaler, _FaultAbort_OneofUnmarshaler, _FaultAbort_OneofSizer, []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

func _FaultAbort_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultAbort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *FaultAbort_HttpStatus:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		return fmt.Errorf("FaultAbort.ErrorType has unexpected type %T", x)
	}
	return nil
}

func _FaultAbort_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultAbort)
	switch tag {
	case 2: // error_type.http_status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ErrorType = &FaultAbort_HttpStatus{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _FaultAbort_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultAbort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *FaultAbort_HttpStatus:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object.
	Delay *v2.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percentage
	// <envoy_api_field_config.filter.http.fault.v2.FaultAbort.percentage>` field.
	// The filter will check the request's headers against all the specified
	// headers in the filter config. A match will happen if all the headers in the
	// config are present in the request with the same values (or based on
	// presence if the *value* field is not in the config).
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	// The maximum number of faults that can be active at a single time via the configured fault
	// filter. Note that because this setting can be overridden at the route level, it's possible
	// for the number of active faults to be greater than this value (if injected via a different
	// route). If not specified, defaults to unlimited. This setting can be overridden via
	// `runtime <config_http_filters_fault_injection_runtime>` and any faults that are not injected
	// due to overflow will be indicated via the `faults_overflow
	// <config_http_filters_fault_injection_stats>` stat.
	//
	// .. attention::
	//   Like other :ref:`circuit breakers <arch_overview_circuit_break>` in Envoy, this is a fuzzy
	//   limit. It's possible for the number of active faults to rise slightly above the configured
	//   amount due to the implementation details.
	MaxActiveFaults *types.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	// The response rate limit to be applied to the response body of the stream. When configured,
	// the percentage can be overridden by the :ref:`fault.http.rate_limit.response_percent
	// <config_http_filters_fault_injection_runtime>` runtime key.
	//
	// .. attention::
	//  This is a per-stream limit versus a connection level limit. This means that concurrent streams
	//  will each get an independent limit.
	ResponseRateLimit    *v2.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_26070db6b6576d5c, []int{1}
}
func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return m.Size()
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *types.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v2.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v2.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v2/fault.proto", fileDescriptor_26070db6b6576d5c)
}

var fileDescriptor_26070db6b6576d5c = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x4d, 0xdb, 0x6d, 0xae, 0xa6, 0x75, 0xd9, 0x81, 0xa8, 0x82, 0xaa, 0xdb, 0x29, 0x20,
	0xe6, 0xa0, 0xec, 0x88, 0x40, 0x5a, 0x87, 0xa6, 0x82, 0xf8, 0x51, 0x85, 0xc1, 0x01, 0x21, 0x45,
	0x6e, 0xf3, 0xda, 0x46, 0x4a, 0x63, 0xcb, 0x76, 0xb2, 0xf6, 0x3f, 0xe1, 0x4f, 0x41, 0x9c, 0x76,
	0x9c, 0xc4, 0x85, 0x23, 0x47, 0xd4, 0xdb, 0xfe, 0x0b, 0x64, 0xbb, 0x61, 0x1c, 0x2a, 0x8d, 0x4b,
	0xe4, 0x3c, 0x7f, 0xdf, 0xfb, 0xde, 0xf7, 0xf9, 0xe1, 0x63, 0xc8, 0x4b, 0xb6, 0x0c, 0xc6, 0x2c,
	0x9f, 0xa4, 0xd3, 0x60, 0x92, 0x66, 0x0a, 0x44, 0x30, 0x53, 0x8a, 0x07, 0x13, 0x5a, 0x64, 0x2a,
	0x28, 0x43, 0x7b, 0x20, 0x5c, 0x30, 0xc5, 0xdc, 0x43, 0x03, 0x27, 0x16, 0x4e, 0x2c, 0x9c, 0x68,
	0x38, 0xb1, 0xa8, 0x32, 0xec, 0x74, 0x6d, 0x47, 0xca, 0x53, 0x4d, 0x16, 0xac, 0x50, 0x60, 0xbf,
	0xb6, 0x45, 0xc7, 0xdf, 0xa4, 0xb8, 0x49, 0xac, 0xe3, 0x59, 0xa4, 0x5a, 0x72, 0x08, 0x38, 0x88,
	0x31, 0xe4, 0xd5, 0x4d, 0x77, 0xca, 0xd8, 0x34, 0x83, 0xc0, 0xfc, 0x8d, 0x8a, 0x49, 0x70, 0x29,
	0x28, 0xe7, 0x20, 0xe4, 0xfa, 0xfe, 0x7e, 0x49, 0xb3, 0x34, 0xa1, 0x0a, 0x82, 0xea, 0x60, 0x2f,
	0x8e, 0xbe, 0x22, 0x8c, 0xcf, 0xb5, 0xc4, 0xe9, 0x88, 0x09, 0xe5, 0x3e, 0xc5, 0x2d, 0x3d, 0x7c,
	0x2c, 0x15, 0x55, 0x85, 0xf4, 0x6a, 0x3d, 0xe4, 0xef, 0xf6, 0x77, 0xbf, 0xdf, 0x5c, 0x39, 0xdb,
	0x8f, 0x9b, 0xed, 0x5f, 0x75, 0xff, 0x1a, 0x0d, 0xee, 0x45, 0x58, 0x63, 0x3e, 0x18, 0x88, 0xfb,
	0x1c, 0xe3, 0xf5, 0x28, 0x74, 0x0a, 0x9e, 0xd3, 0x43, 0x7e, 0x2b, 0x7c, 0x48, 0x6c, 0x2a, 0x7a,
	0x50, 0x72, 0x2e, 0xe8, 0x58, 0xa5, 0x2c, 0xa7, 0xd9, 0xd0, 0xe2, 0xa2, 0x7f, 0x08, 0xfd, 0x03,
	0x8c, 0x41, 0x08, 0x26, 0x62, 0x8d, 0x75, 0x1b, 0xdf, 0x6e, 0xae, 0x1c, 0xf4, 0xba, 0xbe, 0x8d,
	0xda, 0xb5, 0xa3, 0x1f, 0x0e, 0xde, 0x19, 0x5c, 0x5c, 0x0c, 0xcd, 0x78, 0xee, 0x0b, 0xdc, 0x48,
	0x20, 0xa3, 0x4b, 0x0f, 0x19, 0x09, 0x9f, 0x6c, 0x0a, 0xbe, 0xca, 0x9c, 0x18, 0xce, 0x4b, 0x8d,
	0x8f, 0x2c, 0xcd, 0x3d, 0xc3, 0x0d, 0xaa, 0x2d, 0x1a, 0x4f, 0xad, 0xf0, 0x98, 0xdc, 0xf9, 0x70,
	0xe4, 0x36, 0x97, 0xc8, 0x72, 0xdd, 0x47, 0xb8, 0x5d, 0x70, 0xa9, 0x04, 0xd0, 0x79, 0x3c, 0xce,
	0x0a, 0xa9, 0x40, 0x18, 0xcb, 0x3b, 0xd1, 0x5e, 0x55, 0x3f, 0xb3, 0x65, 0xf7, 0x19, 0xde, 0x9a,
	0x01, 0x4d, 0x40, 0x48, 0xaf, 0xde, 0x73, 0xfc, 0x56, 0x78, 0xb8, 0x56, 0xa4, 0x3c, 0xd5, 0xcd,
	0xed, 0x06, 0x0c, 0x0c, 0xe4, 0x2d, 0x55, 0xe3, 0x19, 0x88, 0xa8, 0x62, 0x68, 0x9d, 0x84, 0x5d,
	0xe6, 0x6b, 0xa5, 0x9c, 0x25, 0x20, 0xbd, 0x46, 0xcf, 0xd1, 0x3a, 0xb7, 0xf5, 0x77, 0xba, 0xec,
	0x0e, 0xf0, 0xfe, 0x9c, 0x2e, 0x62, 0x9d, 0x71, 0x09, 0xb1, 0x99, 0x5d, 0x7a, 0x4d, 0xe3, 0xf1,
	0x01, 0xb1, 0x5b, 0x41, 0xaa, 0xad, 0x20, 0x1f, 0x5f, 0xe5, 0xea, 0x24, 0xfc, 0x44, 0xb3, 0x02,
	0xa2, 0xbd, 0x39, 0x5d, 0x9c, 0x1a, 0x96, 0xf1, 0x29, 0xdd, 0x2f, 0xf8, 0x40, 0x80, 0xe4, 0x2c,
	0x97, 0x10, 0x0b, 0xaa, 0x20, 0xce, 0xd2, 0x79, 0xaa, 0xbc, 0x2d, 0xd3, 0xeb, 0xc9, 0x7f, 0xe4,
	0x1d, 0x51, 0x05, 0x6f, 0x34, 0x27, 0xda, 0xaf, 0x1a, 0xfd, 0x2d, 0xf5, 0xdf, 0x5f, 0xaf, 0xba,
	0xe8, 0xe7, 0xaa, 0x8b, 0x7e, 0xaf, 0xba, 0x08, 0x07, 0x29, 0xb3, 0x0d, 0xb9, 0x60, 0x8b, 0xe5,
	0xdd, 0x6f, 0xd1, 0xb7, 0x4b, 0x3a, 0xd4, 0x46, 0x86, 0xe8, 0x73, 0xad, 0x0c, 0x47, 0x4d, 0xe3,
	0xea, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x26, 0x1a, 0x15, 0xb1, 0x03, 0x00, 0x00,
}

func (m *FaultAbort) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultAbort) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ErrorType != nil {
		nn1, err := m.ErrorType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Percentage != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Percentage.Size()))
		n2, err := m.Percentage.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *FaultAbort_HttpStatus) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x10
	i++
	i = encodeVarintFault(dAtA, i, uint64(m.HttpStatus))
	return i, nil
}
func (m *HTTPFault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPFault) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Delay != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Delay.Size()))
		n3, err := m.Delay.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Abort != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Abort.Size()))
		n4, err := m.Abort.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.UpstreamCluster) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFault(dAtA, i, uint64(len(m.UpstreamCluster)))
		i += copy(dAtA[i:], m.UpstreamCluster)
	}
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x22
			i++
			i = encodeVarintFault(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DownstreamNodes) > 0 {
		for _, s := range m.DownstreamNodes {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.MaxActiveFaults != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.MaxActiveFaults.Size()))
		n5, err := m.MaxActiveFaults.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.ResponseRateLimit != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.ResponseRateLimit.Size()))
		n6, err := m.ResponseRateLimit.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFault(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FaultAbort) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ErrorType != nil {
		n += m.ErrorType.Size()
	}
	if m.Percentage != nil {
		l = m.Percentage.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultAbort_HttpStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovFault(uint64(m.HttpStatus))
	return n
}
func (m *HTTPFault) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Delay != nil {
		l = m.Delay.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.Abort != nil {
		l = m.Abort.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.UpstreamCluster)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovFault(uint64(l))
		}
	}
	if len(m.DownstreamNodes) > 0 {
		for _, s := range m.DownstreamNodes {
			l = len(s)
			n += 1 + l + sovFault(uint64(l))
		}
	}
	if m.MaxActiveFaults != nil {
		l = m.MaxActiveFaults.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.ResponseRateLimit != nil {
		l = m.ResponseRateLimit.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFault(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFault(x uint64) (n int) {
	return sovFault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FaultAbort) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
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
			return fmt.Errorf("proto: FaultAbort: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultAbort: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpStatus", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ErrorType = &FaultAbort_HttpStatus{v}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Percentage == nil {
				m.Percentage = &_type.FractionalPercent{}
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
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
func (m *HTTPFault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
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
			return fmt.Errorf("proto: HTTPFault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPFault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Delay == nil {
				m.Delay = &v2.FaultDelay{}
			}
			if err := m.Delay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Abort == nil {
				m.Abort = &FaultAbort{}
			}
			if err := m.Abort.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpstreamCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpstreamCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &route.HeaderMatcher{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownstreamNodes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DownstreamNodes = append(m.DownstreamNodes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxActiveFaults", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MaxActiveFaults == nil {
				m.MaxActiveFaults = &types.UInt32Value{}
			}
			if err := m.MaxActiveFaults.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseRateLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
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
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFault
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ResponseRateLimit == nil {
				m.ResponseRateLimit = &v2.FaultRateLimit{}
			}
			if err := m.ResponseRateLimit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFault
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
func skipFault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFault
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
					return 0, ErrIntOverflowFault
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
					return 0, ErrIntOverflowFault
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
				return 0, ErrInvalidLengthFault
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFault
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFault
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
				next, err := skipFault(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFault
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
	ErrInvalidLengthFault = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFault   = fmt.Errorf("proto: integer overflow")
)

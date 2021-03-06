// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/shadowing/shadowing.proto

package shadowing

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

// Specifies traffic shadowing configuration for the associated route.
// If set, Envoy will send a portion of the route's traffic to the shadowed upstream. This can be a useful way to
// preview a new service's behavior before putting the service in the critical path.
// Note that this plugin is only applicable to routes with upstream destinations (not redirect or direct response routes).
// See here for additional information on Envoy's shadowing capabilities: https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/route/route.proto#envoy-api-msg-route-routeaction-requestmirrorpolicy
type RouteShadowing struct {
	// The upstream to which the shadowed traffic should be sent.
	Upstream *core.ResourceRef `protobuf:"bytes,1,opt,name=upstream,proto3" json:"upstream,omitempty"`
	// This should be a value between 0.0 and 100.0, with up to 6 significant digits.
	Percentage           float32  `protobuf:"fixed32,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteShadowing) Reset()         { *m = RouteShadowing{} }
func (m *RouteShadowing) String() string { return proto.CompactTextString(m) }
func (*RouteShadowing) ProtoMessage()    {}
func (*RouteShadowing) Descriptor() ([]byte, []int) {
	return fileDescriptor_74006d9a50a86d32, []int{0}
}
func (m *RouteShadowing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteShadowing.Unmarshal(m, b)
}
func (m *RouteShadowing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteShadowing.Marshal(b, m, deterministic)
}
func (m *RouteShadowing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteShadowing.Merge(m, src)
}
func (m *RouteShadowing) XXX_Size() int {
	return xxx_messageInfo_RouteShadowing.Size(m)
}
func (m *RouteShadowing) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteShadowing.DiscardUnknown(m)
}

var xxx_messageInfo_RouteShadowing proto.InternalMessageInfo

func (m *RouteShadowing) GetUpstream() *core.ResourceRef {
	if m != nil {
		return m.Upstream
	}
	return nil
}

func (m *RouteShadowing) GetPercentage() float32 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

func init() {
	proto.RegisterType((*RouteShadowing)(nil), "shadowing.options.gloo.solo.io.RouteShadowing")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/shadowing/shadowing.proto", fileDescriptor_74006d9a50a86d32)
}

var fileDescriptor_74006d9a50a86d32 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x95, 0x0a, 0x21, 0x14, 0x24, 0x86, 0x88, 0xa1, 0xed, 0x10, 0x55, 0x4c, 0x5d, 0xb0,
	0x05, 0x88, 0x17, 0xe8, 0x8c, 0x18, 0xc2, 0xc6, 0xe6, 0xb8, 0x37, 0xae, 0x69, 0x9a, 0x63, 0xd9,
	0xd7, 0xb4, 0x8f, 0xc4, 0x23, 0xf0, 0x3c, 0xbc, 0x03, 0x3b, 0xca, 0x1f, 0x65, 0xa9, 0xc4, 0x76,
	0xef, 0x39, 0xc7, 0xc7, 0x9f, 0x6e, 0xfa, 0x6c, 0x2c, 0x6f, 0x62, 0x29, 0x34, 0x76, 0x32, 0xa0,
	0xc6, 0xad, 0x85, 0x34, 0x35, 0x20, 0x9d, 0xc7, 0x1b, 0x69, 0x0e, 0xfd, 0xa6, 0x9c, 0x95, 0xef,
	0x77, 0x12, 0x8e, 0x2d, 0x9a, 0x20, 0xc3, 0x46, 0xad, 0xb1, 0xb7, 0x8d, 0x39, 0x4e, 0xc2, 0x79,
	0x30, 0xb2, 0xfc, 0x28, 0x0c, 0x61, 0xd1, 0x16, 0x88, 0xb6, 0x5b, 0x58, 0xcc, 0xaf, 0x0d, 0x0c,
	0xba, 0xa8, 0x6c, 0xa7, 0xfe, 0xd5, 0x3c, 0x37, 0x80, 0xa9, 0x49, 0x76, 0x5b, 0x19, 0x2b, 0xb9,
	0xf7, 0xca, 0x39, 0xf2, 0xe1, 0x94, 0xbf, 0x8e, 0x5e, 0xb5, 0xed, 0x83, 0x3f, 0xeb, 0xd0, 0xb7,
	0x96, 0x47, 0x50, 0x4f, 0xd5, 0x60, 0x65, 0x74, 0xe0, 0xfe, 0x3f, 0x3a, 0x70, 0xaf, 0xdd, 0x98,
	0xf4, 0xaa, 0x40, 0x64, 0x7a, 0x19, 0x59, 0xb3, 0xc7, 0xf4, 0x22, 0xba, 0xc0, 0x9e, 0xd4, 0x6e,
	0x9a, 0x2c, 0x92, 0xe5, 0xe5, 0xfd, 0x4c, 0x68, 0x78, 0x1a, 0xb9, 0x45, 0x41, 0x01, 0xd1, 0x6b,
	0x2a, 0xa8, 0x2a, 0x7e, 0xa3, 0x59, 0x9e, 0xa6, 0x8e, 0xbc, 0xa6, 0x86, 0x95, 0xa1, 0xe9, 0x64,
	0x91, 0x2c, 0x27, 0xc5, 0x1f, 0x65, 0xf5, 0xf4, 0xf9, 0x7d, 0x96, 0x7c, 0x7c, 0xe5, 0xc9, 0xeb,
	0xea, 0x7f, 0x77, 0x76, 0x5b, 0x73, 0xf2, 0xd6, 0xe5, 0x79, 0x47, 0xff, 0xf0, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0xed, 0x6a, 0xf4, 0xb4, 0x01, 0x00, 0x00,
}

func (this *RouteShadowing) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteShadowing)
	if !ok {
		that2, ok := that.(RouteShadowing)
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
	if !this.Upstream.Equal(that1.Upstream) {
		return false
	}
	if this.Percentage != that1.Percentage {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

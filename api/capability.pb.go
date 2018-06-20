// Code generated by protoc-gen-go. DO NOT EDIT.
// source: capability.proto

package gobgpapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AddPathMode int32

const (
	AddPathMode_MODE_NONE    AddPathMode = 0
	AddPathMode_MODE_RECEIVE AddPathMode = 1
	AddPathMode_MODE_SEND    AddPathMode = 2
	AddPathMode_MODE_BOTH    AddPathMode = 3
)

var AddPathMode_name = map[int32]string{
	0: "MODE_NONE",
	1: "MODE_RECEIVE",
	2: "MODE_SEND",
	3: "MODE_BOTH",
}
var AddPathMode_value = map[string]int32{
	"MODE_NONE":    0,
	"MODE_RECEIVE": 1,
	"MODE_SEND":    2,
	"MODE_BOTH":    3,
}

func (x AddPathMode) String() string {
	return proto.EnumName(AddPathMode_name, int32(x))
}
func (AddPathMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type MultiProtocolCapability struct {
	Family Family `protobuf:"varint,1,opt,name=family,enum=gobgpapi.Family" json:"family,omitempty"`
}

func (m *MultiProtocolCapability) Reset()                    { *m = MultiProtocolCapability{} }
func (m *MultiProtocolCapability) String() string            { return proto.CompactTextString(m) }
func (*MultiProtocolCapability) ProtoMessage()               {}
func (*MultiProtocolCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *MultiProtocolCapability) GetFamily() Family {
	if m != nil {
		return m.Family
	}
	return Family_RESERVED
}

type RouteRefreshCapability struct {
}

func (m *RouteRefreshCapability) Reset()                    { *m = RouteRefreshCapability{} }
func (m *RouteRefreshCapability) String() string            { return proto.CompactTextString(m) }
func (*RouteRefreshCapability) ProtoMessage()               {}
func (*RouteRefreshCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type CarryingLabelInfoCapability struct {
}

func (m *CarryingLabelInfoCapability) Reset()                    { *m = CarryingLabelInfoCapability{} }
func (m *CarryingLabelInfoCapability) String() string            { return proto.CompactTextString(m) }
func (*CarryingLabelInfoCapability) ProtoMessage()               {}
func (*CarryingLabelInfoCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type ExtendedNexthopCapabilityTuple struct {
	NlriFamily Family `protobuf:"varint,1,opt,name=nlri_family,json=nlriFamily,enum=gobgpapi.Family" json:"nlri_family,omitempty"`
	// Nexthop AFI must be either
	// gobgp.IPv4 or
	// gobgp.IPv6.
	NexthopFamily Family `protobuf:"varint,2,opt,name=nexthop_family,json=nexthopFamily,enum=gobgpapi.Family" json:"nexthop_family,omitempty"`
}

func (m *ExtendedNexthopCapabilityTuple) Reset()                    { *m = ExtendedNexthopCapabilityTuple{} }
func (m *ExtendedNexthopCapabilityTuple) String() string            { return proto.CompactTextString(m) }
func (*ExtendedNexthopCapabilityTuple) ProtoMessage()               {}
func (*ExtendedNexthopCapabilityTuple) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ExtendedNexthopCapabilityTuple) GetNlriFamily() Family {
	if m != nil {
		return m.NlriFamily
	}
	return Family_RESERVED
}

func (m *ExtendedNexthopCapabilityTuple) GetNexthopFamily() Family {
	if m != nil {
		return m.NexthopFamily
	}
	return Family_RESERVED
}

type ExtendedNexthopCapability struct {
	Tuples []*ExtendedNexthopCapabilityTuple `protobuf:"bytes,1,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *ExtendedNexthopCapability) Reset()                    { *m = ExtendedNexthopCapability{} }
func (m *ExtendedNexthopCapability) String() string            { return proto.CompactTextString(m) }
func (*ExtendedNexthopCapability) ProtoMessage()               {}
func (*ExtendedNexthopCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *ExtendedNexthopCapability) GetTuples() []*ExtendedNexthopCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type GracefulRestartCapabilityTuple struct {
	Family Family `protobuf:"varint,1,opt,name=family,enum=gobgpapi.Family" json:"family,omitempty"`
	Flags  uint32 `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
}

func (m *GracefulRestartCapabilityTuple) Reset()                    { *m = GracefulRestartCapabilityTuple{} }
func (m *GracefulRestartCapabilityTuple) String() string            { return proto.CompactTextString(m) }
func (*GracefulRestartCapabilityTuple) ProtoMessage()               {}
func (*GracefulRestartCapabilityTuple) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *GracefulRestartCapabilityTuple) GetFamily() Family {
	if m != nil {
		return m.Family
	}
	return Family_RESERVED
}

func (m *GracefulRestartCapabilityTuple) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type GracefulRestartCapability struct {
	Flags  uint32                            `protobuf:"varint,1,opt,name=flags" json:"flags,omitempty"`
	Time   uint32                            `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Tuples []*GracefulRestartCapabilityTuple `protobuf:"bytes,3,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *GracefulRestartCapability) Reset()                    { *m = GracefulRestartCapability{} }
func (m *GracefulRestartCapability) String() string            { return proto.CompactTextString(m) }
func (*GracefulRestartCapability) ProtoMessage()               {}
func (*GracefulRestartCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *GracefulRestartCapability) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *GracefulRestartCapability) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *GracefulRestartCapability) GetTuples() []*GracefulRestartCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type FourOctetASNumberCapability struct {
	As uint32 `protobuf:"varint,1,opt,name=as" json:"as,omitempty"`
}

func (m *FourOctetASNumberCapability) Reset()                    { *m = FourOctetASNumberCapability{} }
func (m *FourOctetASNumberCapability) String() string            { return proto.CompactTextString(m) }
func (*FourOctetASNumberCapability) ProtoMessage()               {}
func (*FourOctetASNumberCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *FourOctetASNumberCapability) GetAs() uint32 {
	if m != nil {
		return m.As
	}
	return 0
}

type AddPathCapabilityTuple struct {
	Family Family      `protobuf:"varint,1,opt,name=family,enum=gobgpapi.Family" json:"family,omitempty"`
	Mode   AddPathMode `protobuf:"varint,2,opt,name=mode,enum=gobgpapi.AddPathMode" json:"mode,omitempty"`
}

func (m *AddPathCapabilityTuple) Reset()                    { *m = AddPathCapabilityTuple{} }
func (m *AddPathCapabilityTuple) String() string            { return proto.CompactTextString(m) }
func (*AddPathCapabilityTuple) ProtoMessage()               {}
func (*AddPathCapabilityTuple) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *AddPathCapabilityTuple) GetFamily() Family {
	if m != nil {
		return m.Family
	}
	return Family_RESERVED
}

func (m *AddPathCapabilityTuple) GetMode() AddPathMode {
	if m != nil {
		return m.Mode
	}
	return AddPathMode_MODE_NONE
}

type AddPathCapability struct {
	Tuples []*AddPathCapabilityTuple `protobuf:"bytes,1,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *AddPathCapability) Reset()                    { *m = AddPathCapability{} }
func (m *AddPathCapability) String() string            { return proto.CompactTextString(m) }
func (*AddPathCapability) ProtoMessage()               {}
func (*AddPathCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *AddPathCapability) GetTuples() []*AddPathCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type EnhancedRouteRefreshCapability struct {
}

func (m *EnhancedRouteRefreshCapability) Reset()                    { *m = EnhancedRouteRefreshCapability{} }
func (m *EnhancedRouteRefreshCapability) String() string            { return proto.CompactTextString(m) }
func (*EnhancedRouteRefreshCapability) ProtoMessage()               {}
func (*EnhancedRouteRefreshCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

type LongLivedGracefulRestartCapabilityTuple struct {
	Family Family `protobuf:"varint,1,opt,name=family,enum=gobgpapi.Family" json:"family,omitempty"`
	Flags  uint32 `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
	Time   uint32 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
}

func (m *LongLivedGracefulRestartCapabilityTuple) Reset() {
	*m = LongLivedGracefulRestartCapabilityTuple{}
}
func (m *LongLivedGracefulRestartCapabilityTuple) String() string { return proto.CompactTextString(m) }
func (*LongLivedGracefulRestartCapabilityTuple) ProtoMessage()    {}
func (*LongLivedGracefulRestartCapabilityTuple) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{11}
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetFamily() Family {
	if m != nil {
		return m.Family
	}
	return Family_RESERVED
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetFlags() uint32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *LongLivedGracefulRestartCapabilityTuple) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

type LongLivedGracefulRestartCapability struct {
	Tuples []*LongLivedGracefulRestartCapabilityTuple `protobuf:"bytes,1,rep,name=tuples" json:"tuples,omitempty"`
}

func (m *LongLivedGracefulRestartCapability) Reset()         { *m = LongLivedGracefulRestartCapability{} }
func (m *LongLivedGracefulRestartCapability) String() string { return proto.CompactTextString(m) }
func (*LongLivedGracefulRestartCapability) ProtoMessage()    {}
func (*LongLivedGracefulRestartCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor2, []int{12}
}

func (m *LongLivedGracefulRestartCapability) GetTuples() []*LongLivedGracefulRestartCapabilityTuple {
	if m != nil {
		return m.Tuples
	}
	return nil
}

type RouteRefreshCiscoCapability struct {
}

func (m *RouteRefreshCiscoCapability) Reset()                    { *m = RouteRefreshCiscoCapability{} }
func (m *RouteRefreshCiscoCapability) String() string            { return proto.CompactTextString(m) }
func (*RouteRefreshCiscoCapability) ProtoMessage()               {}
func (*RouteRefreshCiscoCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

type UnknownCapability struct {
	Code  uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *UnknownCapability) Reset()                    { *m = UnknownCapability{} }
func (m *UnknownCapability) String() string            { return proto.CompactTextString(m) }
func (*UnknownCapability) ProtoMessage()               {}
func (*UnknownCapability) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *UnknownCapability) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UnknownCapability) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*MultiProtocolCapability)(nil), "gobgpapi.MultiProtocolCapability")
	proto.RegisterType((*RouteRefreshCapability)(nil), "gobgpapi.RouteRefreshCapability")
	proto.RegisterType((*CarryingLabelInfoCapability)(nil), "gobgpapi.CarryingLabelInfoCapability")
	proto.RegisterType((*ExtendedNexthopCapabilityTuple)(nil), "gobgpapi.ExtendedNexthopCapabilityTuple")
	proto.RegisterType((*ExtendedNexthopCapability)(nil), "gobgpapi.ExtendedNexthopCapability")
	proto.RegisterType((*GracefulRestartCapabilityTuple)(nil), "gobgpapi.GracefulRestartCapabilityTuple")
	proto.RegisterType((*GracefulRestartCapability)(nil), "gobgpapi.GracefulRestartCapability")
	proto.RegisterType((*FourOctetASNumberCapability)(nil), "gobgpapi.FourOctetASNumberCapability")
	proto.RegisterType((*AddPathCapabilityTuple)(nil), "gobgpapi.AddPathCapabilityTuple")
	proto.RegisterType((*AddPathCapability)(nil), "gobgpapi.AddPathCapability")
	proto.RegisterType((*EnhancedRouteRefreshCapability)(nil), "gobgpapi.EnhancedRouteRefreshCapability")
	proto.RegisterType((*LongLivedGracefulRestartCapabilityTuple)(nil), "gobgpapi.LongLivedGracefulRestartCapabilityTuple")
	proto.RegisterType((*LongLivedGracefulRestartCapability)(nil), "gobgpapi.LongLivedGracefulRestartCapability")
	proto.RegisterType((*RouteRefreshCiscoCapability)(nil), "gobgpapi.RouteRefreshCiscoCapability")
	proto.RegisterType((*UnknownCapability)(nil), "gobgpapi.UnknownCapability")
	proto.RegisterEnum("gobgpapi.AddPathMode", AddPathMode_name, AddPathMode_value)
}

func init() { proto.RegisterFile("capability.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x39, 0xc9, 0x2f, 0x82, 0x49, 0x13, 0xb9, 0x2b, 0x28, 0x29, 0x51, 0xa3, 0x68, 0x2f,
	0x04, 0x24, 0x22, 0xb5, 0x1c, 0xe0, 0x82, 0x44, 0x49, 0x5d, 0x88, 0x94, 0x8f, 0xca, 0x2d, 0xdc,
	0x50, 0xd9, 0xd8, 0x6b, 0x67, 0xc5, 0x7a, 0xd7, 0xb2, 0xd7, 0xa5, 0x39, 0x70, 0xe6, 0xc2, 0x1f,
	0x8d, 0xfc, 0x11, 0xdb, 0xa4, 0x72, 0x5b, 0x21, 0x71, 0x9b, 0xf1, 0xce, 0xbc, 0x79, 0x6f, 0xde,
	0xae, 0x41, 0xb7, 0x88, 0x4f, 0x96, 0x8c, 0x33, 0xb5, 0x1e, 0xf9, 0x81, 0x54, 0x12, 0x3d, 0x70,
	0xe5, 0xd2, 0xf5, 0x89, 0xcf, 0x9e, 0xb6, 0x92, 0x28, 0xfd, 0x8c, 0xc7, 0xf0, 0x64, 0x16, 0x71,
	0xc5, 0xce, 0xe2, 0xcc, 0x92, 0x7c, 0x9c, 0xf7, 0xa1, 0x21, 0x34, 0x1d, 0xe2, 0x31, 0xbe, 0xee,
	0x6a, 0x03, 0x6d, 0xd8, 0x39, 0xd2, 0x47, 0x1b, 0x88, 0xd1, 0x69, 0xf2, 0xdd, 0xcc, 0xce, 0x71,
	0x17, 0xf6, 0x4c, 0x19, 0x29, 0x6a, 0x52, 0x27, 0xa0, 0xe1, 0xaa, 0xc0, 0xc0, 0x07, 0xd0, 0x1b,
	0x93, 0x20, 0x58, 0x33, 0xe1, 0x4e, 0xc9, 0x92, 0xf2, 0x89, 0x70, 0x64, 0xe9, 0xf8, 0x97, 0x06,
	0x7d, 0xe3, 0x5a, 0x51, 0x61, 0x53, 0x7b, 0x4e, 0xaf, 0xd5, 0x4a, 0xfa, 0xc5, 0xe9, 0x45, 0xe4,
	0x73, 0x8a, 0x0e, 0xa1, 0x25, 0x78, 0xc0, 0x2e, 0xef, 0xa0, 0x02, 0x71, 0x51, 0x1a, 0xa3, 0xd7,
	0xd0, 0x11, 0x29, 0xd8, 0xa6, 0xab, 0x56, 0xd1, 0xd5, 0xce, 0xea, 0xd2, 0x14, 0x7f, 0x81, 0xfd,
	0x4a, 0x36, 0xe8, 0x1d, 0x34, 0x55, 0xcc, 0x28, 0xec, 0x6a, 0x83, 0xfa, 0xb0, 0x75, 0x34, 0x2c,
	0xd0, 0x6e, 0x97, 0x60, 0x66, 0x7d, 0xf8, 0x2b, 0xf4, 0x3f, 0x04, 0xc4, 0xa2, 0x4e, 0xc4, 0x4d,
	0x1a, 0x2a, 0x12, 0xa8, 0x6d, 0xb1, 0xf7, 0x5e, 0x39, 0x7a, 0x04, 0xff, 0x3b, 0x9c, 0xb8, 0x61,
	0x22, 0xad, 0x6d, 0xa6, 0x09, 0xfe, 0xa9, 0xc1, 0x7e, 0xe5, 0x88, 0xa2, 0x47, 0x2b, 0xf5, 0x20,
	0x04, 0x0d, 0xc5, 0x3c, 0x9a, 0x01, 0x25, 0x71, 0x49, 0x6b, 0x7d, 0x5b, 0xeb, 0xed, 0x0a, 0x72,
	0xad, 0x2f, 0xa1, 0x77, 0x2a, 0xa3, 0x60, 0x61, 0x29, 0xaa, 0x8e, 0xcf, 0xe7, 0x91, 0xb7, 0xa4,
	0x41, 0x89, 0x4a, 0x07, 0x6a, 0x64, 0xc3, 0xa3, 0x46, 0x42, 0xec, 0xc1, 0xde, 0xb1, 0x6d, 0x9f,
	0x11, 0xb5, 0xfa, 0xfb, 0x95, 0x3c, 0x87, 0x86, 0x27, 0x6d, 0x9a, 0x99, 0xfd, 0xb8, 0xa8, 0xcb,
	0x90, 0x67, 0xd2, 0xa6, 0x66, 0x52, 0x82, 0x67, 0xb0, 0x7b, 0x63, 0x1c, 0x7a, 0xb3, 0x65, 0xf0,
	0xe0, 0x06, 0x42, 0x95, 0xd8, 0x01, 0xf4, 0x0d, 0xb1, 0x22, 0xc2, 0xa2, 0x76, 0xc5, 0x3b, 0xf8,
	0x01, 0xcf, 0xa6, 0x52, 0xb8, 0x53, 0x76, 0x45, 0xed, 0x7f, 0x7b, 0x07, 0x72, 0x3f, 0xeb, 0x85,
	0x9f, 0x58, 0x02, 0xbe, 0x7b, 0x3c, 0x9a, 0x6c, 0x2d, 0xe0, 0xb0, 0x98, 0x7c, 0x4f, 0xf2, 0xf9,
	0x46, 0x0e, 0xa0, 0xf7, 0xc7, 0x26, 0x58, 0x68, 0x95, 0xdf, 0xfd, 0x5b, 0xd8, 0xfd, 0x24, 0xbe,
	0x09, 0xf9, 0x5d, 0x94, 0xc6, 0x23, 0x68, 0x58, 0xb1, 0x7f, 0xe9, 0xad, 0x48, 0xe2, 0x58, 0xe2,
	0x15, 0xe1, 0x51, 0x6a, 0xea, 0x8e, 0x99, 0x26, 0x2f, 0xa6, 0xd0, 0x2a, 0x79, 0x8a, 0xda, 0xf0,
	0x70, 0xb6, 0x38, 0x31, 0x2e, 0xe7, 0x8b, 0xb9, 0xa1, 0xff, 0x87, 0x74, 0xd8, 0x49, 0x52, 0xd3,
	0x18, 0x1b, 0x93, 0xcf, 0x86, 0xae, 0xe5, 0x05, 0xe7, 0xc6, 0xfc, 0x44, 0xaf, 0xe5, 0xe9, 0xfb,
	0xc5, 0xc5, 0x47, 0xbd, 0xbe, 0x6c, 0x26, 0x7f, 0xc2, 0x57, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xd6, 0xbd, 0xd3, 0xc3, 0x34, 0x05, 0x00, 0x00,
}
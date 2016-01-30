// Code generated by protoc-gen-go.
// source: proto/google_security_e2ekeys_v1/e2ekeys.proto
// DO NOT EDIT!

/*
Package google_security_e2ekeys_v1 is a generated protocol buffer package.

It is generated from these files:
	proto/google_security_e2ekeys_v1/e2ekeys.proto

It has these top-level messages:
	HkpLookupRequest
	HttpResponse
*/
package google_security_e2ekeys_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_security_e2ekeys_v2 "github.com/google/e2e-key-server/proto/google_security_e2ekeys_v2"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// HkpLookupRequest contains query parameters for retrieving PGP keys.
type HkpLookupRequest struct {
	// Op specifies the operation to be performed on the keyserver.
	// - "get" returns the pgp key specified in the search parameter.
	// - "index" returns 501 (not implemented).
	// - "vindex" returns 501 (not implemented).
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// Search specifies the email address or key id being queried.
	Search string `protobuf:"bytes,2,opt,name=search" json:"search,omitempty"`
	// Options specifies what output format to use.
	// - "mr" machine readable will set the content type to "application/pgp-keys"
	// - other options will be ignored.
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	// Exact specifies an exact match on search. Always on. If specified in the
	// URL, its value will be ignored.
	Exact string `protobuf:"bytes,4,opt,name=exact" json:"exact,omitempty"`
	// fingerprint is ignored.
	Fingerprint string `protobuf:"bytes,5,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *HkpLookupRequest) Reset()                    { *m = HkpLookupRequest{} }
func (m *HkpLookupRequest) String() string            { return proto.CompactTextString(m) }
func (*HkpLookupRequest) ProtoMessage()               {}
func (*HkpLookupRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// HttpBody represents an http body.
type HttpResponse struct {
	// Header content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type" json:"content_type,omitempty"`
	// The http body itself.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HttpResponse) Reset()                    { *m = HttpResponse{} }
func (m *HttpResponse) String() string            { return proto.CompactTextString(m) }
func (*HttpResponse) ProtoMessage()               {}
func (*HttpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*HkpLookupRequest)(nil), "google.security.e2ekeys.v1.HkpLookupRequest")
	proto.RegisterType((*HttpResponse)(nil), "google.security.e2ekeys.v1.HttpResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for E2EKeyProxy service

type E2EKeyProxyClient interface {
	// GetEntry returns a user's current profile.
	GetEntry(ctx context.Context, in *google_security_e2ekeys_v2.GetEntryRequest, opts ...grpc.CallOption) (*google_security_e2ekeys_v2.Profile, error)
	HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error)
}

type e2EKeyProxyClient struct {
	cc *grpc.ClientConn
}

func NewE2EKeyProxyClient(cc *grpc.ClientConn) E2EKeyProxyClient {
	return &e2EKeyProxyClient{cc}
}

func (c *e2EKeyProxyClient) GetEntry(ctx context.Context, in *google_security_e2ekeys_v2.GetEntryRequest, opts ...grpc.CallOption) (*google_security_e2ekeys_v2.Profile, error) {
	out := new(google_security_e2ekeys_v2.Profile)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v1.E2EKeyProxy/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyProxyClient) HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	out := new(HttpResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v1.E2EKeyProxy/HkpLookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for E2EKeyProxy service

type E2EKeyProxyServer interface {
	// GetEntry returns a user's current profile.
	GetEntry(context.Context, *google_security_e2ekeys_v2.GetEntryRequest) (*google_security_e2ekeys_v2.Profile, error)
	HkpLookup(context.Context, *HkpLookupRequest) (*HttpResponse, error)
}

func RegisterE2EKeyProxyServer(s *grpc.Server, srv E2EKeyProxyServer) {
	s.RegisterService(&_E2EKeyProxy_serviceDesc, srv)
}

func _E2EKeyProxy_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_security_e2ekeys_v2.GetEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyProxyServer).GetEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyProxy_HkpLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(HkpLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyProxyServer).HkpLookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _E2EKeyProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.security.e2ekeys.v1.E2EKeyProxy",
	HandlerType: (*E2EKeyProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _E2EKeyProxy_GetEntry_Handler,
		},
		{
			MethodName: "HkpLookup",
			Handler:    _E2EKeyProxy_HkpLookup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xdf, 0x4e, 0xf2, 0x40,
	0x10, 0xc5, 0x03, 0x1f, 0xf0, 0xc9, 0x50, 0xff, 0x64, 0xf5, 0xa2, 0xe9, 0x95, 0xc1, 0x1b, 0x12,
	0x65, 0x09, 0xf5, 0x19, 0x88, 0x24, 0x9a, 0x68, 0xbc, 0x36, 0x69, 0x60, 0x1d, 0x4a, 0x03, 0xee,
	0xac, 0xbb, 0x53, 0xc2, 0xbe, 0xa5, 0x8f, 0x64, 0x29, 0xc5, 0xa8, 0x09, 0x78, 0xb7, 0x3b, 0xf3,
	0x3b, 0xe7, 0x64, 0x0e, 0x48, 0x63, 0x89, 0x69, 0x90, 0x12, 0xa5, 0x4b, 0x4c, 0x1c, 0xaa, 0xdc,
	0x66, 0xec, 0x13, 0x8c, 0x71, 0x81, 0xde, 0x25, 0xab, 0xe1, 0xa0, 0x7a, 0x6e, 0x41, 0x11, 0x6d,
	0x49, 0xb9, 0x23, 0xe5, 0x6e, 0xbd, 0x1a, 0x46, 0x8f, 0x69, 0xc6, 0xf3, 0x7c, 0x2a, 0x15, 0xbd,
	0x55, 0x86, 0x1b, 0x71, 0xbf, 0x58, 0xf7, 0x1d, 0xda, 0x15, 0xda, 0xc1, 0x1f, 0x61, 0xf1, 0xcf,
	0xb0, 0xae, 0x82, 0xb3, 0xf1, 0xc2, 0x3c, 0x10, 0x2d, 0x72, 0xf3, 0x8c, 0xef, 0x39, 0x3a, 0x16,
	0x00, 0x75, 0x32, 0x61, 0xed, 0xb2, 0xd6, 0x6b, 0x8b, 0x13, 0x68, 0x39, 0x9c, 0x58, 0x35, 0x0f,
	0xeb, 0xe5, 0xff, 0x14, 0xfe, 0x93, 0xe1, 0x8c, 0xb4, 0x0b, 0xff, 0x95, 0x83, 0x63, 0x68, 0xe2,
	0x7a, 0xa2, 0x38, 0x6c, 0x94, 0xdf, 0x73, 0xe8, 0xcc, 0x32, 0x9d, 0xa2, 0x35, 0x36, 0xd3, 0x1c,
	0x36, 0x37, 0xc3, 0x6e, 0x0c, 0xc1, 0x98, 0xb9, 0xf0, 0x77, 0xa6, 0x10, 0xa2, 0xb8, 0x80, 0x40,
	0x91, 0x66, 0xd4, 0x9c, 0xb0, 0x37, 0x58, 0x45, 0x05, 0xd0, 0x98, 0xd2, 0xab, 0x2f, 0x83, 0x82,
	0xf8, 0xa3, 0x06, 0x9d, 0x51, 0x3c, 0xba, 0x47, 0xff, 0x64, 0x69, 0xed, 0xc5, 0x0b, 0x1c, 0xdd,
	0x21, 0x8f, 0x34, 0x5b, 0x2f, 0xae, 0xe5, 0xde, 0x8a, 0x62, 0xb9, 0xa3, 0xaa, 0x6b, 0xa2, 0xab,
	0x43, 0x70, 0xe1, 0x3d, 0xcb, 0x96, 0x28, 0x14, 0xb4, 0xbf, 0x6a, 0x10, 0x37, 0xfb, 0x15, 0x43,
	0xf9, 0xbb, 0xad, 0xa8, 0x77, 0x90, 0xfe, 0x76, 0xf6, 0xb4, 0x55, 0x56, 0x7e, 0xfb, 0x19, 0x00,
	0x00, 0xff, 0xff, 0xf9, 0x45, 0xf3, 0x27, 0x11, 0x02, 0x00, 0x00,
}
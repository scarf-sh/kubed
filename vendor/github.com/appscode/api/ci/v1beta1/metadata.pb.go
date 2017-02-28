// Code generated by protoc-gen-go.
// source: metadata.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerInfoResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Provider  string                  `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	ServerUrl string                  `protobuf:"bytes,3,opt,name=server_url,json=serverUrl" json:"server_url,omitempty"`
	CaCert    string                  `protobuf:"bytes,4,opt,name=ca_cert,json=caCert" json:"ca_cert,omitempty"`
}

func (m *ServerInfoResponse) Reset()                    { *m = ServerInfoResponse{} }
func (m *ServerInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerInfoResponse) ProtoMessage()               {}
func (*ServerInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ServerInfoResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ServerInfoResponse) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *ServerInfoResponse) GetServerUrl() string {
	if m != nil {
		return m.ServerUrl
	}
	return ""
}

func (m *ServerInfoResponse) GetCaCert() string {
	if m != nil {
		return m.CaCert
	}
	return ""
}

func init() {
	proto.RegisterType((*ServerInfoResponse)(nil), "appscode.ci.v1beta1.ServerInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metadata service

type MetadataClient interface {
	ServerInfo(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) ServerInfo(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*ServerInfoResponse, error) {
	out := new(ServerInfoResponse)
	err := grpc.Invoke(ctx, "/appscode.ci.v1beta1.Metadata/ServerInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	ServerInfo(context.Context, *appscode_dtypes.VoidRequest) (*ServerInfoResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_ServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).ServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.ci.v1beta1.Metadata/ServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).ServerInfo(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.ci.v1beta1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerInfo",
			Handler:    _Metadata_ServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0x53, 0xbe, 0x2f, 0xfc, 0xa9, 0xd1, 0x45, 0x5d, 0x30, 0x19, 0x51, 0x09, 0x1b, 0x59,
	0xb5, 0x82, 0x0b, 0x37, 0xae, 0x70, 0xe5, 0xc2, 0x84, 0x0c, 0xd1, 0x85, 0x1b, 0x52, 0x3a, 0x17,
	0xd2, 0x04, 0xe6, 0xd6, 0xb6, 0x90, 0xb8, 0xe5, 0x0d, 0x8c, 0x2b, 0x37, 0xbe, 0x85, 0x4f, 0xe2,
	0x2b, 0xf8, 0x20, 0xc6, 0xce, 0x00, 0x12, 0x74, 0x33, 0xc9, 0xf4, 0x77, 0x4f, 0xcf, 0xb9, 0xa7,
	0xf4, 0x60, 0x06, 0x5e, 0xa6, 0xd2, 0x4b, 0x6e, 0x2c, 0x7a, 0x64, 0x87, 0xd2, 0x18, 0xa7, 0x30,
	0x05, 0xae, 0x34, 0x5f, 0x74, 0x46, 0xe0, 0x65, 0x27, 0x6e, 0x4c, 0x10, 0x27, 0x53, 0x10, 0xd2,
	0x68, 0x21, 0xb3, 0x0c, 0xbd, 0xf4, 0x1a, 0x33, 0x97, 0x4b, 0xe2, 0x93, 0x95, 0xe4, 0x0f, 0x7e,
	0xba, 0xc5, 0x53, 0xff, 0x64, 0xc0, 0x89, 0xf0, 0xcd, 0x07, 0x5a, 0xaf, 0x84, 0xb2, 0x01, 0xd8,
	0x05, 0xd8, 0x9b, 0x6c, 0x8c, 0x09, 0x38, 0x83, 0x99, 0x03, 0x26, 0x68, 0xd9, 0x79, 0xe9, 0xe7,
	0x2e, 0x22, 0x4d, 0xd2, 0xde, 0xeb, 0xd6, 0xf9, 0x3a, 0x5b, 0x7e, 0x09, 0x1f, 0x04, 0x9c, 0x14,
	0x63, 0x2c, 0xa6, 0x55, 0x63, 0x71, 0xa1, 0x53, 0xb0, 0x51, 0xa9, 0x49, 0xda, 0xb5, 0x64, 0xfd,
	0xcf, 0x8e, 0x29, 0x75, 0xc1, 0x62, 0x38, 0xb7, 0xd3, 0xe8, 0x5f, 0xa0, 0xb5, 0xfc, 0xe4, 0xce,
	0x4e, 0x59, 0x9d, 0x56, 0x94, 0x1c, 0x2a, 0xb0, 0x3e, 0xfa, 0x1f, 0x58, 0x59, 0xc9, 0x6b, 0xb0,
	0xbe, 0xfb, 0x46, 0x68, 0xf5, 0xb6, 0xa8, 0x88, 0x3d, 0x13, 0x4a, 0x37, 0x41, 0x59, 0x63, 0x27,
	0xd0, 0x3d, 0xea, 0x34, 0x81, 0xc7, 0x39, 0x38, 0x1f, 0x9f, 0xf1, 0x5f, 0xaa, 0xe4, 0xbb, 0x7b,
	0xb6, 0xae, 0x96, 0xef, 0x51, 0xa9, 0x4a, 0x96, 0x1f, 0x9f, 0x2f, 0xa5, 0x73, 0xc6, 0xc5, 0x70,
	0xab, 0x2f, 0xa5, 0x45, 0xa1, 0x15, 0xab, 0xa7, 0x12, 0xc5, 0x26, 0x3a, 0x1b, 0x63, 0xef, 0x92,
	0x1e, 0x29, 0x9c, 0x6d, 0xbc, 0xa4, 0xd1, 0x3f, 0xfc, 0x7a, 0xfb, 0xab, 0xf0, 0xfd, 0xef, 0xaa,
	0xfb, 0xe4, 0xa1, 0x52, 0x90, 0x51, 0x39, 0x94, 0x7f, 0xf1, 0x15, 0x00, 0x00, 0xff, 0xff, 0x70,
	0x23, 0x81, 0xb5, 0x02, 0x02, 0x00, 0x00,
}
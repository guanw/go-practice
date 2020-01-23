// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

/*
Package stream is a generated protocol buffer package.

It is generated from these files:
	stream.proto

It has these top-level messages:
	Req
	Resp
*/
package stream

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Req struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Resp struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Host    string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Resp) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Resp) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "stream.Req")
	proto.RegisterType((*Resp)(nil), "stream.Resp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DnsService service

type DnsServiceClient interface {
	Echo(ctx context.Context, in *Req, opts ...grpc.CallOption) (DnsService_EchoClient, error)
}

type dnsServiceClient struct {
	cc *grpc.ClientConn
}

func NewDnsServiceClient(cc *grpc.ClientConn) DnsServiceClient {
	return &dnsServiceClient{cc}
}

func (c *dnsServiceClient) Echo(ctx context.Context, in *Req, opts ...grpc.CallOption) (DnsService_EchoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DnsService_serviceDesc.Streams[0], c.cc, "/stream.DnsService/Echo", opts...)
	if err != nil {
		return nil, err
	}
	x := &dnsServiceEchoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DnsService_EchoClient interface {
	Recv() (*Resp, error)
	grpc.ClientStream
}

type dnsServiceEchoClient struct {
	grpc.ClientStream
}

func (x *dnsServiceEchoClient) Recv() (*Resp, error) {
	m := new(Resp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DnsService service

type DnsServiceServer interface {
	Echo(*Req, DnsService_EchoServer) error
}

func RegisterDnsServiceServer(s *grpc.Server, srv DnsServiceServer) {
	s.RegisterService(&_DnsService_serviceDesc, srv)
}

func _DnsService_Echo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Req)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DnsServiceServer).Echo(m, &dnsServiceEchoServer{stream})
}

type DnsService_EchoServer interface {
	Send(*Resp) error
	grpc.ServerStream
}

type dnsServiceEchoServer struct {
	grpc.ServerStream
}

func (x *dnsServiceEchoServer) Send(m *Resp) error {
	return x.ServerStream.SendMsg(m)
}

var _DnsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stream.DnsService",
	HandlerType: (*DnsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Echo",
			Handler:       _DnsService_Echo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x44, 0xb9, 0x98,
	0x83, 0x52, 0x0b, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83,
	0x98, 0x32, 0x53, 0x94, 0x4c, 0xb8, 0x58, 0x82, 0x52, 0x8b, 0x0b, 0x84, 0x24, 0xb8, 0xd8, 0x8b,
	0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xc1, 0x92, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x10, 0x17, 0x4b,
	0x46, 0x7e, 0x71, 0x89, 0x04, 0x13, 0x58, 0x18, 0xcc, 0x36, 0x32, 0xe6, 0xe2, 0x72, 0xc9, 0x2b,
	0x0e, 0x86, 0xaa, 0x50, 0xe5, 0x62, 0x71, 0x4d, 0xce, 0xc8, 0x17, 0xe2, 0xd6, 0x83, 0xda, 0x1c,
	0x94, 0x5a, 0x28, 0xc5, 0x83, 0xe0, 0x14, 0x17, 0x28, 0x31, 0x18, 0x30, 0x26, 0xb1, 0x81, 0x1d,
	0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x30, 0xad, 0x54, 0xb0, 0xa0, 0x00, 0x00, 0x00,
}
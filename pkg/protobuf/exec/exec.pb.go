// Code generated by protoc-gen-go.
// source: pkg/protobuf/exec/exec.proto
// DO NOT EDIT!

/*
Package exec is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/exec/exec.proto

It has these top-level messages:
	CommandRequest
	CommandResponse
*/
package exec

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

type CommandRequest struct {
	Name    string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Command []string `protobuf:"bytes,2,rep,name=command" json:"command,omitempty"`
}

func (m *CommandRequest) Reset()                    { *m = CommandRequest{} }
func (m *CommandRequest) String() string            { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()               {}
func (*CommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommandRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommandRequest) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

type CommandResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *CommandResponse) Reset()                    { *m = CommandResponse{} }
func (m *CommandResponse) String() string            { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()               {}
func (*CommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommandResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*CommandRequest)(nil), "exec.CommandRequest")
	proto.RegisterType((*CommandResponse)(nil), "exec.CommandResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Exec service

type ExecClient interface {
	Command(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (Exec_CommandClient, error)
}

type execClient struct {
	cc *grpc.ClientConn
}

func NewExecClient(cc *grpc.ClientConn) ExecClient {
	return &execClient{cc}
}

func (c *execClient) Command(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (Exec_CommandClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Exec_serviceDesc.Streams[0], c.cc, "/exec.Exec/Command", opts...)
	if err != nil {
		return nil, err
	}
	x := &execCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Exec_CommandClient interface {
	Recv() (*CommandResponse, error)
	grpc.ClientStream
}

type execCommandClient struct {
	grpc.ClientStream
}

func (x *execCommandClient) Recv() (*CommandResponse, error) {
	m := new(CommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Exec service

type ExecServer interface {
	Command(*CommandRequest, Exec_CommandServer) error
}

func RegisterExecServer(s *grpc.Server, srv ExecServer) {
	s.RegisterService(&_Exec_serviceDesc, srv)
}

func _Exec_Command_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommandRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExecServer).Command(m, &execCommandServer{stream})
}

type Exec_CommandServer interface {
	Send(*CommandResponse) error
	grpc.ServerStream
}

type execCommandServer struct {
	grpc.ServerStream
}

func (x *execCommandServer) Send(m *CommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Exec_serviceDesc = grpc.ServiceDesc{
	ServiceName: "exec.Exec",
	HandlerType: (*ExecServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Command",
			Handler:       _Exec_Command_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/protobuf/exec/exec.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/exec/exec.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0xc8, 0x4e, 0xd7,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0xad, 0x48, 0x4d, 0x06, 0x13, 0x7a,
	0x60, 0x21, 0x21, 0x16, 0x10, 0x5b, 0xc9, 0x8e, 0x8b, 0xcf, 0x39, 0x3f, 0x37, 0x37, 0x31, 0x2f,
	0x25, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0x86, 0xa8, 0x92,
	0x60, 0x52, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0x54, 0xb9, 0xf8, 0xe1, 0xfa, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x41, 0x06, 0x94, 0xa4, 0x56, 0x94, 0xc0, 0x0c, 0x00, 0xb1, 0x8d, 0x1c,
	0xb8, 0x58, 0x5c, 0x2b, 0x52, 0x93, 0x85, 0x2c, 0xb8, 0xd8, 0xa1, 0xca, 0x85, 0x44, 0xf4, 0xc0,
	0x8e, 0x41, 0xb5, 0x5d, 0x4a, 0x14, 0x4d, 0x14, 0x62, 0xa6, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0xd5,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x3c, 0x06, 0x4d, 0xd5, 0x00, 0x00, 0x00,
}

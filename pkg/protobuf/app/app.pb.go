// Code generated by protoc-gen-go.
// source: pkg/protobuf/app/app.proto
// DO NOT EDIT!

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/app/app.proto

It has these top-level messages:
	CreateRequest
	Empty
*/
package app

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

type CreateRequest struct {
	Name        string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Team        string                   `protobuf:"bytes,2,opt,name=team" json:"team,omitempty"`
	ProcessType string                   `protobuf:"bytes,3,opt,name=process_type,json=processType" json:"process_type,omitempty"`
	Limits      *CreateRequest_Limits    `protobuf:"bytes,4,opt,name=limits" json:"limits,omitempty"`
	AutoScale   *CreateRequest_AutoScale `protobuf:"bytes,5,opt,name=auto_scale,json=autoScale" json:"auto_scale,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *CreateRequest) GetProcessType() string {
	if m != nil {
		return m.ProcessType
	}
	return ""
}

func (m *CreateRequest) GetLimits() *CreateRequest_Limits {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *CreateRequest) GetAutoScale() *CreateRequest_AutoScale {
	if m != nil {
		return m.AutoScale
	}
	return nil
}

type CreateRequest_Limits struct {
	Default        []*CreateRequest_Limits_LimitRangeQuantity `protobuf:"bytes,1,rep,name=default" json:"default,omitempty"`
	DefaultRequest []*CreateRequest_Limits_LimitRangeQuantity `protobuf:"bytes,2,rep,name=default_request,json=defaultRequest" json:"default_request,omitempty"`
}

func (m *CreateRequest_Limits) Reset()                    { *m = CreateRequest_Limits{} }
func (m *CreateRequest_Limits) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_Limits) ProtoMessage()               {}
func (*CreateRequest_Limits) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *CreateRequest_Limits) GetDefault() []*CreateRequest_Limits_LimitRangeQuantity {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *CreateRequest_Limits) GetDefaultRequest() []*CreateRequest_Limits_LimitRangeQuantity {
	if m != nil {
		return m.DefaultRequest
	}
	return nil
}

type CreateRequest_Limits_LimitRangeQuantity struct {
	Quantity string `protobuf:"bytes,1,opt,name=quantity" json:"quantity,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *CreateRequest_Limits_LimitRangeQuantity) Reset() {
	*m = CreateRequest_Limits_LimitRangeQuantity{}
}
func (m *CreateRequest_Limits_LimitRangeQuantity) String() string { return proto.CompactTextString(m) }
func (*CreateRequest_Limits_LimitRangeQuantity) ProtoMessage()    {}
func (*CreateRequest_Limits_LimitRangeQuantity) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0, 0}
}

func (m *CreateRequest_Limits_LimitRangeQuantity) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *CreateRequest_Limits_LimitRangeQuantity) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type CreateRequest_AutoScale struct {
	CpuTargetUtilization int64 `protobuf:"varint,1,opt,name=cpu_target_utilization,json=cpuTargetUtilization" json:"cpu_target_utilization,omitempty"`
	Max                  int64 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min                  int64 `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
}

func (m *CreateRequest_AutoScale) Reset()                    { *m = CreateRequest_AutoScale{} }
func (m *CreateRequest_AutoScale) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_AutoScale) ProtoMessage()               {}
func (*CreateRequest_AutoScale) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *CreateRequest_AutoScale) GetCpuTargetUtilization() int64 {
	if m != nil {
		return m.CpuTargetUtilization
	}
	return 0
}

func (m *CreateRequest_AutoScale) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *CreateRequest_AutoScale) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "app.CreateRequest")
	proto.RegisterType((*CreateRequest_Limits)(nil), "app.CreateRequest.Limits")
	proto.RegisterType((*CreateRequest_Limits_LimitRangeQuantity)(nil), "app.CreateRequest.Limits.LimitRangeQuantity")
	proto.RegisterType((*CreateRequest_AutoScale)(nil), "app.CreateRequest.AutoScale")
	proto.RegisterType((*Empty)(nil), "app.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for App service

type AppClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error)
}

type appClient struct {
	cc *grpc.ClientConn
}

func NewAppClient(cc *grpc.ClientConn) AppClient {
	return &appClient{cc}
}

func (c *appClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppServer interface {
	Create(context.Context, *CreateRequest) (*Empty, error)
}

func RegisterAppServer(s *grpc.Server, srv AppServer) {
	s.RegisterService(&_App_serviceDesc, srv)
}

func _App_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _App_serviceDesc = grpc.ServiceDesc{
	ServiceName: "app.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _App_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/app/app.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/app/app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0x4b, 0x4f, 0xea, 0x40,
	0x14, 0x4e, 0x29, 0x94, 0xcb, 0xe1, 0xbe, 0x32, 0xb9, 0xb9, 0xe9, 0x6d, 0xee, 0x02, 0x59, 0x75,
	0x61, 0x20, 0xa2, 0x3b, 0x57, 0xc4, 0xe8, 0x8a, 0x8d, 0x23, 0xac, 0x9b, 0xa1, 0x1e, 0xc8, 0xc4,
	0x3e, 0x86, 0xf6, 0x4c, 0x62, 0xfd, 0x47, 0xfe, 0x40, 0xf7, 0x66, 0xa6, 0x03, 0xc6, 0xa0, 0x1b,
	0x17, 0x4d, 0xbf, 0xf9, 0x5e, 0x27, 0x3d, 0x53, 0x88, 0xd4, 0xc3, 0x76, 0xaa, 0xaa, 0x92, 0xca,
	0xb5, 0xde, 0x4c, 0x85, 0x52, 0xe6, 0x99, 0x58, 0x82, 0xf9, 0x42, 0xa9, 0xf1, 0x73, 0x17, 0x7e,
	0x5c, 0x55, 0x28, 0x08, 0x39, 0xee, 0x34, 0xd6, 0xc4, 0x18, 0x74, 0x0b, 0x91, 0x63, 0xe8, 0x8d,
	0xbc, 0x78, 0xc0, 0x2d, 0x36, 0x1c, 0xa1, 0xc8, 0xc3, 0x4e, 0xcb, 0x19, 0xcc, 0x4e, 0xe0, 0xbb,
	0xaa, 0xca, 0x14, 0xeb, 0x3a, 0xa1, 0x46, 0x61, 0xe8, 0x5b, 0x6d, 0xe8, 0xb8, 0x65, 0xa3, 0x90,
	0x9d, 0x41, 0x90, 0xc9, 0x5c, 0x52, 0x1d, 0x76, 0x47, 0x5e, 0x3c, 0x9c, 0xfd, 0x9b, 0x98, 0xe9,
	0xef, 0xc6, 0x4d, 0x16, 0xd6, 0xc0, 0x9d, 0x91, 0x5d, 0x02, 0x08, 0x4d, 0x65, 0x52, 0xa7, 0x22,
	0xc3, 0xb0, 0x67, 0x63, 0xff, 0x3f, 0x88, 0xcd, 0x35, 0x95, 0x77, 0xc6, 0xc3, 0x07, 0x62, 0x0f,
	0xa3, 0x17, 0x0f, 0x82, 0xb6, 0x8f, 0xdd, 0x40, 0xff, 0x1e, 0x37, 0x42, 0x67, 0x14, 0x7a, 0x23,
	0x3f, 0x1e, 0xce, 0x4e, 0x3f, 0x9d, 0xdd, 0xbe, 0xb8, 0x28, 0xb6, 0x78, 0xab, 0x45, 0x41, 0x92,
	0x1a, 0xbe, 0x0f, 0xb3, 0x15, 0xfc, 0x72, 0x30, 0xa9, 0xda, 0x54, 0xd8, 0xf9, 0x42, 0xdf, 0x4f,
	0x57, 0xe2, 0x9c, 0xd1, 0x02, 0xd8, 0xb1, 0x8b, 0x45, 0xf0, 0x6d, 0xe7, 0xb0, 0x5b, 0xff, 0xe1,
	0x6c, 0xb4, 0x0a, 0xeb, 0x52, 0x57, 0x29, 0xba, 0x6b, 0x38, 0x9c, 0x23, 0x84, 0xc1, 0x61, 0x1f,
	0xec, 0x02, 0xfe, 0xa6, 0x4a, 0x27, 0x24, 0xaa, 0x2d, 0x52, 0xa2, 0x49, 0x66, 0xf2, 0x49, 0x90,
	0x2c, 0x0b, 0x5b, 0xe9, 0xf3, 0x3f, 0xa9, 0xd2, 0x4b, 0x2b, 0xae, 0xde, 0x34, 0xf6, 0x1b, 0xfc,
	0x5c, 0x3c, 0xda, 0x66, 0x9f, 0x1b, 0x68, 0x19, 0x59, 0xd8, 0x6b, 0x35, 0x8c, 0x2c, 0xc6, 0x7d,
	0xe8, 0x5d, 0xe7, 0x8a, 0x9a, 0xd9, 0x14, 0xfc, 0xb9, 0x52, 0x2c, 0x86, 0xa0, 0xfd, 0x7e, 0xc6,
	0x8e, 0x97, 0x11, 0x81, 0xe5, 0x6c, 0x60, 0x1d, 0xd8, 0x3f, 0xee, 0xfc, 0x35, 0x00, 0x00, 0xff,
	0xff, 0xde, 0x2f, 0xa0, 0x35, 0x8f, 0x02, 0x00, 0x00,
}

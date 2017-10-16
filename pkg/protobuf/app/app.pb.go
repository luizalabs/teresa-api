// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/protobuf/app/app.proto

/*
Package app is a generated protocol buffer package.

It is generated from these files:
	pkg/protobuf/app/app.proto

It has these top-level messages:
	CreateRequest
	ListResponse
	LogsRequest
	LogsResponse
	InfoRequest
	InfoResponse
	SetEnvRequest
	UnsetEnvRequest
	SetAutoscaleRequest
	DeleteRequest
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
	Host        string                   `protobuf:"bytes,6,opt,name=host" json:"host,omitempty"`
	Limits      *CreateRequest_Limits    `protobuf:"bytes,4,opt,name=limits" json:"limits,omitempty"`
	Autoscale   *CreateRequest_Autoscale `protobuf:"bytes,5,opt,name=autoscale" json:"autoscale,omitempty"`
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

func (m *CreateRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *CreateRequest) GetLimits() *CreateRequest_Limits {
	if m != nil {
		return m.Limits
	}
	return nil
}

func (m *CreateRequest) GetAutoscale() *CreateRequest_Autoscale {
	if m != nil {
		return m.Autoscale
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

type CreateRequest_Autoscale struct {
	CpuTargetUtilization int32 `protobuf:"varint,1,opt,name=cpu_target_utilization,json=cpuTargetUtilization" json:"cpu_target_utilization,omitempty"`
	Max                  int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min                  int32 `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
}

func (m *CreateRequest_Autoscale) Reset()                    { *m = CreateRequest_Autoscale{} }
func (m *CreateRequest_Autoscale) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest_Autoscale) ProtoMessage()               {}
func (*CreateRequest_Autoscale) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *CreateRequest_Autoscale) GetCpuTargetUtilization() int32 {
	if m != nil {
		return m.CpuTargetUtilization
	}
	return 0
}

func (m *CreateRequest_Autoscale) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *CreateRequest_Autoscale) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

type ListResponse struct {
	Apps []*ListResponse_App `protobuf:"bytes,1,rep,name=apps" json:"apps,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListResponse) GetApps() []*ListResponse_App {
	if m != nil {
		return m.Apps
	}
	return nil
}

type ListResponse_App struct {
	Team string   `protobuf:"bytes,1,opt,name=team" json:"team,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Urls []string `protobuf:"bytes,3,rep,name=urls" json:"urls,omitempty"`
}

func (m *ListResponse_App) Reset()                    { *m = ListResponse_App{} }
func (m *ListResponse_App) String() string            { return proto.CompactTextString(m) }
func (*ListResponse_App) ProtoMessage()               {}
func (*ListResponse_App) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ListResponse_App) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *ListResponse_App) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListResponse_App) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

type LogsRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Lines  int64  `protobuf:"varint,2,opt,name=lines" json:"lines,omitempty"`
	Follow bool   `protobuf:"varint,3,opt,name=follow" json:"follow,omitempty"`
}

func (m *LogsRequest) Reset()                    { *m = LogsRequest{} }
func (m *LogsRequest) String() string            { return proto.CompactTextString(m) }
func (*LogsRequest) ProtoMessage()               {}
func (*LogsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogsRequest) GetLines() int64 {
	if m != nil {
		return m.Lines
	}
	return 0
}

func (m *LogsRequest) GetFollow() bool {
	if m != nil {
		return m.Follow
	}
	return false
}

type LogsResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *LogsResponse) Reset()                    { *m = LogsResponse{} }
func (m *LogsResponse) String() string            { return proto.CompactTextString(m) }
func (*LogsResponse) ProtoMessage()               {}
func (*LogsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogsResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type InfoRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InfoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type InfoResponse struct {
	Team      string                  `protobuf:"bytes,1,opt,name=team" json:"team,omitempty"`
	Addresses []*InfoResponse_Address `protobuf:"bytes,2,rep,name=addresses" json:"addresses,omitempty"`
	EnvVars   []*InfoResponse_EnvVar  `protobuf:"bytes,3,rep,name=env_vars,json=envVars" json:"env_vars,omitempty"`
	Status    *InfoResponse_Status    `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Autoscale *InfoResponse_Autoscale `protobuf:"bytes,5,opt,name=autoscale" json:"autoscale,omitempty"`
	Limits    *InfoResponse_Limits    `protobuf:"bytes,6,opt,name=limits" json:"limits,omitempty"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InfoResponse) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *InfoResponse) GetAddresses() []*InfoResponse_Address {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *InfoResponse) GetEnvVars() []*InfoResponse_EnvVar {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

func (m *InfoResponse) GetStatus() *InfoResponse_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *InfoResponse) GetAutoscale() *InfoResponse_Autoscale {
	if m != nil {
		return m.Autoscale
	}
	return nil
}

func (m *InfoResponse) GetLimits() *InfoResponse_Limits {
	if m != nil {
		return m.Limits
	}
	return nil
}

type InfoResponse_Address struct {
	Hostname string `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
}

func (m *InfoResponse_Address) Reset()                    { *m = InfoResponse_Address{} }
func (m *InfoResponse_Address) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Address) ProtoMessage()               {}
func (*InfoResponse_Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *InfoResponse_Address) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

type InfoResponse_EnvVar struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *InfoResponse_EnvVar) Reset()                    { *m = InfoResponse_EnvVar{} }
func (m *InfoResponse_EnvVar) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_EnvVar) ProtoMessage()               {}
func (*InfoResponse_EnvVar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *InfoResponse_EnvVar) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *InfoResponse_EnvVar) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type InfoResponse_Status struct {
	Cpu  int32                      `protobuf:"varint,1,opt,name=cpu" json:"cpu,omitempty"`
	Pods []*InfoResponse_Status_Pod `protobuf:"bytes,3,rep,name=pods" json:"pods,omitempty"`
}

func (m *InfoResponse_Status) Reset()                    { *m = InfoResponse_Status{} }
func (m *InfoResponse_Status) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Status) ProtoMessage()               {}
func (*InfoResponse_Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2} }

func (m *InfoResponse_Status) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *InfoResponse_Status) GetPods() []*InfoResponse_Status_Pod {
	if m != nil {
		return m.Pods
	}
	return nil
}

type InfoResponse_Status_Pod struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	State    string `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	Age      int64  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
	Restarts int32  `protobuf:"varint,4,opt,name=restarts" json:"restarts,omitempty"`
}

func (m *InfoResponse_Status_Pod) Reset()                    { *m = InfoResponse_Status_Pod{} }
func (m *InfoResponse_Status_Pod) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Status_Pod) ProtoMessage()               {}
func (*InfoResponse_Status_Pod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2, 0} }

func (m *InfoResponse_Status_Pod) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InfoResponse_Status_Pod) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *InfoResponse_Status_Pod) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *InfoResponse_Status_Pod) GetRestarts() int32 {
	if m != nil {
		return m.Restarts
	}
	return 0
}

type InfoResponse_Autoscale struct {
	CpuTargetUtilization int32 `protobuf:"varint,1,opt,name=cpu_target_utilization,json=cpuTargetUtilization" json:"cpu_target_utilization,omitempty"`
	Max                  int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min                  int32 `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
}

func (m *InfoResponse_Autoscale) Reset()                    { *m = InfoResponse_Autoscale{} }
func (m *InfoResponse_Autoscale) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Autoscale) ProtoMessage()               {}
func (*InfoResponse_Autoscale) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 3} }

func (m *InfoResponse_Autoscale) GetCpuTargetUtilization() int32 {
	if m != nil {
		return m.CpuTargetUtilization
	}
	return 0
}

func (m *InfoResponse_Autoscale) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *InfoResponse_Autoscale) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

type InfoResponse_Limits struct {
	Default        []*InfoResponse_Limits_LimitRangeQuantity `protobuf:"bytes,1,rep,name=default" json:"default,omitempty"`
	DefaultRequest []*InfoResponse_Limits_LimitRangeQuantity `protobuf:"bytes,2,rep,name=default_request,json=defaultRequest" json:"default_request,omitempty"`
}

func (m *InfoResponse_Limits) Reset()                    { *m = InfoResponse_Limits{} }
func (m *InfoResponse_Limits) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse_Limits) ProtoMessage()               {}
func (*InfoResponse_Limits) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 4} }

func (m *InfoResponse_Limits) GetDefault() []*InfoResponse_Limits_LimitRangeQuantity {
	if m != nil {
		return m.Default
	}
	return nil
}

func (m *InfoResponse_Limits) GetDefaultRequest() []*InfoResponse_Limits_LimitRangeQuantity {
	if m != nil {
		return m.DefaultRequest
	}
	return nil
}

type InfoResponse_Limits_LimitRangeQuantity struct {
	Quantity string `protobuf:"bytes,1,opt,name=quantity" json:"quantity,omitempty"`
	Resource string `protobuf:"bytes,2,opt,name=resource" json:"resource,omitempty"`
}

func (m *InfoResponse_Limits_LimitRangeQuantity) Reset() {
	*m = InfoResponse_Limits_LimitRangeQuantity{}
}
func (m *InfoResponse_Limits_LimitRangeQuantity) String() string { return proto.CompactTextString(m) }
func (*InfoResponse_Limits_LimitRangeQuantity) ProtoMessage()    {}
func (*InfoResponse_Limits_LimitRangeQuantity) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 4, 0}
}

func (m *InfoResponse_Limits_LimitRangeQuantity) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *InfoResponse_Limits_LimitRangeQuantity) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

type SetEnvRequest struct {
	Name    string                  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	EnvVars []*SetEnvRequest_EnvVar `protobuf:"bytes,2,rep,name=env_vars,json=envVars" json:"env_vars,omitempty"`
}

func (m *SetEnvRequest) Reset()                    { *m = SetEnvRequest{} }
func (m *SetEnvRequest) String() string            { return proto.CompactTextString(m) }
func (*SetEnvRequest) ProtoMessage()               {}
func (*SetEnvRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SetEnvRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetEnvRequest) GetEnvVars() []*SetEnvRequest_EnvVar {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

type SetEnvRequest_EnvVar struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *SetEnvRequest_EnvVar) Reset()                    { *m = SetEnvRequest_EnvVar{} }
func (m *SetEnvRequest_EnvVar) String() string            { return proto.CompactTextString(m) }
func (*SetEnvRequest_EnvVar) ProtoMessage()               {}
func (*SetEnvRequest_EnvVar) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *SetEnvRequest_EnvVar) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetEnvRequest_EnvVar) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UnsetEnvRequest struct {
	Name    string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	EnvVars []string `protobuf:"bytes,2,rep,name=env_vars,json=envVars" json:"env_vars,omitempty"`
}

func (m *UnsetEnvRequest) Reset()                    { *m = UnsetEnvRequest{} }
func (m *UnsetEnvRequest) String() string            { return proto.CompactTextString(m) }
func (*UnsetEnvRequest) ProtoMessage()               {}
func (*UnsetEnvRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UnsetEnvRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UnsetEnvRequest) GetEnvVars() []string {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

type SetAutoscaleRequest struct {
	Name      string                         `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Autoscale *SetAutoscaleRequest_Autoscale `protobuf:"bytes,2,opt,name=autoscale" json:"autoscale,omitempty"`
}

func (m *SetAutoscaleRequest) Reset()                    { *m = SetAutoscaleRequest{} }
func (m *SetAutoscaleRequest) String() string            { return proto.CompactTextString(m) }
func (*SetAutoscaleRequest) ProtoMessage()               {}
func (*SetAutoscaleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SetAutoscaleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetAutoscaleRequest) GetAutoscale() *SetAutoscaleRequest_Autoscale {
	if m != nil {
		return m.Autoscale
	}
	return nil
}

type SetAutoscaleRequest_Autoscale struct {
	CpuTargetUtilization int32 `protobuf:"varint,1,opt,name=cpu_target_utilization,json=cpuTargetUtilization" json:"cpu_target_utilization,omitempty"`
	Max                  int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
	Min                  int32 `protobuf:"varint,3,opt,name=min" json:"min,omitempty"`
}

func (m *SetAutoscaleRequest_Autoscale) Reset()         { *m = SetAutoscaleRequest_Autoscale{} }
func (m *SetAutoscaleRequest_Autoscale) String() string { return proto.CompactTextString(m) }
func (*SetAutoscaleRequest_Autoscale) ProtoMessage()    {}
func (*SetAutoscaleRequest_Autoscale) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

func (m *SetAutoscaleRequest_Autoscale) GetCpuTargetUtilization() int32 {
	if m != nil {
		return m.CpuTargetUtilization
	}
	return 0
}

func (m *SetAutoscaleRequest_Autoscale) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *SetAutoscaleRequest_Autoscale) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

type DeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "app.CreateRequest")
	proto.RegisterType((*CreateRequest_Limits)(nil), "app.CreateRequest.Limits")
	proto.RegisterType((*CreateRequest_Limits_LimitRangeQuantity)(nil), "app.CreateRequest.Limits.LimitRangeQuantity")
	proto.RegisterType((*CreateRequest_Autoscale)(nil), "app.CreateRequest.Autoscale")
	proto.RegisterType((*ListResponse)(nil), "app.ListResponse")
	proto.RegisterType((*ListResponse_App)(nil), "app.ListResponse.App")
	proto.RegisterType((*LogsRequest)(nil), "app.LogsRequest")
	proto.RegisterType((*LogsResponse)(nil), "app.LogsResponse")
	proto.RegisterType((*InfoRequest)(nil), "app.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "app.InfoResponse")
	proto.RegisterType((*InfoResponse_Address)(nil), "app.InfoResponse.Address")
	proto.RegisterType((*InfoResponse_EnvVar)(nil), "app.InfoResponse.EnvVar")
	proto.RegisterType((*InfoResponse_Status)(nil), "app.InfoResponse.Status")
	proto.RegisterType((*InfoResponse_Status_Pod)(nil), "app.InfoResponse.Status.Pod")
	proto.RegisterType((*InfoResponse_Autoscale)(nil), "app.InfoResponse.Autoscale")
	proto.RegisterType((*InfoResponse_Limits)(nil), "app.InfoResponse.Limits")
	proto.RegisterType((*InfoResponse_Limits_LimitRangeQuantity)(nil), "app.InfoResponse.Limits.LimitRangeQuantity")
	proto.RegisterType((*SetEnvRequest)(nil), "app.SetEnvRequest")
	proto.RegisterType((*SetEnvRequest_EnvVar)(nil), "app.SetEnvRequest.EnvVar")
	proto.RegisterType((*UnsetEnvRequest)(nil), "app.UnsetEnvRequest")
	proto.RegisterType((*SetAutoscaleRequest)(nil), "app.SetAutoscaleRequest")
	proto.RegisterType((*SetAutoscaleRequest_Autoscale)(nil), "app.SetAutoscaleRequest.Autoscale")
	proto.RegisterType((*DeleteRequest)(nil), "app.DeleteRequest")
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
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (App_LogsClient, error)
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	SetEnv(ctx context.Context, in *SetEnvRequest, opts ...grpc.CallOption) (*Empty, error)
	UnsetEnv(ctx context.Context, in *UnsetEnvRequest, opts ...grpc.CallOption) (*Empty, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error)
	SetAutoscale(ctx context.Context, in *SetAutoscaleRequest, opts ...grpc.CallOption) (*Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error)
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

func (c *appClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (App_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_App_serviceDesc.Streams[0], c.cc, "/app.App/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &appLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type App_LogsClient interface {
	Recv() (*LogsResponse, error)
	grpc.ClientStream
}

type appLogsClient struct {
	grpc.ClientStream
}

func (x *appLogsClient) Recv() (*LogsResponse, error) {
	m := new(LogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/app.App/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) SetEnv(ctx context.Context, in *SetEnvRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/SetEnv", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) UnsetEnv(ctx context.Context, in *UnsetEnvRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/UnsetEnv", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/app.App/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) SetAutoscale(ctx context.Context, in *SetAutoscaleRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/SetAutoscale", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/app.App/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for App service

type AppServer interface {
	Create(context.Context, *CreateRequest) (*Empty, error)
	Logs(*LogsRequest, App_LogsServer) error
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	SetEnv(context.Context, *SetEnvRequest) (*Empty, error)
	UnsetEnv(context.Context, *UnsetEnvRequest) (*Empty, error)
	List(context.Context, *Empty) (*ListResponse, error)
	SetAutoscale(context.Context, *SetAutoscaleRequest) (*Empty, error)
	Delete(context.Context, *DeleteRequest) (*Empty, error)
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

func _App_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AppServer).Logs(m, &appLogsServer{stream})
}

type App_LogsServer interface {
	Send(*LogsResponse) error
	grpc.ServerStream
}

type appLogsServer struct {
	grpc.ServerStream
}

func (x *appLogsServer) Send(m *LogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _App_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_SetEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).SetEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/SetEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).SetEnv(ctx, req.(*SetEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_UnsetEnv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetEnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).UnsetEnv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/UnsetEnv",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).UnsetEnv(ctx, req.(*UnsetEnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_SetAutoscale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAutoscaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).SetAutoscale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/SetAutoscale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).SetAutoscale(ctx, req.(*SetAutoscaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _App_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/app.App/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).Delete(ctx, req.(*DeleteRequest))
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
		{
			MethodName: "Info",
			Handler:    _App_Info_Handler,
		},
		{
			MethodName: "SetEnv",
			Handler:    _App_SetEnv_Handler,
		},
		{
			MethodName: "UnsetEnv",
			Handler:    _App_UnsetEnv_Handler,
		},
		{
			MethodName: "List",
			Handler:    _App_List_Handler,
		},
		{
			MethodName: "SetAutoscale",
			Handler:    _App_SetAutoscale_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _App_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _App_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/protobuf/app/app.proto",
}

func init() { proto.RegisterFile("pkg/protobuf/app/app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 889 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0x97, 0x63, 0xc7, 0x49, 0x26, 0x39, 0x5a, 0x96, 0xa3, 0x72, 0x4d, 0x1f, 0x52, 0xa3, 0x4a,
	0x41, 0x2d, 0x69, 0x48, 0x4f, 0x42, 0xf0, 0xd4, 0x08, 0x82, 0x84, 0x14, 0x89, 0xb2, 0x77, 0xc7,
	0x1b, 0x8a, 0xb6, 0xc9, 0x26, 0x58, 0x75, 0xec, 0xad, 0x77, 0x1d, 0x2e, 0x88, 0x6f, 0xc0, 0x23,
	0x1f, 0x83, 0xaf, 0xc0, 0xa7, 0xe0, 0x63, 0xf0, 0x8a, 0x78, 0x47, 0xfb, 0xc7, 0x8e, 0x9d, 0x7f,
	0xe8, 0x90, 0xb8, 0x87, 0x28, 0xb3, 0xb3, 0xbf, 0x99, 0x9d, 0x9d, 0xfd, 0xcd, 0x8c, 0xc1, 0x67,
	0x6f, 0x96, 0xcf, 0x59, 0x9a, 0x88, 0xe4, 0x75, 0xb6, 0x78, 0x4e, 0x18, 0x93, 0xbf, 0xbe, 0x52,
	0x20, 0x9b, 0x30, 0x16, 0xfc, 0xee, 0xc0, 0xd9, 0x17, 0x29, 0x25, 0x82, 0x62, 0xfa, 0x36, 0xa3,
	0x5c, 0x20, 0x04, 0x4e, 0x4c, 0x56, 0xd4, 0xb3, 0xba, 0x56, 0xaf, 0x85, 0x95, 0x2c, 0x75, 0x82,
	0x92, 0x95, 0x57, 0xd3, 0x3a, 0x29, 0xa3, 0xc7, 0xd0, 0x61, 0x69, 0x32, 0xa3, 0x9c, 0x4f, 0xc5,
	0x86, 0x51, 0xcf, 0x56, 0x7b, 0x6d, 0xa3, 0xbb, 0xda, 0x30, 0x65, 0xf6, 0x43, 0xc2, 0x85, 0xe7,
	0x6a, 0x33, 0x29, 0xa3, 0x4f, 0xc0, 0x8d, 0xc2, 0x55, 0x28, 0xb8, 0xe7, 0x74, 0xad, 0x5e, 0x7b,
	0xf8, 0xb0, 0x2f, 0x23, 0xaa, 0x84, 0xd0, 0x9f, 0x28, 0x00, 0x36, 0x40, 0xf4, 0x39, 0xb4, 0x48,
	0x26, 0x12, 0x3e, 0x23, 0x11, 0xf5, 0xea, 0xca, 0xea, 0xd1, 0x01, 0xab, 0x51, 0x8e, 0xc1, 0x5b,
	0xb8, 0xff, 0xb7, 0x05, 0xae, 0x76, 0x87, 0xbe, 0x82, 0xc6, 0x9c, 0x2e, 0x48, 0x16, 0x09, 0xcf,
	0xea, 0xda, 0xbd, 0xf6, 0xf0, 0xd9, 0xd1, 0xa3, 0xf5, 0x1f, 0x26, 0xf1, 0x92, 0x7e, 0x9b, 0x91,
	0x58, 0x84, 0x62, 0x83, 0x73, 0x63, 0x74, 0x0d, 0xf7, 0x8c, 0x38, 0x4d, 0xb5, 0x95, 0x57, 0xfb,
	0x0f, 0xfe, 0xde, 0x31, 0x4e, 0x0c, 0xd2, 0x9f, 0x00, 0xda, 0x47, 0x21, 0x1f, 0x9a, 0x6f, 0x8d,
	0x6c, 0x5e, 0xa4, 0x58, 0xcb, 0xbd, 0x94, 0xf2, 0x24, 0x4b, 0x67, 0xd4, 0xbc, 0x4c, 0xb1, 0xf6,
	0x29, 0xb4, 0x8a, 0x7c, 0xa0, 0x0b, 0x78, 0x30, 0x63, 0xd9, 0x54, 0x90, 0x74, 0x49, 0xc5, 0x34,
	0x13, 0x61, 0x14, 0xfe, 0x44, 0x44, 0x98, 0xc4, 0xca, 0x65, 0x1d, 0x9f, 0xcf, 0x58, 0x76, 0xa5,
	0x36, 0xaf, 0xb7, 0x7b, 0xe8, 0x3e, 0xd8, 0x2b, 0x72, 0xa3, 0x3c, 0xd7, 0xb1, 0x14, 0x95, 0x26,
	0x8c, 0xd5, 0x4b, 0x4b, 0x4d, 0x18, 0x07, 0x3f, 0x43, 0x67, 0x12, 0x72, 0x81, 0x29, 0x67, 0x49,
	0xcc, 0x29, 0xfa, 0x08, 0x1c, 0xc2, 0x18, 0x37, 0x09, 0x7e, 0x5f, 0x25, 0xa4, 0x0c, 0xe8, 0x8f,
	0x18, 0xc3, 0x0a, 0xe2, 0x8f, 0xc0, 0x1e, 0x31, 0x56, 0x50, 0xcb, 0x2a, 0x51, 0x2b, 0xa7, 0x60,
	0xad, 0x4a, 0xc1, 0x2c, 0x8d, 0xb8, 0x67, 0x77, 0x6d, 0xa9, 0x93, 0x72, 0xf0, 0x0d, 0xb4, 0x27,
	0xc9, 0x92, 0x9f, 0x62, 0xee, 0x39, 0xd4, 0xa3, 0x30, 0xa6, 0x5c, 0xf9, 0xb2, 0xb1, 0x5e, 0xa0,
	0x07, 0xe0, 0x2e, 0x92, 0x28, 0x4a, 0x7e, 0x54, 0x77, 0x69, 0x62, 0xb3, 0x0a, 0x02, 0xe8, 0x68,
	0x87, 0xe6, 0x3a, 0x2a, 0xb8, 0x1b, 0xb1, 0x0d, 0xee, 0x46, 0x04, 0x8f, 0xa1, 0xfd, 0x75, 0xbc,
	0x48, 0x4e, 0x1c, 0x1a, 0xfc, 0xda, 0x80, 0x8e, 0xc6, 0x94, 0xfd, 0xec, 0x5c, 0xf2, 0x53, 0x68,
	0x91, 0xf9, 0x3c, 0xa5, 0x9c, 0xab, 0xe8, 0xec, 0xa2, 0x16, 0xca, 0x96, 0xfd, 0x91, 0x86, 0xe0,
	0x2d, 0x16, 0xbd, 0x80, 0x26, 0x8d, 0xd7, 0xd3, 0x35, 0x49, 0x75, 0x36, 0xda, 0x43, 0x6f, 0xdf,
	0x6e, 0x1c, 0xaf, 0xbf, 0x23, 0x29, 0x6e, 0x50, 0xf5, 0xcf, 0xd1, 0x00, 0x5c, 0x2e, 0x88, 0xc8,
	0xf2, 0xb2, 0x3b, 0x60, 0x72, 0xa9, 0xf6, 0xb1, 0xc1, 0xa1, 0xcf, 0xf6, 0xab, 0xee, 0x83, 0x03,
	0xf1, 0x1d, 0x28, 0x3a, 0x79, 0x98, 0xa9, 0x71, 0xf7, 0xd8, 0x61, 0xd5, 0x12, 0xf7, 0x9f, 0x40,
	0xc3, 0xdc, 0x54, 0xb2, 0x5a, 0x36, 0x8a, 0x52, 0x52, 0x8b, 0xb5, 0x3f, 0x00, 0x57, 0x5f, 0x4c,
	0x52, 0xf1, 0x0d, 0xcd, 0x4b, 0x42, 0x8a, 0xf2, 0xa5, 0xd7, 0x24, 0xca, 0x72, 0xd6, 0xe8, 0x85,
	0xff, 0x9b, 0x05, 0xae, 0xbe, 0x98, 0x34, 0x99, 0xb1, 0xcc, 0x50, 0x5e, 0x8a, 0x68, 0x00, 0x0e,
	0x4b, 0xe6, 0x79, 0x16, 0x1f, 0x1d, 0x4b, 0x49, 0xff, 0x55, 0x32, 0xc7, 0x0a, 0xe9, 0x7f, 0x0f,
	0xf6, 0xab, 0x64, 0x7e, 0x8c, 0x69, 0x32, 0x73, 0xc5, 0xf9, 0x6a, 0x21, 0x0f, 0x25, 0x4b, 0xdd,
	0x1c, 0x6d, 0x2c, 0x45, 0x53, 0xb5, 0x82, 0xa4, 0xa6, 0x05, 0xd6, 0x71, 0xb1, 0xbe, 0xa3, 0xaa,
	0xf5, 0xff, 0xda, 0x36, 0xc5, 0xf1, 0x6e, 0x53, 0x7c, 0x7a, 0xec, 0xad, 0x4e, 0xf6, 0xc4, 0xab,
	0x63, 0x3d, 0xf1, 0x56, 0xee, 0xfe, 0xd7, 0x96, 0x18, 0xfc, 0x62, 0xc1, 0xd9, 0x25, 0x15, 0xe3,
	0x78, 0x7d, 0xaa, 0x61, 0x5c, 0x94, 0xaa, 0xab, 0x5c, 0x95, 0x15, 0xcb, 0xdd, 0xf2, 0xba, 0x3d,
	0x31, 0x83, 0x97, 0x70, 0xef, 0x3a, 0xe6, 0xff, 0x1a, 0xce, 0xc3, 0x9d, 0x70, 0x5a, 0xc5, 0x99,
	0xc1, 0x1f, 0x16, 0xbc, 0x77, 0x49, 0xc5, 0xb6, 0x02, 0x4f, 0xb8, 0x79, 0x59, 0x2e, 0xe6, 0x9a,
	0x2a, 0xca, 0x20, 0xbf, 0xd6, 0xae, 0x83, 0xc3, 0x83, 0xf4, 0x8e, 0x06, 0xca, 0x87, 0x70, 0xf6,
	0x25, 0x8d, 0xe8, 0xc9, 0xcf, 0x91, 0xa0, 0x01, 0xf5, 0xf1, 0x8a, 0x89, 0xcd, 0xf0, 0xcf, 0x9a,
	0x1e, 0x22, 0x3d, 0x70, 0xf5, 0xd8, 0x45, 0x68, 0x7f, 0x06, 0xfb, 0xa0, 0x74, 0xca, 0x02, 0x7d,
	0x0c, 0x8e, 0xec, 0xf0, 0xe8, 0xbe, 0x1e, 0x4d, 0xdb, 0xe9, 0xe1, 0xbf, 0x5b, 0xd2, 0x68, 0xa6,
	0x0e, 0x2c, 0xf4, 0x14, 0x1c, 0xc9, 0x5d, 0x03, 0x2f, 0xf5, 0x7d, 0x03, 0xaf, 0x74, 0xf9, 0x1e,
	0xb8, 0x9a, 0x25, 0x26, 0x8a, 0x0a, 0x65, 0x2a, 0x51, 0x3c, 0x83, 0x66, 0xfe, 0xf8, 0xe8, 0x5c,
	0xe9, 0x77, 0xb8, 0x50, 0x41, 0x3f, 0x01, 0x47, 0xce, 0x50, 0x54, 0xd2, 0xe5, 0xd1, 0x96, 0x67,
	0xef, 0x05, 0x74, 0xca, 0xaf, 0x89, 0xbc, 0x63, 0x0f, 0x5c, 0x71, 0xde, 0x03, 0x57, 0x27, 0xdc,
	0x04, 0x5d, 0xc9, 0x7e, 0x19, 0xf9, 0xda, 0x55, 0x9f, 0x8d, 0x2f, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x14, 0x9c, 0x89, 0x4b, 0x54, 0x0a, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/scheduler.proto

package scheduler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExecutionMethod int32

const (
	ExecutionMethod_EXECUTION_METHOD_INVALID ExecutionMethod = 0
	ExecutionMethod_EXECUTION_METHOD_POST    ExecutionMethod = 1
	ExecutionMethod_EXECUTION_METHOD_GET     ExecutionMethod = 2
)

var ExecutionMethod_name = map[int32]string{
	0: "EXECUTION_METHOD_INVALID",
	1: "EXECUTION_METHOD_POST",
	2: "EXECUTION_METHOD_GET",
}

var ExecutionMethod_value = map[string]int32{
	"EXECUTION_METHOD_INVALID": 0,
	"EXECUTION_METHOD_POST":    1,
	"EXECUTION_METHOD_GET":     2,
}

func (x ExecutionMethod) String() string {
	return proto.EnumName(ExecutionMethod_name, int32(x))
}

func (ExecutionMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49bf124e81298781, []int{0}
}

type ScheduleRemoteTaskRequest struct {
	Id                   string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Retries              int32                  `protobuf:"varint,2,opt,name=retries,proto3" json:"retries,omitempty"`
	CreateTime           *timestamp.Timestamp   `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	ScheduleTime         *timestamp.Timestamp   `protobuf:"bytes,4,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	Config               *RemoteExecutionConfig `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ScheduleRemoteTaskRequest) Reset()         { *m = ScheduleRemoteTaskRequest{} }
func (m *ScheduleRemoteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*ScheduleRemoteTaskRequest) ProtoMessage()    {}
func (*ScheduleRemoteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49bf124e81298781, []int{0}
}

func (m *ScheduleRemoteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleRemoteTaskRequest.Unmarshal(m, b)
}
func (m *ScheduleRemoteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleRemoteTaskRequest.Marshal(b, m, deterministic)
}
func (m *ScheduleRemoteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleRemoteTaskRequest.Merge(m, src)
}
func (m *ScheduleRemoteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_ScheduleRemoteTaskRequest.Size(m)
}
func (m *ScheduleRemoteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleRemoteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleRemoteTaskRequest proto.InternalMessageInfo

func (m *ScheduleRemoteTaskRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ScheduleRemoteTaskRequest) GetRetries() int32 {
	if m != nil {
		return m.Retries
	}
	return 0
}

func (m *ScheduleRemoteTaskRequest) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ScheduleRemoteTaskRequest) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *ScheduleRemoteTaskRequest) GetConfig() *RemoteExecutionConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type ScheduleRemoteTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleRemoteTaskResponse) Reset()         { *m = ScheduleRemoteTaskResponse{} }
func (m *ScheduleRemoteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*ScheduleRemoteTaskResponse) ProtoMessage()    {}
func (*ScheduleRemoteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49bf124e81298781, []int{1}
}

func (m *ScheduleRemoteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleRemoteTaskResponse.Unmarshal(m, b)
}
func (m *ScheduleRemoteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleRemoteTaskResponse.Marshal(b, m, deterministic)
}
func (m *ScheduleRemoteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleRemoteTaskResponse.Merge(m, src)
}
func (m *ScheduleRemoteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_ScheduleRemoteTaskResponse.Size(m)
}
func (m *ScheduleRemoteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleRemoteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleRemoteTaskResponse proto.InternalMessageInfo

type RemoteExecutionConfig struct {
	Method               ExecutionMethod   `protobuf:"varint,1,opt,name=method,proto3,enum=scheduler.v1.ExecutionMethod" json:"method,omitempty"`
	Url                  string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Headers              map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Timeout              uint32            `protobuf:"varint,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	ExpectedCodes        []uint32          `protobuf:"varint,6,rep,packed,name=expected_codes,json=expectedCodes,proto3" json:"expected_codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RemoteExecutionConfig) Reset()         { *m = RemoteExecutionConfig{} }
func (m *RemoteExecutionConfig) String() string { return proto.CompactTextString(m) }
func (*RemoteExecutionConfig) ProtoMessage()    {}
func (*RemoteExecutionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_49bf124e81298781, []int{2}
}

func (m *RemoteExecutionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteExecutionConfig.Unmarshal(m, b)
}
func (m *RemoteExecutionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteExecutionConfig.Marshal(b, m, deterministic)
}
func (m *RemoteExecutionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteExecutionConfig.Merge(m, src)
}
func (m *RemoteExecutionConfig) XXX_Size() int {
	return xxx_messageInfo_RemoteExecutionConfig.Size(m)
}
func (m *RemoteExecutionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteExecutionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteExecutionConfig proto.InternalMessageInfo

func (m *RemoteExecutionConfig) GetMethod() ExecutionMethod {
	if m != nil {
		return m.Method
	}
	return ExecutionMethod_EXECUTION_METHOD_INVALID
}

func (m *RemoteExecutionConfig) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RemoteExecutionConfig) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *RemoteExecutionConfig) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *RemoteExecutionConfig) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RemoteExecutionConfig) GetExpectedCodes() []uint32 {
	if m != nil {
		return m.ExpectedCodes
	}
	return nil
}

type ClearScheduleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearScheduleRequest) Reset()         { *m = ClearScheduleRequest{} }
func (m *ClearScheduleRequest) String() string { return proto.CompactTextString(m) }
func (*ClearScheduleRequest) ProtoMessage()    {}
func (*ClearScheduleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_49bf124e81298781, []int{3}
}

func (m *ClearScheduleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearScheduleRequest.Unmarshal(m, b)
}
func (m *ClearScheduleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearScheduleRequest.Marshal(b, m, deterministic)
}
func (m *ClearScheduleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearScheduleRequest.Merge(m, src)
}
func (m *ClearScheduleRequest) XXX_Size() int {
	return xxx_messageInfo_ClearScheduleRequest.Size(m)
}
func (m *ClearScheduleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearScheduleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClearScheduleRequest proto.InternalMessageInfo

type ClearScheduleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClearScheduleResponse) Reset()         { *m = ClearScheduleResponse{} }
func (m *ClearScheduleResponse) String() string { return proto.CompactTextString(m) }
func (*ClearScheduleResponse) ProtoMessage()    {}
func (*ClearScheduleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_49bf124e81298781, []int{4}
}

func (m *ClearScheduleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClearScheduleResponse.Unmarshal(m, b)
}
func (m *ClearScheduleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClearScheduleResponse.Marshal(b, m, deterministic)
}
func (m *ClearScheduleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClearScheduleResponse.Merge(m, src)
}
func (m *ClearScheduleResponse) XXX_Size() int {
	return xxx_messageInfo_ClearScheduleResponse.Size(m)
}
func (m *ClearScheduleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClearScheduleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClearScheduleResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("scheduler.v1.ExecutionMethod", ExecutionMethod_name, ExecutionMethod_value)
	proto.RegisterType((*ScheduleRemoteTaskRequest)(nil), "scheduler.v1.ScheduleRemoteTaskRequest")
	proto.RegisterType((*ScheduleRemoteTaskResponse)(nil), "scheduler.v1.ScheduleRemoteTaskResponse")
	proto.RegisterType((*RemoteExecutionConfig)(nil), "scheduler.v1.RemoteExecutionConfig")
	proto.RegisterMapType((map[string]string)(nil), "scheduler.v1.RemoteExecutionConfig.HeadersEntry")
	proto.RegisterType((*ClearScheduleRequest)(nil), "scheduler.v1.ClearScheduleRequest")
	proto.RegisterType((*ClearScheduleResponse)(nil), "scheduler.v1.ClearScheduleResponse")
}

func init() {
	proto.RegisterFile("v1/scheduler.proto", fileDescriptor_49bf124e81298781)
}

var fileDescriptor_49bf124e81298781 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xfd, 0x25, 0xdd, 0x3a, 0xf5, 0xf6, 0xcf, 0xaf, 0xb2, 0x5a, 0xc8, 0xa2, 0x21, 0xaa, 0x4c,
	0x88, 0x88, 0x87, 0x94, 0x15, 0x21, 0xa1, 0xed, 0x01, 0x41, 0x17, 0xb1, 0x22, 0xb6, 0x22, 0x37,
	0xa0, 0x89, 0x97, 0x2a, 0x4d, 0xee, 0xda, 0x68, 0x6d, 0x5d, 0x1c, 0xa7, 0x5a, 0xdf, 0xf8, 0x62,
	0x7c, 0x0c, 0xbe, 0x0f, 0x8a, 0xdd, 0x74, 0xeb, 0x5a, 0x18, 0x6f, 0xf6, 0xf5, 0x39, 0xf7, 0xf8,
	0x1c, 0x5f, 0x03, 0x99, 0x1f, 0x35, 0xe3, 0x60, 0x84, 0x61, 0x32, 0x46, 0xee, 0xcc, 0x38, 0x13,
	0x8c, 0x94, 0x6e, 0x0b, 0xf3, 0x23, 0xf3, 0xe9, 0x90, 0xb1, 0xe1, 0x18, 0x9b, 0xf2, 0x6c, 0x90,
	0x5c, 0x35, 0x45, 0x34, 0xc1, 0x58, 0xf8, 0x93, 0x99, 0x82, 0x5b, 0x3f, 0x74, 0xd8, 0xef, 0x2d,
	0x19, 0x14, 0x27, 0x4c, 0xa0, 0xe7, 0xc7, 0xd7, 0x14, 0xbf, 0x27, 0x18, 0x0b, 0x52, 0x01, 0x3d,
	0x0a, 0x0d, 0xad, 0xa1, 0xd9, 0x05, 0xaa, 0x47, 0x21, 0x31, 0x60, 0x8f, 0xa3, 0xe0, 0x11, 0xc6,
	0x86, 0xde, 0xd0, 0xec, 0x5d, 0x9a, 0x6d, 0xc9, 0x09, 0x14, 0x03, 0x8e, 0xbe, 0xc0, 0x7e, 0xaa,
	0x60, 0xe4, 0x1a, 0x9a, 0x5d, 0x6c, 0x99, 0x8e, 0x92, 0x77, 0x32, 0x79, 0xc7, 0xcb, 0xe4, 0x29,
	0x28, 0x78, 0x5a, 0x20, 0x6f, 0xa1, 0x9c, 0xdd, 0x5a, 0xd1, 0x77, 0x1e, 0xa4, 0xaf, 0x6c, 0xca,
	0x06, 0x27, 0x90, 0x0f, 0xd8, 0xf4, 0x2a, 0x1a, 0x1a, 0xbb, 0x92, 0x79, 0xe8, 0xdc, 0x4d, 0xc1,
	0x51, 0xc6, 0xdc, 0x1b, 0x0c, 0x12, 0x11, 0xb1, 0x69, 0x5b, 0x42, 0xe9, 0x92, 0x62, 0x1d, 0x80,
	0xb9, 0x2d, 0x81, 0x78, 0xc6, 0xa6, 0x31, 0x5a, 0x3f, 0x75, 0xa8, 0x6f, 0xe5, 0x93, 0xd7, 0x90,
	0x9f, 0xa0, 0x18, 0x31, 0x15, 0x50, 0xa5, 0xf5, 0x64, 0x5d, 0x74, 0x05, 0x3f, 0x97, 0x20, 0xba,
	0x04, 0x93, 0x2a, 0xe4, 0x12, 0x3e, 0x96, 0xf9, 0x15, 0x68, 0xba, 0x24, 0x1f, 0x61, 0x6f, 0x84,
	0x7e, 0x88, 0x3c, 0x36, 0x72, 0x8d, 0x9c, 0x5d, 0x6c, 0xbd, 0xfc, 0x87, 0xeb, 0x3b, 0x67, 0x8a,
	0xe2, 0x4e, 0x05, 0x5f, 0xd0, 0xac, 0x01, 0x21, 0xb0, 0x33, 0x60, 0xe1, 0x42, 0x26, 0x58, 0xa2,
	0x72, 0x9d, 0xbe, 0x5a, 0x9a, 0x2a, 0x4b, 0x84, 0x8c, 0xa7, 0x4c, 0xb3, 0x2d, 0x79, 0x06, 0x15,
	0xbc, 0x99, 0x61, 0x20, 0x30, 0xec, 0x07, 0x2c, 0xc4, 0xd8, 0xc8, 0x37, 0x72, 0x76, 0x99, 0x96,
	0xb3, 0x6a, 0x3b, 0x2d, 0x9a, 0xc7, 0x50, 0xba, 0xab, 0x96, 0x5a, 0xb8, 0xc6, 0xc5, 0x72, 0x2e,
	0xd2, 0x25, 0xa9, 0xc1, 0xee, 0xdc, 0x1f, 0x27, 0xb8, 0xb4, 0xa5, 0x36, 0xc7, 0xfa, 0x1b, 0xcd,
	0x7a, 0x04, 0xb5, 0xf6, 0x18, 0x7d, 0x7e, 0x1b, 0xb1, 0x1c, 0x2d, 0xeb, 0x31, 0xd4, 0xef, 0xd5,
	0x55, 0xe0, 0x2f, 0x42, 0xf8, 0xff, 0x5e, 0x74, 0xe4, 0x00, 0x0c, 0xf7, 0xd2, 0x6d, 0x7f, 0xf1,
	0x3a, 0xdd, 0x8b, 0xfe, 0xb9, 0xeb, 0x9d, 0x75, 0x4f, 0xfb, 0x9d, 0x8b, 0xaf, 0xef, 0x3e, 0x75,
	0x4e, 0xab, 0xff, 0x91, 0x7d, 0xa8, 0x6f, 0x9c, 0x7e, 0xee, 0xf6, 0xbc, 0xaa, 0x46, 0x0c, 0xa8,
	0x6d, 0x1c, 0x7d, 0x70, 0xbd, 0xaa, 0xde, 0xfa, 0xa5, 0x41, 0x35, 0x93, 0xe6, 0x3d, 0xe4, 0xf3,
	0x28, 0x40, 0x32, 0x04, 0xb2, 0x39, 0x09, 0xe4, 0xf9, 0xfa, 0x6b, 0xfc, 0xf1, 0xb7, 0x98, 0xf6,
	0xc3, 0x40, 0xe5, 0x91, 0x5c, 0x42, 0x79, 0xcd, 0x3c, 0xb1, 0xd6, 0xa9, 0xdb, 0x12, 0x33, 0x0f,
	0xff, 0x8a, 0x51, 0x9d, 0xdf, 0x17, 0xbf, 0x15, 0x56, 0xa8, 0x41, 0x5e, 0x7e, 0x9c, 0x57, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xee, 0x66, 0x8b, 0x17, 0x28, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SchedulerServiceClient is the client API for SchedulerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerServiceClient interface {
	ScheduleRemoteTask(ctx context.Context, in *ScheduleRemoteTaskRequest, opts ...grpc.CallOption) (*ScheduleRemoteTaskResponse, error)
	ClearSchedule(ctx context.Context, in *ClearScheduleRequest, opts ...grpc.CallOption) (*ClearScheduleResponse, error)
}

type schedulerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerServiceClient(cc grpc.ClientConnInterface) SchedulerServiceClient {
	return &schedulerServiceClient{cc}
}

func (c *schedulerServiceClient) ScheduleRemoteTask(ctx context.Context, in *ScheduleRemoteTaskRequest, opts ...grpc.CallOption) (*ScheduleRemoteTaskResponse, error) {
	out := new(ScheduleRemoteTaskResponse)
	err := c.cc.Invoke(ctx, "/scheduler.v1.SchedulerService/ScheduleRemoteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerServiceClient) ClearSchedule(ctx context.Context, in *ClearScheduleRequest, opts ...grpc.CallOption) (*ClearScheduleResponse, error) {
	out := new(ClearScheduleResponse)
	err := c.cc.Invoke(ctx, "/scheduler.v1.SchedulerService/ClearSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServiceServer is the server API for SchedulerService service.
type SchedulerServiceServer interface {
	ScheduleRemoteTask(context.Context, *ScheduleRemoteTaskRequest) (*ScheduleRemoteTaskResponse, error)
	ClearSchedule(context.Context, *ClearScheduleRequest) (*ClearScheduleResponse, error)
}

// UnimplementedSchedulerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSchedulerServiceServer struct {
}

func (*UnimplementedSchedulerServiceServer) ScheduleRemoteTask(ctx context.Context, req *ScheduleRemoteTaskRequest) (*ScheduleRemoteTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleRemoteTask not implemented")
}
func (*UnimplementedSchedulerServiceServer) ClearSchedule(ctx context.Context, req *ClearScheduleRequest) (*ClearScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearSchedule not implemented")
}

func RegisterSchedulerServiceServer(s *grpc.Server, srv SchedulerServiceServer) {
	s.RegisterService(&_SchedulerService_serviceDesc, srv)
}

func _SchedulerService_ScheduleRemoteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRemoteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ScheduleRemoteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v1.SchedulerService/ScheduleRemoteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ScheduleRemoteTask(ctx, req.(*ScheduleRemoteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchedulerService_ClearSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServiceServer).ClearSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.v1.SchedulerService/ClearSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServiceServer).ClearSchedule(ctx, req.(*ClearScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchedulerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.v1.SchedulerService",
	HandlerType: (*SchedulerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleRemoteTask",
			Handler:    _SchedulerService_ScheduleRemoteTask_Handler,
		},
		{
			MethodName: "ClearSchedule",
			Handler:    _SchedulerService_ClearSchedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/scheduler.proto",
}

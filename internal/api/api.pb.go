// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: internal/api/api.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PhpFpmServiceRequest_Action int32

const (
	PhpFpmServiceRequest_RESTART PhpFpmServiceRequest_Action = 0
	PhpFpmServiceRequest_STOP    PhpFpmServiceRequest_Action = 1
	PhpFpmServiceRequest_START   PhpFpmServiceRequest_Action = 2
)

// Enum value maps for PhpFpmServiceRequest_Action.
var (
	PhpFpmServiceRequest_Action_name = map[int32]string{
		0: "RESTART",
		1: "STOP",
		2: "START",
	}
	PhpFpmServiceRequest_Action_value = map[string]int32{
		"RESTART": 0,
		"STOP":    1,
		"START":   2,
	}
)

func (x PhpFpmServiceRequest_Action) Enum() *PhpFpmServiceRequest_Action {
	p := new(PhpFpmServiceRequest_Action)
	*p = x
	return p
}

func (x PhpFpmServiceRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhpFpmServiceRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_api_api_proto_enumTypes[0].Descriptor()
}

func (PhpFpmServiceRequest_Action) Type() protoreflect.EnumType {
	return &file_internal_api_api_proto_enumTypes[0]
}

func (x PhpFpmServiceRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhpFpmServiceRequest_Action.Descriptor instead.
func (PhpFpmServiceRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{0, 0}
}

type NginxServiceRequest_Action int32

const (
	NginxServiceRequest_RESTART NginxServiceRequest_Action = 0
	NginxServiceRequest_STOP    NginxServiceRequest_Action = 1
	NginxServiceRequest_START   NginxServiceRequest_Action = 2
)

// Enum value maps for NginxServiceRequest_Action.
var (
	NginxServiceRequest_Action_name = map[int32]string{
		0: "RESTART",
		1: "STOP",
		2: "START",
	}
	NginxServiceRequest_Action_value = map[string]int32{
		"RESTART": 0,
		"STOP":    1,
		"START":   2,
	}
)

func (x NginxServiceRequest_Action) Enum() *NginxServiceRequest_Action {
	p := new(NginxServiceRequest_Action)
	*p = x
	return p
}

func (x NginxServiceRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NginxServiceRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_api_api_proto_enumTypes[1].Descriptor()
}

func (NginxServiceRequest_Action) Type() protoreflect.EnumType {
	return &file_internal_api_api_proto_enumTypes[1]
}

func (x NginxServiceRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NginxServiceRequest_Action.Descriptor instead.
func (NginxServiceRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{1, 0}
}

type PhpFpmServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string                      `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Action  PhpFpmServiceRequest_Action `protobuf:"varint,2,opt,name=action,proto3,enum=api.PhpFpmServiceRequest_Action" json:"action,omitempty"`
}

func (x *PhpFpmServiceRequest) Reset() {
	*x = PhpFpmServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhpFpmServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhpFpmServiceRequest) ProtoMessage() {}

func (x *PhpFpmServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhpFpmServiceRequest.ProtoReflect.Descriptor instead.
func (*PhpFpmServiceRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *PhpFpmServiceRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PhpFpmServiceRequest) GetAction() PhpFpmServiceRequest_Action {
	if x != nil {
		return x.Action
	}
	return PhpFpmServiceRequest_RESTART
}

type NginxServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action NginxServiceRequest_Action `protobuf:"varint,1,opt,name=action,proto3,enum=api.NginxServiceRequest_Action" json:"action,omitempty"`
}

func (x *NginxServiceRequest) Reset() {
	*x = NginxServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NginxServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NginxServiceRequest) ProtoMessage() {}

func (x *NginxServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NginxServiceRequest.ProtoReflect.Descriptor instead.
func (*NginxServiceRequest) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *NginxServiceRequest) GetAction() NginxServiceRequest_Action {
	if x != nil {
		return x.Action
	}
	return NginxServiceRequest_RESTART
}

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_internal_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_api_api_proto protoreflect.FileDescriptor

var file_internal_api_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x96, 0x01,
	0x0a, 0x14, 0x50, 0x68, 0x70, 0x46, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x68, 0x70, 0x46, 0x70, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53,
	0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x22, 0x7a, 0x0a, 0x13, 0x4e, 0x67, 0x69, 0x6e, 0x78, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x4e, 0x67, 0x69, 0x6e, 0x78, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x02, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x94, 0x01, 0x0a, 0x0c, 0x4e, 0x69, 0x74, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x42, 0x0a, 0x0d, 0x50, 0x68, 0x70, 0x46, 0x70, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x68, 0x70, 0x46, 0x70, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x4e, 0x67, 0x69, 0x6e, 0x78, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x67, 0x69, 0x6e, 0x78,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_api_api_proto_rawDescOnce sync.Once
	file_internal_api_api_proto_rawDescData = file_internal_api_api_proto_rawDesc
)

func file_internal_api_api_proto_rawDescGZIP() []byte {
	file_internal_api_api_proto_rawDescOnce.Do(func() {
		file_internal_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_api_api_proto_rawDescData)
	})
	return file_internal_api_api_proto_rawDescData
}

var file_internal_api_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_internal_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_api_api_proto_goTypes = []interface{}{
	(PhpFpmServiceRequest_Action)(0), // 0: api.PhpFpmServiceRequest.Action
	(NginxServiceRequest_Action)(0),  // 1: api.NginxServiceRequest.Action
	(*PhpFpmServiceRequest)(nil),     // 2: api.PhpFpmServiceRequest
	(*NginxServiceRequest)(nil),      // 3: api.NginxServiceRequest
	(*ServiceResponse)(nil),          // 4: api.ServiceResponse
}
var file_internal_api_api_proto_depIdxs = []int32{
	0, // 0: api.PhpFpmServiceRequest.action:type_name -> api.PhpFpmServiceRequest.Action
	1, // 1: api.NginxServiceRequest.action:type_name -> api.NginxServiceRequest.Action
	2, // 2: api.NitroService.PhpFpmService:input_type -> api.PhpFpmServiceRequest
	3, // 3: api.NitroService.NginxService:input_type -> api.NginxServiceRequest
	4, // 4: api.NitroService.PhpFpmService:output_type -> api.ServiceResponse
	4, // 5: api.NitroService.NginxService:output_type -> api.ServiceResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_api_api_proto_init() }
func file_internal_api_api_proto_init() {
	if File_internal_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhpFpmServiceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NginxServiceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_internal_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_api_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_api_api_proto_goTypes,
		DependencyIndexes: file_internal_api_api_proto_depIdxs,
		EnumInfos:         file_internal_api_api_proto_enumTypes,
		MessageInfos:      file_internal_api_api_proto_msgTypes,
	}.Build()
	File_internal_api_api_proto = out.File
	file_internal_api_api_proto_rawDesc = nil
	file_internal_api_api_proto_goTypes = nil
	file_internal_api_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NitroServiceClient is the client API for NitroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NitroServiceClient interface {
	PhpFpmService(ctx context.Context, in *PhpFpmServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error)
	NginxService(ctx context.Context, in *NginxServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error)
}

type nitroServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNitroServiceClient(cc grpc.ClientConnInterface) NitroServiceClient {
	return &nitroServiceClient{cc}
}

func (c *nitroServiceClient) PhpFpmService(ctx context.Context, in *PhpFpmServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/api.NitroService/PhpFpmService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nitroServiceClient) NginxService(ctx context.Context, in *NginxServiceRequest, opts ...grpc.CallOption) (*ServiceResponse, error) {
	out := new(ServiceResponse)
	err := c.cc.Invoke(ctx, "/api.NitroService/NginxService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NitroServiceServer is the server API for NitroService service.
type NitroServiceServer interface {
	PhpFpmService(context.Context, *PhpFpmServiceRequest) (*ServiceResponse, error)
	NginxService(context.Context, *NginxServiceRequest) (*ServiceResponse, error)
}

// UnimplementedNitroServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNitroServiceServer struct {
}

func (*UnimplementedNitroServiceServer) PhpFpmService(context.Context, *PhpFpmServiceRequest) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhpFpmService not implemented")
}
func (*UnimplementedNitroServiceServer) NginxService(context.Context, *NginxServiceRequest) (*ServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NginxService not implemented")
}

func RegisterNitroServiceServer(s *grpc.Server, srv NitroServiceServer) {
	s.RegisterService(&_NitroService_serviceDesc, srv)
}

func _NitroService_PhpFpmService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhpFpmServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NitroServiceServer).PhpFpmService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NitroService/PhpFpmService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NitroServiceServer).PhpFpmService(ctx, req.(*PhpFpmServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NitroService_NginxService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NginxServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NitroServiceServer).NginxService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.NitroService/NginxService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NitroServiceServer).NginxService(ctx, req.(*NginxServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NitroService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.NitroService",
	HandlerType: (*NitroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PhpFpmService",
			Handler:    _NitroService_PhpFpmService_Handler,
		},
		{
			MethodName: "NginxService",
			Handler:    _NitroService_NginxService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/api.proto",
}

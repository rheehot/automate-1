// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: interservice/keyval/keyval.proto

package keyval

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

type PutKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" toml:"key,omitempty" mapstructure:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" toml:"value,omitempty" mapstructure:"value,omitempty"`
}

func (x *PutKeyRequest) Reset() {
	*x = PutKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_keyval_keyval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutKeyRequest) ProtoMessage() {}

func (x *PutKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_keyval_keyval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutKeyRequest.ProtoReflect.Descriptor instead.
func (*PutKeyRequest) Descriptor() ([]byte, []int) {
	return file_interservice_keyval_keyval_proto_rawDescGZIP(), []int{0}
}

func (x *PutKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutKeyRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type PutKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutKeyResponse) Reset() {
	*x = PutKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_keyval_keyval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutKeyResponse) ProtoMessage() {}

func (x *PutKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_keyval_keyval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutKeyResponse.ProtoReflect.Descriptor instead.
func (*PutKeyResponse) Descriptor() ([]byte, []int) {
	return file_interservice_keyval_keyval_proto_rawDescGZIP(), []int{1}
}

var File_interservice_keyval_keyval_proto protoreflect.FileDescriptor

var file_interservice_keyval_keyval_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6b,
	0x65, 0x79, 0x76, 0x61, 0x6c, 0x2f, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x22,
	0x37, 0x0a, 0x0d, 0x50, 0x75, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x50, 0x75, 0x74, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x72, 0x0a, 0x0d, 0x4b, 0x65,
	0x79, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x50,
	0x75, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6b, 0x65, 0x79,
	0x76, 0x61, 0x6c, 0x2e, 0x50, 0x75, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x6c, 0x2e,
	0x50, 0x75, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65,
	0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6b, 0x65, 0x79, 0x76,
	0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_keyval_keyval_proto_rawDescOnce sync.Once
	file_interservice_keyval_keyval_proto_rawDescData = file_interservice_keyval_keyval_proto_rawDesc
)

func file_interservice_keyval_keyval_proto_rawDescGZIP() []byte {
	file_interservice_keyval_keyval_proto_rawDescOnce.Do(func() {
		file_interservice_keyval_keyval_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_keyval_keyval_proto_rawDescData)
	})
	return file_interservice_keyval_keyval_proto_rawDescData
}

var file_interservice_keyval_keyval_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_interservice_keyval_keyval_proto_goTypes = []interface{}{
	(*PutKeyRequest)(nil),  // 0: chef.automate.domain.keyval.PutKeyRequest
	(*PutKeyResponse)(nil), // 1: chef.automate.domain.keyval.PutKeyResponse
}
var file_interservice_keyval_keyval_proto_depIdxs = []int32{
	0, // 0: chef.automate.domain.keyval.KeyvalService.PutKey:input_type -> chef.automate.domain.keyval.PutKeyRequest
	1, // 1: chef.automate.domain.keyval.KeyvalService.PutKey:output_type -> chef.automate.domain.keyval.PutKeyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_keyval_keyval_proto_init() }
func file_interservice_keyval_keyval_proto_init() {
	if File_interservice_keyval_keyval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_keyval_keyval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutKeyRequest); i {
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
		file_interservice_keyval_keyval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutKeyResponse); i {
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
			RawDescriptor: file_interservice_keyval_keyval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_interservice_keyval_keyval_proto_goTypes,
		DependencyIndexes: file_interservice_keyval_keyval_proto_depIdxs,
		MessageInfos:      file_interservice_keyval_keyval_proto_msgTypes,
	}.Build()
	File_interservice_keyval_keyval_proto = out.File
	file_interservice_keyval_keyval_proto_rawDesc = nil
	file_interservice_keyval_keyval_proto_goTypes = nil
	file_interservice_keyval_keyval_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeyvalServiceClient is the client API for KeyvalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyvalServiceClient interface {
	PutKey(ctx context.Context, in *PutKeyRequest, opts ...grpc.CallOption) (*PutKeyResponse, error)
}

type keyvalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyvalServiceClient(cc grpc.ClientConnInterface) KeyvalServiceClient {
	return &keyvalServiceClient{cc}
}

func (c *keyvalServiceClient) PutKey(ctx context.Context, in *PutKeyRequest, opts ...grpc.CallOption) (*PutKeyResponse, error) {
	out := new(PutKeyResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.domain.keyval.KeyvalService/PutKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyvalServiceServer is the server API for KeyvalService service.
type KeyvalServiceServer interface {
	PutKey(context.Context, *PutKeyRequest) (*PutKeyResponse, error)
}

// UnimplementedKeyvalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeyvalServiceServer struct {
}

func (*UnimplementedKeyvalServiceServer) PutKey(context.Context, *PutKeyRequest) (*PutKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutKey not implemented")
}

func RegisterKeyvalServiceServer(s *grpc.Server, srv KeyvalServiceServer) {
	s.RegisterService(&_KeyvalService_serviceDesc, srv)
}

func _KeyvalService_PutKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyvalServiceServer).PutKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.domain.keyval.KeyvalService/PutKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyvalServiceServer).PutKey(ctx, req.(*PutKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyvalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.domain.keyval.KeyvalService",
	HandlerType: (*KeyvalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutKey",
			Handler:    _KeyvalService_PutKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interservice/keyval/keyval.proto",
}
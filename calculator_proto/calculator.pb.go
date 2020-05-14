// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: calculator.proto

package calculator_proto

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Message structure
type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculateRequest) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *CalculateRequest) GetB() float64 {
	if x != nil {
		return x.B
	}
	return 0
}

type CalculateResponnse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A      float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B      float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
	Result float64 `protobuf:"fixed64,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculateResponnse) Reset() {
	*x = CalculateResponnse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponnse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponnse) ProtoMessage() {}

func (x *CalculateResponnse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponnse.ProtoReflect.Descriptor instead.
func (*CalculateResponnse) Descriptor() ([]byte, []int) {
	return file_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateResponnse) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *CalculateResponnse) GetB() float64 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *CalculateResponnse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_proto protoreflect.FileDescriptor

var file_calculator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x01, 0x62, 0x22, 0x48, 0x0a, 0x12, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x6e, 0x73, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xcf,
	0x02, 0x0a, 0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x4f, 0x0a, 0x03,
	0x53, 0x75, 0x6d, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a,
	0x03, 0x53, 0x75, 0x62, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x03, 0x4d, 0x75, 0x6c, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x03, 0x44, 0x69, 0x76, 0x12, 0x22, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x6e, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_proto_rawDescOnce sync.Once
	file_calculator_proto_rawDescData = file_calculator_proto_rawDesc
)

func file_calculator_proto_rawDescGZIP() []byte {
	file_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_proto_rawDescData)
	})
	return file_calculator_proto_rawDescData
}

var file_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_proto_goTypes = []interface{}{
	(*CalculateRequest)(nil),   // 0: calculator_proto.CalculateRequest
	(*CalculateResponnse)(nil), // 1: calculator_proto.CalculateResponnse
}
var file_calculator_proto_depIdxs = []int32{
	0, // 0: calculator_proto.Calculate.Sum:input_type -> calculator_proto.CalculateRequest
	0, // 1: calculator_proto.Calculate.Sub:input_type -> calculator_proto.CalculateRequest
	0, // 2: calculator_proto.Calculate.Mul:input_type -> calculator_proto.CalculateRequest
	0, // 3: calculator_proto.Calculate.Div:input_type -> calculator_proto.CalculateRequest
	1, // 4: calculator_proto.Calculate.Sum:output_type -> calculator_proto.CalculateResponnse
	1, // 5: calculator_proto.Calculate.Sub:output_type -> calculator_proto.CalculateResponnse
	1, // 6: calculator_proto.Calculate.Mul:output_type -> calculator_proto.CalculateResponnse
	1, // 7: calculator_proto.Calculate.Div:output_type -> calculator_proto.CalculateResponnse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_proto_init() }
func file_calculator_proto_init() {
	if File_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRequest); i {
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
		file_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponnse); i {
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
			RawDescriptor: file_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_proto_msgTypes,
	}.Build()
	File_calculator_proto = out.File
	file_calculator_proto_rawDesc = nil
	file_calculator_proto_goTypes = nil
	file_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculateClient is the client API for Calculate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateClient interface {
	Sum(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error)
	Sub(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error)
	Mul(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error)
	Div(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error)
}

type calculateClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateClient(cc grpc.ClientConnInterface) CalculateClient {
	return &calculateClient{cc}
}

func (c *calculateClient) Sum(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error) {
	out := new(CalculateResponnse)
	err := c.cc.Invoke(ctx, "/calculator_proto.Calculate/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateClient) Sub(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error) {
	out := new(CalculateResponnse)
	err := c.cc.Invoke(ctx, "/calculator_proto.Calculate/Sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateClient) Mul(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error) {
	out := new(CalculateResponnse)
	err := c.cc.Invoke(ctx, "/calculator_proto.Calculate/Mul", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateClient) Div(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponnse, error) {
	out := new(CalculateResponnse)
	err := c.cc.Invoke(ctx, "/calculator_proto.Calculate/Div", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServer is the server API for Calculate service.
type CalculateServer interface {
	Sum(context.Context, *CalculateRequest) (*CalculateResponnse, error)
	Sub(context.Context, *CalculateRequest) (*CalculateResponnse, error)
	Mul(context.Context, *CalculateRequest) (*CalculateResponnse, error)
	Div(context.Context, *CalculateRequest) (*CalculateResponnse, error)
}

// UnimplementedCalculateServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServer struct {
}

func (*UnimplementedCalculateServer) Sum(context.Context, *CalculateRequest) (*CalculateResponnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculateServer) Sub(context.Context, *CalculateRequest) (*CalculateResponnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (*UnimplementedCalculateServer) Mul(context.Context, *CalculateRequest) (*CalculateResponnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mul not implemented")
}
func (*UnimplementedCalculateServer) Div(context.Context, *CalculateRequest) (*CalculateResponnse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Div not implemented")
}

func RegisterCalculateServer(s *grpc.Server, srv CalculateServer) {
	s.RegisterService(&_Calculate_serviceDesc, srv)
}

func _Calculate_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator_proto.Calculate/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Sum(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculate_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator_proto.Calculate/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Sub(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculate_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator_proto.Calculate/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Mul(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculate_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator_proto.Calculate/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServer).Div(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator_proto.Calculate",
	HandlerType: (*CalculateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Calculate_Sum_Handler,
		},
		{
			MethodName: "Sub",
			Handler:    _Calculate_Sub_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _Calculate_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _Calculate_Div_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator.proto",
}

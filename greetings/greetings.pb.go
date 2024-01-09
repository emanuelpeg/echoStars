// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: greetings/greetings.proto

package greetings

import (
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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetings_greetings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetings_greetings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_greetings_greetings_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetings_greetings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetings_greetings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_greetings_greetings_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HelloWoldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HelloWoldRequest) Reset() {
	*x = HelloWoldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetings_greetings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWoldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWoldRequest) ProtoMessage() {}

func (x *HelloWoldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greetings_greetings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWoldRequest.ProtoReflect.Descriptor instead.
func (*HelloWoldRequest) Descriptor() ([]byte, []int) {
	return file_greetings_greetings_proto_rawDescGZIP(), []int{2}
}

type HelloWoldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloWoldResponse) Reset() {
	*x = HelloWoldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greetings_greetings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloWoldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloWoldResponse) ProtoMessage() {}

func (x *HelloWoldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greetings_greetings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloWoldResponse.ProtoReflect.Descriptor instead.
func (*HelloWoldResponse) Descriptor() ([]byte, []int) {
	return file_greetings_greetings_proto_rawDescGZIP(), []int{3}
}

func (x *HelloWoldResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_greetings_greetings_proto protoreflect.FileDescriptor

var file_greetings_greetings_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x29, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2d, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9e,
	0x01, 0x0a, 0x12, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x72, 0x70, 0x63, 0x12, 0x3c, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x17,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x09, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57, 0x6f, 0x6c, 0x64,
	0x12, 0x1b, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x57,
	0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x16, 0x5a, 0x14, 0x65, 0x63, 0x68, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x73, 0x2f, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greetings_greetings_proto_rawDescOnce sync.Once
	file_greetings_greetings_proto_rawDescData = file_greetings_greetings_proto_rawDesc
)

func file_greetings_greetings_proto_rawDescGZIP() []byte {
	file_greetings_greetings_proto_rawDescOnce.Do(func() {
		file_greetings_greetings_proto_rawDescData = protoimpl.X.CompressGZIP(file_greetings_greetings_proto_rawDescData)
	})
	return file_greetings_greetings_proto_rawDescData
}

var file_greetings_greetings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greetings_greetings_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),      // 0: greetings.HelloRequest
	(*HelloResponse)(nil),     // 1: greetings.HelloResponse
	(*HelloWoldRequest)(nil),  // 2: greetings.HelloWoldRequest
	(*HelloWoldResponse)(nil), // 3: greetings.HelloWoldResponse
}
var file_greetings_greetings_proto_depIdxs = []int32{
	0, // 0: greetings.GreeterServiceGrpc.Hello:input_type -> greetings.HelloRequest
	2, // 1: greetings.GreeterServiceGrpc.HelloWold:input_type -> greetings.HelloWoldRequest
	1, // 2: greetings.GreeterServiceGrpc.Hello:output_type -> greetings.HelloResponse
	3, // 3: greetings.GreeterServiceGrpc.HelloWold:output_type -> greetings.HelloWoldResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greetings_greetings_proto_init() }
func file_greetings_greetings_proto_init() {
	if File_greetings_greetings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greetings_greetings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_greetings_greetings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_greetings_greetings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWoldRequest); i {
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
		file_greetings_greetings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloWoldResponse); i {
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
			RawDescriptor: file_greetings_greetings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greetings_greetings_proto_goTypes,
		DependencyIndexes: file_greetings_greetings_proto_depIdxs,
		MessageInfos:      file_greetings_greetings_proto_msgTypes,
	}.Build()
	File_greetings_greetings_proto = out.File
	file_greetings_greetings_proto_rawDesc = nil
	file_greetings_greetings_proto_goTypes = nil
	file_greetings_greetings_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	ragù          v0.2.3
// source: pkg/test/testdata/plugins/ext/ext.proto

package ext

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request string `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *FooRequest) Reset() {
	*x = FooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooRequest) ProtoMessage() {}

func (x *FooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooRequest.ProtoReflect.Descriptor instead.
func (*FooRequest) Descriptor() ([]byte, []int) {
	return file_pkg_test_testdata_plugins_ext_ext_proto_rawDescGZIP(), []int{0}
}

func (x *FooRequest) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

type FooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *FooResponse) Reset() {
	*x = FooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooResponse) ProtoMessage() {}

func (x *FooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooResponse.ProtoReflect.Descriptor instead.
func (*FooResponse) Descriptor() ([]byte, []int) {
	return file_pkg_test_testdata_plugins_ext_ext_proto_rawDescGZIP(), []int{1}
}

func (x *FooResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_pkg_test_testdata_plugins_ext_ext_proto protoreflect.FileDescriptor

var file_pkg_test_testdata_plugins_ext_ext_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x2f,
	0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x78, 0x74, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0a, 0x46,
	0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x11, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x23,
	0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x00, 0x3a, 0x00, 0x32, 0xdb, 0x02, 0x0a, 0x03, 0x45, 0x78, 0x74, 0x12, 0xd1, 0x02, 0x0a, 0x03,
	0x46, 0x6f, 0x6f, 0x12, 0x0f, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa2, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x3a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5a, 0x06, 0x12, 0x04, 0x2f, 0x66, 0x6f, 0x6f,
	0x5a, 0x0f, 0x3a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x04, 0x2f, 0x66, 0x6f,
	0x6f, 0x5a, 0x06, 0x2a, 0x04, 0x2f, 0x66, 0x6f, 0x6f, 0x5a, 0x0f, 0x3a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x32, 0x04, 0x2f, 0x66, 0x6f, 0x6f, 0x22, 0x04, 0x2f, 0x66, 0x6f, 0x6f,
	0xba, 0x3e, 0xd7, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0xbf, 0x01, 0x7b, 0x70, 0x6f, 0x73,
	0x74, 0x3a, 0x22, 0x2f, 0x66, 0x6f, 0x6f, 0x22, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x22, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x7b, 0x67, 0x65, 0x74, 0x3a,
	0x22, 0x2f, 0x66, 0x6f, 0x6f, 0x22, 0x7d, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x7b, 0x70, 0x75, 0x74, 0x3a,
	0x22, 0x2f, 0x66, 0x6f, 0x6f, 0x22, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x22, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x7d, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x3a, 0x22, 0x2f, 0x66, 0x6f, 0x6f, 0x22, 0x7d, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x7b, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x3a, 0x22, 0x2f, 0x66, 0x6f, 0x6f, 0x22, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x3a,
	0x22, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7d, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x1a,
	0x00, 0x32, 0x38, 0x0a, 0x04, 0x45, 0x78, 0x74, 0x32, 0x12, 0x2e, 0x0a, 0x03, 0x46, 0x6f, 0x6f,
	0x12, 0x0f, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64,
	0x61, 0x74, 0x61, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_test_testdata_plugins_ext_ext_proto_rawDescOnce sync.Once
	file_pkg_test_testdata_plugins_ext_ext_proto_rawDescData = file_pkg_test_testdata_plugins_ext_ext_proto_rawDesc
)

func file_pkg_test_testdata_plugins_ext_ext_proto_rawDescGZIP() []byte {
	file_pkg_test_testdata_plugins_ext_ext_proto_rawDescOnce.Do(func() {
		file_pkg_test_testdata_plugins_ext_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_test_testdata_plugins_ext_ext_proto_rawDescData)
	})
	return file_pkg_test_testdata_plugins_ext_ext_proto_rawDescData
}

var file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_test_testdata_plugins_ext_ext_proto_goTypes = []interface{}{
	(*FooRequest)(nil),  // 0: ext.FooRequest
	(*FooResponse)(nil), // 1: ext.FooResponse
}
var file_pkg_test_testdata_plugins_ext_ext_proto_depIdxs = []int32{
	0, // 0: ext.Ext.Foo:input_type -> ext.FooRequest
	0, // 1: ext.Ext2.Foo:input_type -> ext.FooRequest
	1, // 2: ext.Ext.Foo:output_type -> ext.FooResponse
	1, // 3: ext.Ext2.Foo:output_type -> ext.FooResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_test_testdata_plugins_ext_ext_proto_init() }
func file_pkg_test_testdata_plugins_ext_ext_proto_init() {
	if File_pkg_test_testdata_plugins_ext_ext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooRequest); i {
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
		file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooResponse); i {
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
			RawDescriptor: file_pkg_test_testdata_plugins_ext_ext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_test_testdata_plugins_ext_ext_proto_goTypes,
		DependencyIndexes: file_pkg_test_testdata_plugins_ext_ext_proto_depIdxs,
		MessageInfos:      file_pkg_test_testdata_plugins_ext_ext_proto_msgTypes,
	}.Build()
	File_pkg_test_testdata_plugins_ext_ext_proto = out.File
	file_pkg_test_testdata_plugins_ext_ext_proto_rawDesc = nil
	file_pkg_test_testdata_plugins_ext_ext_proto_goTypes = nil
	file_pkg_test_testdata_plugins_ext_ext_proto_depIdxs = nil
}

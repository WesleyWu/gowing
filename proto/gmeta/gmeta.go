// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: gmeta.proto

package gmeta

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

type Meta int32

const (
	Meta__ Meta = 0
)

// Enum value maps for Meta.
var (
	Meta_name = map[int32]string{
		0: "_",
	}
	Meta_value = map[string]int32{
		"_": 0,
	}
)

func (x Meta) Enum() *Meta {
	p := new(Meta)
	*p = x
	return p
}

func (x Meta) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Meta) Descriptor() protoreflect.EnumDescriptor {
	return file_goframe_proto_enumTypes[0].Descriptor()
}

func (Meta) Type() protoreflect.EnumType {
	return &file_goframe_proto_enumTypes[0]
}

func (x Meta) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Meta.Descriptor instead.
func (Meta) EnumDescriptor() ([]byte, []int) {
	return file_goframe_proto_rawDescGZIP(), []int{0}
}

var File_goframe_proto protoreflect.FileDescriptor

var file_goframe_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x6f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2a, 0x0d, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x05, 0x0a, 0x01, 0x5f, 0x10, 0x00, 0x42,
	0x0e, 0x5a, 0x0c, 0x67, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x6d, 0x65, 0x74, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goframe_proto_rawDescOnce sync.Once
	file_goframe_proto_rawDescData = file_goframe_proto_rawDesc
)

func file_goframe_proto_rawDescGZIP() []byte {
	file_goframe_proto_rawDescOnce.Do(func() {
		file_goframe_proto_rawDescData = protoimpl.X.CompressGZIP(file_goframe_proto_rawDescData)
	})
	return file_goframe_proto_rawDescData
}

var file_goframe_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_goframe_proto_goTypes = []interface{}{
	(Meta)(0), // 0: gowing.protobuf.Meta
}
var file_goframe_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goframe_proto_init() }
func file_goframe_proto_init() {
	if File_goframe_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goframe_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_goframe_proto_goTypes,
		DependencyIndexes: file_goframe_proto_depIdxs,
		EnumInfos:         file_goframe_proto_enumTypes,
	}.Build()
	File_goframe_proto = out.File
	file_goframe_proto_rawDesc = nil
	file_goframe_proto_goTypes = nil
	file_goframe_proto_depIdxs = nil
}

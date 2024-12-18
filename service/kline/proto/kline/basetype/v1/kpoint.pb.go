// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: kline/basetype/v1/kpoint.proto

package v1

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

type KPointType int32

const (
	KPointType_KPointTypeUnknown KPointType = 0
	KPointType_FiveSecond        KPointType = 10
	KPointType_OneMinute         KPointType = 20
	KPointType_TenMinute         KPointType = 30
	KPointType_OneHour           KPointType = 40
	KPointType_OneDay            KPointType = 50
	KPointType_OneWeek           KPointType = 60
	KPointType_OneMonth          KPointType = 70
)

// Enum value maps for KPointType.
var (
	KPointType_name = map[int32]string{
		0:  "KPointTypeUnknown",
		10: "FiveSecond",
		20: "OneMinute",
		30: "TenMinute",
		40: "OneHour",
		50: "OneDay",
		60: "OneWeek",
		70: "OneMonth",
	}
	KPointType_value = map[string]int32{
		"KPointTypeUnknown": 0,
		"FiveSecond":        10,
		"OneMinute":         20,
		"TenMinute":         30,
		"OneHour":           40,
		"OneDay":            50,
		"OneWeek":           60,
		"OneMonth":          70,
	}
)

func (x KPointType) Enum() *KPointType {
	p := new(KPointType)
	*p = x
	return p
}

func (x KPointType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KPointType) Descriptor() protoreflect.EnumDescriptor {
	return file_kline_basetype_v1_kpoint_proto_enumTypes[0].Descriptor()
}

func (KPointType) Type() protoreflect.EnumType {
	return &file_kline_basetype_v1_kpoint_proto_enumTypes[0]
}

func (x KPointType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KPointType.Descriptor instead.
func (KPointType) EnumDescriptor() ([]byte, []int) {
	return file_kline_basetype_v1_kpoint_proto_rawDescGZIP(), []int{0}
}

var File_kline_basetype_v1_kpoint_proto protoreflect.FileDescriptor

var file_kline_basetype_v1_kpoint_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2a, 0x85, 0x01, 0x0a, 0x0a, 0x4b,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x4b, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x69, 0x76, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x0a,
	0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x6e, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x10, 0x14, 0x12,
	0x0d, 0x0a, 0x09, 0x54, 0x65, 0x6e, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x10, 0x1e, 0x12, 0x0b,
	0x0a, 0x07, 0x4f, 0x6e, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x10, 0x28, 0x12, 0x0a, 0x0a, 0x06, 0x4f,
	0x6e, 0x65, 0x44, 0x61, 0x79, 0x10, 0x32, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x6e, 0x65, 0x57, 0x65,
	0x65, 0x6b, 0x10, 0x3c, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x6e, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x10, 0x46, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x32, 0x35, 0x35, 0x31, 0x39, 0x2f, 0x6c, 0x69, 0x6e,
	0x65, 0x72, 0x61, 0x2d, 0x64, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x6b, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x6c,
	0x69, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kline_basetype_v1_kpoint_proto_rawDescOnce sync.Once
	file_kline_basetype_v1_kpoint_proto_rawDescData = file_kline_basetype_v1_kpoint_proto_rawDesc
)

func file_kline_basetype_v1_kpoint_proto_rawDescGZIP() []byte {
	file_kline_basetype_v1_kpoint_proto_rawDescOnce.Do(func() {
		file_kline_basetype_v1_kpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_kline_basetype_v1_kpoint_proto_rawDescData)
	})
	return file_kline_basetype_v1_kpoint_proto_rawDescData
}

var file_kline_basetype_v1_kpoint_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kline_basetype_v1_kpoint_proto_goTypes = []interface{}{
	(KPointType)(0), // 0: basetype.KPointType
}
var file_kline_basetype_v1_kpoint_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_kline_basetype_v1_kpoint_proto_init() }
func file_kline_basetype_v1_kpoint_proto_init() {
	if File_kline_basetype_v1_kpoint_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kline_basetype_v1_kpoint_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kline_basetype_v1_kpoint_proto_goTypes,
		DependencyIndexes: file_kline_basetype_v1_kpoint_proto_depIdxs,
		EnumInfos:         file_kline_basetype_v1_kpoint_proto_enumTypes,
	}.Build()
	File_kline_basetype_v1_kpoint_proto = out.File
	file_kline_basetype_v1_kpoint_proto_rawDesc = nil
	file_kline_basetype_v1_kpoint_proto_goTypes = nil
	file_kline_basetype_v1_kpoint_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/quality/quality.proto

package quality

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

type Quality int32

const (
	Quality_UNSPECIFIED Quality = 0
	Quality_Q360P       Quality = 360
	Quality_Q480P       Quality = 480
	Quality_Q720P       Quality = 720
	Quality_Q1080P      Quality = 1080
)

// Enum value maps for Quality.
var (
	Quality_name = map[int32]string{
		0:    "UNSPECIFIED",
		360:  "Q360P",
		480:  "Q480P",
		720:  "Q720P",
		1080: "Q1080P",
	}
	Quality_value = map[string]int32{
		"UNSPECIFIED": 0,
		"Q360P":       360,
		"Q480P":       480,
		"Q720P":       720,
		"Q1080P":      1080,
	}
)

func (x Quality) Enum() *Quality {
	p := new(Quality)
	*p = x
	return p
}

func (x Quality) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Quality) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_quality_quality_proto_enumTypes[0].Descriptor()
}

func (Quality) Type() protoreflect.EnumType {
	return &file_proto_quality_quality_proto_enumTypes[0]
}

func (x Quality) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Quality.Descriptor instead.
func (Quality) EnumDescriptor() ([]byte, []int) {
	return file_proto_quality_quality_proto_rawDescGZIP(), []int{0}
}

var File_proto_quality_quality_proto protoreflect.FileDescriptor

var file_proto_quality_quality_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2f,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x2a, 0x4b, 0x0a, 0x07, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x05, 0x51, 0x33, 0x36, 0x30, 0x50, 0x10, 0xe8, 0x02, 0x12, 0x0a,
	0x0a, 0x05, 0x51, 0x34, 0x38, 0x30, 0x50, 0x10, 0xe0, 0x03, 0x12, 0x0a, 0x0a, 0x05, 0x51, 0x37,
	0x32, 0x30, 0x50, 0x10, 0xd0, 0x05, 0x12, 0x0b, 0x0a, 0x06, 0x51, 0x31, 0x30, 0x38, 0x30, 0x50,
	0x10, 0xb8, 0x08, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x71, 0x75, 0x6f, 0x63, 0x62, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x2d, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quality_quality_proto_rawDescOnce sync.Once
	file_proto_quality_quality_proto_rawDescData = file_proto_quality_quality_proto_rawDesc
)

func file_proto_quality_quality_proto_rawDescGZIP() []byte {
	file_proto_quality_quality_proto_rawDescOnce.Do(func() {
		file_proto_quality_quality_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quality_quality_proto_rawDescData)
	})
	return file_proto_quality_quality_proto_rawDescData
}

var file_proto_quality_quality_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_quality_quality_proto_goTypes = []interface{}{
	(Quality)(0), // 0: quality.Quality
}
var file_proto_quality_quality_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_quality_quality_proto_init() }
func file_proto_quality_quality_proto_init() {
	if File_proto_quality_quality_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_quality_quality_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_quality_quality_proto_goTypes,
		DependencyIndexes: file_proto_quality_quality_proto_depIdxs,
		EnumInfos:         file_proto_quality_quality_proto_enumTypes,
	}.Build()
	File_proto_quality_quality_proto = out.File
	file_proto_quality_quality_proto_rawDesc = nil
	file_proto_quality_quality_proto_goTypes = nil
	file_proto_quality_quality_proto_depIdxs = nil
}

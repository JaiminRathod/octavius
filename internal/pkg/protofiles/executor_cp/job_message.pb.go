// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: internal/pkg/protofiles/executor_cp/job_message.proto

package executor_cp

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_internal_pkg_protofiles_executor_cp_job_message_proto protoreflect.FileDescriptor

var file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDesc = []byte{
	0x0a, 0x35, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x6f, 0x72, 0x5f, 0x63, 0x70, 0x2f, 0x6a, 0x6f, 0x62, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x2e, 0x5a, 0x2c, 0x6f, 0x63, 0x74, 0x61, 0x76,
	0x69, 0x75, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_internal_pkg_protofiles_executor_cp_job_message_proto_goTypes = []interface{}{}
var file_internal_pkg_protofiles_executor_cp_job_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pkg_protofiles_executor_cp_job_message_proto_init() }
func file_internal_pkg_protofiles_executor_cp_job_message_proto_init() {
	if File_internal_pkg_protofiles_executor_cp_job_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_protofiles_executor_cp_job_message_proto_goTypes,
		DependencyIndexes: file_internal_pkg_protofiles_executor_cp_job_message_proto_depIdxs,
	}.Build()
	File_internal_pkg_protofiles_executor_cp_job_message_proto = out.File
	file_internal_pkg_protofiles_executor_cp_job_message_proto_rawDesc = nil
	file_internal_pkg_protofiles_executor_cp_job_message_proto_goTypes = nil
	file_internal_pkg_protofiles_executor_cp_job_message_proto_depIdxs = nil
}

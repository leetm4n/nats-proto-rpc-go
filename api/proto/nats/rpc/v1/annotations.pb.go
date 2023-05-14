// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: nats/rpc/v1/annotations.proto

package natsrpcv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_nats_rpc_v1_annotations_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         133337,
		Name:          "nats.rpc.v1.isFileGenerationTarget",
		Tag:           "varint,133337,opt,name=isFileGenerationTarget",
		Filename:      "nats/rpc/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         133338,
		Name:          "nats.rpc.v1.isServiceGenerationTarget",
		Tag:           "varint,133338,opt,name=isServiceGenerationTarget",
		Filename:      "nats/rpc/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         133339,
		Name:          "nats.rpc.v1.serviceVersion",
		Tag:           "bytes,133339,opt,name=serviceVersion",
		Filename:      "nats/rpc/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         133340,
		Name:          "nats.rpc.v1.serviceSubject",
		Tag:           "bytes,133340,opt,name=serviceSubject",
		Filename:      "nats/rpc/v1/annotations.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         133341,
		Name:          "nats.rpc.v1.methodSubject",
		Tag:           "bytes,133341,opt,name=methodSubject",
		Filename:      "nats/rpc/v1/annotations.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// Flags if services within proto file are targets of generation
	//
	// optional bool isFileGenerationTarget = 133337;
	E_IsFileGenerationTarget = &file_nats_rpc_v1_annotations_proto_extTypes[0]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// Flags if service is target of generation
	//
	// optional bool isServiceGenerationTarget = 133338;
	E_IsServiceGenerationTarget = &file_nats_rpc_v1_annotations_proto_extTypes[1]
	// Setting this will set the nats micro.service version
	//
	// optional string serviceVersion = 133339;
	E_ServiceVersion = &file_nats_rpc_v1_annotations_proto_extTypes[2]
	// Overrides default subject name of service
	//
	// optional string serviceSubject = 133340;
	E_ServiceSubject = &file_nats_rpc_v1_annotations_proto_extTypes[3]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// Overrides default subject name of method
	//
	// optional string methodSubject = 133341;
	E_MethodSubject = &file_nats_rpc_v1_annotations_proto_extTypes[4]
)

var File_nats_rpc_v1_annotations_proto protoreflect.FileDescriptor

var file_nats_rpc_v1_annotations_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x61, 0x74, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x6e, 0x61, 0x74, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x56,
	0x0a, 0x16, 0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x91, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16,
	0x69, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x3a, 0x5f, 0x0a, 0x19, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xda, 0x91, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x69, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x3a, 0x49, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdb, 0x91, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x3a, 0x49, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdc, 0x91, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x46, 0x0a,
	0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdd,
	0x91, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0xb7, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x61,
	0x74, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x65, 0x65, 0x74, 0x6d, 0x34,
	0x6e, 0x2f, 0x6e, 0x61, 0x74, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x72, 0x70, 0x63,
	0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x61,
	0x74, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x61, 0x74, 0x73, 0x72, 0x70,
	0x63, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x52, 0x58, 0xaa, 0x02, 0x0b, 0x4e, 0x61, 0x74, 0x73,
	0x2e, 0x52, 0x70, 0x63, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x4e, 0x61, 0x74, 0x73, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x4e, 0x61, 0x74, 0x73, 0x5c, 0x52, 0x70, 0x63,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0d, 0x4e, 0x61, 0x74, 0x73, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_nats_rpc_v1_annotations_proto_goTypes = []interface{}{
	(*descriptorpb.FileOptions)(nil),    // 0: google.protobuf.FileOptions
	(*descriptorpb.ServiceOptions)(nil), // 1: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),  // 2: google.protobuf.MethodOptions
}
var file_nats_rpc_v1_annotations_proto_depIdxs = []int32{
	0, // 0: nats.rpc.v1.isFileGenerationTarget:extendee -> google.protobuf.FileOptions
	1, // 1: nats.rpc.v1.isServiceGenerationTarget:extendee -> google.protobuf.ServiceOptions
	1, // 2: nats.rpc.v1.serviceVersion:extendee -> google.protobuf.ServiceOptions
	1, // 3: nats.rpc.v1.serviceSubject:extendee -> google.protobuf.ServiceOptions
	2, // 4: nats.rpc.v1.methodSubject:extendee -> google.protobuf.MethodOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	0, // [0:5] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nats_rpc_v1_annotations_proto_init() }
func file_nats_rpc_v1_annotations_proto_init() {
	if File_nats_rpc_v1_annotations_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nats_rpc_v1_annotations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 5,
			NumServices:   0,
		},
		GoTypes:           file_nats_rpc_v1_annotations_proto_goTypes,
		DependencyIndexes: file_nats_rpc_v1_annotations_proto_depIdxs,
		ExtensionInfos:    file_nats_rpc_v1_annotations_proto_extTypes,
	}.Build()
	File_nats_rpc_v1_annotations_proto = out.File
	file_nats_rpc_v1_annotations_proto_rawDesc = nil
	file_nats_rpc_v1_annotations_proto_goTypes = nil
	file_nats_rpc_v1_annotations_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: example/service/v1/service.proto

package exampleservicev1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/leetm4n/nats-proto-rpc-go/api/proto/nats/rpc/v1"
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

// SendMessageRequest is the request for SendMessage
type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message itself
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The recipient email of the message
	RecipientId string `protobuf:"bytes,2,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_example_service_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SendMessageRequest) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

// SendMessageResponse is the response of SendMessage
type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the message sent
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_example_service_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetMessageRequest is the request for GetMessage
type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message id itself
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_example_service_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessageRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// GetMessageResponse is the response of GetMessage
type GetMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id of the message
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The message itself
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_example_service_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_example_service_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_example_service_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessageResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetMessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_example_service_v1_service_proto protoreflect.FileDescriptor

var file_example_service_v1_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x6e, 0x61, 0x74, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64,
	0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xfb, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x25, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0xea, 0x8d, 0x41, 0x0c, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0xd0, 0x8d, 0x41, 0x01, 0xda, 0x8d,
	0x41, 0x05, 0x31, 0x2e, 0x30, 0x2e, 0x30, 0xe2, 0x8d, 0x41, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xed, 0x01, 0xc8, 0x8d, 0x41, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x65, 0x65, 0x74, 0x6d, 0x34, 0x6e, 0x2f, 0x6e, 0x61, 0x74, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x45, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x14, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_example_service_v1_service_proto_rawDescOnce sync.Once
	file_example_service_v1_service_proto_rawDescData = file_example_service_v1_service_proto_rawDesc
)

func file_example_service_v1_service_proto_rawDescGZIP() []byte {
	file_example_service_v1_service_proto_rawDescOnce.Do(func() {
		file_example_service_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_example_service_v1_service_proto_rawDescData)
	})
	return file_example_service_v1_service_proto_rawDescData
}

var file_example_service_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_example_service_v1_service_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),  // 0: example.service.v1.SendMessageRequest
	(*SendMessageResponse)(nil), // 1: example.service.v1.SendMessageResponse
	(*GetMessageRequest)(nil),   // 2: example.service.v1.GetMessageRequest
	(*GetMessageResponse)(nil),  // 3: example.service.v1.GetMessageResponse
}
var file_example_service_v1_service_proto_depIdxs = []int32{
	0, // 0: example.service.v1.TestService.SendMessage:input_type -> example.service.v1.SendMessageRequest
	2, // 1: example.service.v1.TestService.GetMessage:input_type -> example.service.v1.GetMessageRequest
	1, // 2: example.service.v1.TestService.SendMessage:output_type -> example.service.v1.SendMessageResponse
	3, // 3: example.service.v1.TestService.GetMessage:output_type -> example.service.v1.GetMessageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_example_service_v1_service_proto_init() }
func file_example_service_v1_service_proto_init() {
	if File_example_service_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_example_service_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_example_service_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_example_service_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_example_service_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse); i {
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
			RawDescriptor: file_example_service_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_example_service_v1_service_proto_goTypes,
		DependencyIndexes: file_example_service_v1_service_proto_depIdxs,
		MessageInfos:      file_example_service_v1_service_proto_msgTypes,
	}.Build()
	File_example_service_v1_service_proto = out.File
	file_example_service_v1_service_proto_rawDesc = nil
	file_example_service_v1_service_proto_goTypes = nil
	file_example_service_v1_service_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: natsPb.proto

package natsPb

import (
	proto "github.com/golang/protobuf/proto"
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

// is user exist req and res
type IsUserExist_REQUESTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *IsUserExist_REQUESTRequest) Reset() {
	*x = IsUserExist_REQUESTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_natsPb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserExist_REQUESTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExist_REQUESTRequest) ProtoMessage() {}

func (x *IsUserExist_REQUESTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_natsPb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExist_REQUESTRequest.ProtoReflect.Descriptor instead.
func (*IsUserExist_REQUESTRequest) Descriptor() ([]byte, []int) {
	return file_natsPb_proto_rawDescGZIP(), []int{0}
}

func (x *IsUserExist_REQUESTRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type IsUserExist_REQUESTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
}

func (x *IsUserExist_REQUESTResponse) Reset() {
	*x = IsUserExist_REQUESTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_natsPb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserExist_REQUESTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExist_REQUESTResponse) ProtoMessage() {}

func (x *IsUserExist_REQUESTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_natsPb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExist_REQUESTResponse.ProtoReflect.Descriptor instead.
func (*IsUserExist_REQUESTResponse) Descriptor() ([]byte, []int) {
	return file_natsPb_proto_rawDescGZIP(), []int{1}
}

func (x *IsUserExist_REQUESTResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

type DeleteUserPosts_EVENT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserPosts_EVENT) Reset() {
	*x = DeleteUserPosts_EVENT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_natsPb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserPosts_EVENT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserPosts_EVENT) ProtoMessage() {}

func (x *DeleteUserPosts_EVENT) ProtoReflect() protoreflect.Message {
	mi := &file_natsPb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserPosts_EVENT.ProtoReflect.Descriptor instead.
func (*DeleteUserPosts_EVENT) Descriptor() ([]byte, []int) {
	return file_natsPb_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteUserPosts_EVENT) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_natsPb_proto protoreflect.FileDescriptor

var file_natsPb_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6e, 0x61, 0x74, 0x73, 0x50, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6e, 0x61, 0x74, 0x73, 0x50, 0x62, 0x22, 0x35, 0x0a, 0x1b, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x34, 0x0a,
	0x1c, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x42, 0x09, 0x5a, 0x07,
	0x2f, 0x6e, 0x61, 0x74, 0x73, 0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_natsPb_proto_rawDescOnce sync.Once
	file_natsPb_proto_rawDescData = file_natsPb_proto_rawDesc
)

func file_natsPb_proto_rawDescGZIP() []byte {
	file_natsPb_proto_rawDescOnce.Do(func() {
		file_natsPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_natsPb_proto_rawDescData)
	})
	return file_natsPb_proto_rawDescData
}

var file_natsPb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_natsPb_proto_goTypes = []interface{}{
	(*IsUserExist_REQUESTRequest)(nil),  // 0: natsPb.IsUserExist_REQUEST_request
	(*IsUserExist_REQUESTResponse)(nil), // 1: natsPb.IsUserExist_REQUEST_response
	(*DeleteUserPosts_EVENT)(nil),       // 2: natsPb.DeleteUserPosts_EVENT
}
var file_natsPb_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_natsPb_proto_init() }
func file_natsPb_proto_init() {
	if File_natsPb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_natsPb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserExist_REQUESTRequest); i {
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
		file_natsPb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserExist_REQUESTResponse); i {
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
		file_natsPb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserPosts_EVENT); i {
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
			RawDescriptor: file_natsPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_natsPb_proto_goTypes,
		DependencyIndexes: file_natsPb_proto_depIdxs,
		MessageInfos:      file_natsPb_proto_msgTypes,
	}.Build()
	File_natsPb_proto = out.File
	file_natsPb_proto_rawDesc = nil
	file_natsPb_proto_goTypes = nil
	file_natsPb_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protobufs/fileRequests.proto

package pbs

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

type PostFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path       string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Permission uint32 `protobuf:"varint,2,opt,name=permission,proto3" json:"permission,omitempty"`
}

func (x *PostFileRequest) Reset() {
	*x = PostFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobufs_fileRequests_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFileRequest) ProtoMessage() {}

func (x *PostFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobufs_fileRequests_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFileRequest.ProtoReflect.Descriptor instead.
func (*PostFileRequest) Descriptor() ([]byte, []int) {
	return file_protobufs_fileRequests_proto_rawDescGZIP(), []int{0}
}

func (x *PostFileRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PostFileRequest) GetPermission() uint32 {
	if x != nil {
		return x.Permission
	}
	return 0
}

var File_protobufs_fileRequests_proto protoreflect.FileDescriptor

var file_protobufs_fileRequests_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45,
	0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x70, 0x62, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobufs_fileRequests_proto_rawDescOnce sync.Once
	file_protobufs_fileRequests_proto_rawDescData = file_protobufs_fileRequests_proto_rawDesc
)

func file_protobufs_fileRequests_proto_rawDescGZIP() []byte {
	file_protobufs_fileRequests_proto_rawDescOnce.Do(func() {
		file_protobufs_fileRequests_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufs_fileRequests_proto_rawDescData)
	})
	return file_protobufs_fileRequests_proto_rawDescData
}

var file_protobufs_fileRequests_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobufs_fileRequests_proto_goTypes = []any{
	(*PostFileRequest)(nil), // 0: PostFileRequest
}
var file_protobufs_fileRequests_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobufs_fileRequests_proto_init() }
func file_protobufs_fileRequests_proto_init() {
	if File_protobufs_fileRequests_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobufs_fileRequests_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PostFileRequest); i {
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
			RawDescriptor: file_protobufs_fileRequests_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobufs_fileRequests_proto_goTypes,
		DependencyIndexes: file_protobufs_fileRequests_proto_depIdxs,
		MessageInfos:      file_protobufs_fileRequests_proto_msgTypes,
	}.Build()
	File_protobufs_fileRequests_proto = out.File
	file_protobufs_fileRequests_proto_rawDesc = nil
	file_protobufs_fileRequests_proto_goTypes = nil
	file_protobufs_fileRequests_proto_depIdxs = nil
}
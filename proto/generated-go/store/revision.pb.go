// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: store/revision.proto

package store

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RevisionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: projects/{project}/releases/{release}
	// Can be empty.
	Release string `protobuf:"bytes,1,opt,name=release,proto3" json:"release,omitempty"`
	// The sheet that holds the content.
	// Format: projects/{project}/sheets/{sheet}
	Sheet string `protobuf:"bytes,2,opt,name=sheet,proto3" json:"sheet,omitempty"`
	// The SHA256 hash value of the sheet.
	SheetSha256 string          `protobuf:"bytes,3,opt,name=sheet_sha256,json=sheetSha256,proto3" json:"sheet_sha256,omitempty"`
	Type        ReleaseFileType `protobuf:"varint,4,opt,name=type,proto3,enum=bytebase.store.ReleaseFileType" json:"type,omitempty"`
	Version     string          `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	// The name of the file in the release.
	// Expressed as a path, e.g. `2.2/V0001_create_table.sql`
	// Can be empty.
	File string `protobuf:"bytes,6,opt,name=file,proto3" json:"file,omitempty"`
	// The task run associated with the revision.
	// Can be empty.
	// Format: projects/{project}/rollouts/{rollout}/stages/{stage}/tasks/{task}/taskRuns/{taskRun}
	TaskRun string `protobuf:"bytes,7,opt,name=task_run,json=taskRun,proto3" json:"task_run,omitempty"`
}

func (x *RevisionPayload) Reset() {
	*x = RevisionPayload{}
	mi := &file_store_revision_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevisionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevisionPayload) ProtoMessage() {}

func (x *RevisionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_store_revision_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevisionPayload.ProtoReflect.Descriptor instead.
func (*RevisionPayload) Descriptor() ([]byte, []int) {
	return file_store_revision_proto_rawDescGZIP(), []int{0}
}

func (x *RevisionPayload) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *RevisionPayload) GetSheet() string {
	if x != nil {
		return x.Sheet
	}
	return ""
}

func (x *RevisionPayload) GetSheetSha256() string {
	if x != nil {
		return x.SheetSha256
	}
	return ""
}

func (x *RevisionPayload) GetType() ReleaseFileType {
	if x != nil {
		return x.Type
	}
	return ReleaseFileType_TYPE_UNSPECIFIED
}

func (x *RevisionPayload) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RevisionPayload) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *RevisionPayload) GetTaskRun() string {
	if x != nil {
		return x.TaskRun
	}
	return ""
}

var File_store_revision_proto protoreflect.FileDescriptor

var file_store_revision_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x41, 0x16,
	0x0a, 0x14, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x2d, 0x0a, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17,
	0xfa, 0x41, 0x14, 0x0a, 0x12, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x53, 0x68, 0x65, 0x65, 0x74, 0x52, 0x05, 0x73, 0x68, 0x65, 0x65, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x68, 0x65, 0x65, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x68, 0x65, 0x65, 0x74, 0x53, 0x68, 0x61, 0x32, 0x35,
	0x36, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1f, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0xfa, 0x41, 0x16, 0x0a, 0x14, 0x62, 0x79, 0x74,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75,
	0x6e, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_revision_proto_rawDescOnce sync.Once
	file_store_revision_proto_rawDescData = file_store_revision_proto_rawDesc
)

func file_store_revision_proto_rawDescGZIP() []byte {
	file_store_revision_proto_rawDescOnce.Do(func() {
		file_store_revision_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_revision_proto_rawDescData)
	})
	return file_store_revision_proto_rawDescData
}

var file_store_revision_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_revision_proto_goTypes = []any{
	(*RevisionPayload)(nil), // 0: bytebase.store.RevisionPayload
	(ReleaseFileType)(0),    // 1: bytebase.store.ReleaseFileType
}
var file_store_revision_proto_depIdxs = []int32{
	1, // 0: bytebase.store.RevisionPayload.type:type_name -> bytebase.store.ReleaseFileType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_store_revision_proto_init() }
func file_store_revision_proto_init() {
	if File_store_revision_proto != nil {
		return
	}
	file_store_release_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_revision_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_revision_proto_goTypes,
		DependencyIndexes: file_store_revision_proto_depIdxs,
		MessageInfos:      file_store_revision_proto_msgTypes,
	}.Build()
	File_store_revision_proto = out.File
	file_store_revision_proto_rawDesc = nil
	file_store_revision_proto_goTypes = nil
	file_store_revision_proto_depIdxs = nil
}

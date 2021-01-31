// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: RecordTemplateProto.proto

package pb_templates

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

type RecordTemplateProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Human readable name to quickly identify type (non-unique)
	FriendlyName string `protobuf:"bytes,1,opt,name=friendlyName,proto3" json:"friendlyName,omitempty"`
	// Description of the purpose behind this new type
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Compiled descriptor including dependencies; Defines fields
	DescriptorSetProto []byte `protobuf:"bytes,4,opt,name=DescriptorSetProto,proto3" json:"DescriptorSetProto,omitempty"`
	// List of unique template identifiers required for use with this template
	Extends []uint32 `protobuf:"fixed32,13,rep,packed,name=extends,proto3" json:"extends,omitempty"`
	// Populated by oipd with the unique identifier for this type
	Identifier uint32 `protobuf:"fixed32,10,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *RecordTemplateProto) Reset() {
	*x = RecordTemplateProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RecordTemplateProto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordTemplateProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordTemplateProto) ProtoMessage() {}

func (x *RecordTemplateProto) ProtoReflect() protoreflect.Message {
	mi := &file_RecordTemplateProto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordTemplateProto.ProtoReflect.Descriptor instead.
func (*RecordTemplateProto) Descriptor() ([]byte, []int) {
	return file_RecordTemplateProto_proto_rawDescGZIP(), []int{0}
}

func (x *RecordTemplateProto) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *RecordTemplateProto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RecordTemplateProto) GetDescriptorSetProto() []byte {
	if x != nil {
		return x.DescriptorSetProto
	}
	return nil
}

func (x *RecordTemplateProto) GetExtends() []uint32 {
	if x != nil {
		return x.Extends
	}
	return nil
}

func (x *RecordTemplateProto) GetIdentifier() uint32 {
	if x != nil {
		return x.Identifier
	}
	return 0
}

var File_RecordTemplateProto_proto protoreflect.FileDescriptor

var file_RecordTemplateProto_proto_rawDesc = []byte{
	0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x69, 0x70,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22,
	0xcb, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x07, 0x52, 0x07,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0d, 0x42, 0x3d, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x69, 0x70, 0x77,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x5f, 0x6f, 0x69,
	0x70, 0x35, 0x2f, 0x70, 0x62, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x3b,
	0x70, 0x62, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RecordTemplateProto_proto_rawDescOnce sync.Once
	file_RecordTemplateProto_proto_rawDescData = file_RecordTemplateProto_proto_rawDesc
)

func file_RecordTemplateProto_proto_rawDescGZIP() []byte {
	file_RecordTemplateProto_proto_rawDescOnce.Do(func() {
		file_RecordTemplateProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_RecordTemplateProto_proto_rawDescData)
	})
	return file_RecordTemplateProto_proto_rawDescData
}

var file_RecordTemplateProto_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RecordTemplateProto_proto_goTypes = []interface{}{
	(*RecordTemplateProto)(nil), // 0: oipProto.templates.RecordTemplateProto
}
var file_RecordTemplateProto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RecordTemplateProto_proto_init() }
func file_RecordTemplateProto_proto_init() {
	if File_RecordTemplateProto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RecordTemplateProto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordTemplateProto); i {
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
			RawDescriptor: file_RecordTemplateProto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RecordTemplateProto_proto_goTypes,
		DependencyIndexes: file_RecordTemplateProto_proto_depIdxs,
		MessageInfos:      file_RecordTemplateProto_proto_msgTypes,
	}.Build()
	File_RecordTemplateProto_proto = out.File
	file_RecordTemplateProto_proto_rawDesc = nil
	file_RecordTemplateProto_proto_goTypes = nil
	file_RecordTemplateProto_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: historian.proto

package pb_historian

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

type DataPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version                  int32   `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	PubKey                   []byte  `protobuf:"bytes,2,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	MiningRigRentalsLast10   float64 `protobuf:"fixed64,3,opt,name=MiningRigRentalsLast10,proto3" json:"MiningRigRentalsLast10,omitempty"`
	MiningRigRentalsLast24Hr float64 `protobuf:"fixed64,4,opt,name=MiningRigRentalsLast24Hr,proto3" json:"MiningRigRentalsLast24Hr,omitempty"`
	AutominerPoolHashrate    float64 `protobuf:"fixed64,5,opt,name=AutominerPoolHashrate,proto3" json:"AutominerPoolHashrate,omitempty"`
	FloNetHashRate           float64 `protobuf:"fixed64,6,opt,name=FloNetHashRate,proto3" json:"FloNetHashRate,omitempty"`
	FloMarketPriceBTC        float64 `protobuf:"fixed64,7,opt,name=FloMarketPriceBTC,proto3" json:"FloMarketPriceBTC,omitempty"`
	FloMarketPriceUSD        float64 `protobuf:"fixed64,8,opt,name=FloMarketPriceUSD,proto3" json:"FloMarketPriceUSD,omitempty"`
	LtcMarketPriceUSD        float64 `protobuf:"fixed64,9,opt,name=LtcMarketPriceUSD,proto3" json:"LtcMarketPriceUSD,omitempty"`
	NiceHashLast             float64 `protobuf:"fixed64,10,opt,name=NiceHashLast,proto3" json:"NiceHashLast,omitempty"`
	NiceHash24Hr             float64 `protobuf:"fixed64,11,opt,name=NiceHash24hr,proto3" json:"NiceHash24hr,omitempty"`
}

func (x *DataPoint) Reset() {
	*x = DataPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_historian_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPoint) ProtoMessage() {}

func (x *DataPoint) ProtoReflect() protoreflect.Message {
	mi := &file_historian_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPoint.ProtoReflect.Descriptor instead.
func (*DataPoint) Descriptor() ([]byte, []int) {
	return file_historian_proto_rawDescGZIP(), []int{0}
}

func (x *DataPoint) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DataPoint) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *DataPoint) GetMiningRigRentalsLast10() float64 {
	if x != nil {
		return x.MiningRigRentalsLast10
	}
	return 0
}

func (x *DataPoint) GetMiningRigRentalsLast24Hr() float64 {
	if x != nil {
		return x.MiningRigRentalsLast24Hr
	}
	return 0
}

func (x *DataPoint) GetAutominerPoolHashrate() float64 {
	if x != nil {
		return x.AutominerPoolHashrate
	}
	return 0
}

func (x *DataPoint) GetFloNetHashRate() float64 {
	if x != nil {
		return x.FloNetHashRate
	}
	return 0
}

func (x *DataPoint) GetFloMarketPriceBTC() float64 {
	if x != nil {
		return x.FloMarketPriceBTC
	}
	return 0
}

func (x *DataPoint) GetFloMarketPriceUSD() float64 {
	if x != nil {
		return x.FloMarketPriceUSD
	}
	return 0
}

func (x *DataPoint) GetLtcMarketPriceUSD() float64 {
	if x != nil {
		return x.LtcMarketPriceUSD
	}
	return 0
}

func (x *DataPoint) GetNiceHashLast() float64 {
	if x != nil {
		return x.NiceHashLast
	}
	return 0
}

func (x *DataPoint) GetNiceHash24Hr() float64 {
	if x != nil {
		return x.NiceHash24Hr
	}
	return 0
}

type Payout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version int32 `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *Payout) Reset() {
	*x = Payout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_historian_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payout) ProtoMessage() {}

func (x *Payout) ProtoReflect() protoreflect.Message {
	mi := &file_historian_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payout.ProtoReflect.Descriptor instead.
func (*Payout) Descriptor() ([]byte, []int) {
	return file_historian_proto_rawDescGZIP(), []int{1}
}

func (x *Payout) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_historian_proto protoreflect.FileDescriptor

var file_historian_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x6f, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x6e, 0x22, 0xe1, 0x03, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x16, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x69, 0x67, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x31, 0x30, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x16, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x69, 0x67,
	0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x31, 0x30, 0x12, 0x3a, 0x0a,
	0x18, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x69, 0x67, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x73, 0x4c, 0x61, 0x73, 0x74, 0x32, 0x34, 0x48, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x18, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x69, 0x67, 0x52, 0x65, 0x6e, 0x74, 0x61, 0x6c,
	0x73, 0x4c, 0x61, 0x73, 0x74, 0x32, 0x34, 0x48, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x41, 0x75, 0x74,
	0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x61, 0x73, 0x68, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x15, 0x41, 0x75, 0x74, 0x6f, 0x6d, 0x69,
	0x6e, 0x65, 0x72, 0x50, 0x6f, 0x6f, 0x6c, 0x48, 0x61, 0x73, 0x68, 0x72, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x46, 0x6c, 0x6f, 0x4e, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x46, 0x6c, 0x6f, 0x4e, 0x65, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x6c, 0x6f, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x54, 0x43, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x11, 0x46, 0x6c, 0x6f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x42, 0x54, 0x43, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x6c, 0x6f, 0x4d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x11, 0x46, 0x6c, 0x6f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x55, 0x53, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x4c, 0x74, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x53, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11,
	0x4c, 0x74, 0x63, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x55, 0x53,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x69, 0x63, 0x65, 0x48, 0x61, 0x73, 0x68, 0x4c, 0x61, 0x73,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x4e, 0x69, 0x63, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x69, 0x63, 0x65, 0x48, 0x61, 0x73,
	0x68, 0x32, 0x34, 0x68, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x4e, 0x69, 0x63,
	0x65, 0x48, 0x61, 0x73, 0x68, 0x32, 0x34, 0x68, 0x72, 0x22, 0x22, 0x0a, 0x06, 0x50, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x69, 0x70, 0x77,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x5f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6e, 0x3b, 0x70, 0x62, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_historian_proto_rawDescOnce sync.Once
	file_historian_proto_rawDescData = file_historian_proto_rawDesc
)

func file_historian_proto_rawDescGZIP() []byte {
	file_historian_proto_rawDescOnce.Do(func() {
		file_historian_proto_rawDescData = protoimpl.X.CompressGZIP(file_historian_proto_rawDescData)
	})
	return file_historian_proto_rawDescData
}

var file_historian_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_historian_proto_goTypes = []interface{}{
	(*DataPoint)(nil), // 0: oipProto.historian.DataPoint
	(*Payout)(nil),    // 1: oipProto.historian.Payout
}
var file_historian_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_historian_proto_init() }
func file_historian_proto_init() {
	if File_historian_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_historian_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPoint); i {
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
		file_historian_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payout); i {
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
			RawDescriptor: file_historian_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_historian_proto_goTypes,
		DependencyIndexes: file_historian_proto_depIdxs,
		MessageInfos:      file_historian_proto_msgTypes,
	}.Build()
	File_historian_proto = out.File
	file_historian_proto_rawDesc = nil
	file_historian_proto_goTypes = nil
	file_historian_proto_depIdxs = nil
}

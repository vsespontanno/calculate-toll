// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: types/ptypes.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type None struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *None) Reset() {
	*x = None{}
	mi := &file_types_ptypes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *None) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*None) ProtoMessage() {}

func (x *None) ProtoReflect() protoreflect.Message {
	mi := &file_types_ptypes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use None.ProtoReflect.Descriptor instead.
func (*None) Descriptor() ([]byte, []int) {
	return file_types_ptypes_proto_rawDescGZIP(), []int{0}
}

type GetInvoiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ObuID         int32                  `protobuf:"varint,1,opt,name=ObuID,proto3" json:"ObuID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInvoiceRequest) Reset() {
	*x = GetInvoiceRequest{}
	mi := &file_types_ptypes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvoiceRequest) ProtoMessage() {}

func (x *GetInvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_ptypes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvoiceRequest.ProtoReflect.Descriptor instead.
func (*GetInvoiceRequest) Descriptor() ([]byte, []int) {
	return file_types_ptypes_proto_rawDescGZIP(), []int{1}
}

func (x *GetInvoiceRequest) GetObuID() int32 {
	if x != nil {
		return x.ObuID
	}
	return 0
}

type AggregateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ObuID         int32                  `protobuf:"varint,1,opt,name=ObuID,proto3" json:"ObuID,omitempty"`
	Value         float64                `protobuf:"fixed64,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Unix          int64                  `protobuf:"varint,3,opt,name=Unix,proto3" json:"Unix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AggregateRequest) Reset() {
	*x = AggregateRequest{}
	mi := &file_types_ptypes_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateRequest) ProtoMessage() {}

func (x *AggregateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_ptypes_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateRequest.ProtoReflect.Descriptor instead.
func (*AggregateRequest) Descriptor() ([]byte, []int) {
	return file_types_ptypes_proto_rawDescGZIP(), []int{2}
}

func (x *AggregateRequest) GetObuID() int32 {
	if x != nil {
		return x.ObuID
	}
	return 0
}

func (x *AggregateRequest) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *AggregateRequest) GetUnix() int64 {
	if x != nil {
		return x.Unix
	}
	return 0
}

var File_types_ptypes_proto protoreflect.FileDescriptor

const file_types_ptypes_proto_rawDesc = "" +
	"\n" +
	"\x12types/ptypes.proto\"\x06\n" +
	"\x04None\")\n" +
	"\x11GetInvoiceRequest\x12\x14\n" +
	"\x05ObuID\x18\x01 \x01(\x05R\x05ObuID\"R\n" +
	"\x10AggregateRequest\x12\x14\n" +
	"\x05ObuID\x18\x01 \x01(\x05R\x05ObuID\x12\x14\n" +
	"\x05Value\x18\x02 \x01(\x01R\x05Value\x12\x12\n" +
	"\x04Unix\x18\x03 \x01(\x03R\x04Unix23\n" +
	"\n" +
	"Aggregator\x12%\n" +
	"\tAggregate\x12\x11.AggregateRequest\x1a\x05.NoneB.Z,github.com/vsespontanno/calculate-toll/typesb\x06proto3"

var (
	file_types_ptypes_proto_rawDescOnce sync.Once
	file_types_ptypes_proto_rawDescData []byte
)

func file_types_ptypes_proto_rawDescGZIP() []byte {
	file_types_ptypes_proto_rawDescOnce.Do(func() {
		file_types_ptypes_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_types_ptypes_proto_rawDesc), len(file_types_ptypes_proto_rawDesc)))
	})
	return file_types_ptypes_proto_rawDescData
}

var file_types_ptypes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_types_ptypes_proto_goTypes = []any{
	(*None)(nil),              // 0: None
	(*GetInvoiceRequest)(nil), // 1: GetInvoiceRequest
	(*AggregateRequest)(nil),  // 2: AggregateRequest
}
var file_types_ptypes_proto_depIdxs = []int32{
	2, // 0: Aggregator.Aggregate:input_type -> AggregateRequest
	0, // 1: Aggregator.Aggregate:output_type -> None
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_ptypes_proto_init() }
func file_types_ptypes_proto_init() {
	if File_types_ptypes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_types_ptypes_proto_rawDesc), len(file_types_ptypes_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_types_ptypes_proto_goTypes,
		DependencyIndexes: file_types_ptypes_proto_depIdxs,
		MessageInfos:      file_types_ptypes_proto_msgTypes,
	}.Build()
	File_types_ptypes_proto = out.File
	file_types_ptypes_proto_goTypes = nil
	file_types_ptypes_proto_depIdxs = nil
}

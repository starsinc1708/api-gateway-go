// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: shipping_query.proto

package telegram_api

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

type ShippingQuery struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From            *User                  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	InvoicePayload  string                 `protobuf:"bytes,3,opt,name=invoice_payload,json=invoicePayload,proto3" json:"invoice_payload,omitempty"`
	ShippingAddress *ShippingAddress       `protobuf:"bytes,4,opt,name=shipping_address,json=shippingAddress,proto3" json:"shipping_address,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ShippingQuery) Reset() {
	*x = ShippingQuery{}
	mi := &file_shipping_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShippingQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShippingQuery) ProtoMessage() {}

func (x *ShippingQuery) ProtoReflect() protoreflect.Message {
	mi := &file_shipping_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShippingQuery.ProtoReflect.Descriptor instead.
func (*ShippingQuery) Descriptor() ([]byte, []int) {
	return file_shipping_query_proto_rawDescGZIP(), []int{0}
}

func (x *ShippingQuery) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ShippingQuery) GetFrom() *User {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ShippingQuery) GetInvoicePayload() string {
	if x != nil {
		return x.InvoicePayload
	}
	return ""
}

func (x *ShippingQuery) GetShippingAddress() *ShippingAddress {
	if x != nil {
		return x.ShippingAddress
	}
	return nil
}

var File_shipping_query_proto protoreflect.FileDescriptor

var file_shipping_query_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x3c, 0x0a, 0x10, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0f,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_shipping_query_proto_rawDescOnce sync.Once
	file_shipping_query_proto_rawDescData []byte
)

func file_shipping_query_proto_rawDescGZIP() []byte {
	file_shipping_query_proto_rawDescOnce.Do(func() {
		file_shipping_query_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shipping_query_proto_rawDesc), len(file_shipping_query_proto_rawDesc)))
	})
	return file_shipping_query_proto_rawDescData
}

var file_shipping_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shipping_query_proto_goTypes = []any{
	(*ShippingQuery)(nil),   // 0: shipping_query
	(*User)(nil),            // 1: user
	(*ShippingAddress)(nil), // 2: shipping_address
}
var file_shipping_query_proto_depIdxs = []int32{
	1, // 0: shipping_query.from:type_name -> user
	2, // 1: shipping_query.shipping_address:type_name -> shipping_address
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shipping_query_proto_init() }
func file_shipping_query_proto_init() {
	if File_shipping_query_proto != nil {
		return
	}
	file_user_proto_init()
	file_shipping_address_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shipping_query_proto_rawDesc), len(file_shipping_query_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shipping_query_proto_goTypes,
		DependencyIndexes: file_shipping_query_proto_depIdxs,
		MessageInfos:      file_shipping_query_proto_msgTypes,
	}.Build()
	File_shipping_query_proto = out.File
	file_shipping_query_proto_goTypes = nil
	file_shipping_query_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: encrypted_passport_element.proto

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

type EncryptedPassportElement struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          string                 `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data          *string                `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
	PhoneNumber   *string                `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3,oneof" json:"phone_number,omitempty"`
	Email         *string                `protobuf:"bytes,4,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Files         []*PassportFile        `protobuf:"bytes,5,rep,name=files,proto3" json:"files,omitempty"`
	FrontSide     *PassportFile          `protobuf:"bytes,6,opt,name=front_side,json=frontSide,proto3,oneof" json:"front_side,omitempty"`
	ReverseSide   *PassportFile          `protobuf:"bytes,7,opt,name=reverse_side,json=reverseSide,proto3,oneof" json:"reverse_side,omitempty"`
	Selfie        *PassportFile          `protobuf:"bytes,8,opt,name=selfie,proto3,oneof" json:"selfie,omitempty"`
	Translation   []*PassportFile        `protobuf:"bytes,9,rep,name=translation,proto3" json:"translation,omitempty"`
	Hash          string                 `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EncryptedPassportElement) Reset() {
	*x = EncryptedPassportElement{}
	mi := &file_encrypted_passport_element_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EncryptedPassportElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedPassportElement) ProtoMessage() {}

func (x *EncryptedPassportElement) ProtoReflect() protoreflect.Message {
	mi := &file_encrypted_passport_element_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedPassportElement.ProtoReflect.Descriptor instead.
func (*EncryptedPassportElement) Descriptor() ([]byte, []int) {
	return file_encrypted_passport_element_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptedPassportElement) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EncryptedPassportElement) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

func (x *EncryptedPassportElement) GetPhoneNumber() string {
	if x != nil && x.PhoneNumber != nil {
		return *x.PhoneNumber
	}
	return ""
}

func (x *EncryptedPassportElement) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *EncryptedPassportElement) GetFiles() []*PassportFile {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *EncryptedPassportElement) GetFrontSide() *PassportFile {
	if x != nil {
		return x.FrontSide
	}
	return nil
}

func (x *EncryptedPassportElement) GetReverseSide() *PassportFile {
	if x != nil {
		return x.ReverseSide
	}
	return nil
}

func (x *EncryptedPassportElement) GetSelfie() *PassportFile {
	if x != nil {
		return x.Selfie
	}
	return nil
}

func (x *EncryptedPassportElement) GetTranslation() []*PassportFile {
	if x != nil {
		return x.Translation
	}
	return nil
}

func (x *EncryptedPassportElement) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_encrypted_passport_element_proto protoreflect.FileDescriptor

var file_encrypted_passport_element_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x03, 0x0a, 0x1a, 0x65, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x48, 0x03, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x69, 0x64, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x36, 0x0a, 0x0c, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x64, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x65, 0x53, 0x69, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x66,
	0x69, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x48, 0x05, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x66,
	0x69, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x6d,
	0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_encrypted_passport_element_proto_rawDescOnce sync.Once
	file_encrypted_passport_element_proto_rawDescData []byte
)

func file_encrypted_passport_element_proto_rawDescGZIP() []byte {
	file_encrypted_passport_element_proto_rawDescOnce.Do(func() {
		file_encrypted_passport_element_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_encrypted_passport_element_proto_rawDesc), len(file_encrypted_passport_element_proto_rawDesc)))
	})
	return file_encrypted_passport_element_proto_rawDescData
}

var file_encrypted_passport_element_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_encrypted_passport_element_proto_goTypes = []any{
	(*EncryptedPassportElement)(nil), // 0: encrypted_passport_element
	(*PassportFile)(nil),             // 1: passport_file
}
var file_encrypted_passport_element_proto_depIdxs = []int32{
	1, // 0: encrypted_passport_element.files:type_name -> passport_file
	1, // 1: encrypted_passport_element.front_side:type_name -> passport_file
	1, // 2: encrypted_passport_element.reverse_side:type_name -> passport_file
	1, // 3: encrypted_passport_element.selfie:type_name -> passport_file
	1, // 4: encrypted_passport_element.translation:type_name -> passport_file
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_encrypted_passport_element_proto_init() }
func file_encrypted_passport_element_proto_init() {
	if File_encrypted_passport_element_proto != nil {
		return
	}
	file_passport_file_proto_init()
	file_encrypted_passport_element_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_encrypted_passport_element_proto_rawDesc), len(file_encrypted_passport_element_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encrypted_passport_element_proto_goTypes,
		DependencyIndexes: file_encrypted_passport_element_proto_depIdxs,
		MessageInfos:      file_encrypted_passport_element_proto_msgTypes,
	}.Build()
	File_encrypted_passport_element_proto = out.File
	file_encrypted_passport_element_proto_goTypes = nil
	file_encrypted_passport_element_proto_depIdxs = nil
}

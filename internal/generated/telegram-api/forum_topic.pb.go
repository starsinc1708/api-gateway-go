// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: forum_topic.proto

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

type ForumTopic struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	MessageThreadId   int32                  `protobuf:"varint,1,opt,name=message_thread_id,json=messageThreadId,proto3" json:"message_thread_id,omitempty"`
	Name              string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IconColor         int32                  `protobuf:"varint,3,opt,name=icon_color,json=iconColor,proto3" json:"icon_color,omitempty"`
	IconCustomEmojiId *string                `protobuf:"bytes,4,opt,name=icon_custom_emoji_id,json=iconCustomEmojiId,proto3,oneof" json:"icon_custom_emoji_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ForumTopic) Reset() {
	*x = ForumTopic{}
	mi := &file_forum_topic_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForumTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForumTopic) ProtoMessage() {}

func (x *ForumTopic) ProtoReflect() protoreflect.Message {
	mi := &file_forum_topic_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForumTopic.ProtoReflect.Descriptor instead.
func (*ForumTopic) Descriptor() ([]byte, []int) {
	return file_forum_topic_proto_rawDescGZIP(), []int{0}
}

func (x *ForumTopic) GetMessageThreadId() int32 {
	if x != nil {
		return x.MessageThreadId
	}
	return 0
}

func (x *ForumTopic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ForumTopic) GetIconColor() int32 {
	if x != nil {
		return x.IconColor
	}
	return 0
}

func (x *ForumTopic) GetIconCustomEmojiId() string {
	if x != nil && x.IconCustomEmojiId != nil {
		return *x.IconCustomEmojiId
	}
	return ""
}

type ForumTopicCreated struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Name              string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IconColor         int32                  `protobuf:"varint,2,opt,name=icon_color,json=iconColor,proto3" json:"icon_color,omitempty"`
	IconCustomEmojiId *string                `protobuf:"bytes,3,opt,name=icon_custom_emoji_id,json=iconCustomEmojiId,proto3,oneof" json:"icon_custom_emoji_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ForumTopicCreated) Reset() {
	*x = ForumTopicCreated{}
	mi := &file_forum_topic_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForumTopicCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForumTopicCreated) ProtoMessage() {}

func (x *ForumTopicCreated) ProtoReflect() protoreflect.Message {
	mi := &file_forum_topic_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForumTopicCreated.ProtoReflect.Descriptor instead.
func (*ForumTopicCreated) Descriptor() ([]byte, []int) {
	return file_forum_topic_proto_rawDescGZIP(), []int{1}
}

func (x *ForumTopicCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ForumTopicCreated) GetIconColor() int32 {
	if x != nil {
		return x.IconColor
	}
	return 0
}

func (x *ForumTopicCreated) GetIconCustomEmojiId() string {
	if x != nil && x.IconCustomEmojiId != nil {
		return *x.IconCustomEmojiId
	}
	return ""
}

type ForumTopicEdited struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Name              *string                `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	IconCustomEmojiId *string                `protobuf:"bytes,2,opt,name=icon_custom_emoji_id,json=iconCustomEmojiId,proto3,oneof" json:"icon_custom_emoji_id,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ForumTopicEdited) Reset() {
	*x = ForumTopicEdited{}
	mi := &file_forum_topic_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForumTopicEdited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForumTopicEdited) ProtoMessage() {}

func (x *ForumTopicEdited) ProtoReflect() protoreflect.Message {
	mi := &file_forum_topic_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForumTopicEdited.ProtoReflect.Descriptor instead.
func (*ForumTopicEdited) Descriptor() ([]byte, []int) {
	return file_forum_topic_proto_rawDescGZIP(), []int{2}
}

func (x *ForumTopicEdited) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ForumTopicEdited) GetIconCustomEmojiId() string {
	if x != nil && x.IconCustomEmojiId != nil {
		return *x.IconCustomEmojiId
	}
	return ""
}

type ForumTopicClosed struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ForumTopicClosed) Reset() {
	*x = ForumTopicClosed{}
	mi := &file_forum_topic_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForumTopicClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForumTopicClosed) ProtoMessage() {}

func (x *ForumTopicClosed) ProtoReflect() protoreflect.Message {
	mi := &file_forum_topic_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForumTopicClosed.ProtoReflect.Descriptor instead.
func (*ForumTopicClosed) Descriptor() ([]byte, []int) {
	return file_forum_topic_proto_rawDescGZIP(), []int{3}
}

type ForumTopicReopened struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ForumTopicReopened) Reset() {
	*x = ForumTopicReopened{}
	mi := &file_forum_topic_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForumTopicReopened) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForumTopicReopened) ProtoMessage() {}

func (x *ForumTopicReopened) ProtoReflect() protoreflect.Message {
	mi := &file_forum_topic_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForumTopicReopened.ProtoReflect.Descriptor instead.
func (*ForumTopicReopened) Descriptor() ([]byte, []int) {
	return file_forum_topic_proto_rawDescGZIP(), []int{4}
}

type GeneralForumTopicHidden struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeneralForumTopicHidden) Reset() {
	*x = GeneralForumTopicHidden{}
	mi := &file_forum_topic_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeneralForumTopicHidden) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralForumTopicHidden) ProtoMessage() {}

func (x *GeneralForumTopicHidden) ProtoReflect() protoreflect.Message {
	mi := &file_forum_topic_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralForumTopicHidden.ProtoReflect.Descriptor instead.
func (*GeneralForumTopicHidden) Descriptor() ([]byte, []int) {
	return file_forum_topic_proto_rawDescGZIP(), []int{5}
}

type GeneralForumTopicUnhidden struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeneralForumTopicUnhidden) Reset() {
	*x = GeneralForumTopicUnhidden{}
	mi := &file_forum_topic_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeneralForumTopicUnhidden) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralForumTopicUnhidden) ProtoMessage() {}

func (x *GeneralForumTopicUnhidden) ProtoReflect() protoreflect.Message {
	mi := &file_forum_topic_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralForumTopicUnhidden.ProtoReflect.Descriptor instead.
func (*GeneralForumTopicUnhidden) Descriptor() ([]byte, []int) {
	return file_forum_topic_proto_rawDescGZIP(), []int{6}
}

var File_forum_topic_proto protoreflect.FileDescriptor

var file_forum_topic_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x34, 0x0a, 0x14, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x11, 0x69, 0x63, 0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6d,
	0x6f, 0x6a, 0x69, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x69, 0x63, 0x6f,
	0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x5f, 0x69,
	0x64, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x14,
	0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6d, 0x6f, 0x6a,
	0x69, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x69, 0x63,
	0x6f, 0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x65, 0x6d, 0x6f, 0x6a, 0x69, 0x5f, 0x69, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x12,
	0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x65, 0x64, 0x69, 0x74,
	0x65, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x14, 0x69,
	0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6d, 0x6f, 0x6a, 0x69,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x11, 0x69, 0x63, 0x6f,
	0x6e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x45, 0x6d, 0x6f, 0x6a, 0x69, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x69,
	0x63, 0x6f, 0x6e, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x65, 0x6d, 0x6f, 0x6a, 0x69,
	0x5f, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x66, 0x6f, 0x72,
	0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x6f, 0x70, 0x65, 0x6e, 0x65,
	0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x66, 0x6f, 0x72,
	0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22,
	0x1e, 0x0a, 0x1c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x5f, 0x66, 0x6f, 0x72, 0x75, 0x6d,
	0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x5f, 0x75, 0x6e, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x42,
	0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x67, 0x72, 0x61, 0x6d, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_forum_topic_proto_rawDescOnce sync.Once
	file_forum_topic_proto_rawDescData []byte
)

func file_forum_topic_proto_rawDescGZIP() []byte {
	file_forum_topic_proto_rawDescOnce.Do(func() {
		file_forum_topic_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_forum_topic_proto_rawDesc), len(file_forum_topic_proto_rawDesc)))
	})
	return file_forum_topic_proto_rawDescData
}

var file_forum_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_forum_topic_proto_goTypes = []any{
	(*ForumTopic)(nil),                // 0: forum_topic
	(*ForumTopicCreated)(nil),         // 1: forum_topic_created
	(*ForumTopicEdited)(nil),          // 2: forum_topic_edited
	(*ForumTopicClosed)(nil),          // 3: forum_topic_closed
	(*ForumTopicReopened)(nil),        // 4: forum_topic_reopened
	(*GeneralForumTopicHidden)(nil),   // 5: general_forum_topic_hidden
	(*GeneralForumTopicUnhidden)(nil), // 6: general_forum_topic_unhidden
}
var file_forum_topic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_forum_topic_proto_init() }
func file_forum_topic_proto_init() {
	if File_forum_topic_proto != nil {
		return
	}
	file_forum_topic_proto_msgTypes[0].OneofWrappers = []any{}
	file_forum_topic_proto_msgTypes[1].OneofWrappers = []any{}
	file_forum_topic_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_forum_topic_proto_rawDesc), len(file_forum_topic_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_forum_topic_proto_goTypes,
		DependencyIndexes: file_forum_topic_proto_depIdxs,
		MessageInfos:      file_forum_topic_proto_msgTypes,
	}.Build()
	File_forum_topic_proto = out.File
	file_forum_topic_proto_goTypes = nil
	file_forum_topic_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: instamadilloCoreTypeText/InstamadilloCoreTypeText.proto

package instamadilloCoreTypeText

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	instamadilloCoreTypeMedia "github.com/tuusuario/whatsmeow-backend/proto/instamadilloCoreTypeMedia"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Text_FormatStyle int32

const (
	Text_TEXT_FORMAT_STYLE_UNSPECIFIED   Text_FormatStyle = 0
	Text_TEXT_FORMAT_STYLE_BOLD          Text_FormatStyle = 1
	Text_TEXT_FORMAT_STYLE_ITALIC        Text_FormatStyle = 2
	Text_TEXT_FORMAT_STYLE_STRIKETHROUGH Text_FormatStyle = 3
	Text_TEXT_FORMAT_STYLE_UNDERLINE     Text_FormatStyle = 4
	Text_TEXT_FORMAT_STYLE_INVALID       Text_FormatStyle = 5
)

// Enum value maps for Text_FormatStyle.
var (
	Text_FormatStyle_name = map[int32]string{
		0: "TEXT_FORMAT_STYLE_UNSPECIFIED",
		1: "TEXT_FORMAT_STYLE_BOLD",
		2: "TEXT_FORMAT_STYLE_ITALIC",
		3: "TEXT_FORMAT_STYLE_STRIKETHROUGH",
		4: "TEXT_FORMAT_STYLE_UNDERLINE",
		5: "TEXT_FORMAT_STYLE_INVALID",
	}
	Text_FormatStyle_value = map[string]int32{
		"TEXT_FORMAT_STYLE_UNSPECIFIED":   0,
		"TEXT_FORMAT_STYLE_BOLD":          1,
		"TEXT_FORMAT_STYLE_ITALIC":        2,
		"TEXT_FORMAT_STYLE_STRIKETHROUGH": 3,
		"TEXT_FORMAT_STYLE_UNDERLINE":     4,
		"TEXT_FORMAT_STYLE_INVALID":       5,
	}
)

func (x Text_FormatStyle) Enum() *Text_FormatStyle {
	p := new(Text_FormatStyle)
	*p = x
	return p
}

func (x Text_FormatStyle) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Text_FormatStyle) Descriptor() protoreflect.EnumDescriptor {
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_enumTypes[0].Descriptor()
}

func (Text_FormatStyle) Type() protoreflect.EnumType {
	return &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_enumTypes[0]
}

func (x Text_FormatStyle) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Text_FormatStyle) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Text_FormatStyle(num)
	return nil
}

// Deprecated: Use Text_FormatStyle.Descriptor instead.
func (Text_FormatStyle) EnumDescriptor() ([]byte, []int) {
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescGZIP(), []int{0, 0}
}

type Text struct {
	state                        protoimpl.MessageState         `protogen:"open.v1"`
	Text                         *string                        `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	IsSuggestedReply             *bool                          `protobuf:"varint,2,opt,name=isSuggestedReply" json:"isSuggestedReply,omitempty"`
	PostbackPayload              *string                        `protobuf:"bytes,3,opt,name=postbackPayload" json:"postbackPayload,omitempty"`
	PowerUpData                  *PowerUpsData                  `protobuf:"bytes,4,opt,name=powerUpData" json:"powerUpData,omitempty"`
	Commands                     []*CommandRangeData            `protobuf:"bytes,5,rep,name=commands" json:"commands,omitempty"`
	AnimatedEmojiCharacterRanges []*AnimatedEmojiCharacterRange `protobuf:"bytes,6,rep,name=animatedEmojiCharacterRanges" json:"animatedEmojiCharacterRanges,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *Text) Reset() {
	*x = Text{}
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Text) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Text) ProtoMessage() {}

func (x *Text) ProtoReflect() protoreflect.Message {
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Text.ProtoReflect.Descriptor instead.
func (*Text) Descriptor() ([]byte, []int) {
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescGZIP(), []int{0}
}

func (x *Text) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *Text) GetIsSuggestedReply() bool {
	if x != nil && x.IsSuggestedReply != nil {
		return *x.IsSuggestedReply
	}
	return false
}

func (x *Text) GetPostbackPayload() string {
	if x != nil && x.PostbackPayload != nil {
		return *x.PostbackPayload
	}
	return ""
}

func (x *Text) GetPowerUpData() *PowerUpsData {
	if x != nil {
		return x.PowerUpData
	}
	return nil
}

func (x *Text) GetCommands() []*CommandRangeData {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *Text) GetAnimatedEmojiCharacterRanges() []*AnimatedEmojiCharacterRange {
	if x != nil {
		return x.AnimatedEmojiCharacterRanges
	}
	return nil
}

type PowerUpsData struct {
	state           protoimpl.MessageState                          `protogen:"open.v1"`
	Style           *int32                                          `protobuf:"varint,1,opt,name=style" json:"style,omitempty"`
	MediaAttachment *instamadilloCoreTypeMedia.CommonMediaTransport `protobuf:"bytes,2,opt,name=mediaAttachment" json:"mediaAttachment,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PowerUpsData) Reset() {
	*x = PowerUpsData{}
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PowerUpsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerUpsData) ProtoMessage() {}

func (x *PowerUpsData) ProtoReflect() protoreflect.Message {
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerUpsData.ProtoReflect.Descriptor instead.
func (*PowerUpsData) Descriptor() ([]byte, []int) {
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescGZIP(), []int{1}
}

func (x *PowerUpsData) GetStyle() int32 {
	if x != nil && x.Style != nil {
		return *x.Style
	}
	return 0
}

func (x *PowerUpsData) GetMediaAttachment() *instamadilloCoreTypeMedia.CommonMediaTransport {
	if x != nil {
		return x.MediaAttachment
	}
	return nil
}

type CommandRangeData struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Offset           *int32                 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Length           *int32                 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Type             *int32                 `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	FBID             *string                `protobuf:"bytes,4,opt,name=FBID" json:"FBID,omitempty"`
	UserOrThreadFbid *string                `protobuf:"bytes,5,opt,name=userOrThreadFbid" json:"userOrThreadFbid,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CommandRangeData) Reset() {
	*x = CommandRangeData{}
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandRangeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRangeData) ProtoMessage() {}

func (x *CommandRangeData) ProtoReflect() protoreflect.Message {
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRangeData.ProtoReflect.Descriptor instead.
func (*CommandRangeData) Descriptor() ([]byte, []int) {
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescGZIP(), []int{2}
}

func (x *CommandRangeData) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *CommandRangeData) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *CommandRangeData) GetType() int32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *CommandRangeData) GetFBID() string {
	if x != nil && x.FBID != nil {
		return *x.FBID
	}
	return ""
}

func (x *CommandRangeData) GetUserOrThreadFbid() string {
	if x != nil && x.UserOrThreadFbid != nil {
		return *x.UserOrThreadFbid
	}
	return ""
}

type FormattedText struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Offset        *int32                 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Length        *int32                 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	Style         *Text_FormatStyle      `protobuf:"varint,3,opt,name=style,enum=InstamadilloCoreTypeText.Text_FormatStyle" json:"style,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FormattedText) Reset() {
	*x = FormattedText{}
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FormattedText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormattedText) ProtoMessage() {}

func (x *FormattedText) ProtoReflect() protoreflect.Message {
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormattedText.ProtoReflect.Descriptor instead.
func (*FormattedText) Descriptor() ([]byte, []int) {
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescGZIP(), []int{3}
}

func (x *FormattedText) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *FormattedText) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

func (x *FormattedText) GetStyle() Text_FormatStyle {
	if x != nil && x.Style != nil {
		return *x.Style
	}
	return Text_TEXT_FORMAT_STYLE_UNSPECIFIED
}

type AnimatedEmojiCharacterRange struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Offset        *int32                 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Length        *int32                 `protobuf:"varint,2,opt,name=length" json:"length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AnimatedEmojiCharacterRange) Reset() {
	*x = AnimatedEmojiCharacterRange{}
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnimatedEmojiCharacterRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimatedEmojiCharacterRange) ProtoMessage() {}

func (x *AnimatedEmojiCharacterRange) ProtoReflect() protoreflect.Message {
	mi := &file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimatedEmojiCharacterRange.ProtoReflect.Descriptor instead.
func (*AnimatedEmojiCharacterRange) Descriptor() ([]byte, []int) {
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescGZIP(), []int{4}
}

func (x *AnimatedEmojiCharacterRange) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *AnimatedEmojiCharacterRange) GetLength() int32 {
	if x != nil && x.Length != nil {
		return *x.Length
	}
	return 0
}

var File_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto protoreflect.FileDescriptor

const file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDesc = "" +
	"\n" +
	"7instamadilloCoreTypeText/InstamadilloCoreTypeText.proto\x12\x18InstamadilloCoreTypeText\x1a9instamadilloCoreTypeMedia/InstamadilloCoreTypeMedia.proto\"\xcf\x04\n" +
	"\x04Text\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\x12*\n" +
	"\x10isSuggestedReply\x18\x02 \x01(\bR\x10isSuggestedReply\x12(\n" +
	"\x0fpostbackPayload\x18\x03 \x01(\tR\x0fpostbackPayload\x12H\n" +
	"\vpowerUpData\x18\x04 \x01(\v2&.InstamadilloCoreTypeText.PowerUpsDataR\vpowerUpData\x12F\n" +
	"\bcommands\x18\x05 \x03(\v2*.InstamadilloCoreTypeText.CommandRangeDataR\bcommands\x12y\n" +
	"\x1canimatedEmojiCharacterRanges\x18\x06 \x03(\v25.InstamadilloCoreTypeText.AnimatedEmojiCharacterRangeR\x1canimatedEmojiCharacterRanges\"\xcf\x01\n" +
	"\vFormatStyle\x12!\n" +
	"\x1dTEXT_FORMAT_STYLE_UNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x16TEXT_FORMAT_STYLE_BOLD\x10\x01\x12\x1c\n" +
	"\x18TEXT_FORMAT_STYLE_ITALIC\x10\x02\x12#\n" +
	"\x1fTEXT_FORMAT_STYLE_STRIKETHROUGH\x10\x03\x12\x1f\n" +
	"\x1bTEXT_FORMAT_STYLE_UNDERLINE\x10\x04\x12\x1d\n" +
	"\x19TEXT_FORMAT_STYLE_INVALID\x10\x05\"\x7f\n" +
	"\fPowerUpsData\x12\x14\n" +
	"\x05style\x18\x01 \x01(\x05R\x05style\x12Y\n" +
	"\x0fmediaAttachment\x18\x02 \x01(\v2/.InstamadilloCoreTypeMedia.CommonMediaTransportR\x0fmediaAttachment\"\x96\x01\n" +
	"\x10CommandRangeData\x12\x16\n" +
	"\x06offset\x18\x01 \x01(\x05R\x06offset\x12\x16\n" +
	"\x06length\x18\x02 \x01(\x05R\x06length\x12\x12\n" +
	"\x04type\x18\x03 \x01(\x05R\x04type\x12\x12\n" +
	"\x04FBID\x18\x04 \x01(\tR\x04FBID\x12*\n" +
	"\x10userOrThreadFbid\x18\x05 \x01(\tR\x10userOrThreadFbid\"\x81\x01\n" +
	"\rFormattedText\x12\x16\n" +
	"\x06offset\x18\x01 \x01(\x05R\x06offset\x12\x16\n" +
	"\x06length\x18\x02 \x01(\x05R\x06length\x12@\n" +
	"\x05style\x18\x03 \x01(\x0e2*.InstamadilloCoreTypeText.Text.FormatStyleR\x05style\"M\n" +
	"\x1bAnimatedEmojiCharacterRange\x12\x16\n" +
	"\x06offset\x18\x01 \x01(\x05R\x06offset\x12\x16\n" +
	"\x06length\x18\x02 \x01(\x05R\x06lengthB4Z2github.com/tuusuario/whatsmeow-backend/proto/instamadilloCoreTypeText"

var (
	file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescOnce sync.Once
	file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescData []byte
)

func file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescGZIP() []byte {
	file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescOnce.Do(func() {
		file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDesc), len(file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDesc)))
	})
	return file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDescData
}

var file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_goTypes = []any{
	(Text_FormatStyle)(0),               // 0: InstamadilloCoreTypeText.Text.FormatStyle
	(*Text)(nil),                        // 1: InstamadilloCoreTypeText.Text
	(*PowerUpsData)(nil),                // 2: InstamadilloCoreTypeText.PowerUpsData
	(*CommandRangeData)(nil),            // 3: InstamadilloCoreTypeText.CommandRangeData
	(*FormattedText)(nil),               // 4: InstamadilloCoreTypeText.FormattedText
	(*AnimatedEmojiCharacterRange)(nil), // 5: InstamadilloCoreTypeText.AnimatedEmojiCharacterRange
	(*instamadilloCoreTypeMedia.CommonMediaTransport)(nil), // 6: InstamadilloCoreTypeMedia.CommonMediaTransport
}
var file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_depIdxs = []int32{
	2, // 0: InstamadilloCoreTypeText.Text.powerUpData:type_name -> InstamadilloCoreTypeText.PowerUpsData
	3, // 1: InstamadilloCoreTypeText.Text.commands:type_name -> InstamadilloCoreTypeText.CommandRangeData
	5, // 2: InstamadilloCoreTypeText.Text.animatedEmojiCharacterRanges:type_name -> InstamadilloCoreTypeText.AnimatedEmojiCharacterRange
	6, // 3: InstamadilloCoreTypeText.PowerUpsData.mediaAttachment:type_name -> InstamadilloCoreTypeMedia.CommonMediaTransport
	0, // 4: InstamadilloCoreTypeText.FormattedText.style:type_name -> InstamadilloCoreTypeText.Text.FormatStyle
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_init() }
func file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_init() {
	if File_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDesc), len(file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_goTypes,
		DependencyIndexes: file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_depIdxs,
		EnumInfos:         file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_enumTypes,
		MessageInfos:      file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_msgTypes,
	}.Build()
	File_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto = out.File
	file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_goTypes = nil
	file_instamadilloCoreTypeText_InstamadilloCoreTypeText_proto_depIdxs = nil
}

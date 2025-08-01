// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: instamadilloCoreTypeCollection/InstamadilloCoreTypeCollection.proto

package instamadilloCoreTypeCollection

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

type Collection struct {
	state         protoimpl.MessageState             `protogen:"open.v1"`
	Name          *string                            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Media         []*instamadilloCoreTypeMedia.Media `protobuf:"bytes,2,rep,name=media" json:"media,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Collection) Reset() {
	*x = Collection{}
	mi := &file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDescGZIP(), []int{0}
}

func (x *Collection) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Collection) GetMedia() []*instamadilloCoreTypeMedia.Media {
	if x != nil {
		return x.Media
	}
	return nil
}

var File_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto protoreflect.FileDescriptor

const file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDesc = "" +
	"\n" +
	"CinstamadilloCoreTypeCollection/InstamadilloCoreTypeCollection.proto\x12\x1eInstamadilloCoreTypeCollection\x1a9instamadilloCoreTypeMedia/InstamadilloCoreTypeMedia.proto\"X\n" +
	"\n" +
	"Collection\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x126\n" +
	"\x05media\x18\x02 \x03(\v2 .InstamadilloCoreTypeMedia.MediaR\x05mediaB:Z8github.com/tuusuario/whatsmeow-backend/proto/instamadilloCoreTypeCollection"

var (
	file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDescOnce sync.Once
	file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDescData []byte
)

func file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDescGZIP() []byte {
	file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDescOnce.Do(func() {
		file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDesc), len(file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDesc)))
	})
	return file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDescData
}

var file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_goTypes = []any{
	(*Collection)(nil),                      // 0: InstamadilloCoreTypeCollection.Collection
	(*instamadilloCoreTypeMedia.Media)(nil), // 1: InstamadilloCoreTypeMedia.Media
}
var file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_depIdxs = []int32{
	1, // 0: InstamadilloCoreTypeCollection.Collection.media:type_name -> InstamadilloCoreTypeMedia.Media
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_init() }
func file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_init() {
	if File_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDesc), len(file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_goTypes,
		DependencyIndexes: file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_depIdxs,
		MessageInfos:      file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_msgTypes,
	}.Build()
	File_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto = out.File
	file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_goTypes = nil
	file_instamadilloCoreTypeCollection_InstamadilloCoreTypeCollection_proto_depIdxs = nil
}

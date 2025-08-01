// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: waDeviceCapabilities/WAProtobufsDeviceCapabilities.proto

package waDeviceCapabilities

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeviceCapabilities_ChatLockSupportLevel int32

const (
	DeviceCapabilities_NONE    DeviceCapabilities_ChatLockSupportLevel = 0
	DeviceCapabilities_MINIMAL DeviceCapabilities_ChatLockSupportLevel = 1
	DeviceCapabilities_FULL    DeviceCapabilities_ChatLockSupportLevel = 2
)

// Enum value maps for DeviceCapabilities_ChatLockSupportLevel.
var (
	DeviceCapabilities_ChatLockSupportLevel_name = map[int32]string{
		0: "NONE",
		1: "MINIMAL",
		2: "FULL",
	}
	DeviceCapabilities_ChatLockSupportLevel_value = map[string]int32{
		"NONE":    0,
		"MINIMAL": 1,
		"FULL":    2,
	}
)

func (x DeviceCapabilities_ChatLockSupportLevel) Enum() *DeviceCapabilities_ChatLockSupportLevel {
	p := new(DeviceCapabilities_ChatLockSupportLevel)
	*p = x
	return p
}

func (x DeviceCapabilities_ChatLockSupportLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceCapabilities_ChatLockSupportLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_enumTypes[0].Descriptor()
}

func (DeviceCapabilities_ChatLockSupportLevel) Type() protoreflect.EnumType {
	return &file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_enumTypes[0]
}

func (x DeviceCapabilities_ChatLockSupportLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DeviceCapabilities_ChatLockSupportLevel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DeviceCapabilities_ChatLockSupportLevel(num)
	return nil
}

// Deprecated: Use DeviceCapabilities_ChatLockSupportLevel.Descriptor instead.
func (DeviceCapabilities_ChatLockSupportLevel) EnumDescriptor() ([]byte, []int) {
	return file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescGZIP(), []int{0, 0}
}

type DeviceCapabilities struct {
	state                protoimpl.MessageState                   `protogen:"open.v1"`
	ChatLockSupportLevel *DeviceCapabilities_ChatLockSupportLevel `protobuf:"varint,1,opt,name=chatLockSupportLevel,enum=WAProtobufsDeviceCapabilities.DeviceCapabilities_ChatLockSupportLevel" json:"chatLockSupportLevel,omitempty"`
	LidMigration         *DeviceCapabilities_LIDMigration         `protobuf:"bytes,2,opt,name=lidMigration" json:"lidMigration,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *DeviceCapabilities) Reset() {
	*x = DeviceCapabilities{}
	mi := &file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCapabilities) ProtoMessage() {}

func (x *DeviceCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCapabilities.ProtoReflect.Descriptor instead.
func (*DeviceCapabilities) Descriptor() ([]byte, []int) {
	return file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceCapabilities) GetChatLockSupportLevel() DeviceCapabilities_ChatLockSupportLevel {
	if x != nil && x.ChatLockSupportLevel != nil {
		return *x.ChatLockSupportLevel
	}
	return DeviceCapabilities_NONE
}

func (x *DeviceCapabilities) GetLidMigration() *DeviceCapabilities_LIDMigration {
	if x != nil {
		return x.LidMigration
	}
	return nil
}

type DeviceCapabilities_LIDMigration struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	ChatDbMigrationTimestamp *uint64                `protobuf:"varint,1,opt,name=chatDbMigrationTimestamp" json:"chatDbMigrationTimestamp,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *DeviceCapabilities_LIDMigration) Reset() {
	*x = DeviceCapabilities_LIDMigration{}
	mi := &file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceCapabilities_LIDMigration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceCapabilities_LIDMigration) ProtoMessage() {}

func (x *DeviceCapabilities_LIDMigration) ProtoReflect() protoreflect.Message {
	mi := &file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceCapabilities_LIDMigration.ProtoReflect.Descriptor instead.
func (*DeviceCapabilities_LIDMigration) Descriptor() ([]byte, []int) {
	return file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DeviceCapabilities_LIDMigration) GetChatDbMigrationTimestamp() uint64 {
	if x != nil && x.ChatDbMigrationTimestamp != nil {
		return *x.ChatDbMigrationTimestamp
	}
	return 0
}

var File_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto protoreflect.FileDescriptor

const file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDesc = "" +
	"\n" +
	"8waDeviceCapabilities/WAProtobufsDeviceCapabilities.proto\x12\x1dWAProtobufsDeviceCapabilities\"\xf9\x02\n" +
	"\x12DeviceCapabilities\x12z\n" +
	"\x14chatLockSupportLevel\x18\x01 \x01(\x0e2F.WAProtobufsDeviceCapabilities.DeviceCapabilities.ChatLockSupportLevelR\x14chatLockSupportLevel\x12b\n" +
	"\flidMigration\x18\x02 \x01(\v2>.WAProtobufsDeviceCapabilities.DeviceCapabilities.LIDMigrationR\flidMigration\x1aJ\n" +
	"\fLIDMigration\x12:\n" +
	"\x18chatDbMigrationTimestamp\x18\x01 \x01(\x04R\x18chatDbMigrationTimestamp\"7\n" +
	"\x14ChatLockSupportLevel\x12\b\n" +
	"\x04NONE\x10\x00\x12\v\n" +
	"\aMINIMAL\x10\x01\x12\b\n" +
	"\x04FULL\x10\x02B0Z.github.com/tuusuario/whatsmeow-backend/proto/waDeviceCapabilities"

var (
	file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescOnce sync.Once
	file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescData []byte
)

func file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescGZIP() []byte {
	file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescOnce.Do(func() {
		file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDesc), len(file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDesc)))
	})
	return file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDescData
}

var file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_goTypes = []any{
	(DeviceCapabilities_ChatLockSupportLevel)(0), // 0: WAProtobufsDeviceCapabilities.DeviceCapabilities.ChatLockSupportLevel
	(*DeviceCapabilities)(nil),                   // 1: WAProtobufsDeviceCapabilities.DeviceCapabilities
	(*DeviceCapabilities_LIDMigration)(nil),      // 2: WAProtobufsDeviceCapabilities.DeviceCapabilities.LIDMigration
}
var file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_depIdxs = []int32{
	0, // 0: WAProtobufsDeviceCapabilities.DeviceCapabilities.chatLockSupportLevel:type_name -> WAProtobufsDeviceCapabilities.DeviceCapabilities.ChatLockSupportLevel
	2, // 1: WAProtobufsDeviceCapabilities.DeviceCapabilities.lidMigration:type_name -> WAProtobufsDeviceCapabilities.DeviceCapabilities.LIDMigration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_init() }
func file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_init() {
	if File_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDesc), len(file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_goTypes,
		DependencyIndexes: file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_depIdxs,
		EnumInfos:         file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_enumTypes,
		MessageInfos:      file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_msgTypes,
	}.Build()
	File_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto = out.File
	file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_goTypes = nil
	file_waDeviceCapabilities_WAProtobufsDeviceCapabilities_proto_depIdxs = nil
}

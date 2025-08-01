// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: waCert/WACert.proto

package waCert

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

type NoiseCertificate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Details       []byte                 `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	Signature     []byte                 `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoiseCertificate) Reset() {
	*x = NoiseCertificate{}
	mi := &file_waCert_WACert_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoiseCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoiseCertificate) ProtoMessage() {}

func (x *NoiseCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_waCert_WACert_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoiseCertificate.ProtoReflect.Descriptor instead.
func (*NoiseCertificate) Descriptor() ([]byte, []int) {
	return file_waCert_WACert_proto_rawDescGZIP(), []int{0}
}

func (x *NoiseCertificate) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *NoiseCertificate) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type CertChain struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Leaf          *CertChain_NoiseCertificate `protobuf:"bytes,1,opt,name=leaf" json:"leaf,omitempty"`
	Intermediate  *CertChain_NoiseCertificate `protobuf:"bytes,2,opt,name=intermediate" json:"intermediate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CertChain) Reset() {
	*x = CertChain{}
	mi := &file_waCert_WACert_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CertChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertChain) ProtoMessage() {}

func (x *CertChain) ProtoReflect() protoreflect.Message {
	mi := &file_waCert_WACert_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertChain.ProtoReflect.Descriptor instead.
func (*CertChain) Descriptor() ([]byte, []int) {
	return file_waCert_WACert_proto_rawDescGZIP(), []int{1}
}

func (x *CertChain) GetLeaf() *CertChain_NoiseCertificate {
	if x != nil {
		return x.Leaf
	}
	return nil
}

func (x *CertChain) GetIntermediate() *CertChain_NoiseCertificate {
	if x != nil {
		return x.Intermediate
	}
	return nil
}

type NoiseCertificate_Details struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Serial        *uint32                `protobuf:"varint,1,opt,name=serial" json:"serial,omitempty"`
	Issuer        *string                `protobuf:"bytes,2,opt,name=issuer" json:"issuer,omitempty"`
	Expires       *uint64                `protobuf:"varint,3,opt,name=expires" json:"expires,omitempty"`
	Subject       *string                `protobuf:"bytes,4,opt,name=subject" json:"subject,omitempty"`
	Key           []byte                 `protobuf:"bytes,5,opt,name=key" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoiseCertificate_Details) Reset() {
	*x = NoiseCertificate_Details{}
	mi := &file_waCert_WACert_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoiseCertificate_Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoiseCertificate_Details) ProtoMessage() {}

func (x *NoiseCertificate_Details) ProtoReflect() protoreflect.Message {
	mi := &file_waCert_WACert_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoiseCertificate_Details.ProtoReflect.Descriptor instead.
func (*NoiseCertificate_Details) Descriptor() ([]byte, []int) {
	return file_waCert_WACert_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NoiseCertificate_Details) GetSerial() uint32 {
	if x != nil && x.Serial != nil {
		return *x.Serial
	}
	return 0
}

func (x *NoiseCertificate_Details) GetIssuer() string {
	if x != nil && x.Issuer != nil {
		return *x.Issuer
	}
	return ""
}

func (x *NoiseCertificate_Details) GetExpires() uint64 {
	if x != nil && x.Expires != nil {
		return *x.Expires
	}
	return 0
}

func (x *NoiseCertificate_Details) GetSubject() string {
	if x != nil && x.Subject != nil {
		return *x.Subject
	}
	return ""
}

func (x *NoiseCertificate_Details) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type CertChain_NoiseCertificate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Details       []byte                 `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	Signature     []byte                 `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CertChain_NoiseCertificate) Reset() {
	*x = CertChain_NoiseCertificate{}
	mi := &file_waCert_WACert_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CertChain_NoiseCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertChain_NoiseCertificate) ProtoMessage() {}

func (x *CertChain_NoiseCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_waCert_WACert_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertChain_NoiseCertificate.ProtoReflect.Descriptor instead.
func (*CertChain_NoiseCertificate) Descriptor() ([]byte, []int) {
	return file_waCert_WACert_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CertChain_NoiseCertificate) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *CertChain_NoiseCertificate) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type CertChain_NoiseCertificate_Details struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Serial        *uint32                `protobuf:"varint,1,opt,name=serial" json:"serial,omitempty"`
	IssuerSerial  *uint32                `protobuf:"varint,2,opt,name=issuerSerial" json:"issuerSerial,omitempty"`
	Key           []byte                 `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	NotBefore     *uint64                `protobuf:"varint,4,opt,name=notBefore" json:"notBefore,omitempty"`
	NotAfter      *uint64                `protobuf:"varint,5,opt,name=notAfter" json:"notAfter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CertChain_NoiseCertificate_Details) Reset() {
	*x = CertChain_NoiseCertificate_Details{}
	mi := &file_waCert_WACert_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CertChain_NoiseCertificate_Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertChain_NoiseCertificate_Details) ProtoMessage() {}

func (x *CertChain_NoiseCertificate_Details) ProtoReflect() protoreflect.Message {
	mi := &file_waCert_WACert_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertChain_NoiseCertificate_Details.ProtoReflect.Descriptor instead.
func (*CertChain_NoiseCertificate_Details) Descriptor() ([]byte, []int) {
	return file_waCert_WACert_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *CertChain_NoiseCertificate_Details) GetSerial() uint32 {
	if x != nil && x.Serial != nil {
		return *x.Serial
	}
	return 0
}

func (x *CertChain_NoiseCertificate_Details) GetIssuerSerial() uint32 {
	if x != nil && x.IssuerSerial != nil {
		return *x.IssuerSerial
	}
	return 0
}

func (x *CertChain_NoiseCertificate_Details) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *CertChain_NoiseCertificate_Details) GetNotBefore() uint64 {
	if x != nil && x.NotBefore != nil {
		return *x.NotBefore
	}
	return 0
}

func (x *CertChain_NoiseCertificate_Details) GetNotAfter() uint64 {
	if x != nil && x.NotAfter != nil {
		return *x.NotAfter
	}
	return 0
}

var File_waCert_WACert_proto protoreflect.FileDescriptor

const file_waCert_WACert_proto_rawDesc = "" +
	"\n" +
	"\x13waCert/WACert.proto\x12\x06WACert\"\xcb\x01\n" +
	"\x10NoiseCertificate\x12\x18\n" +
	"\adetails\x18\x01 \x01(\fR\adetails\x12\x1c\n" +
	"\tsignature\x18\x02 \x01(\fR\tsignature\x1a\x7f\n" +
	"\aDetails\x12\x16\n" +
	"\x06serial\x18\x01 \x01(\rR\x06serial\x12\x16\n" +
	"\x06issuer\x18\x02 \x01(\tR\x06issuer\x12\x18\n" +
	"\aexpires\x18\x03 \x01(\x04R\aexpires\x12\x18\n" +
	"\asubject\x18\x04 \x01(\tR\asubject\x12\x10\n" +
	"\x03key\x18\x05 \x01(\fR\x03key\"\xec\x02\n" +
	"\tCertChain\x126\n" +
	"\x04leaf\x18\x01 \x01(\v2\".WACert.CertChain.NoiseCertificateR\x04leaf\x12F\n" +
	"\fintermediate\x18\x02 \x01(\v2\".WACert.CertChain.NoiseCertificateR\fintermediate\x1a\xde\x01\n" +
	"\x10NoiseCertificate\x12\x18\n" +
	"\adetails\x18\x01 \x01(\fR\adetails\x12\x1c\n" +
	"\tsignature\x18\x02 \x01(\fR\tsignature\x1a\x91\x01\n" +
	"\aDetails\x12\x16\n" +
	"\x06serial\x18\x01 \x01(\rR\x06serial\x12\"\n" +
	"\fissuerSerial\x18\x02 \x01(\rR\fissuerSerial\x12\x10\n" +
	"\x03key\x18\x03 \x01(\fR\x03key\x12\x1c\n" +
	"\tnotBefore\x18\x04 \x01(\x04R\tnotBefore\x12\x1a\n" +
	"\bnotAfter\x18\x05 \x01(\x04R\bnotAfterB\"Z github.com/tuusuario/whatsmeow-backend/proto/waCert"

var (
	file_waCert_WACert_proto_rawDescOnce sync.Once
	file_waCert_WACert_proto_rawDescData []byte
)

func file_waCert_WACert_proto_rawDescGZIP() []byte {
	file_waCert_WACert_proto_rawDescOnce.Do(func() {
		file_waCert_WACert_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_waCert_WACert_proto_rawDesc), len(file_waCert_WACert_proto_rawDesc)))
	})
	return file_waCert_WACert_proto_rawDescData
}

var file_waCert_WACert_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_waCert_WACert_proto_goTypes = []any{
	(*NoiseCertificate)(nil),                   // 0: WACert.NoiseCertificate
	(*CertChain)(nil),                          // 1: WACert.CertChain
	(*NoiseCertificate_Details)(nil),           // 2: WACert.NoiseCertificate.Details
	(*CertChain_NoiseCertificate)(nil),         // 3: WACert.CertChain.NoiseCertificate
	(*CertChain_NoiseCertificate_Details)(nil), // 4: WACert.CertChain.NoiseCertificate.Details
}
var file_waCert_WACert_proto_depIdxs = []int32{
	3, // 0: WACert.CertChain.leaf:type_name -> WACert.CertChain.NoiseCertificate
	3, // 1: WACert.CertChain.intermediate:type_name -> WACert.CertChain.NoiseCertificate
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_waCert_WACert_proto_init() }
func file_waCert_WACert_proto_init() {
	if File_waCert_WACert_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_waCert_WACert_proto_rawDesc), len(file_waCert_WACert_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waCert_WACert_proto_goTypes,
		DependencyIndexes: file_waCert_WACert_proto_depIdxs,
		MessageInfos:      file_waCert_WACert_proto_msgTypes,
	}.Build()
	File_waCert_WACert_proto = out.File
	file_waCert_WACert_proto_goTypes = nil
	file_waCert_WACert_proto_depIdxs = nil
}

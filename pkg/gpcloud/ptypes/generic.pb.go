// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: generic.proto

package ptypes

import (
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

type ComputeResourceType int32

const (
	ComputeResourceType_BARE_METAL ComputeResourceType = 0
	ComputeResourceType_VIRTUAL    ComputeResourceType = 1
)

// Enum value maps for ComputeResourceType.
var (
	ComputeResourceType_name = map[int32]string{
		0: "BARE_METAL",
		1: "VIRTUAL",
	}
	ComputeResourceType_value = map[string]int32{
		"BARE_METAL": 0,
		"VIRTUAL":    1,
	}
)

func (x ComputeResourceType) Enum() *ComputeResourceType {
	p := new(ComputeResourceType)
	*p = x
	return p
}

func (x ComputeResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComputeResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_generic_proto_enumTypes[0].Descriptor()
}

func (ComputeResourceType) Type() protoreflect.EnumType {
	return &file_generic_proto_enumTypes[0]
}

func (x ComputeResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComputeResourceType.Descriptor instead.
func (ComputeResourceType) EnumDescriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{0}
}

type Platform int32

const (
	Platform_UNKNOWN Platform = 0
	Platform_WINDOWS Platform = 1
	Platform_MAC_OS  Platform = 2
	Platform_LINUX   Platform = 3
	Platform_ANDROID Platform = 4
	Platform_IOS     Platform = 5
)

// Enum value maps for Platform.
var (
	Platform_name = map[int32]string{
		0: "UNKNOWN",
		1: "WINDOWS",
		2: "MAC_OS",
		3: "LINUX",
		4: "ANDROID",
		5: "IOS",
	}
	Platform_value = map[string]int32{
		"UNKNOWN": 0,
		"WINDOWS": 1,
		"MAC_OS":  2,
		"LINUX":   3,
		"ANDROID": 4,
		"IOS":     5,
	}
)

func (x Platform) Enum() *Platform {
	p := new(Platform)
	*p = x
	return p
}

func (x Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_generic_proto_enumTypes[1].Descriptor()
}

func (Platform) Type() protoreflect.EnumType {
	return &file_generic_proto_enumTypes[1]
}

func (x Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Platform.Descriptor instead.
func (Platform) EnumDescriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{1}
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes    []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *File) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int64 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// ISO 4217 code
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	// Formatted string of price
	Formatted string `protobuf:"bytes,3,opt,name=formatted,proto3" json:"formatted,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{1}
}

func (x *Price) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Price) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Price) GetFormatted() string {
	if x != nil {
		return x.Formatted
	}
	return ""
}

type EmptyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResponse) Reset() {
	*x = EmptyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResponse) ProtoMessage() {}

func (x *EmptyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResponse.ProtoReflect.Descriptor instead.
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{2}
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{3}
}

var File_generic_proto protoreflect.FileDescriptor

var file_generic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x39, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x59, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2a, 0x32, 0x0a, 0x13, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x42, 0x41, 0x52, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, 0x4c, 0x10, 0x01, 0x2a,
	0x51, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x43, 0x5f, 0x4f, 0x53, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53,
	0x10, 0x05, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x3b, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_generic_proto_rawDescOnce sync.Once
	file_generic_proto_rawDescData = file_generic_proto_rawDesc
)

func file_generic_proto_rawDescGZIP() []byte {
	file_generic_proto_rawDescOnce.Do(func() {
		file_generic_proto_rawDescData = protoimpl.X.CompressGZIP(file_generic_proto_rawDescData)
	})
	return file_generic_proto_rawDescData
}

var file_generic_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_generic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_generic_proto_goTypes = []interface{}{
	(ComputeResourceType)(0), // 0: api.ComputeResourceType
	(Platform)(0),            // 1: api.Platform
	(*File)(nil),             // 2: api.File
	(*Price)(nil),            // 3: api.Price
	(*EmptyResponse)(nil),    // 4: api.EmptyResponse
	(*EmptyRequest)(nil),     // 5: api.EmptyRequest
}
var file_generic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_generic_proto_init() }
func file_generic_proto_init() {
	if File_generic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_generic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_generic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_generic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResponse); i {
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
		file_generic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
			RawDescriptor: file_generic_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_generic_proto_goTypes,
		DependencyIndexes: file_generic_proto_depIdxs,
		EnumInfos:         file_generic_proto_enumTypes,
		MessageInfos:      file_generic_proto_msgTypes,
	}.Build()
	File_generic_proto = out.File
	file_generic_proto_rawDesc = nil
	file_generic_proto_goTypes = nil
	file_generic_proto_depIdxs = nil
}

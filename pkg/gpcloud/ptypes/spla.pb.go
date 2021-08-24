// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: spla.proto

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

type SplaPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base price for the first 8 cores
	BasePrice *Price `protobuf:"bytes,1,opt,name=base_price,json=basePrice,proto3" json:"base_price,omitempty"`
	// Additional price for 2 further CPU cores
	Additional_2CoresPrice *Price `protobuf:"bytes,2,opt,name=additional_2_cores_price,json=additional2CoresPrice,proto3" json:"additional_2_cores_price,omitempty"`
}

func (x *SplaPrice) Reset() {
	*x = SplaPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spla_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SplaPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SplaPrice) ProtoMessage() {}

func (x *SplaPrice) ProtoReflect() protoreflect.Message {
	mi := &file_spla_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SplaPrice.ProtoReflect.Descriptor instead.
func (*SplaPrice) Descriptor() ([]byte, []int) {
	return file_spla_proto_rawDescGZIP(), []int{0}
}

func (x *SplaPrice) GetBasePrice() *Price {
	if x != nil {
		return x.BasePrice
	}
	return nil
}

func (x *SplaPrice) GetAdditional_2CoresPrice() *Price {
	if x != nil {
		return x.Additional_2CoresPrice
	}
	return nil
}

type GetSplaPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *GetSplaPriceRequest) Reset() {
	*x = GetSplaPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spla_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSplaPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSplaPriceRequest) ProtoMessage() {}

func (x *GetSplaPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spla_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSplaPriceRequest.ProtoReflect.Descriptor instead.
func (*GetSplaPriceRequest) Descriptor() ([]byte, []int) {
	return file_spla_proto_rawDescGZIP(), []int{1}
}

func (x *GetSplaPriceRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type GetSplaPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price *SplaPrice `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetSplaPriceResponse) Reset() {
	*x = GetSplaPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spla_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSplaPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSplaPriceResponse) ProtoMessage() {}

func (x *GetSplaPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spla_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSplaPriceResponse.ProtoReflect.Descriptor instead.
func (*GetSplaPriceResponse) Descriptor() ([]byte, []int) {
	return file_spla_proto_rawDescGZIP(), []int{2}
}

func (x *GetSplaPriceResponse) GetPrice() *SplaPrice {
	if x != nil {
		return x.Price
	}
	return nil
}

var File_spla_proto protoreflect.FileDescriptor

var file_spla_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x70, 0x6c, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x70, 0x6c, 0x61, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x09, 0x53, 0x70, 0x6c, 0x61, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a,
	0x18, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x32, 0x5f, 0x63, 0x6f,
	0x72, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x15, 0x61, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x32, 0x43, 0x6f, 0x72, 0x65, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x22, 0x34, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x70, 0x6c, 0x61, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53,
	0x70, 0x6c, 0x61, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x6c, 0x61, 0x2e, 0x53, 0x70, 0x6c, 0x61, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spla_proto_rawDescOnce sync.Once
	file_spla_proto_rawDescData = file_spla_proto_rawDesc
)

func file_spla_proto_rawDescGZIP() []byte {
	file_spla_proto_rawDescOnce.Do(func() {
		file_spla_proto_rawDescData = protoimpl.X.CompressGZIP(file_spla_proto_rawDescData)
	})
	return file_spla_proto_rawDescData
}

var file_spla_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spla_proto_goTypes = []interface{}{
	(*SplaPrice)(nil),            // 0: api.spla.SplaPrice
	(*GetSplaPriceRequest)(nil),  // 1: api.spla.GetSplaPriceRequest
	(*GetSplaPriceResponse)(nil), // 2: api.spla.GetSplaPriceResponse
	(*Price)(nil),                // 3: api.Price
}
var file_spla_proto_depIdxs = []int32{
	3, // 0: api.spla.SplaPrice.base_price:type_name -> api.Price
	3, // 1: api.spla.SplaPrice.additional_2_cores_price:type_name -> api.Price
	0, // 2: api.spla.GetSplaPriceResponse.price:type_name -> api.spla.SplaPrice
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spla_proto_init() }
func file_spla_proto_init() {
	if File_spla_proto != nil {
		return
	}
	file_generic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spla_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SplaPrice); i {
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
		file_spla_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSplaPriceRequest); i {
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
		file_spla_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSplaPriceResponse); i {
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
			RawDescriptor: file_spla_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spla_proto_goTypes,
		DependencyIndexes: file_spla_proto_depIdxs,
		MessageInfos:      file_spla_proto_msgTypes,
	}.Build()
	File_spla_proto = out.File
	file_spla_proto_rawDesc = nil
	file_spla_proto_goTypes = nil
	file_spla_proto_depIdxs = nil
}

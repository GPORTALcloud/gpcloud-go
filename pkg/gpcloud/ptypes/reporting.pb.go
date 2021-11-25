// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: ptypes/reporting.proto

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

type UserReporting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Country code of billing user
	CountryCode string `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Net sum for all bills
	SumNet   *Price `protobuf:"bytes,2,opt,name=sum_net,json=sumNet,proto3" json:"sum_net,omitempty"`
	SumGross *Price `protobuf:"bytes,3,opt,name=sum_gross,json=sumGross,proto3" json:"sum_gross,omitempty"`
}

func (x *UserReporting) Reset() {
	*x = UserReporting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_reporting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserReporting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserReporting) ProtoMessage() {}

func (x *UserReporting) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_reporting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserReporting.ProtoReflect.Descriptor instead.
func (*UserReporting) Descriptor() ([]byte, []int) {
	return file_ptypes_reporting_proto_rawDescGZIP(), []int{0}
}

func (x *UserReporting) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *UserReporting) GetSumNet() *Price {
	if x != nil {
		return x.SumNet
	}
	return nil
}

func (x *UserReporting) GetSumGross() *Price {
	if x != nil {
		return x.SumGross
	}
	return nil
}

type ServerReportingFlavour struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Server flavour
	Flavour string `protobuf:"bytes,1,opt,name=flavour,proto3" json:"flavour,omitempty"`
	// Amount of rented compute resources
	RentedProducts int64 `protobuf:"varint,2,opt,name=rented_products,json=rentedProducts,proto3" json:"rented_products,omitempty"`
	// Amount of hours rented compute resources
	RentedHoursProducts int64 `protobuf:"varint,3,opt,name=rented_hours_products,json=rentedHoursProducts,proto3" json:"rented_hours_products,omitempty"`
	// Amount of unique servers
	RentedUniqueServers int64 `protobuf:"varint,4,opt,name=rented_unique_servers,json=rentedUniqueServers,proto3" json:"rented_unique_servers,omitempty"`
}

func (x *ServerReportingFlavour) Reset() {
	*x = ServerReportingFlavour{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_reporting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReportingFlavour) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReportingFlavour) ProtoMessage() {}

func (x *ServerReportingFlavour) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_reporting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReportingFlavour.ProtoReflect.Descriptor instead.
func (*ServerReportingFlavour) Descriptor() ([]byte, []int) {
	return file_ptypes_reporting_proto_rawDescGZIP(), []int{1}
}

func (x *ServerReportingFlavour) GetFlavour() string {
	if x != nil {
		return x.Flavour
	}
	return ""
}

func (x *ServerReportingFlavour) GetRentedProducts() int64 {
	if x != nil {
		return x.RentedProducts
	}
	return 0
}

func (x *ServerReportingFlavour) GetRentedHoursProducts() int64 {
	if x != nil {
		return x.RentedHoursProducts
	}
	return 0
}

func (x *ServerReportingFlavour) GetRentedUniqueServers() int64 {
	if x != nil {
		return x.RentedUniqueServers
	}
	return 0
}

type ServerReporting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Datacenter location
	DatacenterShort string                    `protobuf:"bytes,1,opt,name=datacenter_short,json=datacenterShort,proto3" json:"datacenter_short,omitempty"`
	Flavours        []*ServerReportingFlavour `protobuf:"bytes,2,rep,name=flavours,proto3" json:"flavours,omitempty"`
}

func (x *ServerReporting) Reset() {
	*x = ServerReporting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_reporting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReporting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReporting) ProtoMessage() {}

func (x *ServerReporting) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_reporting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReporting.ProtoReflect.Descriptor instead.
func (*ServerReporting) Descriptor() ([]byte, []int) {
	return file_ptypes_reporting_proto_rawDescGZIP(), []int{2}
}

func (x *ServerReporting) GetDatacenterShort() string {
	if x != nil {
		return x.DatacenterShort
	}
	return ""
}

func (x *ServerReporting) GetFlavours() []*ServerReportingFlavour {
	if x != nil {
		return x.Flavours
	}
	return nil
}

type Reporting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User reportings
	UserReportings []*UserReporting `protobuf:"bytes,1,rep,name=user_reportings,json=userReportings,proto3" json:"user_reportings,omitempty"`
	// Server reportings
	ServerReportings []*ServerReporting `protobuf:"bytes,2,rep,name=server_reportings,json=serverReportings,proto3" json:"server_reportings,omitempty"`
}

func (x *Reporting) Reset() {
	*x = Reporting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_reporting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reporting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reporting) ProtoMessage() {}

func (x *Reporting) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_reporting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reporting.ProtoReflect.Descriptor instead.
func (*Reporting) Descriptor() ([]byte, []int) {
	return file_ptypes_reporting_proto_rawDescGZIP(), []int{3}
}

func (x *Reporting) GetUserReportings() []*UserReporting {
	if x != nil {
		return x.UserReportings
	}
	return nil
}

func (x *Reporting) GetServerReportings() []*ServerReporting {
	if x != nil {
		return x.ServerReportings
	}
	return nil
}

var File_ptypes_reporting_proto protoreflect.FileDescriptor

var file_ptypes_reporting_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x14, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x23, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x5f, 0x6e, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52,
	0x06, 0x73, 0x75, 0x6d, 0x4e, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x73, 0x75, 0x6d, 0x5f, 0x67,
	0x72, 0x6f, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x75, 0x6d, 0x47, 0x72, 0x6f, 0x73, 0x73,
	0x22, 0xc3, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x66,
	0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6c,
	0x61, 0x76, 0x6f, 0x75, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x72, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x32,
	0x0a, 0x15, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x72,
	0x65, 0x6e, 0x74, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x13, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x64, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x7f, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x12, 0x41, 0x0a, 0x08, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x52, 0x08, 0x66,
	0x6c, 0x61, 0x76, 0x6f, 0x75, 0x72, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4b, 0x0a, 0x11,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x8e, 0x01, 0x0a, 0x11, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x42,
	0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x52, 0x58, 0xaa, 0x02, 0x0d,
	0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x0d,
	0x41, 0x70, 0x69, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x19,
	0x41, 0x70, 0x69, 0x5c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ptypes_reporting_proto_rawDescOnce sync.Once
	file_ptypes_reporting_proto_rawDescData = file_ptypes_reporting_proto_rawDesc
)

func file_ptypes_reporting_proto_rawDescGZIP() []byte {
	file_ptypes_reporting_proto_rawDescOnce.Do(func() {
		file_ptypes_reporting_proto_rawDescData = protoimpl.X.CompressGZIP(file_ptypes_reporting_proto_rawDescData)
	})
	return file_ptypes_reporting_proto_rawDescData
}

var file_ptypes_reporting_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ptypes_reporting_proto_goTypes = []interface{}{
	(*UserReporting)(nil),          // 0: api.reporting.UserReporting
	(*ServerReportingFlavour)(nil), // 1: api.reporting.ServerReportingFlavour
	(*ServerReporting)(nil),        // 2: api.reporting.ServerReporting
	(*Reporting)(nil),              // 3: api.reporting.Reporting
	(*Price)(nil),                  // 4: api.Price
}
var file_ptypes_reporting_proto_depIdxs = []int32{
	4, // 0: api.reporting.UserReporting.sum_net:type_name -> api.Price
	4, // 1: api.reporting.UserReporting.sum_gross:type_name -> api.Price
	1, // 2: api.reporting.ServerReporting.flavours:type_name -> api.reporting.ServerReportingFlavour
	0, // 3: api.reporting.Reporting.user_reportings:type_name -> api.reporting.UserReporting
	2, // 4: api.reporting.Reporting.server_reportings:type_name -> api.reporting.ServerReporting
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ptypes_reporting_proto_init() }
func file_ptypes_reporting_proto_init() {
	if File_ptypes_reporting_proto != nil {
		return
	}
	file_ptypes_generic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ptypes_reporting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserReporting); i {
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
		file_ptypes_reporting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReportingFlavour); i {
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
		file_ptypes_reporting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerReporting); i {
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
		file_ptypes_reporting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reporting); i {
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
			RawDescriptor: file_ptypes_reporting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ptypes_reporting_proto_goTypes,
		DependencyIndexes: file_ptypes_reporting_proto_depIdxs,
		MessageInfos:      file_ptypes_reporting_proto_msgTypes,
	}.Build()
	File_ptypes_reporting_proto = out.File
	file_ptypes_reporting_proto_rawDesc = nil
	file_ptypes_reporting_proto_goTypes = nil
	file_ptypes_reporting_proto_depIdxs = nil
}
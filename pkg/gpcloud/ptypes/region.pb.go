// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: region.proto

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

type DataCenterStatus int32

const (
	DataCenterStatus_AVAILABLE   DataCenterStatus = 0
	DataCenterStatus_COMING_SOON DataCenterStatus = 1
	DataCenterStatus_DISABLED    DataCenterStatus = 2
)

// Enum value maps for DataCenterStatus.
var (
	DataCenterStatus_name = map[int32]string{
		0: "AVAILABLE",
		1: "COMING_SOON",
		2: "DISABLED",
	}
	DataCenterStatus_value = map[string]int32{
		"AVAILABLE":   0,
		"COMING_SOON": 1,
		"DISABLED":    2,
	}
)

func (x DataCenterStatus) Enum() *DataCenterStatus {
	p := new(DataCenterStatus)
	*p = x
	return p
}

func (x DataCenterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataCenterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_region_proto_enumTypes[0].Descriptor()
}

func (DataCenterStatus) Type() protoreflect.EnumType {
	return &file_region_proto_enumTypes[0]
}

func (x DataCenterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataCenterStatus.Descriptor instead.
func (DataCenterStatus) EnumDescriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{0}
}

type DataCenter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Datacenter ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Datacenter name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Datacenter name
	Short        string `protobuf:"bytes,3,opt,name=short,proto3" json:"short,omitempty"`
	ServerPrefix string `protobuf:"bytes,4,opt,name=server_prefix,json=serverPrefix,proto3" json:"server_prefix,omitempty"`
	// DataCenter Region
	Region *Region `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	// DataCenter status
	Status DataCenterStatus `protobuf:"varint,6,opt,name=status,proto3,enum=api.region.DataCenterStatus" json:"status,omitempty"`
}

func (x *DataCenter) Reset() {
	*x = DataCenter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataCenter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataCenter) ProtoMessage() {}

func (x *DataCenter) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataCenter.ProtoReflect.Descriptor instead.
func (*DataCenter) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{0}
}

func (x *DataCenter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataCenter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataCenter) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *DataCenter) GetServerPrefix() string {
	if x != nil {
		return x.ServerPrefix
	}
	return ""
}

func (x *DataCenter) GetRegion() *Region {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *DataCenter) GetStatus() DataCenterStatus {
	if x != nil {
		return x.Status
	}
	return DataCenterStatus_AVAILABLE
}

// Region Represents a region
type Region struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Region ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Region name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Country code ISO 3166-1 alpha-2
	CountryCode string        `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	Datacenters []*DataCenter `protobuf:"bytes,4,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
}

func (x *Region) Reset() {
	*x = Region{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{1}
}

func (x *Region) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Region) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Region) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Region) GetDatacenters() []*DataCenter {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

type ListDataCenterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datacenters []*DataCenter `protobuf:"bytes,1,rep,name=datacenters,proto3" json:"datacenters,omitempty"`
}

func (x *ListDataCenterResponse) Reset() {
	*x = ListDataCenterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDataCenterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDataCenterResponse) ProtoMessage() {}

func (x *ListDataCenterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDataCenterResponse.ProtoReflect.Descriptor instead.
func (*ListDataCenterResponse) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{2}
}

func (x *ListDataCenterResponse) GetDatacenters() []*DataCenter {
	if x != nil {
		return x.Datacenters
	}
	return nil
}

type AdminListRegionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regions []*Region `protobuf:"bytes,1,rep,name=regions,proto3" json:"regions,omitempty"`
}

func (x *AdminListRegionsResponse) Reset() {
	*x = AdminListRegionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminListRegionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminListRegionsResponse) ProtoMessage() {}

func (x *AdminListRegionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminListRegionsResponse.ProtoReflect.Descriptor instead.
func (*AdminListRegionsResponse) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{3}
}

func (x *AdminListRegionsResponse) GetRegions() []*Region {
	if x != nil {
		return x.Regions
	}
	return nil
}

type AdminCreateRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CountryCode string `protobuf:"bytes,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *AdminCreateRegionRequest) Reset() {
	*x = AdminCreateRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateRegionRequest) ProtoMessage() {}

func (x *AdminCreateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateRegionRequest.ProtoReflect.Descriptor instead.
func (*AdminCreateRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{4}
}

func (x *AdminCreateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminCreateRegionRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type AdminUpdateRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CountryCode string `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *AdminUpdateRegionRequest) Reset() {
	*x = AdminUpdateRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateRegionRequest) ProtoMessage() {}

func (x *AdminUpdateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateRegionRequest.ProtoReflect.Descriptor instead.
func (*AdminUpdateRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{5}
}

func (x *AdminUpdateRegionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdminUpdateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminUpdateRegionRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type AdminAddDatacenterToRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Short        string           `protobuf:"bytes,3,opt,name=short,proto3" json:"short,omitempty"`
	ServerPrefix string           `protobuf:"bytes,4,opt,name=server_prefix,json=serverPrefix,proto3" json:"server_prefix,omitempty"`
	Status       DataCenterStatus `protobuf:"varint,5,opt,name=status,proto3,enum=api.region.DataCenterStatus" json:"status,omitempty"`
}

func (x *AdminAddDatacenterToRegionRequest) Reset() {
	*x = AdminAddDatacenterToRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminAddDatacenterToRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminAddDatacenterToRegionRequest) ProtoMessage() {}

func (x *AdminAddDatacenterToRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminAddDatacenterToRegionRequest.ProtoReflect.Descriptor instead.
func (*AdminAddDatacenterToRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{6}
}

func (x *AdminAddDatacenterToRegionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdminAddDatacenterToRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminAddDatacenterToRegionRequest) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *AdminAddDatacenterToRegionRequest) GetServerPrefix() string {
	if x != nil {
		return x.ServerPrefix
	}
	return ""
}

func (x *AdminAddDatacenterToRegionRequest) GetStatus() DataCenterStatus {
	if x != nil {
		return x.Status
	}
	return DataCenterStatus_AVAILABLE
}

type AdminUpdateDatacenterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Short        string           `protobuf:"bytes,3,opt,name=short,proto3" json:"short,omitempty"`
	ServerPrefix string           `protobuf:"bytes,4,opt,name=server_prefix,json=serverPrefix,proto3" json:"server_prefix,omitempty"`
	Status       DataCenterStatus `protobuf:"varint,5,opt,name=status,proto3,enum=api.region.DataCenterStatus" json:"status,omitempty"`
}

func (x *AdminUpdateDatacenterRequest) Reset() {
	*x = AdminUpdateDatacenterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateDatacenterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateDatacenterRequest) ProtoMessage() {}

func (x *AdminUpdateDatacenterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateDatacenterRequest.ProtoReflect.Descriptor instead.
func (*AdminUpdateDatacenterRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{7}
}

func (x *AdminUpdateDatacenterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AdminUpdateDatacenterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminUpdateDatacenterRequest) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *AdminUpdateDatacenterRequest) GetServerPrefix() string {
	if x != nil {
		return x.ServerPrefix
	}
	return ""
}

func (x *AdminUpdateDatacenterRequest) GetStatus() DataCenterStatus {
	if x != nil {
		return x.Status
	}
	return DataCenterStatus_AVAILABLE
}

type AdminDeleteDatacenterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AdminDeleteDatacenterRequest) Reset() {
	*x = AdminDeleteDatacenterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteDatacenterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteDatacenterRequest) ProtoMessage() {}

func (x *AdminDeleteDatacenterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteDatacenterRequest.ProtoReflect.Descriptor instead.
func (*AdminDeleteDatacenterRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{8}
}

func (x *AdminDeleteDatacenterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AdminDeleteRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AdminDeleteRegionRequest) Reset() {
	*x = AdminDeleteRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteRegionRequest) ProtoMessage() {}

func (x *AdminDeleteRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteRegionRequest.ProtoReflect.Descriptor instead.
func (*AdminDeleteRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{9}
}

func (x *AdminDeleteRegionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AdminGetRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AdminGetRegionRequest) Reset() {
	*x = AdminGetRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetRegionRequest) ProtoMessage() {}

func (x *AdminGetRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetRegionRequest.ProtoReflect.Descriptor instead.
func (*AdminGetRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{10}
}

func (x *AdminGetRegionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_region_proto protoreflect.FileDescriptor

var file_region_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0xcd, 0x01, 0x0a, 0x0a, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2a, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x52, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x48, 0x0a, 0x18, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x51, 0x0a, 0x18, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x61, 0x0a, 0x18, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x21, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x54, 0x6f, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x34, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x1c, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x18, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x2a, 0x40, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x4f, 0x4d, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4f,
	0x4f, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44,
	0x10, 0x02, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x3b, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_region_proto_rawDescOnce sync.Once
	file_region_proto_rawDescData = file_region_proto_rawDesc
)

func file_region_proto_rawDescGZIP() []byte {
	file_region_proto_rawDescOnce.Do(func() {
		file_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_region_proto_rawDescData)
	})
	return file_region_proto_rawDescData
}

var file_region_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_region_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_region_proto_goTypes = []interface{}{
	(DataCenterStatus)(0),                     // 0: api.region.DataCenterStatus
	(*DataCenter)(nil),                        // 1: api.region.DataCenter
	(*Region)(nil),                            // 2: api.region.Region
	(*ListDataCenterResponse)(nil),            // 3: api.region.ListDataCenterResponse
	(*AdminListRegionsResponse)(nil),          // 4: api.region.AdminListRegionsResponse
	(*AdminCreateRegionRequest)(nil),          // 5: api.region.AdminCreateRegionRequest
	(*AdminUpdateRegionRequest)(nil),          // 6: api.region.AdminUpdateRegionRequest
	(*AdminAddDatacenterToRegionRequest)(nil), // 7: api.region.AdminAddDatacenterToRegionRequest
	(*AdminUpdateDatacenterRequest)(nil),      // 8: api.region.AdminUpdateDatacenterRequest
	(*AdminDeleteDatacenterRequest)(nil),      // 9: api.region.AdminDeleteDatacenterRequest
	(*AdminDeleteRegionRequest)(nil),          // 10: api.region.AdminDeleteRegionRequest
	(*AdminGetRegionRequest)(nil),             // 11: api.region.AdminGetRegionRequest
}
var file_region_proto_depIdxs = []int32{
	2, // 0: api.region.DataCenter.region:type_name -> api.region.Region
	0, // 1: api.region.DataCenter.status:type_name -> api.region.DataCenterStatus
	1, // 2: api.region.Region.datacenters:type_name -> api.region.DataCenter
	1, // 3: api.region.ListDataCenterResponse.datacenters:type_name -> api.region.DataCenter
	2, // 4: api.region.AdminListRegionsResponse.regions:type_name -> api.region.Region
	0, // 5: api.region.AdminAddDatacenterToRegionRequest.status:type_name -> api.region.DataCenterStatus
	0, // 6: api.region.AdminUpdateDatacenterRequest.status:type_name -> api.region.DataCenterStatus
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_region_proto_init() }
func file_region_proto_init() {
	if File_region_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_region_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataCenter); i {
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
		file_region_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Region); i {
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
		file_region_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDataCenterResponse); i {
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
		file_region_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminListRegionsResponse); i {
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
		file_region_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateRegionRequest); i {
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
		file_region_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateRegionRequest); i {
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
		file_region_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminAddDatacenterToRegionRequest); i {
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
		file_region_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateDatacenterRequest); i {
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
		file_region_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteDatacenterRequest); i {
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
		file_region_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteRegionRequest); i {
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
		file_region_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetRegionRequest); i {
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
			RawDescriptor: file_region_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_region_proto_goTypes,
		DependencyIndexes: file_region_proto_depIdxs,
		EnumInfos:         file_region_proto_enumTypes,
		MessageInfos:      file_region_proto_msgTypes,
	}.Build()
	File_region_proto = out.File
	file_region_proto_rawDesc = nil
	file_region_proto_goTypes = nil
	file_region_proto_depIdxs = nil
}
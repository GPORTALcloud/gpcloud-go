// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: metadata.proto

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

type MetadataInterfaceType int32

const (
	MetadataInterfaceType_PUBLIC  MetadataInterfaceType = 0
	MetadataInterfaceType_PRIVATE MetadataInterfaceType = 1
)

// Enum value maps for MetadataInterfaceType.
var (
	MetadataInterfaceType_name = map[int32]string{
		0: "PUBLIC",
		1: "PRIVATE",
	}
	MetadataInterfaceType_value = map[string]int32{
		"PUBLIC":  0,
		"PRIVATE": 1,
	}
)

func (x MetadataInterfaceType) Enum() *MetadataInterfaceType {
	p := new(MetadataInterfaceType)
	*p = x
	return p
}

func (x MetadataInterfaceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetadataInterfaceType) Descriptor() protoreflect.EnumDescriptor {
	return file_metadata_proto_enumTypes[0].Descriptor()
}

func (MetadataInterfaceType) Type() protoreflect.EnumType {
	return &file_metadata_proto_enumTypes[0]
}

func (x MetadataInterfaceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetadataInterfaceType.Descriptor instead.
func (MetadataInterfaceType) EnumDescriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0}
}

type MetadataVLAN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tagged bool  `protobuf:"varint,2,opt,name=tagged,proto3" json:"tagged,omitempty"`
}

func (x *MetadataVLAN) Reset() {
	*x = MetadataVLAN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataVLAN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataVLAN) ProtoMessage() {}

func (x *MetadataVLAN) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataVLAN.ProtoReflect.Descriptor instead.
func (*MetadataVLAN) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataVLAN) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MetadataVLAN) GetTagged() bool {
	if x != nil {
		return x.Tagged
	}
	return false
}

type MetadataInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MetadataInterfaceType `protobuf:"varint,1,opt,name=type,proto3,enum=api.metadata.MetadataInterfaceType" json:"type,omitempty"`
	// IP Address
	Ipv4 *MetadataInterfaceIP `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"` //MetadataInterfaceIP ipv6 = 3;
	// Mac address
	Mac  string        `protobuf:"bytes,4,opt,name=mac,proto3" json:"mac,omitempty"`
	Vlan *MetadataVLAN `protobuf:"bytes,5,opt,name=vlan,proto3" json:"vlan,omitempty"`
}

func (x *MetadataInterface) Reset() {
	*x = MetadataInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataInterface) ProtoMessage() {}

func (x *MetadataInterface) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataInterface.ProtoReflect.Descriptor instead.
func (*MetadataInterface) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataInterface) GetType() MetadataInterfaceType {
	if x != nil {
		return x.Type
	}
	return MetadataInterfaceType_PUBLIC
}

func (x *MetadataInterface) GetIpv4() *MetadataInterfaceIP {
	if x != nil {
		return x.Ipv4
	}
	return nil
}

func (x *MetadataInterface) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *MetadataInterface) GetVlan() *MetadataVLAN {
	if x != nil {
		return x.Vlan
	}
	return nil
}

type MetadataInterfaceIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Netmask   string `protobuf:"bytes,2,opt,name=netmask,proto3" json:"netmask,omitempty"`
	Prefix    string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Gateway   string `protobuf:"bytes,4,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (x *MetadataInterfaceIP) Reset() {
	*x = MetadataInterfaceIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataInterfaceIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataInterfaceIP) ProtoMessage() {}

func (x *MetadataInterfaceIP) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataInterfaceIP.ProtoReflect.Descriptor instead.
func (*MetadataInterfaceIP) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *MetadataInterfaceIP) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *MetadataInterfaceIP) GetNetmask() string {
	if x != nil {
		return x.Netmask
	}
	return ""
}

func (x *MetadataInterfaceIP) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *MetadataInterfaceIP) GetGateway() string {
	if x != nil {
		return x.Gateway
	}
	return ""
}

type MetadataDNS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Nameserver
	Nameservers []string `protobuf:"bytes,1,rep,name=nameservers,proto3" json:"nameservers,omitempty"`
	// Search domains
	Search []string `protobuf:"bytes,2,rep,name=search,proto3" json:"search,omitempty"`
}

func (x *MetadataDNS) Reset() {
	*x = MetadataDNS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataDNS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataDNS) ProtoMessage() {}

func (x *MetadataDNS) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataDNS.ProtoReflect.Descriptor instead.
func (*MetadataDNS) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *MetadataDNS) GetNameservers() []string {
	if x != nil {
		return x.Nameservers
	}
	return nil
}

func (x *MetadataDNS) GetSearch() []string {
	if x != nil {
		return x.Search
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Compute resource type
	Type ComputeResourceType `protobuf:"varint,2,opt,name=type,proto3,enum=api.ComputeResourceType" json:"type,omitempty"`
	// FQDN
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Cloud Init configs
	VendorData string `protobuf:"bytes,4,opt,name=vendor_data,json=vendorData,proto3" json:"vendor_data,omitempty"`
	UserData   string `protobuf:"bytes,5,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
	// Cloud region
	Region string `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	// SSH public keys
	PublicKeys []string `protobuf:"bytes,7,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	// Network interfaces
	Interfaces []*MetadataInterface `protobuf:"bytes,8,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	// DNS settings
	Dns *MetadataDNS `protobuf:"bytes,9,opt,name=dns,proto3" json:"dns,omitempty"`
	// Tags
	Tags map[string]string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{4}
}

func (x *Metadata) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Metadata) GetType() ComputeResourceType {
	if x != nil {
		return x.Type
	}
	return ComputeResourceType_BARE_METAL
}

func (x *Metadata) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *Metadata) GetVendorData() string {
	if x != nil {
		return x.VendorData
	}
	return ""
}

func (x *Metadata) GetUserData() string {
	if x != nil {
		return x.UserData
	}
	return ""
}

func (x *Metadata) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Metadata) GetPublicKeys() []string {
	if x != nil {
		return x.PublicKeys
	}
	return nil
}

func (x *Metadata) GetInterfaces() []*MetadataInterface {
	if x != nil {
		return x.Interfaces
	}
	return nil
}

func (x *Metadata) GetDns() *MetadataDNS {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *Metadata) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type MetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *MetadataRequest) Reset() {
	*x = MetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataRequest) ProtoMessage() {}

func (x *MetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataRequest.ProtoReflect.Descriptor instead.
func (*MetadataRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{5}
}

func (x *MetadataRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type MetadataPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *MetadataPasswordResponse) Reset() {
	*x = MetadataPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataPasswordResponse) ProtoMessage() {}

func (x *MetadataPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataPasswordResponse.ProtoReflect.Descriptor instead.
func (*MetadataPasswordResponse) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{6}
}

func (x *MetadataPasswordResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type NetworkBootDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *NetworkBootDataRequest) Reset() {
	*x = NetworkBootDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkBootDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkBootDataRequest) ProtoMessage() {}

func (x *NetworkBootDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkBootDataRequest.ProtoReflect.Descriptor instead.
func (*NetworkBootDataRequest) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{7}
}

func (x *NetworkBootDataRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type NetworkBootDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script string `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
}

func (x *NetworkBootDataResponse) Reset() {
	*x = NetworkBootDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkBootDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkBootDataResponse) ProtoMessage() {}

func (x *NetworkBootDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkBootDataResponse.ProtoReflect.Descriptor instead.
func (*NetworkBootDataResponse) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{8}
}

func (x *NetworkBootDataResponse) GetScript() string {
	if x != nil {
		return x.Script
	}
	return ""
}

var File_metadata_proto protoreflect.FileDescriptor

var file_metadata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0d,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x0c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x4c, 0x41, 0x4e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x64, 0x22, 0xc5, 0x01, 0x0a, 0x11, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x49, 0x50, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x61, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x2e, 0x0a,
	0x04, 0x76, 0x6c, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x56, 0x4c, 0x41, 0x4e, 0x52, 0x04, 0x76, 0x6c, 0x61, 0x6e, 0x22, 0x80, 0x01,
	0x0a, 0x13, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x49, 0x50, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x22, 0x47, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x44, 0x4e, 0x53, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0xb8, 0x03, 0x0a, 0x08, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x0a, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x44, 0x4e, 0x53,
	0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54,
	0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x21, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x36, 0x0a, 0x18, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x28, 0x0a, 0x16, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x6f, 0x6f, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x31, 0x0a, 0x17, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x42, 0x6f, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2a, 0x30, 0x0a, 0x15,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x01, 0x42, 0x16,
	0x5a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b,
	0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_proto_rawDescOnce sync.Once
	file_metadata_proto_rawDescData = file_metadata_proto_rawDesc
)

func file_metadata_proto_rawDescGZIP() []byte {
	file_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_proto_rawDescData)
	})
	return file_metadata_proto_rawDescData
}

var file_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_metadata_proto_goTypes = []interface{}{
	(MetadataInterfaceType)(0),       // 0: api.metadata.MetadataInterfaceType
	(*MetadataVLAN)(nil),             // 1: api.metadata.MetadataVLAN
	(*MetadataInterface)(nil),        // 2: api.metadata.MetadataInterface
	(*MetadataInterfaceIP)(nil),      // 3: api.metadata.MetadataInterfaceIP
	(*MetadataDNS)(nil),              // 4: api.metadata.MetadataDNS
	(*Metadata)(nil),                 // 5: api.metadata.Metadata
	(*MetadataRequest)(nil),          // 6: api.metadata.MetadataRequest
	(*MetadataPasswordResponse)(nil), // 7: api.metadata.MetadataPasswordResponse
	(*NetworkBootDataRequest)(nil),   // 8: api.metadata.NetworkBootDataRequest
	(*NetworkBootDataResponse)(nil),  // 9: api.metadata.NetworkBootDataResponse
	nil,                              // 10: api.metadata.Metadata.TagsEntry
	(ComputeResourceType)(0),         // 11: api.ComputeResourceType
}
var file_metadata_proto_depIdxs = []int32{
	0,  // 0: api.metadata.MetadataInterface.type:type_name -> api.metadata.MetadataInterfaceType
	3,  // 1: api.metadata.MetadataInterface.ipv4:type_name -> api.metadata.MetadataInterfaceIP
	1,  // 2: api.metadata.MetadataInterface.vlan:type_name -> api.metadata.MetadataVLAN
	11, // 3: api.metadata.Metadata.type:type_name -> api.ComputeResourceType
	2,  // 4: api.metadata.Metadata.interfaces:type_name -> api.metadata.MetadataInterface
	4,  // 5: api.metadata.Metadata.dns:type_name -> api.metadata.MetadataDNS
	10, // 6: api.metadata.Metadata.tags:type_name -> api.metadata.Metadata.TagsEntry
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_metadata_proto_init() }
func file_metadata_proto_init() {
	if File_metadata_proto != nil {
		return
	}
	file_generic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataVLAN); i {
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
		file_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataInterface); i {
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
		file_metadata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataInterfaceIP); i {
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
		file_metadata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataDNS); i {
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
		file_metadata_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_metadata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataRequest); i {
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
		file_metadata_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataPasswordResponse); i {
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
		file_metadata_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkBootDataRequest); i {
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
		file_metadata_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkBootDataResponse); i {
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
			RawDescriptor: file_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_proto_depIdxs,
		EnumInfos:         file_metadata_proto_enumTypes,
		MessageInfos:      file_metadata_proto_msgTypes,
	}.Build()
	File_metadata_proto = out.File
	file_metadata_proto_rawDesc = nil
	file_metadata_proto_goTypes = nil
	file_metadata_proto_depIdxs = nil
}

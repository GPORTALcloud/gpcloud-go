// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: ptypes/agent.proto

package ptypes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AgentType int32

const (
	AgentType_NETWORK_AGENT AgentType = 0
)

// Enum value maps for AgentType.
var (
	AgentType_name = map[int32]string{
		0: "NETWORK_AGENT",
	}
	AgentType_value = map[string]int32{
		"NETWORK_AGENT": 0,
	}
)

func (x AgentType) Enum() *AgentType {
	p := new(AgentType)
	*p = x
	return p
}

func (x AgentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AgentType) Descriptor() protoreflect.EnumDescriptor {
	return file_ptypes_agent_proto_enumTypes[0].Descriptor()
}

func (AgentType) Type() protoreflect.EnumType {
	return &file_ptypes_agent_proto_enumTypes[0]
}

func (x AgentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AgentType.Descriptor instead.
func (AgentType) EnumDescriptor() ([]byte, []int) {
	return file_ptypes_agent_proto_rawDescGZIP(), []int{0}
}

type AgentCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateAuthority string `protobuf:"bytes,1,opt,name=certificate_authority,json=certificateAuthority,proto3" json:"certificate_authority,omitempty"`
	Certificate          string `protobuf:"bytes,2,opt,name=certificate,proto3" json:"certificate,omitempty"`
	PrivateKey           string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
}

func (x *AgentCertificate) Reset() {
	*x = AgentCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCertificate) ProtoMessage() {}

func (x *AgentCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCertificate.ProtoReflect.Descriptor instead.
func (*AgentCertificate) Descriptor() ([]byte, []int) {
	return file_ptypes_agent_proto_rawDescGZIP(), []int{0}
}

func (x *AgentCertificate) GetCertificateAuthority() string {
	if x != nil {
		return x.CertificateAuthority
	}
	return ""
}

func (x *AgentCertificate) GetCertificate() string {
	if x != nil {
		return x.Certificate
	}
	return ""
}

func (x *AgentCertificate) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          AgentType              `protobuf:"varint,2,opt,name=type,proto3,enum=api.agent.AgentType" json:"type,omitempty"`
	Datacenter    *DataCenter            `protobuf:"bytes,3,opt,name=datacenter,proto3" json:"datacenter,omitempty"`
	Fqdn          string                 `protobuf:"bytes,4,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	Certificate   *AgentCertificate      `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	LastContactAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_contact_at,json=lastContactAt,proto3" json:"last_contact_at,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_ptypes_agent_proto_rawDescGZIP(), []int{1}
}

func (x *Agent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Agent) GetType() AgentType {
	if x != nil {
		return x.Type
	}
	return AgentType_NETWORK_AGENT
}

func (x *Agent) GetDatacenter() *DataCenter {
	if x != nil {
		return x.Datacenter
	}
	return nil
}

func (x *Agent) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *Agent) GetCertificate() *AgentCertificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Agent) GetLastContactAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastContactAt
	}
	return nil
}

type AdminCreateAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         AgentType `protobuf:"varint,1,opt,name=type,proto3,enum=api.agent.AgentType" json:"type,omitempty"`
	DatacenterId string    `protobuf:"bytes,2,opt,name=datacenter_id,json=datacenterId,proto3" json:"datacenter_id,omitempty"`
	Fqdn         string    `protobuf:"bytes,3,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
}

func (x *AdminCreateAgentRequest) Reset() {
	*x = AdminCreateAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateAgentRequest) ProtoMessage() {}

func (x *AdminCreateAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateAgentRequest.ProtoReflect.Descriptor instead.
func (*AdminCreateAgentRequest) Descriptor() ([]byte, []int) {
	return file_ptypes_agent_proto_rawDescGZIP(), []int{2}
}

func (x *AdminCreateAgentRequest) GetType() AgentType {
	if x != nil {
		return x.Type
	}
	return AgentType_NETWORK_AGENT
}

func (x *AdminCreateAgentRequest) GetDatacenterId() string {
	if x != nil {
		return x.DatacenterId
	}
	return ""
}

func (x *AdminCreateAgentRequest) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

type AdminDeleteAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AdminDeleteAgentRequest) Reset() {
	*x = AdminDeleteAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteAgentRequest) ProtoMessage() {}

func (x *AdminDeleteAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteAgentRequest.ProtoReflect.Descriptor instead.
func (*AdminDeleteAgentRequest) Descriptor() ([]byte, []int) {
	return file_ptypes_agent_proto_rawDescGZIP(), []int{3}
}

func (x *AdminDeleteAgentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AdminListAgentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agents []*Agent `protobuf:"bytes,1,rep,name=agents,proto3" json:"agents,omitempty"`
}

func (x *AdminListAgentsResponse) Reset() {
	*x = AdminListAgentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ptypes_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminListAgentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminListAgentsResponse) ProtoMessage() {}

func (x *AdminListAgentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ptypes_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminListAgentsResponse.ProtoReflect.Descriptor instead.
func (*AdminListAgentsResponse) Descriptor() ([]byte, []int) {
	return file_ptypes_agent_proto_rawDescGZIP(), []int{4}
}

func (x *AdminListAgentsResponse) GetAgents() []*Agent {
	if x != nil {
		return x.Agents
	}
	return nil
}

var File_ptypes_agent_proto protoreflect.FileDescriptor

var file_ptypes_agent_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x22, 0x90, 0x02, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x71,
	0x64, 0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x41, 0x74, 0x22, 0x7c, 0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x71, 0x64, 0x6e, 0x22, 0x29, 0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x43,
	0x0a, 0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x2a, 0x1e, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x5f, 0x41, 0x47, 0x45, 0x4e,
	0x54, 0x10, 0x00, 0x42, 0x76, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x42, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x41, 0x41, 0x58, 0xaa, 0x02,
	0x09, 0x41, 0x70, 0x69, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0xca, 0x02, 0x09, 0x41, 0x70, 0x69,
	0x5c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0xe2, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ptypes_agent_proto_rawDescOnce sync.Once
	file_ptypes_agent_proto_rawDescData = file_ptypes_agent_proto_rawDesc
)

func file_ptypes_agent_proto_rawDescGZIP() []byte {
	file_ptypes_agent_proto_rawDescOnce.Do(func() {
		file_ptypes_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_ptypes_agent_proto_rawDescData)
	})
	return file_ptypes_agent_proto_rawDescData
}

var file_ptypes_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ptypes_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ptypes_agent_proto_goTypes = []interface{}{
	(AgentType)(0),                  // 0: api.agent.AgentType
	(*AgentCertificate)(nil),        // 1: api.agent.AgentCertificate
	(*Agent)(nil),                   // 2: api.agent.Agent
	(*AdminCreateAgentRequest)(nil), // 3: api.agent.AdminCreateAgentRequest
	(*AdminDeleteAgentRequest)(nil), // 4: api.agent.AdminDeleteAgentRequest
	(*AdminListAgentsResponse)(nil), // 5: api.agent.AdminListAgentsResponse
	(*DataCenter)(nil),              // 6: api.region.DataCenter
	(*timestamppb.Timestamp)(nil),   // 7: google.protobuf.Timestamp
}
var file_ptypes_agent_proto_depIdxs = []int32{
	0, // 0: api.agent.Agent.type:type_name -> api.agent.AgentType
	6, // 1: api.agent.Agent.datacenter:type_name -> api.region.DataCenter
	1, // 2: api.agent.Agent.certificate:type_name -> api.agent.AgentCertificate
	7, // 3: api.agent.Agent.last_contact_at:type_name -> google.protobuf.Timestamp
	0, // 4: api.agent.AdminCreateAgentRequest.type:type_name -> api.agent.AgentType
	2, // 5: api.agent.AdminListAgentsResponse.agents:type_name -> api.agent.Agent
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ptypes_agent_proto_init() }
func file_ptypes_agent_proto_init() {
	if File_ptypes_agent_proto != nil {
		return
	}
	file_ptypes_region_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ptypes_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentCertificate); i {
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
		file_ptypes_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
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
		file_ptypes_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateAgentRequest); i {
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
		file_ptypes_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteAgentRequest); i {
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
		file_ptypes_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminListAgentsResponse); i {
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
			RawDescriptor: file_ptypes_agent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ptypes_agent_proto_goTypes,
		DependencyIndexes: file_ptypes_agent_proto_depIdxs,
		EnumInfos:         file_ptypes_agent_proto_enumTypes,
		MessageInfos:      file_ptypes_agent_proto_msgTypes,
	}.Build()
	File_ptypes_agent_proto = out.File
	file_ptypes_agent_proto_rawDesc = nil
	file_ptypes_agent_proto_goTypes = nil
	file_ptypes_agent_proto_depIdxs = nil
}

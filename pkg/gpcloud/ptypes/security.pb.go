// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: security.proto

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

// SSHKey Represents a SSH public key
type SSHKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey   string                 `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Fingerprint string                 `protobuf:"bytes,4,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	Type        string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Bits        int32                  `protobuf:"varint,6,opt,name=bits,proto3" json:"bits,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *SSHKey) Reset() {
	*x = SSHKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSHKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSHKey) ProtoMessage() {}

func (x *SSHKey) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSHKey.ProtoReflect.Descriptor instead.
func (*SSHKey) Descriptor() ([]byte, []int) {
	return file_security_proto_rawDescGZIP(), []int{0}
}

func (x *SSHKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SSHKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSHKey) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *SSHKey) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *SSHKey) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SSHKey) GetBits() int32 {
	if x != nil {
		return x.Bits
	}
	return 0
}

func (x *SSHKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// UserBackupCode Represents a 2FA backup code for a user
type UserBackupCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Used bool   `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	Code string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *UserBackupCode) Reset() {
	*x = UserBackupCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBackupCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBackupCode) ProtoMessage() {}

func (x *UserBackupCode) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBackupCode.ProtoReflect.Descriptor instead.
func (*UserBackupCode) Descriptor() ([]byte, []int) {
	return file_security_proto_rawDescGZIP(), []int{1}
}

func (x *UserBackupCode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserBackupCode) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *UserBackupCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_security_proto protoreflect.FileDescriptor

var file_security_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd0, 0x01, 0x0a, 0x06, 0x53, 0x53, 0x48, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a,
	0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x69, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x62, 0x69, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x48, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x75, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x16, 0x5a, 0x14,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x70, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_proto_rawDescOnce sync.Once
	file_security_proto_rawDescData = file_security_proto_rawDesc
)

func file_security_proto_rawDescGZIP() []byte {
	file_security_proto_rawDescOnce.Do(func() {
		file_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_proto_rawDescData)
	})
	return file_security_proto_rawDescData
}

var file_security_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_security_proto_goTypes = []interface{}{
	(*SSHKey)(nil),                // 0: api.security.SSHKey
	(*UserBackupCode)(nil),        // 1: api.security.UserBackupCode
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_security_proto_depIdxs = []int32{
	2, // 0: api.security.SSHKey.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_security_proto_init() }
func file_security_proto_init() {
	if File_security_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSHKey); i {
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
		file_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBackupCode); i {
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
			RawDescriptor: file_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_proto_goTypes,
		DependencyIndexes: file_security_proto_depIdxs,
		MessageInfos:      file_security_proto_msgTypes,
	}.Build()
	File_security_proto = out.File
	file_security_proto_rawDesc = nil
	file_security_proto_goTypes = nil
	file_security_proto_depIdxs = nil
}
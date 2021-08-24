// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: basic.proto

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

type BasicProject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *BasicProject) Reset() {
	*x = BasicProject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicProject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicProject) ProtoMessage() {}

func (x *BasicProject) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicProject.ProtoReflect.Descriptor instead.
func (*BasicProject) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{0}
}

func (x *BasicProject) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BasicProject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BasicProject) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type BasicUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName string `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Avatar   string `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *BasicUser) Reset() {
	*x = BasicUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicUser) ProtoMessage() {}

func (x *BasicUser) ProtoReflect() protoreflect.Message {
	mi := &file_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicUser.ProtoReflect.Descriptor instead.
func (*BasicUser) Descriptor() ([]byte, []int) {
	return file_basic_proto_rawDescGZIP(), []int{1}
}

func (x *BasicUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BasicUser) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *BasicUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BasicUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

var File_basic_proto protoreflect.FileDescriptor

var file_basic_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x22, 0x4a, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x22, 0x66, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x16, 0x5a, 0x14,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x70, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basic_proto_rawDescOnce sync.Once
	file_basic_proto_rawDescData = file_basic_proto_rawDesc
)

func file_basic_proto_rawDescGZIP() []byte {
	file_basic_proto_rawDescOnce.Do(func() {
		file_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_proto_rawDescData)
	})
	return file_basic_proto_rawDescData
}

var file_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_basic_proto_goTypes = []interface{}{
	(*BasicProject)(nil), // 0: api.basic.BasicProject
	(*BasicUser)(nil),    // 1: api.basic.BasicUser
}
var file_basic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basic_proto_init() }
func file_basic_proto_init() {
	if File_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicProject); i {
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
		file_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicUser); i {
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
			RawDescriptor: file_basic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basic_proto_goTypes,
		DependencyIndexes: file_basic_proto_depIdxs,
		MessageInfos:      file_basic_proto_msgTypes,
	}.Build()
	File_basic_proto = out.File
	file_basic_proto_rawDesc = nil
	file_basic_proto_goTypes = nil
	file_basic_proto_depIdxs = nil
}

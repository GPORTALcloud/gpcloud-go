// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: billing.proto

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

type Product int32

const (
	Product_CREDIT          Product = 0
	Product_VOUCHER         Product = 1
	Product_COMPUTE         Product = 2
	Product_WINDOWS_LICENSE Product = 3
	Product_SUPPORT         Product = 4
	Product_TRAFFIC         Product = 5
)

// Enum value maps for Product.
var (
	Product_name = map[int32]string{
		0: "CREDIT",
		1: "VOUCHER",
		2: "COMPUTE",
		3: "WINDOWS_LICENSE",
		4: "SUPPORT",
		5: "TRAFFIC",
	}
	Product_value = map[string]int32{
		"CREDIT":          0,
		"VOUCHER":         1,
		"COMPUTE":         2,
		"WINDOWS_LICENSE": 3,
		"SUPPORT":         4,
		"TRAFFIC":         5,
	}
)

func (x Product) Enum() *Product {
	p := new(Product)
	*p = x
	return p
}

func (x Product) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Product) Descriptor() protoreflect.EnumDescriptor {
	return file_billing_proto_enumTypes[0].Descriptor()
}

func (Product) Type() protoreflect.EnumType {
	return &file_billing_proto_enumTypes[0]
}

func (x Product) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Product.Descriptor instead.
func (Product) EnumDescriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{0}
}

type BillStatus int32

const (
	BillStatus_UNPAID    BillStatus = 0
	BillStatus_PAID      BillStatus = 1
	BillStatus_CANCELLED BillStatus = 2
	BillStatus_REFUNDED  BillStatus = 3
	BillStatus_ERROR     BillStatus = 4
)

// Enum value maps for BillStatus.
var (
	BillStatus_name = map[int32]string{
		0: "UNPAID",
		1: "PAID",
		2: "CANCELLED",
		3: "REFUNDED",
		4: "ERROR",
	}
	BillStatus_value = map[string]int32{
		"UNPAID":    0,
		"PAID":      1,
		"CANCELLED": 2,
		"REFUNDED":  3,
		"ERROR":     4,
	}
)

func (x BillStatus) Enum() *BillStatus {
	p := new(BillStatus)
	*p = x
	return p
}

func (x BillStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BillStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_billing_proto_enumTypes[1].Descriptor()
}

func (BillStatus) Type() protoreflect.EnumType {
	return &file_billing_proto_enumTypes[1]
}

func (x BillStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BillStatus.Descriptor instead.
func (BillStatus) EnumDescriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{1}
}

type Bill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number             string                 `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Final              bool                   `protobuf:"varint,3,opt,name=final,proto3" json:"final,omitempty"`
	Currency           string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	Vat                int32                  `protobuf:"varint,5,opt,name=vat,proto3" json:"vat,omitempty"`
	Net                *Price                 `protobuf:"bytes,6,opt,name=net,proto3" json:"net,omitempty"`
	Gross              *Price                 `protobuf:"bytes,7,opt,name=gross,proto3" json:"gross,omitempty"`
	Status             BillStatus             `protobuf:"varint,8,opt,name=status,proto3,enum=api.billing.BillStatus" json:"status,omitempty"`
	PaymentMethod      PaymentMethod          `protobuf:"varint,9,opt,name=payment_method,json=paymentMethod,proto3,enum=api.payment.PaymentMethod" json:"payment_method,omitempty"`
	Items              []*BillItem            `protobuf:"bytes,10,rep,name=items,proto3" json:"items,omitempty"`
	OutstandingNet     *Price                 `protobuf:"bytes,11,opt,name=outstanding_net,json=outstandingNet,proto3" json:"outstanding_net,omitempty"`
	OutstandingGross   *Price                 `protobuf:"bytes,12,opt,name=outstanding_gross,json=outstandingGross,proto3" json:"outstanding_gross,omitempty"`
	DueAt              *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=due_at,json=dueAt,proto3" json:"due_at,omitempty"`
	CreatedAt          *timestamppb.Timestamp `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	User               *BasicUser             `protobuf:"bytes,15,opt,name=user,proto3" json:"user,omitempty"`
	Project            *BasicProject          `protobuf:"bytes,16,opt,name=project,proto3" json:"project,omitempty"`
	OutstandingBalance *Price                 `protobuf:"bytes,17,opt,name=outstanding_balance,json=outstandingBalance,proto3" json:"outstanding_balance,omitempty"`
}

func (x *Bill) Reset() {
	*x = Bill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bill) ProtoMessage() {}

func (x *Bill) ProtoReflect() protoreflect.Message {
	mi := &file_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bill.ProtoReflect.Descriptor instead.
func (*Bill) Descriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{0}
}

func (x *Bill) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bill) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Bill) GetFinal() bool {
	if x != nil {
		return x.Final
	}
	return false
}

func (x *Bill) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Bill) GetVat() int32 {
	if x != nil {
		return x.Vat
	}
	return 0
}

func (x *Bill) GetNet() *Price {
	if x != nil {
		return x.Net
	}
	return nil
}

func (x *Bill) GetGross() *Price {
	if x != nil {
		return x.Gross
	}
	return nil
}

func (x *Bill) GetStatus() BillStatus {
	if x != nil {
		return x.Status
	}
	return BillStatus_UNPAID
}

func (x *Bill) GetPaymentMethod() PaymentMethod {
	if x != nil {
		return x.PaymentMethod
	}
	return PaymentMethod_CREDIT_CARD
}

func (x *Bill) GetItems() []*BillItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Bill) GetOutstandingNet() *Price {
	if x != nil {
		return x.OutstandingNet
	}
	return nil
}

func (x *Bill) GetOutstandingGross() *Price {
	if x != nil {
		return x.OutstandingGross
	}
	return nil
}

func (x *Bill) GetDueAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DueAt
	}
	return nil
}

func (x *Bill) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Bill) GetUser() *BasicUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Bill) GetProject() *BasicProject {
	if x != nil {
		return x.Project
	}
	return nil
}

func (x *Bill) GetOutstandingBalance() *Price {
	if x != nil {
		return x.OutstandingBalance
	}
	return nil
}

type BillItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product     Product                `protobuf:"varint,2,opt,name=product,proto3,enum=api.billing.Product" json:"product,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       *Price                 `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	StartedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ended_at,json=endedAt,proto3" json:"ended_at,omitempty"`
}

func (x *BillItem) Reset() {
	*x = BillItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillItem) ProtoMessage() {}

func (x *BillItem) ProtoReflect() protoreflect.Message {
	mi := &file_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillItem.ProtoReflect.Descriptor instead.
func (*BillItem) Descriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{1}
}

func (x *BillItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BillItem) GetProduct() Product {
	if x != nil {
		return x.Product
	}
	return Product_CREDIT
}

func (x *BillItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BillItem) GetPrice() *Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *BillItem) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *BillItem) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

var File_billing_proto protoreflect.FileDescriptor

var file_billing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x05, 0x0a, 0x04, 0x42, 0x69, 0x6c,
	0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x76,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76, 0x61, 0x74, 0x12, 0x1c, 0x0a,
	0x03, 0x6e, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x03, 0x6e, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x12, 0x2f, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x41,
	0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x12, 0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x33,
	0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x65,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x4e, 0x65, 0x74, 0x12, 0x37, 0x0a, 0x11, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x73,
	0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x06,
	0x64, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x64, 0x75, 0x65, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x13, 0x6f, 0x75, 0x74, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x12, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x08, 0x42, 0x69, 0x6c, 0x6c, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x5e, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x56, 0x4f, 0x55, 0x43, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43,
	0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x49, 0x4e, 0x44,
	0x4f, 0x57, 0x53, 0x5f, 0x4c, 0x49, 0x43, 0x45, 0x4e, 0x53, 0x45, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x52,
	0x41, 0x46, 0x46, 0x49, 0x43, 0x10, 0x05, 0x2a, 0x4a, 0x0a, 0x0a, 0x42, 0x69, 0x6c, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x50, 0x41, 0x49, 0x44, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x41, 0x49, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x46, 0x55, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x04, 0x42, 0x16, 0x5a, 0x14, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x3b, 0x70, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_billing_proto_rawDescOnce sync.Once
	file_billing_proto_rawDescData = file_billing_proto_rawDesc
)

func file_billing_proto_rawDescGZIP() []byte {
	file_billing_proto_rawDescOnce.Do(func() {
		file_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_billing_proto_rawDescData)
	})
	return file_billing_proto_rawDescData
}

var file_billing_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_billing_proto_goTypes = []interface{}{
	(Product)(0),                  // 0: api.billing.Product
	(BillStatus)(0),               // 1: api.billing.BillStatus
	(*Bill)(nil),                  // 2: api.billing.Bill
	(*BillItem)(nil),              // 3: api.billing.BillItem
	(*Price)(nil),                 // 4: api.Price
	(PaymentMethod)(0),            // 5: api.payment.PaymentMethod
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*BasicUser)(nil),             // 7: api.basic.BasicUser
	(*BasicProject)(nil),          // 8: api.basic.BasicProject
}
var file_billing_proto_depIdxs = []int32{
	4,  // 0: api.billing.Bill.net:type_name -> api.Price
	4,  // 1: api.billing.Bill.gross:type_name -> api.Price
	1,  // 2: api.billing.Bill.status:type_name -> api.billing.BillStatus
	5,  // 3: api.billing.Bill.payment_method:type_name -> api.payment.PaymentMethod
	3,  // 4: api.billing.Bill.items:type_name -> api.billing.BillItem
	4,  // 5: api.billing.Bill.outstanding_net:type_name -> api.Price
	4,  // 6: api.billing.Bill.outstanding_gross:type_name -> api.Price
	6,  // 7: api.billing.Bill.due_at:type_name -> google.protobuf.Timestamp
	6,  // 8: api.billing.Bill.created_at:type_name -> google.protobuf.Timestamp
	7,  // 9: api.billing.Bill.user:type_name -> api.basic.BasicUser
	8,  // 10: api.billing.Bill.project:type_name -> api.basic.BasicProject
	4,  // 11: api.billing.Bill.outstanding_balance:type_name -> api.Price
	0,  // 12: api.billing.BillItem.product:type_name -> api.billing.Product
	4,  // 13: api.billing.BillItem.price:type_name -> api.Price
	6,  // 14: api.billing.BillItem.started_at:type_name -> google.protobuf.Timestamp
	6,  // 15: api.billing.BillItem.ended_at:type_name -> google.protobuf.Timestamp
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_billing_proto_init() }
func file_billing_proto_init() {
	if File_billing_proto != nil {
		return
	}
	file_generic_proto_init()
	file_basic_proto_init()
	file_payment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bill); i {
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
		file_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillItem); i {
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
			RawDescriptor: file_billing_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_billing_proto_goTypes,
		DependencyIndexes: file_billing_proto_depIdxs,
		EnumInfos:         file_billing_proto_enumTypes,
		MessageInfos:      file_billing_proto_msgTypes,
	}.Build()
	File_billing_proto = out.File
	file_billing_proto_rawDesc = nil
	file_billing_proto_goTypes = nil
	file_billing_proto_depIdxs = nil
}
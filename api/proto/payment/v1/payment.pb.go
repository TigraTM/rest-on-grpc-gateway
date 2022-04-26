// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: api/proto/payment/v1/payment.proto

package pb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	decimal "google.golang.org/genproto/googleapis/type/decimal"
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

type CreatePaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum         *decimal.Decimal `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	CompanyName string           `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Category    string           `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	UserId      int64            `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreatePaymentRequest) Reset() {
	*x = CreatePaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentRequest) ProtoMessage() {}

func (x *CreatePaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentRequest.ProtoReflect.Descriptor instead.
func (*CreatePaymentRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaymentRequest) GetSum() *decimal.Decimal {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *CreatePaymentRequest) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CreatePaymentRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreatePaymentRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type CreatePaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePaymentResponse) Reset() {
	*x = CreatePaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaymentResponse) ProtoMessage() {}

func (x *CreatePaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaymentResponse.ProtoReflect.Descriptor instead.
func (*CreatePaymentResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{1}
}

type GetBalanceByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Currency string `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *GetBalanceByUserIDRequest) Reset() {
	*x = GetBalanceByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceByUserIDRequest) ProtoMessage() {}

func (x *GetBalanceByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceByUserIDRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetBalanceByUserIDRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type GetBalanceByUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64            `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Sum      *decimal.Decimal `protobuf:"bytes,2,opt,name=sum,proto3" json:"sum,omitempty"`
	Currency string           `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *GetBalanceByUserIDResponse) Reset() {
	*x = GetBalanceByUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceByUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceByUserIDResponse) ProtoMessage() {}

func (x *GetBalanceByUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceByUserIDResponse.ProtoReflect.Descriptor instead.
func (*GetBalanceByUserIDResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetBalanceByUserIDResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetBalanceByUserIDResponse) GetSum() *decimal.Decimal {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *GetBalanceByUserIDResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type TransferBetweenUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum           *decimal.Decimal `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	SenderId      int64            `protobuf:"varint,2,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecipientId   int64            `protobuf:"varint,3,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	RecipientName string           `protobuf:"bytes,4,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
}

func (x *TransferBetweenUsersRequest) Reset() {
	*x = TransferBetweenUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferBetweenUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferBetweenUsersRequest) ProtoMessage() {}

func (x *TransferBetweenUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferBetweenUsersRequest.ProtoReflect.Descriptor instead.
func (*TransferBetweenUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{4}
}

func (x *TransferBetweenUsersRequest) GetSum() *decimal.Decimal {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *TransferBetweenUsersRequest) GetSenderId() int64 {
	if x != nil {
		return x.SenderId
	}
	return 0
}

func (x *TransferBetweenUsersRequest) GetRecipientId() int64 {
	if x != nil {
		return x.RecipientId
	}
	return 0
}

func (x *TransferBetweenUsersRequest) GetRecipientName() string {
	if x != nil {
		return x.RecipientName
	}
	return ""
}

type TransferBetweenUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum           *decimal.Decimal `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
	RecipientId   int64            `protobuf:"varint,2,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
	RecipientName string           `protobuf:"bytes,3,opt,name=recipient_name,json=recipientName,proto3" json:"recipient_name,omitempty"`
}

func (x *TransferBetweenUsersResponse) Reset() {
	*x = TransferBetweenUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferBetweenUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferBetweenUsersResponse) ProtoMessage() {}

func (x *TransferBetweenUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferBetweenUsersResponse.ProtoReflect.Descriptor instead.
func (*TransferBetweenUsersResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{5}
}

func (x *TransferBetweenUsersResponse) GetSum() *decimal.Decimal {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *TransferBetweenUsersResponse) GetRecipientId() int64 {
	if x != nil {
		return x.RecipientId
	}
	return 0
}

func (x *TransferBetweenUsersResponse) GetRecipientName() string {
	if x != nil {
		return x.RecipientName
	}
	return ""
}

type GetPaymentsHistoryByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	SortBy string `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	Sort   string `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
}

func (x *GetPaymentsHistoryByUserIDRequest) Reset() {
	*x = GetPaymentsHistoryByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentsHistoryByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentsHistoryByUserIDRequest) ProtoMessage() {}

func (x *GetPaymentsHistoryByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentsHistoryByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentsHistoryByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{6}
}

func (x *GetPaymentsHistoryByUserIDRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetPaymentsHistoryByUserIDRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPaymentsHistoryByUserIDRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetPaymentsHistoryByUserIDRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *GetPaymentsHistoryByUserIDRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type GetPaymentsHistoryByUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payments []*Payment `protobuf:"bytes,1,rep,name=payments,proto3" json:"payments,omitempty"`
	Total    int64      `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetPaymentsHistoryByUserIDResponse) Reset() {
	*x = GetPaymentsHistoryByUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentsHistoryByUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentsHistoryByUserIDResponse) ProtoMessage() {}

func (x *GetPaymentsHistoryByUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentsHistoryByUserIDResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentsHistoryByUserIDResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{7}
}

func (x *GetPaymentsHistoryByUserIDResponse) GetPayments() []*Payment {
	if x != nil {
		return x.Payments
	}
	return nil
}

func (x *GetPaymentsHistoryByUserIDResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateAt    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	Sum         *decimal.Decimal       `protobuf:"bytes,3,opt,name=sum,proto3" json:"sum,omitempty"`
	CompanyName string                 `protobuf:"bytes,4,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	Category    string                 `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_payment_v1_payment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_payment_v1_payment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_api_proto_payment_v1_payment_proto_rawDescGZIP(), []int{8}
}

func (x *Payment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetCreateAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateAt
	}
	return nil
}

func (x *Payment) GetSum() *decimal.Decimal {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *Payment) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *Payment) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_api_proto_payment_v1_payment_proto protoreflect.FileDescriptor

var file_api_proto_payment_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6,
	0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04,
	0x10, 0x01, 0x18, 0x32, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x5d, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x98,
	0x01, 0x03, 0xd0, 0x01, 0x01, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22,
	0x79, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0xc1, 0x01, 0x0a, 0x1b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x03, 0x73, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x0e,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52,
	0x0d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x90,
	0x01, 0x0a, 0x1c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x65, 0x74, 0x77, 0x65,
	0x65, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0xd5, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x22, 0x05, 0x18, 0xf4, 0x03, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x28, 0x00, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x62, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x12, 0x72, 0x10, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x52, 0x06, 0x73, 0x6f, 0x72,
	0x74, 0x42, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0x72, 0x0b, 0x52, 0x03, 0x61, 0x73, 0x63, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0x75, 0x0a, 0x22, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x08, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0xd9, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa,
	0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x32, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01,
	0x18, 0x32, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0xb2, 0x05, 0x0a,
	0x0a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x50, 0x49, 0x12, 0x8c, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17,
	0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0xa2, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0xaa, 0x01, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x65, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x65, 0x74, 0x77, 0x65,
	0x65, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x22, 0x20, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xc2, 0x01, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x37, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x42, 0x2e, 0x5a, 0x2c, 0x72, 0x65, 0x73, 0x74, 0x2d, 0x6f, 0x6e, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_payment_v1_payment_proto_rawDescOnce sync.Once
	file_api_proto_payment_v1_payment_proto_rawDescData = file_api_proto_payment_v1_payment_proto_rawDesc
)

func file_api_proto_payment_v1_payment_proto_rawDescGZIP() []byte {
	file_api_proto_payment_v1_payment_proto_rawDescOnce.Do(func() {
		file_api_proto_payment_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_payment_v1_payment_proto_rawDescData)
	})
	return file_api_proto_payment_v1_payment_proto_rawDescData
}

var file_api_proto_payment_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_proto_payment_v1_payment_proto_goTypes = []interface{}{
	(*CreatePaymentRequest)(nil),               // 0: api.proto.payment.v1.CreatePaymentRequest
	(*CreatePaymentResponse)(nil),              // 1: api.proto.payment.v1.CreatePaymentResponse
	(*GetBalanceByUserIDRequest)(nil),          // 2: api.proto.payment.v1.GetBalanceByUserIDRequest
	(*GetBalanceByUserIDResponse)(nil),         // 3: api.proto.payment.v1.GetBalanceByUserIDResponse
	(*TransferBetweenUsersRequest)(nil),        // 4: api.proto.payment.v1.TransferBetweenUsersRequest
	(*TransferBetweenUsersResponse)(nil),       // 5: api.proto.payment.v1.TransferBetweenUsersResponse
	(*GetPaymentsHistoryByUserIDRequest)(nil),  // 6: api.proto.payment.v1.GetPaymentsHistoryByUserIDRequest
	(*GetPaymentsHistoryByUserIDResponse)(nil), // 7: api.proto.payment.v1.GetPaymentsHistoryByUserIDResponse
	(*Payment)(nil),                            // 8: api.proto.payment.v1.Payment
	(*decimal.Decimal)(nil),                    // 9: google.type.Decimal
	(*timestamppb.Timestamp)(nil),              // 10: google.protobuf.Timestamp
}
var file_api_proto_payment_v1_payment_proto_depIdxs = []int32{
	9,  // 0: api.proto.payment.v1.CreatePaymentRequest.sum:type_name -> google.type.Decimal
	9,  // 1: api.proto.payment.v1.GetBalanceByUserIDResponse.sum:type_name -> google.type.Decimal
	9,  // 2: api.proto.payment.v1.TransferBetweenUsersRequest.sum:type_name -> google.type.Decimal
	9,  // 3: api.proto.payment.v1.TransferBetweenUsersResponse.sum:type_name -> google.type.Decimal
	8,  // 4: api.proto.payment.v1.GetPaymentsHistoryByUserIDResponse.payments:type_name -> api.proto.payment.v1.Payment
	10, // 5: api.proto.payment.v1.Payment.create_at:type_name -> google.protobuf.Timestamp
	9,  // 6: api.proto.payment.v1.Payment.sum:type_name -> google.type.Decimal
	0,  // 7: api.proto.payment.v1.PaymentAPI.CreatePayment:input_type -> api.proto.payment.v1.CreatePaymentRequest
	2,  // 8: api.proto.payment.v1.PaymentAPI.GetBalanceByUserID:input_type -> api.proto.payment.v1.GetBalanceByUserIDRequest
	4,  // 9: api.proto.payment.v1.PaymentAPI.TransferBetweenUsers:input_type -> api.proto.payment.v1.TransferBetweenUsersRequest
	6,  // 10: api.proto.payment.v1.PaymentAPI.GetPaymentsHistoryByUserID:input_type -> api.proto.payment.v1.GetPaymentsHistoryByUserIDRequest
	1,  // 11: api.proto.payment.v1.PaymentAPI.CreatePayment:output_type -> api.proto.payment.v1.CreatePaymentResponse
	3,  // 12: api.proto.payment.v1.PaymentAPI.GetBalanceByUserID:output_type -> api.proto.payment.v1.GetBalanceByUserIDResponse
	5,  // 13: api.proto.payment.v1.PaymentAPI.TransferBetweenUsers:output_type -> api.proto.payment.v1.TransferBetweenUsersResponse
	7,  // 14: api.proto.payment.v1.PaymentAPI.GetPaymentsHistoryByUserID:output_type -> api.proto.payment.v1.GetPaymentsHistoryByUserIDResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_api_proto_payment_v1_payment_proto_init() }
func file_api_proto_payment_v1_payment_proto_init() {
	if File_api_proto_payment_v1_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_payment_v1_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentRequest); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaymentResponse); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceByUserIDRequest); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceByUserIDResponse); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferBetweenUsersRequest); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferBetweenUsersResponse); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentsHistoryByUserIDRequest); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentsHistoryByUserIDResponse); i {
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
		file_api_proto_payment_v1_payment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
			RawDescriptor: file_api_proto_payment_v1_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_payment_v1_payment_proto_goTypes,
		DependencyIndexes: file_api_proto_payment_v1_payment_proto_depIdxs,
		MessageInfos:      file_api_proto_payment_v1_payment_proto_msgTypes,
	}.Build()
	File_api_proto_payment_v1_payment_proto = out.File
	file_api_proto_payment_v1_payment_proto_rawDesc = nil
	file_api_proto_payment_v1_payment_proto_goTypes = nil
	file_api_proto_payment_v1_payment_proto_depIdxs = nil
}

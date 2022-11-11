// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: account.proto

package proto

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

// NEW ACCOUNT
type NewAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdNationality int64  `protobuf:"varint,1,opt,name=IdNationality,proto3" json:"IdNationality,omitempty"`
	Firstname     string `protobuf:"bytes,2,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname      string `protobuf:"bytes,3,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Email         string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password      string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	PhonePrefix   string `protobuf:"bytes,6,opt,name=PhonePrefix,proto3" json:"PhonePrefix,omitempty"`
	Ddd           string `protobuf:"bytes,7,opt,name=Ddd,proto3" json:"Ddd,omitempty"`
	Phone         string `protobuf:"bytes,8,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *NewAccountRequest) Reset() {
	*x = NewAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAccountRequest) ProtoMessage() {}

func (x *NewAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAccountRequest.ProtoReflect.Descriptor instead.
func (*NewAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *NewAccountRequest) GetIdNationality() int64 {
	if x != nil {
		return x.IdNationality
	}
	return 0
}

func (x *NewAccountRequest) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *NewAccountRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *NewAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewAccountRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *NewAccountRequest) GetPhonePrefix() string {
	if x != nil {
		return x.PhonePrefix
	}
	return ""
}

func (x *NewAccountRequest) GetDdd() string {
	if x != nil {
		return x.Ddd
	}
	return ""
}

func (x *NewAccountRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

// ACCOUNT DETAILS
type AccountDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Firstname       string `protobuf:"bytes,2,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname        string `protobuf:"bytes,3,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Email           string `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password        string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	PhonePrefix     int64  `protobuf:"varint,6,opt,name=PhonePrefix,proto3" json:"PhonePrefix,omitempty"`
	Ddd             string `protobuf:"bytes,7,opt,name=Ddd,proto3" json:"Ddd,omitempty"`
	Phone           string `protobuf:"bytes,8,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Deleted         bool   `protobuf:"varint,9,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	IdNationality   int64  `protobuf:"varint,10,opt,name=IdNationality,proto3" json:"IdNationality,omitempty"`
	NationalityName string `protobuf:"bytes,11,opt,name=NationalityName,proto3" json:"NationalityName,omitempty"`
	NationalityKey  string `protobuf:"bytes,12,opt,name=NationalityKey,proto3" json:"NationalityKey,omitempty"`
}

func (x *AccountDetails) Reset() {
	*x = AccountDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountDetails) ProtoMessage() {}

func (x *AccountDetails) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountDetails.ProtoReflect.Descriptor instead.
func (*AccountDetails) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountDetails) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AccountDetails) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *AccountDetails) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *AccountDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AccountDetails) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AccountDetails) GetPhonePrefix() int64 {
	if x != nil {
		return x.PhonePrefix
	}
	return 0
}

func (x *AccountDetails) GetDdd() string {
	if x != nil {
		return x.Ddd
	}
	return ""
}

func (x *AccountDetails) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AccountDetails) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *AccountDetails) GetIdNationality() int64 {
	if x != nil {
		return x.IdNationality
	}
	return 0
}

func (x *AccountDetails) GetNationalityName() string {
	if x != nil {
		return x.NationalityName
	}
	return ""
}

func (x *AccountDetails) GetNationalityKey() string {
	if x != nil {
		return x.NationalityKey
	}
	return ""
}

type AccountDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *AccountDetails   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Config  *AppConfigDetails `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *AccountDetailsResponse) Reset() {
	*x = AccountDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountDetailsResponse) ProtoMessage() {}

func (x *AccountDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountDetailsResponse.ProtoReflect.Descriptor instead.
func (*AccountDetailsResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccountDetailsResponse) GetAccount() *AccountDetails {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *AccountDetailsResponse) GetConfig() *AppConfigDetails {
	if x != nil {
		return x.Config
	}
	return nil
}

type EditAccountInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	IdNationality int64  `protobuf:"varint,2,opt,name=IdNationality,proto3" json:"IdNationality,omitempty"`
	Firstname     string `protobuf:"bytes,3,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname      string `protobuf:"bytes,4,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	Email         string `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`
	Password      string `protobuf:"bytes,6,opt,name=Password,proto3" json:"Password,omitempty"`
	PhonePrefix   string `protobuf:"bytes,7,opt,name=PhonePrefix,proto3" json:"PhonePrefix,omitempty"`
	Ddd           string `protobuf:"bytes,8,opt,name=Ddd,proto3" json:"Ddd,omitempty"`
	Phone         string `protobuf:"bytes,9,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Deleted       bool   `protobuf:"varint,10,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
}

func (x *EditAccountInfos) Reset() {
	*x = EditAccountInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditAccountInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditAccountInfos) ProtoMessage() {}

func (x *EditAccountInfos) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditAccountInfos.ProtoReflect.Descriptor instead.
func (*EditAccountInfos) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *EditAccountInfos) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditAccountInfos) GetIdNationality() int64 {
	if x != nil {
		return x.IdNationality
	}
	return 0
}

func (x *EditAccountInfos) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *EditAccountInfos) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *EditAccountInfos) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EditAccountInfos) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *EditAccountInfos) GetPhonePrefix() string {
	if x != nil {
		return x.PhonePrefix
	}
	return ""
}

func (x *EditAccountInfos) GetDdd() string {
	if x != nil {
		return x.Ddd
	}
	return ""
}

func (x *EditAccountInfos) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *EditAccountInfos) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type EditAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *EditAccountInfos `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Config  *AppConfig        `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *EditAccountRequest) Reset() {
	*x = EditAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditAccountRequest) ProtoMessage() {}

func (x *EditAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditAccountRequest.ProtoReflect.Descriptor instead.
func (*EditAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *EditAccountRequest) GetAccount() *EditAccountInfos {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *EditAccountRequest) GetConfig() *AppConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x64, 0x4e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x49, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x10, 0x0a,
	0x03, 0x44, 0x64, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x64, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xe8, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x64, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x44, 0x64, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x64, 0x4e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x49, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x28,
	0x0a, 0x0f, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x4e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79,
	0x22, 0x7a, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x98, 0x02, 0x0a,
	0x10, 0x45, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x49, 0x64, 0x4e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x44, 0x64, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x44, 0x64, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x71, 0x0a, 0x12, 0x45, 0x64, 0x69, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x32, 0xa5, 0x02, 0x0a, 0x0e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64,
	0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x64, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x07, 0x52,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x64, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_account_proto_goTypes = []interface{}{
	(*NewAccountRequest)(nil),      // 0: proto.NewAccountRequest
	(*AccountDetails)(nil),         // 1: proto.AccountDetails
	(*AccountDetailsResponse)(nil), // 2: proto.AccountDetailsResponse
	(*EditAccountInfos)(nil),       // 3: proto.EditAccountInfos
	(*EditAccountRequest)(nil),     // 4: proto.EditAccountRequest
	(*AppConfigDetails)(nil),       // 5: proto.AppConfigDetails
	(*AppConfig)(nil),              // 6: proto.AppConfig
	(*Id)(nil),                     // 7: proto.Id
	(*StatusResponse)(nil),         // 8: proto.StatusResponse
}
var file_account_proto_depIdxs = []int32{
	1, // 0: proto.AccountDetailsResponse.account:type_name -> proto.AccountDetails
	5, // 1: proto.AccountDetailsResponse.config:type_name -> proto.AppConfigDetails
	3, // 2: proto.EditAccountRequest.account:type_name -> proto.EditAccountInfos
	6, // 3: proto.EditAccountRequest.config:type_name -> proto.AppConfig
	0, // 4: proto.AccountService.Create:input_type -> proto.NewAccountRequest
	7, // 5: proto.AccountService.Details:input_type -> proto.Id
	4, // 6: proto.AccountService.Edit:input_type -> proto.EditAccountRequest
	7, // 7: proto.AccountService.Delete:input_type -> proto.Id
	7, // 8: proto.AccountService.Restore:input_type -> proto.Id
	8, // 9: proto.AccountService.Create:output_type -> proto.StatusResponse
	2, // 10: proto.AccountService.Details:output_type -> proto.AccountDetailsResponse
	8, // 11: proto.AccountService.Edit:output_type -> proto.StatusResponse
	8, // 12: proto.AccountService.Delete:output_type -> proto.StatusResponse
	2, // 13: proto.AccountService.Restore:output_type -> proto.AccountDetailsResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	file_config_proto_init()
	file_generic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAccountRequest); i {
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
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountDetails); i {
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
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountDetailsResponse); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditAccountInfos); i {
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
		file_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditAccountRequest); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}

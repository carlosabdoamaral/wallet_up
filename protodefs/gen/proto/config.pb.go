// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: config.proto

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

type AppConfigDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdConfig          int64  `protobuf:"varint,1,opt,name=IdConfig,proto3" json:"IdConfig,omitempty"`
	AppTheme          string `protobuf:"bytes,2,opt,name=AppTheme,proto3" json:"AppTheme,omitempty"`
	BiometryActivated bool   `protobuf:"varint,3,opt,name=BiometryActivated,proto3" json:"BiometryActivated,omitempty"`
	AlertOnEmail      bool   `protobuf:"varint,4,opt,name=AlertOnEmail,proto3" json:"AlertOnEmail,omitempty"`
	AlertOnMobile     string `protobuf:"bytes,5,opt,name=AlertOnMobile,proto3" json:"AlertOnMobile,omitempty"`
	AppLanguageId     int64  `protobuf:"varint,6,opt,name=AppLanguageId,proto3" json:"AppLanguageId,omitempty"`
	AppLanguage       string `protobuf:"bytes,7,opt,name=AppLanguage,proto3" json:"AppLanguage,omitempty"`
	AppLanguageKey    string `protobuf:"bytes,8,opt,name=AppLanguageKey,proto3" json:"AppLanguageKey,omitempty"`
}

func (x *AppConfigDetails) Reset() {
	*x = AppConfigDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfigDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfigDetails) ProtoMessage() {}

func (x *AppConfigDetails) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfigDetails.ProtoReflect.Descriptor instead.
func (*AppConfigDetails) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *AppConfigDetails) GetIdConfig() int64 {
	if x != nil {
		return x.IdConfig
	}
	return 0
}

func (x *AppConfigDetails) GetAppTheme() string {
	if x != nil {
		return x.AppTheme
	}
	return ""
}

func (x *AppConfigDetails) GetBiometryActivated() bool {
	if x != nil {
		return x.BiometryActivated
	}
	return false
}

func (x *AppConfigDetails) GetAlertOnEmail() bool {
	if x != nil {
		return x.AlertOnEmail
	}
	return false
}

func (x *AppConfigDetails) GetAlertOnMobile() string {
	if x != nil {
		return x.AlertOnMobile
	}
	return ""
}

func (x *AppConfigDetails) GetAppLanguageId() int64 {
	if x != nil {
		return x.AppLanguageId
	}
	return 0
}

func (x *AppConfigDetails) GetAppLanguage() string {
	if x != nil {
		return x.AppLanguage
	}
	return ""
}

func (x *AppConfigDetails) GetAppLanguageKey() string {
	if x != nil {
		return x.AppLanguageKey
	}
	return ""
}

type AppConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	IdLanguage        int64  `protobuf:"varint,2,opt,name=Id_language,json=IdLanguage,proto3" json:"Id_language,omitempty"`
	IdCurrency        int64  `protobuf:"varint,3,opt,name=Id_currency,json=IdCurrency,proto3" json:"Id_currency,omitempty"`
	Theme             string `protobuf:"bytes,4,opt,name=Theme,proto3" json:"Theme,omitempty"`
	BiometryActivated bool   `protobuf:"varint,5,opt,name=BiometryActivated,proto3" json:"BiometryActivated,omitempty"`
	AlertOnEmail      bool   `protobuf:"varint,6,opt,name=AlertOnEmail,proto3" json:"AlertOnEmail,omitempty"`
	AlertOnMobile     bool   `protobuf:"varint,7,opt,name=AlertOnMobile,proto3" json:"AlertOnMobile,omitempty"`
}

func (x *AppConfig) Reset() {
	*x = AppConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfig) ProtoMessage() {}

func (x *AppConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfig.ProtoReflect.Descriptor instead.
func (*AppConfig) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *AppConfig) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppConfig) GetIdLanguage() int64 {
	if x != nil {
		return x.IdLanguage
	}
	return 0
}

func (x *AppConfig) GetIdCurrency() int64 {
	if x != nil {
		return x.IdCurrency
	}
	return 0
}

func (x *AppConfig) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *AppConfig) GetBiometryActivated() bool {
	if x != nil {
		return x.BiometryActivated
	}
	return false
}

func (x *AppConfig) GetAlertOnEmail() bool {
	if x != nil {
		return x.AlertOnEmail
	}
	return false
}

func (x *AppConfig) GetAlertOnMobile() bool {
	if x != nil {
		return x.AlertOnMobile
	}
	return false
}

type NewAppConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdLanguage        int64  `protobuf:"varint,1,opt,name=Id_language,json=IdLanguage,proto3" json:"Id_language,omitempty"`
	IdCurrency        int64  `protobuf:"varint,2,opt,name=Id_currency,json=IdCurrency,proto3" json:"Id_currency,omitempty"`
	Theme             string `protobuf:"bytes,3,opt,name=Theme,proto3" json:"Theme,omitempty"`
	BiometryActivated bool   `protobuf:"varint,4,opt,name=BiometryActivated,proto3" json:"BiometryActivated,omitempty"`
	AlertOnEmail      bool   `protobuf:"varint,5,opt,name=AlertOnEmail,proto3" json:"AlertOnEmail,omitempty"`
	AlertOnMobile     bool   `protobuf:"varint,6,opt,name=AlertOnMobile,proto3" json:"AlertOnMobile,omitempty"`
}

func (x *NewAppConfigRequest) Reset() {
	*x = NewAppConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewAppConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAppConfigRequest) ProtoMessage() {}

func (x *NewAppConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAppConfigRequest.ProtoReflect.Descriptor instead.
func (*NewAppConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *NewAppConfigRequest) GetIdLanguage() int64 {
	if x != nil {
		return x.IdLanguage
	}
	return 0
}

func (x *NewAppConfigRequest) GetIdCurrency() int64 {
	if x != nil {
		return x.IdCurrency
	}
	return 0
}

func (x *NewAppConfigRequest) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *NewAppConfigRequest) GetBiometryActivated() bool {
	if x != nil {
		return x.BiometryActivated
	}
	return false
}

func (x *NewAppConfigRequest) GetAlertOnEmail() bool {
	if x != nil {
		return x.AlertOnEmail
	}
	return false
}

func (x *NewAppConfigRequest) GetAlertOnMobile() bool {
	if x != nil {
		return x.AlertOnMobile
	}
	return false
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x02, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x49, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x54, 0x68, 0x65, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x70, 0x70, 0x54, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x42, 0x69,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e, 0x4d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x4f, 0x6e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x70, 0x70,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x41, 0x70, 0x70, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x70, 0x70, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x70, 0x70, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x22, 0xeb, 0x01, 0x0a, 0x09, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x49, 0x64, 0x5f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x64,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x49, 0x64, 0x5f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49,
	0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x42, 0x69, 0x6f, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a,
	0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f,
	0x6e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x49, 0x64, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x49, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x49, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x42, 0x69, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x42, 0x69, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x4f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x4f, 0x6e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4f, 0x6e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x32,
	0x82, 0x01, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_proto_goTypes = []interface{}{
	(*AppConfigDetails)(nil),    // 0: proto.AppConfigDetails
	(*AppConfig)(nil),           // 1: proto.AppConfig
	(*NewAppConfigRequest)(nil), // 2: proto.NewAppConfigRequest
	(*StatusResponse)(nil),      // 3: proto.StatusResponse
}
var file_config_proto_depIdxs = []int32{
	2, // 0: proto.AppConfigService.Create:input_type -> proto.NewAppConfigRequest
	1, // 1: proto.AppConfigService.Details:input_type -> proto.AppConfig
	1, // 2: proto.AppConfigService.Create:output_type -> proto.AppConfig
	3, // 3: proto.AppConfigService.Details:output_type -> proto.StatusResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	file_generic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfigDetails); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfig); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewAppConfigRequest); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}

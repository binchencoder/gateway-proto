// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: data/data.proto

package data

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The unique service ID. Please sort in alphanumeric order.
//
// Services or apps under the same product should be placed in one section
// with the same naming prefix. For example, vexillary itself has 4 services
// so their service IDs could be defined as:
//
//   <blank line>
//   //### 2700-2709 Approval services.
//   <blank line>
//   APPROVAL_GRPC_SERVICE   = 2700;
//   APPROVAL_SCHEDULE_SERVICE = 27001;
//   <blank line>
//
type ServiceId int32

const (
	ServiceId_SERVICE_NONE ServiceId = 0
	//### 11 - 99: Infrastructure services
	ServiceId_VEXILLARY_SERVICE ServiceId = 11
	ServiceId_EASE_GATEWAY      ServiceId = 12
	ServiceId_CUSTOM_GATEWAY    ServiceId = 13
	ServiceId_SKYLB_FRONTEND    ServiceId = 14 // SkyLB managment UI.
	//### 2700-2709 Approval services.
	ServiceId_APPROVAL_GRPC_SERVICE     ServiceId = 2700
	ServiceId_APPROVAL_SCHEDULE_SERVICE ServiceId = 2701
	//### 2800-2819 Account system services.
	ServiceId_ACCOUNT_SERVICE ServiceId = 2801
	//### Test services begins from 2147470000.
	ServiceId_SHARED_TEST_CLIENT_SERVICE ServiceId = 2147470000 // public test client service.
	ServiceId_SHARED_TEST_SERVER_SERVICE ServiceId = 2147470001 // public test server service.
	ServiceId_CUSTOM_EASE_GATEWAY_TEST   ServiceId = 2147470006 // public custom ease-gateway service.
	///### for config server default settings.
	ServiceId__ ServiceId = 2147483647 // math.MaxInt32
)

// Enum value maps for ServiceId.
var (
	ServiceId_name = map[int32]string{
		0:          "SERVICE_NONE",
		11:         "VEXILLARY_SERVICE",
		12:         "EASE_GATEWAY",
		13:         "CUSTOM_GATEWAY",
		14:         "SKYLB_FRONTEND",
		2700:       "APPROVAL_GRPC_SERVICE",
		2701:       "APPROVAL_SCHEDULE_SERVICE",
		2801:       "ACCOUNT_SERVICE",
		2147470000: "SHARED_TEST_CLIENT_SERVICE",
		2147470001: "SHARED_TEST_SERVER_SERVICE",
		2147470006: "CUSTOM_EASE_GATEWAY_TEST",
		2147483647: "_",
	}
	ServiceId_value = map[string]int32{
		"SERVICE_NONE":               0,
		"VEXILLARY_SERVICE":          11,
		"EASE_GATEWAY":               12,
		"CUSTOM_GATEWAY":             13,
		"SKYLB_FRONTEND":             14,
		"APPROVAL_GRPC_SERVICE":      2700,
		"APPROVAL_SCHEDULE_SERVICE":  2701,
		"ACCOUNT_SERVICE":            2801,
		"SHARED_TEST_CLIENT_SERVICE": 2147470000,
		"SHARED_TEST_SERVER_SERVICE": 2147470001,
		"CUSTOM_EASE_GATEWAY_TEST":   2147470006,
		"_":                          2147483647,
	}
)

func (x ServiceId) Enum() *ServiceId {
	p := new(ServiceId)
	*p = x
	return p
}

func (x ServiceId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceId) Descriptor() protoreflect.EnumDescriptor {
	return file_data_data_proto_enumTypes[0].Descriptor()
}

func (ServiceId) Type() protoreflect.EnumType {
	return &file_data_data_proto_enumTypes[0]
}

func (x ServiceId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceId.Descriptor instead.
func (ServiceId) EnumDescriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0}
}

// AppType 定义应用类型.
type AppType int32

const (
	AppType_ALL      AppType = 0 // All app types.
	AppType_APPROVAL AppType = 1 // 审批应用.
)

// Enum value maps for AppType.
var (
	AppType_name = map[int32]string{
		0: "ALL",
		1: "APPROVAL",
	}
	AppType_value = map[string]int32{
		"ALL":      0,
		"APPROVAL": 1,
	}
)

func (x AppType) Enum() *AppType {
	p := new(AppType)
	*p = x
	return p
}

func (x AppType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppType) Descriptor() protoreflect.EnumDescriptor {
	return file_data_data_proto_enumTypes[1].Descriptor()
}

func (AppType) Type() protoreflect.EnumType {
	return &file_data_data_proto_enumTypes[1]
}

func (x AppType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppType.Descriptor instead.
func (AppType) EnumDescriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{1}
}

// Locale represents a specific geographical, political, or cultural region.
type Locale int32

const (
	Locale_LOCALE_NONE Locale = 0
	// ENGLISH
	Locale_EN Locale = 1
	// FRENCH
	Locale_FR Locale = 2
	// GERMAN
	Locale_DE Locale = 3
	// ITALIAN
	Locale_IT Locale = 4
	// JAPANESE
	Locale_JA Locale = 5
	// KOREAN
	Locale_CO Locale = 6
	// CHINESE
	Locale_ZH Locale = 7
	// SIMPLIFIED_CHINESE
	Locale_ZH_CN Locale = 101
	// TRADITIONAL_CHINESE
	Locale_ZH_TW Locale = 102
	// FRANCE
	Locale_FR_FR Locale = 103
	// GERMANY
	Locale_DE_DE Locale = 104
	// ITALY
	Locale_IT_IT Locale = 105
	// JAPAN
	Locale_JA_JP Locale = 106
	// KOREA
	Locale_KO_KR Locale = 107
	// UK
	Locale_EN_GB Locale = 108
	// US
	Locale_EN_US Locale = 109
	// CANADA
	Locale_EN_CA Locale = 110
	// CANADA_FRENCH
	Locale_FR_CA Locale = 111
)

// Enum value maps for Locale.
var (
	Locale_name = map[int32]string{
		0:   "LOCALE_NONE",
		1:   "EN",
		2:   "FR",
		3:   "DE",
		4:   "IT",
		5:   "JA",
		6:   "CO",
		7:   "ZH",
		101: "ZH_CN",
		102: "ZH_TW",
		103: "FR_FR",
		104: "DE_DE",
		105: "IT_IT",
		106: "JA_JP",
		107: "KO_KR",
		108: "EN_GB",
		109: "EN_US",
		110: "EN_CA",
		111: "FR_CA",
	}
	Locale_value = map[string]int32{
		"LOCALE_NONE": 0,
		"EN":          1,
		"FR":          2,
		"DE":          3,
		"IT":          4,
		"JA":          5,
		"CO":          6,
		"ZH":          7,
		"ZH_CN":       101,
		"ZH_TW":       102,
		"FR_FR":       103,
		"DE_DE":       104,
		"IT_IT":       105,
		"JA_JP":       106,
		"KO_KR":       107,
		"EN_GB":       108,
		"EN_US":       109,
		"EN_CA":       110,
		"FR_CA":       111,
	}
)

func (x Locale) Enum() *Locale {
	p := new(Locale)
	*p = x
	return p
}

func (x Locale) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Locale) Descriptor() protoreflect.EnumDescriptor {
	return file_data_data_proto_enumTypes[2].Descriptor()
}

func (Locale) Type() protoreflect.EnumType {
	return &file_data_data_proto_enumTypes[2]
}

func (x Locale) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Locale.Descriptor instead.
func (Locale) EnumDescriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{2}
}

// UserDetails provides core user information.
type UserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User's identity.
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// uid.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// cid.
	CompanyId string `protobuf:"bytes,3,opt,name=company_id,json=companyId,proto3" json:"company_id,omitempty"`
	// User's Locale.
	Locale Locale `protobuf:"varint,4,opt,name=locale,proto3,enum=data.Locale" json:"locale,omitempty"`
}

func (x *UserDetails) Reset() {
	*x = UserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetails) ProtoMessage() {}

func (x *UserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_data_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetails.ProtoReflect.Descriptor instead.
func (*UserDetails) Descriptor() ([]byte, []int) {
	return file_data_data_proto_rawDescGZIP(), []int{0}
}

func (x *UserDetails) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *UserDetails) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserDetails) GetCompanyId() string {
	if x != nil {
		return x.CompanyId
	}
	return ""
}

func (x *UserDetails) GetLocale() Locale {
	if x != nil {
		return x.Locale
	}
	return Locale_LOCALE_NONE
}

var File_data_data_proto protoreflect.FileDescriptor

var file_data_data_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8a, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x2a, 0xb5, 0x02, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x4f,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x56, 0x45, 0x58, 0x49, 0x4c, 0x4c, 0x41, 0x52,
	0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x45,
	0x41, 0x53, 0x45, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10, 0x0c, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x47, 0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x10,
	0x0d, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x4b, 0x59, 0x4c, 0x42, 0x5f, 0x46, 0x52, 0x4f, 0x4e, 0x54,
	0x45, 0x4e, 0x44, 0x10, 0x0e, 0x12, 0x1a, 0x0a, 0x15, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41,
	0x4c, 0x5f, 0x47, 0x52, 0x50, 0x43, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x8c,
	0x15, 0x12, 0x1e, 0x0a, 0x19, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x5f, 0x53, 0x43,
	0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x8d,
	0x15, 0x12, 0x14, 0x0a, 0x0f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0xf1, 0x15, 0x12, 0x22, 0x0a, 0x1a, 0x53, 0x48, 0x41, 0x52, 0x45,
	0x44, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0xb0, 0x95, 0xff, 0xff, 0x07, 0x12, 0x22, 0x0a, 0x1a, 0x53,
	0x48, 0x41, 0x52, 0x45, 0x44, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0xb1, 0x95, 0xff, 0xff, 0x07, 0x12,
	0x20, 0x0a, 0x18, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f, 0x45, 0x41, 0x53, 0x45, 0x5f, 0x47,
	0x41, 0x54, 0x45, 0x57, 0x41, 0x59, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0xb6, 0x95, 0xff, 0xff,
	0x07, 0x12, 0x09, 0x0a, 0x01, 0x5f, 0x10, 0xff, 0xff, 0xff, 0xff, 0x07, 0x2a, 0x20, 0x0a, 0x07,
	0x41, 0x70, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x41, 0x4c, 0x10, 0x01, 0x2a, 0xca,
	0x01, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x43,
	0x41, 0x4c, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x45, 0x4e,
	0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x46, 0x52, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x44, 0x45,
	0x10, 0x03, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x54, 0x10, 0x04, 0x12, 0x06, 0x0a, 0x02, 0x4a, 0x41,
	0x10, 0x05, 0x12, 0x06, 0x0a, 0x02, 0x43, 0x4f, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x5a, 0x48,
	0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x5a, 0x48, 0x5f, 0x43, 0x4e, 0x10, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x5a, 0x48, 0x5f, 0x54, 0x57, 0x10, 0x66, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x52, 0x5f, 0x46,
	0x52, 0x10, 0x67, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x10, 0x68, 0x12, 0x09,
	0x0a, 0x05, 0x49, 0x54, 0x5f, 0x49, 0x54, 0x10, 0x69, 0x12, 0x09, 0x0a, 0x05, 0x4a, 0x41, 0x5f,
	0x4a, 0x50, 0x10, 0x6a, 0x12, 0x09, 0x0a, 0x05, 0x4b, 0x4f, 0x5f, 0x4b, 0x52, 0x10, 0x6b, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x4e, 0x5f, 0x47, 0x42, 0x10, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e,
	0x5f, 0x55, 0x53, 0x10, 0x6d, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4e, 0x5f, 0x43, 0x41, 0x10, 0x6e,
	0x12, 0x09, 0x0a, 0x05, 0x46, 0x52, 0x5f, 0x43, 0x41, 0x10, 0x6f, 0x42, 0x54, 0x0a, 0x1a, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x69, 0x6e, 0x63, 0x68, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_data_proto_rawDescOnce sync.Once
	file_data_data_proto_rawDescData = file_data_data_proto_rawDesc
)

func file_data_data_proto_rawDescGZIP() []byte {
	file_data_data_proto_rawDescOnce.Do(func() {
		file_data_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_data_proto_rawDescData)
	})
	return file_data_data_proto_rawDescData
}

var file_data_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_data_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_data_data_proto_goTypes = []interface{}{
	(ServiceId)(0),      // 0: data.ServiceId
	(AppType)(0),        // 1: data.AppType
	(Locale)(0),         // 2: data.Locale
	(*UserDetails)(nil), // 3: data.UserDetails
}
var file_data_data_proto_depIdxs = []int32{
	2, // 0: data.UserDetails.locale:type_name -> data.Locale
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_data_data_proto_init() }
func file_data_data_proto_init() {
	if File_data_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetails); i {
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
			RawDescriptor: file_data_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_data_proto_goTypes,
		DependencyIndexes: file_data_data_proto_depIdxs,
		EnumInfos:         file_data_data_proto_enumTypes,
		MessageInfos:      file_data_data_proto_msgTypes,
	}.Build()
	File_data_data_proto = out.File
	file_data_data_proto_rawDesc = nil
	file_data_data_proto_goTypes = nil
	file_data_data_proto_depIdxs = nil
}

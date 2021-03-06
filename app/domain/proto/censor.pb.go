// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.3
// source: censor.proto

package censorPb

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

type RET_CODE int32

const (
	RET_CODE_SUCCESS       RET_CODE = 0
	RET_CODE_ERROR         RET_CODE = 1
	RET_CODE_INVALID_PARAM RET_CODE = 2
	RET_CODE_UID_TOO_MANY  RET_CODE = 3
	RET_CODE_REDIS_ERR     RET_CODE = 4
)

// Enum value maps for RET_CODE.
var (
	RET_CODE_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
		2: "INVALID_PARAM",
		3: "UID_TOO_MANY",
		4: "REDIS_ERR",
	}
	RET_CODE_value = map[string]int32{
		"SUCCESS":       0,
		"ERROR":         1,
		"INVALID_PARAM": 2,
		"UID_TOO_MANY":  3,
		"REDIS_ERR":     4,
	}
)

func (x RET_CODE) Enum() *RET_CODE {
	p := new(RET_CODE)
	*p = x
	return p
}

func (x RET_CODE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RET_CODE) Descriptor() protoreflect.EnumDescriptor {
	return file_censor_proto_enumTypes[0].Descriptor()
}

func (RET_CODE) Type() protoreflect.EnumType {
	return &file_censor_proto_enumTypes[0]
}

func (x RET_CODE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RET_CODE.Descriptor instead.
func (RET_CODE) EnumDescriptor() ([]byte, []int) {
	return file_censor_proto_rawDescGZIP(), []int{0}
}

type CensorTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CensorTextRequest) Reset() {
	*x = CensorTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_censor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CensorTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CensorTextRequest) ProtoMessage() {}

func (x *CensorTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_censor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CensorTextRequest.ProtoReflect.Descriptor instead.
func (*CensorTextRequest) Descriptor() ([]byte, []int) {
	return file_censor_proto_rawDescGZIP(), []int{0}
}

func (x *CensorTextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CensorImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *CensorImageRequest) Reset() {
	*x = CensorImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_censor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CensorImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CensorImageRequest) ProtoMessage() {}

func (x *CensorImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_censor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CensorImageRequest.ProtoReflect.Descriptor instead.
func (*CensorImageRequest) Descriptor() ([]byte, []int) {
	return file_censor_proto_rawDescGZIP(), []int{1}
}

func (x *CensorImageRequest) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type CensorResultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    RET_CODE      `protobuf:"varint,1,opt,name=code,proto3,enum=proto.RET_CODE" json:"code,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    *CensorResult `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CensorResultResponse) Reset() {
	*x = CensorResultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_censor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CensorResultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CensorResultResponse) ProtoMessage() {}

func (x *CensorResultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_censor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CensorResultResponse.ProtoReflect.Descriptor instead.
func (*CensorResultResponse) Descriptor() ([]byte, []int) {
	return file_censor_proto_rawDescGZIP(), []int{2}
}

func (x *CensorResultResponse) GetCode() RET_CODE {
	if x != nil {
		return x.Code
	}
	return RET_CODE_SUCCESS
}

func (x *CensorResultResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CensorResultResponse) GetData() *CensorResult {
	if x != nil {
		return x.Data
	}
	return nil
}

type CensorResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CensorContent   string         `protobuf:"bytes,1,opt,name=censor_content,json=censorContent,proto3" json:"censor_content,omitempty"`
	CensorType      int32          `protobuf:"varint,2,opt,name=censor_type,json=censorType,proto3" json:"censor_type,omitempty"`
	Suggestion      string         `protobuf:"bytes,3,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	InterceptStatus bool           `protobuf:"varint,4,opt,name=intercept_status,json=interceptStatus,proto3" json:"intercept_status,omitempty"`
	ReviewLabel     []*ReviewLabel `protobuf:"bytes,5,rep,name=review_label,json=reviewLabel,proto3" json:"review_label,omitempty"`
}

func (x *CensorResult) Reset() {
	*x = CensorResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_censor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CensorResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CensorResult) ProtoMessage() {}

func (x *CensorResult) ProtoReflect() protoreflect.Message {
	mi := &file_censor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CensorResult.ProtoReflect.Descriptor instead.
func (*CensorResult) Descriptor() ([]byte, []int) {
	return file_censor_proto_rawDescGZIP(), []int{3}
}

func (x *CensorResult) GetCensorContent() string {
	if x != nil {
		return x.CensorContent
	}
	return ""
}

func (x *CensorResult) GetCensorType() int32 {
	if x != nil {
		return x.CensorType
	}
	return 0
}

func (x *CensorResult) GetSuggestion() string {
	if x != nil {
		return x.Suggestion
	}
	return ""
}

func (x *CensorResult) GetInterceptStatus() bool {
	if x != nil {
		return x.InterceptStatus
	}
	return false
}

func (x *CensorResult) GetReviewLabel() []*ReviewLabel {
	if x != nil {
		return x.ReviewLabel
	}
	return nil
}

type ReviewLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReviewContent string  `protobuf:"bytes,1,opt,name=review_content,json=reviewContent,proto3" json:"review_content,omitempty"`
	Label         string  `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Rate          float64 `protobuf:"fixed64,3,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *ReviewLabel) Reset() {
	*x = ReviewLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_censor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviewLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewLabel) ProtoMessage() {}

func (x *ReviewLabel) ProtoReflect() protoreflect.Message {
	mi := &file_censor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewLabel.ProtoReflect.Descriptor instead.
func (*ReviewLabel) Descriptor() ([]byte, []int) {
	return file_censor_proto_rawDescGZIP(), []int{4}
}

func (x *ReviewLabel) GetReviewContent() string {
	if x != nil {
		return x.ReviewContent
	}
	return ""
}

func (x *ReviewLabel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ReviewLabel) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_censor_proto protoreflect.FileDescriptor

var file_censor_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2a,
	0x0a, 0x12, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x7e, 0x0a, 0x14, 0x43, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xd8, 0x01, 0x0a, 0x0c, 0x43,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35,
	0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x5e, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x72, 0x61, 0x74, 0x65, 0x2a, 0x56, 0x0a, 0x08, 0x52, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x10, 0x03, 0x12, 0x0d,
	0x0a, 0x09, 0x52, 0x45, 0x44, 0x49, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x10, 0x04, 0x32, 0xa3, 0x01,
	0x0a, 0x11, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b, 0x43, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x63, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x50, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_censor_proto_rawDescOnce sync.Once
	file_censor_proto_rawDescData = file_censor_proto_rawDesc
)

func file_censor_proto_rawDescGZIP() []byte {
	file_censor_proto_rawDescOnce.Do(func() {
		file_censor_proto_rawDescData = protoimpl.X.CompressGZIP(file_censor_proto_rawDescData)
	})
	return file_censor_proto_rawDescData
}

var file_censor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_censor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_censor_proto_goTypes = []interface{}{
	(RET_CODE)(0),                // 0: proto.RET_CODE
	(*CensorTextRequest)(nil),    // 1: proto.CensorTextRequest
	(*CensorImageRequest)(nil),   // 2: proto.CensorImageRequest
	(*CensorResultResponse)(nil), // 3: proto.CensorResultResponse
	(*CensorResult)(nil),         // 4: proto.CensorResult
	(*ReviewLabel)(nil),          // 5: proto.ReviewLabel
}
var file_censor_proto_depIdxs = []int32{
	0, // 0: proto.CensorResultResponse.code:type_name -> proto.RET_CODE
	4, // 1: proto.CensorResultResponse.data:type_name -> proto.CensorResult
	5, // 2: proto.CensorResult.review_label:type_name -> proto.ReviewLabel
	1, // 3: proto.CensorGrpcService.CensorText:input_type -> proto.CensorTextRequest
	2, // 4: proto.CensorGrpcService.CensorImage:input_type -> proto.CensorImageRequest
	3, // 5: proto.CensorGrpcService.CensorText:output_type -> proto.CensorResultResponse
	3, // 6: proto.CensorGrpcService.CensorImage:output_type -> proto.CensorResultResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_censor_proto_init() }
func file_censor_proto_init() {
	if File_censor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_censor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CensorTextRequest); i {
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
		file_censor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CensorImageRequest); i {
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
		file_censor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CensorResultResponse); i {
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
		file_censor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CensorResult); i {
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
		file_censor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviewLabel); i {
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
			RawDescriptor: file_censor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_censor_proto_goTypes,
		DependencyIndexes: file_censor_proto_depIdxs,
		EnumInfos:         file_censor_proto_enumTypes,
		MessageInfos:      file_censor_proto_msgTypes,
	}.Build()
	File_censor_proto = out.File
	file_censor_proto_rawDesc = nil
	file_censor_proto_goTypes = nil
	file_censor_proto_depIdxs = nil
}

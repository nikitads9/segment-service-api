// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.6
// source: segment.proto

package segment_service_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddSegmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slug string `protobuf:"bytes,1,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *AddSegmentRequest) Reset() {
	*x = AddSegmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSegmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSegmentRequest) ProtoMessage() {}

func (x *AddSegmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_segment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSegmentRequest.ProtoReflect.Descriptor instead.
func (*AddSegmentRequest) Descriptor() ([]byte, []int) {
	return file_segment_proto_rawDescGZIP(), []int{0}
}

func (x *AddSegmentRequest) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

type AddSegmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddSegmentResponse) Reset() {
	*x = AddSegmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSegmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSegmentResponse) ProtoMessage() {}

func (x *AddSegmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_segment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSegmentResponse.ProtoReflect.Descriptor instead.
func (*AddSegmentResponse) Descriptor() ([]byte, []int) {
	return file_segment_proto_rawDescGZIP(), []int{1}
}

func (x *AddSegmentResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RemoveSegmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RemoveSegmentRequest) Reset() {
	*x = RemoveSegmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSegmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSegmentRequest) ProtoMessage() {}

func (x *RemoveSegmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_segment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSegmentRequest.ProtoReflect.Descriptor instead.
func (*RemoveSegmentRequest) Descriptor() ([]byte, []int) {
	return file_segment_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveSegmentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ModifySegmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlugToAdd    []string `protobuf:"bytes,1,rep,name=slug_to_add,json=slugs-to-add,proto3" json:"slug_to_add,omitempty"`
	SlugToRemove []string `protobuf:"bytes,2,rep,name=slug_to_remove,json=slugs-to-remove,proto3" json:"slug_to_remove,omitempty"`
	Id           int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ModifySegmentsRequest) Reset() {
	*x = ModifySegmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySegmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySegmentsRequest) ProtoMessage() {}

func (x *ModifySegmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_segment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySegmentsRequest.ProtoReflect.Descriptor instead.
func (*ModifySegmentsRequest) Descriptor() ([]byte, []int) {
	return file_segment_proto_rawDescGZIP(), []int{3}
}

func (x *ModifySegmentsRequest) GetSlugToAdd() []string {
	if x != nil {
		return x.SlugToAdd
	}
	return nil
}

func (x *ModifySegmentsRequest) GetSlugToRemove() []string {
	if x != nil {
		return x.SlugToRemove
	}
	return nil
}

func (x *ModifySegmentsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSegmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSegmentsRequest) Reset() {
	*x = GetSegmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSegmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSegmentsRequest) ProtoMessage() {}

func (x *GetSegmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_segment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSegmentsRequest.ProtoReflect.Descriptor instead.
func (*GetSegmentsRequest) Descriptor() ([]byte, []int) {
	return file_segment_proto_rawDescGZIP(), []int{4}
}

func (x *GetSegmentsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetSegmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSegmentsResponse) Reset() {
	*x = GetSegmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_segment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSegmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSegmentsResponse) ProtoMessage() {}

func (x *GetSegmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_segment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSegmentsResponse.ProtoReflect.Descriptor instead.
func (*GetSegmentsResponse) Descriptor() ([]byte, []int) {
	return file_segment_proto_rawDescGZIP(), []int{5}
}

var File_segment_proto protoreflect.FileDescriptor

var file_segment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06,
	0x72, 0x04, 0x10, 0x01, 0x18, 0x14, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x22, 0x24, 0x0a, 0x12,
	0x41, 0x64, 0x64, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x2f, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x7c, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0b,
	0x73, 0x6c, 0x75, 0x67, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x6c, 0x75, 0x67, 0x73, 0x2d, 0x74, 0x6f, 0x2d, 0x61, 0x64, 0x64, 0x12,
	0x27, 0x0a, 0x0e, 0x73, 0x6c, 0x75, 0x67, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x6c, 0x75, 0x67, 0x73, 0x2d, 0x74,
	0x6f, 0x2d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfc, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x56, 0x31, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x76, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x64,
	0x64, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x29, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x2a, 0x14, 0x2f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x32, 0x82, 0x02, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x56,
	0x31, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x32, 0x0c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x81, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x12, 0x17, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x74, 0x2d, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x4e, 0x5a, 0x4c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x6b, 0x69, 0x74, 0x61,
	0x64, 0x73, 0x39, 0x2f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x3b, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_segment_proto_rawDescOnce sync.Once
	file_segment_proto_rawDescData = file_segment_proto_rawDesc
)

func file_segment_proto_rawDescGZIP() []byte {
	file_segment_proto_rawDescOnce.Do(func() {
		file_segment_proto_rawDescData = protoimpl.X.CompressGZIP(file_segment_proto_rawDescData)
	})
	return file_segment_proto_rawDescData
}

var file_segment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_segment_proto_goTypes = []interface{}{
	(*AddSegmentRequest)(nil),     // 0: segment.service.api.AddSegmentRequest
	(*AddSegmentResponse)(nil),    // 1: segment.service.api.AddSegmentResponse
	(*RemoveSegmentRequest)(nil),  // 2: segment.service.api.RemoveSegmentRequest
	(*ModifySegmentsRequest)(nil), // 3: segment.service.api.ModifySegmentsRequest
	(*GetSegmentsRequest)(nil),    // 4: segment.service.api.GetSegmentsRequest
	(*GetSegmentsResponse)(nil),   // 5: segment.service.api.GetSegmentsResponse
	(*emptypb.Empty)(nil),         // 6: google.protobuf.Empty
}
var file_segment_proto_depIdxs = []int32{
	0, // 0: segment.service.api.SegmentV1Service.AddSegment:input_type -> segment.service.api.AddSegmentRequest
	2, // 1: segment.service.api.SegmentV1Service.RemoveSegment:input_type -> segment.service.api.RemoveSegmentRequest
	3, // 2: segment.service.api.UserV1Service.ModifySegments:input_type -> segment.service.api.ModifySegmentsRequest
	4, // 3: segment.service.api.UserV1Service.GetSegments:input_type -> segment.service.api.GetSegmentsRequest
	1, // 4: segment.service.api.SegmentV1Service.AddSegment:output_type -> segment.service.api.AddSegmentResponse
	6, // 5: segment.service.api.SegmentV1Service.RemoveSegment:output_type -> google.protobuf.Empty
	6, // 6: segment.service.api.UserV1Service.ModifySegments:output_type -> google.protobuf.Empty
	5, // 7: segment.service.api.UserV1Service.GetSegments:output_type -> segment.service.api.GetSegmentsResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_segment_proto_init() }
func file_segment_proto_init() {
	if File_segment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_segment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSegmentRequest); i {
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
		file_segment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSegmentResponse); i {
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
		file_segment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSegmentRequest); i {
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
		file_segment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifySegmentsRequest); i {
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
		file_segment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSegmentsRequest); i {
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
		file_segment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSegmentsResponse); i {
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
			RawDescriptor: file_segment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_segment_proto_goTypes,
		DependencyIndexes: file_segment_proto_depIdxs,
		MessageInfos:      file_segment_proto_msgTypes,
	}.Build()
	File_segment_proto = out.File
	file_segment_proto_rawDesc = nil
	file_segment_proto_goTypes = nil
	file_segment_proto_depIdxs = nil
}

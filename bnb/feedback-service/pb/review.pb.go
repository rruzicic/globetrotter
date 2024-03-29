// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: review.proto

package pb

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

type HostReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedOn  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	ModifiedOn *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=modified_on,json=modifiedOn,proto3" json:"modified_on,omitempty"`
	DeletedOn  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deleted_on,json=deletedOn,proto3" json:"deleted_on,omitempty"`
	Rating     int32                  `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	UserId     string                 `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
	HostId     string                 `protobuf:"bytes,7,opt,name=hostId,proto3" json:"hostId,omitempty"`
}

func (x *HostReview) Reset() {
	*x = HostReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostReview) ProtoMessage() {}

func (x *HostReview) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostReview.ProtoReflect.Descriptor instead.
func (*HostReview) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{0}
}

func (x *HostReview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HostReview) GetCreatedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedOn
	}
	return nil
}

func (x *HostReview) GetModifiedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedOn
	}
	return nil
}

func (x *HostReview) GetDeletedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedOn
	}
	return nil
}

func (x *HostReview) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *HostReview) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *HostReview) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

type AccommodationReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedOn       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	ModifiedOn      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=modified_on,json=modifiedOn,proto3" json:"modified_on,omitempty"`
	DeletedOn       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deleted_on,json=deletedOn,proto3" json:"deleted_on,omitempty"`
	Rating          int32                  `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	UserId          string                 `protobuf:"bytes,6,opt,name=userId,proto3" json:"userId,omitempty"`
	AccommodationId string                 `protobuf:"bytes,7,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
}

func (x *AccommodationReview) Reset() {
	*x = AccommodationReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationReview) ProtoMessage() {}

func (x *AccommodationReview) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationReview.ProtoReflect.Descriptor instead.
func (*AccommodationReview) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{1}
}

func (x *AccommodationReview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccommodationReview) GetCreatedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedOn
	}
	return nil
}

func (x *AccommodationReview) GetModifiedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedOn
	}
	return nil
}

func (x *AccommodationReview) GetDeletedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedOn
	}
	return nil
}

func (x *AccommodationReview) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *AccommodationReview) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AccommodationReview) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

type RequestReviewById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestReviewById) Reset() {
	*x = RequestReviewById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReviewById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReviewById) ProtoMessage() {}

func (x *RequestReviewById) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReviewById.ProtoReflect.Descriptor instead.
func (*RequestReviewById) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{2}
}

func (x *RequestReviewById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RequestReviewsByUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestReviewsByUserId) Reset() {
	*x = RequestReviewsByUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReviewsByUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReviewsByUserId) ProtoMessage() {}

func (x *RequestReviewsByUserId) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReviewsByUserId.ProtoReflect.Descriptor instead.
func (*RequestReviewsByUserId) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{3}
}

func (x *RequestReviewsByUserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RequestReviewsByHostId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestReviewsByHostId) Reset() {
	*x = RequestReviewsByHostId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReviewsByHostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReviewsByHostId) ProtoMessage() {}

func (x *RequestReviewsByHostId) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReviewsByHostId.ProtoReflect.Descriptor instead.
func (*RequestReviewsByHostId) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{4}
}

func (x *RequestReviewsByHostId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RequestReviewsByAccommodationId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestReviewsByAccommodationId) Reset() {
	*x = RequestReviewsByAccommodationId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReviewsByAccommodationId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReviewsByAccommodationId) ProtoMessage() {}

func (x *RequestReviewsByAccommodationId) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReviewsByAccommodationId.ProtoReflect.Descriptor instead.
func (*RequestReviewsByAccommodationId) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{5}
}

func (x *RequestReviewsByAccommodationId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AvgRatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvgRating float32 `protobuf:"fixed32,1,opt,name=avgRating,proto3" json:"avgRating,omitempty"`
}

func (x *AvgRatingResponse) Reset() {
	*x = AvgRatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvgRatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvgRatingResponse) ProtoMessage() {}

func (x *AvgRatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvgRatingResponse.ProtoReflect.Descriptor instead.
func (*AvgRatingResponse) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{6}
}

func (x *AvgRatingResponse) GetAvgRating() float32 {
	if x != nil {
		return x.AvgRating
	}
	return 0
}

type RequestAvgRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestAvgRating) Reset() {
	*x = RequestAvgRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAvgRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAvgRating) ProtoMessage() {}

func (x *RequestAvgRating) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAvgRating.ProtoReflect.Descriptor instead.
func (*RequestAvgRating) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{7}
}

func (x *RequestAvgRating) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AvgRatingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId    string  `protobuf:"bytes,1,opt,name=hostId,proto3" json:"hostId,omitempty"`
	AvgRating float32 `protobuf:"fixed32,2,opt,name=avgRating,proto3" json:"avgRating,omitempty"`
}

func (x *AvgRatingEvent) Reset() {
	*x = AvgRatingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvgRatingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvgRatingEvent) ProtoMessage() {}

func (x *AvgRatingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvgRatingEvent.ProtoReflect.Descriptor instead.
func (*AvgRatingEvent) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{8}
}

func (x *AvgRatingEvent) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *AvgRatingEvent) GetAvgRating() float32 {
	if x != nil {
		return x.AvgRating
	}
	return 0
}

type EmptyReviewMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReviewMsg) Reset() {
	*x = EmptyReviewMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_review_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReviewMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReviewMsg) ProtoMessage() {}

func (x *EmptyReviewMsg) ProtoReflect() protoreflect.Message {
	mi := &file_review_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReviewMsg.ProtoReflect.Descriptor instead.
func (*EmptyReviewMsg) Descriptor() ([]byte, []int) {
	return file_review_proto_rawDescGZIP(), []int{9}
}

var File_review_proto protoreflect.FileDescriptor

var file_review_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x3b, 0x0a,
	0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0xb2, 0x02,
	0x0a, 0x13, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e,
	0x12, 0x3b, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x39, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x28, 0x0a, 0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x42, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x1f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31,
	0x0a, 0x11, 0x41, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x61, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x22, 0x22, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x76, 0x67, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0e, 0x41, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x61, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x10, 0x0a,
	0x0e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x32,
	0xe3, 0x05, 0x0a, 0x0f, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x79, 0x49, 0x64, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22,
	0x00, 0x12, 0x48, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x30, 0x01, 0x12, 0x48, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42, 0x79, 0x48,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x14, 0x43, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x76, 0x67, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4d, 0x73, 0x67, 0x1a, 0x17,
	0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6c, 0x0a, 0x28, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x1d, 0x43, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x46, 0x6f, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x41, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x41, 0x76, 0x67, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x72, 0x75, 0x7a, 0x69, 0x63, 0x69, 0x63, 0x2f, 0x67, 0x6c, 0x6f,
	0x62, 0x65, 0x74, 0x72, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x62, 0x6e, 0x62, 0x2f, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_review_proto_rawDescOnce sync.Once
	file_review_proto_rawDescData = file_review_proto_rawDesc
)

func file_review_proto_rawDescGZIP() []byte {
	file_review_proto_rawDescOnce.Do(func() {
		file_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_review_proto_rawDescData)
	})
	return file_review_proto_rawDescData
}

var file_review_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_review_proto_goTypes = []interface{}{
	(*HostReview)(nil),                      // 0: pb.HostReview
	(*AccommodationReview)(nil),             // 1: pb.AccommodationReview
	(*RequestReviewById)(nil),               // 2: pb.RequestReviewById
	(*RequestReviewsByUserId)(nil),          // 3: pb.RequestReviewsByUserId
	(*RequestReviewsByHostId)(nil),          // 4: pb.RequestReviewsByHostId
	(*RequestReviewsByAccommodationId)(nil), // 5: pb.RequestReviewsByAccommodationId
	(*AvgRatingResponse)(nil),               // 6: pb.AvgRatingResponse
	(*RequestAvgRating)(nil),                // 7: pb.RequestAvgRating
	(*AvgRatingEvent)(nil),                  // 8: pb.AvgRatingEvent
	(*EmptyReviewMsg)(nil),                  // 9: pb.EmptyReviewMsg
	(*timestamppb.Timestamp)(nil),           // 10: google.protobuf.Timestamp
}
var file_review_proto_depIdxs = []int32{
	10, // 0: pb.HostReview.created_on:type_name -> google.protobuf.Timestamp
	10, // 1: pb.HostReview.modified_on:type_name -> google.protobuf.Timestamp
	10, // 2: pb.HostReview.deleted_on:type_name -> google.protobuf.Timestamp
	10, // 3: pb.AccommodationReview.created_on:type_name -> google.protobuf.Timestamp
	10, // 4: pb.AccommodationReview.modified_on:type_name -> google.protobuf.Timestamp
	10, // 5: pb.AccommodationReview.deleted_on:type_name -> google.protobuf.Timestamp
	2,  // 6: pb.FeedbackService.GetHostReviewById:input_type -> pb.RequestReviewById
	3,  // 7: pb.FeedbackService.GetHostReviewsByUserId:input_type -> pb.RequestReviewsByUserId
	4,  // 8: pb.FeedbackService.GetHostReviewsByHostId:input_type -> pb.RequestReviewsByHostId
	7,  // 9: pb.FeedbackService.CalcAvgRatingForHost:input_type -> pb.RequestAvgRating
	9,  // 10: pb.FeedbackService.GetAllAccommodationReviews:input_type -> pb.EmptyReviewMsg
	2,  // 11: pb.FeedbackService.GetAccommodationReviewById:input_type -> pb.RequestReviewById
	3,  // 12: pb.FeedbackService.GetAccommodationReviewsByUserId:input_type -> pb.RequestReviewsByUserId
	5,  // 13: pb.FeedbackService.GetAccommodationReviewsByAccommodationId:input_type -> pb.RequestReviewsByAccommodationId
	7,  // 14: pb.FeedbackService.CalcAvgRatingForAccommodation:input_type -> pb.RequestAvgRating
	0,  // 15: pb.FeedbackService.GetHostReviewById:output_type -> pb.HostReview
	0,  // 16: pb.FeedbackService.GetHostReviewsByUserId:output_type -> pb.HostReview
	0,  // 17: pb.FeedbackService.GetHostReviewsByHostId:output_type -> pb.HostReview
	6,  // 18: pb.FeedbackService.CalcAvgRatingForHost:output_type -> pb.AvgRatingResponse
	1,  // 19: pb.FeedbackService.GetAllAccommodationReviews:output_type -> pb.AccommodationReview
	1,  // 20: pb.FeedbackService.GetAccommodationReviewById:output_type -> pb.AccommodationReview
	1,  // 21: pb.FeedbackService.GetAccommodationReviewsByUserId:output_type -> pb.AccommodationReview
	1,  // 22: pb.FeedbackService.GetAccommodationReviewsByAccommodationId:output_type -> pb.AccommodationReview
	6,  // 23: pb.FeedbackService.CalcAvgRatingForAccommodation:output_type -> pb.AvgRatingResponse
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_review_proto_init() }
func file_review_proto_init() {
	if File_review_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_review_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostReview); i {
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
		file_review_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationReview); i {
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
		file_review_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReviewById); i {
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
		file_review_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReviewsByUserId); i {
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
		file_review_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReviewsByHostId); i {
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
		file_review_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReviewsByAccommodationId); i {
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
		file_review_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvgRatingResponse); i {
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
		file_review_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAvgRating); i {
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
		file_review_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvgRatingEvent); i {
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
		file_review_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReviewMsg); i {
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
			RawDescriptor: file_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_review_proto_goTypes,
		DependencyIndexes: file_review_proto_depIdxs,
		MessageInfos:      file_review_proto_msgTypes,
	}.Build()
	File_review_proto = out.File
	file_review_proto_rawDesc = nil
	file_review_proto_goTypes = nil
	file_review_proto_depIdxs = nil
}

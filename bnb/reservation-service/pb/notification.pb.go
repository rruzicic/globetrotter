// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: notification.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ReservationNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId   string                 `protobuf:"bytes,1,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	UserId            string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	StartDate         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate           *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	NumOfGuests       int32                  `protobuf:"varint,5,opt,name=numOfGuests,proto3" json:"numOfGuests,omitempty"`
	IsApproved        bool                   `protobuf:"varint,6,opt,name=isApproved,proto3" json:"isApproved,omitempty"`
	TotalPrice        float32                `protobuf:"fixed32,7,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	AccommodationName string                 `protobuf:"bytes,8,opt,name=accommodationName,proto3" json:"accommodationName,omitempty"`
}

func (x *ReservationNotification) Reset() {
	*x = ReservationNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationNotification) ProtoMessage() {}

func (x *ReservationNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationNotification.ProtoReflect.Descriptor instead.
func (*ReservationNotification) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

func (x *ReservationNotification) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *ReservationNotification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationNotification) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *ReservationNotification) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *ReservationNotification) GetNumOfGuests() int32 {
	if x != nil {
		return x.NumOfGuests
	}
	return 0
}

func (x *ReservationNotification) GetIsApproved() bool {
	if x != nil {
		return x.IsApproved
	}
	return false
}

func (x *ReservationNotification) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *ReservationNotification) GetAccommodationName() string {
	if x != nil {
		return x.AccommodationName
	}
	return ""
}

type HostRatingNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RatedId string `protobuf:"bytes,1,opt,name=ratedId,proto3" json:"ratedId,omitempty"`
	RaterId string `protobuf:"bytes,2,opt,name=raterId,proto3" json:"raterId,omitempty"`
	Rating  int64  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *HostRatingNotification) Reset() {
	*x = HostRatingNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostRatingNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostRatingNotification) ProtoMessage() {}

func (x *HostRatingNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostRatingNotification.ProtoReflect.Descriptor instead.
func (*HostRatingNotification) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{1}
}

func (x *HostRatingNotification) GetRatedId() string {
	if x != nil {
		return x.RatedId
	}
	return ""
}

func (x *HostRatingNotification) GetRaterId() string {
	if x != nil {
		return x.RaterId
	}
	return ""
}

func (x *HostRatingNotification) GetRating() int64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type AccommodationRatingNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId           string `protobuf:"bytes,1,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	RatedId           string `protobuf:"bytes,2,opt,name=ratedId,proto3" json:"ratedId,omitempty"`
	RaterId           string `protobuf:"bytes,3,opt,name=raterId,proto3" json:"raterId,omitempty"`
	Rating            int64  `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
	AccommodationName string `protobuf:"bytes,5,opt,name=accommodationName,proto3" json:"accommodationName,omitempty"`
}

func (x *AccommodationRatingNotification) Reset() {
	*x = AccommodationRatingNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccommodationRatingNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccommodationRatingNotification) ProtoMessage() {}

func (x *AccommodationRatingNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccommodationRatingNotification.ProtoReflect.Descriptor instead.
func (*AccommodationRatingNotification) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{2}
}

func (x *AccommodationRatingNotification) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *AccommodationRatingNotification) GetRatedId() string {
	if x != nil {
		return x.RatedId
	}
	return ""
}

func (x *AccommodationRatingNotification) GetRaterId() string {
	if x != nil {
		return x.RaterId
	}
	return ""
}

func (x *AccommodationRatingNotification) GetRating() int64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *AccommodationRatingNotification) GetAccommodationName() string {
	if x != nil {
		return x.AccommodationName
	}
	return ""
}

type ReservationResponseNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	AccommodationId   string `protobuf:"bytes,2,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	AccommodationName string `protobuf:"bytes,3,opt,name=accommodationName,proto3" json:"accommodationName,omitempty"`
	Approved          bool   `protobuf:"varint,4,opt,name=approved,proto3" json:"approved,omitempty"`
}

func (x *ReservationResponseNotification) Reset() {
	*x = ReservationResponseNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationResponseNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationResponseNotification) ProtoMessage() {}

func (x *ReservationResponseNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationResponseNotification.ProtoReflect.Descriptor instead.
func (*ReservationResponseNotification) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{3}
}

func (x *ReservationResponseNotification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ReservationResponseNotification) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *ReservationResponseNotification) GetAccommodationName() string {
	if x != nil {
		return x.AccommodationName
	}
	return ""
}

func (x *ReservationResponseNotification) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

type HostStatusNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *HostStatusNotification) Reset() {
	*x = HostStatusNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostStatusNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostStatusNotification) ProtoMessage() {}

func (x *HostStatusNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostStatusNotification.ProtoReflect.Descriptor instead.
func (*HostStatusNotification) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{4}
}

func (x *HostStatusNotification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_notification_proto protoreflect.FileDescriptor

var file_notification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x02, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34,
	0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x4f, 0x66, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x4f, 0x66,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x41, 0x70,
	0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x64, 0x0a, 0x16, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x61, 0x74, 0x65, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xb5, 0x01, 0x0a, 0x1f, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x1f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x64, 0x22, 0x30, 0x0a, 0x16, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x32, 0xe9, 0x03, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x12,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x13, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64,
	0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x09, 0x48, 0x6f, 0x73, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x12, 0x41, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x54, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x11, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x72, 0x75, 0x7a, 0x69, 0x63, 0x69, 0x63, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x65, 0x74, 0x72, 0x6f,
	0x74, 0x74, 0x65, 0x72, 0x2f, 0x62, 0x6e, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_proto_rawDescOnce sync.Once
	file_notification_proto_rawDescData = file_notification_proto_rawDesc
)

func file_notification_proto_rawDescGZIP() []byte {
	file_notification_proto_rawDescOnce.Do(func() {
		file_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_proto_rawDescData)
	})
	return file_notification_proto_rawDescData
}

var file_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_notification_proto_goTypes = []interface{}{
	(*ReservationNotification)(nil),         // 0: pb.ReservationNotification
	(*HostRatingNotification)(nil),          // 1: pb.HostRatingNotification
	(*AccommodationRatingNotification)(nil), // 2: pb.AccommodationRatingNotification
	(*ReservationResponseNotification)(nil), // 3: pb.ReservationResponseNotification
	(*HostStatusNotification)(nil),          // 4: pb.HostStatusNotification
	(*timestamppb.Timestamp)(nil),           // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                   // 6: google.protobuf.Empty
}
var file_notification_proto_depIdxs = []int32{
	5, // 0: pb.ReservationNotification.startDate:type_name -> google.protobuf.Timestamp
	5, // 1: pb.ReservationNotification.endDate:type_name -> google.protobuf.Timestamp
	0, // 2: pb.NotificationService.ReservationCreated:input_type -> pb.ReservationNotification
	0, // 3: pb.NotificationService.ReservationCanceled:input_type -> pb.ReservationNotification
	1, // 4: pb.NotificationService.HostRated:input_type -> pb.HostRatingNotification
	2, // 5: pb.NotificationService.AccommodationRated:input_type -> pb.AccommodationRatingNotification
	3, // 6: pb.NotificationService.ReservationResponse:input_type -> pb.ReservationResponseNotification
	4, // 7: pb.NotificationService.HostStatusChanged:input_type -> pb.HostStatusNotification
	6, // 8: pb.NotificationService.ReservationCreated:output_type -> google.protobuf.Empty
	6, // 9: pb.NotificationService.ReservationCanceled:output_type -> google.protobuf.Empty
	6, // 10: pb.NotificationService.HostRated:output_type -> google.protobuf.Empty
	6, // 11: pb.NotificationService.AccommodationRated:output_type -> google.protobuf.Empty
	6, // 12: pb.NotificationService.ReservationResponse:output_type -> google.protobuf.Empty
	6, // 13: pb.NotificationService.HostStatusChanged:output_type -> google.protobuf.Empty
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_proto_init() }
func file_notification_proto_init() {
	if File_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationNotification); i {
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
		file_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostRatingNotification); i {
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
		file_notification_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccommodationRatingNotification); i {
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
		file_notification_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationResponseNotification); i {
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
		file_notification_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostStatusNotification); i {
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
			RawDescriptor: file_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notification_proto_goTypes,
		DependencyIndexes: file_notification_proto_depIdxs,
		MessageInfos:      file_notification_proto_msgTypes,
	}.Build()
	File_notification_proto = out.File
	file_notification_proto_rawDesc = nil
	file_notification_proto_goTypes = nil
	file_notification_proto_depIdxs = nil
}

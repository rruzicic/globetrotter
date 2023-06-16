// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: reservation.proto

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

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId string                 `protobuf:"bytes,1,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	UserId          string                 `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	StartDate       *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	NumOfGuests     int32                  `protobuf:"varint,5,opt,name=numOfGuests,proto3" json:"numOfGuests,omitempty"`
	IsApproved      bool                   `protobuf:"varint,6,opt,name=isApproved,proto3" json:"isApproved,omitempty"`
	TotalPrice      float32                `protobuf:"fixed32,7,opt,name=totalPrice,proto3" json:"totalPrice,omitempty"`
	Id              string                 `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *Reservation) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *Reservation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Reservation) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *Reservation) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

func (x *Reservation) GetNumOfGuests() int32 {
	if x != nil {
		return x.NumOfGuests
	}
	return 0
}

func (x *Reservation) GetIsApproved() bool {
	if x != nil {
		return x.IsApproved
	}
	return false
}

func (x *Reservation) GetTotalPrice() float32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

func (x *Reservation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RequestReservationById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestReservationById) Reset() {
	*x = RequestReservationById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReservationById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReservationById) ProtoMessage() {}

func (x *RequestReservationById) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReservationById.ProtoReflect.Descriptor instead.
func (*RequestReservationById) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *RequestReservationById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RequestReservationsByAccommodationId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestReservationsByAccommodationId) Reset() {
	*x = RequestReservationsByAccommodationId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReservationsByAccommodationId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReservationsByAccommodationId) ProtoMessage() {}

func (x *RequestReservationsByAccommodationId) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReservationsByAccommodationId.ProtoReflect.Descriptor instead.
func (*RequestReservationsByAccommodationId) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *RequestReservationsByAccommodationId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RequestUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RequestUserId) Reset() {
	*x = RequestUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestUserId) ProtoMessage() {}

func (x *RequestUserId) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestUserId.ProtoReflect.Descriptor instead.
func (*RequestUserId) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{3}
}

func (x *RequestUserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyResMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResMsg) Reset() {
	*x = EmptyResMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reservation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResMsg) ProtoMessage() {}

func (x *EmptyResMsg) ProtoReflect() protoreflect.Message {
	mi := &file_reservation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResMsg.ProtoReflect.Descriptor instead.
func (*EmptyResMsg) Descriptor() ([]byte, []int) {
	return file_reservation_proto_rawDescGZIP(), []int{4}
}

var File_reservation_proto protoreflect.FileDescriptor

var file_reservation_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x75,
	0x6d, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x6e, 0x75, 0x6d, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x36, 0x0a, 0x24, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f,
	0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x0d, 0x0a, 0x0b, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x4d, 0x73, 0x67, 0x32, 0xd5,
	0x03, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x4d, 0x73, 0x67, 0x1a, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x28, 0x2e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x4b, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x79, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a,
	0x1d, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x72, 0x75, 0x7a, 0x69, 0x63, 0x69, 0x63, 0x2f, 0x67, 0x6c,
	0x6f, 0x62, 0x65, 0x74, 0x72, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x62, 0x6e, 0x62, 0x2f, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_reservation_proto_rawDescOnce sync.Once
	file_reservation_proto_rawDescData = file_reservation_proto_rawDesc
)

func file_reservation_proto_rawDescGZIP() []byte {
	file_reservation_proto_rawDescOnce.Do(func() {
		file_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_reservation_proto_rawDescData)
	})
	return file_reservation_proto_rawDescData
}

var file_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_reservation_proto_goTypes = []interface{}{
	(*Reservation)(nil),                          // 0: pb.Reservation
	(*RequestReservationById)(nil),               // 1: pb.RequestReservationById
	(*RequestReservationsByAccommodationId)(nil), // 2: pb.RequestReservationsByAccommodationId
	(*RequestUserId)(nil),                        // 3: pb.RequestUserId
	(*EmptyResMsg)(nil),                          // 4: pb.EmptyResMsg
	(*timestamppb.Timestamp)(nil),                // 5: google.protobuf.Timestamp
}
var file_reservation_proto_depIdxs = []int32{
	5, // 0: pb.Reservation.startDate:type_name -> google.protobuf.Timestamp
	5, // 1: pb.Reservation.endDate:type_name -> google.protobuf.Timestamp
	4, // 2: pb.ReservationService.GetAllReservations:input_type -> pb.EmptyResMsg
	1, // 3: pb.ReservationService.GetReservationById:input_type -> pb.RequestReservationById
	2, // 4: pb.ReservationService.GetReservationsByAccommodationId:input_type -> pb.RequestReservationsByAccommodationId
	3, // 5: pb.ReservationService.GetActiveReservationsByUser:input_type -> pb.RequestUserId
	3, // 6: pb.ReservationService.GetFutureActiveReservationsByHost:input_type -> pb.RequestUserId
	3, // 7: pb.ReservationService.GetFinishedReservationsByUser:input_type -> pb.RequestUserId
	0, // 8: pb.ReservationService.GetAllReservations:output_type -> pb.Reservation
	0, // 9: pb.ReservationService.GetReservationById:output_type -> pb.Reservation
	0, // 10: pb.ReservationService.GetReservationsByAccommodationId:output_type -> pb.Reservation
	0, // 11: pb.ReservationService.GetActiveReservationsByUser:output_type -> pb.Reservation
	0, // 12: pb.ReservationService.GetFutureActiveReservationsByHost:output_type -> pb.Reservation
	0, // 13: pb.ReservationService.GetFinishedReservationsByUser:output_type -> pb.Reservation
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_reservation_proto_init() }
func file_reservation_proto_init() {
	if File_reservation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reservation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reservation); i {
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
		file_reservation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReservationById); i {
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
		file_reservation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReservationsByAccommodationId); i {
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
		file_reservation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestUserId); i {
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
		file_reservation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResMsg); i {
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
			RawDescriptor: file_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_reservation_proto_goTypes,
		DependencyIndexes: file_reservation_proto_depIdxs,
		MessageInfos:      file_reservation_proto_msgTypes,
	}.Build()
	File_reservation_proto = out.File
	file_reservation_proto_rawDesc = nil
	file_reservation_proto_goTypes = nil
	file_reservation_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: recommendation.proto

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

type GraphAccommodation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Location string  `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Price    float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	MongoId  string  `protobuf:"bytes,4,opt,name=mongoId,proto3" json:"mongoId,omitempty"`
}

func (x *GraphAccommodation) Reset() {
	*x = GraphAccommodation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphAccommodation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphAccommodation) ProtoMessage() {}

func (x *GraphAccommodation) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphAccommodation.ProtoReflect.Descriptor instead.
func (*GraphAccommodation) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{0}
}

func (x *GraphAccommodation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphAccommodation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *GraphAccommodation) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GraphAccommodation) GetMongoId() string {
	if x != nil {
		return x.MongoId
	}
	return ""
}

type GraphReservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MongoId              string                 `protobuf:"bytes,1,opt,name=mongoId,proto3" json:"mongoId,omitempty"`
	UserMongoId          string                 `protobuf:"bytes,2,opt,name=userMongoId,proto3" json:"userMongoId,omitempty"`
	AccommodationMongoId string                 `protobuf:"bytes,3,opt,name=accommodationMongoId,proto3" json:"accommodationMongoId,omitempty"`
	ReservationEnd       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=reservationEnd,proto3" json:"reservationEnd,omitempty"`
}

func (x *GraphReservation) Reset() {
	*x = GraphReservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphReservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphReservation) ProtoMessage() {}

func (x *GraphReservation) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphReservation.ProtoReflect.Descriptor instead.
func (*GraphReservation) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{1}
}

func (x *GraphReservation) GetMongoId() string {
	if x != nil {
		return x.MongoId
	}
	return ""
}

func (x *GraphReservation) GetUserMongoId() string {
	if x != nil {
		return x.UserMongoId
	}
	return ""
}

func (x *GraphReservation) GetAccommodationMongoId() string {
	if x != nil {
		return x.AccommodationMongoId
	}
	return ""
}

func (x *GraphReservation) GetReservationEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.ReservationEnd
	}
	return nil
}

type GraphReview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value                int32  `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	MongoId              string `protobuf:"bytes,2,opt,name=mongoId,proto3" json:"mongoId,omitempty"`
	UserMongoId          string `protobuf:"bytes,3,opt,name=userMongoId,proto3" json:"userMongoId,omitempty"`
	AccommodationMongoId string `protobuf:"bytes,4,opt,name=accommodationMongoId,proto3" json:"accommodationMongoId,omitempty"`
}

func (x *GraphReview) Reset() {
	*x = GraphReview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphReview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphReview) ProtoMessage() {}

func (x *GraphReview) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphReview.ProtoReflect.Descriptor instead.
func (*GraphReview) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{2}
}

func (x *GraphReview) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *GraphReview) GetMongoId() string {
	if x != nil {
		return x.MongoId
	}
	return ""
}

func (x *GraphReview) GetUserMongoId() string {
	if x != nil {
		return x.UserMongoId
	}
	return ""
}

func (x *GraphReview) GetAccommodationMongoId() string {
	if x != nil {
		return x.AccommodationMongoId
	}
	return ""
}

type GraphUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MongoId string `protobuf:"bytes,2,opt,name=mongoId,proto3" json:"mongoId,omitempty"`
}

func (x *GraphUser) Reset() {
	*x = GraphUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphUser) ProtoMessage() {}

func (x *GraphUser) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphUser.ProtoReflect.Descriptor instead.
func (*GraphUser) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{3}
}

func (x *GraphUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GraphUser) GetMongoId() string {
	if x != nil {
		return x.MongoId
	}
	return ""
}

type GraphEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GraphEmpty) Reset() {
	*x = GraphEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommendation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphEmpty) ProtoMessage() {}

func (x *GraphEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_recommendation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphEmpty.ProtoReflect.Descriptor instead.
func (*GraphEmpty) Descriptor() ([]byte, []int) {
	return file_recommendation_proto_rawDescGZIP(), []int{4}
}

var File_recommendation_proto protoreflect.FileDescriptor

var file_recommendation_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x12, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x49,
	0x64, 0x22, 0xc6, 0x01, 0x0a, 0x10, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x67, 0x6f,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x14, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x14,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e,
	0x67, 0x6f, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64,
	0x22, 0x39, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xbf, 0x05, 0x0a, 0x1d, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x44, 0x42, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x41, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x72, 0x75, 0x7a, 0x69, 0x63,
	0x69, 0x63, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x65, 0x74, 0x72, 0x6f, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x62, 0x6e, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recommendation_proto_rawDescOnce sync.Once
	file_recommendation_proto_rawDescData = file_recommendation_proto_rawDesc
)

func file_recommendation_proto_rawDescGZIP() []byte {
	file_recommendation_proto_rawDescOnce.Do(func() {
		file_recommendation_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommendation_proto_rawDescData)
	})
	return file_recommendation_proto_rawDescData
}

var file_recommendation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_recommendation_proto_goTypes = []interface{}{
	(*GraphAccommodation)(nil),    // 0: pb.GraphAccommodation
	(*GraphReservation)(nil),      // 1: pb.GraphReservation
	(*GraphReview)(nil),           // 2: pb.GraphReview
	(*GraphUser)(nil),             // 3: pb.GraphUser
	(*GraphEmpty)(nil),            // 4: pb.GraphEmpty
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_recommendation_proto_depIdxs = []int32{
	5,  // 0: pb.GraphReservation.reservationEnd:type_name -> google.protobuf.Timestamp
	0,  // 1: pb.RecommendationServiceDBEvents.CreateAccommodation:input_type -> pb.GraphAccommodation
	3,  // 2: pb.RecommendationServiceDBEvents.CreateUser:input_type -> pb.GraphUser
	1,  // 3: pb.RecommendationServiceDBEvents.CreateReservation:input_type -> pb.GraphReservation
	2,  // 4: pb.RecommendationServiceDBEvents.CreateReview:input_type -> pb.GraphReview
	0,  // 5: pb.RecommendationServiceDBEvents.DeleteAccommodation:input_type -> pb.GraphAccommodation
	3,  // 6: pb.RecommendationServiceDBEvents.DeleteUser:input_type -> pb.GraphUser
	1,  // 7: pb.RecommendationServiceDBEvents.DeleteReservation:input_type -> pb.GraphReservation
	2,  // 8: pb.RecommendationServiceDBEvents.DeleteReview:input_type -> pb.GraphReview
	0,  // 9: pb.RecommendationServiceDBEvents.UpdateAccommodation:input_type -> pb.GraphAccommodation
	3,  // 10: pb.RecommendationServiceDBEvents.UpdateUser:input_type -> pb.GraphUser
	1,  // 11: pb.RecommendationServiceDBEvents.UpdateReservation:input_type -> pb.GraphReservation
	2,  // 12: pb.RecommendationServiceDBEvents.UpdateReview:input_type -> pb.GraphReview
	4,  // 13: pb.RecommendationServiceDBEvents.CreateAccommodation:output_type -> pb.GraphEmpty
	4,  // 14: pb.RecommendationServiceDBEvents.CreateUser:output_type -> pb.GraphEmpty
	4,  // 15: pb.RecommendationServiceDBEvents.CreateReservation:output_type -> pb.GraphEmpty
	4,  // 16: pb.RecommendationServiceDBEvents.CreateReview:output_type -> pb.GraphEmpty
	4,  // 17: pb.RecommendationServiceDBEvents.DeleteAccommodation:output_type -> pb.GraphEmpty
	4,  // 18: pb.RecommendationServiceDBEvents.DeleteUser:output_type -> pb.GraphEmpty
	4,  // 19: pb.RecommendationServiceDBEvents.DeleteReservation:output_type -> pb.GraphEmpty
	4,  // 20: pb.RecommendationServiceDBEvents.DeleteReview:output_type -> pb.GraphEmpty
	4,  // 21: pb.RecommendationServiceDBEvents.UpdateAccommodation:output_type -> pb.GraphEmpty
	4,  // 22: pb.RecommendationServiceDBEvents.UpdateUser:output_type -> pb.GraphEmpty
	4,  // 23: pb.RecommendationServiceDBEvents.UpdateReservation:output_type -> pb.GraphEmpty
	4,  // 24: pb.RecommendationServiceDBEvents.UpdateReview:output_type -> pb.GraphEmpty
	13, // [13:25] is the sub-list for method output_type
	1,  // [1:13] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_recommendation_proto_init() }
func file_recommendation_proto_init() {
	if File_recommendation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recommendation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphAccommodation); i {
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
		file_recommendation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphReservation); i {
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
		file_recommendation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphReview); i {
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
		file_recommendation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphUser); i {
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
		file_recommendation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphEmpty); i {
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
			RawDescriptor: file_recommendation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recommendation_proto_goTypes,
		DependencyIndexes: file_recommendation_proto_depIdxs,
		MessageInfos:      file_recommendation_proto_msgTypes,
	}.Build()
	File_recommendation_proto = out.File
	file_recommendation_proto_rawDesc = nil
	file_recommendation_proto_goTypes = nil
	file_recommendation_proto_depIdxs = nil
}

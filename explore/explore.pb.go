// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.0
// source: explore/explore.proto

package explore

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

type Gender int32

const (
	Gender_GENDER_MALE   Gender = 0
	Gender_GENDER_FEMALE Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_MALE",
		1: "GENDER_FEMALE",
	}
	Gender_value = map[string]int32{
		"GENDER_MALE":   0,
		"GENDER_FEMALE": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_explore_explore_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_explore_explore_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_explore_explore_proto_rawDescGZIP(), []int{0}
}

type LikedYou int32

const (
	LikedYou_LIKED_YOU_NEW    LikedYou = 0
	LikedYou_LIKED_YOU_SWIPED LikedYou = 1
)

// Enum value maps for LikedYou.
var (
	LikedYou_name = map[int32]string{
		0: "LIKED_YOU_NEW",
		1: "LIKED_YOU_SWIPED",
	}
	LikedYou_value = map[string]int32{
		"LIKED_YOU_NEW":    0,
		"LIKED_YOU_SWIPED": 1,
	}
)

func (x LikedYou) Enum() *LikedYou {
	p := new(LikedYou)
	*p = x
	return p
}

func (x LikedYou) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LikedYou) Descriptor() protoreflect.EnumDescriptor {
	return file_explore_explore_proto_enumTypes[1].Descriptor()
}

func (LikedYou) Type() protoreflect.EnumType {
	return &file_explore_explore_proto_enumTypes[1]
}

func (x LikedYou) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LikedYou.Descriptor instead.
func (LikedYou) EnumDescriptor() ([]byte, []int) {
	return file_explore_explore_proto_rawDescGZIP(), []int{1}
}

type ExploreProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp         uint32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MarriageProfileId uint32 `protobuf:"varint,2,opt,name=marriage_profile_id,json=marriageProfileId,proto3" json:"marriage_profile_id,omitempty"`
}

func (x *ExploreProfile) Reset() {
	*x = ExploreProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explore_explore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExploreProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExploreProfile) ProtoMessage() {}

func (x *ExploreProfile) ProtoReflect() protoreflect.Message {
	mi := &file_explore_explore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExploreProfile.ProtoReflect.Descriptor instead.
func (*ExploreProfile) Descriptor() ([]byte, []int) {
	return file_explore_explore_proto_rawDescGZIP(), []int{0}
}

func (x *ExploreProfile) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ExploreProfile) GetMarriageProfileId() uint32 {
	if x != nil {
		return x.MarriageProfileId
	}
	return 0
}

type LikedYouRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gender            Gender   `protobuf:"varint,1,opt,name=gender,proto3,enum=explore.Gender" json:"gender,omitempty"`
	Limit             uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Filter            LikedYou `protobuf:"varint,3,opt,name=filter,proto3,enum=explore.LikedYou" json:"filter,omitempty"`
	MarriageProfileId uint32   `protobuf:"varint,4,opt,name=marriage_profile_id,json=marriageProfileId,proto3" json:"marriage_profile_id,omitempty"`
}

func (x *LikedYouRequest) Reset() {
	*x = LikedYouRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explore_explore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikedYouRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikedYouRequest) ProtoMessage() {}

func (x *LikedYouRequest) ProtoReflect() protoreflect.Message {
	mi := &file_explore_explore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikedYouRequest.ProtoReflect.Descriptor instead.
func (*LikedYouRequest) Descriptor() ([]byte, []int) {
	return file_explore_explore_proto_rawDescGZIP(), []int{1}
}

func (x *LikedYouRequest) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GENDER_MALE
}

func (x *LikedYouRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *LikedYouRequest) GetFilter() LikedYou {
	if x != nil {
		return x.Filter
	}
	return LikedYou_LIKED_YOU_NEW
}

func (x *LikedYouRequest) GetMarriageProfileId() uint32 {
	if x != nil {
		return x.MarriageProfileId
	}
	return 0
}

type LikedYouResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Profiles []*ExploreProfile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *LikedYouResponse) Reset() {
	*x = LikedYouResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explore_explore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikedYouResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikedYouResponse) ProtoMessage() {}

func (x *LikedYouResponse) ProtoReflect() protoreflect.Message {
	mi := &file_explore_explore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikedYouResponse.ProtoReflect.Descriptor instead.
func (*LikedYouResponse) Descriptor() ([]byte, []int) {
	return file_explore_explore_proto_rawDescGZIP(), []int{2}
}

func (x *LikedYouResponse) GetProfiles() []*ExploreProfile {
	if x != nil {
		return x.Profiles
	}
	return nil
}

type PutSwipeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorMarriageProfileId     uint32 `protobuf:"varint,1,opt,name=actor_marriage_profile_id,json=actorMarriageProfileId,proto3" json:"actor_marriage_profile_id,omitempty"`
	RecipientMarriageProfileId uint32 `protobuf:"varint,2,opt,name=recipient_marriage_profile_id,json=recipientMarriageProfileId,proto3" json:"recipient_marriage_profile_id,omitempty"`
	ActorGender                Gender `protobuf:"varint,3,opt,name=actor_gender,json=actorGender,proto3,enum=explore.Gender" json:"actor_gender,omitempty"`
	Timestamp                  uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Like                       bool   `protobuf:"varint,5,opt,name=like,proto3" json:"like,omitempty"` // True if the actor liked on the recipient, false if they have passed
}

func (x *PutSwipeRequest) Reset() {
	*x = PutSwipeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explore_explore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutSwipeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutSwipeRequest) ProtoMessage() {}

func (x *PutSwipeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_explore_explore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutSwipeRequest.ProtoReflect.Descriptor instead.
func (*PutSwipeRequest) Descriptor() ([]byte, []int) {
	return file_explore_explore_proto_rawDescGZIP(), []int{3}
}

func (x *PutSwipeRequest) GetActorMarriageProfileId() uint32 {
	if x != nil {
		return x.ActorMarriageProfileId
	}
	return 0
}

func (x *PutSwipeRequest) GetRecipientMarriageProfileId() uint32 {
	if x != nil {
		return x.RecipientMarriageProfileId
	}
	return 0
}

func (x *PutSwipeRequest) GetActorGender() Gender {
	if x != nil {
		return x.ActorGender
	}
	return Gender_GENDER_MALE
}

func (x *PutSwipeRequest) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PutSwipeRequest) GetLike() bool {
	if x != nil {
		return x.Like
	}
	return false
}

type PutSwipeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutSwipeResponse) Reset() {
	*x = PutSwipeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_explore_explore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutSwipeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutSwipeResponse) ProtoMessage() {}

func (x *PutSwipeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_explore_explore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutSwipeResponse.ProtoReflect.Descriptor instead.
func (*PutSwipeResponse) Descriptor() ([]byte, []int) {
	return file_explore_explore_proto_rawDescGZIP(), []int{4}
}

var File_explore_explore_proto protoreflect.FileDescriptor

var file_explore_explore_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x22, 0x5e, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d,
	0x61, 0x72, 0x72, 0x69, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0xab, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69,
	0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x13, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x61, 0x72,
	0x72, 0x69, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x47,
	0x0a, 0x10, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x45,
	0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0xf5, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x74, 0x53,
	0x77, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x4d, 0x61, 0x72, 0x72, 0x69, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x72, 0x72, 0x69, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1a, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x72, 0x72, 0x69, 0x61, 0x67, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0c, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0f, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6b, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x22,
	0x12, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x53, 0x77, 0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2a, 0x2c, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a,
	0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x11,
	0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10,
	0x01, 0x2a, 0x33, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x12, 0x11, 0x0a,
	0x0d, 0x4c, 0x49, 0x4b, 0x45, 0x44, 0x5f, 0x59, 0x4f, 0x55, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x4c, 0x49, 0x4b, 0x45, 0x44, 0x5f, 0x59, 0x4f, 0x55, 0x5f, 0x53, 0x57,
	0x49, 0x50, 0x45, 0x44, 0x10, 0x01, 0x32, 0x92, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x59, 0x6f, 0x75, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e,
	0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59, 0x6f, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x59,
	0x6f, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x50, 0x75,
	0x74, 0x53, 0x77, 0x69, 0x70, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65,
	0x2e, 0x50, 0x75, 0x74, 0x53, 0x77, 0x69, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x75, 0x74, 0x53, 0x77,
	0x69, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x28, 0x5a, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6c, 0x69, 0x6f, 0x6e, 0x2f,
	0x6d, 0x75, 0x7a, 0x7a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2f, 0x65, 0x78,
	0x70, 0x6c, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_explore_explore_proto_rawDescOnce sync.Once
	file_explore_explore_proto_rawDescData = file_explore_explore_proto_rawDesc
)

func file_explore_explore_proto_rawDescGZIP() []byte {
	file_explore_explore_proto_rawDescOnce.Do(func() {
		file_explore_explore_proto_rawDescData = protoimpl.X.CompressGZIP(file_explore_explore_proto_rawDescData)
	})
	return file_explore_explore_proto_rawDescData
}

var file_explore_explore_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_explore_explore_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_explore_explore_proto_goTypes = []interface{}{
	(Gender)(0),              // 0: explore.Gender
	(LikedYou)(0),            // 1: explore.LikedYou
	(*ExploreProfile)(nil),   // 2: explore.ExploreProfile
	(*LikedYouRequest)(nil),  // 3: explore.LikedYouRequest
	(*LikedYouResponse)(nil), // 4: explore.LikedYouResponse
	(*PutSwipeRequest)(nil),  // 5: explore.PutSwipeRequest
	(*PutSwipeResponse)(nil), // 6: explore.PutSwipeResponse
}
var file_explore_explore_proto_depIdxs = []int32{
	0, // 0: explore.LikedYouRequest.gender:type_name -> explore.Gender
	1, // 1: explore.LikedYouRequest.filter:type_name -> explore.LikedYou
	2, // 2: explore.LikedYouResponse.profiles:type_name -> explore.ExploreProfile
	0, // 3: explore.PutSwipeRequest.actor_gender:type_name -> explore.Gender
	3, // 4: explore.ExploreService.LikedYou:input_type -> explore.LikedYouRequest
	5, // 5: explore.ExploreService.PutSwipe:input_type -> explore.PutSwipeRequest
	4, // 6: explore.ExploreService.LikedYou:output_type -> explore.LikedYouResponse
	6, // 7: explore.ExploreService.PutSwipe:output_type -> explore.PutSwipeResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_explore_explore_proto_init() }
func file_explore_explore_proto_init() {
	if File_explore_explore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_explore_explore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExploreProfile); i {
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
		file_explore_explore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikedYouRequest); i {
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
		file_explore_explore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikedYouResponse); i {
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
		file_explore_explore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutSwipeRequest); i {
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
		file_explore_explore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutSwipeResponse); i {
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
			RawDescriptor: file_explore_explore_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_explore_explore_proto_goTypes,
		DependencyIndexes: file_explore_explore_proto_depIdxs,
		EnumInfos:         file_explore_explore_proto_enumTypes,
		MessageInfos:      file_explore_explore_proto_msgTypes,
	}.Build()
	File_explore_explore_proto = out.File
	file_explore_explore_proto_rawDesc = nil
	file_explore_explore_proto_goTypes = nil
	file_explore_explore_proto_depIdxs = nil
}

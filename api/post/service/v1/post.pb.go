// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: api/post/service/v1/post.proto

package v1

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

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text   string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_post_service_v1_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_post_service_v1_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_api_post_service_v1_post_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreatePostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePostReply) Reset() {
	*x = CreatePostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_post_service_v1_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostReply) ProtoMessage() {}

func (x *CreatePostReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_post_service_v1_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostReply.ProtoReflect.Descriptor instead.
func (*CreatePostReply) Descriptor() ([]byte, []int) {
	return file_api_post_service_v1_post_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPostByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostByIDRequest) Reset() {
	*x = GetPostByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_post_service_v1_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIDRequest) ProtoMessage() {}

func (x *GetPostByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_post_service_v1_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIDRequest.ProtoReflect.Descriptor instead.
func (*GetPostByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_post_service_v1_post_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPostByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Text   string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *GetPostByIDReply) Reset() {
	*x = GetPostByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_post_service_v1_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIDReply) ProtoMessage() {}

func (x *GetPostByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_post_service_v1_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIDReply.ProtoReflect.Descriptor instead.
func (*GetPostByIDReply) Descriptor() ([]byte, []int) {
	return file_api_post_service_v1_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostByIDReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPostByIDReply) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetPostByIDReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetPostByIDReply) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetPostsByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostsByUserIDRequest) Reset() {
	*x = GetPostsByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_post_service_v1_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsByUserIDRequest) ProtoMessage() {}

func (x *GetPostsByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_post_service_v1_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetPostsByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_api_post_service_v1_post_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostsByUserIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPostsByUserIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*GetPostsByUserIDReply_Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostsByUserIDReply) Reset() {
	*x = GetPostsByUserIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_post_service_v1_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsByUserIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsByUserIDReply) ProtoMessage() {}

func (x *GetPostsByUserIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_post_service_v1_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsByUserIDReply.ProtoReflect.Descriptor instead.
func (*GetPostsByUserIDReply) Descriptor() ([]byte, []int) {
	return file_api_post_service_v1_post_proto_rawDescGZIP(), []int{5}
}

func (x *GetPostsByUserIDReply) GetPosts() []*GetPostsByUserIDReply_Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type GetPostsByUserIDReply_Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Text   string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *GetPostsByUserIDReply_Post) Reset() {
	*x = GetPostsByUserIDReply_Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_post_service_v1_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsByUserIDReply_Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsByUserIDReply_Post) ProtoMessage() {}

func (x *GetPostsByUserIDReply_Post) ProtoReflect() protoreflect.Message {
	mi := &file_api_post_service_v1_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsByUserIDReply_Post.ProtoReflect.Descriptor instead.
func (*GetPostsByUserIDReply_Post) Descriptor() ([]byte, []int) {
	return file_api_post_service_v1_post_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetPostsByUserIDReply_Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetPostsByUserIDReply_Post) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetPostsByUserIDReply_Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetPostsByUserIDReply_Post) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_api_post_service_v1_post_proto protoreflect.FileDescriptor

var file_api_post_service_v1_post_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x55, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x21, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x29, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x45, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x1a, 0x58, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x32, 0xb5, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x5c, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x3e, 0x0a, 0x13, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x25, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_post_service_v1_post_proto_rawDescOnce sync.Once
	file_api_post_service_v1_post_proto_rawDescData = file_api_post_service_v1_post_proto_rawDesc
)

func file_api_post_service_v1_post_proto_rawDescGZIP() []byte {
	file_api_post_service_v1_post_proto_rawDescOnce.Do(func() {
		file_api_post_service_v1_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_post_service_v1_post_proto_rawDescData)
	})
	return file_api_post_service_v1_post_proto_rawDescData
}

var file_api_post_service_v1_post_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_post_service_v1_post_proto_goTypes = []interface{}{
	(*CreatePostRequest)(nil),          // 0: api.post.service.v1.CreatePostRequest
	(*CreatePostReply)(nil),            // 1: api.post.service.v1.CreatePostReply
	(*GetPostByIDRequest)(nil),         // 2: api.post.service.v1.GetPostByIDRequest
	(*GetPostByIDReply)(nil),           // 3: api.post.service.v1.GetPostByIDReply
	(*GetPostsByUserIDRequest)(nil),    // 4: api.post.service.v1.GetPostsByUserIDRequest
	(*GetPostsByUserIDReply)(nil),      // 5: api.post.service.v1.GetPostsByUserIDReply
	(*GetPostsByUserIDReply_Post)(nil), // 6: api.post.service.v1.GetPostsByUserIDReply.Post
}
var file_api_post_service_v1_post_proto_depIdxs = []int32{
	6, // 0: api.post.service.v1.GetPostsByUserIDReply.posts:type_name -> api.post.service.v1.GetPostsByUserIDReply.Post
	0, // 1: api.post.service.v1.Post.CreatePost:input_type -> api.post.service.v1.CreatePostRequest
	2, // 2: api.post.service.v1.Post.GetPostByID:input_type -> api.post.service.v1.GetPostByIDRequest
	4, // 3: api.post.service.v1.Post.GetPostsByUserID:input_type -> api.post.service.v1.GetPostsByUserIDRequest
	1, // 4: api.post.service.v1.Post.CreatePost:output_type -> api.post.service.v1.CreatePostReply
	3, // 5: api.post.service.v1.Post.GetPostByID:output_type -> api.post.service.v1.GetPostByIDReply
	5, // 6: api.post.service.v1.Post.GetPostsByUserID:output_type -> api.post.service.v1.GetPostsByUserIDReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_post_service_v1_post_proto_init() }
func file_api_post_service_v1_post_proto_init() {
	if File_api_post_service_v1_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_post_service_v1_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_api_post_service_v1_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostReply); i {
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
		file_api_post_service_v1_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIDRequest); i {
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
		file_api_post_service_v1_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIDReply); i {
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
		file_api_post_service_v1_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsByUserIDRequest); i {
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
		file_api_post_service_v1_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsByUserIDReply); i {
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
		file_api_post_service_v1_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsByUserIDReply_Post); i {
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
			RawDescriptor: file_api_post_service_v1_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_post_service_v1_post_proto_goTypes,
		DependencyIndexes: file_api_post_service_v1_post_proto_depIdxs,
		MessageInfos:      file_api_post_service_v1_post_proto_msgTypes,
	}.Build()
	File_api_post_service_v1_post_proto = out.File
	file_api_post_service_v1_post_proto_rawDesc = nil
	file_api_post_service_v1_post_proto_goTypes = nil
	file_api_post_service_v1_post_proto_depIdxs = nil
}

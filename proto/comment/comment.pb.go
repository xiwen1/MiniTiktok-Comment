// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: comment.proto

package comment

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CommentActionRequest_CommentActionTypeCode int32

const (
	CommentActionRequest_PUBLISH CommentActionRequest_CommentActionTypeCode = 0
	CommentActionRequest_DELETE  CommentActionRequest_CommentActionTypeCode = 1
)

// Enum value maps for CommentActionRequest_CommentActionTypeCode.
var (
	CommentActionRequest_CommentActionTypeCode_name = map[int32]string{
		0: "PUBLISH",
		1: "DELETE",
	}
	CommentActionRequest_CommentActionTypeCode_value = map[string]int32{
		"PUBLISH": 0,
		"DELETE":  1,
	}
)

func (x CommentActionRequest_CommentActionTypeCode) Enum() *CommentActionRequest_CommentActionTypeCode {
	p := new(CommentActionRequest_CommentActionTypeCode)
	*p = x
	return p
}

func (x CommentActionRequest_CommentActionTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentActionRequest_CommentActionTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_proto_enumTypes[0].Descriptor()
}

func (CommentActionRequest_CommentActionTypeCode) Type() protoreflect.EnumType {
	return &file_comment_proto_enumTypes[0]
}

func (x CommentActionRequest_CommentActionTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentActionRequest_CommentActionTypeCode.Descriptor instead.
func (CommentActionRequest_CommentActionTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1, 0}
}

type CommentActionResponse_CommentStatusTypeCode int32

const (
	CommentActionResponse_SUCCESS CommentActionResponse_CommentStatusTypeCode = 0
	CommentActionResponse_FAIL    CommentActionResponse_CommentStatusTypeCode = 1
)

// Enum value maps for CommentActionResponse_CommentStatusTypeCode.
var (
	CommentActionResponse_CommentStatusTypeCode_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	CommentActionResponse_CommentStatusTypeCode_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x CommentActionResponse_CommentStatusTypeCode) Enum() *CommentActionResponse_CommentStatusTypeCode {
	p := new(CommentActionResponse_CommentStatusTypeCode)
	*p = x
	return p
}

func (x CommentActionResponse_CommentStatusTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentActionResponse_CommentStatusTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_proto_enumTypes[1].Descriptor()
}

func (CommentActionResponse_CommentStatusTypeCode) Type() protoreflect.EnumType {
	return &file_comment_proto_enumTypes[1]
}

func (x CommentActionResponse_CommentStatusTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentActionResponse_CommentStatusTypeCode.Descriptor instead.
func (CommentActionResponse_CommentStatusTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2, 0}
}

type CommentListResponse_ListStatusTypeCode int32

const (
	CommentListResponse_SUCCESS CommentListResponse_ListStatusTypeCode = 0
	CommentListResponse_FAIL    CommentListResponse_ListStatusTypeCode = 1
)

// Enum value maps for CommentListResponse_ListStatusTypeCode.
var (
	CommentListResponse_ListStatusTypeCode_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	CommentListResponse_ListStatusTypeCode_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x CommentListResponse_ListStatusTypeCode) Enum() *CommentListResponse_ListStatusTypeCode {
	p := new(CommentListResponse_ListStatusTypeCode)
	*p = x
	return p
}

func (x CommentListResponse_ListStatusTypeCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentListResponse_ListStatusTypeCode) Descriptor() protoreflect.EnumDescriptor {
	return file_comment_proto_enumTypes[2].Descriptor()
}

func (CommentListResponse_ListStatusTypeCode) Type() protoreflect.EnumType {
	return &file_comment_proto_enumTypes[2]
}

func (x CommentListResponse_ListStatusTypeCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentListResponse_ListStatusTypeCode.Descriptor instead.
func (CommentListResponse_ListStatusTypeCode) EnumDescriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{4, 0}
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 视频评论id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 评论用户的信息
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// 评论内容
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// 评论发布日期, 格式为 mm-dd
	CreateDate string `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

type CommentActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户鉴权token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 视频id
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	// 状态码, 如上
	ActionType CommentActionRequest_CommentActionTypeCode `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3,enum=mini_tiktok.proto.comment.CommentActionRequest_CommentActionTypeCode" json:"action_type,omitempty"`
	// 用户评论内容, 当action_type的值为PUBLISH时使用
	CommentText string `protobuf:"bytes,4,opt,name=comment_text,json=commentText,proto3" json:"comment_text,omitempty"`
	// 需要删除的评论的id, 当action_type的值为DELETE时使用
	CommentId int64 `protobuf:"varint,5,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
}

func (x *CommentActionRequest) Reset() {
	*x = CommentActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentActionRequest) ProtoMessage() {}

func (x *CommentActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentActionRequest.ProtoReflect.Descriptor instead.
func (*CommentActionRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{1}
}

func (x *CommentActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentActionRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CommentActionRequest) GetActionType() CommentActionRequest_CommentActionTypeCode {
	if x != nil {
		return x.ActionType
	}
	return CommentActionRequest_PUBLISH
}

func (x *CommentActionRequest) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

func (x *CommentActionRequest) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type CommentActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码, 如上
	StatusCode CommentActionResponse_CommentStatusTypeCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=mini_tiktok.proto.comment.CommentActionResponse_CommentStatusTypeCode" json:"status_code,omitempty"`
	// 返回状态描述
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// 评论成功则返回评论内容
	Comment *Comment `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentActionResponse) Reset() {
	*x = CommentActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentActionResponse) ProtoMessage() {}

func (x *CommentActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentActionResponse.ProtoReflect.Descriptor instead.
func (*CommentActionResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{2}
}

func (x *CommentActionResponse) GetStatusCode() CommentActionResponse_CommentStatusTypeCode {
	if x != nil {
		return x.StatusCode
	}
	return CommentActionResponse_SUCCESS
}

func (x *CommentActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CommentActionResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户鉴权token
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// 视频id
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
}

func (x *CommentListRequest) Reset() {
	*x = CommentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListRequest) ProtoMessage() {}

func (x *CommentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListRequest.ProtoReflect.Descriptor instead.
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CommentListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentListRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type CommentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 状态码, 如上
	StatusCode CommentListResponse_ListStatusTypeCode `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=mini_tiktok.proto.comment.CommentListResponse_ListStatusTypeCode" json:"status_code,omitempty"`
	// 返回状态信息
	StatusMsg string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	// 评论列表
	CommentList []*Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"`
}

func (x *CommentListResponse) Reset() {
	*x = CommentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResponse) ProtoMessage() {}

func (x *CommentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResponse.ProtoReflect.Descriptor instead.
func (*CommentListResponse) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CommentListResponse) GetStatusCode() CommentListResponse_ListStatusTypeCode {
	if x != nil {
		return x.StatusCode
	}
	return CommentListResponse_SUCCESS
}

func (x *CommentListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CommentListResponse) GetCommentList() []*Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 用户id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 用户名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 关注总数
	FollowCount uint32 `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`
	// 粉丝总数
	FollowerCount uint32 `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`
	// true-已关注，false-未关注
	IsFollow bool `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_comment_proto_rawDescGZIP(), []int{5}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() uint32 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() uint32 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

var File_comment_proto protoreflect.FileDescriptor

var file_comment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x22, 0xa3, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x66,
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x22, 0x8d, 0x02, 0x0a, 0x15, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x3c, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x15, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x22, 0x8c, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x45, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x2b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01,
	0x22, 0x91, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x32, 0xff, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xb9, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f,
	0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x3f, 0x22, 0x3a, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67,
	0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2d, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x2d, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x3a,
	0x01, 0x2a, 0x12, 0xb1, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x43, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3d, 0x22, 0x38, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62,
	0x61, 0x77, 0x6c, 0x69, 0x6e, 0x67, 0x2d, 0x6d, 0x69, 0x6e, 0x69, 0x64, 0x6f, 0x75, 0x79, 0x69,
	0x6e, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x67,
	0x72, 0x70, 0x63, 0x3a, 0x01, 0x2a, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comment_proto_rawDescOnce sync.Once
	file_comment_proto_rawDescData = file_comment_proto_rawDesc
)

func file_comment_proto_rawDescGZIP() []byte {
	file_comment_proto_rawDescOnce.Do(func() {
		file_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_comment_proto_rawDescData)
	})
	return file_comment_proto_rawDescData
}

var file_comment_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_comment_proto_goTypes = []interface{}{
	(CommentActionRequest_CommentActionTypeCode)(0),  // 0: mini_tiktok.proto.comment.CommentActionRequest.CommentActionTypeCode
	(CommentActionResponse_CommentStatusTypeCode)(0), // 1: mini_tiktok.proto.comment.CommentActionResponse.CommentStatusTypeCode
	(CommentListResponse_ListStatusTypeCode)(0),      // 2: mini_tiktok.proto.comment.CommentListResponse.ListStatusTypeCode
	(*Comment)(nil),               // 3: mini_tiktok.proto.comment.Comment
	(*CommentActionRequest)(nil),  // 4: mini_tiktok.proto.comment.CommentActionRequest
	(*CommentActionResponse)(nil), // 5: mini_tiktok.proto.comment.CommentActionResponse
	(*CommentListRequest)(nil),    // 6: mini_tiktok.proto.comment.CommentListRequest
	(*CommentListResponse)(nil),   // 7: mini_tiktok.proto.comment.CommentListResponse
	(*User)(nil),                  // 8: mini_tiktok.proto.comment.User
}
var file_comment_proto_depIdxs = []int32{
	8, // 0: mini_tiktok.proto.comment.Comment.user:type_name -> mini_tiktok.proto.comment.User
	0, // 1: mini_tiktok.proto.comment.CommentActionRequest.action_type:type_name -> mini_tiktok.proto.comment.CommentActionRequest.CommentActionTypeCode
	1, // 2: mini_tiktok.proto.comment.CommentActionResponse.status_code:type_name -> mini_tiktok.proto.comment.CommentActionResponse.CommentStatusTypeCode
	3, // 3: mini_tiktok.proto.comment.CommentActionResponse.comment:type_name -> mini_tiktok.proto.comment.Comment
	2, // 4: mini_tiktok.proto.comment.CommentListResponse.status_code:type_name -> mini_tiktok.proto.comment.CommentListResponse.ListStatusTypeCode
	3, // 5: mini_tiktok.proto.comment.CommentListResponse.comment_list:type_name -> mini_tiktok.proto.comment.Comment
	4, // 6: mini_tiktok.proto.comment.CommentAction.CommentAction:input_type -> mini_tiktok.proto.comment.CommentActionRequest
	6, // 7: mini_tiktok.proto.comment.CommentAction.CommentList:input_type -> mini_tiktok.proto.comment.CommentListRequest
	5, // 8: mini_tiktok.proto.comment.CommentAction.CommentAction:output_type -> mini_tiktok.proto.comment.CommentActionResponse
	7, // 9: mini_tiktok.proto.comment.CommentAction.CommentList:output_type -> mini_tiktok.proto.comment.CommentListResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_comment_proto_init() }
func file_comment_proto_init() {
	if File_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentActionRequest); i {
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
		file_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentActionResponse); i {
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
		file_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListRequest); i {
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
		file_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResponse); i {
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
		file_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_comment_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comment_proto_goTypes,
		DependencyIndexes: file_comment_proto_depIdxs,
		EnumInfos:         file_comment_proto_enumTypes,
		MessageInfos:      file_comment_proto_msgTypes,
	}.Build()
	File_comment_proto = out.File
	file_comment_proto_rawDesc = nil
	file_comment_proto_goTypes = nil
	file_comment_proto_depIdxs = nil
}

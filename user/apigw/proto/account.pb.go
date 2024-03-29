// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: proto/account.proto

package proto

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

type ReqSignin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ReqSignin) Reset() {
	*x = ReqSignin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSignin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSignin) ProtoMessage() {}

func (x *ReqSignin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSignin.ProtoReflect.Descriptor instead.
func (*ReqSignin) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{0}
}

func (x *ReqSignin) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ReqSignin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResSignin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Uid   int64  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ResSignin) Reset() {
	*x = ResSignin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResSignin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResSignin) ProtoMessage() {}

func (x *ResSignin) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResSignin.ProtoReflect.Descriptor instead.
func (*ResSignin) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{1}
}

func (x *ResSignin) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ResSignin) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ReqSignup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ReqSignup) Reset() {
	*x = ReqSignup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSignup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSignup) ProtoMessage() {}

func (x *ReqSignup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSignup.ProtoReflect.Descriptor instead.
func (*ReqSignup) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{2}
}

func (x *ReqSignup) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ReqSignup) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ResSignup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ResSignup) Reset() {
	*x = ResSignup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResSignup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResSignup) ProtoMessage() {}

func (x *ResSignup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResSignup.ProtoReflect.Descriptor instead.
func (*ResSignup) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{3}
}

func (x *ResSignup) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ResSignup) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type ReqGetUserFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ReqGetUserFile) Reset() {
	*x = ReqGetUserFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGetUserFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetUserFile) ProtoMessage() {}

func (x *ReqGetUserFile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetUserFile.ProtoReflect.Descriptor instead.
func (*ReqGetUserFile) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{4}
}

func (x *ReqGetUserFile) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ResGetUserFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userfile []*UserFile `protobuf:"bytes,1,rep,name=userfile,proto3" json:"userfile,omitempty"`
}

func (x *ResGetUserFile) Reset() {
	*x = ResGetUserFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResGetUserFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetUserFile) ProtoMessage() {}

func (x *ResGetUserFile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetUserFile.ProtoReflect.Descriptor instead.
func (*ResGetUserFile) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{5}
}

func (x *ResGetUserFile) GetUserfile() []*UserFile {
	if x != nil {
		return x.Userfile
	}
	return nil
}

type UserFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid   int64  `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Filehash string `protobuf:"bytes,3,opt,name=filehash,proto3" json:"filehash,omitempty"`
	Filename string `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	Filesize int64  `protobuf:"varint,5,opt,name=filesize,proto3" json:"filesize,omitempty"`
}

func (x *UserFile) Reset() {
	*x = UserFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFile) ProtoMessage() {}

func (x *UserFile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFile.ProtoReflect.Descriptor instead.
func (*UserFile) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{6}
}

func (x *UserFile) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserFile) GetUserid() int64 {
	if x != nil {
		return x.Userid
	}
	return 0
}

func (x *UserFile) GetFilehash() string {
	if x != nil {
		return x.Filehash
	}
	return ""
}

func (x *UserFile) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *UserFile) GetFilesize() int64 {
	if x != nil {
		return x.Filesize
	}
	return 0
}

type ReqReName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filehash string `protobuf:"bytes,1,opt,name=filehash,proto3" json:"filehash,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Uid      int64  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ReqReName) Reset() {
	*x = ReqReName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqReName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqReName) ProtoMessage() {}

func (x *ReqReName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqReName.ProtoReflect.Descriptor instead.
func (*ReqReName) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{7}
}

func (x *ReqReName) GetFilehash() string {
	if x != nil {
		return x.Filehash
	}
	return ""
}

func (x *ReqReName) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ReqReName) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ResReName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok      bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResReName) Reset() {
	*x = ResReName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResReName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResReName) ProtoMessage() {}

func (x *ResReName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResReName.ProtoReflect.Descriptor instead.
func (*ResReName) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{8}
}

func (x *ResReName) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *ResReName) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ReqUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ReqUserInfo) Reset() {
	*x = ReqUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUserInfo) ProtoMessage() {}

func (x *ReqUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUserInfo.ProtoReflect.Descriptor instead.
func (*ReqUserInfo) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{9}
}

func (x *ReqUserInfo) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type MetaFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Filehash string `protobuf:"bytes,2,opt,name=filehash,proto3" json:"filehash,omitempty"`
	Filename string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Filesize int64  `protobuf:"varint,4,opt,name=filesize,proto3" json:"filesize,omitempty"`
	Location string `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *MetaFile) Reset() {
	*x = MetaFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaFile) ProtoMessage() {}

func (x *MetaFile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaFile.ProtoReflect.Descriptor instead.
func (*MetaFile) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{10}
}

func (x *MetaFile) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MetaFile) GetFilehash() string {
	if x != nil {
		return x.Filehash
	}
	return ""
}

func (x *MetaFile) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *MetaFile) GetFilesize() int64 {
	if x != nil {
		return x.Filesize
	}
	return 0
}

func (x *MetaFile) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type ResUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ResUserInfo) Reset() {
	*x = ResUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_account_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResUserInfo) ProtoMessage() {}

func (x *ResUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_account_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResUserInfo.ProtoReflect.Descriptor instead.
func (*ResUserInfo) Descriptor() ([]byte, []int) {
	return file_proto_account_proto_rawDescGZIP(), []int{11}
}

func (x *ResUserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResUserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_proto_account_proto protoreflect.FileDescriptor

var file_proto_account_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x09,
	0x52, 0x65, 0x71, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x33, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x53, 0x69, 0x67,
	0x6e, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x39, 0x0a, 0x09, 0x52,
	0x65, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2b, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x22, 0x55, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x52, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x09, 0x52, 0x65, 0x73,
	0x52, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x1f, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x22, 0x8a, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x39,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x9f, 0x02, 0x0a, 0x0e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x71, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x11,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x41, 0x6c,
	0x6c, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x0a, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x52, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_account_proto_rawDescOnce sync.Once
	file_proto_account_proto_rawDescData = file_proto_account_proto_rawDesc
)

func file_proto_account_proto_rawDescGZIP() []byte {
	file_proto_account_proto_rawDescOnce.Do(func() {
		file_proto_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_account_proto_rawDescData)
	})
	return file_proto_account_proto_rawDescData
}

var file_proto_account_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_account_proto_goTypes = []interface{}{
	(*ReqSignin)(nil),      // 0: proto.ReqSignin
	(*ResSignin)(nil),      // 1: proto.ResSignin
	(*ReqSignup)(nil),      // 2: proto.ReqSignup
	(*ResSignup)(nil),      // 3: proto.ResSignup
	(*ReqGetUserFile)(nil), // 4: proto.ReqGetUserFile
	(*ResGetUserFile)(nil), // 5: proto.ResGetUserFile
	(*UserFile)(nil),       // 6: proto.UserFile
	(*ReqReName)(nil),      // 7: proto.ReqReName
	(*ResReName)(nil),      // 8: proto.ResReName
	(*ReqUserInfo)(nil),    // 9: proto.ReqUserInfo
	(*MetaFile)(nil),       // 10: proto.MetaFile
	(*ResUserInfo)(nil),    // 11: proto.ResUserInfo
}
var file_proto_account_proto_depIdxs = []int32{
	6,  // 0: proto.ResGetUserFile.userfile:type_name -> proto.UserFile
	0,  // 1: proto.AccountService.Signin:input_type -> proto.ReqSignin
	2,  // 2: proto.AccountService.Signup:input_type -> proto.ReqSignup
	4,  // 3: proto.AccountService.SelectUserFileAll:input_type -> proto.ReqGetUserFile
	7,  // 4: proto.AccountService.RenameFile:input_type -> proto.ReqReName
	9,  // 5: proto.AccountService.UserInfo:input_type -> proto.ReqUserInfo
	1,  // 6: proto.AccountService.Signin:output_type -> proto.ResSignin
	3,  // 7: proto.AccountService.Signup:output_type -> proto.ResSignup
	5,  // 8: proto.AccountService.SelectUserFileAll:output_type -> proto.ResGetUserFile
	8,  // 9: proto.AccountService.RenameFile:output_type -> proto.ResReName
	11, // 10: proto.AccountService.UserInfo:output_type -> proto.ResUserInfo
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_account_proto_init() }
func file_proto_account_proto_init() {
	if File_proto_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSignin); i {
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
		file_proto_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResSignin); i {
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
		file_proto_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSignup); i {
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
		file_proto_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResSignup); i {
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
		file_proto_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGetUserFile); i {
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
		file_proto_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResGetUserFile); i {
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
		file_proto_account_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFile); i {
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
		file_proto_account_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqReName); i {
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
		file_proto_account_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResReName); i {
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
		file_proto_account_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUserInfo); i {
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
		file_proto_account_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaFile); i {
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
		file_proto_account_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResUserInfo); i {
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
			RawDescriptor: file_proto_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_account_proto_goTypes,
		DependencyIndexes: file_proto_account_proto_depIdxs,
		MessageInfos:      file_proto_account_proto_msgTypes,
	}.Build()
	File_proto_account_proto = out.File
	file_proto_account_proto_rawDesc = nil
	file_proto_account_proto_goTypes = nil
	file_proto_account_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api.proto

package api

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

type RequestVoteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term         int32 `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	CandidateID  int32 `protobuf:"varint,2,opt,name=CandidateID,proto3" json:"CandidateID,omitempty"`
	LastLogIndex int32 `protobuf:"varint,3,opt,name=LastLogIndex,proto3" json:"LastLogIndex,omitempty"`
	LastLogTerm  int32 `protobuf:"varint,4,opt,name=LastLogTerm,proto3" json:"LastLogTerm,omitempty"`
}

func (x *RequestVoteArgs) Reset() {
	*x = RequestVoteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteArgs) ProtoMessage() {}

func (x *RequestVoteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteArgs.ProtoReflect.Descriptor instead.
func (*RequestVoteArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *RequestVoteArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteArgs) GetCandidateID() int32 {
	if x != nil {
		return x.CandidateID
	}
	return 0
}

func (x *RequestVoteArgs) GetLastLogIndex() int32 {
	if x != nil {
		return x.LastLogIndex
	}
	return 0
}

func (x *RequestVoteArgs) GetLastLogTerm() int32 {
	if x != nil {
		return x.LastLogTerm
	}
	return 0
}

type RequestVoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term        int32 `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	VoteGranted bool  `protobuf:"varint,2,opt,name=VoteGranted,proto3" json:"VoteGranted,omitempty"`
}

func (x *RequestVoteReply) Reset() {
	*x = RequestVoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestVoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestVoteReply) ProtoMessage() {}

func (x *RequestVoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestVoteReply.ProtoReflect.Descriptor instead.
func (*RequestVoteReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *RequestVoteReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *RequestVoteReply) GetVoteGranted() bool {
	if x != nil {
		return x.VoteGranted
	}
	return false
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term   int32  `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	Index  int32  `protobuf:"varint,2,opt,name=Index,proto3" json:"Index,omitempty"`
	Entry  []byte `protobuf:"bytes,3,opt,name=Entry,proto3" json:"Entry,omitempty"`
	Status bool   `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *Log) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *Log) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Log) GetEntry() []byte {
	if x != nil {
		return x.Entry
	}
	return nil
}

func (x *Log) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type AppendEntriesArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term           int32  `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	LeaderID       int32  `protobuf:"varint,2,opt,name=LeaderID,proto3" json:"LeaderID,omitempty"`
	PrevLogIndex   int32  `protobuf:"varint,3,opt,name=PrevLogIndex,proto3" json:"PrevLogIndex,omitempty"`
	PrevLogTerm    int32  `protobuf:"varint,4,opt,name=PrevLogTerm,proto3" json:"PrevLogTerm,omitempty"`
	LeaderCommitID int32  `protobuf:"varint,5,opt,name=LeaderCommitID,proto3" json:"LeaderCommitID,omitempty"`
	Logs           []*Log `protobuf:"bytes,6,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (x *AppendEntriesArgs) Reset() {
	*x = AppendEntriesArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesArgs) ProtoMessage() {}

func (x *AppendEntriesArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesArgs.ProtoReflect.Descriptor instead.
func (*AppendEntriesArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *AppendEntriesArgs) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesArgs) GetLeaderID() int32 {
	if x != nil {
		return x.LeaderID
	}
	return 0
}

func (x *AppendEntriesArgs) GetPrevLogIndex() int32 {
	if x != nil {
		return x.PrevLogIndex
	}
	return 0
}

func (x *AppendEntriesArgs) GetPrevLogTerm() int32 {
	if x != nil {
		return x.PrevLogTerm
	}
	return 0
}

func (x *AppendEntriesArgs) GetLeaderCommitID() int32 {
	if x != nil {
		return x.LeaderCommitID
	}
	return 0
}

func (x *AppendEntriesArgs) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type AppendEntriesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term    int32 `protobuf:"varint,1,opt,name=Term,proto3" json:"Term,omitempty"`
	Success bool  `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AppendEntriesReply) Reset() {
	*x = AppendEntriesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppendEntriesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppendEntriesReply) ProtoMessage() {}

func (x *AppendEntriesReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppendEntriesReply.ProtoReflect.Descriptor instead.
func (*AppendEntriesReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *AppendEntriesReply) GetTerm() int32 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *AppendEntriesReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type NewEntryArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry []byte `protobuf:"bytes,1,opt,name=Entry,proto3" json:"Entry,omitempty"`
}

func (x *NewEntryArgs) Reset() {
	*x = NewEntryArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEntryArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEntryArgs) ProtoMessage() {}

func (x *NewEntryArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEntryArgs.ProtoReflect.Descriptor instead.
func (*NewEntryArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *NewEntryArgs) GetEntry() []byte {
	if x != nil {
		return x.Entry
	}
	return nil
}

type NewEntryReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	LeaderID int32  `protobuf:"varint,3,opt,name=leaderID,proto3" json:"leaderID,omitempty"`
}

func (x *NewEntryReply) Reset() {
	*x = NewEntryReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEntryReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEntryReply) ProtoMessage() {}

func (x *NewEntryReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEntryReply.ProtoReflect.Descriptor instead.
func (*NewEntryReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *NewEntryReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *NewEntryReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *NewEntryReply) GetLeaderID() int32 {
	if x != nil {
		return x.LeaderID
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54,
	0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x4c, 0x61, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x61, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x4c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x22, 0x48, 0x0a, 0x10, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54,
	0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x56, 0x6f, 0x74, 0x65, 0x47, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0x5d, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xcb, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x41, 0x72, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65,
	0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72,
	0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0c, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20,
	0x0a, 0x0b, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x50, 0x72, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x54, 0x65, 0x72, 0x6d,
	0x12, 0x26, 0x0a, 0x0e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x4c, 0x6f,
	0x67, 0x73, 0x22, 0x42, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x54, 0x65, 0x72, 0x6d, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x41, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x57, 0x0a, 0x0d,
	0x4e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x32, 0x74, 0x0a, 0x04, 0x72, 0x61, 0x66, 0x74, 0x12, 0x32, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x11,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x38, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x12, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x13, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x32, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x0d, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x72, 0x67, 0x73, 0x1a,
	0x0e, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x22, 0x5a, 0x20, 0x6c, 0x78, 0x68, 0x30, 0x32, 0x37, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(*RequestVoteArgs)(nil),    // 0: RequestVoteArgs
	(*RequestVoteReply)(nil),   // 1: RequestVoteReply
	(*Log)(nil),                // 2: Log
	(*AppendEntriesArgs)(nil),  // 3: AppendEntriesArgs
	(*AppendEntriesReply)(nil), // 4: AppendEntriesReply
	(*NewEntryArgs)(nil),       // 5: NewEntryArgs
	(*NewEntryReply)(nil),      // 6: NewEntryReply
}
var file_api_proto_depIdxs = []int32{
	2, // 0: AppendEntriesArgs.Logs:type_name -> Log
	0, // 1: raft.RequestVote:input_type -> RequestVoteArgs
	3, // 2: raft.AppendEntries:input_type -> AppendEntriesArgs
	5, // 3: state.NewEntry:input_type -> NewEntryArgs
	1, // 4: raft.RequestVote:output_type -> RequestVoteReply
	4, // 5: raft.AppendEntries:output_type -> AppendEntriesReply
	6, // 6: state.NewEntry:output_type -> NewEntryReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteArgs); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestVoteReply); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesArgs); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppendEntriesReply); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEntryArgs); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEntryReply); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

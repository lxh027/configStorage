// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/register/register.proto

package register

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

type RegisterRaftArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	RaftID     string `protobuf:"bytes,2,opt,name=RaftID,proto3" json:"RaftID,omitempty"`
	Host       string `protobuf:"bytes,3,opt,name=Host,proto3" json:"Host,omitempty"`
	RaftPort   string `protobuf:"bytes,4,opt,name=RaftPort,proto3" json:"RaftPort,omitempty"`
	ClientPort string `protobuf:"bytes,5,opt,name=ClientPort,proto3" json:"ClientPort,omitempty"`
}

func (x *RegisterRaftArgs) Reset() {
	*x = RegisterRaftArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_register_register_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRaftArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRaftArgs) ProtoMessage() {}

func (x *RegisterRaftArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_register_register_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRaftArgs.ProtoReflect.Descriptor instead.
func (*RegisterRaftArgs) Descriptor() ([]byte, []int) {
	return file_api_register_register_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRaftArgs) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *RegisterRaftArgs) GetRaftID() string {
	if x != nil {
		return x.RaftID
	}
	return ""
}

func (x *RegisterRaftArgs) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RegisterRaftArgs) GetRaftPort() string {
	if x != nil {
		return x.RaftPort
	}
	return ""
}

func (x *RegisterRaftArgs) GetClientPort() string {
	if x != nil {
		return x.ClientPort
	}
	return ""
}

type RegisterRaftReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK bool `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
}

func (x *RegisterRaftReply) Reset() {
	*x = RegisterRaftReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_register_register_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRaftReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRaftReply) ProtoMessage() {}

func (x *RegisterRaftReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_register_register_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRaftReply.ProtoReflect.Descriptor instead.
func (*RegisterRaftReply) Descriptor() ([]byte, []int) {
	return file_api_register_register_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRaftReply) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

type GetRaftRegistrationsArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	RaftID string `protobuf:"bytes,2,opt,name=RaftID,proto3" json:"RaftID,omitempty"`
}

func (x *GetRaftRegistrationsArgs) Reset() {
	*x = GetRaftRegistrationsArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_register_register_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaftRegistrationsArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaftRegistrationsArgs) ProtoMessage() {}

func (x *GetRaftRegistrationsArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_register_register_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaftRegistrationsArgs.ProtoReflect.Descriptor instead.
func (*GetRaftRegistrationsArgs) Descriptor() ([]byte, []int) {
	return file_api_register_register_proto_rawDescGZIP(), []int{2}
}

func (x *GetRaftRegistrationsArgs) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *GetRaftRegistrationsArgs) GetRaftID() string {
	if x != nil {
		return x.RaftID
	}
	return ""
}

type GetRaftRegistrationsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OK     bool   `protobuf:"varint,1,opt,name=OK,proto3" json:"OK,omitempty"`
	Config []byte `protobuf:"bytes,2,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (x *GetRaftRegistrationsReply) Reset() {
	*x = GetRaftRegistrationsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_register_register_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaftRegistrationsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaftRegistrationsReply) ProtoMessage() {}

func (x *GetRaftRegistrationsReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_register_register_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaftRegistrationsReply.ProtoReflect.Descriptor instead.
func (*GetRaftRegistrationsReply) Descriptor() ([]byte, []int) {
	return file_api_register_register_proto_rawDescGZIP(), []int{3}
}

func (x *GetRaftRegistrationsReply) GetOK() bool {
	if x != nil {
		return x.OK
	}
	return false
}

func (x *GetRaftRegistrationsReply) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type UnregisterRaftArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	RaftID string `protobuf:"bytes,2,opt,name=RaftID,proto3" json:"RaftID,omitempty"`
	Idx    int64  `protobuf:"varint,3,opt,name=Idx,proto3" json:"Idx,omitempty"`
}

func (x *UnregisterRaftArgs) Reset() {
	*x = UnregisterRaftArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_register_register_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterRaftArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterRaftArgs) ProtoMessage() {}

func (x *UnregisterRaftArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_register_register_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterRaftArgs.ProtoReflect.Descriptor instead.
func (*UnregisterRaftArgs) Descriptor() ([]byte, []int) {
	return file_api_register_register_proto_rawDescGZIP(), []int{4}
}

func (x *UnregisterRaftArgs) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UnregisterRaftArgs) GetRaftID() string {
	if x != nil {
		return x.RaftID
	}
	return ""
}

func (x *UnregisterRaftArgs) GetIdx() int64 {
	if x != nil {
		return x.Idx
	}
	return 0
}

type UnregisterRaftReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnregisterRaftReply) Reset() {
	*x = UnregisterRaftReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_register_register_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterRaftReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterRaftReply) ProtoMessage() {}

func (x *UnregisterRaftReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_register_register_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterRaftReply.ProtoReflect.Descriptor instead.
func (*UnregisterRaftReply) Descriptor() ([]byte, []int) {
	return file_api_register_register_proto_rawDescGZIP(), []int{5}
}

var File_api_register_register_proto protoreflect.FileDescriptor

var file_api_register_register_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01,
	0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x66, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x61, 0x66, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x61, 0x66, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x52, 0x61, 0x66, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x23, 0x0a, 0x11,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4f,
	0x4b, 0x22, 0x44, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x61, 0x66, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x72, 0x67, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x52, 0x61, 0x66, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x52, 0x61, 0x66, 0x74, 0x49, 0x44, 0x22, 0x43, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x61,
	0x66, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x4b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x4f, 0x4b, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x50, 0x0a, 0x12,
	0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x41, 0x72,
	0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x55, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x66, 0x74, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x61, 0x66, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03,
	0x49, 0x64, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x49, 0x64, 0x78, 0x22, 0x15,
	0x0a, 0x13, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xd1, 0x01, 0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x12, 0x35, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x12, 0x11, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x61, 0x66, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x12, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4d, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x52, 0x61, 0x66, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x66, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x72, 0x67, 0x73,
	0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x66, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0e,
	0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x12, 0x13,
	0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x61, 0x66, 0x74, 0x41,
	0x72, 0x67, 0x73, 0x1a, 0x14, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x61, 0x66, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x23, 0x5a, 0x21, 0x6c, 0x78, 0x68,
	0x30, 0x32, 0x37, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_register_register_proto_rawDescOnce sync.Once
	file_api_register_register_proto_rawDescData = file_api_register_register_proto_rawDesc
)

func file_api_register_register_proto_rawDescGZIP() []byte {
	file_api_register_register_proto_rawDescOnce.Do(func() {
		file_api_register_register_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_register_register_proto_rawDescData)
	})
	return file_api_register_register_proto_rawDescData
}

var file_api_register_register_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_register_register_proto_goTypes = []interface{}{
	(*RegisterRaftArgs)(nil),          // 0: RegisterRaftArgs
	(*RegisterRaftReply)(nil),         // 1: RegisterRaftReply
	(*GetRaftRegistrationsArgs)(nil),  // 2: GetRaftRegistrationsArgs
	(*GetRaftRegistrationsReply)(nil), // 3: GetRaftRegistrationsReply
	(*UnregisterRaftArgs)(nil),        // 4: UnregisterRaftArgs
	(*UnregisterRaftReply)(nil),       // 5: UnregisterRaftReply
}
var file_api_register_register_proto_depIdxs = []int32{
	0, // 0: registerRaft.RegisterRaft:input_type -> RegisterRaftArgs
	2, // 1: registerRaft.GetRaftRegistrations:input_type -> GetRaftRegistrationsArgs
	4, // 2: registerRaft.UnregisterRaft:input_type -> UnregisterRaftArgs
	1, // 3: registerRaft.RegisterRaft:output_type -> RegisterRaftReply
	3, // 4: registerRaft.GetRaftRegistrations:output_type -> GetRaftRegistrationsReply
	5, // 5: registerRaft.UnregisterRaft:output_type -> UnregisterRaftReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_register_register_proto_init() }
func file_api_register_register_proto_init() {
	if File_api_register_register_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_register_register_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRaftArgs); i {
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
		file_api_register_register_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRaftReply); i {
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
		file_api_register_register_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaftRegistrationsArgs); i {
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
		file_api_register_register_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaftRegistrationsReply); i {
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
		file_api_register_register_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterRaftArgs); i {
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
		file_api_register_register_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterRaftReply); i {
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
			RawDescriptor: file_api_register_register_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_register_register_proto_goTypes,
		DependencyIndexes: file_api_register_register_proto_depIdxs,
		MessageInfos:      file_api_register_register_proto_msgTypes,
	}.Build()
	File_api_register_register_proto = out.File
	file_api_register_register_proto_rawDesc = nil
	file_api_register_register_proto_goTypes = nil
	file_api_register_register_proto_depIdxs = nil
}
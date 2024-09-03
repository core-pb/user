// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: user/v1/user.proto

package user

import (
	v1 "github.com/core-pb/dt/time/v1"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64           `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty" bun:",pk,autoincrement"`
	Username  string           `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Disable   bool             `protobuf:"varint,3,opt,name=disable,proto3" json:"disable,omitempty"`
	Data      *structpb.Struct `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty" bun:"type:jsonb"`
	Info      *structpb.Struct `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty" bun:"type:jsonb"`
	CreatedAt *v1.Time         `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bun:"type:timestamptz"`
	UpdatedAt *v1.Time         `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" bun:"type:timestamptz"`
	DeletedAt *v1.Time         `protobuf:"bytes,16,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty" bun:"type:timestamptz,soft_delete,nullzero"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[0]
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
	return file_user_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

func (x *User) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *User) GetInfo() *structpb.Struct {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *User) GetCreatedAt() *v1.Time {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *v1.Time {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *User) GetDeletedAt() *v1.Time {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	UserAuth []*UserAuth `protobuf:"bytes,2,rep,name=user_auth,json=userAuth,proto3" json:"user_auth,omitempty"`
	UserTag  []*UserTag  `protobuf:"bytes,3,rep,name=user_tag,json=userTag,proto3" json:"user_tag,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_user_v1_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_user_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserDetail) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserDetail) GetUserAuth() []*UserAuth {
	if x != nil {
		return x.UserAuth
	}
	return nil
}

func (x *UserDetail) GetUserTag() []*UserTag {
	if x != nil {
		return x.UserTag
	}
	return nil
}

var File_user_v1_user_proto protoreflect.FileDescriptor

var file_user_v1_user_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x12, 0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x1c, 0x9a, 0x84, 0x9e, 0x03,
	0x17, 0x62, 0x75, 0x6e, 0x3a, 0x22, 0x2c, 0x70, 0x6b, 0x2c, 0x61, 0x75, 0x74, 0x6f, 0x69, 0x6e,
	0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10,
	0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x6a, 0x73, 0x6f, 0x6e, 0x62, 0x22,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x15, 0x9a,
	0x84, 0x9e, 0x03, 0x10, 0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x6a, 0x73,
	0x6f, 0x6e, 0x62, 0x22, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x1b, 0x9a,
	0x84, 0x9e, 0x03, 0x16, 0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x74, 0x7a, 0x22, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x49, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x1b, 0x9a, 0x84, 0x9e, 0x03, 0x16, 0x62,
	0x75, 0x6e, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x74, 0x7a, 0x22, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x63, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x30, 0x9a, 0x84, 0x9e, 0x03, 0x2b, 0x62, 0x75, 0x6e, 0x3a, 0x22, 0x74,
	0x79, 0x70, 0x65, 0x3a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x74, 0x7a, 0x2c,
	0x73, 0x6f, 0x66, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2c, 0x6e, 0x75, 0x6c, 0x6c,
	0x7a, 0x65, 0x72, 0x6f, 0x22, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x61, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x67, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x61, 0x67, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_v1_user_proto_rawDescOnce sync.Once
	file_user_v1_user_proto_rawDescData = file_user_v1_user_proto_rawDesc
)

func file_user_v1_user_proto_rawDescGZIP() []byte {
	file_user_v1_user_proto_rawDescOnce.Do(func() {
		file_user_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_v1_user_proto_rawDescData)
	})
	return file_user_v1_user_proto_rawDescData
}

var file_user_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_v1_user_proto_goTypes = []any{
	(*User)(nil),            // 0: user.v1.User
	(*UserDetail)(nil),      // 1: user.v1.UserDetail
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
	(*v1.Time)(nil),         // 3: time.v1.Time
	(*UserAuth)(nil),        // 4: user.v1.UserAuth
	(*UserTag)(nil),         // 5: user.v1.UserTag
}
var file_user_v1_user_proto_depIdxs = []int32{
	2, // 0: user.v1.User.data:type_name -> google.protobuf.Struct
	2, // 1: user.v1.User.info:type_name -> google.protobuf.Struct
	3, // 2: user.v1.User.created_at:type_name -> time.v1.Time
	3, // 3: user.v1.User.updated_at:type_name -> time.v1.Time
	3, // 4: user.v1.User.deleted_at:type_name -> time.v1.Time
	0, // 5: user.v1.UserDetail.user:type_name -> user.v1.User
	4, // 6: user.v1.UserDetail.user_auth:type_name -> user.v1.UserAuth
	5, // 7: user.v1.UserDetail.user_tag:type_name -> user.v1.UserTag
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_user_v1_user_proto_init() }
func file_user_v1_user_proto_init() {
	if File_user_v1_user_proto != nil {
		return
	}
	file_user_v1_user_auth_proto_init()
	file_user_v1_user_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_v1_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_user_v1_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserDetail); i {
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
	file_user_v1_user_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_v1_user_proto_goTypes,
		DependencyIndexes: file_user_v1_user_proto_depIdxs,
		MessageInfos:      file_user_v1_user_proto_msgTypes,
	}.Build()
	File_user_v1_user_proto = out.File
	file_user_v1_user_proto_rawDesc = nil
	file_user_v1_user_proto_goTypes = nil
	file_user_v1_user_proto_depIdxs = nil
}

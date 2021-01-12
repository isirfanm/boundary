// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: controller/storage/iam/store/v1/role_grant.proto

package store

import (
	timestamp "github.com/hashicorp/boundary/internal/db/timestamp"
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

type RoleGrant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// create_time from the RDBMS
	// @inject_tag: `gorm:"default:current_timestamp"`
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" gorm:"default:current_timestamp"`
	// role_id is the ID of the role this is a part of
	// @inject_tag: gorm:"primary_key"
	RoleId string `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty" gorm:"primary_key"`
	// raw_grant is the string grant value as provided by the user
	// @inject_tag: `gorm:"default:null"`
	RawGrant string `protobuf:"bytes,3,opt,name=raw_grant,json=rawGrant,proto3" json:"raw_grant,omitempty" gorm:"default:null"`
	// canonical_grant is the canonical string representation of the grant value.
	// We use this as the unique constraint.
	// @inject_tag: gorm:"primary_key"
	CanonicalGrant string `protobuf:"bytes,4,opt,name=canonical_grant,json=canonicalGrant,proto3" json:"canonical_grant,omitempty" gorm:"primary_key"`
}

func (x *RoleGrant) Reset() {
	*x = RoleGrant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_storage_iam_store_v1_role_grant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleGrant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleGrant) ProtoMessage() {}

func (x *RoleGrant) ProtoReflect() protoreflect.Message {
	mi := &file_controller_storage_iam_store_v1_role_grant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleGrant.ProtoReflect.Descriptor instead.
func (*RoleGrant) Descriptor() ([]byte, []int) {
	return file_controller_storage_iam_store_v1_role_grant_proto_rawDescGZIP(), []int{0}
}

func (x *RoleGrant) GetCreateTime() *timestamp.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *RoleGrant) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *RoleGrant) GetRawGrant() string {
	if x != nil {
		return x.RawGrant
	}
	return ""
}

func (x *RoleGrant) GetCanonicalGrant() string {
	if x != nil {
		return x.CanonicalGrant
	}
	return ""
}

var File_controller_storage_iam_store_v1_role_grant_proto protoreflect.FileDescriptor

var file_controller_storage_iam_store_v1_role_grant_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12,
	0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x67, 0x72, 0x61,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x61, 0x77, 0x47, 0x72, 0x61,
	0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x67, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x6e,
	0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63,
	0x6f, 0x72, 0x70, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_storage_iam_store_v1_role_grant_proto_rawDescOnce sync.Once
	file_controller_storage_iam_store_v1_role_grant_proto_rawDescData = file_controller_storage_iam_store_v1_role_grant_proto_rawDesc
)

func file_controller_storage_iam_store_v1_role_grant_proto_rawDescGZIP() []byte {
	file_controller_storage_iam_store_v1_role_grant_proto_rawDescOnce.Do(func() {
		file_controller_storage_iam_store_v1_role_grant_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_storage_iam_store_v1_role_grant_proto_rawDescData)
	})
	return file_controller_storage_iam_store_v1_role_grant_proto_rawDescData
}

var file_controller_storage_iam_store_v1_role_grant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_controller_storage_iam_store_v1_role_grant_proto_goTypes = []interface{}{
	(*RoleGrant)(nil),           // 0: controller.storage.iam.store.v1.RoleGrant
	(*timestamp.Timestamp)(nil), // 1: controller.storage.timestamp.v1.Timestamp
}
var file_controller_storage_iam_store_v1_role_grant_proto_depIdxs = []int32{
	1, // 0: controller.storage.iam.store.v1.RoleGrant.create_time:type_name -> controller.storage.timestamp.v1.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_controller_storage_iam_store_v1_role_grant_proto_init() }
func file_controller_storage_iam_store_v1_role_grant_proto_init() {
	if File_controller_storage_iam_store_v1_role_grant_proto != nil {
		return
	}
	file_controller_storage_iam_store_v1_scope_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_controller_storage_iam_store_v1_role_grant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleGrant); i {
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
			RawDescriptor: file_controller_storage_iam_store_v1_role_grant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_storage_iam_store_v1_role_grant_proto_goTypes,
		DependencyIndexes: file_controller_storage_iam_store_v1_role_grant_proto_depIdxs,
		MessageInfos:      file_controller_storage_iam_store_v1_role_grant_proto_msgTypes,
	}.Build()
	File_controller_storage_iam_store_v1_role_grant_proto = out.File
	file_controller_storage_iam_store_v1_role_grant_proto_rawDesc = nil
	file_controller_storage_iam_store_v1_role_grant_proto_goTypes = nil
	file_controller_storage_iam_store_v1_role_grant_proto_depIdxs = nil
}

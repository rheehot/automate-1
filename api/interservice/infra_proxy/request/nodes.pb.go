// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: interservice/infra_proxy/request/nodes.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type AffectedNodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Chef object type (e.g. 'cookbooks', 'roles', 'chef_environment').
	ChefType string `protobuf:"bytes,3,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty" toml:"chef_type,omitempty" mapstructure:"chef_type,omitempty"`
	// Chef object name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Chef object version.
	Version string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
}

func (x *AffectedNodes) Reset() {
	*x = AffectedNodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AffectedNodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AffectedNodes) ProtoMessage() {}

func (x *AffectedNodes) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AffectedNodes.ProtoReflect.Descriptor instead.
func (*AffectedNodes) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *AffectedNodes) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *AffectedNodes) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *AffectedNodes) GetChefType() string {
	if x != nil {
		return x.ChefType
	}
	return ""
}

func (x *AffectedNodes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AffectedNodes) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type DeleteNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
}

func (x *DeleteNode) Reset() {
	*x = DeleteNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_nodes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNode) ProtoMessage() {}

func (x *DeleteNode) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_nodes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNode.ProtoReflect.Descriptor instead.
func (*DeleteNode) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteNode) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *DeleteNode) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *DeleteNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty" toml:"org_id,omitempty" mapstructure:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty" toml:"server_id,omitempty" mapstructure:"server_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Node environment.
	Environment string `protobuf:"bytes,4,opt,name=environment,proto3" json:"environment,omitempty" toml:"environment,omitempty" mapstructure:"environment,omitempty"`
	// Node policy name.
	PolicyName string `protobuf:"bytes,5,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty" toml:"policy_name,omitempty" mapstructure:"policy_name,omitempty"`
	// Node policy group.
	PolicyGroup string `protobuf:"bytes,6,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty" toml:"policy_group,omitempty" mapstructure:"policy_group,omitempty"`
	// Node run-list.
	RunList []string `protobuf:"bytes,7,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty" toml:"run_list,omitempty" mapstructure:"run_list,omitempty"`
	// Node automatic attributes JSON.
	AutomaticAttributes *_struct.Struct `protobuf:"bytes,8,opt,name=automatic_attributes,json=automaticAttributes,proto3" json:"automatic_attributes,omitempty" toml:"automatic_attributes,omitempty" mapstructure:"automatic_attributes,omitempty"`
	// Node default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,9,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty" toml:"default_attributes,omitempty" mapstructure:"default_attributes,omitempty"`
	// Node normal attributes JSON.
	NormalAttributes *_struct.Struct `protobuf:"bytes,10,opt,name=normal_attributes,json=normalAttributes,proto3" json:"normal_attributes,omitempty" toml:"normal_attributes,omitempty" mapstructure:"normal_attributes,omitempty"`
	// Node override attributes JSON.
	OverrideAttributes *_struct.Struct `protobuf:"bytes,11,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty" toml:"override_attributes,omitempty" mapstructure:"override_attributes,omitempty"`
}

func (x *UpdateNode) Reset() {
	*x = UpdateNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_infra_proxy_request_nodes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNode) ProtoMessage() {}

func (x *UpdateNode) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_infra_proxy_request_nodes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNode.ProtoReflect.Descriptor instead.
func (*UpdateNode) Descriptor() ([]byte, []int) {
	return file_interservice_infra_proxy_request_nodes_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNode) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *UpdateNode) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *UpdateNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNode) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *UpdateNode) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *UpdateNode) GetPolicyGroup() string {
	if x != nil {
		return x.PolicyGroup
	}
	return ""
}

func (x *UpdateNode) GetRunList() []string {
	if x != nil {
		return x.RunList
	}
	return nil
}

func (x *UpdateNode) GetAutomaticAttributes() *_struct.Struct {
	if x != nil {
		return x.AutomaticAttributes
	}
	return nil
}

func (x *UpdateNode) GetDefaultAttributes() *_struct.Struct {
	if x != nil {
		return x.DefaultAttributes
	}
	return nil
}

func (x *UpdateNode) GetNormalAttributes() *_struct.Struct {
	if x != nil {
		return x.NormalAttributes
	}
	return nil
}

func (x *UpdateNode) GetOverrideAttributes() *_struct.Struct {
	if x != nil {
		return x.OverrideAttributes
	}
	return nil
}

var File_interservice_infra_proxy_request_nodes_proto protoreflect.FileDescriptor

var file_interservice_infra_proxy_request_nodes_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0d, 0x41, 0x66, 0x66, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x68, 0x65, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x68, 0x65, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf9, 0x03,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72,
	0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x75,
	0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x69, 0x63, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x13, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x12, 0x46, 0x0a, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x11, 0x6e, 0x6f, 0x72,
	0x6d, 0x61, 0x6c, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x10, 0x6e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x48, 0x0a, 0x13, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x12, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_interservice_infra_proxy_request_nodes_proto_rawDescOnce sync.Once
	file_interservice_infra_proxy_request_nodes_proto_rawDescData = file_interservice_infra_proxy_request_nodes_proto_rawDesc
)

func file_interservice_infra_proxy_request_nodes_proto_rawDescGZIP() []byte {
	file_interservice_infra_proxy_request_nodes_proto_rawDescOnce.Do(func() {
		file_interservice_infra_proxy_request_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_infra_proxy_request_nodes_proto_rawDescData)
	})
	return file_interservice_infra_proxy_request_nodes_proto_rawDescData
}

var file_interservice_infra_proxy_request_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interservice_infra_proxy_request_nodes_proto_goTypes = []interface{}{
	(*AffectedNodes)(nil),  // 0: chef.automate.domain.infra_proxy.request.AffectedNodes
	(*DeleteNode)(nil),     // 1: chef.automate.domain.infra_proxy.request.DeleteNode
	(*UpdateNode)(nil),     // 2: chef.automate.domain.infra_proxy.request.UpdateNode
	(*_struct.Struct)(nil), // 3: google.protobuf.Struct
}
var file_interservice_infra_proxy_request_nodes_proto_depIdxs = []int32{
	3, // 0: chef.automate.domain.infra_proxy.request.UpdateNode.automatic_attributes:type_name -> google.protobuf.Struct
	3, // 1: chef.automate.domain.infra_proxy.request.UpdateNode.default_attributes:type_name -> google.protobuf.Struct
	3, // 2: chef.automate.domain.infra_proxy.request.UpdateNode.normal_attributes:type_name -> google.protobuf.Struct
	3, // 3: chef.automate.domain.infra_proxy.request.UpdateNode.override_attributes:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_interservice_infra_proxy_request_nodes_proto_init() }
func file_interservice_infra_proxy_request_nodes_proto_init() {
	if File_interservice_infra_proxy_request_nodes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_infra_proxy_request_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AffectedNodes); i {
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
		file_interservice_infra_proxy_request_nodes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNode); i {
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
		file_interservice_infra_proxy_request_nodes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNode); i {
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
			RawDescriptor: file_interservice_infra_proxy_request_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_infra_proxy_request_nodes_proto_goTypes,
		DependencyIndexes: file_interservice_infra_proxy_request_nodes_proto_depIdxs,
		MessageInfos:      file_interservice_infra_proxy_request_nodes_proto_msgTypes,
	}.Build()
	File_interservice_infra_proxy_request_nodes_proto = out.File
	file_interservice_infra_proxy_request_nodes_proto_rawDesc = nil
	file_interservice_infra_proxy_request_nodes_proto_goTypes = nil
	file_interservice_infra_proxy_request_nodes_proto_depIdxs = nil
}

// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: google/iam/v1/iam_policy.proto

package iampb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for `SetIamPolicy` method.
type SetIamPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED: The resource for which the policy is being specified.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// REQUIRED: The complete policy to be applied to the `resource`. The size of
	// the policy is limited to a few 10s of KB. An empty policy is a
	// valid policy but certain Cloud Platform services (such as Projects)
	// might reject them.
	Policy *Policy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only
	// the fields in the mask will be modified. If no mask is provided, the
	// following default mask is used:
	//
	// `paths: "bindings, etag"`
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *SetIamPolicyRequest) Reset() {
	*x = SetIamPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_iam_v1_iam_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetIamPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIamPolicyRequest) ProtoMessage() {}

func (x *SetIamPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_iam_v1_iam_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIamPolicyRequest.ProtoReflect.Descriptor instead.
func (*SetIamPolicyRequest) Descriptor() ([]byte, []int) {
	return file_google_iam_v1_iam_policy_proto_rawDescGZIP(), []int{0}
}

func (x *SetIamPolicyRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *SetIamPolicyRequest) GetPolicy() *Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *SetIamPolicyRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request message for `GetIamPolicy` method.
type GetIamPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED: The resource for which the policy is being requested.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// OPTIONAL: A `GetPolicyOptions` object for specifying options to
	// `GetIamPolicy`.
	Options *GetPolicyOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *GetIamPolicyRequest) Reset() {
	*x = GetIamPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_iam_v1_iam_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIamPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIamPolicyRequest) ProtoMessage() {}

func (x *GetIamPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_iam_v1_iam_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIamPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetIamPolicyRequest) Descriptor() ([]byte, []int) {
	return file_google_iam_v1_iam_policy_proto_rawDescGZIP(), []int{1}
}

func (x *GetIamPolicyRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *GetIamPolicyRequest) GetOptions() *GetPolicyOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

// Request message for `TestIamPermissions` method.
type TestIamPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED: The resource for which the policy detail is being requested.
	// See the operation documentation for the appropriate value for this field.
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The set of permissions to check for the `resource`. Permissions with
	// wildcards (such as '*' or 'storage.*') are not allowed. For more
	// information see
	// [IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).
	Permissions []string `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *TestIamPermissionsRequest) Reset() {
	*x = TestIamPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_iam_v1_iam_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestIamPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestIamPermissionsRequest) ProtoMessage() {}

func (x *TestIamPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_iam_v1_iam_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestIamPermissionsRequest.ProtoReflect.Descriptor instead.
func (*TestIamPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_google_iam_v1_iam_policy_proto_rawDescGZIP(), []int{2}
}

func (x *TestIamPermissionsRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *TestIamPermissionsRequest) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

// Response message for `TestIamPermissions` method.
type TestIamPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A subset of `TestPermissionsRequest.permissions` that the caller is
	// allowed.
	Permissions []string `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *TestIamPermissionsResponse) Reset() {
	*x = TestIamPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_iam_v1_iam_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestIamPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestIamPermissionsResponse) ProtoMessage() {}

func (x *TestIamPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_iam_v1_iam_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestIamPermissionsResponse.ProtoReflect.Descriptor instead.
func (*TestIamPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_google_iam_v1_iam_policy_proto_rawDescGZIP(), []int{3}
}

func (x *TestIamPermissionsResponse) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_google_iam_v1_iam_policy_proto protoreflect.FileDescriptor

var file_google_iam_v1_iam_policy_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x61, 0x6d, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01,
	0x0a, 0x13, 0x53, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x03, 0x0a,
	0x01, 0x2a, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x77, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x03, 0x0a, 0x01,
	0x2a, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x69, 0x0a, 0x19, 0x54, 0x65, 0x73, 0x74, 0x49, 0x61,
	0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x03, 0x0a, 0x01, 0x2a,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x3e, 0x0a, 0x1a, 0x54, 0x65, 0x73, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x32, 0xb4, 0x03, 0x0a, 0x09, 0x49, 0x41, 0x4d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x74, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12,
	0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x3d, 0x2a, 0x2a, 0x7d, 0x3a, 0x73, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x74, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3d, 0x2a, 0x2a, 0x7d, 0x3a, 0x67, 0x65, 0x74, 0x49,
	0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x9a, 0x01, 0x0a, 0x12,
	0x54, 0x65, 0x73, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22,
	0x24, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3d, 0x2a,
	0x2a, 0x7d, 0x3a, 0x74, 0x65, 0x73, 0x74, 0x49, 0x61, 0x6d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x1a, 0x1e, 0xca, 0x41, 0x1b, 0x69, 0x61, 0x6d,
	0x2d, 0x6d, 0x65, 0x74, 0x61, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0x7f, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x49,
	0x61, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x29, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x69,
	0x61, 0x6d, 0x70, 0x62, 0x3b, 0x69, 0x61, 0x6d, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x13,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x49, 0x61, 0x6d,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x49, 0x61, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_iam_v1_iam_policy_proto_rawDescOnce sync.Once
	file_google_iam_v1_iam_policy_proto_rawDescData = file_google_iam_v1_iam_policy_proto_rawDesc
)

func file_google_iam_v1_iam_policy_proto_rawDescGZIP() []byte {
	file_google_iam_v1_iam_policy_proto_rawDescOnce.Do(func() {
		file_google_iam_v1_iam_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_iam_v1_iam_policy_proto_rawDescData)
	})
	return file_google_iam_v1_iam_policy_proto_rawDescData
}

var file_google_iam_v1_iam_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_iam_v1_iam_policy_proto_goTypes = []interface{}{
	(*SetIamPolicyRequest)(nil),        // 0: google.iam.v1.SetIamPolicyRequest
	(*GetIamPolicyRequest)(nil),        // 1: google.iam.v1.GetIamPolicyRequest
	(*TestIamPermissionsRequest)(nil),  // 2: google.iam.v1.TestIamPermissionsRequest
	(*TestIamPermissionsResponse)(nil), // 3: google.iam.v1.TestIamPermissionsResponse
	(*Policy)(nil),                     // 4: google.iam.v1.Policy
	(*fieldmaskpb.FieldMask)(nil),      // 5: google.protobuf.FieldMask
	(*GetPolicyOptions)(nil),           // 6: google.iam.v1.GetPolicyOptions
}
var file_google_iam_v1_iam_policy_proto_depIdxs = []int32{
	4, // 0: google.iam.v1.SetIamPolicyRequest.policy:type_name -> google.iam.v1.Policy
	5, // 1: google.iam.v1.SetIamPolicyRequest.update_mask:type_name -> google.protobuf.FieldMask
	6, // 2: google.iam.v1.GetIamPolicyRequest.options:type_name -> google.iam.v1.GetPolicyOptions
	0, // 3: google.iam.v1.IAMPolicy.SetIamPolicy:input_type -> google.iam.v1.SetIamPolicyRequest
	1, // 4: google.iam.v1.IAMPolicy.GetIamPolicy:input_type -> google.iam.v1.GetIamPolicyRequest
	2, // 5: google.iam.v1.IAMPolicy.TestIamPermissions:input_type -> google.iam.v1.TestIamPermissionsRequest
	4, // 6: google.iam.v1.IAMPolicy.SetIamPolicy:output_type -> google.iam.v1.Policy
	4, // 7: google.iam.v1.IAMPolicy.GetIamPolicy:output_type -> google.iam.v1.Policy
	3, // 8: google.iam.v1.IAMPolicy.TestIamPermissions:output_type -> google.iam.v1.TestIamPermissionsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_iam_v1_iam_policy_proto_init() }
func file_google_iam_v1_iam_policy_proto_init() {
	if File_google_iam_v1_iam_policy_proto != nil {
		return
	}
	file_google_iam_v1_options_proto_init()
	file_google_iam_v1_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_iam_v1_iam_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetIamPolicyRequest); i {
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
		file_google_iam_v1_iam_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIamPolicyRequest); i {
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
		file_google_iam_v1_iam_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestIamPermissionsRequest); i {
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
		file_google_iam_v1_iam_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestIamPermissionsResponse); i {
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
			RawDescriptor: file_google_iam_v1_iam_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_iam_v1_iam_policy_proto_goTypes,
		DependencyIndexes: file_google_iam_v1_iam_policy_proto_depIdxs,
		MessageInfos:      file_google_iam_v1_iam_policy_proto_msgTypes,
	}.Build()
	File_google_iam_v1_iam_policy_proto = out.File
	file_google_iam_v1_iam_policy_proto_rawDesc = nil
	file_google_iam_v1_iam_policy_proto_goTypes = nil
	file_google_iam_v1_iam_policy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IAMPolicyClient is the client API for IAMPolicy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IAMPolicyClient interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	//
	// Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
	SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a `NOT_FOUND` error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error)
}

type iAMPolicyClient struct {
	cc grpc.ClientConnInterface
}

func NewIAMPolicyClient(cc grpc.ClientConnInterface) IAMPolicyClient {
	return &iAMPolicyClient{cc}
}

func (c *iAMPolicyClient) SetIamPolicy(ctx context.Context, in *SetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/SetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) GetIamPolicy(ctx context.Context, in *GetIamPolicyRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/GetIamPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iAMPolicyClient) TestIamPermissions(ctx context.Context, in *TestIamPermissionsRequest, opts ...grpc.CallOption) (*TestIamPermissionsResponse, error) {
	out := new(TestIamPermissionsResponse)
	err := c.cc.Invoke(ctx, "/google.iam.v1.IAMPolicy/TestIamPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IAMPolicyServer is the server API for IAMPolicy service.
type IAMPolicyServer interface {
	// Sets the access control policy on the specified resource. Replaces any
	// existing policy.
	//
	// Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
	SetIamPolicy(context.Context, *SetIamPolicyRequest) (*Policy, error)
	// Gets the access control policy for a resource.
	// Returns an empty policy if the resource exists and does not have a policy
	// set.
	GetIamPolicy(context.Context, *GetIamPolicyRequest) (*Policy, error)
	// Returns permissions that a caller has on the specified resource.
	// If the resource does not exist, this will return an empty set of
	// permissions, not a `NOT_FOUND` error.
	//
	// Note: This operation is designed to be used for building permission-aware
	// UIs and command-line tools, not for authorization checking. This operation
	// may "fail open" without warning.
	TestIamPermissions(context.Context, *TestIamPermissionsRequest) (*TestIamPermissionsResponse, error)
}

// UnimplementedIAMPolicyServer can be embedded to have forward compatible implementations.
type UnimplementedIAMPolicyServer struct {
}

func (*UnimplementedIAMPolicyServer) SetIamPolicy(context.Context, *SetIamPolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIamPolicy not implemented")
}
func (*UnimplementedIAMPolicyServer) GetIamPolicy(context.Context, *GetIamPolicyRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIamPolicy not implemented")
}
func (*UnimplementedIAMPolicyServer) TestIamPermissions(context.Context, *TestIamPermissionsRequest) (*TestIamPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestIamPermissions not implemented")
}

func RegisterIAMPolicyServer(s *grpc.Server, srv IAMPolicyServer) {
	s.RegisterService(&_IAMPolicy_serviceDesc, srv)
}

func _IAMPolicy_SetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).SetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/SetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).SetIamPolicy(ctx, req.(*SetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMPolicy_GetIamPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIamPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).GetIamPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/GetIamPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).GetIamPolicy(ctx, req.(*GetIamPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IAMPolicy_TestIamPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestIamPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IAMPolicyServer).TestIamPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.iam.v1.IAMPolicy/TestIamPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IAMPolicyServer).TestIamPermissions(ctx, req.(*TestIamPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IAMPolicy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.iam.v1.IAMPolicy",
	HandlerType: (*IAMPolicyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIamPolicy",
			Handler:    _IAMPolicy_SetIamPolicy_Handler,
		},
		{
			MethodName: "GetIamPolicy",
			Handler:    _IAMPolicy_GetIamPolicy_Handler,
		},
		{
			MethodName: "TestIamPermissions",
			Handler:    _IAMPolicy_TestIamPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/iam/v1/iam_policy.proto",
}

// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: pbcatalog/v2beta1/vip.proto

package catalogv2beta1

import (
	_ "github.com/hashicorp/consul/proto-public/pbresource"
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

type VirtualIPs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ips []*IP `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"`
}

func (x *VirtualIPs) Reset() {
	*x = VirtualIPs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcatalog_v2beta1_vip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualIPs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualIPs) ProtoMessage() {}

func (x *VirtualIPs) ProtoReflect() protoreflect.Message {
	mi := &file_pbcatalog_v2beta1_vip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirtualIPs.ProtoReflect.Descriptor instead.
func (*VirtualIPs) Descriptor() ([]byte, []int) {
	return file_pbcatalog_v2beta1_vip_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualIPs) GetIps() []*IP {
	if x != nil {
		return x.Ips
	}
	return nil
}

type IP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address is the string IPv4 address.
	// This could also store IPv6 in the future.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// generated indicates whether Consul generated or it is user-provided
	// (e.g. a ClusterIP of the Kubernetes service).
	Generated bool `protobuf:"varint,2,opt,name=generated,proto3" json:"generated,omitempty"`
}

func (x *IP) Reset() {
	*x = IP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcatalog_v2beta1_vip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IP) ProtoMessage() {}

func (x *IP) ProtoReflect() protoreflect.Message {
	mi := &file_pbcatalog_v2beta1_vip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IP.ProtoReflect.Descriptor instead.
func (*IP) Descriptor() ([]byte, []int) {
	return file_pbcatalog_v2beta1_vip_proto_rawDescGZIP(), []int{1}
}

func (x *IP) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *IP) GetGenerated() bool {
	if x != nil {
		return x.Generated
	}
	return false
}

var File_pbcatalog_v2beta1_vip_proto protoreflect.FileDescriptor

var file_pbcatalog_v2beta1_vip_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x76, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a,
	0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a,
	0x0a, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x50, 0x73, 0x12, 0x36, 0x0a, 0x03, 0x69,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x50, 0x52, 0x03,
	0x69, 0x70, 0x73, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x22, 0x3c, 0x0a, 0x02, 0x49,
	0x50, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x9e, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x42, 0x08, 0x56, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x49,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x43, 0xaa,
	0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x2e, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xca, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2c, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pbcatalog_v2beta1_vip_proto_rawDescOnce sync.Once
	file_pbcatalog_v2beta1_vip_proto_rawDescData = file_pbcatalog_v2beta1_vip_proto_rawDesc
)

func file_pbcatalog_v2beta1_vip_proto_rawDescGZIP() []byte {
	file_pbcatalog_v2beta1_vip_proto_rawDescOnce.Do(func() {
		file_pbcatalog_v2beta1_vip_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbcatalog_v2beta1_vip_proto_rawDescData)
	})
	return file_pbcatalog_v2beta1_vip_proto_rawDescData
}

var file_pbcatalog_v2beta1_vip_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbcatalog_v2beta1_vip_proto_goTypes = []interface{}{
	(*VirtualIPs)(nil), // 0: hashicorp.consul.catalog.v2beta1.VirtualIPs
	(*IP)(nil),         // 1: hashicorp.consul.catalog.v2beta1.IP
}
var file_pbcatalog_v2beta1_vip_proto_depIdxs = []int32{
	1, // 0: hashicorp.consul.catalog.v2beta1.VirtualIPs.ips:type_name -> hashicorp.consul.catalog.v2beta1.IP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pbcatalog_v2beta1_vip_proto_init() }
func file_pbcatalog_v2beta1_vip_proto_init() {
	if File_pbcatalog_v2beta1_vip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbcatalog_v2beta1_vip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualIPs); i {
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
		file_pbcatalog_v2beta1_vip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IP); i {
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
			RawDescriptor: file_pbcatalog_v2beta1_vip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbcatalog_v2beta1_vip_proto_goTypes,
		DependencyIndexes: file_pbcatalog_v2beta1_vip_proto_depIdxs,
		MessageInfos:      file_pbcatalog_v2beta1_vip_proto_msgTypes,
	}.Build()
	File_pbcatalog_v2beta1_vip_proto = out.File
	file_pbcatalog_v2beta1_vip_proto_rawDesc = nil
	file_pbcatalog_v2beta1_vip_proto_goTypes = nil
	file_pbcatalog_v2beta1_vip_proto_depIdxs = nil
}

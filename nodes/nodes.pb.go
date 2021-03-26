// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: nodes/nodes.proto

package nodes

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	IP                string `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port              uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	MaxConnections    uint32 `protobuf:"varint,4,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	ActiveConnections uint32 `protobuf:"varint,5,opt,name=active_connections,json=activeConnections,proto3" json:"active_connections,omitempty"`
}

func (x *NodeStatus) Reset() {
	*x = NodeStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nodes_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeStatus) ProtoMessage() {}

func (x *NodeStatus) ProtoReflect() protoreflect.Message {
	mi := &file_nodes_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeStatus.ProtoReflect.Descriptor instead.
func (*NodeStatus) Descriptor() ([]byte, []int) {
	return file_nodes_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *NodeStatus) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *NodeStatus) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *NodeStatus) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *NodeStatus) GetMaxConnections() uint32 {
	if x != nil {
		return x.MaxConnections
	}
	return 0
}

func (x *NodeStatus) GetActiveConnections() uint32 {
	if x != nil {
		return x.ActiveConnections
	}
	return 0
}

var File_nodes_nodes_proto protoreflect.FileDescriptor

var file_nodes_nodes_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61,
	0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0x43, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x38, 0x0a,
	0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x65, 0x67, 0x6f, 0x68, 0x6f, 0x6c, 0x69, 0x76,
	0x65, 0x69, 0x72, 0x61, 0x2f, 0x67, 0x6f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nodes_nodes_proto_rawDescOnce sync.Once
	file_nodes_nodes_proto_rawDescData = file_nodes_nodes_proto_rawDesc
)

func file_nodes_nodes_proto_rawDescGZIP() []byte {
	file_nodes_nodes_proto_rawDescOnce.Do(func() {
		file_nodes_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_nodes_nodes_proto_rawDescData)
	})
	return file_nodes_nodes_proto_rawDescData
}

var file_nodes_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_nodes_nodes_proto_goTypes = []interface{}{
	(*NodeStatus)(nil),    // 0: nodes.NodeStatus
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_nodes_nodes_proto_depIdxs = []int32{
	0, // 0: nodes.Manager.SetStatus:input_type -> nodes.NodeStatus
	1, // 1: nodes.Manager.SetStatus:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nodes_nodes_proto_init() }
func file_nodes_nodes_proto_init() {
	if File_nodes_nodes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nodes_nodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeStatus); i {
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
			RawDescriptor: file_nodes_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nodes_nodes_proto_goTypes,
		DependencyIndexes: file_nodes_nodes_proto_depIdxs,
		MessageInfos:      file_nodes_nodes_proto_msgTypes,
	}.Build()
	File_nodes_nodes_proto = out.File
	file_nodes_nodes_proto_rawDesc = nil
	file_nodes_nodes_proto_goTypes = nil
	file_nodes_nodes_proto_depIdxs = nil
}

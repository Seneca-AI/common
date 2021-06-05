// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.2
// source: devops.proto

package _type

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

type SenecaServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId                                string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ServerExternalIp                         string `protobuf:"bytes,2,opt,name=server_external_ip,json=serverExternalIp,proto3" json:"server_external_ip,omitempty"`
	ServerPort                               string `protobuf:"bytes,3,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	ServerVmName                             string `protobuf:"bytes,4,opt,name=server_vm_name,json=serverVmName,proto3" json:"server_vm_name,omitempty"`
	ServerVmZone                             string `protobuf:"bytes,5,opt,name=server_vm_zone,json=serverVmZone,proto3" json:"server_vm_zone,omitempty"`
	ServerPathToGoogleApplicationCredentials string `protobuf:"bytes,6,opt,name=server_path_to_google_application_credentials,json=serverPathToGoogleApplicationCredentials,proto3" json:"server_path_to_google_application_credentials,omitempty"`
	ServerPathToGoogleOauthCredentials       string `protobuf:"bytes,7,opt,name=server_path_to_google_oauth_credentials,json=serverPathToGoogleOauthCredentials,proto3" json:"server_path_to_google_oauth_credentials,omitempty"`
	// Whether the server should be updated when main gets pushed to.
	ReceiveMainPushes bool  `protobuf:"varint,8,opt,name=receive_main_pushes,json=receiveMainPushes,proto3" json:"receive_main_pushes,omitempty"`
	DatastoreId       int64 `protobuf:"varint,9,opt,name=datastore_id,json=datastoreId,proto3" json:"datastore_id,omitempty"`
}

func (x *SenecaServer) Reset() {
	*x = SenecaServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_devops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SenecaServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SenecaServer) ProtoMessage() {}

func (x *SenecaServer) ProtoReflect() protoreflect.Message {
	mi := &file_devops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SenecaServer.ProtoReflect.Descriptor instead.
func (*SenecaServer) Descriptor() ([]byte, []int) {
	return file_devops_proto_rawDescGZIP(), []int{0}
}

func (x *SenecaServer) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *SenecaServer) GetServerExternalIp() string {
	if x != nil {
		return x.ServerExternalIp
	}
	return ""
}

func (x *SenecaServer) GetServerPort() string {
	if x != nil {
		return x.ServerPort
	}
	return ""
}

func (x *SenecaServer) GetServerVmName() string {
	if x != nil {
		return x.ServerVmName
	}
	return ""
}

func (x *SenecaServer) GetServerVmZone() string {
	if x != nil {
		return x.ServerVmZone
	}
	return ""
}

func (x *SenecaServer) GetServerPathToGoogleApplicationCredentials() string {
	if x != nil {
		return x.ServerPathToGoogleApplicationCredentials
	}
	return ""
}

func (x *SenecaServer) GetServerPathToGoogleOauthCredentials() string {
	if x != nil {
		return x.ServerPathToGoogleOauthCredentials
	}
	return ""
}

func (x *SenecaServer) GetReceiveMainPushes() bool {
	if x != nil {
		return x.ReceiveMainPushes
	}
	return false
}

func (x *SenecaServer) GetDatastoreId() int64 {
	if x != nil {
		return x.DatastoreId
	}
	return 0
}

var File_devops_proto protoreflect.FileDescriptor

var file_devops_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x64, 0x65, 0x76, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x22, 0xd1, 0x03, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x65, 0x63, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49,
	0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f,
	0x72, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x6d, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x56, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x76, 0x6d, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x6d, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x5f,
	0x0a, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x6f,
	0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x28, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x61, 0x74,
	0x68, 0x54, 0x6f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x53, 0x0a, 0x27, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74,
	0x6f, 0x5f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x22, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x61, 0x74, 0x68, 0x54, 0x6f, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f,
	0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x75,
	0x73, 0x68, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_devops_proto_rawDescOnce sync.Once
	file_devops_proto_rawDescData = file_devops_proto_rawDesc
)

func file_devops_proto_rawDescGZIP() []byte {
	file_devops_proto_rawDescOnce.Do(func() {
		file_devops_proto_rawDescData = protoimpl.X.CompressGZIP(file_devops_proto_rawDescData)
	})
	return file_devops_proto_rawDescData
}

var file_devops_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_devops_proto_goTypes = []interface{}{
	(*SenecaServer)(nil), // 0: api.SenecaServer
}
var file_devops_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_devops_proto_init() }
func file_devops_proto_init() {
	if File_devops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_devops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SenecaServer); i {
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
			RawDescriptor: file_devops_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_devops_proto_goTypes,
		DependencyIndexes: file_devops_proto_depIdxs,
		MessageInfos:      file_devops_proto_msgTypes,
	}.Build()
	File_devops_proto = out.File
	file_devops_proto_rawDesc = nil
	file_devops_proto_goTypes = nil
	file_devops_proto_depIdxs = nil
}

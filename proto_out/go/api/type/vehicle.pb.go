// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.16.0
// source: vehicle.proto

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

// key(uuid)
type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid              string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FriendlyName      string   `protobuf:"bytes,2,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	PrimaryDriverId   string   `protobuf:"bytes,3,opt,name=primary_driver_id,json=primaryDriverId,proto3" json:"primary_driver_id,omitempty"`
	SecondaryDriverId []string `protobuf:"bytes,4,rep,name=secondary_driver_id,json=secondaryDriverId,proto3" json:"secondary_driver_id,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *Vehicle) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Vehicle) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *Vehicle) GetPrimaryDriverId() string {
	if x != nil {
		return x.PrimaryDriverId
	}
	return ""
}

func (x *Vehicle) GetSecondaryDriverId() []string {
	if x != nil {
		return x.SecondaryDriverId
	}
	return nil
}

var File_vehicle_proto protoreflect.FileDescriptor

var file_vehicle_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x13, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vehicle_proto_rawDescOnce sync.Once
	file_vehicle_proto_rawDescData = file_vehicle_proto_rawDesc
)

func file_vehicle_proto_rawDescGZIP() []byte {
	file_vehicle_proto_rawDescOnce.Do(func() {
		file_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_vehicle_proto_rawDescData)
	})
	return file_vehicle_proto_rawDescData
}

var file_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vehicle_proto_goTypes = []interface{}{
	(*Vehicle)(nil), // 0: Vehicle
}
var file_vehicle_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vehicle_proto_init() }
func file_vehicle_proto_init() {
	if File_vehicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
			RawDescriptor: file_vehicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vehicle_proto_goTypes,
		DependencyIndexes: file_vehicle_proto_depIdxs,
		MessageInfos:      file_vehicle_proto_msgTypes,
	}.Build()
	File_vehicle_proto = out.File
	file_vehicle_proto_rawDesc = nil
	file_vehicle_proto_goTypes = nil
	file_vehicle_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.16.0
// source: common.proto

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

type Latitude_LatDirection int32

const (
	Latitude_UNKNOWN Latitude_LatDirection = 0
	Latitude_NORTH   Latitude_LatDirection = 1
	Latitude_SOUTH   Latitude_LatDirection = 2
)

// Enum value maps for Latitude_LatDirection.
var (
	Latitude_LatDirection_name = map[int32]string{
		0: "UNKNOWN",
		1: "NORTH",
		2: "SOUTH",
	}
	Latitude_LatDirection_value = map[string]int32{
		"UNKNOWN": 0,
		"NORTH":   1,
		"SOUTH":   2,
	}
)

func (x Latitude_LatDirection) Enum() *Latitude_LatDirection {
	p := new(Latitude_LatDirection)
	*p = x
	return p
}

func (x Latitude_LatDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Latitude_LatDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (Latitude_LatDirection) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x Latitude_LatDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Latitude_LatDirection.Descriptor instead.
func (Latitude_LatDirection) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1, 0}
}

type Longitude_LongDirection int32

const (
	Longitude_UNKNOWN Longitude_LongDirection = 0
	Longitude_EAST    Longitude_LongDirection = 1
	Longitude_WEST    Longitude_LongDirection = 2
)

// Enum value maps for Longitude_LongDirection.
var (
	Longitude_LongDirection_name = map[int32]string{
		0: "UNKNOWN",
		1: "EAST",
		2: "WEST",
	}
	Longitude_LongDirection_value = map[string]int32{
		"UNKNOWN": 0,
		"EAST":    1,
		"WEST":    2,
	}
)

func (x Longitude_LongDirection) Enum() *Longitude_LongDirection {
	p := new(Longitude_LongDirection)
	*p = x
	return p
}

func (x Longitude_LongDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Longitude_LongDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (Longitude_LongDirection) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x Longitude_LongDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Longitude_LongDirection.Descriptor instead.
func (Longitude_LongDirection) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2, 0}
}

type Source_Type int32

const (
	Source_UNKNOWN      Source_Type = 0
	Source_RAW_VIDEO    Source_Type = 1
	Source_RAW_LOCATION Source_Type = 2
	Source_RAW_MOTION   Source_Type = 3
)

// Enum value maps for Source_Type.
var (
	Source_Type_name = map[int32]string{
		0: "UNKNOWN",
		1: "RAW_VIDEO",
		2: "RAW_LOCATION",
		3: "RAW_MOTION",
	}
	Source_Type_value = map[string]int32{
		"UNKNOWN":      0,
		"RAW_VIDEO":    1,
		"RAW_LOCATION": 2,
		"RAW_MOTION":   3,
	}
)

func (x Source_Type) Enum() *Source_Type {
	p := new(Source_Type)
	*p = x
	return p
}

func (x Source_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Source_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (Source_Type) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x Source_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Source_Type.Descriptor instead.
func (Source_Type) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5, 0}
}

type TimePeriod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimeMs int64 `protobuf:"varint,1,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
	EndTimeMs   int64 `protobuf:"varint,2,opt,name=end_time_ms,json=endTimeMs,proto3" json:"end_time_ms,omitempty"`
}

func (x *TimePeriod) Reset() {
	*x = TimePeriod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimePeriod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimePeriod) ProtoMessage() {}

func (x *TimePeriod) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimePeriod.ProtoReflect.Descriptor instead.
func (*TimePeriod) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *TimePeriod) GetStartTimeMs() int64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

func (x *TimePeriod) GetEndTimeMs() int64 {
	if x != nil {
		return x.EndTimeMs
	}
	return 0
}

type Latitude struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Degrees       float64               `protobuf:"fixed64,1,opt,name=degrees,proto3" json:"degrees,omitempty"`
	DegreeMinutes float64               `protobuf:"fixed64,2,opt,name=degree_minutes,json=degreeMinutes,proto3" json:"degree_minutes,omitempty"`
	DegreeSeconds float64               `protobuf:"fixed64,3,opt,name=degree_seconds,json=degreeSeconds,proto3" json:"degree_seconds,omitempty"`
	LatDirection  Latitude_LatDirection `protobuf:"varint,4,opt,name=lat_direction,json=latDirection,proto3,enum=api.Latitude_LatDirection" json:"lat_direction,omitempty"`
}

func (x *Latitude) Reset() {
	*x = Latitude{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Latitude) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Latitude) ProtoMessage() {}

func (x *Latitude) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Latitude.ProtoReflect.Descriptor instead.
func (*Latitude) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Latitude) GetDegrees() float64 {
	if x != nil {
		return x.Degrees
	}
	return 0
}

func (x *Latitude) GetDegreeMinutes() float64 {
	if x != nil {
		return x.DegreeMinutes
	}
	return 0
}

func (x *Latitude) GetDegreeSeconds() float64 {
	if x != nil {
		return x.DegreeSeconds
	}
	return 0
}

func (x *Latitude) GetLatDirection() Latitude_LatDirection {
	if x != nil {
		return x.LatDirection
	}
	return Latitude_UNKNOWN
}

type Longitude struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Degrees       float64                 `protobuf:"fixed64,1,opt,name=degrees,proto3" json:"degrees,omitempty"`
	DegreeMinutes float64                 `protobuf:"fixed64,2,opt,name=degree_minutes,json=degreeMinutes,proto3" json:"degree_minutes,omitempty"`
	DegreeSeconds float64                 `protobuf:"fixed64,3,opt,name=degree_seconds,json=degreeSeconds,proto3" json:"degree_seconds,omitempty"`
	LongDirection Longitude_LongDirection `protobuf:"varint,4,opt,name=long_direction,json=longDirection,proto3,enum=api.Longitude_LongDirection" json:"long_direction,omitempty"`
}

func (x *Longitude) Reset() {
	*x = Longitude{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Longitude) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Longitude) ProtoMessage() {}

func (x *Longitude) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Longitude.ProtoReflect.Descriptor instead.
func (*Longitude) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Longitude) GetDegrees() float64 {
	if x != nil {
		return x.Degrees
	}
	return 0
}

func (x *Longitude) GetDegreeMinutes() float64 {
	if x != nil {
		return x.DegreeMinutes
	}
	return 0
}

func (x *Longitude) GetDegreeSeconds() float64 {
	if x != nil {
		return x.DegreeSeconds
	}
	return 0
}

func (x *Longitude) GetLongDirection() Longitude_LongDirection {
	if x != nil {
		return x.LongDirection
	}
	return Longitude_UNKNOWN
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat  *Latitude  `protobuf:"bytes,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Long *Longitude `protobuf:"bytes,2,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetLat() *Latitude {
	if x != nil {
		return x.Lat
	}
	return nil
}

func (x *Location) GetLong() *Longitude {
	if x != nil {
		return x.Long
	}
	return nil
}

type Motion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VelocityMph      float64 `protobuf:"fixed64,1,opt,name=velocity_mph,json=velocityMph,proto3" json:"velocity_mph,omitempty"`
	AccelerationMphS float64 `protobuf:"fixed64,2,opt,name=acceleration_mph_s,json=accelerationMphS,proto3" json:"acceleration_mph_s,omitempty"`
}

func (x *Motion) Reset() {
	*x = Motion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Motion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Motion) ProtoMessage() {}

func (x *Motion) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Motion.ProtoReflect.Descriptor instead.
func (*Motion) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *Motion) GetVelocityMph() float64 {
	if x != nil {
		return x.VelocityMph
	}
	return 0
}

func (x *Motion) GetAccelerationMphS() float64 {
	if x != nil {
		return x.AccelerationMphS
	}
	return 0
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId   string      `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	SourceType Source_Type `protobuf:"varint,2,opt,name=source_type,json=sourceType,proto3,enum=api.Source_Type" json:"source_type,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *Source) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *Source) GetSourceType() Source_Type {
	if x != nil {
		return x.SourceType
	}
	return Source_UNKNOWN
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x22, 0x50, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x22, 0xe6, 0x01, 0x0a, 0x08, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x4d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x6c, 0x61,
	0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x2e, 0x4c, 0x61, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6c,
	0x61, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x31, 0x0a, 0x0c, 0x4c,
	0x61, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4e, 0x4f, 0x52, 0x54,
	0x48, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x4f, 0x55, 0x54, 0x48, 0x10, 0x02, 0x22, 0xea,
	0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x64,
	0x65, 0x67, 0x72, 0x65, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65,
	0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d,
	0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x43, 0x0a, 0x0e, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x2e, 0x4c, 0x6f, 0x6e,
	0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6c, 0x6f, 0x6e, 0x67,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x0d, 0x4c, 0x6f, 0x6e,
	0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x41, 0x53, 0x54, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x45, 0x53, 0x54, 0x10, 0x02, 0x22, 0x4f, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x22, 0x59, 0x0a, 0x06,
	0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x6d, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x76, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x4d, 0x70, 0x68, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x63, 0x63,
	0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x70, 0x68, 0x5f, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x70, 0x68, 0x53, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x31, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x57, 0x5f, 0x56,
	0x49, 0x44, 0x45, 0x4f, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x41, 0x57, 0x5f, 0x4c, 0x4f,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x41, 0x57, 0x5f,
	0x4d, 0x4f, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_goTypes = []interface{}{
	(Latitude_LatDirection)(0),   // 0: api.Latitude.LatDirection
	(Longitude_LongDirection)(0), // 1: api.Longitude.LongDirection
	(Source_Type)(0),             // 2: api.Source.Type
	(*TimePeriod)(nil),           // 3: api.TimePeriod
	(*Latitude)(nil),             // 4: api.Latitude
	(*Longitude)(nil),            // 5: api.Longitude
	(*Location)(nil),             // 6: api.Location
	(*Motion)(nil),               // 7: api.Motion
	(*Source)(nil),               // 8: api.Source
}
var file_common_proto_depIdxs = []int32{
	0, // 0: api.Latitude.lat_direction:type_name -> api.Latitude.LatDirection
	1, // 1: api.Longitude.long_direction:type_name -> api.Longitude.LongDirection
	4, // 2: api.Location.lat:type_name -> api.Latitude
	5, // 3: api.Location.long:type_name -> api.Longitude
	2, // 4: api.Source.source_type:type_name -> api.Source.Type
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimePeriod); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Latitude); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Longitude); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Motion); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Source); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}

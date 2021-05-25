// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.1
// source: external.proto

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

type ExternalSource_ExternalSourceType int32

const (
	ExternalSource_UNKNOWN       ExternalSource_ExternalSourceType = 0
	ExternalSource_DASHCAM_VIDEO ExternalSource_ExternalSourceType = 1
)

// Enum value maps for ExternalSource_ExternalSourceType.
var (
	ExternalSource_ExternalSourceType_name = map[int32]string{
		0: "UNKNOWN",
		1: "DASHCAM_VIDEO",
	}
	ExternalSource_ExternalSourceType_value = map[string]int32{
		"UNKNOWN":       0,
		"DASHCAM_VIDEO": 1,
	}
)

func (x ExternalSource_ExternalSourceType) Enum() *ExternalSource_ExternalSourceType {
	p := new(ExternalSource_ExternalSourceType)
	*p = x
	return p
}

func (x ExternalSource_ExternalSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExternalSource_ExternalSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_external_proto_enumTypes[0].Descriptor()
}

func (ExternalSource_ExternalSourceType) Type() protoreflect.EnumType {
	return &file_external_proto_enumTypes[0]
}

func (x ExternalSource_ExternalSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExternalSource_ExternalSourceType.Descriptor instead.
func (ExternalSource_ExternalSourceType) EnumDescriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{1, 0}
}

type Trip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimeMs      int64               `protobuf:"varint,1,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
	EndTimeMs        int64               `protobuf:"varint,2,opt,name=end_time_ms,json=endTimeMs,proto3" json:"end_time_ms,omitempty"`
	Event            []*Event            `protobuf:"bytes,3,rep,name=event,proto3" json:"event,omitempty"`
	DrivingCondition []*DrivingCondition `protobuf:"bytes,4,rep,name=driving_condition,json=drivingCondition,proto3" json:"driving_condition,omitempty"`
}

func (x *Trip) Reset() {
	*x = Trip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trip) ProtoMessage() {}

func (x *Trip) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trip.ProtoReflect.Descriptor instead.
func (*Trip) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{0}
}

func (x *Trip) GetStartTimeMs() int64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

func (x *Trip) GetEndTimeMs() int64 {
	if x != nil {
		return x.EndTimeMs
	}
	return 0
}

func (x *Trip) GetEvent() []*Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *Trip) GetDrivingCondition() []*DrivingCondition {
	if x != nil {
		return x.DrivingCondition
	}
	return nil
}

type ExternalSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceType ExternalSource_ExternalSourceType `protobuf:"varint,1,opt,name=source_type,json=sourceType,proto3,enum=api.ExternalSource_ExternalSourceType" json:"source_type,omitempty"`
	VideoUrl   string                            `protobuf:"bytes,2,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
}

func (x *ExternalSource) Reset() {
	*x = ExternalSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalSource) ProtoMessage() {}

func (x *ExternalSource) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalSource.ProtoReflect.Descriptor instead.
func (*ExternalSource) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalSource) GetSourceType() ExternalSource_ExternalSourceType {
	if x != nil {
		return x.SourceType
	}
	return ExternalSource_UNKNOWN
}

func (x *ExternalSource) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

type DrivingCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConditionType  []ConditionType `protobuf:"varint,2,rep,packed,name=condition_type,json=conditionType,proto3,enum=api.ConditionType" json:"condition_type,omitempty"`
	Severity       []float64       `protobuf:"fixed64,5,rep,packed,name=severity,proto3" json:"severity,omitempty"`
	StartTimeMs    int64           `protobuf:"varint,6,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
	EndTimeMs      int64           `protobuf:"varint,7,opt,name=end_time_ms,json=endTimeMs,proto3" json:"end_time_ms,omitempty"`
	ExternalSource *ExternalSource `protobuf:"bytes,8,opt,name=external_source,json=externalSource,proto3" json:"external_source,omitempty"`
}

func (x *DrivingCondition) Reset() {
	*x = DrivingCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrivingCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrivingCondition) ProtoMessage() {}

func (x *DrivingCondition) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrivingCondition.ProtoReflect.Descriptor instead.
func (*DrivingCondition) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{2}
}

func (x *DrivingCondition) GetConditionType() []ConditionType {
	if x != nil {
		return x.ConditionType
	}
	return nil
}

func (x *DrivingCondition) GetSeverity() []float64 {
	if x != nil {
		return x.Severity
	}
	return nil
}

func (x *DrivingCondition) GetStartTimeMs() int64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

func (x *DrivingCondition) GetEndTimeMs() int64 {
	if x != nil {
		return x.EndTimeMs
	}
	return 0
}

func (x *DrivingCondition) GetExternalSource() *ExternalSource {
	if x != nil {
		return x.ExternalSource
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType EventType `protobuf:"varint,3,opt,name=event_type,json=eventType,proto3,enum=api.EventType" json:"event_type,omitempty"`
	// Value depends on type.  For example, FAST_ACCELERATION will be in mph/s.
	Value          float64         `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
	Severity       float64         `protobuf:"fixed64,5,opt,name=severity,proto3" json:"severity,omitempty"`
	TimestampMs    int64           `protobuf:"varint,6,opt,name=timestamp_ms,json=timestampMs,proto3" json:"timestamp_ms,omitempty"`
	ExternalSource *ExternalSource `protobuf:"bytes,7,opt,name=external_source,json=externalSource,proto3" json:"external_source,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{3}
}

func (x *Event) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_UNKNOWN_EVENT_TYPE
}

func (x *Event) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Event) GetSeverity() float64 {
	if x != nil {
		return x.Severity
	}
	return 0
}

func (x *Event) GetTimestampMs() int64 {
	if x != nil {
		return x.TimestampMs
	}
	return 0
}

func (x *Event) GetExternalSource() *ExternalSource {
	if x != nil {
		return x.ExternalSource
	}
	return nil
}

type TripListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartTimeMs int64  `protobuf:"varint,2,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
	EndTimeMs   int64  `protobuf:"varint,3,opt,name=end_time_ms,json=endTimeMs,proto3" json:"end_time_ms,omitempty"`
}

func (x *TripListRequest) Reset() {
	*x = TripListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TripListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripListRequest) ProtoMessage() {}

func (x *TripListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripListRequest.ProtoReflect.Descriptor instead.
func (*TripListRequest) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{4}
}

func (x *TripListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TripListRequest) GetStartTimeMs() int64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

func (x *TripListRequest) GetEndTimeMs() int64 {
	if x != nil {
		return x.EndTimeMs
	}
	return 0
}

type TripListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	UserId string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Trip   []*Trip `protobuf:"bytes,3,rep,name=trip,proto3" json:"trip,omitempty"`
}

func (x *TripListResponse) Reset() {
	*x = TripListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TripListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripListResponse) ProtoMessage() {}

func (x *TripListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripListResponse.ProtoReflect.Descriptor instead.
func (*TripListResponse) Descriptor() ([]byte, []int) {
	return file_external_proto_rawDescGZIP(), []int{5}
}

func (x *TripListResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *TripListResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TripListResponse) GetTrip() []*Trip {
	if x != nil {
		return x.Trip
	}
	return nil
}

var File_external_proto protoreflect.FileDescriptor

var file_external_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x04, 0x54, 0x72, 0x69, 0x70, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73,
	0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73,
	0x12, 0x20, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x42, 0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0b, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x22,
	0x34, 0x0a, 0x12, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x41, 0x53, 0x48, 0x43, 0x41, 0x4d, 0x5f, 0x56, 0x49,
	0x44, 0x45, 0x4f, 0x10, 0x01, 0x22, 0xeb, 0x01, 0x0a, 0x10, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x01, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x3c, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0xc9, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a,
	0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d,
	0x73, 0x12, 0x3c, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22,
	0x6e, 0x0a, 0x0f, 0x54, 0x72, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12,
	0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x22,
	0x6f, 0x0a, 0x10, 0x54, 0x72, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x04, 0x74, 0x72, 0x69, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x72, 0x69, 0x70, 0x52, 0x04, 0x74, 0x72, 0x69, 0x70,
	0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_proto_rawDescOnce sync.Once
	file_external_proto_rawDescData = file_external_proto_rawDesc
)

func file_external_proto_rawDescGZIP() []byte {
	file_external_proto_rawDescOnce.Do(func() {
		file_external_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_proto_rawDescData)
	})
	return file_external_proto_rawDescData
}

var file_external_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_external_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_proto_goTypes = []interface{}{
	(ExternalSource_ExternalSourceType)(0), // 0: api.ExternalSource.ExternalSourceType
	(*Trip)(nil),                           // 1: api.Trip
	(*ExternalSource)(nil),                 // 2: api.ExternalSource
	(*DrivingCondition)(nil),               // 3: api.DrivingCondition
	(*Event)(nil),                          // 4: api.Event
	(*TripListRequest)(nil),                // 5: api.TripListRequest
	(*TripListResponse)(nil),               // 6: api.TripListResponse
	(ConditionType)(0),                     // 7: api.ConditionType
	(EventType)(0),                         // 8: api.EventType
	(*Header)(nil),                         // 9: api.Header
}
var file_external_proto_depIdxs = []int32{
	4, // 0: api.Trip.event:type_name -> api.Event
	3, // 1: api.Trip.driving_condition:type_name -> api.DrivingCondition
	0, // 2: api.ExternalSource.source_type:type_name -> api.ExternalSource.ExternalSourceType
	7, // 3: api.DrivingCondition.condition_type:type_name -> api.ConditionType
	2, // 4: api.DrivingCondition.external_source:type_name -> api.ExternalSource
	8, // 5: api.Event.event_type:type_name -> api.EventType
	2, // 6: api.Event.external_source:type_name -> api.ExternalSource
	9, // 7: api.TripListResponse.header:type_name -> api.Header
	1, // 8: api.TripListResponse.trip:type_name -> api.Trip
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_external_proto_init() }
func file_external_proto_init() {
	if File_external_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_external_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trip); i {
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
		file_external_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalSource); i {
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
		file_external_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrivingCondition); i {
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
		file_external_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_external_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TripListRequest); i {
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
		file_external_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TripListResponse); i {
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
			RawDescriptor: file_external_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_proto_goTypes,
		DependencyIndexes: file_external_proto_depIdxs,
		EnumInfos:         file_external_proto_enumTypes,
		MessageInfos:      file_external_proto_msgTypes,
	}.Build()
	File_external_proto = out.File
	file_external_proto_rawDesc = nil
	file_external_proto_goTypes = nil
	file_external_proto_depIdxs = nil
}

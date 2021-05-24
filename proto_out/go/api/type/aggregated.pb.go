// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.0
// source: aggregated.proto

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

type TripInternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	StartTimeMs int64  `protobuf:"varint,3,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
	EndTimeMs   int64  `protobuf:"varint,4,opt,name=end_time_ms,json=endTimeMs,proto3" json:"end_time_ms,omitempty"`
}

func (x *TripInternal) Reset() {
	*x = TripInternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TripInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TripInternal) ProtoMessage() {}

func (x *TripInternal) ProtoReflect() protoreflect.Message {
	mi := &file_aggregated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TripInternal.ProtoReflect.Descriptor instead.
func (*TripInternal) Descriptor() ([]byte, []int) {
	return file_aggregated_proto_rawDescGZIP(), []int{0}
}

func (x *TripInternal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TripInternal) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TripInternal) GetStartTimeMs() int64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

func (x *TripInternal) GetEndTimeMs() int64 {
	if x != nil {
		return x.EndTimeMs
	}
	return 0
}

type DrivingConditionInternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string        `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TripId        string        `protobuf:"bytes,3,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
	ConditionType ConditionType `protobuf:"varint,4,opt,name=condition_type,json=conditionType,proto3,enum=api.ConditionType" json:"condition_type,omitempty"`
	Severity      float64       `protobuf:"fixed64,5,opt,name=severity,proto3" json:"severity,omitempty"`
	StartTimeMs   int64         `protobuf:"varint,6,opt,name=start_time_ms,json=startTimeMs,proto3" json:"start_time_ms,omitempty"`
	EndTimeMs     int64         `protobuf:"varint,7,opt,name=end_time_ms,json=endTimeMs,proto3" json:"end_time_ms,omitempty"`
	Source        *Source       `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	// The algorithms this drivingCondition has gone through.
	AlgoTag []string `protobuf:"bytes,9,rep,name=algoTag,proto3" json:"algoTag,omitempty"`
}

func (x *DrivingConditionInternal) Reset() {
	*x = DrivingConditionInternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrivingConditionInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrivingConditionInternal) ProtoMessage() {}

func (x *DrivingConditionInternal) ProtoReflect() protoreflect.Message {
	mi := &file_aggregated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrivingConditionInternal.ProtoReflect.Descriptor instead.
func (*DrivingConditionInternal) Descriptor() ([]byte, []int) {
	return file_aggregated_proto_rawDescGZIP(), []int{1}
}

func (x *DrivingConditionInternal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DrivingConditionInternal) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DrivingConditionInternal) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

func (x *DrivingConditionInternal) GetConditionType() ConditionType {
	if x != nil {
		return x.ConditionType
	}
	return ConditionType_NONE_CONDITION_TYPE
}

func (x *DrivingConditionInternal) GetSeverity() float64 {
	if x != nil {
		return x.Severity
	}
	return 0
}

func (x *DrivingConditionInternal) GetStartTimeMs() int64 {
	if x != nil {
		return x.StartTimeMs
	}
	return 0
}

func (x *DrivingConditionInternal) GetEndTimeMs() int64 {
	if x != nil {
		return x.EndTimeMs
	}
	return 0
}

func (x *DrivingConditionInternal) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *DrivingConditionInternal) GetAlgoTag() []string {
	if x != nil {
		return x.AlgoTag
	}
	return nil
}

type EventInternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string    `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TripId    string    `protobuf:"bytes,3,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
	EventType EventType `protobuf:"varint,4,opt,name=event_type,json=eventType,proto3,enum=api.EventType" json:"event_type,omitempty"`
	// Value depends on type.  For example, FAST_ACCELERATION will be in mph/s.
	Value       float64 `protobuf:"fixed64,5,opt,name=value,proto3" json:"value,omitempty"`
	Severity    float64 `protobuf:"fixed64,6,opt,name=severity,proto3" json:"severity,omitempty"`
	TimestampMs int64   `protobuf:"varint,7,opt,name=timestamp_ms,json=timestampMs,proto3" json:"timestamp_ms,omitempty"`
	Source      *Source `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	// The algorithms this drivingCondition has gone through.
	AlgoTag []string `protobuf:"bytes,9,rep,name=algoTag,proto3" json:"algoTag,omitempty"`
}

func (x *EventInternal) Reset() {
	*x = EventInternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventInternal) ProtoMessage() {}

func (x *EventInternal) ProtoReflect() protoreflect.Message {
	mi := &file_aggregated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventInternal.ProtoReflect.Descriptor instead.
func (*EventInternal) Descriptor() ([]byte, []int) {
	return file_aggregated_proto_rawDescGZIP(), []int{2}
}

func (x *EventInternal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventInternal) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventInternal) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

func (x *EventInternal) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_UNKNOWN_EVENT_TYPE
}

func (x *EventInternal) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *EventInternal) GetSeverity() float64 {
	if x != nil {
		return x.Severity
	}
	return 0
}

func (x *EventInternal) GetTimestampMs() int64 {
	if x != nil {
		return x.TimestampMs
	}
	return 0
}

func (x *EventInternal) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *EventInternal) GetAlgoTag() []string {
	if x != nil {
		return x.AlgoTag
	}
	return nil
}

type EventCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string         `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Event  *EventInternal `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventCreateRequest) Reset() {
	*x = EventCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCreateRequest) ProtoMessage() {}

func (x *EventCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventCreateRequest.ProtoReflect.Descriptor instead.
func (*EventCreateRequest) Descriptor() ([]byte, []int) {
	return file_aggregated_proto_rawDescGZIP(), []int{3}
}

func (x *EventCreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventCreateRequest) GetEvent() *EventInternal {
	if x != nil {
		return x.Event
	}
	return nil
}

type EventCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	UserId string         `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Event  *EventInternal `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *EventCreateResponse) Reset() {
	*x = EventCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventCreateResponse) ProtoMessage() {}

func (x *EventCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventCreateResponse.ProtoReflect.Descriptor instead.
func (*EventCreateResponse) Descriptor() ([]byte, []int) {
	return file_aggregated_proto_rawDescGZIP(), []int{4}
}

func (x *EventCreateResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *EventCreateResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EventCreateResponse) GetEvent() *EventInternal {
	if x != nil {
		return x.Event
	}
	return nil
}

type DrivingConditionCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId           string                    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DrivingCondition *DrivingConditionInternal `protobuf:"bytes,2,opt,name=driving_condition,json=drivingCondition,proto3" json:"driving_condition,omitempty"`
}

func (x *DrivingConditionCreateRequest) Reset() {
	*x = DrivingConditionCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregated_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrivingConditionCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrivingConditionCreateRequest) ProtoMessage() {}

func (x *DrivingConditionCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aggregated_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrivingConditionCreateRequest.ProtoReflect.Descriptor instead.
func (*DrivingConditionCreateRequest) Descriptor() ([]byte, []int) {
	return file_aggregated_proto_rawDescGZIP(), []int{5}
}

func (x *DrivingConditionCreateRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DrivingConditionCreateRequest) GetDrivingCondition() *DrivingConditionInternal {
	if x != nil {
		return x.DrivingCondition
	}
	return nil
}

type DrivingConditionCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header           *Header                   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	UserId           string                    `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DrivingCondition *DrivingConditionInternal `protobuf:"bytes,3,opt,name=driving_condition,json=drivingCondition,proto3" json:"driving_condition,omitempty"`
}

func (x *DrivingConditionCreateResponse) Reset() {
	*x = DrivingConditionCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aggregated_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrivingConditionCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrivingConditionCreateResponse) ProtoMessage() {}

func (x *DrivingConditionCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aggregated_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrivingConditionCreateResponse.ProtoReflect.Descriptor instead.
func (*DrivingConditionCreateResponse) Descriptor() ([]byte, []int) {
	return file_aggregated_proto_rawDescGZIP(), []int{6}
}

func (x *DrivingConditionCreateResponse) GetHeader() *Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *DrivingConditionCreateResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DrivingConditionCreateResponse) GetDrivingCondition() *DrivingConditionInternal {
	if x != nil {
		return x.DrivingCondition
	}
	return nil
}

var File_aggregated_proto protoreflect.FileDescriptor

var file_aggregated_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x0c, 0x54, 0x72, 0x69, 0x70, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x4d, 0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x4d, 0x73, 0x22, 0xb6, 0x02, 0x0a, 0x18, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1e, 0x0a, 0x0b,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x23, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x67, 0x6f, 0x54, 0x61, 0x67, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x67, 0x6f, 0x54, 0x61, 0x67, 0x22, 0x94, 0x02, 0x0a, 0x0d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x4d, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x67, 0x6f,
	0x54, 0x61, 0x67, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x67, 0x6f, 0x54,
	0x61, 0x67, 0x22, 0x57, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x13, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x1d, 0x44,
	0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52,
	0x10, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xaa, 0x01, 0x0a, 0x1e, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x4a, 0x0a, 0x11, 0x64, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x10, 0x64, 0x72,
	0x69, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a,
	0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_aggregated_proto_rawDescOnce sync.Once
	file_aggregated_proto_rawDescData = file_aggregated_proto_rawDesc
)

func file_aggregated_proto_rawDescGZIP() []byte {
	file_aggregated_proto_rawDescOnce.Do(func() {
		file_aggregated_proto_rawDescData = protoimpl.X.CompressGZIP(file_aggregated_proto_rawDescData)
	})
	return file_aggregated_proto_rawDescData
}

var file_aggregated_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_aggregated_proto_goTypes = []interface{}{
	(*TripInternal)(nil),                   // 0: api.TripInternal
	(*DrivingConditionInternal)(nil),       // 1: api.DrivingConditionInternal
	(*EventInternal)(nil),                  // 2: api.EventInternal
	(*EventCreateRequest)(nil),             // 3: api.EventCreateRequest
	(*EventCreateResponse)(nil),            // 4: api.EventCreateResponse
	(*DrivingConditionCreateRequest)(nil),  // 5: api.DrivingConditionCreateRequest
	(*DrivingConditionCreateResponse)(nil), // 6: api.DrivingConditionCreateResponse
	(ConditionType)(0),                     // 7: api.ConditionType
	(*Source)(nil),                         // 8: api.Source
	(EventType)(0),                         // 9: api.EventType
	(*Header)(nil),                         // 10: api.Header
}
var file_aggregated_proto_depIdxs = []int32{
	7,  // 0: api.DrivingConditionInternal.condition_type:type_name -> api.ConditionType
	8,  // 1: api.DrivingConditionInternal.source:type_name -> api.Source
	9,  // 2: api.EventInternal.event_type:type_name -> api.EventType
	8,  // 3: api.EventInternal.source:type_name -> api.Source
	2,  // 4: api.EventCreateRequest.event:type_name -> api.EventInternal
	10, // 5: api.EventCreateResponse.header:type_name -> api.Header
	2,  // 6: api.EventCreateResponse.event:type_name -> api.EventInternal
	1,  // 7: api.DrivingConditionCreateRequest.driving_condition:type_name -> api.DrivingConditionInternal
	10, // 8: api.DrivingConditionCreateResponse.header:type_name -> api.Header
	1,  // 9: api.DrivingConditionCreateResponse.driving_condition:type_name -> api.DrivingConditionInternal
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_aggregated_proto_init() }
func file_aggregated_proto_init() {
	if File_aggregated_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aggregated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TripInternal); i {
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
		file_aggregated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrivingConditionInternal); i {
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
		file_aggregated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventInternal); i {
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
		file_aggregated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCreateRequest); i {
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
		file_aggregated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventCreateResponse); i {
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
		file_aggregated_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrivingConditionCreateRequest); i {
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
		file_aggregated_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrivingConditionCreateResponse); i {
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
			RawDescriptor: file_aggregated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aggregated_proto_goTypes,
		DependencyIndexes: file_aggregated_proto_depIdxs,
		MessageInfos:      file_aggregated_proto_msgTypes,
	}.Build()
	File_aggregated_proto = out.File
	file_aggregated_proto_rawDesc = nil
	file_aggregated_proto_goTypes = nil
	file_aggregated_proto_depIdxs = nil
}

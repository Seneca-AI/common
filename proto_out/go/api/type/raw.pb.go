// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.0
// source: raw.proto

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

// key(user_id, id)
type RawVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// In the form <USER_ID>.<CREATE_TIMESTAMP_MS>.RAW_VIDEO.mp4
	CloudStorageFileName string `protobuf:"bytes,3,opt,name=cloud_storage_file_name,json=cloudStorageFileName,proto3" json:"cloud_storage_file_name,omitempty"`
	// When the video begins in ms.
	CreateTimeMs int64    `protobuf:"varint,4,opt,name=create_time_ms,json=createTimeMs,proto3" json:"create_time_ms,omitempty"`
	CutVideoId   []string `protobuf:"bytes,5,rep,name=cut_video_id,json=cutVideoId,proto3" json:"cut_video_id,omitempty"`
	DurationMs   int64    `protobuf:"varint,6,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	// Used for giving users feedback on which files resulted in error.
	OriginalFileName string `protobuf:"bytes,7,opt,name=original_file_name,json=originalFileName,proto3" json:"original_file_name,omitempty"`
	FailureReason    string `protobuf:"bytes,8,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"`
	OriginalFileId   string `protobuf:"bytes,9,opt,name=original_file_id,json=originalFileId,proto3" json:"original_file_id,omitempty"`
}

func (x *RawVideo) Reset() {
	*x = RawVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawVideo) ProtoMessage() {}

func (x *RawVideo) ProtoReflect() protoreflect.Message {
	mi := &file_raw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawVideo.ProtoReflect.Descriptor instead.
func (*RawVideo) Descriptor() ([]byte, []int) {
	return file_raw_proto_rawDescGZIP(), []int{0}
}

func (x *RawVideo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RawVideo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RawVideo) GetCloudStorageFileName() string {
	if x != nil {
		return x.CloudStorageFileName
	}
	return ""
}

func (x *RawVideo) GetCreateTimeMs() int64 {
	if x != nil {
		return x.CreateTimeMs
	}
	return 0
}

func (x *RawVideo) GetCutVideoId() []string {
	if x != nil {
		return x.CutVideoId
	}
	return nil
}

func (x *RawVideo) GetDurationMs() int64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *RawVideo) GetOriginalFileName() string {
	if x != nil {
		return x.OriginalFileName
	}
	return ""
}

func (x *RawVideo) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

func (x *RawVideo) GetOriginalFileId() string {
	if x != nil {
		return x.OriginalFileId
	}
	return ""
}

type RawVideoProcessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoName  string `protobuf:"bytes,2,opt,name=video_name,json=videoName,proto3" json:"video_name,omitempty"`
	LocalPath  string `protobuf:"bytes,3,opt,name=local_path,json=localPath,proto3" json:"local_path,omitempty"`
	VideoBytes []byte `protobuf:"bytes,4,opt,name=video_bytes,json=videoBytes,proto3" json:"video_bytes,omitempty"`
}

func (x *RawVideoProcessRequest) Reset() {
	*x = RawVideoProcessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawVideoProcessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawVideoProcessRequest) ProtoMessage() {}

func (x *RawVideoProcessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_raw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawVideoProcessRequest.ProtoReflect.Descriptor instead.
func (*RawVideoProcessRequest) Descriptor() ([]byte, []int) {
	return file_raw_proto_rawDescGZIP(), []int{1}
}

func (x *RawVideoProcessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RawVideoProcessRequest) GetVideoName() string {
	if x != nil {
		return x.VideoName
	}
	return ""
}

func (x *RawVideoProcessRequest) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

func (x *RawVideoProcessRequest) GetVideoBytes() []byte {
	if x != nil {
		return x.VideoBytes
	}
	return nil
}

type RawVideoProcessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RawVideoId string `protobuf:"bytes,1,opt,name=raw_video_id,json=rawVideoId,proto3" json:"raw_video_id,omitempty"`
}

func (x *RawVideoProcessResponse) Reset() {
	*x = RawVideoProcessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawVideoProcessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawVideoProcessResponse) ProtoMessage() {}

func (x *RawVideoProcessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_raw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawVideoProcessResponse.ProtoReflect.Descriptor instead.
func (*RawVideoProcessResponse) Descriptor() ([]byte, []int) {
	return file_raw_proto_rawDescGZIP(), []int{2}
}

func (x *RawVideoProcessResponse) GetRawVideoId() string {
	if x != nil {
		return x.RawVideoId
	}
	return ""
}

// key(user_id, id)
type CutVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// In the form <USER_ID>.<CREATE_TIMESTAMP_MS>.CUT_VIDEO.mp4
	CloudStorageFileName string `protobuf:"bytes,3,opt,name=cloud_storage_file_name,json=cloudStorageFileName,proto3" json:"cloud_storage_file_name,omitempty"`
	// When the video begins.
	CreateTimeMs int64  `protobuf:"varint,4,opt,name=create_time_ms,json=createTimeMs,proto3" json:"create_time_ms,omitempty"`
	RawVideoId   string `protobuf:"bytes,5,opt,name=raw_video_id,json=rawVideoId,proto3" json:"raw_video_id,omitempty"`
	DurationMs   int64  `protobuf:"varint,6,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
}

func (x *CutVideo) Reset() {
	*x = CutVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CutVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CutVideo) ProtoMessage() {}

func (x *CutVideo) ProtoReflect() protoreflect.Message {
	mi := &file_raw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CutVideo.ProtoReflect.Descriptor instead.
func (*CutVideo) Descriptor() ([]byte, []int) {
	return file_raw_proto_rawDescGZIP(), []int{3}
}

func (x *CutVideo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CutVideo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CutVideo) GetCloudStorageFileName() string {
	if x != nil {
		return x.CloudStorageFileName
	}
	return ""
}

func (x *CutVideo) GetCreateTimeMs() int64 {
	if x != nil {
		return x.CreateTimeMs
	}
	return 0
}

func (x *CutVideo) GetRawVideoId() string {
	if x != nil {
		return x.RawVideoId
	}
	return ""
}

func (x *CutVideo) GetDurationMs() int64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

// key(user_id, id)
type RawLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id          string    `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Location    *Location `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	TimestampMs int64     `protobuf:"varint,4,opt,name=timestamp_ms,json=timestampMs,proto3" json:"timestamp_ms,omitempty"`
	Source      *Source   `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *RawLocation) Reset() {
	*x = RawLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawLocation) ProtoMessage() {}

func (x *RawLocation) ProtoReflect() protoreflect.Message {
	mi := &file_raw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawLocation.ProtoReflect.Descriptor instead.
func (*RawLocation) Descriptor() ([]byte, []int) {
	return file_raw_proto_rawDescGZIP(), []int{4}
}

func (x *RawLocation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RawLocation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RawLocation) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *RawLocation) GetTimestampMs() int64 {
	if x != nil {
		return x.TimestampMs
	}
	return 0
}

func (x *RawLocation) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

// key(user_id, id)
type RawMotion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Id          string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Motion      *Motion `protobuf:"bytes,3,opt,name=motion,proto3" json:"motion,omitempty"`
	TimestampMs int64   `protobuf:"varint,4,opt,name=timestamp_ms,json=timestampMs,proto3" json:"timestamp_ms,omitempty"`
	Source      *Source `protobuf:"bytes,5,opt,name=source,proto3" json:"source,omitempty"`
	// The algorithms this rawMotion has gone through.
	AlgoTag []string `protobuf:"bytes,6,rep,name=algo_tag,json=algoTag,proto3" json:"algo_tag,omitempty"`
}

func (x *RawMotion) Reset() {
	*x = RawMotion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_raw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawMotion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawMotion) ProtoMessage() {}

func (x *RawMotion) ProtoReflect() protoreflect.Message {
	mi := &file_raw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawMotion.ProtoReflect.Descriptor instead.
func (*RawMotion) Descriptor() ([]byte, []int) {
	return file_raw_proto_rawDescGZIP(), []int{5}
}

func (x *RawMotion) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RawMotion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RawMotion) GetMotion() *Motion {
	if x != nil {
		return x.Motion
	}
	return nil
}

func (x *RawMotion) GetTimestampMs() int64 {
	if x != nil {
		return x.TimestampMs
	}
	return 0
}

func (x *RawMotion) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *RawMotion) GetAlgoTag() []string {
	if x != nil {
		return x.AlgoTag
	}
	return nil
}

var File_raw_proto protoreflect.FileDescriptor

var file_raw_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2,
	0x02, 0x0a, 0x08, 0x52, 0x61, 0x77, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x4d,
	0x73, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x75, 0x74, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x74, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x16, 0x52, 0x61, 0x77, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x17, 0x52, 0x61, 0x77, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x49, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x08, 0x43, 0x75, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x61, 0x77, 0x5f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x61,
	0x77, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x52, 0x61,
	0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73,
	0x12, 0x23, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x09, 0x52, 0x61, 0x77, 0x4d, 0x6f, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x06,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x4d, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x67,
	0x6f, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x67,
	0x6f, 0x54, 0x61, 0x67, 0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_raw_proto_rawDescOnce sync.Once
	file_raw_proto_rawDescData = file_raw_proto_rawDesc
)

func file_raw_proto_rawDescGZIP() []byte {
	file_raw_proto_rawDescOnce.Do(func() {
		file_raw_proto_rawDescData = protoimpl.X.CompressGZIP(file_raw_proto_rawDescData)
	})
	return file_raw_proto_rawDescData
}

var file_raw_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_raw_proto_goTypes = []interface{}{
	(*RawVideo)(nil),                // 0: api.RawVideo
	(*RawVideoProcessRequest)(nil),  // 1: api.RawVideoProcessRequest
	(*RawVideoProcessResponse)(nil), // 2: api.RawVideoProcessResponse
	(*CutVideo)(nil),                // 3: api.CutVideo
	(*RawLocation)(nil),             // 4: api.RawLocation
	(*RawMotion)(nil),               // 5: api.RawMotion
	(*Location)(nil),                // 6: api.Location
	(*Source)(nil),                  // 7: api.Source
	(*Motion)(nil),                  // 8: api.Motion
}
var file_raw_proto_depIdxs = []int32{
	6, // 0: api.RawLocation.location:type_name -> api.Location
	7, // 1: api.RawLocation.source:type_name -> api.Source
	8, // 2: api.RawMotion.motion:type_name -> api.Motion
	7, // 3: api.RawMotion.source:type_name -> api.Source
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_raw_proto_init() }
func file_raw_proto_init() {
	if File_raw_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_raw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawVideo); i {
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
		file_raw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawVideoProcessRequest); i {
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
		file_raw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawVideoProcessResponse); i {
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
		file_raw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CutVideo); i {
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
		file_raw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawLocation); i {
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
		file_raw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawMotion); i {
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
			RawDescriptor: file_raw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_raw_proto_goTypes,
		DependencyIndexes: file_raw_proto_depIdxs,
		MessageInfos:      file_raw_proto_msgTypes,
	}.Build()
	File_raw_proto = out.File
	file_raw_proto_rawDesc = nil
	file_raw_proto_goTypes = nil
	file_raw_proto_depIdxs = nil
}

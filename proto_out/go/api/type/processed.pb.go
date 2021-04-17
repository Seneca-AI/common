// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: processed.proto

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

type LaneChangesForVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumFrames int64 `protobuf:"varint,1,opt,name=num_frames,json=numFrames,proto3" json:"num_frames,omitempty"`
	// The length(lane_change_for_frame) must equal num_frames.
	LaneChangeForFrame []*LaneChangesForVideo_LaneChangeForFrame `protobuf:"bytes,2,rep,name=lane_change_for_frame,json=laneChangeForFrame,proto3" json:"lane_change_for_frame,omitempty"`
}

func (x *LaneChangesForVideo) Reset() {
	*x = LaneChangesForVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaneChangesForVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneChangesForVideo) ProtoMessage() {}

func (x *LaneChangesForVideo) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneChangesForVideo.ProtoReflect.Descriptor instead.
func (*LaneChangesForVideo) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{0}
}

func (x *LaneChangesForVideo) GetNumFrames() int64 {
	if x != nil {
		return x.NumFrames
	}
	return 0
}

func (x *LaneChangesForVideo) GetLaneChangeForFrame() []*LaneChangesForVideo_LaneChangeForFrame {
	if x != nil {
		return x.LaneChangeForFrame
	}
	return nil
}

type LaneChangesForVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId             string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SimpleStorageVideoUrl string `protobuf:"bytes,2,opt,name=simple_storage_video_url,json=simpleStorageVideoUrl,proto3" json:"simple_storage_video_url,omitempty"`
}

func (x *LaneChangesForVideoRequest) Reset() {
	*x = LaneChangesForVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaneChangesForVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneChangesForVideoRequest) ProtoMessage() {}

func (x *LaneChangesForVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneChangesForVideoRequest.ProtoReflect.Descriptor instead.
func (*LaneChangesForVideoRequest) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{1}
}

func (x *LaneChangesForVideoRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *LaneChangesForVideoRequest) GetSimpleStorageVideoUrl() string {
	if x != nil {
		return x.SimpleStorageVideoUrl
	}
	return ""
}

type LaneChangesForVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request this response is for.
	RequestId           string               `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	LaneChangesForVideo *LaneChangesForVideo `protobuf:"bytes,2,opt,name=lane_changes_for_video,json=laneChangesForVideo,proto3" json:"lane_changes_for_video,omitempty"`
}

func (x *LaneChangesForVideoResponse) Reset() {
	*x = LaneChangesForVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaneChangesForVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneChangesForVideoResponse) ProtoMessage() {}

func (x *LaneChangesForVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneChangesForVideoResponse.ProtoReflect.Descriptor instead.
func (*LaneChangesForVideoResponse) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{2}
}

func (x *LaneChangesForVideoResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *LaneChangesForVideoResponse) GetLaneChangesForVideo() *LaneChangesForVideo {
	if x != nil {
		return x.LaneChangesForVideo
	}
	return nil
}

type FollowingDistanceForVideo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumFrames int64 `protobuf:"varint,1,opt,name=num_frames,json=numFrames,proto3" json:"num_frames,omitempty"`
	// The length(lane_change_for_frame) must equal num_frames.
	FollowingDistanceForFrame []*FollowingDistanceForVideo_FollowingDistanceForFrame `protobuf:"bytes,2,rep,name=following_distance_for_frame,json=followingDistanceForFrame,proto3" json:"following_distance_for_frame,omitempty"`
}

func (x *FollowingDistanceForVideo) Reset() {
	*x = FollowingDistanceForVideo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingDistanceForVideo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingDistanceForVideo) ProtoMessage() {}

func (x *FollowingDistanceForVideo) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingDistanceForVideo.ProtoReflect.Descriptor instead.
func (*FollowingDistanceForVideo) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{3}
}

func (x *FollowingDistanceForVideo) GetNumFrames() int64 {
	if x != nil {
		return x.NumFrames
	}
	return 0
}

func (x *FollowingDistanceForVideo) GetFollowingDistanceForFrame() []*FollowingDistanceForVideo_FollowingDistanceForFrame {
	if x != nil {
		return x.FollowingDistanceForFrame
	}
	return nil
}

type FollowingDistanceForVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId             string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	SimpleStorageVideoUrl string `protobuf:"bytes,2,opt,name=simple_storage_video_url,json=simpleStorageVideoUrl,proto3" json:"simple_storage_video_url,omitempty"`
}

func (x *FollowingDistanceForVideoRequest) Reset() {
	*x = FollowingDistanceForVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingDistanceForVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingDistanceForVideoRequest) ProtoMessage() {}

func (x *FollowingDistanceForVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingDistanceForVideoRequest.ProtoReflect.Descriptor instead.
func (*FollowingDistanceForVideoRequest) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{4}
}

func (x *FollowingDistanceForVideoRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *FollowingDistanceForVideoRequest) GetSimpleStorageVideoUrl() string {
	if x != nil {
		return x.SimpleStorageVideoUrl
	}
	return ""
}

type FollowingDistanceForVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request this response is for.
	RequestId           string                     `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	LaneChangesForVideo *FollowingDistanceForVideo `protobuf:"bytes,2,opt,name=lane_changes_for_video,json=laneChangesForVideo,proto3" json:"lane_changes_for_video,omitempty"`
}

func (x *FollowingDistanceForVideoResponse) Reset() {
	*x = FollowingDistanceForVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingDistanceForVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingDistanceForVideoResponse) ProtoMessage() {}

func (x *FollowingDistanceForVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingDistanceForVideoResponse.ProtoReflect.Descriptor instead.
func (*FollowingDistanceForVideoResponse) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{5}
}

func (x *FollowingDistanceForVideoResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *FollowingDistanceForVideoResponse) GetLaneChangesForVideo() *FollowingDistanceForVideo {
	if x != nil {
		return x.LaneChangesForVideo
	}
	return nil
}

type LaneChangesForVideo_LaneChangeForFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameIndex int64 `protobuf:"varint,1,opt,name=frame_index,json=frameIndex,proto3" json:"frame_index,omitempty"`
	// This will likely be deprecated as we learn to have a better way
	// of measuring lane changes.
	LaneChange bool `protobuf:"varint,101,opt,name=lane_change,json=laneChange,proto3" json:"lane_change,omitempty"`
}

func (x *LaneChangesForVideo_LaneChangeForFrame) Reset() {
	*x = LaneChangesForVideo_LaneChangeForFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaneChangesForVideo_LaneChangeForFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneChangesForVideo_LaneChangeForFrame) ProtoMessage() {}

func (x *LaneChangesForVideo_LaneChangeForFrame) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneChangesForVideo_LaneChangeForFrame.ProtoReflect.Descriptor instead.
func (*LaneChangesForVideo_LaneChangeForFrame) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LaneChangesForVideo_LaneChangeForFrame) GetFrameIndex() int64 {
	if x != nil {
		return x.FrameIndex
	}
	return 0
}

func (x *LaneChangesForVideo_LaneChangeForFrame) GetLaneChange() bool {
	if x != nil {
		return x.LaneChange
	}
	return false
}

type FollowingDistanceForVideo_FollowingDistanceForFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameIndex int64 `protobuf:"varint,1,opt,name=frame_index,json=frameIndex,proto3" json:"frame_index,omitempty"`
	// This will likely be deprecated as we learn to have a better way
	// of measuring following distance.
	IsTooClose bool `protobuf:"varint,101,opt,name=is_too_close,json=isTooClose,proto3" json:"is_too_close,omitempty"`
}

func (x *FollowingDistanceForVideo_FollowingDistanceForFrame) Reset() {
	*x = FollowingDistanceForVideo_FollowingDistanceForFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_processed_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingDistanceForVideo_FollowingDistanceForFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingDistanceForVideo_FollowingDistanceForFrame) ProtoMessage() {}

func (x *FollowingDistanceForVideo_FollowingDistanceForFrame) ProtoReflect() protoreflect.Message {
	mi := &file_processed_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingDistanceForVideo_FollowingDistanceForFrame.ProtoReflect.Descriptor instead.
func (*FollowingDistanceForVideo_FollowingDistanceForFrame) Descriptor() ([]byte, []int) {
	return file_processed_proto_rawDescGZIP(), []int{3, 0}
}

func (x *FollowingDistanceForVideo_FollowingDistanceForFrame) GetFrameIndex() int64 {
	if x != nil {
		return x.FrameIndex
	}
	return 0
}

func (x *FollowingDistanceForVideo_FollowingDistanceForFrame) GetIsTooClose() bool {
	if x != nil {
		return x.IsTooClose
	}
	return false
}

var File_processed_proto protoreflect.FileDescriptor

var file_processed_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x13, 0x4c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d,
	0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e,
	0x75, 0x6d, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x15, 0x6c, 0x61, 0x6e, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x4c, 0x61, 0x6e, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x4c, 0x61,
	0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x52, 0x12, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x1a, 0x56, 0x0a, 0x12, 0x4c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x46, 0x6f, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x72,
	0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x61, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x1a,
	0x4c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x72, 0x6c, 0x22, 0x87, 0x01, 0x0a, 0x1b, 0x4c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x49, 0x0a, 0x16, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x4c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x46,
	0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x13, 0x6c, 0x61, 0x6e, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x91, 0x02, 0x0a,
	0x19, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75,
	0x6d, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6e, 0x75, 0x6d, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x75, 0x0a, 0x1c, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x66, 0x6f, 0x72, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x19, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x1a, 0x5e, 0x0a, 0x19, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x74, 0x6f, 0x6f, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x18, 0x65,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x54, 0x6f, 0x6f, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x22, 0x7a, 0x0a, 0x20, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x72, 0x6c, 0x22, 0x93, 0x01, 0x0a,
	0x21, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x4f, 0x0a, 0x16, 0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x73, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x13, 0x6c,
	0x61, 0x6e, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_processed_proto_rawDescOnce sync.Once
	file_processed_proto_rawDescData = file_processed_proto_rawDesc
)

func file_processed_proto_rawDescGZIP() []byte {
	file_processed_proto_rawDescOnce.Do(func() {
		file_processed_proto_rawDescData = protoimpl.X.CompressGZIP(file_processed_proto_rawDescData)
	})
	return file_processed_proto_rawDescData
}

var file_processed_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_processed_proto_goTypes = []interface{}{
	(*LaneChangesForVideo)(nil),                                 // 0: LaneChangesForVideo
	(*LaneChangesForVideoRequest)(nil),                          // 1: LaneChangesForVideoRequest
	(*LaneChangesForVideoResponse)(nil),                         // 2: LaneChangesForVideoResponse
	(*FollowingDistanceForVideo)(nil),                           // 3: FollowingDistanceForVideo
	(*FollowingDistanceForVideoRequest)(nil),                    // 4: FollowingDistanceForVideoRequest
	(*FollowingDistanceForVideoResponse)(nil),                   // 5: FollowingDistanceForVideoResponse
	(*LaneChangesForVideo_LaneChangeForFrame)(nil),              // 6: LaneChangesForVideo.LaneChangeForFrame
	(*FollowingDistanceForVideo_FollowingDistanceForFrame)(nil), // 7: FollowingDistanceForVideo.FollowingDistanceForFrame
}
var file_processed_proto_depIdxs = []int32{
	6, // 0: LaneChangesForVideo.lane_change_for_frame:type_name -> LaneChangesForVideo.LaneChangeForFrame
	0, // 1: LaneChangesForVideoResponse.lane_changes_for_video:type_name -> LaneChangesForVideo
	7, // 2: FollowingDistanceForVideo.following_distance_for_frame:type_name -> FollowingDistanceForVideo.FollowingDistanceForFrame
	3, // 3: FollowingDistanceForVideoResponse.lane_changes_for_video:type_name -> FollowingDistanceForVideo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_processed_proto_init() }
func file_processed_proto_init() {
	if File_processed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_processed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaneChangesForVideo); i {
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
		file_processed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaneChangesForVideoRequest); i {
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
		file_processed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaneChangesForVideoResponse); i {
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
		file_processed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingDistanceForVideo); i {
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
		file_processed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingDistanceForVideoRequest); i {
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
		file_processed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingDistanceForVideoResponse); i {
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
		file_processed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaneChangesForVideo_LaneChangeForFrame); i {
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
		file_processed_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingDistanceForVideo_FollowingDistanceForFrame); i {
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
			RawDescriptor: file_processed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_processed_proto_goTypes,
		DependencyIndexes: file_processed_proto_depIdxs,
		MessageInfos:      file_processed_proto_msgTypes,
	}.Build()
	File_processed_proto = out.File
	file_processed_proto_rawDesc = nil
	file_processed_proto_goTypes = nil
	file_processed_proto_depIdxs = nil
}

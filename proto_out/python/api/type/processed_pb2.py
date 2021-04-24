# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: processed.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='processed.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\010api/type',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0fprocessed.proto\"\xb1\x01\n\x13LaneChangesForVideo\x12\x12\n\nnum_frames\x18\x01 \x01(\x03\x12\x46\n\x15lane_change_for_frame\x18\x02 \x03(\x0b\x32\'.LaneChangesForVideo.LaneChangeForFrame\x1a>\n\x12LaneChangeForFrame\x12\x13\n\x0b\x66rame_index\x18\x01 \x01(\x03\x12\x13\n\x0blane_change\x18\x65 \x01(\x08\"R\n\x1aLaneChangesForVideoRequest\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12 \n\x18simple_storage_video_url\x18\x02 \x01(\t\"g\n\x1bLaneChangesForVideoResponse\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12\x34\n\x16lane_changes_for_video\x18\x02 \x01(\x0b\x32\x14.LaneChangesForVideo\"\xd3\x01\n\x19\x46ollowingDistanceForVideo\x12\x12\n\nnum_frames\x18\x01 \x01(\x03\x12Z\n\x1c\x66ollowing_distance_for_frame\x18\x02 \x03(\x0b\x32\x34.FollowingDistanceForVideo.FollowingDistanceForFrame\x1a\x46\n\x19\x46ollowingDistanceForFrame\x12\x13\n\x0b\x66rame_index\x18\x01 \x01(\x03\x12\x14\n\x0cis_too_close\x18\x65 \x01(\x08\"X\n FollowingDistanceForVideoRequest\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12 \n\x18simple_storage_video_url\x18\x02 \x01(\t\"y\n!FollowingDistanceForVideoResponse\x12\x12\n\nrequest_id\x18\x01 \x01(\t\x12@\n\x1c\x66ollowing_distance_for_video\x18\x02 \x01(\x0b\x32\x1a.FollowingDistanceForVideoB\nZ\x08\x61pi/typeb\x06proto3'
)




_LANECHANGESFORVIDEO_LANECHANGEFORFRAME = _descriptor.Descriptor(
  name='LaneChangeForFrame',
  full_name='LaneChangesForVideo.LaneChangeForFrame',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='frame_index', full_name='LaneChangesForVideo.LaneChangeForFrame.frame_index', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='lane_change', full_name='LaneChangesForVideo.LaneChangeForFrame.lane_change', index=1,
      number=101, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=135,
  serialized_end=197,
)

_LANECHANGESFORVIDEO = _descriptor.Descriptor(
  name='LaneChangesForVideo',
  full_name='LaneChangesForVideo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='num_frames', full_name='LaneChangesForVideo.num_frames', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='lane_change_for_frame', full_name='LaneChangesForVideo.lane_change_for_frame', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_LANECHANGESFORVIDEO_LANECHANGEFORFRAME, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=20,
  serialized_end=197,
)


_LANECHANGESFORVIDEOREQUEST = _descriptor.Descriptor(
  name='LaneChangesForVideoRequest',
  full_name='LaneChangesForVideoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='LaneChangesForVideoRequest.request_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='simple_storage_video_url', full_name='LaneChangesForVideoRequest.simple_storage_video_url', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=199,
  serialized_end=281,
)


_LANECHANGESFORVIDEORESPONSE = _descriptor.Descriptor(
  name='LaneChangesForVideoResponse',
  full_name='LaneChangesForVideoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='LaneChangesForVideoResponse.request_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='lane_changes_for_video', full_name='LaneChangesForVideoResponse.lane_changes_for_video', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=283,
  serialized_end=386,
)


_FOLLOWINGDISTANCEFORVIDEO_FOLLOWINGDISTANCEFORFRAME = _descriptor.Descriptor(
  name='FollowingDistanceForFrame',
  full_name='FollowingDistanceForVideo.FollowingDistanceForFrame',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='frame_index', full_name='FollowingDistanceForVideo.FollowingDistanceForFrame.frame_index', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='is_too_close', full_name='FollowingDistanceForVideo.FollowingDistanceForFrame.is_too_close', index=1,
      number=101, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=530,
  serialized_end=600,
)

_FOLLOWINGDISTANCEFORVIDEO = _descriptor.Descriptor(
  name='FollowingDistanceForVideo',
  full_name='FollowingDistanceForVideo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='num_frames', full_name='FollowingDistanceForVideo.num_frames', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='following_distance_for_frame', full_name='FollowingDistanceForVideo.following_distance_for_frame', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_FOLLOWINGDISTANCEFORVIDEO_FOLLOWINGDISTANCEFORFRAME, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=389,
  serialized_end=600,
)


_FOLLOWINGDISTANCEFORVIDEOREQUEST = _descriptor.Descriptor(
  name='FollowingDistanceForVideoRequest',
  full_name='FollowingDistanceForVideoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='FollowingDistanceForVideoRequest.request_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='simple_storage_video_url', full_name='FollowingDistanceForVideoRequest.simple_storage_video_url', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=602,
  serialized_end=690,
)


_FOLLOWINGDISTANCEFORVIDEORESPONSE = _descriptor.Descriptor(
  name='FollowingDistanceForVideoResponse',
  full_name='FollowingDistanceForVideoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='FollowingDistanceForVideoResponse.request_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='following_distance_for_video', full_name='FollowingDistanceForVideoResponse.following_distance_for_video', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=692,
  serialized_end=813,
)

_LANECHANGESFORVIDEO_LANECHANGEFORFRAME.containing_type = _LANECHANGESFORVIDEO
_LANECHANGESFORVIDEO.fields_by_name['lane_change_for_frame'].message_type = _LANECHANGESFORVIDEO_LANECHANGEFORFRAME
_LANECHANGESFORVIDEORESPONSE.fields_by_name['lane_changes_for_video'].message_type = _LANECHANGESFORVIDEO
_FOLLOWINGDISTANCEFORVIDEO_FOLLOWINGDISTANCEFORFRAME.containing_type = _FOLLOWINGDISTANCEFORVIDEO
_FOLLOWINGDISTANCEFORVIDEO.fields_by_name['following_distance_for_frame'].message_type = _FOLLOWINGDISTANCEFORVIDEO_FOLLOWINGDISTANCEFORFRAME
_FOLLOWINGDISTANCEFORVIDEORESPONSE.fields_by_name['following_distance_for_video'].message_type = _FOLLOWINGDISTANCEFORVIDEO
DESCRIPTOR.message_types_by_name['LaneChangesForVideo'] = _LANECHANGESFORVIDEO
DESCRIPTOR.message_types_by_name['LaneChangesForVideoRequest'] = _LANECHANGESFORVIDEOREQUEST
DESCRIPTOR.message_types_by_name['LaneChangesForVideoResponse'] = _LANECHANGESFORVIDEORESPONSE
DESCRIPTOR.message_types_by_name['FollowingDistanceForVideo'] = _FOLLOWINGDISTANCEFORVIDEO
DESCRIPTOR.message_types_by_name['FollowingDistanceForVideoRequest'] = _FOLLOWINGDISTANCEFORVIDEOREQUEST
DESCRIPTOR.message_types_by_name['FollowingDistanceForVideoResponse'] = _FOLLOWINGDISTANCEFORVIDEORESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

LaneChangesForVideo = _reflection.GeneratedProtocolMessageType('LaneChangesForVideo', (_message.Message,), {

  'LaneChangeForFrame' : _reflection.GeneratedProtocolMessageType('LaneChangeForFrame', (_message.Message,), {
    'DESCRIPTOR' : _LANECHANGESFORVIDEO_LANECHANGEFORFRAME,
    '__module__' : 'processed_pb2'
    # @@protoc_insertion_point(class_scope:LaneChangesForVideo.LaneChangeForFrame)
    })
  ,
  'DESCRIPTOR' : _LANECHANGESFORVIDEO,
  '__module__' : 'processed_pb2'
  # @@protoc_insertion_point(class_scope:LaneChangesForVideo)
  })
_sym_db.RegisterMessage(LaneChangesForVideo)
_sym_db.RegisterMessage(LaneChangesForVideo.LaneChangeForFrame)

LaneChangesForVideoRequest = _reflection.GeneratedProtocolMessageType('LaneChangesForVideoRequest', (_message.Message,), {
  'DESCRIPTOR' : _LANECHANGESFORVIDEOREQUEST,
  '__module__' : 'processed_pb2'
  # @@protoc_insertion_point(class_scope:LaneChangesForVideoRequest)
  })
_sym_db.RegisterMessage(LaneChangesForVideoRequest)

LaneChangesForVideoResponse = _reflection.GeneratedProtocolMessageType('LaneChangesForVideoResponse', (_message.Message,), {
  'DESCRIPTOR' : _LANECHANGESFORVIDEORESPONSE,
  '__module__' : 'processed_pb2'
  # @@protoc_insertion_point(class_scope:LaneChangesForVideoResponse)
  })
_sym_db.RegisterMessage(LaneChangesForVideoResponse)

FollowingDistanceForVideo = _reflection.GeneratedProtocolMessageType('FollowingDistanceForVideo', (_message.Message,), {

  'FollowingDistanceForFrame' : _reflection.GeneratedProtocolMessageType('FollowingDistanceForFrame', (_message.Message,), {
    'DESCRIPTOR' : _FOLLOWINGDISTANCEFORVIDEO_FOLLOWINGDISTANCEFORFRAME,
    '__module__' : 'processed_pb2'
    # @@protoc_insertion_point(class_scope:FollowingDistanceForVideo.FollowingDistanceForFrame)
    })
  ,
  'DESCRIPTOR' : _FOLLOWINGDISTANCEFORVIDEO,
  '__module__' : 'processed_pb2'
  # @@protoc_insertion_point(class_scope:FollowingDistanceForVideo)
  })
_sym_db.RegisterMessage(FollowingDistanceForVideo)
_sym_db.RegisterMessage(FollowingDistanceForVideo.FollowingDistanceForFrame)

FollowingDistanceForVideoRequest = _reflection.GeneratedProtocolMessageType('FollowingDistanceForVideoRequest', (_message.Message,), {
  'DESCRIPTOR' : _FOLLOWINGDISTANCEFORVIDEOREQUEST,
  '__module__' : 'processed_pb2'
  # @@protoc_insertion_point(class_scope:FollowingDistanceForVideoRequest)
  })
_sym_db.RegisterMessage(FollowingDistanceForVideoRequest)

FollowingDistanceForVideoResponse = _reflection.GeneratedProtocolMessageType('FollowingDistanceForVideoResponse', (_message.Message,), {
  'DESCRIPTOR' : _FOLLOWINGDISTANCEFORVIDEORESPONSE,
  '__module__' : 'processed_pb2'
  # @@protoc_insertion_point(class_scope:FollowingDistanceForVideoResponse)
  })
_sym_db.RegisterMessage(FollowingDistanceForVideoResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

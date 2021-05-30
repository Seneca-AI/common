# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: raw.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import common_pb2 as common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='raw.proto',
  package='api',
  syntax='proto3',
  serialized_options=b'Z\010api/type',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\traw.proto\x12\x03\x61pi\x1a\x0c\x63ommon.proto\"\xf0\x01\n\x08RawVideo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x1f\n\x17\x63loud_storage_file_name\x18\x03 \x01(\t\x12\x16\n\x0e\x63reate_time_ms\x18\x04 \x01(\x03\x12\x14\n\x0c\x63ut_video_id\x18\x05 \x03(\t\x12\x13\n\x0b\x64uration_ms\x18\x06 \x01(\x03\x12\x1a\n\x12original_file_name\x18\x07 \x01(\t\x12\x16\n\x0e\x66\x61ilure_reason\x18\x08 \x01(\t\x12\x18\n\x10original_file_id\x18\t \x01(\t\x12\x15\n\ralgos_version\x18\n \x01(\x01\"f\n\x16RawVideoProcessRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x12\n\nvideo_name\x18\x02 \x01(\t\x12\x12\n\nlocal_path\x18\x03 \x01(\t\x12\x13\n\x0bvideo_bytes\x18\x04 \x01(\x0c\"/\n\x17RawVideoProcessResponse\x12\x14\n\x0craw_video_id\x18\x01 \x01(\t\"\x8b\x01\n\x08\x43utVideo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x1f\n\x17\x63loud_storage_file_name\x18\x03 \x01(\t\x12\x16\n\x0e\x63reate_time_ms\x18\x04 \x01(\x03\x12\x14\n\x0craw_video_id\x18\x05 \x01(\t\x12\x13\n\x0b\x64uration_ms\x18\x06 \x01(\x03\"\xa7\x01\n\x0bRawLocation\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x1f\n\x08location\x18\x03 \x01(\x0b\x32\r.api.Location\x12\x14\n\x0ctimestamp_ms\x18\x04 \x01(\x03\x12\x1b\n\x06source\x18\x05 \x01(\x0b\x32\x0b.api.Source\x12\x10\n\x08\x61lgo_tag\x18\x06 \x03(\t\x12\x15\n\ralgos_version\x18\x07 \x01(\x01\"\xa1\x01\n\tRawMotion\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x1b\n\x06motion\x18\x03 \x01(\x0b\x32\x0b.api.Motion\x12\x14\n\x0ctimestamp_ms\x18\x04 \x01(\x03\x12\x1b\n\x06source\x18\x05 \x01(\x0b\x32\x0b.api.Source\x12\x10\n\x08\x61lgo_tag\x18\x06 \x03(\t\x12\x15\n\ralgos_version\x18\x07 \x01(\x01\x42\nZ\x08\x61pi/typeb\x06proto3'
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_RAWVIDEO = _descriptor.Descriptor(
  name='RawVideo',
  full_name='api.RawVideo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.RawVideo.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='api.RawVideo.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cloud_storage_file_name', full_name='api.RawVideo.cloud_storage_file_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='create_time_ms', full_name='api.RawVideo.create_time_ms', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cut_video_id', full_name='api.RawVideo.cut_video_id', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='duration_ms', full_name='api.RawVideo.duration_ms', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='original_file_name', full_name='api.RawVideo.original_file_name', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='failure_reason', full_name='api.RawVideo.failure_reason', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='original_file_id', full_name='api.RawVideo.original_file_id', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='algos_version', full_name='api.RawVideo.algos_version', index=9,
      number=10, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=33,
  serialized_end=273,
)


_RAWVIDEOPROCESSREQUEST = _descriptor.Descriptor(
  name='RawVideoProcessRequest',
  full_name='api.RawVideoProcessRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.RawVideoProcessRequest.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='video_name', full_name='api.RawVideoProcessRequest.video_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='local_path', full_name='api.RawVideoProcessRequest.local_path', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='video_bytes', full_name='api.RawVideoProcessRequest.video_bytes', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
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
  serialized_start=275,
  serialized_end=377,
)


_RAWVIDEOPROCESSRESPONSE = _descriptor.Descriptor(
  name='RawVideoProcessResponse',
  full_name='api.RawVideoProcessResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='raw_video_id', full_name='api.RawVideoProcessResponse.raw_video_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=379,
  serialized_end=426,
)


_CUTVIDEO = _descriptor.Descriptor(
  name='CutVideo',
  full_name='api.CutVideo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.CutVideo.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='api.CutVideo.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cloud_storage_file_name', full_name='api.CutVideo.cloud_storage_file_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='create_time_ms', full_name='api.CutVideo.create_time_ms', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='raw_video_id', full_name='api.CutVideo.raw_video_id', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='duration_ms', full_name='api.CutVideo.duration_ms', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=429,
  serialized_end=568,
)


_RAWLOCATION = _descriptor.Descriptor(
  name='RawLocation',
  full_name='api.RawLocation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.RawLocation.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='api.RawLocation.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='location', full_name='api.RawLocation.location', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp_ms', full_name='api.RawLocation.timestamp_ms', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='source', full_name='api.RawLocation.source', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='algo_tag', full_name='api.RawLocation.algo_tag', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='algos_version', full_name='api.RawLocation.algos_version', index=6,
      number=7, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=571,
  serialized_end=738,
)


_RAWMOTION = _descriptor.Descriptor(
  name='RawMotion',
  full_name='api.RawMotion',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.RawMotion.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='api.RawMotion.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='motion', full_name='api.RawMotion.motion', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp_ms', full_name='api.RawMotion.timestamp_ms', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='source', full_name='api.RawMotion.source', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='algo_tag', full_name='api.RawMotion.algo_tag', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='algos_version', full_name='api.RawMotion.algos_version', index=6,
      number=7, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
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
  serialized_start=741,
  serialized_end=902,
)

_RAWLOCATION.fields_by_name['location'].message_type = common__pb2._LOCATION
_RAWLOCATION.fields_by_name['source'].message_type = common__pb2._SOURCE
_RAWMOTION.fields_by_name['motion'].message_type = common__pb2._MOTION
_RAWMOTION.fields_by_name['source'].message_type = common__pb2._SOURCE
DESCRIPTOR.message_types_by_name['RawVideo'] = _RAWVIDEO
DESCRIPTOR.message_types_by_name['RawVideoProcessRequest'] = _RAWVIDEOPROCESSREQUEST
DESCRIPTOR.message_types_by_name['RawVideoProcessResponse'] = _RAWVIDEOPROCESSRESPONSE
DESCRIPTOR.message_types_by_name['CutVideo'] = _CUTVIDEO
DESCRIPTOR.message_types_by_name['RawLocation'] = _RAWLOCATION
DESCRIPTOR.message_types_by_name['RawMotion'] = _RAWMOTION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RawVideo = _reflection.GeneratedProtocolMessageType('RawVideo', (_message.Message,), {
  'DESCRIPTOR' : _RAWVIDEO,
  '__module__' : 'raw_pb2'
  # @@protoc_insertion_point(class_scope:api.RawVideo)
  })
_sym_db.RegisterMessage(RawVideo)

RawVideoProcessRequest = _reflection.GeneratedProtocolMessageType('RawVideoProcessRequest', (_message.Message,), {
  'DESCRIPTOR' : _RAWVIDEOPROCESSREQUEST,
  '__module__' : 'raw_pb2'
  # @@protoc_insertion_point(class_scope:api.RawVideoProcessRequest)
  })
_sym_db.RegisterMessage(RawVideoProcessRequest)

RawVideoProcessResponse = _reflection.GeneratedProtocolMessageType('RawVideoProcessResponse', (_message.Message,), {
  'DESCRIPTOR' : _RAWVIDEOPROCESSRESPONSE,
  '__module__' : 'raw_pb2'
  # @@protoc_insertion_point(class_scope:api.RawVideoProcessResponse)
  })
_sym_db.RegisterMessage(RawVideoProcessResponse)

CutVideo = _reflection.GeneratedProtocolMessageType('CutVideo', (_message.Message,), {
  'DESCRIPTOR' : _CUTVIDEO,
  '__module__' : 'raw_pb2'
  # @@protoc_insertion_point(class_scope:api.CutVideo)
  })
_sym_db.RegisterMessage(CutVideo)

RawLocation = _reflection.GeneratedProtocolMessageType('RawLocation', (_message.Message,), {
  'DESCRIPTOR' : _RAWLOCATION,
  '__module__' : 'raw_pb2'
  # @@protoc_insertion_point(class_scope:api.RawLocation)
  })
_sym_db.RegisterMessage(RawLocation)

RawMotion = _reflection.GeneratedProtocolMessageType('RawMotion', (_message.Message,), {
  'DESCRIPTOR' : _RAWMOTION,
  '__module__' : 'raw_pb2'
  # @@protoc_insertion_point(class_scope:api.RawMotion)
  })
_sym_db.RegisterMessage(RawMotion)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

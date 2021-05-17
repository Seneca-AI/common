# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='common.proto',
  package='api',
  syntax='proto3',
  serialized_options=b'Z\010api/type',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0c\x63ommon.proto\x12\x03\x61pi\"\'\n\x06Header\x12\x0c\n\x04\x63ode\x18\x01 \x01(\x03\x12\x0f\n\x07message\x18\x02 \x01(\t\"8\n\nTimePeriod\x12\x15\n\rstart_time_ms\x18\x01 \x01(\x03\x12\x13\n\x0b\x65nd_time_ms\x18\x02 \x01(\x03\"\xb1\x01\n\x08Latitude\x12\x0f\n\x07\x64\x65grees\x18\x01 \x01(\x01\x12\x16\n\x0e\x64\x65gree_minutes\x18\x02 \x01(\x01\x12\x16\n\x0e\x64\x65gree_seconds\x18\x03 \x01(\x01\x12\x31\n\rlat_direction\x18\x04 \x01(\x0e\x32\x1a.api.Latitude.LatDirection\"1\n\x0cLatDirection\x12\x0b\n\x07UNKNOWN\x10\x00\x12\t\n\x05NORTH\x10\x01\x12\t\n\x05SOUTH\x10\x02\"\xb4\x01\n\tLongitude\x12\x0f\n\x07\x64\x65grees\x18\x01 \x01(\x01\x12\x16\n\x0e\x64\x65gree_minutes\x18\x02 \x01(\x01\x12\x16\n\x0e\x64\x65gree_seconds\x18\x03 \x01(\x01\x12\x34\n\x0elong_direction\x18\x04 \x01(\x0e\x32\x1c.api.Longitude.LongDirection\"0\n\rLongDirection\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x08\n\x04\x45\x41ST\x10\x01\x12\x08\n\x04WEST\x10\x02\"D\n\x08Location\x12\x1a\n\x03lat\x18\x01 \x01(\x0b\x32\r.api.Latitude\x12\x1c\n\x04long\x18\x02 \x01(\x0b\x32\x0e.api.Longitude\":\n\x06Motion\x12\x14\n\x0cvelocity_mph\x18\x01 \x01(\x01\x12\x1a\n\x12\x61\x63\x63\x65leration_mph_s\x18\x02 \x01(\x01\"\x88\x01\n\x06Source\x12\x11\n\tsource_id\x18\x01 \x01(\t\x12%\n\x0bsource_type\x18\x02 \x01(\x0e\x32\x10.api.Source.Type\"D\n\x04Type\x12\x0b\n\x07UNKNOWN\x10\x00\x12\r\n\tRAW_VIDEO\x10\x01\x12\x10\n\x0cRAW_LOCATION\x10\x02\x12\x0e\n\nRAW_MOTION\x10\x03*\x80\x01\n\rConditionType\x12\x17\n\x13NONE_CONDITION_TYPE\x10\x00\x12\t\n\x05NIGHT\x10\x01\x12\x08\n\x04RAIN\x10\x02\x12\x08\n\x04SNOW\x10\x03\x12\x07\n\x03ICE\x10\x04\x12\t\n\x05URBAN\x10\x05\x12\x0b\n\x07HIGHWAY\x10\x06\x12\t\n\x05RURAL\x10\x07\x12\x0b\n\x07TRAFFIC\x10\x08*b\n\tEventType\x12\x16\n\x12UNKNOWN_EVENT_TYPE\x10\x00\x12\x0f\n\x0bLANE_CHANGE\x10\x01\x12\x15\n\x11\x46\x41ST_ACCELERATION\x10\x02\x12\x15\n\x11\x46\x41ST_DECELERATION\x10\x03\x42\nZ\x08\x61pi/typeb\x06proto3'
)

_CONDITIONTYPE = _descriptor.EnumDescriptor(
  name='ConditionType',
  full_name='api.ConditionType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE_CONDITION_TYPE', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NIGHT', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RAIN', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SNOW', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ICE', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='URBAN', index=5, number=5,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='HIGHWAY', index=6, number=6,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RURAL', index=7, number=7,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='TRAFFIC', index=8, number=8,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=753,
  serialized_end=881,
)
_sym_db.RegisterEnumDescriptor(_CONDITIONTYPE)

ConditionType = enum_type_wrapper.EnumTypeWrapper(_CONDITIONTYPE)
_EVENTTYPE = _descriptor.EnumDescriptor(
  name='EventType',
  full_name='api.EventType',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN_EVENT_TYPE', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='LANE_CHANGE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='FAST_ACCELERATION', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='FAST_DECELERATION', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=883,
  serialized_end=981,
)
_sym_db.RegisterEnumDescriptor(_EVENTTYPE)

EventType = enum_type_wrapper.EnumTypeWrapper(_EVENTTYPE)
NONE_CONDITION_TYPE = 0
NIGHT = 1
RAIN = 2
SNOW = 3
ICE = 4
URBAN = 5
HIGHWAY = 6
RURAL = 7
TRAFFIC = 8
UNKNOWN_EVENT_TYPE = 0
LANE_CHANGE = 1
FAST_ACCELERATION = 2
FAST_DECELERATION = 3


_LATITUDE_LATDIRECTION = _descriptor.EnumDescriptor(
  name='LatDirection',
  full_name='api.Latitude.LatDirection',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NORTH', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='SOUTH', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=249,
  serialized_end=298,
)
_sym_db.RegisterEnumDescriptor(_LATITUDE_LATDIRECTION)

_LONGITUDE_LONGDIRECTION = _descriptor.EnumDescriptor(
  name='LongDirection',
  full_name='api.Longitude.LongDirection',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='EAST', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='WEST', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=433,
  serialized_end=481,
)
_sym_db.RegisterEnumDescriptor(_LONGITUDE_LONGDIRECTION)

_SOURCE_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='api.Source.Type',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RAW_VIDEO', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RAW_LOCATION', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='RAW_MOTION', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=682,
  serialized_end=750,
)
_sym_db.RegisterEnumDescriptor(_SOURCE_TYPE)


_HEADER = _descriptor.Descriptor(
  name='Header',
  full_name='api.Header',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='api.Header.code', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='message', full_name='api.Header.message', index=1,
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
  serialized_start=21,
  serialized_end=60,
)


_TIMEPERIOD = _descriptor.Descriptor(
  name='TimePeriod',
  full_name='api.TimePeriod',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='start_time_ms', full_name='api.TimePeriod.start_time_ms', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time_ms', full_name='api.TimePeriod.end_time_ms', index=1,
      number=2, type=3, cpp_type=2, label=1,
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
  serialized_start=62,
  serialized_end=118,
)


_LATITUDE = _descriptor.Descriptor(
  name='Latitude',
  full_name='api.Latitude',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='degrees', full_name='api.Latitude.degrees', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='degree_minutes', full_name='api.Latitude.degree_minutes', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='degree_seconds', full_name='api.Latitude.degree_seconds', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='lat_direction', full_name='api.Latitude.lat_direction', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _LATITUDE_LATDIRECTION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=121,
  serialized_end=298,
)


_LONGITUDE = _descriptor.Descriptor(
  name='Longitude',
  full_name='api.Longitude',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='degrees', full_name='api.Longitude.degrees', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='degree_minutes', full_name='api.Longitude.degree_minutes', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='degree_seconds', full_name='api.Longitude.degree_seconds', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='long_direction', full_name='api.Longitude.long_direction', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _LONGITUDE_LONGDIRECTION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=301,
  serialized_end=481,
)


_LOCATION = _descriptor.Descriptor(
  name='Location',
  full_name='api.Location',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='lat', full_name='api.Location.lat', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='long', full_name='api.Location.long', index=1,
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
  serialized_start=483,
  serialized_end=551,
)


_MOTION = _descriptor.Descriptor(
  name='Motion',
  full_name='api.Motion',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='velocity_mph', full_name='api.Motion.velocity_mph', index=0,
      number=1, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='acceleration_mph_s', full_name='api.Motion.acceleration_mph_s', index=1,
      number=2, type=1, cpp_type=5, label=1,
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
  serialized_start=553,
  serialized_end=611,
)


_SOURCE = _descriptor.Descriptor(
  name='Source',
  full_name='api.Source',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='source_id', full_name='api.Source.source_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='source_type', full_name='api.Source.source_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _SOURCE_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=614,
  serialized_end=750,
)

_LATITUDE.fields_by_name['lat_direction'].enum_type = _LATITUDE_LATDIRECTION
_LATITUDE_LATDIRECTION.containing_type = _LATITUDE
_LONGITUDE.fields_by_name['long_direction'].enum_type = _LONGITUDE_LONGDIRECTION
_LONGITUDE_LONGDIRECTION.containing_type = _LONGITUDE
_LOCATION.fields_by_name['lat'].message_type = _LATITUDE
_LOCATION.fields_by_name['long'].message_type = _LONGITUDE
_SOURCE.fields_by_name['source_type'].enum_type = _SOURCE_TYPE
_SOURCE_TYPE.containing_type = _SOURCE
DESCRIPTOR.message_types_by_name['Header'] = _HEADER
DESCRIPTOR.message_types_by_name['TimePeriod'] = _TIMEPERIOD
DESCRIPTOR.message_types_by_name['Latitude'] = _LATITUDE
DESCRIPTOR.message_types_by_name['Longitude'] = _LONGITUDE
DESCRIPTOR.message_types_by_name['Location'] = _LOCATION
DESCRIPTOR.message_types_by_name['Motion'] = _MOTION
DESCRIPTOR.message_types_by_name['Source'] = _SOURCE
DESCRIPTOR.enum_types_by_name['ConditionType'] = _CONDITIONTYPE
DESCRIPTOR.enum_types_by_name['EventType'] = _EVENTTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Header = _reflection.GeneratedProtocolMessageType('Header', (_message.Message,), {
  'DESCRIPTOR' : _HEADER,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:api.Header)
  })
_sym_db.RegisterMessage(Header)

TimePeriod = _reflection.GeneratedProtocolMessageType('TimePeriod', (_message.Message,), {
  'DESCRIPTOR' : _TIMEPERIOD,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:api.TimePeriod)
  })
_sym_db.RegisterMessage(TimePeriod)

Latitude = _reflection.GeneratedProtocolMessageType('Latitude', (_message.Message,), {
  'DESCRIPTOR' : _LATITUDE,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:api.Latitude)
  })
_sym_db.RegisterMessage(Latitude)

Longitude = _reflection.GeneratedProtocolMessageType('Longitude', (_message.Message,), {
  'DESCRIPTOR' : _LONGITUDE,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:api.Longitude)
  })
_sym_db.RegisterMessage(Longitude)

Location = _reflection.GeneratedProtocolMessageType('Location', (_message.Message,), {
  'DESCRIPTOR' : _LOCATION,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:api.Location)
  })
_sym_db.RegisterMessage(Location)

Motion = _reflection.GeneratedProtocolMessageType('Motion', (_message.Message,), {
  'DESCRIPTOR' : _MOTION,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:api.Motion)
  })
_sym_db.RegisterMessage(Motion)

Source = _reflection.GeneratedProtocolMessageType('Source', (_message.Message,), {
  'DESCRIPTOR' : _SOURCE,
  '__module__' : 'common_pb2'
  # @@protoc_insertion_point(class_scope:api.Source)
  })
_sym_db.RegisterMessage(Source)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

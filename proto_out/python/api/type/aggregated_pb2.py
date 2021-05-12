# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: aggregated.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import common_pb2 as common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='aggregated.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\010api/type',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10\x61ggregated.proto\x1a\x0c\x63ommon.proto\"\xf0\x01\n\x10\x44rivingCondition\x12\x37\n\x0e\x63ondition_type\x18\x01 \x01(\x0e\x32\x1f.DrivingCondition.ConditionType\x12\x10\n\x08severity\x18\x02 \x01(\x01\x12\x1b\n\x06source\x18\x03 \x01(\x0b\x32\x0b.api.Source\"t\n\rConditionType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\t\n\x05NIGHT\x10\x01\x12\x08\n\x04RAIN\x10\x02\x12\x08\n\x04SNOW\x10\x03\x12\x07\n\x03ICE\x10\x04\x12\t\n\x05URBAN\x10\x05\x12\x0b\n\x07HIGHWAY\x10\x06\x12\t\n\x05RURAL\x10\x07\x12\x0b\n\x07TRAFFIC\x10\x08\"\x86\x01\n\rDrivingPeriod\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x15\n\rstart_time_ms\x18\x03 \x01(\x03\x12\x13\n\x0b\x65nd_time_ms\x18\x04 \x01(\x03\x12,\n\x11\x64riving_condition\x18\x05 \x03(\x0b\x32\x11.DrivingCondition\"\xca\x01\n\x08SubEvent\x12\'\n\nevent_type\x18\x01 \x01(\x0e\x32\x13.SubEvent.EventType\x12\r\n\x05value\x18\x02 \x01(\x01\x12\x10\n\x08severity\x18\x03 \x01(\x01\x12\x1b\n\x06source\x18\x04 \x01(\x0b\x32\x0b.api.Source\"W\n\tEventType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\x0f\n\x0bLANE_CHANGE\x10\x01\x12\x15\n\x11\x46\x41ST_ACCELERATION\x10\x02\x12\x15\n\x11\x46\x41ST_DECELERATION\x10\x03\"Y\n\x05\x45vent\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x15\n\rstart_time_ms\x18\x03 \x01(\x03\x12\x1c\n\tsub_event\x18\x04 \x03(\x0b\x32\t.SubEventB\nZ\x08\x61pi/typeb\x06proto3'
  ,
  dependencies=[common__pb2.DESCRIPTOR,])



_DRIVINGCONDITION_CONDITIONTYPE = _descriptor.EnumDescriptor(
  name='ConditionType',
  full_name='DrivingCondition.ConditionType',
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
  serialized_start=159,
  serialized_end=275,
)
_sym_db.RegisterEnumDescriptor(_DRIVINGCONDITION_CONDITIONTYPE)

_SUBEVENT_EVENTTYPE = _descriptor.EnumDescriptor(
  name='EventType',
  full_name='SubEvent.EventType',
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
  serialized_start=530,
  serialized_end=617,
)
_sym_db.RegisterEnumDescriptor(_SUBEVENT_EVENTTYPE)


_DRIVINGCONDITION = _descriptor.Descriptor(
  name='DrivingCondition',
  full_name='DrivingCondition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='condition_type', full_name='DrivingCondition.condition_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='severity', full_name='DrivingCondition.severity', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='source', full_name='DrivingCondition.source', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _DRIVINGCONDITION_CONDITIONTYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=35,
  serialized_end=275,
)


_DRIVINGPERIOD = _descriptor.Descriptor(
  name='DrivingPeriod',
  full_name='DrivingPeriod',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='DrivingPeriod.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='DrivingPeriod.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time_ms', full_name='DrivingPeriod.start_time_ms', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time_ms', full_name='DrivingPeriod.end_time_ms', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='driving_condition', full_name='DrivingPeriod.driving_condition', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=278,
  serialized_end=412,
)


_SUBEVENT = _descriptor.Descriptor(
  name='SubEvent',
  full_name='SubEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='event_type', full_name='SubEvent.event_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='SubEvent.value', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='severity', full_name='SubEvent.severity', index=2,
      number=3, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='source', full_name='SubEvent.source', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _SUBEVENT_EVENTTYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=415,
  serialized_end=617,
)


_EVENT = _descriptor.Descriptor(
  name='Event',
  full_name='Event',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='Event.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='Event.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time_ms', full_name='Event.start_time_ms', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sub_event', full_name='Event.sub_event', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=619,
  serialized_end=708,
)

_DRIVINGCONDITION.fields_by_name['condition_type'].enum_type = _DRIVINGCONDITION_CONDITIONTYPE
_DRIVINGCONDITION.fields_by_name['source'].message_type = common__pb2._SOURCE
_DRIVINGCONDITION_CONDITIONTYPE.containing_type = _DRIVINGCONDITION
_DRIVINGPERIOD.fields_by_name['driving_condition'].message_type = _DRIVINGCONDITION
_SUBEVENT.fields_by_name['event_type'].enum_type = _SUBEVENT_EVENTTYPE
_SUBEVENT.fields_by_name['source'].message_type = common__pb2._SOURCE
_SUBEVENT_EVENTTYPE.containing_type = _SUBEVENT
_EVENT.fields_by_name['sub_event'].message_type = _SUBEVENT
DESCRIPTOR.message_types_by_name['DrivingCondition'] = _DRIVINGCONDITION
DESCRIPTOR.message_types_by_name['DrivingPeriod'] = _DRIVINGPERIOD
DESCRIPTOR.message_types_by_name['SubEvent'] = _SUBEVENT
DESCRIPTOR.message_types_by_name['Event'] = _EVENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DrivingCondition = _reflection.GeneratedProtocolMessageType('DrivingCondition', (_message.Message,), {
  'DESCRIPTOR' : _DRIVINGCONDITION,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:DrivingCondition)
  })
_sym_db.RegisterMessage(DrivingCondition)

DrivingPeriod = _reflection.GeneratedProtocolMessageType('DrivingPeriod', (_message.Message,), {
  'DESCRIPTOR' : _DRIVINGPERIOD,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:DrivingPeriod)
  })
_sym_db.RegisterMessage(DrivingPeriod)

SubEvent = _reflection.GeneratedProtocolMessageType('SubEvent', (_message.Message,), {
  'DESCRIPTOR' : _SUBEVENT,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:SubEvent)
  })
_sym_db.RegisterMessage(SubEvent)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), {
  'DESCRIPTOR' : _EVENT,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:Event)
  })
_sym_db.RegisterMessage(Event)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

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
  package='api',
  syntax='proto3',
  serialized_options=b'Z\010api/type',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10\x61ggregated.proto\x12\x03\x61pi\x1a\x0c\x63ommon.proto\"W\n\x0cTripInternal\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x15\n\rstart_time_ms\x18\x03 \x01(\x03\x12\x13\n\x0b\x65nd_time_ms\x18\x04 \x01(\x03\"\xe0\x01\n\x18\x44rivingConditionInternal\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x0f\n\x07trip_id\x18\x03 \x01(\t\x12*\n\x0e\x63ondition_type\x18\x04 \x01(\x0e\x32\x12.api.ConditionType\x12\x10\n\x08severity\x18\x05 \x01(\x01\x12\x15\n\rstart_time_ms\x18\x06 \x01(\x03\x12\x13\n\x0b\x65nd_time_ms\x18\x07 \x01(\x03\x12\x1b\n\x06source\x18\x08 \x01(\x0b\x32\x0b.api.Source\x12\x0f\n\x07\x61lgoTag\x18\t \x03(\t\"\xc6\x01\n\rEventInternal\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x0f\n\x07trip_id\x18\x03 \x01(\t\x12\"\n\nevent_type\x18\x04 \x01(\x0e\x32\x0e.api.EventType\x12\r\n\x05value\x18\x05 \x01(\x01\x12\x10\n\x08severity\x18\x06 \x01(\x01\x12\x14\n\x0ctimestamp_ms\x18\x07 \x01(\x03\x12\x1b\n\x06source\x18\x08 \x01(\x0b\x32\x0b.api.Source\x12\x0f\n\x07\x61lgoTag\x18\t \x03(\t\"H\n\x12\x45ventCreateRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12!\n\x05\x65vent\x18\x02 \x01(\x0b\x32\x12.api.EventInternal\"f\n\x13\x45ventCreateResponse\x12\x1b\n\x06header\x18\x01 \x01(\x0b\x32\x0b.api.Header\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12!\n\x05\x65vent\x18\x03 \x01(\x0b\x32\x12.api.EventInternal\"j\n\x1d\x44rivingConditionCreateRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x38\n\x11\x64riving_condition\x18\x02 \x01(\x0b\x32\x1d.api.DrivingConditionInternal\"\x88\x01\n\x1e\x44rivingConditionCreateResponse\x12\x1b\n\x06header\x18\x01 \x01(\x0b\x32\x0b.api.Header\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x38\n\x11\x64riving_condition\x18\x03 \x01(\x0b\x32\x1d.api.DrivingConditionInternalB\nZ\x08\x61pi/typeb\x06proto3'
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_TRIPINTERNAL = _descriptor.Descriptor(
  name='TripInternal',
  full_name='api.TripInternal',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.TripInternal.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.TripInternal.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time_ms', full_name='api.TripInternal.start_time_ms', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time_ms', full_name='api.TripInternal.end_time_ms', index=3,
      number=4, type=3, cpp_type=2, label=1,
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
  serialized_start=39,
  serialized_end=126,
)


_DRIVINGCONDITIONINTERNAL = _descriptor.Descriptor(
  name='DrivingConditionInternal',
  full_name='api.DrivingConditionInternal',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.DrivingConditionInternal.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.DrivingConditionInternal.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='trip_id', full_name='api.DrivingConditionInternal.trip_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='condition_type', full_name='api.DrivingConditionInternal.condition_type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='severity', full_name='api.DrivingConditionInternal.severity', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time_ms', full_name='api.DrivingConditionInternal.start_time_ms', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time_ms', full_name='api.DrivingConditionInternal.end_time_ms', index=6,
      number=7, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='source', full_name='api.DrivingConditionInternal.source', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='algoTag', full_name='api.DrivingConditionInternal.algoTag', index=8,
      number=9, type=9, cpp_type=9, label=3,
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
  serialized_start=129,
  serialized_end=353,
)


_EVENTINTERNAL = _descriptor.Descriptor(
  name='EventInternal',
  full_name='api.EventInternal',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.EventInternal.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.EventInternal.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='trip_id', full_name='api.EventInternal.trip_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='event_type', full_name='api.EventInternal.event_type', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='api.EventInternal.value', index=4,
      number=5, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='severity', full_name='api.EventInternal.severity', index=5,
      number=6, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp_ms', full_name='api.EventInternal.timestamp_ms', index=6,
      number=7, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='source', full_name='api.EventInternal.source', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='algoTag', full_name='api.EventInternal.algoTag', index=8,
      number=9, type=9, cpp_type=9, label=3,
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
  serialized_start=356,
  serialized_end=554,
)


_EVENTCREATEREQUEST = _descriptor.Descriptor(
  name='EventCreateRequest',
  full_name='api.EventCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.EventCreateRequest.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='event', full_name='api.EventCreateRequest.event', index=1,
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
  serialized_start=556,
  serialized_end=628,
)


_EVENTCREATERESPONSE = _descriptor.Descriptor(
  name='EventCreateResponse',
  full_name='api.EventCreateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='api.EventCreateResponse.header', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.EventCreateResponse.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='event', full_name='api.EventCreateResponse.event', index=2,
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
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=630,
  serialized_end=732,
)


_DRIVINGCONDITIONCREATEREQUEST = _descriptor.Descriptor(
  name='DrivingConditionCreateRequest',
  full_name='api.DrivingConditionCreateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.DrivingConditionCreateRequest.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='driving_condition', full_name='api.DrivingConditionCreateRequest.driving_condition', index=1,
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
  serialized_start=734,
  serialized_end=840,
)


_DRIVINGCONDITIONCREATERESPONSE = _descriptor.Descriptor(
  name='DrivingConditionCreateResponse',
  full_name='api.DrivingConditionCreateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='header', full_name='api.DrivingConditionCreateResponse.header', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_id', full_name='api.DrivingConditionCreateResponse.user_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='driving_condition', full_name='api.DrivingConditionCreateResponse.driving_condition', index=2,
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
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=843,
  serialized_end=979,
)

_DRIVINGCONDITIONINTERNAL.fields_by_name['condition_type'].enum_type = common__pb2._CONDITIONTYPE
_DRIVINGCONDITIONINTERNAL.fields_by_name['source'].message_type = common__pb2._SOURCE
_EVENTINTERNAL.fields_by_name['event_type'].enum_type = common__pb2._EVENTTYPE
_EVENTINTERNAL.fields_by_name['source'].message_type = common__pb2._SOURCE
_EVENTCREATEREQUEST.fields_by_name['event'].message_type = _EVENTINTERNAL
_EVENTCREATERESPONSE.fields_by_name['header'].message_type = common__pb2._HEADER
_EVENTCREATERESPONSE.fields_by_name['event'].message_type = _EVENTINTERNAL
_DRIVINGCONDITIONCREATEREQUEST.fields_by_name['driving_condition'].message_type = _DRIVINGCONDITIONINTERNAL
_DRIVINGCONDITIONCREATERESPONSE.fields_by_name['header'].message_type = common__pb2._HEADER
_DRIVINGCONDITIONCREATERESPONSE.fields_by_name['driving_condition'].message_type = _DRIVINGCONDITIONINTERNAL
DESCRIPTOR.message_types_by_name['TripInternal'] = _TRIPINTERNAL
DESCRIPTOR.message_types_by_name['DrivingConditionInternal'] = _DRIVINGCONDITIONINTERNAL
DESCRIPTOR.message_types_by_name['EventInternal'] = _EVENTINTERNAL
DESCRIPTOR.message_types_by_name['EventCreateRequest'] = _EVENTCREATEREQUEST
DESCRIPTOR.message_types_by_name['EventCreateResponse'] = _EVENTCREATERESPONSE
DESCRIPTOR.message_types_by_name['DrivingConditionCreateRequest'] = _DRIVINGCONDITIONCREATEREQUEST
DESCRIPTOR.message_types_by_name['DrivingConditionCreateResponse'] = _DRIVINGCONDITIONCREATERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TripInternal = _reflection.GeneratedProtocolMessageType('TripInternal', (_message.Message,), {
  'DESCRIPTOR' : _TRIPINTERNAL,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:api.TripInternal)
  })
_sym_db.RegisterMessage(TripInternal)

DrivingConditionInternal = _reflection.GeneratedProtocolMessageType('DrivingConditionInternal', (_message.Message,), {
  'DESCRIPTOR' : _DRIVINGCONDITIONINTERNAL,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:api.DrivingConditionInternal)
  })
_sym_db.RegisterMessage(DrivingConditionInternal)

EventInternal = _reflection.GeneratedProtocolMessageType('EventInternal', (_message.Message,), {
  'DESCRIPTOR' : _EVENTINTERNAL,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:api.EventInternal)
  })
_sym_db.RegisterMessage(EventInternal)

EventCreateRequest = _reflection.GeneratedProtocolMessageType('EventCreateRequest', (_message.Message,), {
  'DESCRIPTOR' : _EVENTCREATEREQUEST,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:api.EventCreateRequest)
  })
_sym_db.RegisterMessage(EventCreateRequest)

EventCreateResponse = _reflection.GeneratedProtocolMessageType('EventCreateResponse', (_message.Message,), {
  'DESCRIPTOR' : _EVENTCREATERESPONSE,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:api.EventCreateResponse)
  })
_sym_db.RegisterMessage(EventCreateResponse)

DrivingConditionCreateRequest = _reflection.GeneratedProtocolMessageType('DrivingConditionCreateRequest', (_message.Message,), {
  'DESCRIPTOR' : _DRIVINGCONDITIONCREATEREQUEST,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:api.DrivingConditionCreateRequest)
  })
_sym_db.RegisterMessage(DrivingConditionCreateRequest)

DrivingConditionCreateResponse = _reflection.GeneratedProtocolMessageType('DrivingConditionCreateResponse', (_message.Message,), {
  'DESCRIPTOR' : _DRIVINGCONDITIONCREATERESPONSE,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:api.DrivingConditionCreateResponse)
  })
_sym_db.RegisterMessage(DrivingConditionCreateResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)

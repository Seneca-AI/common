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




DESCRIPTOR = _descriptor.FileDescriptor(
  name='aggregated.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\010api/type',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10\x61ggregated.proto\"@\n\x0e\x41ggregatedTrip\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x11\n\tevent_ids\x18\x04 \x03(\tB\nZ\x08\x61pi/typeb\x06proto3'
)




_AGGREGATEDTRIP = _descriptor.Descriptor(
  name='AggregatedTrip',
  full_name='AggregatedTrip',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='AggregatedTrip.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='AggregatedTrip.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='event_ids', full_name='AggregatedTrip.event_ids', index=2,
      number=4, type=9, cpp_type=9, label=3,
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
  serialized_start=20,
  serialized_end=84,
)

DESCRIPTOR.message_types_by_name['AggregatedTrip'] = _AGGREGATEDTRIP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AggregatedTrip = _reflection.GeneratedProtocolMessageType('AggregatedTrip', (_message.Message,), {
  'DESCRIPTOR' : _AGGREGATEDTRIP,
  '__module__' : 'aggregated_pb2'
  # @@protoc_insertion_point(class_scope:AggregatedTrip)
  })
_sym_db.RegisterMessage(AggregatedTrip)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
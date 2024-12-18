# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/main.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10proto/main.proto\x12\x04main\" \n\x04Item\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"\x8d\x04\n\x06Record\x12\n\n\x02id\x18\x01 \x01(\t\x12\x15\n\roriginal_name\x18\x02 \x01(\t\x12\x14\n\x0c\x63leaned_name\x18\x03 \x01(\t\x12\x0e\n\x06rule_1\x18\x04 \x01(\x08\x12\x0e\n\x06rule_2\x18\x05 \x01(\x08\x12\x1a\n\x12rule_2_description\x18\x06 \x01(\t\x12\x0e\n\x06rule_3\x18\x07 \x01(\x08\x12\x1a\n\x12rule_3_description\x18\x08 \x01(\t\x12\x0e\n\x06rule_4\x18\t \x01(\x08\x12\x1a\n\x12rule_4_description\x18\n \x01(\t\x12\x0e\n\x06rule_5\x18\x0b \x01(\x08\x12\x1a\n\x12rule_5_description\x18\x0c \x01(\t\x12\x0e\n\x06rule_6\x18\r \x01(\x08\x12\x1a\n\x12rule_6_description\x18\x0e \x01(\t\x12\x0e\n\x06rule_7\x18\x0f \x01(\x08\x12\x1a\n\x12rule_7_description\x18\x10 \x01(\t\x12\x0e\n\x06rule_8\x18\x11 \x01(\x08\x12\x1a\n\x12rule_8_description\x18\x12 \x01(\t\x12\x0e\n\x06rule_9\x18\x13 \x01(\x08\x12\x1a\n\x12rule_9_description\x18\x14 \x01(\t\x12\x0f\n\x07rule_10\x18\x15 \x01(\x08\x12\x1b\n\x13rule_10_description\x18\x16 \x01(\t\x12\x0f\n\x07rule_17\x18\x17 \x01(\x08\x12\x1b\n\x13rule_17_description\x18\x18 \x01(\t\"\xd4\x02\n\x07Summary\x12\x13\n\x0btotal_names\x18\x01 \x01(\r\x12\x1a\n\x12rule_1_false_count\x18\x02 \x01(\r\x12\x1a\n\x12rule_2_error_count\x18\x03 \x01(\r\x12\x1a\n\x12rule_3_error_count\x18\x04 \x01(\r\x12\x1a\n\x12rule_4_error_count\x18\x05 \x01(\r\x12\x1a\n\x12rule_5_error_count\x18\x06 \x01(\r\x12\x1a\n\x12rule_6_error_count\x18\x07 \x01(\r\x12\x1a\n\x12rule_7_error_count\x18\x08 \x01(\r\x12\x1a\n\x12rule_8_error_count\x18\t \x01(\r\x12\x1a\n\x12rule_9_error_count\x18\n \x01(\r\x12\x1b\n\x13rule_10_error_count\x18\x0b \x01(\r\x12\x1b\n\x13rule_17_error_count\x18\x0c \x01(\r\"-\n\x10NormalizeRequest\x12\x19\n\x05items\x18\x01 \x03(\x0b\x32\n.main.Item\"R\n\x11NormalizeResponse\x12\x1d\n\x07records\x18\x01 \x03(\x0b\x32\x0c.main.Record\x12\x1e\n\x07summary\x18\x02 \x01(\x0b\x32\r.main.Summary\"-\n\x10\x44uplicateRequest\x12\x19\n\x05items\x18\x01 \x03(\x0b\x32\n.main.Item\"M\n\x11\x44uplicateResponse\x12\x0f\n\x07message\x18\x01 \x01(\t\x12\x13\n\x0btotal_items\x18\x02 \x01(\x05\x12\x12\n\nduplicates\x18\x03 \x03(\t2P\n\x10NormalizeService\x12<\n\tNormalize\x12\x16.main.NormalizeRequest\x1a\x17.main.NormalizeResponse2V\n\x12NormalizePyService\x12@\n\rGetDuplicates\x12\x16.main.DuplicateRequest\x1a\x17.main.DuplicateResponseB\x06Z\x04/apib\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.main_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\004/api'
  _ITEM._serialized_start=26
  _ITEM._serialized_end=58
  _RECORD._serialized_start=61
  _RECORD._serialized_end=586
  _SUMMARY._serialized_start=589
  _SUMMARY._serialized_end=929
  _NORMALIZEREQUEST._serialized_start=931
  _NORMALIZEREQUEST._serialized_end=976
  _NORMALIZERESPONSE._serialized_start=978
  _NORMALIZERESPONSE._serialized_end=1060
  _DUPLICATEREQUEST._serialized_start=1062
  _DUPLICATEREQUEST._serialized_end=1107
  _DUPLICATERESPONSE._serialized_start=1109
  _DUPLICATERESPONSE._serialized_end=1186
  _NORMALIZESERVICE._serialized_start=1188
  _NORMALIZESERVICE._serialized_end=1268
  _NORMALIZEPYSERVICE._serialized_start=1270
  _NORMALIZEPYSERVICE._serialized_end=1356
# @@protoc_insertion_point(module_scope)
// source: aggregated.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

goog.provide('proto.api.DrivingConditionInternal');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.api.Source');

goog.forwardDeclare('proto.api.ConditionType');
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.api.DrivingConditionInternal = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.DrivingConditionInternal, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.api.DrivingConditionInternal.displayName = 'proto.api.DrivingConditionInternal';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.api.DrivingConditionInternal.prototype.toObject = function(opt_includeInstance) {
  return proto.api.DrivingConditionInternal.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.DrivingConditionInternal} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.DrivingConditionInternal.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    userId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    tripId: jspb.Message.getFieldWithDefault(msg, 3, ""),
    conditionType: jspb.Message.getFieldWithDefault(msg, 4, 0),
    severity: jspb.Message.getFloatingPointFieldWithDefault(msg, 5, 0.0),
    startTimeMs: jspb.Message.getFieldWithDefault(msg, 6, 0),
    endTimeMs: jspb.Message.getFieldWithDefault(msg, 7, 0),
    source: (f = msg.getSource()) && proto.api.Source.toObject(includeInstance, f),
    algoTag: jspb.Message.getFieldWithDefault(msg, 9, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.api.DrivingConditionInternal}
 */
proto.api.DrivingConditionInternal.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.DrivingConditionInternal;
  return proto.api.DrivingConditionInternal.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.DrivingConditionInternal} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.DrivingConditionInternal}
 */
proto.api.DrivingConditionInternal.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setTripId(value);
      break;
    case 4:
      var value = /** @type {!proto.api.ConditionType} */ (reader.readEnum());
      msg.setConditionType(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setSeverity(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setStartTimeMs(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEndTimeMs(value);
      break;
    case 8:
      var value = new proto.api.Source;
      reader.readMessage(value,proto.api.Source.deserializeBinaryFromReader);
      msg.setSource(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setAlgoTag(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.api.DrivingConditionInternal.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.DrivingConditionInternal.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.DrivingConditionInternal} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.DrivingConditionInternal.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTripId();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getConditionType();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = message.getSeverity();
  if (f !== 0.0) {
    writer.writeDouble(
      5,
      f
    );
  }
  f = message.getStartTimeMs();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getEndTimeMs();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getSource();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      proto.api.Source.serializeBinaryToWriter
    );
  }
  f = message.getAlgoTag();
  if (f.length > 0) {
    writer.writeString(
      9,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.api.DrivingConditionInternal.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string user_id = 2;
 * @return {string}
 */
proto.api.DrivingConditionInternal.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setUserId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string trip_id = 3;
 * @return {string}
 */
proto.api.DrivingConditionInternal.prototype.getTripId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setTripId = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional ConditionType condition_type = 4;
 * @return {!proto.api.ConditionType}
 */
proto.api.DrivingConditionInternal.prototype.getConditionType = function() {
  return /** @type {!proto.api.ConditionType} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.api.ConditionType} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setConditionType = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * optional double severity = 5;
 * @return {number}
 */
proto.api.DrivingConditionInternal.prototype.getSeverity = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 5, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setSeverity = function(value) {
  return jspb.Message.setProto3FloatField(this, 5, value);
};


/**
 * optional int64 start_time_ms = 6;
 * @return {number}
 */
proto.api.DrivingConditionInternal.prototype.getStartTimeMs = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setStartTimeMs = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int64 end_time_ms = 7;
 * @return {number}
 */
proto.api.DrivingConditionInternal.prototype.getEndTimeMs = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {number} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setEndTimeMs = function(value) {
  return jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional Source source = 8;
 * @return {?proto.api.Source}
 */
proto.api.DrivingConditionInternal.prototype.getSource = function() {
  return /** @type{?proto.api.Source} */ (
    jspb.Message.getWrapperField(this, proto.api.Source, 8));
};


/**
 * @param {?proto.api.Source|undefined} value
 * @return {!proto.api.DrivingConditionInternal} returns this
*/
proto.api.DrivingConditionInternal.prototype.setSource = function(value) {
  return jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.clearSource = function() {
  return this.setSource(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.api.DrivingConditionInternal.prototype.hasSource = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional string algo_tag = 9;
 * @return {string}
 */
proto.api.DrivingConditionInternal.prototype.getAlgoTag = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.DrivingConditionInternal} returns this
 */
proto.api.DrivingConditionInternal.prototype.setAlgoTag = function(value) {
  return jspb.Message.setProto3StringField(this, 9, value);
};


// source: common.proto
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

goog.provide('proto.api.Latitude');
goog.provide('proto.api.Latitude.LatDirection');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');

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
proto.api.Latitude = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.Latitude, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.api.Latitude.displayName = 'proto.api.Latitude';
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
proto.api.Latitude.prototype.toObject = function(opt_includeInstance) {
  return proto.api.Latitude.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.Latitude} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.Latitude.toObject = function(includeInstance, msg) {
  var f, obj = {
    degrees: jspb.Message.getFieldWithDefault(msg, 1, 0),
    degreeMinutes: jspb.Message.getFieldWithDefault(msg, 2, 0),
    degreeSeconds: jspb.Message.getFloatingPointFieldWithDefault(msg, 3, 0.0),
    latDirection: jspb.Message.getFieldWithDefault(msg, 4, 0)
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
 * @return {!proto.api.Latitude}
 */
proto.api.Latitude.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.Latitude;
  return proto.api.Latitude.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.Latitude} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.Latitude}
 */
proto.api.Latitude.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setDegrees(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setDegreeMinutes(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setDegreeSeconds(value);
      break;
    case 4:
      var value = /** @type {!proto.api.Latitude.LatDirection} */ (reader.readEnum());
      msg.setLatDirection(value);
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
proto.api.Latitude.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.Latitude.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.Latitude} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.Latitude.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getDegrees();
  if (f !== 0) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = message.getDegreeMinutes();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getDegreeSeconds();
  if (f !== 0.0) {
    writer.writeDouble(
      3,
      f
    );
  }
  f = message.getLatDirection();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.api.Latitude.LatDirection = {
  UNKNOWN: 0,
  NORTH: 1,
  SOUTH: 2
};

/**
 * optional int32 degrees = 1;
 * @return {number}
 */
proto.api.Latitude.prototype.getDegrees = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.api.Latitude} returns this
 */
proto.api.Latitude.prototype.setDegrees = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int32 degree_minutes = 2;
 * @return {number}
 */
proto.api.Latitude.prototype.getDegreeMinutes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.api.Latitude} returns this
 */
proto.api.Latitude.prototype.setDegreeMinutes = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional double degree_seconds = 3;
 * @return {number}
 */
proto.api.Latitude.prototype.getDegreeSeconds = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.api.Latitude} returns this
 */
proto.api.Latitude.prototype.setDegreeSeconds = function(value) {
  return jspb.Message.setProto3FloatField(this, 3, value);
};


/**
 * optional LatDirection lat_direction = 4;
 * @return {!proto.api.Latitude.LatDirection}
 */
proto.api.Latitude.prototype.getLatDirection = function() {
  return /** @type {!proto.api.Latitude.LatDirection} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.api.Latitude.LatDirection} value
 * @return {!proto.api.Latitude} returns this
 */
proto.api.Latitude.prototype.setLatDirection = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};



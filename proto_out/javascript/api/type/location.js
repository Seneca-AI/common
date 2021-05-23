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

goog.provide('proto.api.Location');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.api.Latitude');
goog.require('proto.api.Longitude');

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
proto.api.Location = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.Location, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.api.Location.displayName = 'proto.api.Location';
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
proto.api.Location.prototype.toObject = function(opt_includeInstance) {
  return proto.api.Location.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.Location} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.Location.toObject = function(includeInstance, msg) {
  var f, obj = {
    lat: (f = msg.getLat()) && proto.api.Latitude.toObject(includeInstance, f),
    pb_long: (f = msg.getLong()) && proto.api.Longitude.toObject(includeInstance, f)
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
 * @return {!proto.api.Location}
 */
proto.api.Location.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.Location;
  return proto.api.Location.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.Location} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.Location}
 */
proto.api.Location.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.api.Latitude;
      reader.readMessage(value,proto.api.Latitude.deserializeBinaryFromReader);
      msg.setLat(value);
      break;
    case 2:
      var value = new proto.api.Longitude;
      reader.readMessage(value,proto.api.Longitude.deserializeBinaryFromReader);
      msg.setLong(value);
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
proto.api.Location.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.Location.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.Location} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.Location.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLat();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.api.Latitude.serializeBinaryToWriter
    );
  }
  f = message.getLong();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.api.Longitude.serializeBinaryToWriter
    );
  }
};


/**
 * optional Latitude lat = 1;
 * @return {?proto.api.Latitude}
 */
proto.api.Location.prototype.getLat = function() {
  return /** @type{?proto.api.Latitude} */ (
    jspb.Message.getWrapperField(this, proto.api.Latitude, 1));
};


/**
 * @param {?proto.api.Latitude|undefined} value
 * @return {!proto.api.Location} returns this
*/
proto.api.Location.prototype.setLat = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.api.Location} returns this
 */
proto.api.Location.prototype.clearLat = function() {
  return this.setLat(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.api.Location.prototype.hasLat = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Longitude long = 2;
 * @return {?proto.api.Longitude}
 */
proto.api.Location.prototype.getLong = function() {
  return /** @type{?proto.api.Longitude} */ (
    jspb.Message.getWrapperField(this, proto.api.Longitude, 2));
};


/**
 * @param {?proto.api.Longitude|undefined} value
 * @return {!proto.api.Location} returns this
*/
proto.api.Location.prototype.setLong = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.api.Location} returns this
 */
proto.api.Location.prototype.clearLong = function() {
  return this.setLong(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.api.Location.prototype.hasLong = function() {
  return jspb.Message.getField(this, 2) != null;
};



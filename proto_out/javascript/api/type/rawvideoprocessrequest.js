// source: raw.proto
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

goog.provide('proto.api.RawVideoProcessRequest');

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
proto.api.RawVideoProcessRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.RawVideoProcessRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.api.RawVideoProcessRequest.displayName = 'proto.api.RawVideoProcessRequest';
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
proto.api.RawVideoProcessRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.api.RawVideoProcessRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.RawVideoProcessRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.RawVideoProcessRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    userId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    videoName: jspb.Message.getFieldWithDefault(msg, 2, ""),
    localPath: jspb.Message.getFieldWithDefault(msg, 3, ""),
    videoBytes: msg.getVideoBytes_asB64()
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
 * @return {!proto.api.RawVideoProcessRequest}
 */
proto.api.RawVideoProcessRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.RawVideoProcessRequest;
  return proto.api.RawVideoProcessRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.RawVideoProcessRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.RawVideoProcessRequest}
 */
proto.api.RawVideoProcessRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setVideoName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setLocalPath(value);
      break;
    case 4:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setVideoBytes(value);
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
proto.api.RawVideoProcessRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.RawVideoProcessRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.RawVideoProcessRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.RawVideoProcessRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getVideoName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getLocalPath();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVideoBytes_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      4,
      f
    );
  }
};


/**
 * optional string user_id = 1;
 * @return {string}
 */
proto.api.RawVideoProcessRequest.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.RawVideoProcessRequest} returns this
 */
proto.api.RawVideoProcessRequest.prototype.setUserId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string video_name = 2;
 * @return {string}
 */
proto.api.RawVideoProcessRequest.prototype.getVideoName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.RawVideoProcessRequest} returns this
 */
proto.api.RawVideoProcessRequest.prototype.setVideoName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string local_path = 3;
 * @return {string}
 */
proto.api.RawVideoProcessRequest.prototype.getLocalPath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.RawVideoProcessRequest} returns this
 */
proto.api.RawVideoProcessRequest.prototype.setLocalPath = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional bytes video_bytes = 4;
 * @return {string}
 */
proto.api.RawVideoProcessRequest.prototype.getVideoBytes = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * optional bytes video_bytes = 4;
 * This is a type-conversion wrapper around `getVideoBytes()`
 * @return {string}
 */
proto.api.RawVideoProcessRequest.prototype.getVideoBytes_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getVideoBytes()));
};


/**
 * optional bytes video_bytes = 4;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getVideoBytes()`
 * @return {!Uint8Array}
 */
proto.api.RawVideoProcessRequest.prototype.getVideoBytes_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getVideoBytes()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.api.RawVideoProcessRequest} returns this
 */
proto.api.RawVideoProcessRequest.prototype.setVideoBytes = function(value) {
  return jspb.Message.setProto3BytesField(this, 4, value);
};



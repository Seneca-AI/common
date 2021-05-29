// source: devops.proto
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

goog.provide('proto.api.SenecaServer');

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
proto.api.SenecaServer = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.api.SenecaServer, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.api.SenecaServer.displayName = 'proto.api.SenecaServer';
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
proto.api.SenecaServer.prototype.toObject = function(opt_includeInstance) {
  return proto.api.SenecaServer.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.api.SenecaServer} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.SenecaServer.toObject = function(includeInstance, msg) {
  var f, obj = {
    projectId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    serverExternalIp: jspb.Message.getFieldWithDefault(msg, 2, ""),
    serverPort: jspb.Message.getFieldWithDefault(msg, 3, ""),
    serverVmName: jspb.Message.getFieldWithDefault(msg, 4, ""),
    serverVmZone: jspb.Message.getFieldWithDefault(msg, 5, ""),
    serverPathToGoogleApplicationCredentials: jspb.Message.getFieldWithDefault(msg, 6, ""),
    serverPathToGoogleOauthCredentials: jspb.Message.getFieldWithDefault(msg, 7, ""),
    receiveMainPushes: jspb.Message.getBooleanFieldWithDefault(msg, 8, false)
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
 * @return {!proto.api.SenecaServer}
 */
proto.api.SenecaServer.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.api.SenecaServer;
  return proto.api.SenecaServer.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.api.SenecaServer} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.api.SenecaServer}
 */
proto.api.SenecaServer.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setProjectId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setServerExternalIp(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setServerPort(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setServerVmName(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setServerVmZone(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setServerPathToGoogleApplicationCredentials(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setServerPathToGoogleOauthCredentials(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setReceiveMainPushes(value);
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
proto.api.SenecaServer.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.api.SenecaServer.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.api.SenecaServer} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.api.SenecaServer.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getProjectId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getServerExternalIp();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getServerPort();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getServerVmName();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getServerVmZone();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getServerPathToGoogleApplicationCredentials();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getServerPathToGoogleOauthCredentials();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getReceiveMainPushes();
  if (f) {
    writer.writeBool(
      8,
      f
    );
  }
};


/**
 * optional string project_id = 1;
 * @return {string}
 */
proto.api.SenecaServer.prototype.getProjectId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setProjectId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string server_external_ip = 2;
 * @return {string}
 */
proto.api.SenecaServer.prototype.getServerExternalIp = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setServerExternalIp = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string server_port = 3;
 * @return {string}
 */
proto.api.SenecaServer.prototype.getServerPort = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setServerPort = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string server_vm_name = 4;
 * @return {string}
 */
proto.api.SenecaServer.prototype.getServerVmName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setServerVmName = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string server_vm_zone = 5;
 * @return {string}
 */
proto.api.SenecaServer.prototype.getServerVmZone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setServerVmZone = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string server_path_to_google_application_credentials = 6;
 * @return {string}
 */
proto.api.SenecaServer.prototype.getServerPathToGoogleApplicationCredentials = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setServerPathToGoogleApplicationCredentials = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string server_path_to_google_oauth_credentials = 7;
 * @return {string}
 */
proto.api.SenecaServer.prototype.getServerPathToGoogleOauthCredentials = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setServerPathToGoogleOauthCredentials = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional bool receive_main_pushes = 8;
 * @return {boolean}
 */
proto.api.SenecaServer.prototype.getReceiveMainPushes = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.api.SenecaServer} returns this
 */
proto.api.SenecaServer.prototype.setReceiveMainPushes = function(value) {
  return jspb.Message.setProto3BooleanField(this, 8, value);
};



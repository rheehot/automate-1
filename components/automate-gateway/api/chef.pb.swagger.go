package api

func init() {
	Swagger.Add("chef", `{
  "swagger": "2.0",
  "info": {
    "title": "api/external/ingest/chef.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/ingest/events/chef/action": {
      "post": {
        "operationId": "ProcessChefAction",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseProcessChefActionResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestAction"
            }
          }
        ],
        "tags": [
          "ChefIngester"
        ]
      }
    },
    "/ingest/events/chef/liveness": {
      "post": {
        "operationId": "ProcessLivenessPing",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseProcessLivenessResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestLiveness"
            }
          }
        ],
        "tags": [
          "ChefIngester"
        ]
      }
    },
    "/ingest/events/chef/node-multiple-deletes": {
      "post": {
        "operationId": "ProcessMultipleNodeDeletes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseProcessMultipleNodeDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestMultipleNodeDeleteRequest"
            }
          }
        ],
        "tags": [
          "ChefIngester"
        ]
      }
    },
    "/ingest/events/chef/nodedelete": {
      "post": {
        "operationId": "ProcessNodeDelete",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseProcessNodeDeleteResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestDelete"
            }
          }
        ],
        "tags": [
          "ChefIngester"
        ]
      }
    },
    "/ingest/events/chef/run": {
      "post": {
        "operationId": "ProcessChefRun",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/responseProcessChefRunResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/requestRun"
            }
          }
        ],
        "tags": [
          "ChefIngester"
        ]
      }
    },
    "/ingest/version": {
      "get": {
        "operationId": "GetVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/versionVersionInfo"
            }
          }
        },
        "tags": [
          "ChefIngester"
        ]
      }
    }
  },
  "definitions": {
    "protobufListValue": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufValue"
          },
          "description": "Repeated field of dynamically typed values."
        }
      },
      "description": "`+"`"+`ListValue`+"`"+` is a wrapper around a repeated field of values.\n\nThe JSON representation for `+"`"+`ListValue`+"`"+` is JSON array."
    },
    "protobufNullValue": {
      "type": "string",
      "enum": [
        "NULL_VALUE"
      ],
      "default": "NULL_VALUE",
      "description": "`+"`"+`NullValue`+"`"+` is a singleton enumeration to represent the null value for the\n`+"`"+`Value`+"`"+` type union.\n\n The JSON representation for `+"`"+`NullValue`+"`"+` is JSON `+"`"+`null`+"`"+`.\n\n - NULL_VALUE: Null value."
    },
    "protobufStruct": {
      "type": "object",
      "properties": {
        "fields": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/protobufValue"
          },
          "description": "Unordered map of dynamically typed values."
        }
      },
      "description": "`+"`"+`Struct`+"`"+` represents a structured data value, consisting of fields\nwhich map to dynamically typed values. In some languages, `+"`"+`Struct`+"`"+`\nmight be supported by a native representation. For example, in\nscripting languages like JS a struct is represented as an\nobject. The details of that representation are described together\nwith the proto support for the language.\n\nThe JSON representation for `+"`"+`Struct`+"`"+` is JSON object."
    },
    "protobufValue": {
      "type": "object",
      "properties": {
        "null_value": {
          "$ref": "#/definitions/protobufNullValue",
          "description": "Represents a null value."
        },
        "number_value": {
          "type": "number",
          "format": "double",
          "description": "Represents a double value."
        },
        "string_value": {
          "type": "string",
          "description": "Represents a string value."
        },
        "bool_value": {
          "type": "boolean",
          "format": "boolean",
          "description": "Represents a boolean value."
        },
        "struct_value": {
          "$ref": "#/definitions/protobufStruct",
          "description": "Represents a structured value."
        },
        "list_value": {
          "$ref": "#/definitions/protobufListValue",
          "description": "Represents a repeated `+"`"+`Value`+"`"+`."
        }
      },
      "description": "`+"`"+`Value`+"`"+` represents a dynamically typed value which can be either\nnull, a number, a string, a boolean, a recursive struct value, or a\nlist of values. A producer of value is expected to set one of that\nvariants, absence of any variant indicates an error.\n\nThe JSON representation for `+"`"+`Value`+"`"+` is JSON value."
    },
    "requestAction": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID of the action message itself"
        },
        "message_type": {
          "type": "string"
        },
        "message_version": {
          "type": "string"
        },
        "entity_name": {
          "type": "string"
        },
        "entity_type": {
          "type": "string"
        },
        "task": {
          "type": "string"
        },
        "organization_name": {
          "type": "string"
        },
        "remote_hostname": {
          "type": "string"
        },
        "run_id": {
          "type": "string"
        },
        "content": {
          "type": "string",
          "format": "byte",
          "description": "This new field called 'content' is being used to send the entire raw JSON\nmessage in bytes, this field is heavily used by the gateway for the DataCollector\nFunc Handler that will send the Action message to the (receiver) ingest-service\nthat will manually unmarshal the message from this field if it is provided.\nThe main purpose of this field it to improve the performance of ingestion when\nthe requests comes in REST/HTTP format."
        },
        "node_id": {
          "type": "string"
        },
        "recorded_at": {
          "type": "string"
        },
        "remote_request_id": {
          "type": "string"
        },
        "request_id": {
          "type": "string"
        },
        "requestor_name": {
          "type": "string"
        },
        "requestor_type": {
          "type": "string"
        },
        "service_hostname": {
          "type": "string"
        },
        "user_agent": {
          "type": "string"
        },
        "parent_type": {
          "type": "string"
        },
        "parent_name": {
          "type": "string"
        },
        "revision_id": {
          "type": "string"
        }
      }
    },
    "requestDelete": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "ID of the action message itself"
        },
        "node_name": {
          "type": "string"
        },
        "organization_name": {
          "type": "string"
        },
        "remote_hostname": {
          "type": "string"
        },
        "service_hostname": {
          "type": "string"
        },
        "node_id": {
          "type": "string"
        }
      }
    },
    "requestDeprecation": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "location": {
          "type": "string"
        }
      }
    },
    "requestDescription": {
      "type": "object",
      "properties": {
        "title": {
          "type": "string"
        },
        "sections": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/protobufStruct"
          }
        }
      }
    },
    "requestError": {
      "type": "object",
      "properties": {
        "class": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "backtrace": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "description": {
          "$ref": "#/definitions/requestDescription"
        }
      }
    },
    "requestExpandedRunList": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "run_list": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/requestRunList"
          }
        }
      }
    },
    "requestLiveness": {
      "type": "object",
      "properties": {
        "event_type": {
          "type": "string"
        },
        "entity_uuid": {
          "type": "string"
        },
        "chef_server_fqdn": {
          "type": "string"
        },
        "source": {
          "type": "string"
        },
        "message_version": {
          "type": "string"
        },
        "organization_name": {
          "type": "string"
        },
        "node_name": {
          "type": "string"
        }
      }
    },
    "requestMultipleNodeDeleteRequest": {
      "type": "object",
      "properties": {
        "node_ids": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "requestResource": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "after": {
          "$ref": "#/definitions/protobufStruct"
        },
        "before": {
          "$ref": "#/definitions/protobufStruct"
        },
        "duration": {
          "type": "string"
        },
        "delta": {
          "type": "string"
        },
        "cookbook_name": {
          "type": "string"
        },
        "cookbook_version": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "recipe_name": {
          "type": "string"
        },
        "ignore_failure": {
          "$ref": "#/definitions/protobufValue"
        },
        "conditional": {
          "type": "string"
        },
        "result": {
          "type": "string"
        }
      }
    },
    "requestRun": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "1 through 15 are for frequently occuring fields\nReserving for shared fields between run_start and run_converge mesages."
        },
        "run_id": {
          "type": "string"
        },
        "entity_uuid": {
          "type": "string"
        },
        "message_version": {
          "type": "string"
        },
        "message_type": {
          "type": "string"
        },
        "node_name": {
          "type": "string"
        },
        "organization_name": {
          "type": "string"
        },
        "start_time": {
          "type": "string"
        },
        "chef_server_fqdn": {
          "type": "string"
        },
        "content": {
          "type": "string",
          "format": "byte",
          "description": "This new field called 'content' is being used to send the entire raw JSON\nmessage in bytes, this field is heavily used by the gateway for the DataCollector\nFunc Handler that will send the Run message to the (receiver) ingest-service\nthat will manually unmarshal the message from this field if it is provided.\nThe main purpose of this field it to improve the performance of ingestion when\nthe requests comes in REST/HTTP format."
        },
        "end_time": {
          "type": "string"
        },
        "status": {
          "type": "string"
        },
        "total_resource_count": {
          "type": "integer",
          "format": "int32"
        },
        "updated_resource_count": {
          "type": "integer",
          "format": "int32"
        },
        "source": {
          "type": "string"
        },
        "expanded_run_list": {
          "$ref": "#/definitions/requestExpandedRunList"
        },
        "resources": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/requestResource"
          }
        },
        "run_list": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "node": {
          "$ref": "#/definitions/protobufStruct"
        },
        "error": {
          "$ref": "#/definitions/requestError"
        },
        "policy_name": {
          "type": "string"
        },
        "policy_group": {
          "type": "string"
        },
        "deprecations": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/requestDeprecation"
          }
        },
        "tags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "requestRunList": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "skipped": {
          "type": "boolean",
          "format": "boolean"
        },
        "children": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/requestRunList"
          }
        }
      }
    },
    "responseProcessChefActionResponse": {
      "type": "object"
    },
    "responseProcessChefRunResponse": {
      "type": "object"
    },
    "responseProcessLivenessResponse": {
      "type": "object"
    },
    "responseProcessMultipleNodeDeleteResponse": {
      "type": "object"
    },
    "responseProcessNodeDeleteResponse": {
      "type": "object"
    },
    "versionVersionInfo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "sha": {
          "type": "string"
        },
        "built": {
          "type": "string"
        }
      }
    }
  }
}
`)
}

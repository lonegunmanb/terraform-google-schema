package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEventarcPipeline = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "User-defined annotations. See https://google.aip.dev/128#annotations.\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The creation time.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up\nto nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and\n\"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "crypto_key_name": {
        "description": "Resource name of a KMS crypto key (managed by the user) used to\nencrypt/decrypt the event data. If not set, an internal Google-owned key\nwill be used to encrypt messages. It must match the pattern\n\"projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name of resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "etag": {
        "computed": true,
        "description": "This checksum is computed by the server based on the value of\nother fields, and might be sent only on create requests to ensure that the\nclient has an up-to-date value before proceeding.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "User labels attached to the Pipeline that can be used to group\nresources. An object containing a list of \"key\": value pairs. Example: {\n\"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Pipeline. Must be unique within the\nlocation of the project and must be in\n'projects/{project}/locations/{location}/pipelines/{pipeline}' format.",
        "description_kind": "plain",
        "type": "string"
      },
      "pipeline_id": {
        "description": "The user-provided ID to be assigned to the Pipeline. It should match the\nformat '^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "uid": {
        "computed": true,
        "description": "Server-assigned unique identifier for the Pipeline. The value\nis a UUID4 string and guaranteed to remain unchanged until the resource is\ndeleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The last-modified time.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up\nto nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and\n\"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "destinations": {
        "block": {
          "attributes": {
            "message_bus": {
              "description": "The resource name of the Message Bus to which events should be\npublished. The Message Bus resource should exist in the same project as\nthe Pipeline. Format:\n'projects/{project}/locations/{location}/messageBuses/{message_bus}'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
              "description": "The resource name of the Pub/Sub topic to which events should be\npublished. Format:\n'projects/{project}/locations/{location}/topics/{topic}'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "workflow": {
              "description": "The resource name of the Workflow whose Executions are triggered by\nthe events. The Workflow resource should be deployed in the same\nproject as the Pipeline. Format:\n'projects/{project}/locations/{location}/workflows/{workflow}'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "authentication_config": {
              "block": {
                "block_types": {
                  "google_oidc": {
                    "block": {
                      "attributes": {
                        "audience": {
                          "description": "Audience to be used to generate the OIDC Token. The audience claim\nidentifies the recipient that the JWT is intended for. If\nunspecified, the destination URI will be used.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "service_account": {
                          "description": "Service account email used to generate the OIDC Token.\nThe principal who calls this API must have\niam.serviceAccounts.actAs permission in the service account. See\nhttps://cloud.google.com/iam/docs/understanding-service-accounts\nfor more information. Eventarc service agents must have\nroles/roles/iam.serviceAccountTokenCreator role to allow the\nPipeline to create OpenID tokens for authenticated requests.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Represents a config used to authenticate with a Google OIDC token using\na GCP service account. Use this authentication method to invoke your\nCloud Run and Cloud Functions destinations or HTTP endpoints that\nsupport Google OIDC.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "oauth_token": {
                    "block": {
                      "attributes": {
                        "scope": {
                          "description": "OAuth scope to be used for generating OAuth access token. If not\nspecified, \"https://www.googleapis.com/auth/cloud-platform\" will be\nused.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "service_account": {
                          "description": "Service account email used to generate the [OAuth\ntoken](https://developers.google.com/identity/protocols/OAuth2).\nThe principal who calls this API must have\niam.serviceAccounts.actAs permission in the service account. See\nhttps://cloud.google.com/iam/docs/understanding-service-accounts\nfor more information. Eventarc service agents must have\nroles/roles/iam.serviceAccountTokenCreator role to allow Pipeline\nto create OAuth2 tokens for authenticated requests.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Contains information needed for generating an\n[OAuth token](https://developers.google.com/identity/protocols/OAuth2).\nThis type of authorization should generally only be used when calling\nGoogle APIs hosted on *.googleapis.com.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Represents a config used to authenticate message requests.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "http_endpoint": {
              "block": {
                "attributes": {
                  "message_binding_template": {
                    "description": "The CEL expression used to modify how the destination-bound HTTP\nrequest is constructed.\n\nIf a binding expression is not specified here, the message\nis treated as a CloudEvent and is mapped to the HTTP request according\nto the CloudEvent HTTP Protocol Binding Binary Content Mode\n(https://github.com/cloudevents/spec/blob/main/cloudevents/bindings/http-protocol-binding.md#31-binary-content-mode).\nIn this representation, all fields except the 'data' and\n'datacontenttype' field on the message are mapped to HTTP request\nheaders with a prefix of 'ce-'.\n\nTo construct the HTTP request payload and the value of the content-type\nHTTP header, the payload format is defined as follows:\n1) Use the output_payload_format_type on the Pipeline.Destination if it\nis set, else:\n2) Use the input_payload_format_type on the Pipeline if it is set,\nelse:\n3) Treat the payload as opaque binary data.\n\nThe 'data' field of the message is converted to the payload format or\nleft as-is for case 3) and then attached as the payload of the HTTP\nrequest. The 'content-type' header on the HTTP request is set to the\npayload format type or left empty for case 3). However, if a mediation\nhas updated the 'datacontenttype' field on the message so that it is\nnot the same as the payload format type but it is still a prefix of the\npayload format type, then the 'content-type' header on the HTTP request\nis set to this 'datacontenttype' value. For example, if the\n'datacontenttype' is \"application/json\" and the payload format type is\n\"application/json; charset=utf-8\", then the 'content-type' header on\nthe HTTP request is set to \"application/json; charset=utf-8\".\n\nIf a non-empty binding expression is specified then this expression is\nused to modify the default CloudEvent HTTP Protocol Binding Binary\nContent representation.\nThe result of the CEL expression must be a map of key/value pairs\nwhich is used as follows:\n- If a map named 'headers' exists on the result of the expression,\nthen its key/value pairs are directly mapped to the HTTP request\nheaders. The headers values are constructed from the corresponding\nvalue type's canonical representation. If the 'headers' field doesn't\nexist then the resulting HTTP request will be the headers of the\nCloudEvent HTTP Binding Binary Content Mode representation of the final\nmessage. Note: If the specified binding expression, has updated the\n'datacontenttype' field on the message so that it is not the same as\nthe payload format type but it is still a prefix of the payload format\ntype, then the 'content-type' header in the 'headers' map is set to\nthis 'datacontenttype' value.\n- If a field named 'body' exists on the result of the expression then\nits value is directly mapped to the body of the request. If the value\nof the 'body' field is of type bytes or string then it is used for\nthe HTTP request body as-is, with no conversion. If the body field is\nof any other type then it is converted to a JSON string. If the body\nfield does not exist then the resulting payload of the HTTP request\nwill be data value of the CloudEvent HTTP Binding Binary Content Mode\nrepresentation of the final message as described earlier.\n- Any other fields in the resulting expression will be ignored.\n\nThe CEL expression may access the incoming CloudEvent message in its\ndefinition, as follows:\n- The 'data' field of the incoming CloudEvent message can be accessed\nusing the 'message.data' value. Subfields of 'message.data' may also be\naccessed if an input_payload_format has been specified on the Pipeline.\n- Each attribute of the incoming CloudEvent message can be accessed\nusing the 'message.' value, where  is replaced with the\nname of the attribute.\n- Existing headers can be accessed in the CEL expression using the\n'headers' variable. The 'headers' variable defines a map of key/value\npairs corresponding to the HTTP headers of the CloudEvent HTTP Binding\nBinary Content Mode representation of the final message as described\nearlier. For example, the following CEL expression can be used to\nconstruct an HTTP request by adding an additional header to the HTTP\nheaders of the CloudEvent HTTP Binding Binary Content Mode\nrepresentation of the final message and by overwriting the body of the\nrequest:\n\n'''\n{\n\"headers\": headers.merge({\"new-header-key\": \"new-header-value\"}),\n\"body\": \"new-body\"\n}\n'''\n- The default binding for the message payload can be accessed using the\n'body' variable. It conatins a string representation of the message\npayload in the format specified by the 'output_payload_format' field.\nIf the 'input_payload_format' field is not set, the 'body'\nvariable contains the same message payload bytes that were published.\n\nAdditionally, the following CEL extension functions are provided for\nuse in this CEL expression:\n- toBase64Url:\nmap.toBase64Url() -\u003e string\n- Converts a CelValue to a base64url encoded string\n- toJsonString: map.toJsonString() -\u003e string\n- Converts a CelValue to a JSON string\n- merge:\nmap1.merge(map2) -\u003e map3\n- Merges the passed CEL map with the existing CEL map the\nfunction is applied to.\n- If the same key exists in both maps, if the key's value is type\nmap both maps are merged else the value from the passed map is\nused.\n- denormalize:\nmap.denormalize() -\u003e map\n- Denormalizes a CEL map such that every value of type map or key\nin the map is expanded to return a single level map.\n- The resulting keys are \".\" separated indices of the map keys.\n- For example:\n{\n\"a\": 1,\n\"b\": {\n\"c\": 2,\n\"d\": 3\n}\n\"e\": [4, 5]\n}\n.denormalize()\n-\u003e {\n\"a\": 1,\n\"b.c\": 2,\n\"b.d\": 3,\n\"e.0\": 4,\n\"e.1\": 5\n}\n- setField:\nmap.setField(key, value) -\u003e message\n- Sets the field of the message with the given key to the\ngiven value.\n- If the field is not present it will be added.\n- If the field is present it will be overwritten.\n- The key can be a dot separated path to set a field in a nested\nmessage.\n- Key must be of type string.\n- Value may be any valid type.\n- removeFields:\nmap.removeFields([key1, key2, ...]) -\u003e message\n- Removes the fields of the map with the given keys.\n- The keys can be a dot separated path to remove a field in a\nnested message.\n- If a key is not found it will be ignored.\n- Keys must be of type string.\n- toMap:\n[map1, map2, ...].toMap() -\u003e map\n- Converts a CEL list of CEL maps to a single CEL map\n- toCloudEventJsonWithPayloadFormat:\nmessage.toCloudEventJsonWithPayloadFormat() -\u003e map\n- Converts a message to the corresponding structure of JSON\nformat for CloudEvents.\n- It converts 'data' to destination payload format\nspecified in 'output_payload_format'. If 'output_payload_format' is\nnot set, the data will remain unchanged.\n- It also sets the corresponding datacontenttype of\nthe CloudEvent, as indicated by\n'output_payload_format'. If no\n'output_payload_format' is set it will use the value of the\n\"datacontenttype\" attribute on the CloudEvent if present, else\nremove \"datacontenttype\" attribute.\n- This function expects that the content of the message will\nadhere to the standard CloudEvent format. If it doesn't then this\nfunction will fail.\n- The result is a CEL map that corresponds to the JSON\nrepresentation of the CloudEvent. To convert that data to a JSON\nstring it can be chained with the toJsonString function.\n\nThe Pipeline expects that the message it receives adheres to the\nstandard CloudEvent format. If it doesn't then the outgoing message\nrequest may fail with a persistent error.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uri": {
                    "description": "The URI of the HTTP enpdoint.\n\nThe value must be a RFC2396 URI string.\nExamples: 'https://svc.us-central1.p.local:8080/route'.\nOnly the HTTPS protocol is supported.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Represents a HTTP endpoint destination.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "network_config": {
              "block": {
                "attributes": {
                  "network_attachment": {
                    "description": "Name of the NetworkAttachment that allows access to the consumer VPC.\nFormat:\n'projects/{PROJECT_ID}/regions/{REGION}/networkAttachments/{NETWORK_ATTACHMENT_NAME}'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Represents a network config to be used for destination resolution and\nconnectivity.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "output_payload_format": {
              "block": {
                "block_types": {
                  "avro": {
                    "block": {
                      "attributes": {
                        "schema_definition": {
                          "description": "The entire schema definition is stored in this field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The format of an AVRO message payload.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "json": {
                    "block": {
                      "description": "The format of a JSON message payload.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "protobuf": {
                    "block": {
                      "attributes": {
                        "schema_definition": {
                          "description": "The entire schema definition is stored in this field.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "The format of a Protobuf message payload.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Represents the format of message data.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of destinations to which messages will be forwarded. Currently,\nexactly one destination is supported per Pipeline.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "input_payload_format": {
        "block": {
          "block_types": {
            "avro": {
              "block": {
                "attributes": {
                  "schema_definition": {
                    "description": "The entire schema definition is stored in this field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The format of an AVRO message payload.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "json": {
              "block": {
                "description": "The format of a JSON message payload.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "protobuf": {
              "block": {
                "attributes": {
                  "schema_definition": {
                    "description": "The entire schema definition is stored in this field.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The format of a Protobuf message payload.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Represents the format of message data.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "logging_config": {
        "block": {
          "attributes": {
            "log_severity": {
              "computed": true,
              "description": "The minimum severity of logs that will be sent to Stackdriver/Platform\nTelemetry. Logs at severitiy ≥ this value will be sent, unless it is NONE. Possible values: [\"NONE\", \"DEBUG\", \"INFO\", \"NOTICE\", \"WARNING\", \"ERROR\", \"CRITICAL\", \"ALERT\", \"EMERGENCY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The configuration for Platform Telemetry logging for Eventarc Advanced\nresources.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "mediations": {
        "block": {
          "block_types": {
            "transformation": {
              "block": {
                "attributes": {
                  "transformation_template": {
                    "description": "The CEL expression template to apply to transform messages.\nThe following CEL extension functions are provided for\nuse in this CEL expression:\n- merge:\nmap1.merge(map2) -\u003e map3\n- Merges the passed CEL map with the existing CEL map the\nfunction is applied to.\n- If the same key exists in both maps, if the key's value is type\nmap both maps are merged else the value from the passed map is\nused.\n- denormalize:\nmap.denormalize() -\u003e map\n- Denormalizes a CEL map such that every value of type map or key\nin the map is expanded to return a single level map.\n- The resulting keys are \".\" separated indices of the map keys.\n- For example:\n{\n\"a\": 1,\n\"b\": {\n\"c\": 2,\n\"d\": 3\n}\n\"e\": [4, 5]\n}\n.denormalize()\n-\u003e {\n\"a\": 1,\n\"b.c\": 2,\n\"b.d\": 3,\n\"e.0\": 4,\n\"e.1\": 5\n}\n- setField:\nmap.setField(key, value) -\u003e message\n- Sets the field of the message with the given key to the\ngiven value.\n- If the field is not present it will be added.\n- If the field is present it will be overwritten.\n- The key can be a dot separated path to set a field in a nested\nmessage.\n- Key must be of type string.\n- Value may be any valid type.\n- removeFields:\nmap.removeFields([key1, key2, ...]) -\u003e message\n- Removes the fields of the map with the given keys.\n- The keys can be a dot separated path to remove a field in a\nnested message.\n- If a key is not found it will be ignored.\n- Keys must be of type string.\n- toMap:\n[map1, map2, ...].toMap() -\u003e map\n- Converts a CEL list of CEL maps to a single CEL map\n- toDestinationPayloadFormat():\nmessage.data.toDestinationPayloadFormat() -\u003e string or bytes\n- Converts the message data to the destination payload format\nspecified in Pipeline.Destination.output_payload_format\n- This function is meant to be applied to the message.data field.\n- If the destination payload format is not set, the function will\nreturn the message data unchanged.\n- toCloudEventJsonWithPayloadFormat:\nmessage.toCloudEventJsonWithPayloadFormat() -\u003e map\n- Converts a message to the corresponding structure of JSON\nformat for CloudEvents\n- This function applies toDestinationPayloadFormat() to the\nmessage data. It also sets the corresponding datacontenttype of\nthe CloudEvent, as indicated by\nPipeline.Destination.output_payload_format. If no\noutput_payload_format is set it will use the existing\ndatacontenttype on the CloudEvent if present, else leave\ndatacontenttype absent.\n- This function expects that the content of the message will\nadhere to the standard CloudEvent format. If it doesn't then this\nfunction will fail.\n- The result is a CEL map that corresponds to the JSON\nrepresentation of the CloudEvent. To convert that data to a JSON\nstring it can be chained with the toJsonString function.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Transformation defines the way to transform an incoming message.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "List of mediation operations to be performed on the message. Currently,\nonly one Transformation operation is allowed in each Pipeline.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "retry_policy": {
        "block": {
          "attributes": {
            "max_attempts": {
              "description": "The maximum number of delivery attempts for any message. The value must\nbe between 1 and 100.\nThe default value for this field is 5.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_retry_delay": {
              "description": "The maximum amount of seconds to wait between retry attempts. The value\nmust be between 1 and 600.\nThe default value for this field is 60.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_retry_delay": {
              "description": "The minimum amount of seconds to wait between retry attempts. The value\nmust be between 1 and 600.\nThe default value for this field is 5.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "The retry policy configuration for the Pipeline. The pipeline\nexponentially backs off in case the destination is non responsive or\nreturns a retryable error code. The default semantics are as follows:\nThe backoff starts with a 5 second delay and doubles the\ndelay after each failed attempt (10 seconds, 20 seconds, 40 seconds, etc.).\nThe delay is capped at 60 seconds by default.\nPlease note that if you set the min_retry_delay and max_retry_delay fields\nto the same value this will make the duration between retries constant.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleEventarcPipelineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEventarcPipeline), &result)
	return &result
}

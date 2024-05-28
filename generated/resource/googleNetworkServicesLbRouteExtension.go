package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesLbRouteExtension = `{
  "block": {
    "attributes": {
      "description": {
        "description": "A human-readable description of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "forwarding_rules": {
        "description": "A list of references to the forwarding rules to which this service extension is attached to.\nAt least one forwarding rule is required. There can be only one LbRouteExtension resource per forwarding rule.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Set of labels associated with the LbRouteExtension resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "load_balancing_scheme": {
        "description": "All backend services and forwarding rules referenced by this extension must share the same load balancing scheme.\nFor more information, refer to [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service) and\n[Supported application load balancers](https://cloud.google.com/service-extensions/docs/callouts-overview#supported-lbs). Possible values: [\"INTERNAL_MANAGED\", \"EXTERNAL_MANAGED\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the route extension",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the LbRouteExtension resource in the following format: projects/{project}/locations/{location}/lbRouteExtensions/{lbRouteExtension}",
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
      }
    },
    "block_types": {
      "extension_chains": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name for this extension chain. The name is logged as part of the HTTP request logs.\nThe name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,\nand can have a maximum length of 63 characters. Additionally, the first character must be a letter\nand the last character must be a letter or a number.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "extensions": {
              "block": {
                "attributes": {
                  "authority": {
                    "description": "The :authority header in the gRPC request sent from Envoy to the extension service.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "fail_open": {
                    "description": "Determines how the proxy behaves if the call to the extension fails or times out.\nWhen set to TRUE, request or response processing continues without error.\nAny subsequent extensions in the extension chain are also executed.\nWhen set to FALSE: * If response headers have not been delivered to the downstream client,\na generic 500 error is returned to the client. The error response can be tailored by\nconfiguring a custom error response in the load balancer.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "forward_headers": {
                    "description": "List of the HTTP headers to forward to the extension (from the client or backend).\nIf omitted, all headers are sent. Each element is a string indicating the header name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "name": {
                    "description": "The name for this extension. The name is logged as part of the HTTP request logs.\nThe name must conform with RFC-1034, is restricted to lower-cased letters, numbers and hyphens,\nand can have a maximum length of 63 characters. Additionally, the first character must be a letter\nand the last a letter or a number.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "service": {
                    "description": "The reference to the service that runs the extension. Must be a reference to a backend service",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "timeout": {
                    "description": "Specifies the timeout for each individual message on the stream. The timeout must be between 10-1000 milliseconds.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "A set of extensions to execute for the matching request.\nAt least one extension is required. Up to 3 extensions can be defined for each extension chain for\nLbTrafficExtension resource. LbRouteExtension chains are limited to 1 extension per extension chain.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "match_condition": {
              "block": {
                "attributes": {
                  "cel_expression": {
                    "description": "A Common Expression Language (CEL) expression that is used to match requests for which the extension chain is executed.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Conditions under which this chain is invoked for a request.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "A set of ordered extension chains that contain the match conditions and extensions to execute.\nMatch conditions for each extension chain are evaluated in sequence for a given request.\nThe first extension chain that has a condition that matches the request is executed.\nAny subsequent extension chains do not execute. Limited to 5 extension chains per resource.",
          "description_kind": "plain"
        },
        "min_items": 1,
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

func GoogleNetworkServicesLbRouteExtensionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesLbRouteExtension), &result)
	return &result
}

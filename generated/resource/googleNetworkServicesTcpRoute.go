package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetworkServicesTcpRoute = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Time the TcpRoute was created in UTC.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A free-text description of the resource. Max length 1024 characters.",
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
      "gateways": {
        "description": "Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.\nEach gateway reference should match the pattern: projects/*/locations/global/gateways/\u003cgateway_name\u003e",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Set of label tags associated with the TcpRoute resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "meshes": {
        "description": "Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.\nEach mesh reference should match the pattern: projects/*/locations/global/meshes/\u003cmesh_name\u003e\nThe attached Mesh should be of a type SIDECAR",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "description": "Name of the TcpRoute resource.",
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
      "self_link": {
        "computed": true,
        "description": "Server-defined URL of this resource.",
        "description_kind": "plain",
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
      "update_time": {
        "computed": true,
        "description": "Time the TcpRoute was updated in UTC.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "rules": {
        "block": {
          "block_types": {
            "action": {
              "block": {
                "attributes": {
                  "idle_timeout": {
                    "description": "Specifies the idle timeout for the selected route. The idle timeout is defined as the period in which there are no bytes sent or received on either the upstream or downstream connection. If not set, the default idle timeout is 30 seconds. If set to 0s, the timeout will be disabled.\n\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "original_destination": {
                    "description": "If true, Router will use the destination IP and port of the original connection as the destination of the request.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "destinations": {
                    "block": {
                      "attributes": {
                        "service_name": {
                          "description": "The URL of a BackendService to route traffic to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "weight": {
                          "description": "Specifies the proportion of requests forwarded to the backend referenced by the serviceName field. This is computed as: weight/Sum(weights in this destination list). For non-zero values, there may be some epsilon from the exact proportion defined here depending on the precision an implementation supports.\nIf only one serviceName is specified and it has a weight greater than 0, 100% of the traffic is forwarded to that backend.\nIf weights are specified for any one service name, they need to be specified for all of them.\nIf weights are unspecified for all services, then, traffic is distributed in equal proportions to all of them.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The destination services to which traffic should be forwarded. At least one destination service is required.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "A detailed rule defining how to route traffic.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "matches": {
              "block": {
                "attributes": {
                  "address": {
                    "description": "Must be specified in the CIDR range format. A CIDR range consists of an IP Address and a prefix length to construct the subnet mask.\nBy default, the prefix length is 32 (i.e. matches a single IP address). Only IPV4 addresses are supported. Examples: \"10.0.0.1\" - matches against this exact IP address. \"10.0.0.0/8\" - matches against any IP address within the 10.0.0.0 subnet and 255.255.255.0 mask. \"0.0.0.0/0\" - matches against any IP address'.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "port": {
                    "description": "Specifies the destination port to match against.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "RouteMatch defines the predicate used to match requests to a given action. Multiple match types are \"OR\"ed for evaluation.\nIf no routeMatch field is specified, this rule will unconditionally match traffic.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Rules that define how traffic is routed and handled. At least one RouteRule must be supplied.\nIf there are multiple rules then the action taken will be the first rule to match.",
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

func GoogleNetworkServicesTcpRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetworkServicesTcpRoute), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvironment = `{
  "block": {
    "attributes": {
      "api_proxy_type": {
        "computed": true,
        "description": "Optional. API Proxy type supported by the environment. The type can be set when creating\nthe Environment and cannot be changed. Possible values: [\"API_PROXY_TYPE_UNSPECIFIED\", \"PROGRAMMABLE\", \"CONFIGURABLE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_type": {
        "computed": true,
        "description": "Optional. Deployment type supported by the environment. The deployment type can be\nset when creating the environment and cannot be changed. When you enable archive\ndeployment, you will be prevented from performing a subset of actions within the\nenvironment, including:\nManaging the deployment of API proxy or shared flow revisions;\nCreating, updating, or deleting resource files;\nCreating, updating, or deleting target servers. Possible values: [\"DEPLOYMENT_TYPE_UNSPECIFIED\", \"PROXY\", \"ARCHIVE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "Description of the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "Display name of the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "forward_proxy_uri": {
        "description": "Optional. URI of the forward proxy to be applied to the runtime instances in this environment. Must be in the format of {scheme}://{hostname}:{port}. Note that the scheme must be one of \"http\" or \"https\", and the port must be supplied.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource ID of the environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee environment,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "Types that can be selected for an Environment. Each of the types are\nlimited by capability and capacity. Refer to Apigee's public documentation\nto understand about each of these types in details.\nAn Apigee org can support heterogeneous Environments. Possible values: [\"ENVIRONMENT_TYPE_UNSPECIFIED\", \"BASE\", \"INTERMEDIATE\", \"COMPREHENSIVE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "client_ip_resolution_config": {
        "block": {
          "block_types": {
            "header_index_algorithm": {
              "block": {
                "attributes": {
                  "ip_header_index": {
                    "description": "The index of the ip in the header. Positive indices 0, 1, 2, 3 chooses indices from the left (first ips). Negative indices -1, -2, -3 chooses indices from the right (last ips).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "ip_header_name": {
                    "description": "The name of the header to extract the client ip from. We are currently only supporting the X-Forwarded-For header.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Resolves the client ip based on a custom header.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The algorithm to resolve IP. This will affect Analytics, API Security, and other features that use the client ip. To remove a client ip resolution config, update the field to an empty value. Example: '{ \"clientIpResolutionConfig\" = {} }' For more information, see: https://cloud.google.com/apigee/docs/api-platform/system-administration/client-ip-resolution",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "current_aggregate_node_count": {
              "computed": true,
              "description": "The current total number of gateway nodes that each environment currently has across\nall instances.",
              "description_kind": "plain",
              "type": "string"
            },
            "max_node_count": {
              "description": "The maximum total number of gateway nodes that the is reserved for all instances that\nhas the specified environment. If not specified, the default is determined by the\nrecommended maximum number of nodes for that gateway.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_node_count": {
              "description": "The minimum total number of gateway nodes that the is reserved for all instances that\nhas the specified environment. If not specified, the default is determined by the\nrecommended minimum number of nodes for that gateway.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "NodeConfig for setting the min/max number of nodes associated with the environment.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "properties": {
        "block": {
          "block_types": {
            "property": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The property key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The property value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "List of all properties in the object.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Key-value pairs that may be used for customizing the environment.",
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

func GoogleApigeeEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvironment), &result)
	return &result
}

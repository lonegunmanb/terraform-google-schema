package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBeyondcorpApplication = `{
  "block": {
    "attributes": {
      "application_id": {
        "description": "Optional. User-settable Application resource ID.\n* Must start with a letter.\n* Must contain between 4-63 characters from '/a-z-/'.\n* Must end with a number or letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "Optional. An arbitrary user-provided name for the Application resource.\nCannot exceed 64 characters.",
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
        "computed": true,
        "description": "Identifier. Name of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_gateways_id": {
        "description": "Part of 'parent'. See documentation of 'projectsId'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Timestamp when the resource was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "endpoint_matchers": {
        "block": {
          "attributes": {
            "hostname": {
              "description": "Required. Hostname of the application.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ports": {
              "description": "Optional. Ports of the application.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "number"
              ]
            }
          },
          "description": "Required. Endpoint matchers associated with an application.\nA combination of hostname and ports as endpoint matcher is used to match\nthe application.\nMatch conditions for OR logic.\nAn array of match conditions to allow for multiple matching criteria.\nThe rule is considered a match if one the conditions are met.\nThe conditions can be one of the following combination\n(Hostname), (Hostname \u0026 Ports)\n\nEXAMPLES:\nHostname - (\"*.abc.com\"), (\"xyz.abc.com\")\nHostname and Ports - (\"abc.com\" and \"22\"), (\"abc.com\" and \"22,33\") etc",
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
      },
      "upstreams": {
        "block": {
          "block_types": {
            "egress_policy": {
              "block": {
                "attributes": {
                  "regions": {
                    "description": "Required. List of regions where the application sends traffic to.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Optional. Routing policy information.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "network": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "Required. Network name is of the format:\n'projects/{project}/global/networks/{network}'",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "Network to forward traffic to.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Optional. List of which upstream resource(s) to forward traffic to.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      }
    },
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBeyondcorpApplicationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBeyondcorpApplication), &result)
	return &result
}

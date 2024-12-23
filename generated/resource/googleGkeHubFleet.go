package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeHubFleet = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time the fleet was created, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "The time the fleet was deleted, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "A user-assigned display name of the Fleet. When present, it must be between 4 to 30 characters.\nAllowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point.",
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
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the fleet resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "code": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "Google-generated UUID for this resource. This is unique across all\nFleet resources. If a Fleet resource is deleted and another\nresource with the same name is created, it gets a different uid.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time the fleet was last updated, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "default_cluster_config": {
        "block": {
          "block_types": {
            "binary_authorization_config": {
              "block": {
                "attributes": {
                  "evaluation_mode": {
                    "description": "Mode of operation for binauthz policy evaluation. Possible values: [\"DISABLED\", \"POLICY_BINDINGS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "policy_bindings": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The relative resource name of the binauthz platform policy to audit. GKE\nplatform policies have the following format:\n'projects/{project_number}/platforms/gke/policies/{policy_id}'.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Binauthz policies that apply to this cluster.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Enable/Disable binary authorization features for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "security_posture_config": {
              "block": {
                "attributes": {
                  "mode": {
                    "description": "Sets which mode to use for Security Posture features. Possible values: [\"DISABLED\", \"BASIC\", \"ENTERPRISE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "vulnerability_mode": {
                    "description": "Sets which mode to use for vulnerability scanning. Possible values: [\"VULNERABILITY_DISABLED\", \"VULNERABILITY_BASIC\", \"VULNERABILITY_ENTERPRISE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Enable/Disable Security Posture features for the cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The default cluster configurations to apply across the fleet.",
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

func GoogleGkeHubFleetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeHubFleet), &result)
	return &result
}

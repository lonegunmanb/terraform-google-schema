package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvironmentAddonsConfig = `{
  "block": {
    "attributes": {
      "analytics_enabled": {
        "description": "Flag to enable/disable Analytics.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "env_id": {
        "description": "The Apigee environment group associated with the Apigee environment,\nin the format 'organizations/{{org_name}}/environments/{{env_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
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

func GoogleApigeeEnvironmentAddonsConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvironmentAddonsConfig), &result)
	return &result
}

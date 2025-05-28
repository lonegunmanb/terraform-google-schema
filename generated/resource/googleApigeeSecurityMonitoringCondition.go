package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeSecurityMonitoringCondition = `{
  "block": {
    "attributes": {
      "condition_id": {
        "description": "Resource ID of the security monitoring condition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp at which this profile was created.",
        "description_kind": "plain",
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
        "description": "Name of the security monitoring condition resource,\nin the format 'organizations/{{org_name}}/securityMonitoringConditions/{{condition_id}}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee Security Monitoring Condition,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile": {
        "description": "ID of security profile of the security monitoring condition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "description": "ID of security profile of the security monitoring condition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "total_deployed_resources": {
        "computed": true,
        "description": "Total number of deployed resources within scope.",
        "description_kind": "plain",
        "type": "number"
      },
      "total_monitored_resources": {
        "computed": true,
        "description": "Total number of monitored resources within this condition.",
        "description_kind": "plain",
        "type": "number"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp at which this profile was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "include_all_resources": {
        "block": {
          "description": "A nested object resource.",
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

func GoogleApigeeSecurityMonitoringConditionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeSecurityMonitoringCondition), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeSecurityProfileV2 = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The timestamp at which this profile was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Description of the security profile.",
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
        "description": "Name of the security profile v2 resource,\nin the format 'organizations/{{org_name}}/securityProfilesV2/{{profile_id}}'.",
        "description_kind": "plain",
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee Security Profile V2,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_id": {
        "description": "Resource ID of the security profile.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp at which this profile was most recently updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "profile_assessment_configs": {
        "block": {
          "attributes": {
            "assessment": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "weight": {
              "description": "The weight of the assessment. Possible values: [\"MINOR\", \"MODERATE\", \"MAJOR\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A map of the assessment name and the assessment config.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func GoogleApigeeSecurityProfileV2Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeSecurityProfileV2), &result)
	return &result
}

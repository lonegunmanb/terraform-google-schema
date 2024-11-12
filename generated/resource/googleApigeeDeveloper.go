package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeDeveloper = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "Time at which the developer was created in milliseconds since epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "email": {
        "description": "Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only..",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "first_name": {
        "description": "First name of the developer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_modified_at": {
        "computed": true,
        "description": "Time at which the developer was last modified in milliseconds since epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_name": {
        "description": "Last name of the developer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee instance,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organizatio_name": {
        "computed": true,
        "description": "Name of the Apigee organization in which the developer resides.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Status of the developer. Valid values are active and inactive.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "description": "User name of the developer. Not used by Apigee hybrid.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "attributes": {
        "block": {
          "attributes": {
            "name": {
              "description": "Key of the attribute",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description": "Value of the attribute",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Developer attributes (name/value pairs). The custom attribute limit is 18.",
          "description_kind": "plain"
        },
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

func GoogleApigeeDeveloperSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeDeveloper), &result)
	return &result
}

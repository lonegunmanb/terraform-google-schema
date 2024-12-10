package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeAppGroup = `{
  "block": {
    "attributes": {
      "app_group_id": {
        "computed": true,
        "description": "Internal identifier that cannot be edited",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_id": {
        "description": "Channel identifier identifies the owner maintaining this grouping.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "channel_uri": {
        "description": "A reference to the associated storefront/marketplace.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "Created time as milliseconds since epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "App group name displayed in the UI",
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
      "last_modified_at": {
        "computed": true,
        "description": "Modified time as milliseconds since epoch.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._-$ %.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "org_id": {
        "description": "The Apigee Organization associated with the Apigee app group,\nin the format 'organizations/{{org_name}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "organization": {
        "computed": true,
        "description": "App group name displayed in the UI",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "description": "Valid values are active or inactive. Note that the status of the AppGroup should be updated via UpdateAppGroupRequest by setting the action as active or inactive. Possible values: [\"active\", \"inactive\"]",
        "description_kind": "plain",
        "optional": true,
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
          "description": "A list of attributes",
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

func GoogleApigeeAppGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeAppGroup), &result)
	return &result
}

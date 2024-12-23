package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleLoggingLogScope = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. The creation timestamp of the log scopes.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "Describes this log scopes.",
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
      "location": {
        "computed": true,
        "description": "The location of the resource. The only supported location is global so far.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the log scope. For example: \\'projects/my-project/locations/global/logScopes/my-log-scope\\'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "computed": true,
        "description": "The parent of the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_names": {
        "description": "Names of one or more parent resources : *  \\'projects/[PROJECT_ID]\\' May alternatively be one or more views : * \\'projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/views/[VIEW_ID]\\' A log scope can include a maximum of 50 projects and a maximum of 100 resources in total.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "update_time": {
        "computed": true,
        "description": "Output only. The last update timestamp of the log scopes.",
        "description_kind": "plain",
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

func GoogleLoggingLogScopeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleLoggingLogScope), &result)
	return &result
}

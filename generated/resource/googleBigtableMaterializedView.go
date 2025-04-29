package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigtableMaterializedView = `{
  "block": {
    "attributes": {
      "deletion_protection": {
        "description": "Set to true to make the MaterializedView protected against deletion.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance": {
        "description": "The name of the instance to create the materialized view within.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "materialized_view_id": {
        "description": "The unique name of the materialized view in the form '[_a-zA-Z0-9][-_.a-zA-Z0-9]*'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The unique name of the requested materialized view. Values are of the form 'projects/\u003cproject\u003e/instances/\u003cinstance\u003e/materializedViews/\u003cmaterializedViewId\u003e'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query": {
        "description": "The materialized view's select query.",
        "description_kind": "plain",
        "required": true,
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

func GoogleBigtableMaterializedViewSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigtableMaterializedView), &result)
	return &result
}

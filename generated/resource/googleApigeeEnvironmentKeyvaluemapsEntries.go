package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApigeeEnvironmentKeyvaluemapsEntries = `{
  "block": {
    "attributes": {
      "env_keyvaluemap_id": {
        "description": "The Apigee environment keyvalumaps Id associated with the Apigee environment,\nin the format 'organizations/{{org_name}}/environments/{{env_name}}/keyvaluemaps/{{keyvaluemap_name}}'.",
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
      "name": {
        "description": "Required. Resource URI that can be used to identify the scope of the key value map entries.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "value": {
        "description": "Required. Data or payload that is being retrieved and associated with the unique key.",
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

func GoogleApigeeEnvironmentKeyvaluemapsEntriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApigeeEnvironmentKeyvaluemapsEntries), &result)
	return &result
}

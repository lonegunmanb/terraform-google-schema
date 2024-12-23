package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleTagsTagKeys = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "keys": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "description": "string",
              "name": "string",
              "namespaced_name": "string",
              "parent": "string",
              "purpose": "string",
              "purpose_data": [
                "map",
                "string"
              ],
              "short_name": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "parent": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleTagsTagKeysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleTagsTagKeys), &result)
	return &result
}

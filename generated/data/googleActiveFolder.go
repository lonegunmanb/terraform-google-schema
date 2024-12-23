package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleActiveFolder = `{
  "block": {
    "attributes": {
      "api_method": {
        "description": "Provides the REST method through which to find the folder. LIST is recommended as it is strongly consistent.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func GoogleActiveFolderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleActiveFolder), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleArtifactRegistryRepositories = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name_filter": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "repositories": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "description": "string",
              "format": "string",
              "id": "string",
              "repository_id": "string",
              "update_time": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleArtifactRegistryRepositoriesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleArtifactRegistryRepositories), &result)
	return &result
}

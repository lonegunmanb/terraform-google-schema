package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComposerUserWorkloadsSecret = `{
  "block": {
    "attributes": {
      "data": {
        "computed": true,
        "description": "A map of the secret data.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "environment": {
        "description": "Name of the environment.",
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
        "description": "Name of the secret.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "description": "The location or Compute Engine region for the environment.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComposerUserWorkloadsSecretSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComposerUserWorkloadsSecret), &result)
	return &result
}

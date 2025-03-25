package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeImages = `{
  "block": {
    "attributes": {
      "filter": {
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
      "images": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "archive_size_bytes": "number",
              "creation_timestamp": "string",
              "description": "string",
              "disk_size_gb": "number",
              "family": "string",
              "image_id": "number",
              "labels": [
                "map",
                "string"
              ],
              "name": "string",
              "self_link": "string",
              "source_disk": "string",
              "source_disk_id": "string",
              "source_image_id": "string"
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeImagesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeImages), &result)
	return &result
}

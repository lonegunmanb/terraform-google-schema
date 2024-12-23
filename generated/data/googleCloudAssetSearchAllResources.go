package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudAssetSearchAllResources = `{
  "block": {
    "attributes": {
      "asset_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "results": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "asset_type": "string",
              "create_time": "string",
              "description": "string",
              "display_name": "string",
              "folders": [
                "list",
                "string"
              ],
              "kms_keys": [
                "list",
                "string"
              ],
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "name": "string",
              "network_tags": [
                "list",
                "string"
              ],
              "organization": "string",
              "parent_asset_type": "string",
              "parent_full_resource_name": "string",
              "project": "string",
              "state": "string",
              "update_time": "string"
            }
          ]
        ]
      },
      "scope": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudAssetSearchAllResourcesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudAssetSearchAllResources), &result)
	return &result
}

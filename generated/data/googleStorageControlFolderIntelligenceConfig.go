package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageControlFolderIntelligenceConfig = `{
  "block": {
    "attributes": {
      "edition_config": {
        "computed": true,
        "description": "Edition configuration of the Storage Intelligence resource. Valid values are INHERIT, TRIAL, DISABLED and STANDARD.",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_intelligence_config": {
        "computed": true,
        "description": "The Intelligence config that is effective for the resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "effective_edition": "string",
              "intelligence_config": "string"
            }
          ]
        ]
      },
      "filter": {
        "computed": true,
        "description": "Filter over location and bucket using include or exclude semantics. Resources that match the include or exclude filter are exclusively included or excluded from the Storage Intelligence plan.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "excluded_cloud_storage_buckets": [
                "list",
                [
                  "object",
                  {
                    "bucket_id_regexes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "excluded_cloud_storage_locations": [
                "list",
                [
                  "object",
                  {
                    "locations": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "included_cloud_storage_buckets": [
                "list",
                [
                  "object",
                  {
                    "bucket_id_regexes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "included_cloud_storage_locations": [
                "list",
                [
                  "object",
                  {
                    "locations": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Identifier of the GCP Folder. For GCP Folder, this field can be folder number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which the Storage Intelligence Config resource is last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleStorageControlFolderIntelligenceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageControlFolderIntelligenceConfig), &result)
	return &result
}

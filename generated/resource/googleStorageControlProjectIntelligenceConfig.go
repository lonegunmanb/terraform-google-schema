package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleStorageControlProjectIntelligenceConfig = `{
  "block": {
    "attributes": {
      "edition_config": {
        "computed": true,
        "description": "Edition configuration of the Storage Intelligence resource. Valid values are INHERIT, TRIAL, DISABLED and STANDARD.",
        "description_kind": "plain",
        "optional": true,
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Identifier of the GCP project. For GCP project, this field can be project name or project number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "trial_config": {
        "computed": true,
        "description": "The trial configuration of the Storage Intelligence resource.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "expire_time": "string"
            }
          ]
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The time at which the Storage Intelligence Config resource is last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "block_types": {
            "excluded_cloud_storage_buckets": {
              "block": {
                "attributes": {
                  "bucket_id_regexes": {
                    "description": "List of bucket id regexes to exclude in the storage intelligence plan.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Buckets to exclude from the Storage Intelligence plan.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "excluded_cloud_storage_locations": {
              "block": {
                "attributes": {
                  "locations": {
                    "description": "List of locations.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Locations to exclude from the Storage Intelligence plan.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "included_cloud_storage_buckets": {
              "block": {
                "attributes": {
                  "bucket_id_regexes": {
                    "description": "List of bucket id regexes to exclude in the storage intelligence plan.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Buckets to include in the Storage Intelligence plan.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "included_cloud_storage_locations": {
              "block": {
                "attributes": {
                  "locations": {
                    "description": "List of locations.",
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Locations to include in the Storage Intelligence plan.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Filter over location and bucket using include or exclude semantics. Resources that match the include or exclude filter are exclusively included or excluded from the Storage Intelligence plan.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func GoogleStorageControlProjectIntelligenceConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleStorageControlProjectIntelligenceConfig), &result)
	return &result
}

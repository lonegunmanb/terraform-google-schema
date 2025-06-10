package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeSnapshotSettings = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "storage_location": {
        "block": {
          "attributes": {
            "policy": {
              "description": "The chosen location policy Possible values: [\"NEAREST_MULTI_REGION\", \"LOCAL_REGION\", \"SPECIFIC_LOCATIONS\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "locations": {
              "block": {
                "attributes": {
                  "location": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Name of the location. It should be one of the Cloud Storage buckets.\nOnly one location can be specified. (should match location)",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "When the policy is SPECIFIC_LOCATIONS, snapshots will be stored in the\nlocations listed in this field. Keys are Cloud Storage bucket locations.\nOnly one location can be specified.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "Policy of which storage location is going to be resolved, and additional data\nthat particularizes how the policy is going to be carried out",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleComputeSnapshotSettingsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeSnapshotSettings), &result)
	return &result
}

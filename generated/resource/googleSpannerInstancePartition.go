package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSpannerInstancePartition = `{
  "block": {
    "attributes": {
      "config": {
        "description": "The name of the instance partition's configuration (similar to a region) which\ndefines the geographic placement and replication of data in this instance partition.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "description": "The descriptive name for this instance partition as it appears in UIs.\nMust be unique per project and between 4 and 30 characters in length.",
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
      "instance": {
        "description": "The instance to create the instance partition in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "A unique identifier for the instance partition, which cannot be changed after\nthe instance partition is created. The name must be between 2 and 64 characters\nand match the regular expression [a-z][a-z0-9\\\\-]{0,61}[a-z0-9].",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_count": {
        "description": "The number of nodes allocated to this instance partition. One node equals\n1000 processing units. Exactly one of either node_count or processing_units\nmust be present.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "processing_units": {
        "description": "The number of processing units allocated to this instance partition.\nExactly one of either node_count or processing_units must be present.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current instance partition state. Possible values are:\nCREATING: The instance partition is being created. Resources are being\nallocated for the instance partition.\nREADY: The instance partition has been allocated resources and is ready for use.",
        "description_kind": "plain",
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

func GoogleSpannerInstancePartitionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSpannerInstancePartition), &result)
	return &result
}

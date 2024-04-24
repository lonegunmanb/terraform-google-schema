package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeMachineTypes = `{
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
      "machine_types": {
        "computed": true,
        "description": "The list of machine types",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "accelerators": [
                "list",
                [
                  "object",
                  {
                    "guest_accelerator_count": "number",
                    "guest_accelerator_type": "string"
                  }
                ]
              ],
              "deprecated": [
                "set",
                [
                  "object",
                  {
                    "replacement": "string",
                    "state": "string"
                  }
                ]
              ],
              "description": "string",
              "guest_cpus": "number",
              "is_shared_cpus": "bool",
              "maximum_persistent_disks": "number",
              "maximum_persistent_disks_size_gb": "number",
              "memory_mb": "number",
              "name": "string",
              "self_link": "string"
            }
          ]
        ]
      },
      "project": {
        "computed": true,
        "description": "Project ID for this request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "zone": {
        "computed": true,
        "description": "The name of the zone for this request.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleComputeMachineTypesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeMachineTypes), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApphubDiscoveredWorkload = `{
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
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workload_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "gcp_project": "string",
              "location": "string",
              "zone": "string"
            }
          ]
        ]
      },
      "workload_reference": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "uri": "string"
            }
          ]
        ]
      },
      "workload_uri": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleApphubDiscoveredWorkloadSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApphubDiscoveredWorkload), &result)
	return &result
}

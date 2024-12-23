package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleApphubDiscoveredService = `{
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
      "service_properties": {
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
      "service_reference": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "path": "string",
              "uri": "string"
            }
          ]
        ]
      },
      "service_uri": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleApphubDiscoveredServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleApphubDiscoveredService), &result)
	return &result
}

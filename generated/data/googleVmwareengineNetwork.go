package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineNetwork = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "User-provided description for this VMware Engine network.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location where the VMwareEngineNetwork should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The ID of the VMwareEngineNetwork.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the VMware Engine network.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "VMware Engine network type. Possible values: [\"LEGACY\", \"STANDARD\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpc_networks": {
        "computed": true,
        "description": "VMware Engine service VPC networks that provide connectivity from a private cloud to customer projects,\nthe internet, and other Google Cloud services.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "network": "string",
              "type": "string"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleVmwareengineNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineNetwork), &result)
	return &result
}

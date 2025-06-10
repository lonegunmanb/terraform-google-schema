package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBeyondcorpSecurityGateway = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Output only. Timestamp when the resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "delegating_service_account": {
        "computed": true,
        "description": "Service account used for operations that involve resources in consumer projects.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "Optional. An arbitrary user-provided name for the SecurityGateway.\nCannot exceed 64 characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "external_ips": {
        "computed": true,
        "description": "Output only. IP addresses that will be used for establishing\nconnection to the endpoints.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "hubs": {
        "computed": true,
        "description": "Optional. Map of Hubs that represents regional data path deployment with GCP region\nas a key.",
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "internet_gateway": [
                "list",
                [
                  "object",
                  {
                    "assigned_ips": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "region": "string"
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
      "location": {
        "computed": true,
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122. Must be omitted or set to 'global'.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. Name of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "security_gateway_id": {
        "description": "Optional. User-settable SecurityGateway resource ID.\n* Must start with a letter.\n* Must contain between 4-63 characters from '/a-z-/'.\n* Must end with a number or letter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "Output only. The operational state of the SecurityGateway.\nPossible values:\nSTATE_UNSPECIFIED\nCREATING\nUPDATING\nDELETING\nRUNNING\nDOWN\nERROR",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Timestamp when the resource was last modified.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBeyondcorpSecurityGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBeyondcorpSecurityGateway), &result)
	return &result
}

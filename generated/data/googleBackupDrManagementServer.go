package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrManagementServer = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the management server (management console)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "management_uri": {
        "computed": true,
        "description": "The management console URI",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "api": "string",
              "web_ui": "string"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of management server (management console)",
        "description_kind": "plain",
        "type": "string"
      },
      "networks": {
        "computed": true,
        "description": "Network details to create management server (management console).",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "network": "string",
              "peering_mode": "string"
            }
          ]
        ]
      },
      "oauth2_client_id": {
        "computed": true,
        "description": "The oauth2ClientId of management console.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of management server (management console). Default value: \"BACKUP_RESTORE\" Possible values: [\"BACKUP_RESTORE\"]",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBackupDrManagementServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrManagementServer), &result)
	return &result
}

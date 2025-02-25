package resource

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
        "description": "The name of management server (management console)",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of management server (management console). Default value: \"BACKUP_RESTORE\" Possible values: [\"BACKUP_RESTORE\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "networks": {
        "block": {
          "attributes": {
            "network": {
              "description": "Network with format 'projects/{{project_id}}/global/networks/{{network_id}}'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "peering_mode": {
              "description": "Type of Network peeringMode Default value: \"PRIVATE_SERVICE_ACCESS\" Possible values: [\"PRIVATE_SERVICE_ACCESS\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Network details to create management server (management console).",
          "description_kind": "plain"
        },
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

func GoogleBackupDrManagementServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrManagementServer), &result)
	return &result
}

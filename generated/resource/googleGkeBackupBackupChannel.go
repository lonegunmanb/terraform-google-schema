package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeBackupBackupChannel = `{
  "block": {
    "attributes": {
      "description": {
        "description": "User specified descriptive string for this BackupChannel.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_project": {
        "description": "The project where Backups are allowed to be stored.\nThe format is 'projects/{project}'.\n{project} can be project number or project id.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "destination_project_id": {
        "computed": true,
        "description": "The project_id where Backups are allowed to be stored.\nExample Project ID: \"my-project-id\".",
        "description_kind": "plain",
        "type": "string"
      },
      "effective_labels": {
        "computed": true,
        "description": "All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "etag": {
        "computed": true,
        "description": "etag is used for optimistic concurrency control as a way to help prevent simultaneous\nupdates of a backup channel from overwriting each other. It is strongly suggested that\nsystems make use of the 'etag' in the read-modify-write cycle to perform BackupChannel updates\nin order to avoid race conditions: An etag is returned in the response to backupChannels.get,\nand systems are expected to put that etag in the request to backupChannels.patch or\nbackupChannels.delete to ensure that their change will be applied to the same version of the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Description: A set of custom labels supplied by the user.\nA list of key-\u003evalue pairs.\nExample: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The region of the Backup Channel.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The full name of the BackupChannel Resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "terraform_labels": {
        "computed": true,
        "description": "The combination of labels configured directly on the resource\n and default labels configured on the provider.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "uid": {
        "computed": true,
        "description": "Server generated, unique identifier of UUID format.",
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

func GoogleGkeBackupBackupChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeBackupBackupChannel), &result)
	return &result
}

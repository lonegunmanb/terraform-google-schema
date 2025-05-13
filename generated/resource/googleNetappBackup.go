package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetappBackup = `{
  "block": {
    "attributes": {
      "backup_region": {
        "computed": true,
        "description": "Region in which backup is stored.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_type": {
        "computed": true,
        "description": "Type of backup, manually created or created by a backup policy. Possible Values : [TYPE_UNSPECIFIED, MANUAL, SCHEDULED]",
        "description_kind": "plain",
        "type": "string"
      },
      "chain_storage_bytes": {
        "computed": true,
        "description": "Backups of a volume build incrementally on top of each other. They form a \"backup chain\".\nTotal size of all backups in a chain in bytes = baseline backup size + sum(incremental backup size)",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Create time of the backup. A timestamp in RFC3339 UTC \"Zulu\" format. Examples: \"2023-06-22T09:13:01.617Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.",
        "description_kind": "plain",
        "optional": true,
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels as key value pairs. Example: '{ \"owner\": \"Bob\", \"department\": \"finance\", \"purpose\": \"testing\" }'.\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of the backup.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the backup. Needs to be unique per location.",
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
      "source_snapshot": {
        "description": "If specified, backup will be created from the given snapshot. If not specified,\nthere will be a new snapshot taken to initiate the backup creation.\nFormat: 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}''",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "source_volume": {
        "description": "ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}''",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the Backup Vault. Possible Values : [STATE_UNSPECIFIED, CREATING, UPLOADING, READY, DELETING, ERROR, UPDATING]",
        "description_kind": "plain",
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
      "vault_name": {
        "description": "Name of the backup vault to store the backup in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "volume_region": {
        "computed": true,
        "description": "Region of the volume from which the backup was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "volume_usage_bytes": {
        "computed": true,
        "description": "Size of the file system when the backup was created. When creating a new volume from the backup, the volume capacity will have to be at least as big.",
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

func GoogleNetappBackupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetappBackup), &result)
	return &result
}

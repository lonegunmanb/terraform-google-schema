package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleNetappBackupVault = `{
  "block": {
    "attributes": {
      "backup_region": {
        "description": "Region in which backup is stored.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_vault_type": {
        "computed": true,
        "description": "Type of the backup vault to be created. Default is IN_REGION. Possible values: [\"BACKUP_VAULT_TYPE_UNSPECIFIED\", \"IN_REGION\", \"CROSS_REGION\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Create time of the backup vault. A timestamp in RFC3339 UTC \"Zulu\" format. Examples: \"2023-06-22T09:13:01.617Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_backup_vault": {
        "computed": true,
        "description": "Name of the Backup vault created in backup region.",
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
        "description": "Location (region) of the backup vault.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the backup vault. Needs to be unique per location.",
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
      "source_backup_vault": {
        "computed": true,
        "description": "Name of the Backup vault created in source region.",
        "description_kind": "plain",
        "type": "string"
      },
      "source_region": {
        "computed": true,
        "description": "Region in which the backup vault is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the Backup Vault.",
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
      }
    },
    "block_types": {
      "backup_retention_policy": {
        "block": {
          "attributes": {
            "backup_minimum_enforced_retention_days": {
              "description": "Minimum retention duration in days for backups in the backup vault.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "daily_backup_immutable": {
              "description": "Indicates if the daily backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "manual_backup_immutable": {
              "description": "Indicates if the manual backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "monthly_backup_immutable": {
              "description": "Indicates if the monthly backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "weekly_backup_immutable": {
              "description": "Indicates if the weekly backups are immutable. At least one of daily_backup_immutable, weekly_backup_immutable, monthly_backup_immutable and manual_backup_immutable must be true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Backup retention policy defining the retention of the backups.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleNetappBackupVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleNetappBackupVault), &result)
	return &result
}

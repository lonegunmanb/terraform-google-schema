package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrBackupPlan = `{
  "block": {
    "attributes": {
      "backup_plan_id": {
        "description": "The ID of the backup plan",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backup_rules": {
        "computed": true,
        "description": "The backup rules for this 'BackupPlan'. There must be at least one 'BackupRule' message.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_retention_days": "number",
              "rule_id": "string",
              "standard_schedule": [
                "list",
                [
                  "object",
                  {
                    "backup_window": [
                      "list",
                      [
                        "object",
                        {
                          "end_hour_of_day": "number",
                          "start_hour_of_day": "number"
                        }
                      ]
                    ],
                    "days_of_month": [
                      "list",
                      "number"
                    ],
                    "days_of_week": [
                      "list",
                      "string"
                    ],
                    "hourly_frequency": "number",
                    "months": [
                      "list",
                      "string"
                    ],
                    "recurrence_type": "string",
                    "time_zone": "string",
                    "week_day_of_month": [
                      "list",
                      [
                        "object",
                        {
                          "day_of_week": "string",
                          "week_of_month": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "backup_vault": {
        "computed": true,
        "description": "Backup vault where the backups gets stored using this Backup plan.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_service_account": {
        "computed": true,
        "description": "The Google Cloud Platform Service Account to be used by the BackupVault for taking backups.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "When the 'BackupPlan' was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description allows for additional details about 'BackupPlan' and its use cases to be provided.",
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
        "description": "The location for the backup plan",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of backup plan resource created",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The resource type to which the 'BackupPlan' will be applied.\nExamples include, \"compute.googleapis.com/Instance\", \"compute.googleapis.com/Disk\", and \"storage.googleapis.com/Bucket\".",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "When the 'BackupPlan' was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBackupDrBackupPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrBackupPlan), &result)
	return &result
}

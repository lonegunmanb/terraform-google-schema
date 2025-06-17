package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrBackupPlanAssociation = `{
  "block": {
    "attributes": {
      "backup_plan": {
        "computed": true,
        "description": "The BP with which resource needs to be created\nNote:\n- A Backup Plan configured for 'compute.googleapis.com/Instance', can only protect instance type resources.\n- A Backup Plan configured for 'compute.googleapis.com/Disk' can be used to protect both standard Disks and Regional Disks resources.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_plan_association_id": {
        "description": "The id of backupplan association",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the instance was created",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source": {
        "computed": true,
        "description": "Resource name of data source which will be used as storage location for backups taken",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_successful_backup_consistency_time": {
        "computed": true,
        "description": "The point in time when the last successful backup was captured from the source",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location for the backupplan association",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of backup plan association resource created",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource": {
        "computed": true,
        "description": "The resource for which BPA needs to be created",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The resource type of workload on which backupplan is applied.\nExamples include, \"compute.googleapis.com/Instance\", \"compute.googleapis.com/Disk\", and \"compute.googleapis.com/RegionDisk\"",
        "description_kind": "plain",
        "type": "string"
      },
      "rules_config_info": {
        "computed": true,
        "description": "Message for rules config info",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "last_backup_error": [
                "list",
                [
                  "object",
                  {
                    "code": "number",
                    "message": "string"
                  }
                ]
              ],
              "last_backup_state": "string",
              "rule_id": "string"
            }
          ]
        ]
      },
      "update_time": {
        "computed": true,
        "description": "The time when the instance was updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleBackupDrBackupPlanAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrBackupPlanAssociation), &result)
	return &result
}

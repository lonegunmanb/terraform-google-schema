package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrBackupPlanAssociation = `{
  "block": {
    "attributes": {
      "backup_plan": {
        "description": "The BP with which resource needs to be created",
        "description_kind": "plain",
        "required": true,
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
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource": {
        "description": "The resource for which BPA needs to be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_type": {
        "description": "The resource type of workload on which backupplan is applied",
        "description_kind": "plain",
        "required": true,
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

func GoogleBackupDrBackupPlanAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrBackupPlanAssociation), &result)
	return &result
}

package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBackupDrDataSource = `{
  "block": {
    "attributes": {
      "backup_config_info": {
        "computed": true,
        "description": "Details of how the resource is configured for backup.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_appliance_backup_config": [
                "list",
                [
                  "object",
                  {
                    "application_name": "string",
                    "backup_appliance_id": "string",
                    "backup_appliance_name": "string",
                    "host_name": "string",
                    "sla_id": "string",
                    "slp_name": "string",
                    "slt_name": "string"
                  }
                ]
              ],
              "gcp_backup_config": [
                "list",
                [
                  "object",
                  {
                    "backup_plan": "string",
                    "backup_plan_association": "string",
                    "backup_plan_description": "string",
                    "backup_plan_rules": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "last_backup_error": [
                "map",
                "string"
              ],
              "last_backup_state": "string",
              "last_successful_backup_consistency_time": "string"
            }
          ]
        ]
      },
      "backup_count": {
        "computed": true,
        "description": "Number of backups in the data source.",
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "config_state": {
        "computed": true,
        "description": "The backup configuration state.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the instance was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_source_backup_appliance_application": {
        "computed": true,
        "description": "The backed up resource is a backup appliance application.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "appliance_id": "string",
              "application_id": "string",
              "application_name": "string",
              "backup_appliance": "string",
              "host_id": "string",
              "hostname": "string",
              "type": "string"
            }
          ]
        ]
      },
      "data_source_gcp_resource": {
        "computed": true,
        "description": "The backed up resource is a Google Cloud resource.\n\t\t\tThe word 'DataSource' was included in the names to indicate that this is\n\t\t\tthe representation of the Google Cloud resource used within the\n\t\t\tDataSource object.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "compute_instance_data_source_properties": [
                "list",
                [
                  "object",
                  {
                    "description": "string",
                    "machine_type": "string",
                    "name": "string",
                    "total_disk_count": "string",
                    "total_disk_size_gb": "string"
                  }
                ]
              ],
              "gcp_resourcename": "string",
              "location": "string",
              "type": "string"
            }
          ]
        ]
      },
      "data_source_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "etag": {
        "computed": true,
        "description": "Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.",
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
        "computed": true,
        "description": "Resource labels to represent user provided metadata.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the datasource to create.\n\t\t\tIt must have the format \"projects/{project}/locations/{location}/backupVaults/{backupvault}/dataSources/{datasource}\".\n\t\t\t'{datasource}' cannot be changed after creation. It must be between 3-63 characters long and must be unique within the backup vault.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The DataSource resource instance state.",
        "description_kind": "plain",
        "type": "string"
      },
      "total_stored_bytes": {
        "computed": true,
        "description": "The number of bytes (metadata and data) stored in this datasource.",
        "description_kind": "plain",
        "type": "string"
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

func GoogleBackupDrDataSourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBackupDrDataSource), &result)
	return &result
}

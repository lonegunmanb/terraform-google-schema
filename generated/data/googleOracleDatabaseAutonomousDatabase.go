package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseAutonomousDatabase = `{
  "block": {
    "attributes": {
      "admin_password": {
        "computed": true,
        "description": "The password for the default ADMIN user.",
        "description_kind": "plain",
        "type": "string"
      },
      "autonomous_database_id": {
        "description": "The ID of the Autonomous Database to create. This value is restricted\nto (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63\ncharacters in length. The value must start with a letter and end with\na letter or a number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cidr": {
        "computed": true,
        "description": "The subnet CIDR range for the Autonmous Database.",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time that the Autonomous Database was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "database": {
        "computed": true,
        "description": "The name of the Autonomous Database. The database name must be unique in\nthe project. The name must begin with a letter and can\ncontain a maximum of 30 alphanumeric characters.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection": {
        "computed": true,
        "description": "Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.",
        "description_kind": "plain",
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "The display name for the Autonomous Database. The name does not have to\nbe unique within your project.",
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
      "entitlement_id": {
        "computed": true,
        "description": "The ID of the subscription entitlement associated with the Autonomous\nDatabase.",
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
        "description": "The labels or tags associated with the Autonomous Database. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. See documentation for resource type 'oracledatabase.googleapis.com/AutonomousDatabaseBackup'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the Autonomous Database resource in the following format:\nprojects/{project}/locations/{region}/autonomousDatabases/{autonomous_database}",
        "description_kind": "plain",
        "type": "string"
      },
      "network": {
        "computed": true,
        "description": "The name of the VPC network used by the Autonomous Database.\nFormat: projects/{project}/global/networks/{network}",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_network": {
        "computed": true,
        "description": "The name of the OdbNetwork associated with the Autonomous Database.\nFormat:\nprojects/{project}/locations/{location}/odbNetworks/{odb_network}\nIt is optional but if specified, this should match the parent ODBNetwork of\nthe odb_subnet and backup_odb_subnet.",
        "description_kind": "plain",
        "type": "string"
      },
      "odb_subnet": {
        "computed": true,
        "description": "The name of the OdbSubnet associated with the Autonomous Database for\nIP allocation. Format:\nprojects/{project}/locations/{location}/odbNetworks/{odb_network}/odbSubnets/{odb_subnet}",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "properties": {
        "computed": true,
        "description": "The properties of an Autonomous Database.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actual_used_data_storage_size_tb": "number",
              "allocated_storage_size_tb": "number",
              "apex_details": [
                "list",
                [
                  "object",
                  {
                    "apex_version": "string",
                    "ords_version": "string"
                  }
                ]
              ],
              "are_primary_allowlisted_ips_used": "bool",
              "autonomous_container_database_id": "string",
              "available_upgrade_versions": [
                "list",
                "string"
              ],
              "backup_retention_period_days": "number",
              "character_set": "string",
              "compute_count": "number",
              "connection_strings": [
                "list",
                [
                  "object",
                  {
                    "all_connection_strings": [
                      "list",
                      [
                        "object",
                        {
                          "high": "string",
                          "low": "string",
                          "medium": "string"
                        }
                      ]
                    ],
                    "dedicated": "string",
                    "high": "string",
                    "low": "string",
                    "medium": "string",
                    "profiles": [
                      "list",
                      [
                        "object",
                        {
                          "consumer_group": "string",
                          "display_name": "string",
                          "host_format": "string",
                          "is_regional": "bool",
                          "protocol": "string",
                          "session_mode": "string",
                          "syntax_format": "string",
                          "tls_authentication": "string",
                          "value": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "connection_urls": [
                "list",
                [
                  "object",
                  {
                    "apex_uri": "string",
                    "database_transforms_uri": "string",
                    "graph_studio_uri": "string",
                    "machine_learning_notebook_uri": "string",
                    "machine_learning_user_management_uri": "string",
                    "mongo_db_uri": "string",
                    "ords_uri": "string",
                    "sql_dev_web_uri": "string"
                  }
                ]
              ],
              "customer_contacts": [
                "list",
                [
                  "object",
                  {
                    "email": "string"
                  }
                ]
              ],
              "data_safe_state": "string",
              "data_storage_size_gb": "number",
              "data_storage_size_tb": "number",
              "database_management_state": "string",
              "db_edition": "string",
              "db_version": "string",
              "db_workload": "string",
              "failed_data_recovery_duration": "string",
              "is_auto_scaling_enabled": "bool",
              "is_local_data_guard_enabled": "bool",
              "is_storage_auto_scaling_enabled": "bool",
              "license_type": "string",
              "lifecycle_details": "string",
              "local_adg_auto_failover_max_data_loss_limit": "number",
              "local_disaster_recovery_type": "string",
              "local_standby_db": [
                "list",
                [
                  "object",
                  {
                    "data_guard_role_changed_time": "string",
                    "disaster_recovery_role_changed_time": "string",
                    "lag_time_duration": "string",
                    "lifecycle_details": "string",
                    "state": "string"
                  }
                ]
              ],
              "maintenance_begin_time": "string",
              "maintenance_end_time": "string",
              "maintenance_schedule_type": "string",
              "memory_per_oracle_compute_unit_gbs": "number",
              "memory_table_gbs": "number",
              "mtls_connection_required": "bool",
              "n_character_set": "string",
              "next_long_term_backup_time": "string",
              "oci_url": "string",
              "ocid": "string",
              "open_mode": "string",
              "operations_insights_state": "string",
              "peer_db_ids": [
                "list",
                "string"
              ],
              "permission_level": "string",
              "private_endpoint": "string",
              "private_endpoint_ip": "string",
              "private_endpoint_label": "string",
              "refreshable_mode": "string",
              "refreshable_state": "string",
              "role": "string",
              "scheduled_operation_details": [
                "list",
                [
                  "object",
                  {
                    "day_of_week": "string",
                    "start_time": [
                      "list",
                      [
                        "object",
                        {
                          "hours": "number",
                          "minutes": "number",
                          "nanos": "number",
                          "seconds": "number"
                        }
                      ]
                    ],
                    "stop_time": [
                      "list",
                      [
                        "object",
                        {
                          "hours": "number",
                          "minutes": "number",
                          "nanos": "number",
                          "seconds": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "sql_web_developer_url": "string",
              "state": "string",
              "supported_clone_regions": [
                "list",
                "string"
              ],
              "total_auto_backup_storage_size_gbs": "number",
              "used_data_storage_size_tbs": "number"
            }
          ]
        ]
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOracleDatabaseAutonomousDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseAutonomousDatabase), &result)
	return &result
}

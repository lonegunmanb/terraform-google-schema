package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseAutonomousDatabases = `{
  "block": {
    "attributes": {
      "autonomous_databases": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "admin_password": "string",
              "autonomous_database_id": "string",
              "cidr": "string",
              "create_time": "string",
              "database": "string",
              "deletion_protection": "bool",
              "display_name": "string",
              "effective_labels": [
                "map",
                "string"
              ],
              "entitlement_id": "string",
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "name": "string",
              "network": "string",
              "odb_network": "string",
              "odb_subnet": "string",
              "project": "string",
              "properties": [
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
              ],
              "terraform_labels": [
                "map",
                "string"
              ]
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "location",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "The ID of the project in which the dataset is located. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOracleDatabaseAutonomousDatabasesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseAutonomousDatabases), &result)
	return &result
}

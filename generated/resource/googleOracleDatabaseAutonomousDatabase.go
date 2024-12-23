package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseAutonomousDatabase = `{
  "block": {
    "attributes": {
      "admin_password": {
        "description": "The password for the default ADMIN user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "autonomous_database_id": {
        "description": "The ID of the Autonomous Database to create. This value is restricted\nto (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63\ncharacters in length. The value must start with a letter and end with\na letter or a number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cidr": {
        "description": "The subnet CIDR range for the Autonmous Database.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time that the Autonomous Database was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "database": {
        "description": "The name of the Autonomous Database. The database name must be unique in\nthe project. The name must begin with a letter and can\ncontain a maximum of 30 alphanumeric characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "deletion_protection": {
        "description": "Whether or not to allow Terraform to destroy the instance. Unless this field is set to false in Terraform state, a terraform destroy or terraform apply that would delete the instance will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "display_name": {
        "computed": true,
        "description": "The display name for the Autonomous Database. The name does not have to\nbe unique within your project.",
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
        "description": "The labels or tags associated with the Autonomous Database. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "The name of the VPC network used by the Autonomous Database.\nFormat: projects/{project}/global/networks/{network}",
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
      }
    },
    "block_types": {
      "properties": {
        "block": {
          "attributes": {
            "actual_used_data_storage_size_tb": {
              "computed": true,
              "description": "The amount of storage currently being used for user and system data, in\nterabytes.",
              "description_kind": "plain",
              "type": "number"
            },
            "allocated_storage_size_tb": {
              "computed": true,
              "description": "The amount of storage currently allocated for the database tables and\nbilled for, rounded up in terabytes.",
              "description_kind": "plain",
              "type": "number"
            },
            "apex_details": {
              "computed": true,
              "description": "Oracle APEX Application Development.\nhttps://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/datatypes/AutonomousDatabaseApex",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "apex_version": "string",
                    "ords_version": "string"
                  }
                ]
              ]
            },
            "are_primary_allowlisted_ips_used": {
              "computed": true,
              "description": "This field indicates the status of Data Guard and Access control for the\nAutonomous Database. The field's value is null if Data Guard is disabled\nor Access Control is disabled. The field's value is TRUE if both Data Guard\nand Access Control are enabled, and the Autonomous Database is using\nprimary IP access control list (ACL) for standby. The field's value is\nFALSE if both Data Guard and Access Control are enabled, and the Autonomous\nDatabase is using a different IP access control list (ACL) for standby\ncompared to primary.",
              "description_kind": "plain",
              "type": "bool"
            },
            "autonomous_container_database_id": {
              "computed": true,
              "description": "The Autonomous Container Database OCID.",
              "description_kind": "plain",
              "type": "string"
            },
            "available_upgrade_versions": {
              "computed": true,
              "description": "The list of available Oracle Database upgrade versions for an Autonomous\nDatabase.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "backup_retention_period_days": {
              "computed": true,
              "description": "The retention period for the Autonomous Database. This field is specified\nin days, can range from 1 day to 60 days, and has a default value of\n60 days.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "character_set": {
              "description": "The character set for the Autonomous Database. The default is AL32UTF8.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "compute_count": {
              "computed": true,
              "description": "The number of compute servers for the Autonomous Database.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "connection_strings": {
              "computed": true,
              "description": "The connection string used to connect to the Autonomous Database.\nhttps://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/datatypes/AutonomousDatabaseConnectionStrings",
              "description_kind": "plain",
              "type": [
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
              ]
            },
            "connection_urls": {
              "computed": true,
              "description": "The URLs for accessing Oracle Application Express (APEX) and SQL Developer\nWeb with a browser from a Compute instance.\nhttps://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/datatypes/AutonomousDatabaseConnectionUrls",
              "description_kind": "plain",
              "type": [
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
              ]
            },
            "data_safe_state": {
              "computed": true,
              "description": "The current state of the Data Safe registration for the\nAutonomous Database. \n Possible values:\n DATA_SAFE_STATE_UNSPECIFIED\nREGISTERING\nREGISTERED\nDEREGISTERING\nNOT_REGISTERED\nFAILED",
              "description_kind": "plain",
              "type": "string"
            },
            "data_storage_size_gb": {
              "computed": true,
              "description": "The size of the data stored in the database, in gigabytes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "data_storage_size_tb": {
              "computed": true,
              "description": "The size of the data stored in the database, in terabytes.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "database_management_state": {
              "computed": true,
              "description": "The current state of database management for the Autonomous Database. \n Possible values:\n DATABASE_MANAGEMENT_STATE_UNSPECIFIED\nENABLING\nENABLED\nDISABLING\nNOT_ENABLED\nFAILED_ENABLING\nFAILED_DISABLING",
              "description_kind": "plain",
              "type": "string"
            },
            "db_edition": {
              "description": "The edition of the Autonomous Databases. \n Possible values:\n DATABASE_EDITION_UNSPECIFIED\nSTANDARD_EDITION\nENTERPRISE_EDITION",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "db_version": {
              "description": "The Oracle Database version for the Autonomous Database.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "db_workload": {
              "description": "Possible values:\n DB_WORKLOAD_UNSPECIFIED\nOLTP\nDW\nAJD\nAPEX",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "failed_data_recovery_duration": {
              "computed": true,
              "description": "This field indicates the number of seconds of data loss during a Data\nGuard failover.",
              "description_kind": "plain",
              "type": "string"
            },
            "is_auto_scaling_enabled": {
              "description": "This field indicates if auto scaling is enabled for the Autonomous Database\nCPU core count.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "is_local_data_guard_enabled": {
              "computed": true,
              "description": "This field indicates whether the Autonomous Database has local (in-region)\nData Guard enabled.",
              "description_kind": "plain",
              "type": "bool"
            },
            "is_storage_auto_scaling_enabled": {
              "computed": true,
              "description": "This field indicates if auto scaling is enabled for the Autonomous Database\nstorage.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "license_type": {
              "description": "The license type used for the Autonomous Database. \n Possible values:\n LICENSE_TYPE_UNSPECIFIED\nLICENSE_INCLUDED\nBRING_YOUR_OWN_LICENSE",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "lifecycle_details": {
              "computed": true,
              "description": "The details of the current lifestyle state of the Autonomous Database.",
              "description_kind": "plain",
              "type": "string"
            },
            "local_adg_auto_failover_max_data_loss_limit": {
              "computed": true,
              "description": "This field indicates the maximum data loss limit for an Autonomous\nDatabase, in seconds.",
              "description_kind": "plain",
              "type": "number"
            },
            "local_disaster_recovery_type": {
              "computed": true,
              "description": "This field indicates the local disaster recovery (DR) type of an\nAutonomous Database. \n Possible values:\n LOCAL_DISASTER_RECOVERY_TYPE_UNSPECIFIED\nADG\nBACKUP_BASED",
              "description_kind": "plain",
              "type": "string"
            },
            "local_standby_db": {
              "computed": true,
              "description": "Autonomous Data Guard standby database details.\nhttps://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/datatypes/AutonomousDatabaseStandbySummary",
              "description_kind": "plain",
              "type": [
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
              ]
            },
            "maintenance_begin_time": {
              "computed": true,
              "description": "The date and time when maintenance will begin.",
              "description_kind": "plain",
              "type": "string"
            },
            "maintenance_end_time": {
              "computed": true,
              "description": "The date and time when maintenance will end.",
              "description_kind": "plain",
              "type": "string"
            },
            "maintenance_schedule_type": {
              "computed": true,
              "description": "The maintenance schedule of the Autonomous Database. \n Possible values:\n MAINTENANCE_SCHEDULE_TYPE_UNSPECIFIED\nEARLY\nREGULAR",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "memory_per_oracle_compute_unit_gbs": {
              "computed": true,
              "description": "The amount of memory enabled per ECPU, in gigabytes.",
              "description_kind": "plain",
              "type": "number"
            },
            "memory_table_gbs": {
              "computed": true,
              "description": "The memory assigned to in-memory tables in an Autonomous Database.",
              "description_kind": "plain",
              "type": "number"
            },
            "mtls_connection_required": {
              "description": "This field specifies if the Autonomous Database requires mTLS connections.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "n_character_set": {
              "description": "The national character set for the Autonomous Database. The default is\nAL16UTF16.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "next_long_term_backup_time": {
              "computed": true,
              "description": "The long term backup schedule of the Autonomous Database.",
              "description_kind": "plain",
              "type": "string"
            },
            "oci_url": {
              "computed": true,
              "description": "The Oracle Cloud Infrastructure link for the Autonomous Database.",
              "description_kind": "plain",
              "type": "string"
            },
            "ocid": {
              "computed": true,
              "description": "OCID of the Autonomous Database.\nhttps://docs.oracle.com/en-us/iaas/Content/General/Concepts/identifiers.htm#Oracle",
              "description_kind": "plain",
              "type": "string"
            },
            "open_mode": {
              "computed": true,
              "description": "This field indicates the current mode of the Autonomous Database. \n Possible values:\n OPEN_MODE_UNSPECIFIED\nREAD_ONLY\nREAD_WRITE",
              "description_kind": "plain",
              "type": "string"
            },
            "operations_insights_state": {
              "computed": true,
              "description": "Possible values:\n OPERATIONS_INSIGHTS_STATE_UNSPECIFIED\nENABLING\nENABLED\nDISABLING\nNOT_ENABLED\nFAILED_ENABLING\nFAILED_DISABLING",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "peer_db_ids": {
              "computed": true,
              "description": "The list of OCIDs of standby databases located in Autonomous Data Guard\nremote regions that are associated with the source database.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "permission_level": {
              "computed": true,
              "description": "The permission level of the Autonomous Database. \n Possible values:\n PERMISSION_LEVEL_UNSPECIFIED\nRESTRICTED\nUNRESTRICTED",
              "description_kind": "plain",
              "type": "string"
            },
            "private_endpoint": {
              "computed": true,
              "description": "The private endpoint for the Autonomous Database.",
              "description_kind": "plain",
              "type": "string"
            },
            "private_endpoint_ip": {
              "computed": true,
              "description": "The private endpoint IP address for the Autonomous Database.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_endpoint_label": {
              "computed": true,
              "description": "The private endpoint label for the Autonomous Database.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "refreshable_mode": {
              "computed": true,
              "description": "The refresh mode of the cloned Autonomous Database. \n Possible values:\n REFRESHABLE_MODE_UNSPECIFIED\nAUTOMATIC\nMANUAL",
              "description_kind": "plain",
              "type": "string"
            },
            "refreshable_state": {
              "computed": true,
              "description": "The refresh State of the clone. \n Possible values:\n REFRESHABLE_STATE_UNSPECIFIED\nREFRESHING\nNOT_REFRESHING",
              "description_kind": "plain",
              "type": "string"
            },
            "role": {
              "computed": true,
              "description": "The Data Guard role of the Autonomous Database. \n Possible values:\n ROLE_UNSPECIFIED\nPRIMARY\nSTANDBY\nDISABLED_STANDBY\nBACKUP_COPY\nSNAPSHOT_STANDBY",
              "description_kind": "plain",
              "type": "string"
            },
            "scheduled_operation_details": {
              "computed": true,
              "description": "The list and details of the scheduled operations of the Autonomous\nDatabase.",
              "description_kind": "plain",
              "type": [
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
              ]
            },
            "sql_web_developer_url": {
              "computed": true,
              "description": "The SQL Web Developer URL for the Autonomous Database.",
              "description_kind": "plain",
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "Possible values:\n STATE_UNSPECIFIED\nPROVISIONING\nAVAILABLE\nSTOPPING\nSTOPPED\nSTARTING\nTERMINATING\nTERMINATED\nUNAVAILABLE\nRESTORE_IN_PROGRESS\nRESTORE_FAILED\nBACKUP_IN_PROGRESS\nSCALE_IN_PROGRESS\nAVAILABLE_NEEDS_ATTENTION\nUPDATING\nMAINTENANCE_IN_PROGRESS\nRESTARTING\nRECREATING\nROLE_CHANGE_IN_PROGRESS\nUPGRADING\nINACCESSIBLE\nSTANDBY",
              "description_kind": "plain",
              "type": "string"
            },
            "supported_clone_regions": {
              "computed": true,
              "description": "The list of available regions that can be used to create a clone for the\nAutonomous Database.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "total_auto_backup_storage_size_gbs": {
              "computed": true,
              "description": "The storage space used by automatic backups of Autonomous Database, in\ngigabytes.",
              "description_kind": "plain",
              "type": "number"
            },
            "used_data_storage_size_tbs": {
              "computed": true,
              "description": "The storage space used by Autonomous Database, in gigabytes.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "block_types": {
            "customer_contacts": {
              "block": {
                "attributes": {
                  "email": {
                    "description": "The email address used by Oracle to send notifications regarding databases\nand infrastructure.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The list of customer contacts.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "The properties of an Autonomous Database.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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

func GoogleOracleDatabaseAutonomousDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseAutonomousDatabase), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOracleDatabaseCloudExadataInfrastructure = `{
  "block": {
    "attributes": {
      "cloud_exadata_infrastructure_id": {
        "description": "The ID of the Exadata Infrastructure to create. This value is restricted\nto (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$) and must be a maximum of 63\ncharacters in length. The value must start with a letter and end with\na letter or a number.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The date and time that the Exadata Infrastructure was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "User friendly name for this resource.",
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
        "description": "Entitlement ID of the private offer against which this infrastructure\nresource is provisioned.",
        "description_kind": "plain",
        "type": "string"
      },
      "gcp_oracle_zone": {
        "computed": true,
        "description": "GCP location where Oracle Exadata is hosted.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "Labels or tags associated with the resource. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. See documentation for resource type 'oracledatabase.googleapis.com/DbServer'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The name of the Exadata Infrastructure resource with the following format:\nprojects/{project}/locations/{region}/cloudExadataInfrastructures/{cloud_exadata_infrastructure}",
        "description_kind": "plain",
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
            "activated_storage_count": {
              "computed": true,
              "description": "The requested number of additional storage servers activated for the\nExadata Infrastructure.",
              "description_kind": "plain",
              "type": "number"
            },
            "additional_storage_count": {
              "computed": true,
              "description": "The requested number of additional storage servers for the Exadata\nInfrastructure.",
              "description_kind": "plain",
              "type": "number"
            },
            "available_storage_size_gb": {
              "computed": true,
              "description": "The available storage can be allocated to the Exadata Infrastructure\nresource, in gigabytes (GB).",
              "description_kind": "plain",
              "type": "number"
            },
            "compute_count": {
              "description": "The number of compute servers for the Exadata Infrastructure.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cpu_count": {
              "computed": true,
              "description": "The number of enabled CPU cores.",
              "description_kind": "plain",
              "type": "number"
            },
            "data_storage_size_tb": {
              "computed": true,
              "description": "Size, in terabytes, of the DATA disk group.",
              "description_kind": "plain",
              "type": "number"
            },
            "db_node_storage_size_gb": {
              "computed": true,
              "description": "The local node storage allocated in GBs.",
              "description_kind": "plain",
              "type": "number"
            },
            "db_server_version": {
              "computed": true,
              "description": "The software version of the database servers (dom0) in the Exadata\nInfrastructure.",
              "description_kind": "plain",
              "type": "string"
            },
            "max_cpu_count": {
              "computed": true,
              "description": "The total number of CPU cores available.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_data_storage_tb": {
              "computed": true,
              "description": "The total available DATA disk group size.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_db_node_storage_size_gb": {
              "computed": true,
              "description": "The total local node storage available in GBs.",
              "description_kind": "plain",
              "type": "number"
            },
            "max_memory_gb": {
              "computed": true,
              "description": "The total memory available in GBs.",
              "description_kind": "plain",
              "type": "number"
            },
            "memory_size_gb": {
              "computed": true,
              "description": "The memory allocated in GBs.",
              "description_kind": "plain",
              "type": "number"
            },
            "monthly_db_server_version": {
              "computed": true,
              "description": "The monthly software version of the database servers (dom0)\nin the Exadata Infrastructure. Example: 20.1.15",
              "description_kind": "plain",
              "type": "string"
            },
            "monthly_storage_server_version": {
              "computed": true,
              "description": "The monthly software version of the storage servers (cells)\nin the Exadata Infrastructure. Example: 20.1.15",
              "description_kind": "plain",
              "type": "string"
            },
            "next_maintenance_run_id": {
              "computed": true,
              "description": "The OCID of the next maintenance run.",
              "description_kind": "plain",
              "type": "string"
            },
            "next_maintenance_run_time": {
              "computed": true,
              "description": "The time when the next maintenance run will occur.",
              "description_kind": "plain",
              "type": "string"
            },
            "next_security_maintenance_run_time": {
              "computed": true,
              "description": "The time when the next security maintenance run will occur.",
              "description_kind": "plain",
              "type": "string"
            },
            "oci_url": {
              "computed": true,
              "description": "Deep link to the OCI console to view this resource.",
              "description_kind": "plain",
              "type": "string"
            },
            "ocid": {
              "computed": true,
              "description": "OCID of created infra.\nhttps://docs.oracle.com/en-us/iaas/Content/General/Concepts/identifiers.htm#Oracle",
              "description_kind": "plain",
              "type": "string"
            },
            "shape": {
              "description": "The shape of the Exadata Infrastructure. The shape determines the\namount of CPU, storage, and memory resources allocated to the instance.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "state": {
              "computed": true,
              "description": "The current lifecycle state of the Exadata Infrastructure. \n Possible values:\n STATE_UNSPECIFIED\nPROVISIONING\nAVAILABLE\nUPDATING\nTERMINATING\nTERMINATED\nFAILED\nMAINTENANCE_IN_PROGRESS",
              "description_kind": "plain",
              "type": "string"
            },
            "storage_count": {
              "description": "The number of Cloud Exadata storage servers for the Exadata Infrastructure.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "storage_server_version": {
              "computed": true,
              "description": "The software version of the storage servers (cells) in the Exadata\nInfrastructure.",
              "description_kind": "plain",
              "type": "string"
            },
            "total_storage_size_gb": {
              "computed": true,
              "description": "The total storage allocated to the Exadata Infrastructure\nresource, in gigabytes (GB).",
              "description_kind": "plain",
              "optional": true,
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
            },
            "maintenance_window": {
              "block": {
                "attributes": {
                  "custom_action_timeout_mins": {
                    "computed": true,
                    "description": "Determines the amount of time the system will wait before the start of each\ndatabase server patching operation. Custom action timeout is in minutes and\nvalid value is between 15 to 120 (inclusive).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "days_of_week": {
                    "computed": true,
                    "description": "Days during the week when maintenance should be performed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "hours_of_day": {
                    "computed": true,
                    "description": "The window of hours during the day when maintenance should be performed.\nThe window is a 4 hour slot. Valid values are:\n  0 - represents time slot 0:00 - 3:59 UTC\n  4 - represents time slot 4:00 - 7:59 UTC\n  8 - represents time slot 8:00 - 11:59 UTC\n  12 - represents time slot 12:00 - 15:59 UTC\n  16 - represents time slot 16:00 - 19:59 UTC\n  20 - represents time slot 20:00 - 23:59 UTC",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  },
                  "is_custom_action_timeout_enabled": {
                    "computed": true,
                    "description": "If true, enables the configuration of a custom action timeout (waiting\nperiod) between database server patching operations.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "lead_time_week": {
                    "computed": true,
                    "description": "Lead time window allows user to set a lead time to prepare for a down time.\nThe lead time is in weeks and valid value is between 1 to 4.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "months": {
                    "computed": true,
                    "description": "Months during the year when maintenance should be performed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "patching_mode": {
                    "computed": true,
                    "description": "Cloud CloudExadataInfrastructure node patching method, either \"ROLLING\"\n or \"NONROLLING\". Default value is ROLLING. \n Possible values:\n PATCHING_MODE_UNSPECIFIED\nROLLING\nNON_ROLLING",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "preference": {
                    "computed": true,
                    "description": "The maintenance window scheduling preference. \n Possible values:\n MAINTENANCE_WINDOW_PREFERENCE_UNSPECIFIED\nCUSTOM_PREFERENCE\nNO_PREFERENCE",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "weeks_of_month": {
                    "computed": true,
                    "description": "Weeks during the month when maintenance should be performed. Weeks start on\nthe 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7\ndays. Weeks start and end based on calendar dates, not days of the week.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "description": "Maintenance window as defined by Oracle.\nhttps://docs.oracle.com/en-us/iaas/api/#/en/database/20160918/datatypes/MaintenanceWindow",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Various properties of Exadata Infrastructure.",
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

func GoogleOracleDatabaseCloudExadataInfrastructureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOracleDatabaseCloudExadataInfrastructure), &result)
	return &result
}

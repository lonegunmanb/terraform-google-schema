package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbCluster = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "backup_source": {
        "computed": true,
        "description": "Cluster created from backup.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_name": "string"
            }
          ]
        ]
      },
      "cluster_id": {
        "description": "The ID of the alloydb cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cluster_type": {
        "description": "The type of cluster. If not set, defaults to PRIMARY. Default value: \"PRIMARY\" Possible values: [\"PRIMARY\", \"SECONDARY\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "continuous_backup_info": {
        "computed": true,
        "description": "ContinuousBackupInfo describes the continuous backup properties of a cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "earliest_restorable_time": "string",
              "enabled_time": "string",
              "encryption_info": [
                "list",
                [
                  "object",
                  {
                    "encryption_type": "string",
                    "kms_key_versions": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "schedule": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "database_version": {
        "computed": true,
        "description": "The database engine major version. This is an optional field and it's populated at the Cluster creation time. This field cannot be changed after cluster creation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deletion_policy": {
        "description": "Policy to determine if the cluster should be deleted forcefully.\nDeleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.\nDeleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = \"FORCE\" otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.\nPossible values: DEFAULT, FORCE",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "User-settable and human-readable display name for the Cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "effective_annotations": {
        "computed": true,
        "description": "All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
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
      "encryption_info": {
        "computed": true,
        "description": "EncryptionInfo describes the encryption information of a cluster or a backup.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "encryption_type": "string",
              "kms_key_versions": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "etag": {
        "description": "For Resource freshness validation (https://google.aip.dev/154)",
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
        "description": "User-defined labels for the alloydb cluster.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the alloydb cluster should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "migration_source": {
        "computed": true,
        "description": "Cluster created via DMS migration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host_port": "string",
              "reference_id": "string",
              "source_type": "string"
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "The name of the cluster resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "Output only. Reconciling (https://google.aip.dev/128#reconciliation).\nSet to true if the current state of Cluster does not match the user's intended state, and the service is actively updating the resource to reconcile them.\nThis can happen due to user-triggered updates or system actions like failover or maintenance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "Output only. The current serving state of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "subscription_type": {
        "computed": true,
        "description": "The subscrition type of cluster. Possible values: [\"TRIAL\", \"STANDARD\"]",
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
      },
      "trial_metadata": {
        "computed": true,
        "description": "Contains information and all metadata related to TRIAL clusters.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "end_time": "string",
              "grace_end_time": "string",
              "start_time": "string",
              "upgrade_time": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "The system-generated UID of the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "automated_backup_policy": {
        "block": {
          "attributes": {
            "backup_window": {
              "computed": true,
              "description": "The length of the time window during which a backup can be taken. If a backup does not succeed within this time window, it will be canceled and considered failed.\n\nThe backup window must be at least 5 minutes long. There is no upper bound on the window. If not set, it will default to 1 hour.\n\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description": "Whether automated backups are enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "labels": {
              "description": "Labels to apply to backups created using this configuration.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "location": {
              "computed": true,
              "description": "The location where the backup will be stored. Currently, the only supported option is to store the backup in the same region as the cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "encryption_config": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "quantity_based_retention": {
              "block": {
                "attributes": {
                  "count": {
                    "description": "The number of backups to retain.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description": "Quantity-based Backup retention policy to retain recent backups. Conflicts with 'time_based_retention', both can't be set together.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "time_based_retention": {
              "block": {
                "attributes": {
                  "retention_period": {
                    "description": "The retention period.\nA duration in seconds with up to nine fractional digits, terminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Time-based Backup retention policy. Conflicts with 'quantity_based_retention', both can't be set together.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "weekly_schedule": {
              "block": {
                "attributes": {
                  "days_of_week": {
                    "description": "The days of the week to perform a backup. At least one day of the week must be provided. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "start_times": {
                    "block": {
                      "attributes": {
                        "hours": {
                          "description": "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minutes": {
                          "description": "Minutes of hour of day. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "nanos": {
                          "description": "Fractions of seconds in nanoseconds. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Seconds of minutes of the time. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "The times during the day to start a backup. At least one start time must be provided. The start times are assumed to be in UTC and to be an exact hour (e.g., 04:00:00).",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Weekly schedule for the Backup.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "continuous_backup_config": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "Whether continuous backup recovery is enabled. If not set, defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "recovery_window_days": {
              "computed": true,
              "description": "The numbers of days that are eligible to restore from using PITR. To support the entire recovery window, backups and logs are retained for one day more than the recovery window.\n\nIf not set, defaults to 14 days.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "encryption_config": {
              "block": {
                "attributes": {
                  "kms_key_name": {
                    "description": "The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The continuous backup config for this cluster.\n\nIf no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "encryption_config": {
        "block": {
          "attributes": {
            "kms_key_name": {
              "description": "The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "initial_user": {
        "block": {
          "attributes": {
            "password": {
              "description": "The initial password for the user.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "user": {
              "description": "The database username.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Initial user to setup during cluster creation.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_update_policy": {
        "block": {
          "block_types": {
            "maintenance_windows": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Preferred day of the week for maintenance, e.g. MONDAY, TUESDAY, etc. Possible values: [\"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "start_time": {
                    "block": {
                      "attributes": {
                        "hours": {
                          "description": "Hours of day in 24 hour format. Should be from 0 to 23.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "minutes": {
                          "description": "Minutes of hour of day. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "nanos": {
                          "description": "Fractions of seconds in nanoseconds. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Seconds of minutes of the time. Currently, only the value 0 is supported.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Preferred time to start the maintenance operation on the specified day. Maintenance will start within 1 hour of this time.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Preferred windows to perform maintenance. Currently limited to 1.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "MaintenanceUpdatePolicy defines the policy for system updates.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "allocated_ip_range": {
              "description": "The name of the allocated IP range for the private IP AlloyDB cluster. For example: \"google-managed-services-default\".\nIf set, the instance IPs for this cluster will be created in the allocated range.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description": "The resource link for the VPC network in which cluster resources are created and from which they are accessible via Private IP. The network must belong to the same project as the cluster.\nIt is specified in the form: \"projects/{projectNumber}/global/networks/{network_id}\".",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Metadata related to network configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "psc_config": {
        "block": {
          "attributes": {
            "psc_enabled": {
              "description": "Create an instance that allows connections from Private Service Connect endpoints to the instance.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description": "Configuration for Private Service Connect (PSC) for the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "restore_backup_source": {
        "block": {
          "attributes": {
            "backup_name": {
              "description": "The name of the backup that this cluster is restored from.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The source when restoring from a backup. Conflicts with 'restore_continuous_backup_source', both can't be set together.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "restore_continuous_backup_source": {
        "block": {
          "attributes": {
            "cluster": {
              "description": "The name of the source cluster that this cluster is restored from.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "point_in_time": {
              "description": "The point in time that this cluster is restored to, in RFC 3339 format.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The source when restoring via point in time recovery (PITR). Conflicts with 'restore_backup_source', both can't be set together.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "secondary_config": {
        "block": {
          "attributes": {
            "primary_cluster_name": {
              "description": "Name of the primary cluster must be in the format\n'projects/{project}/locations/{location}/clusters/{cluster_id}'",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.",
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

func GoogleAlloydbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbCluster), &result)
	return &result
}

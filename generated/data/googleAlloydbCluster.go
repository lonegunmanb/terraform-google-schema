package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbCluster = `{
  "block": {
    "attributes": {
      "annotations": {
        "computed": true,
        "description": "Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128\nAn object containing a list of \"key\": value pairs. Example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "automated_backup_policy": {
        "computed": true,
        "description": "The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup_window": "string",
              "enabled": "bool",
              "encryption_config": [
                "list",
                [
                  "object",
                  {
                    "kms_key_name": "string"
                  }
                ]
              ],
              "labels": [
                "map",
                "string"
              ],
              "location": "string",
              "quantity_based_retention": [
                "list",
                [
                  "object",
                  {
                    "count": "number"
                  }
                ]
              ],
              "time_based_retention": [
                "list",
                [
                  "object",
                  {
                    "retention_period": "string"
                  }
                ]
              ],
              "weekly_schedule": [
                "list",
                [
                  "object",
                  {
                    "days_of_week": [
                      "list",
                      "string"
                    ],
                    "start_times": [
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
            }
          ]
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
        "computed": true,
        "description": "The type of cluster. If not set, defaults to PRIMARY. Default value: \"PRIMARY\" Possible values: [\"PRIMARY\", \"SECONDARY\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "continuous_backup_config": {
        "computed": true,
        "description": "The continuous backup config for this cluster.\n\nIf no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "encryption_config": [
                "list",
                [
                  "object",
                  {
                    "kms_key_name": "string"
                  }
                ]
              ],
              "recovery_window_days": "number"
            }
          ]
        ]
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
        "description": "The database engine major version. This is an optional field and it's populated at the Cluster creation time.\nNote: Changing this field to a higer version results in upgrading the AlloyDB cluster which is an irreversible change.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_policy": {
        "computed": true,
        "description": "Policy to determine if the cluster should be deleted forcefully.\nDeleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.\nDeleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = \"FORCE\" otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.\nPossible values: DEFAULT, FORCE",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "User-settable and human-readable display name for the Cluster.",
        "description_kind": "plain",
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
      "encryption_config": {
        "computed": true,
        "description": "EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_name": "string"
            }
          ]
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
        "computed": true,
        "description": "For Resource freshness validation (https://google.aip.dev/154)",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "initial_user": {
        "computed": true,
        "description": "Initial user to setup during cluster creation.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "password": "string",
              "user": "string"
            }
          ]
        ]
      },
      "labels": {
        "computed": true,
        "description": "User-defined labels for the alloydb cluster.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location where the alloydb cluster should reside.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maintenance_update_policy": {
        "computed": true,
        "description": "MaintenanceUpdatePolicy defines the policy for system updates.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "maintenance_windows": [
                "list",
                [
                  "object",
                  {
                    "day": "string",
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
                    ]
                  }
                ]
              ]
            }
          ]
        ]
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
      "network_config": {
        "computed": true,
        "description": "Metadata related to network configuration.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocated_ip_range": "string",
              "network": "string"
            }
          ]
        ]
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_config": {
        "computed": true,
        "description": "Configuration for Private Service Connect (PSC) for the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "psc_enabled": "bool",
              "service_owned_project_number": "number"
            }
          ]
        ]
      },
      "reconciling": {
        "computed": true,
        "description": "Output only. Reconciling (https://google.aip.dev/128#reconciliation).\nSet to true if the current state of Cluster does not match the user's intended state, and the service is actively updating the resource to reconcile them.\nThis can happen due to user-triggered updates or system actions like failover or maintenance.",
        "description_kind": "plain",
        "type": "bool"
      },
      "restore_backup_source": {
        "computed": true,
        "description": "The source when restoring from a backup. Conflicts with 'restore_continuous_backup_source', both can't be set together.",
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
      "restore_continuous_backup_source": {
        "computed": true,
        "description": "The source when restoring via point in time recovery (PITR). Conflicts with 'restore_backup_source', both can't be set together.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster": "string",
              "point_in_time": "string"
            }
          ]
        ]
      },
      "secondary_config": {
        "computed": true,
        "description": "Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "primary_cluster_name": "string"
            }
          ]
        ]
      },
      "skip_await_major_version_upgrade": {
        "computed": true,
        "description": "Set to true to skip awaiting on the major version upgrade of the cluster.\nPossible values: true, false\nDefault value: \"true\"",
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
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleAlloydbClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbCluster), &result)
	return &result
}

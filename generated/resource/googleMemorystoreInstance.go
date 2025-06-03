package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleMemorystoreInstance = `{
  "block": {
    "attributes": {
      "authorization_mode": {
        "computed": true,
        "description": "Optional. Immutable. Authorization mode of the instance. Possible values:\n AUTH_DISABLED\nIAM_AUTH",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "backup_collection": {
        "computed": true,
        "description": "The backup collection full resource name.\nExample: projects/{project}/locations/{location}/backupCollections/{collection}",
        "description_kind": "plain",
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Output only. Creation timestamp of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "description": "Optional. If set to true deletion of the instance will fail.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "discovery_endpoints": {
        "computed": true,
        "deprecated": true,
        "description": "Output only. Endpoints clients can connect to the instance through. Currently only one\ndiscovery endpoint is supported.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "network": "string",
              "port": "number"
            }
          ]
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
      "endpoints": {
        "computed": true,
        "description": "Endpoints for the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connections": [
                "list",
                [
                  "object",
                  {
                    "psc_auto_connection": [
                      "list",
                      [
                        "object",
                        {
                          "connection_type": "string",
                          "forwarding_rule": "string",
                          "ip_address": "string",
                          "network": "string",
                          "port": "number",
                          "project_id": "string",
                          "psc_connection_id": "string",
                          "service_attachment": "string"
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
      "engine_configs": {
        "description": "Optional. User-provided engine configurations for the instance.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "engine_version": {
        "computed": true,
        "description": "Optional. Engine version of the instance.",
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
      "instance_id": {
        "description": "Required. The ID to use for the instance, which will become the final component of\nthe instance's resource name.\n\nThis value is subject to the following restrictions:\n\n* Must be 4-63 characters in length\n* Must begin with a letter or digit\n* Must contain only lowercase letters, digits, and hyphens\n* Must not end with a hyphen\n* Must be unique within a location",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "labels": {
        "description": "Optional. Labels to represent user-provided metadata. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122. See documentation for resource type 'memorystore.googleapis.com/CertificateAuthority'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintenance_schedule": {
        "computed": true,
        "description": "Upcoming maintenance schedule.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "end_time": "string",
              "schedule_deadline_time": "string",
              "start_time": "string"
            }
          ]
        ]
      },
      "mode": {
        "computed": true,
        "description": "Optional. cluster or cluster-disabled. \n Possible values:\n CLUSTER\n CLUSTER_DISABLED Possible values: [\"CLUSTER\", \"CLUSTER_DISABLED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. Unique name of the instance.\nFormat: projects/{project}/locations/{location}/instances/{instance}",
        "description_kind": "plain",
        "type": "string"
      },
      "node_config": {
        "computed": true,
        "description": "Represents configuration for nodes of the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "size_gb": "number"
            }
          ]
        ]
      },
      "node_type": {
        "computed": true,
        "description": "Optional. Machine type for individual nodes of the instance. \n Possible values:\n SHARED_CORE_NANO\nHIGHMEM_MEDIUM\nHIGHMEM_XLARGE\nSTANDARD_SMALL",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_attachment_details": {
        "computed": true,
        "description": "Configuration of a service attachment of the cluster, for creating PSC connections.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_type": "string",
              "service_attachment": "string"
            }
          ]
        ]
      },
      "psc_auto_connections": {
        "computed": true,
        "deprecated": true,
        "description": "Output only. User inputs and resource details of the auto-created PSC connections.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_type": "string",
              "forwarding_rule": "string",
              "ip_address": "string",
              "network": "string",
              "port": "number",
              "project_id": "string",
              "psc_connection_id": "string",
              "psc_connection_status": "string",
              "service_attachment": "string"
            }
          ]
        ]
      },
      "replica_count": {
        "computed": true,
        "description": "Optional. Number of replica nodes per shard. If omitted the default is 0 replicas.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shard_count": {
        "description": "Required. Number of shards for the instance.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "Output only. Current state of the instance. \n Possible values:\n CREATING\nACTIVE\nUPDATING\nDELETING",
        "description_kind": "plain",
        "type": "string"
      },
      "state_info": {
        "computed": true,
        "description": "Additional information about the state of the instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "update_info": [
                "list",
                [
                  "object",
                  {
                    "target_engine_version": "string",
                    "target_node_type": "string",
                    "target_replica_count": "number",
                    "target_shard_count": "number"
                  }
                ]
              ]
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
      },
      "transit_encryption_mode": {
        "computed": true,
        "description": "Optional. Immutable. In-transit encryption mode of the instance. \n Possible values:\n TRANSIT_ENCRYPTION_DISABLED\nSERVER_AUTHENTICATION",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "Output only. System assigned, unique identifier for the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Output only. Latest update timestamp of the instance.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "automated_backup_config": {
        "block": {
          "attributes": {
            "retention": {
              "description": "How long to keep automated backups before the backups are deleted.\nThe value should be between 1 day and 365 days. If not specified, the default value is 35 days.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\". The default_value is \"3024000s\"",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "fixed_frequency_schedule": {
              "block": {
                "block_types": {
                  "start_time": {
                    "block": {
                      "attributes": {
                        "hours": {
                          "description": "Hours of a day in 24 hour format. Must be greater than or equal to 0 and typically must be less than or equal to 23.\nAn API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "The start time of every automated backup in UTC.\nIt must be set to the start of an hour. This field is required.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Trigger automated backups at a fixed frequency.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The automated backup config for a instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cross_instance_replication_config": {
        "block": {
          "attributes": {
            "instance_role": {
              "description": "The instance role supports the following values:\n1. 'INSTANCE_ROLE_UNSPECIFIED': This is an independent instance that has never participated in cross instance replication. It allows both reads and writes.\n2. 'NONE': This is an independent instance that previously participated in cross instance replication(either as a 'PRIMARY' or 'SECONDARY' cluster). It allows both reads and writes.\n3. 'PRIMARY': This instance serves as the replication source for secondary instance that are replicating from it. Any data written to it is automatically replicated to its secondary clusters. It allows both reads and writes.\n4. 'SECONDARY': This instance replicates data from the primary instance. It allows only reads. Possible values: [\"INSTANCE_ROLE_UNSPECIFIED\", \"NONE\", \"PRIMARY\", \"SECONDARY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "membership": {
              "computed": true,
              "description": "An output only view of all the member instance participating in cross instance replication. This field is populated for all the member clusters irrespective of their cluster role.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "primary_instance": [
                      "list",
                      [
                        "object",
                        {
                          "instance": "string",
                          "uid": "string"
                        }
                      ]
                    ],
                    "secondary_instance": [
                      "list",
                      [
                        "object",
                        {
                          "instance": "string",
                          "uid": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            },
            "update_time": {
              "computed": true,
              "description": "The last time cross instance replication config was updated.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "primary_instance": {
              "block": {
                "attributes": {
                  "instance": {
                    "description": "The full resource path of the primary instance in the format: projects/{project}/locations/{region}/instances/{instance-id}",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uid": {
                    "computed": true,
                    "description": "The unique id of the primary instance.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "This field is only set for a secondary instance. Details of the primary instance that is used as the replication source for this secondary instance. This is allowed to be set only for clusters whose cluster role is of type 'SECONDARY'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "secondary_instances": {
              "block": {
                "attributes": {
                  "instance": {
                    "description": "The full resource path of the Nth instance in the format: projects/{project}/locations/{region}/instance/{instance-id}",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uid": {
                    "computed": true,
                    "description": "The unique id of the Nth instance.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "List of secondary instances that are replicating from this primary cluster. This is allowed to be set only for instances whose cluster role is of type 'PRIMARY'.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Cross instance replication config",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "desired_auto_created_endpoints": {
        "block": {
          "attributes": {
            "network": {
              "description": "Required. The consumer network where the IP address resides, in the form of\nprojects/{project_id}/global/networks/{network_id}.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "project_id": {
              "description": "Required. The consumer project_id where the forwarding rule is created from.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Immutable. User inputs for the auto-created endpoints connections.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "desired_psc_auto_connections": {
        "block": {
          "attributes": {
            "network": {
              "description": "Required. The consumer network where the IP address resides, in the form of\nprojects/{project_id}/global/networks/{network_id}.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "project_id": {
              "description": "Required. The consumer project_id where the forwarding rule is created from.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description": "'desired_psc_auto_connections' is deprecated  Use 'desired_auto_created_endpoints' instead.",
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "gcs_source": {
        "block": {
          "attributes": {
            "uris": {
              "description": "URIs of the GCS objects to import.\nExample: gs://bucket1/object1, gs://bucket2/folder2/object2",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "GCS source for the instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "maintenance_policy": {
        "block": {
          "attributes": {
            "create_time": {
              "computed": true,
              "description": "The time when the policy was created.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
              "description_kind": "plain",
              "type": "string"
            },
            "update_time": {
              "computed": true,
              "description": "The time when the policy was last updated.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "weekly_maintenance_window": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "The day of week that maintenance updates occur.\n\n- DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.\n- MONDAY: Monday\n- TUESDAY: Tuesday\n- WEDNESDAY: Wednesday\n- THURSDAY: Thursday\n- FRIDAY: Friday\n- SATURDAY: Saturday\n- SUNDAY: Sunday Possible values: [\"DAY_OF_WEEK_UNSPECIFIED\", \"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
                    "computed": true,
                    "description": "Duration of the maintenance window.\nThe current window is fixed at 1 hour.\nA duration in seconds with up to nine fractional digits,\nterminated by 's'. Example: \"3.5s\".",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "block_types": {
                  "start_time": {
                    "block": {
                      "attributes": {
                        "hours": {
                          "description": "Hours of day in 24 hour format. Should be from 0 to 23.\nAn API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "minutes": {
                          "description": "Minutes of hour of day. Must be from 0 to 59.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "nanos": {
                          "description": "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "seconds": {
                          "description": "Seconds of minutes of the time. Must normally be from 0 to 59.\nAn API may allow the value 60 if it allows leap-seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description": "Start time of the window in UTC time.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "Optional. Maintenance window that is applied to resources covered by this policy.\nMinimum 1. For the current version, the maximum number\nof weekly_window is expected to be one.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Maintenance policy for a cluster",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "managed_backup_source": {
        "block": {
          "attributes": {
            "backup": {
              "description": "Example: 'projects/{project}/locations/{location}/backupCollections/{collection}/backups/{backup}'.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Managed backup source for the instance.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "persistence_config": {
        "block": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "Optional. Current persistence mode. \n Possible values:\nDISABLED\nRDB\nAOF Possible values: [\"DISABLED\", \"RDB\", \"AOF\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "aof_config": {
              "block": {
                "attributes": {
                  "append_fsync": {
                    "computed": true,
                    "description": "Optional. The fsync mode. \n Possible values:\n NEVER\nEVERY_SEC\nALWAYS",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for AOF based persistence.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "rdb_config": {
              "block": {
                "attributes": {
                  "rdb_snapshot_period": {
                    "computed": true,
                    "description": "Optional. Period between RDB snapshots. \n Possible values:\n ONE_HOUR\nSIX_HOURS\nTWELVE_HOURS\nTWENTY_FOUR_HOURS",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rdb_snapshot_start_time": {
                    "computed": true,
                    "description": "Optional. Time that the first snapshot was/will be attempted, and to which future\nsnapshots will be aligned. If not provided, the current time will be\nused.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "Configuration for RDB based persistence.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Represents persistence configuration for a instance.",
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
      },
      "zone_distribution_config": {
        "block": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "Optional. Current zone distribution mode. Defaults to MULTI_ZONE. \n Possible values:\n MULTI_ZONE\nSINGLE_ZONE Possible values: [\"MULTI_ZONE\", \"SINGLE_ZONE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zone": {
              "description": "Optional. Defines zone where all resources will be allocated with SINGLE_ZONE mode.\nIgnored for MULTI_ZONE mode.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Zone distribution configuration for allocation of instance resources.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleMemorystoreInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMemorystoreInstance), &result)
	return &result
}

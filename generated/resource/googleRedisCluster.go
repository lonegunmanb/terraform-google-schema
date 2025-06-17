package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleRedisCluster = `{
  "block": {
    "attributes": {
      "authorization_mode": {
        "description": "Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster. Default value: \"AUTH_MODE_DISABLED\" Possible values: [\"AUTH_MODE_UNSPECIFIED\", \"AUTH_MODE_IAM_AUTH\", \"AUTH_MODE_DISABLED\"]",
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
        "description": "The timestamp associated with the cluster creation request. A timestamp in\nRFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional\ndigits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_protection_enabled": {
        "description": "Optional. Indicates if the cluster is deletion protected or not.\nIf the value if set to true, any delete cluster operation will fail.\nDefault value is true.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "discovery_endpoints": {
        "computed": true,
        "description": "Output only. Endpoints created on each given network,\nfor Redis clients to connect to the cluster.\nCurrently only one endpoint is supported.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "port": "number",
              "psc_config": [
                "list",
                [
                  "object",
                  {
                    "network": "string"
                  }
                ]
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
      "kms_key": {
        "description": "The KMS key used to encrypt the at-rest data of the cluster.",
        "description_kind": "plain",
        "optional": true,
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
      "managed_server_ca": {
        "computed": true,
        "description": "Cluster's Certificate Authority. This field will only be populated if Redis Cluster's transit_encryption_mode is TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ca_certs": [
                "list",
                [
                  "object",
                  {
                    "certificates": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "name": {
        "computed": true,
        "description": "Unique name of the resource in this scope including project and location using the form:\nprojects/{projectId}/locations/{locationId}/clusters/{clusterId}",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description": "The nodeType for the Redis cluster.\nIf not provided, REDIS_HIGHMEM_MEDIUM will be used as default Possible values: [\"REDIS_SHARED_CORE_NANO\", \"REDIS_HIGHMEM_MEDIUM\", \"REDIS_HIGHMEM_XLARGE\", \"REDIS_STANDARD_SMALL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "precise_size_gb": {
        "computed": true,
        "description": "Output only. Redis memory precise size in GB for the entire cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_connections": {
        "computed": true,
        "description": "Output only. PSC connections for discovery of the cluster topology and accessing the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "address": "string",
              "forwarding_rule": "string",
              "network": "string",
              "project_id": "string",
              "psc_connection_id": "string"
            }
          ]
        ]
      },
      "psc_service_attachments": {
        "computed": true,
        "description": "Service attachment details to configure Psc connections.",
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
      "redis_configs": {
        "description": "Configure Redis Cluster behavior using a subset of native Redis configuration parameters.\nPlease check Memorystore documentation for the list of supported parameters:\nhttps://cloud.google.com/memorystore/docs/cluster/supported-instance-configurations",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "region": {
        "computed": true,
        "description": "The name of the region of the Redis cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_count": {
        "description": "Optional. The number of replica nodes per shard.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "shard_count": {
        "description": "Required. Number of shards for the Redis cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "size_gb": {
        "computed": true,
        "description": "Output only. Redis memory size in GB for the entire cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "state": {
        "computed": true,
        "description": "The current state of this cluster. Can be CREATING, READY, UPDATING, DELETING and SUSPENDED",
        "description_kind": "plain",
        "type": "string"
      },
      "state_info": {
        "computed": true,
        "description": "Output only. Additional information about the current state of the cluster.",
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
                    "target_replica_count": "number",
                    "target_shard_count": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "transit_encryption_mode": {
        "description": "Optional. The in-transit encryption for the Redis cluster.\nIf not provided, encryption is disabled for the cluster. Default value: \"TRANSIT_ENCRYPTION_MODE_DISABLED\" Possible values: [\"TRANSIT_ENCRYPTION_MODE_UNSPECIFIED\", \"TRANSIT_ENCRYPTION_MODE_DISABLED\", \"TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System assigned, unique identifier for the cluster.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "automated_backup_config": {
        "block": {
          "attributes": {
            "retention": {
              "description": "How long to keep automated backups before the backups are deleted.\nThe value should be between 1 day and 365 days. If not specified, the default value is 35 days.\nA duration in seconds with up to nine fractional digits, ending with 's'. Example: \"3.5s\".",
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
      "cross_cluster_replication_config": {
        "block": {
          "attributes": {
            "cluster_role": {
              "description": "The role of the cluster in cross cluster replication. Supported values are:\n\n1. 'CLUSTER_ROLE_UNSPECIFIED': This is an independent cluster that has never participated in cross cluster replication. It allows both reads and writes.\n\n1. 'NONE': This is an independent cluster that previously participated in cross cluster replication(either as a 'PRIMARY' or 'SECONDARY' cluster). It allows both reads and writes.\n\n1. 'PRIMARY': This cluster serves as the replication source for secondary clusters that are replicating from it. Any data written to it is automatically replicated to its secondary clusters. It allows both reads and writes.\n\n1. 'SECONDARY': This cluster replicates data from the primary cluster. It allows only reads. Possible values: [\"CLUSTER_ROLE_UNSPECIFIED\", \"NONE\", \"PRIMARY\", \"SECONDARY\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "membership": {
              "computed": true,
              "description": "An output only view of all the member clusters participating in cross cluster replication. This field is populated for all the member clusters irrespective of their cluster role.",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "primary_cluster": [
                      "list",
                      [
                        "object",
                        {
                          "cluster": "string",
                          "uid": "string"
                        }
                      ]
                    ],
                    "secondary_clusters": [
                      "list",
                      [
                        "object",
                        {
                          "cluster": "string",
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
              "description": "The last time cross cluster replication config was updated.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "primary_cluster": {
              "block": {
                "attributes": {
                  "cluster": {
                    "description": "The full resource path of the primary cluster in the format: projects/{project}/locations/{region}/clusters/{cluster-id}",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uid": {
                    "computed": true,
                    "description": "The unique id of the primary cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "Details of the primary cluster that is used as the replication source for this secondary cluster. This is allowed to be set only for clusters whose cluster role is of type 'SECONDARY'.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "secondary_clusters": {
              "block": {
                "attributes": {
                  "cluster": {
                    "description": "The full resource path of the secondary cluster in the format: projects/{project}/locations/{region}/clusters/{cluster-id}",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "uid": {
                    "computed": true,
                    "description": "The unique id of the secondary cluster.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description": "List of secondary clusters that are replicating from this primary cluster. This is allowed to be set only for clusters whose cluster role is of type 'PRIMARY'.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Cross cluster replication config",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "gcs_source": {
        "block": {
          "attributes": {
            "uris": {
              "description": "URIs of the GCS objects to import. Example: gs://bucket1/object1, gs://bucket2/folder2/object2",
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description": "Backups stored in Cloud Storage buckets. The Cloud Storage buckets need to be the same region as the clusters.",
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
              "description": "Output only. The time when the policy was created.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
              "description_kind": "plain",
              "type": "string"
            },
            "update_time": {
              "computed": true,
              "description": "Output only. The time when the policy was last updated.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond\nresolution and up to nine fractional digits.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "block_types": {
            "weekly_maintenance_window": {
              "block": {
                "attributes": {
                  "day": {
                    "description": "Required. The day of week that maintenance updates occur.\n\n- DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.\n- MONDAY: Monday\n- TUESDAY: Tuesday\n- WEDNESDAY: Wednesday\n- THURSDAY: Thursday\n- FRIDAY: Friday\n- SATURDAY: Saturday\n- SUNDAY: Sunday Possible values: [\"DAY_OF_WEEK_UNSPECIFIED\", \"MONDAY\", \"TUESDAY\", \"WEDNESDAY\", \"THURSDAY\", \"FRIDAY\", \"SATURDAY\", \"SUNDAY\"]",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
                    "computed": true,
                    "description": "Output only. Duration of the maintenance window.\nThe current window is fixed at 1 hour.\nA duration in seconds with up to nine fractional digits,\nterminated by 's'. Example: \"3.5s\".",
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
                      "description": "Required. Start time of the window in UTC time.",
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
          "description": "Backups that generated and managed by memorystore.",
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
              "description": "Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.\n\n- DISABLED: \tPersistence (both backup and restore) is disabled for the cluster.\n- RDB: RDB based Persistence is enabled.\n- AOF: AOF based Persistence is enabled. Possible values: [\"PERSISTENCE_MODE_UNSPECIFIED\", \"DISABLED\", \"RDB\", \"AOF\"]",
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
                    "description": "Optional. Available fsync modes.\n\n- NO - Do not explicitly call fsync(). Rely on OS defaults.\n- EVERYSEC - Call fsync() once per second in a background thread. A balance between performance and durability.\n- ALWAYS - Call fsync() for earch write command. Possible values: [\"APPEND_FSYNC_UNSPECIFIED\", \"NO\", \"EVERYSEC\", \"ALWAYS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "AOF configuration. This field will be ignored if mode is not AOF.",
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
                    "description": "Optional. Available snapshot periods for scheduling.\n\n- ONE_HOUR:\tSnapshot every 1 hour.\n- SIX_HOURS:\tSnapshot every 6 hours.\n- TWELVE_HOURS:\tSnapshot every 12 hours.\n- TWENTY_FOUR_HOURS:\tSnapshot every 24 hours. Possible values: [\"SNAPSHOT_PERIOD_UNSPECIFIED\", \"ONE_HOUR\", \"SIX_HOURS\", \"TWELVE_HOURS\", \"TWENTY_FOUR_HOURS\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rdb_snapshot_start_time": {
                    "computed": true,
                    "description": "The time that the first snapshot was/will be attempted, and to which\nfuture snapshots will be aligned.\nIf not provided, the current time will be used.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "RDB configuration. This field will be ignored if mode is not RDB.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Persistence config (RDB, AOF) for the cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "psc_configs": {
        "block": {
          "attributes": {
            "network": {
              "description": "Required. The consumer network where the network address of\nthe discovery endpoint will be reserved, in the form of\nprojects/{network_project_id_or_number}/global/networks/{network_id}.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "Required. Each PscConfig configures the consumer network where two\nnetwork addresses will be designated to the cluster for client access.\nCurrently, only one PscConfig is supported.",
          "description_kind": "plain"
        },
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
              "description": "Immutable. The mode for zone distribution for Memorystore Redis cluster.\nIf not provided, MULTI_ZONE will be used as default Possible values: [\"MULTI_ZONE\", \"SINGLE_ZONE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "zone": {
              "description": "Immutable. The zone for single zone Memorystore Redis cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Immutable. Zone distribution config for Memorystore Redis cluster.",
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

func GoogleRedisClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleRedisCluster), &result)
	return &result
}

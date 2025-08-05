package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleRedisCluster = `{
  "block": {
    "attributes": {
      "allow_fewer_zones_deployment": {
        "computed": true,
        "description": "Allows customers to specify if they are okay with deploying a multi-zone\ncluster in less than 3 zones. Once set, if there is a zonal outage during\nthe cluster creation, the cluster will only be deployed in 2 zones, and\nstay within the 2 zones for its lifecycle.",
        "description_kind": "plain",
        "type": "bool"
      },
      "authorization_mode": {
        "computed": true,
        "description": "Optional. The authorization mode of the Redis cluster. If not provided, auth feature is disabled for the cluster. Default value: \"AUTH_MODE_DISABLED\" Possible values: [\"AUTH_MODE_UNSPECIFIED\", \"AUTH_MODE_IAM_AUTH\", \"AUTH_MODE_DISABLED\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "automated_backup_config": {
        "computed": true,
        "description": "The automated backup config for a instance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fixed_frequency_schedule": [
                "list",
                [
                  "object",
                  {
                    "start_time": [
                      "list",
                      [
                        "object",
                        {
                          "hours": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "retention": "string"
            }
          ]
        ]
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
      "cross_cluster_replication_config": {
        "computed": true,
        "description": "Cross cluster replication config",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cluster_role": "string",
              "membership": [
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
              ],
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
              ],
              "update_time": "string"
            }
          ]
        ]
      },
      "deletion_protection_enabled": {
        "computed": true,
        "description": "Optional. Indicates if the cluster is deletion protected or not.\nIf the value if set to true, any delete cluster operation will fail.\nDefault value is true.",
        "description_kind": "plain",
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
      "gcs_source": {
        "computed": true,
        "description": "Backups stored in Cloud Storage buckets. The Cloud Storage buckets need to be the same region as the clusters.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "uris": [
                "set",
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
      "kms_key": {
        "computed": true,
        "description": "The KMS key used to encrypt the at-rest data of the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "maintenance_policy": {
        "computed": true,
        "description": "Maintenance policy for a cluster",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "create_time": "string",
              "update_time": "string",
              "weekly_maintenance_window": [
                "list",
                [
                  "object",
                  {
                    "day": "string",
                    "duration": "string",
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
      "managed_backup_source": {
        "computed": true,
        "description": "Backups that generated and managed by memorystore.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backup": "string"
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
        "description": "Unique name of the resource in this scope including project and location using the form:\nprojects/{projectId}/locations/{locationId}/clusters/{clusterId}",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_type": {
        "computed": true,
        "description": "The nodeType for the Redis cluster.\nIf not provided, REDIS_HIGHMEM_MEDIUM will be used as default Possible values: [\"REDIS_SHARED_CORE_NANO\", \"REDIS_HIGHMEM_MEDIUM\", \"REDIS_HIGHMEM_XLARGE\", \"REDIS_STANDARD_SMALL\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "persistence_config": {
        "computed": true,
        "description": "Persistence config (RDB, AOF) for the cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aof_config": [
                "list",
                [
                  "object",
                  {
                    "append_fsync": "string"
                  }
                ]
              ],
              "mode": "string",
              "rdb_config": [
                "list",
                [
                  "object",
                  {
                    "rdb_snapshot_period": "string",
                    "rdb_snapshot_start_time": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "precise_size_gb": {
        "computed": true,
        "description": "Output only. Redis memory precise size in GB for the entire cluster.",
        "description_kind": "plain",
        "type": "number"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "psc_configs": {
        "computed": true,
        "description": "Required. Each PscConfig configures the consumer network where two\nnetwork addresses will be designated to the cluster for client access.\nCurrently, only one PscConfig is supported.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "network": "string"
            }
          ]
        ]
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
        "computed": true,
        "description": "Configure Redis Cluster behavior using a subset of native Redis configuration parameters.\nPlease check Memorystore documentation for the list of supported parameters:\nhttps://cloud.google.com/memorystore/docs/cluster/supported-instance-configurations",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "region": {
        "description": "The name of the region of the Redis cluster.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "replica_count": {
        "computed": true,
        "description": "Optional. The number of replica nodes per shard.",
        "description_kind": "plain",
        "type": "number"
      },
      "shard_count": {
        "computed": true,
        "description": "Required. Number of shards for the Redis cluster.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "Optional. The in-transit encryption for the Redis cluster.\nIf not provided, encryption is disabled for the cluster. Default value: \"TRANSIT_ENCRYPTION_MODE_DISABLED\" Possible values: [\"TRANSIT_ENCRYPTION_MODE_UNSPECIFIED\", \"TRANSIT_ENCRYPTION_MODE_DISABLED\", \"TRANSIT_ENCRYPTION_MODE_SERVER_AUTHENTICATION\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System assigned, unique identifier for the cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "zone_distribution_config": {
        "computed": true,
        "description": "Immutable. Zone distribution config for Memorystore Redis cluster.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "mode": "string",
              "zone": "string"
            }
          ]
        ]
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

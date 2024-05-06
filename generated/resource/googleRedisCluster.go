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
      "create_time": {
        "computed": true,
        "description": "The timestamp associated with the cluster creation request. A timestamp in\nRFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine fractional\ndigits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
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

func GoogleRedisClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleRedisCluster), &result)
	return &result
}

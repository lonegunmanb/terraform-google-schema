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
            "list",
            "string"
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
        "description": "Optional. Immutable. Engine version of the instance.",
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
      "mode": {
        "computed": true,
        "description": "Optional. Standalone or cluster. \n Possible values:\n CLUSTER\nSTANDALONE Possible values: [\"CLUSTER\", \"STANDALONE\"]",
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
        "description": "Optional. Immutable. Machine type for individual nodes of the instance. \n Possible values:\n SHARED_CORE_NANO\nHIGHMEM_MEDIUM\nHIGHMEM_XLARGE\nSTANDARD_SMALL",
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
      "psc_auto_connections": {
        "computed": true,
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
          "description": "Required. Immutable. User inputs for the auto-created PSC connections.",
          "description_kind": "plain"
        },
        "min_items": 1,
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

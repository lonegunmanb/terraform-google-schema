package data

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
        "description": "Output only. Creation timestamp of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "cross_instance_replication_config": {
        "computed": true,
        "description": "Cross instance replication config",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "instance_role": "string",
              "membership": [
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
              ],
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
              "secondary_instances": [
                "list",
                [
                  "object",
                  {
                    "instance": "string",
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
        "description": "Optional. If set to true deletion of the instance will fail.",
        "description_kind": "plain",
        "type": "bool"
      },
      "desired_auto_created_endpoints": {
        "computed": true,
        "description": "Immutable. User inputs for the auto-created endpoints connections.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "network": "string",
              "project_id": "string"
            }
          ]
        ]
      },
      "desired_psc_auto_connections": {
        "computed": true,
        "description": "'desired_psc_auto_connections' is deprecated  Use 'desired_auto_created_endpoints' instead.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "network": "string",
              "project_id": "string"
            }
          ]
        ]
      },
      "discovery_endpoints": {
        "computed": true,
        "description": "Deprecated. Output only. Endpoints clients can connect to the instance through.",
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
        "computed": true,
        "description": "Optional. User-provided engine configurations for the instance.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "engine_version": {
        "computed": true,
        "description": "Optional. Engine version of the instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "gcs_source": {
        "computed": true,
        "description": "GCS source for the instance.",
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
      "instance_id": {
        "description": "Required. The ID to use for the instance, which will become the final component of\nthe instance's resource name.\n\nThis value is subject to the following restrictions:\n\n* Must be 4-63 characters in length\n* Must begin with a letter or digit\n* Must contain only lowercase letters, digits, and hyphens\n* Must not end with a hyphen\n* Must be unique within a location",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key": {
        "computed": true,
        "description": "The KMS key used to encrypt the at-rest data of the cluster",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "computed": true,
        "description": "Optional. Labels to represent user-provided metadata. \n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122. See documentation for resource type 'memorystore.googleapis.com/CertificateAuthority'.",
        "description_kind": "plain",
        "optional": true,
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
        "description": "Managed backup source for the instance.",
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
        "description": "Instance's Certificate Authority. This field will only be populated if instance's transit_encryption_mode is SERVER_AUTHENTICATION",
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
      "mode": {
        "computed": true,
        "description": "Optional. cluster or cluster-disabled. \n Possible values:\n CLUSTER\n CLUSTER_DISABLED Possible values: [\"CLUSTER\", \"CLUSTER_DISABLED\"]",
        "description_kind": "plain",
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
        "type": "string"
      },
      "persistence_config": {
        "computed": true,
        "description": "Represents persistence configuration for a instance.",
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
      "project": {
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
        "type": "number"
      },
      "shard_count": {
        "computed": true,
        "description": "Required. Number of shards for the instance.",
        "description_kind": "plain",
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
      },
      "zone_distribution_config": {
        "computed": true,
        "description": "Zone distribution configuration for allocation of instance resources.",
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

func GoogleMemorystoreInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleMemorystoreInstance), &result)
	return &result
}

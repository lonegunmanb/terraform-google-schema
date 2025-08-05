package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleManagedKafkaCluster = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description": "The ID to use for the cluster, which will become the final component of the cluster's name. The ID must be 1-63 characters long, and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' to comply with RFC 1035. This value is structured like: 'my-cluster-id'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the cluster was created.",
        "description_kind": "plain",
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "labels": {
        "description": "List of label KEY=VALUE pairs to add. Keys must start with a lowercase character and contain only hyphens (-), underscores ( ), lowercase characters, and numbers. Values must contain only hyphens (-), underscores ( ), lowercase characters, and numbers.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "ID of the location of the Kafka resource. See https://cloud.google.com/managed-kafka/docs/locations for a list of supported locations.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the cluster. Structured like: 'projects/PROJECT_ID/locations/LOCATION/clusters/CLUSTER_ID'.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the cluster. Possible values: 'STATE_UNSPECIFIED', 'CREATING', 'ACTIVE', 'DELETING'.",
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
      "update_time": {
        "computed": true,
        "description": "The time when the cluster was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "capacity_config": {
        "block": {
          "attributes": {
            "memory_bytes": {
              "description": "The memory to provision for the cluster in bytes. The value must be between 1 GiB and 8 GiB per vCPU. Ex. 1024Mi, 4Gi.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vcpu_count": {
              "description": "The number of vCPUs to provision for the cluster. The minimum is 3.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A capacity configuration of a Kafka cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "gcp_config": {
        "block": {
          "attributes": {
            "kms_key": {
              "description": "The Cloud KMS Key name to use for encryption. The key must be located in the same region as the cluster and cannot be changed. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "access_config": {
              "block": {
                "block_types": {
                  "network_configs": {
                    "block": {
                      "attributes": {
                        "subnet": {
                          "description": "Name of the VPC subnet from which the cluster is accessible. Both broker and bootstrap server IP addresses and DNS entries are automatically created in the subnet. There can only be one subnet per network, and the subnet must be located in the same region as the cluster. The project may differ. The name of the subnet must be in the format 'projects/PROJECT_ID/regions/REGION/subnetworks/SUBNET'.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Virtual Private Cloud (VPC) subnets where IP addresses for the Kafka cluster are allocated. To make the cluster available in a VPC, you must specify at least one 'network_configs' block. Max of 10 subnets per cluster. Additional subnets may be specified with additional 'network_configs' blocks.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration of access to the Kafka cluster.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "Configuration properties for a Kafka cluster deployed to Google Cloud Platform.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "rebalance_config": {
        "block": {
          "attributes": {
            "mode": {
              "description": "The rebalance behavior for the cluster. When not specified, defaults to 'NO_REBALANCE'. Possible values: 'MODE_UNSPECIFIED', 'NO_REBALANCE', 'AUTO_REBALANCE_ON_SCALE_UP'.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "Defines rebalancing behavior of a Kafka cluster.",
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
      "tls_config": {
        "block": {
          "attributes": {
            "ssl_principal_mapping_rules": {
              "description": "The rules for mapping mTLS certificate Distinguished Names (DNs) to shortened principal names for Kafka ACLs. This field corresponds exactly to the ssl.principal.mapping.rules broker config and matches the format and syntax defined in the Apache Kafka documentation. Setting or modifying this field will trigger a rolling restart of the Kafka brokers to apply the change. An empty string means that the default Kafka behavior is used. Example: 'RULE:^CN=(.?),OU=ServiceUsers.$/$1@example.com/,DEFAULT'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "trust_config": {
              "block": {
                "block_types": {
                  "cas_configs": {
                    "block": {
                      "attributes": {
                        "ca_pool": {
                          "description": "The name of the CA pool to pull CA certificates from. The CA pool does not need to be in the same project or location as the Kafka cluster. Must be in the format 'projects/PROJECT_ID/locations/LOCATION/caPools/CA_POOL_ID.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description": "Configuration for the Google Certificate Authority Service. To support mTLS, you must specify at least one 'cas_configs' block. A maximum of 10 CA pools can be specified. Additional CA pools may be specified with additional 'cas_configs' blocks.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "The configuration of the broker truststore. If specified, clients can use mTLS for authentication.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "TLS configuration for the Kafka cluster. This is used to configure mTLS authentication. To clear our a TLS configuration that has been previously set, please explicitly add an empty 'tls_config' block.",
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

func GoogleManagedKafkaClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleManagedKafkaCluster), &result)
	return &result
}

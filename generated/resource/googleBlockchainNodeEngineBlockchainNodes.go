package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBlockchainNodeEngineBlockchainNodes = `{
  "block": {
    "attributes": {
      "blockchain_node_id": {
        "description": "ID of the requesting object.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "blockchain_type": {
        "description": "User-provided key-value pairs Possible values: [\"ETHEREUM\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connection_info": {
        "computed": true,
        "description": "The connection information through which to interact with a blockchain node.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "endpoint_info": [
                "list",
                [
                  "object",
                  {
                    "json_rpc_api_endpoint": "string",
                    "websockets_api_endpoint": "string"
                  }
                ]
              ],
              "service_attachment": "string"
            }
          ]
        ]
      },
      "create_time": {
        "computed": true,
        "description": "The timestamp at which the blockchain node was first created.",
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
        "description": "User-provided key-value pairs\n\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "Location of Blockchain Node being created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The fully qualified name of the blockchain node. e.g. projects/my-project/locations/us-central1/blockchainNodes/my-node.",
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
      },
      "update_time": {
        "computed": true,
        "description": "The timestamp at which the blockchain node was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "ethereum_details": {
        "block": {
          "attributes": {
            "additional_endpoints": {
              "computed": true,
              "description": "User-provided key-value pairs",
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "beacon_api_endpoint": "string",
                    "beacon_prometheus_metrics_api_endpoint": "string",
                    "execution_client_prometheus_metrics_api_endpoint": "string"
                  }
                ]
              ]
            },
            "api_enable_admin": {
              "description": "Enables JSON-RPC access to functions in the admin namespace. Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "api_enable_debug": {
              "description": "Enables JSON-RPC access to functions in the debug namespace. Defaults to false.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "consensus_client": {
              "description": "The consensus client Possible values: [\"CONSENSUS_CLIENT_UNSPECIFIED\", \"LIGHTHOUSE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "execution_client": {
              "description": "The execution client Possible values: [\"EXECUTION_CLIENT_UNSPECIFIED\", \"GETH\", \"ERIGON\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "network": {
              "description": "The Ethereum environment being accessed. Possible values: [\"MAINNET\", \"TESTNET_GOERLI_PRATER\", \"TESTNET_SEPOLIA\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "node_type": {
              "description": "The type of Ethereum node. Possible values: [\"LIGHT\", \"FULL\", \"ARCHIVE\"]",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "geth_details": {
              "block": {
                "attributes": {
                  "garbage_collection_mode": {
                    "description": "Blockchain garbage collection modes. Only applicable when NodeType is FULL or ARCHIVE. Possible values: [\"FULL\", \"ARCHIVE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "User-provided key-value pairs",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "validator_config": {
              "block": {
                "attributes": {
                  "mev_relay_urls": {
                    "description": "URLs for MEV-relay services to use for block building. When set, a managed MEV-boost service is configured on the beacon client.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description": "Configuration for validator-related parameters on the beacon client, and for any managed validator client.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "User-provided key-value pairs",
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

func GoogleBlockchainNodeEngineBlockchainNodesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBlockchainNodeEngineBlockchainNodes), &result)
	return &result
}

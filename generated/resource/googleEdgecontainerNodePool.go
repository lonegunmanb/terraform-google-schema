package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleEdgecontainerNodePool = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "The name of the target Distributed Cloud Edge Cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time when the node pool was created.",
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
        "description": "Labels associated with this resource.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "machine_filter": {
        "computed": true,
        "description": "Only machines matching this filter will be allowed to join the node pool.\nThe filtering language accepts strings like \"name=\u003cname\u003e\", and is\ndocumented in more detail in [AIP-160](https://google.aip.dev/160).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name of the node pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_count": {
        "description": "The number of nodes in the pool.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "node_location": {
        "description": "Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: 'us-central1-edge-customer-a'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_version": {
        "computed": true,
        "description": "The lowest release version among all worker nodes.",
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
        "description": "The time when the node pool was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "local_disk_encryption": {
        "block": {
          "attributes": {
            "kms_key": {
              "description": "The Cloud KMS CryptoKey e.g. projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey} to use for protecting node local disks.\nIf not specified, a Google-managed key will be used instead.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kms_key_active_version": {
              "computed": true,
              "description": "The Cloud KMS CryptoKeyVersion currently in use for protecting node local disks. Only applicable if kmsKey is set.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_state": {
              "computed": true,
              "description": "Availability of the Cloud KMS CryptoKey. If not KEY_AVAILABLE, then nodes may go offline as they cannot access their local data.\nThis can be caused by a lack of permissions to use the key, or if the key is disabled or deleted.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Local disk encryption options. This field is only used when enabling CMEK support.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_config": {
        "block": {
          "attributes": {
            "labels": {
              "computed": true,
              "description": "\"The Kubernetes node labels\"",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description": "Configuration for each node in the NodePool",
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

func GoogleEdgecontainerNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleEdgecontainerNodePool), &result)
	return &result
}

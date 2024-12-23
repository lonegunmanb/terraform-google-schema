package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeonpremBareMetalNodePool = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Annotations on the Bare Metal Node Pool.\nThis field has the same restrictions as Kubernetes annotations.\nThe total size of all keys and values combined is limited to 256k.\nKey can have 2 segments: prefix (optional) and name (required),\nseparated by a slash (/).\nPrefix must be a DNS subdomain.\nName must be 63 characters or less, begin and end with alphanumerics,\nwith dashes (-), underscores (_), dots (.), and alphanumerics between.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "bare_metal_cluster": {
        "description": "The cluster this node pool belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "The time the cluster was created, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "delete_time": {
        "computed": true,
        "description": "The time the cluster was deleted, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "description": "The display name for the Bare Metal Node Pool.",
        "description_kind": "plain",
        "optional": true,
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
      "etag": {
        "computed": true,
        "description": "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.\nAllows clients to perform consistent read-modify-writes\nthrough optimistic concurrency control.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location of the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The bare metal node pool name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "reconciling": {
        "computed": true,
        "description": "If set, there are currently changes in flight to the Bare Metal User Cluster.",
        "description_kind": "plain",
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "The current state of this cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "Specifies detailed node pool status.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "conditions": [
                "list",
                [
                  "object",
                  {
                    "last_transition_time": "string",
                    "message": "string",
                    "reason": "string",
                    "state": "string",
                    "type": "string"
                  }
                ]
              ],
              "error_message": "string"
            }
          ]
        ]
      },
      "uid": {
        "computed": true,
        "description": "The unique identifier of the Bare Metal Node Pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time the cluster was last updated, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "node_pool_config": {
        "block": {
          "attributes": {
            "labels": {
              "computed": true,
              "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nFor example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "operating_system": {
              "computed": true,
              "description": "Specifies the nodes operating system (default: LINUX).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "node_configs": {
              "block": {
                "attributes": {
                  "labels": {
                    "description": "The map of Kubernetes labels (key/value pairs) to be applied to\neach node. These will added in addition to any default label(s)\nthat Kubernetes may apply to the node. In case of conflict in\nlabel keys, the applied set may differ depending on the Kubernetes\nversion -- it's best to assume the behavior is undefined and\nconflicts should be avoided. For more information, including usage\nand the valid values, see:\n  - http://kubernetes.io/v1.1/docs/user-guide/labels.html\nAn object containing a list of \"key\": value pairs.\nFor example: { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "node_ip": {
                    "description": "The default IPv4 address for SSH access and Kubernetes node.\nExample: 192.168.0.1",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The list of machine addresses in the Bare Metal Node Pool.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "taints": {
              "block": {
                "attributes": {
                  "effect": {
                    "description": "Specifies the nodes operating system (default: LINUX). Possible values: [\"EFFECT_UNSPECIFIED\", \"PREFER_NO_SCHEDULE\", \"NO_EXECUTE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Key associated with the effect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value associated with the effect.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The initial taints assigned to nodes of this node pool.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description": "Node pool configuration.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleGkeonpremBareMetalNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeonpremBareMetalNodePool), &result)
	return &result
}

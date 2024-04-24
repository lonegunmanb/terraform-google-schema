package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleGkeonpremVmwareNodePool = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "Annotations on the node Pool.\nThis field has the same restrictions as Kubernetes annotations.\nThe total size of all keys and values combined is limited to 256k.\nKey can have 2 segments: prefix (optional) and name (required),\nseparated by a slash (/).\nPrefix must be a DNS subdomain.\nName must be 63 characters or less, begin and end with alphanumerics,\nwith dashes (-), underscores (_), dots (.), and alphanumerics between.\n\n\n**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.\nPlease refer to the field 'effective_annotations' for all of the annotations present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
        "description": "The display name for the node pool.",
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
        "description": "The vmware node pool name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "on_prem_version": {
        "computed": true,
        "description": "Anthos version for the node pool. Defaults to the user cluster version.",
        "description_kind": "plain",
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
        "description": "If set, there are currently changes in flight to the node pool.",
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
        "description": "ResourceStatus representing detailed cluster state.",
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
        "description": "The unique identifier of the node pool.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time the cluster was last updated, in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "vmware_cluster": {
        "description": "The cluster this node pool belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "config": {
        "block": {
          "attributes": {
            "boot_disk_size_gb": {
              "description": "VMware disk size to be used during creation.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "cpus": {
              "description": "The number of CPUs for each node in the node pool.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "enable_load_balancer": {
              "description": "Allow node pool traffic to be load balanced. Only works for clusters with\nMetalLB load balancers.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "image": {
              "description": "The OS image name in vCenter, only valid when using Windows.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "image_type": {
              "description": "The OS image to be used for each node in a node pool.\nCurrently 'cos', 'ubuntu', 'ubuntu_containerd' and 'windows' are supported.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "labels": {
              "computed": true,
              "description": "The map of Kubernetes labels (key/value pairs) to be applied to each node.\nThese will added in addition to any default label(s) that\nKubernetes may apply to the node.\nIn case of conflict in label keys, the applied set may differ depending on\nthe Kubernetes version -- it's best to assume the behavior is undefined\nand conflicts should be avoided.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "memory_mb": {
              "description": "The megabytes of memory for each node in the node pool.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replicas": {
              "description": "The number of nodes in the node pool.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "taints": {
              "block": {
                "attributes": {
                  "effect": {
                    "description": "Available taint effects. Possible values: [\"EFFECT_UNSPECIFIED\", \"NO_SCHEDULE\", \"PREFER_NO_SCHEDULE\", \"NO_EXECUTE\"]",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "key": {
                    "description": "Key associated with the effect.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "Value associated with the effect.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description": "The initial taints assigned to nodes of this node pool.",
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "vsphere_config": {
              "block": {
                "attributes": {
                  "datastore": {
                    "description": "The name of the vCenter datastore. Inherited from the user cluster.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "host_groups": {
                    "description": "Vsphere host groups to apply to all VMs in the node pool",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "tags": {
                    "block": {
                      "attributes": {
                        "category": {
                          "description": "The Vsphere tag category.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "tag": {
                          "description": "The Vsphere tag name.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description": "Tags to apply to VMs.",
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  }
                },
                "description": "Specifies the vSphere config for node pool.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The node configuration of the node pool.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "node_pool_autoscaling": {
        "block": {
          "attributes": {
            "max_replicas": {
              "description": "Maximum number of replicas in the NodePool.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_replicas": {
              "description": "Minimum number of replicas in the NodePool.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description": "Node Pool autoscaling config for the node pool.",
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

func GoogleGkeonpremVmwareNodePoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleGkeonpremVmwareNodePool), &result)
	return &result
}

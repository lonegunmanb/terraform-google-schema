package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineCluster = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "management": {
        "computed": true,
        "description": "True if the cluster is a management cluster; false otherwise.\nThere can only be one management cluster in a private cloud and it has to be the first one.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description": "The ID of the Cluster.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The resource name of the private cloud to create a new cluster in.\nResource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\nFor example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "State of the Cluster.",
        "description_kind": "plain",
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "autoscaling_settings": {
        "block": {
          "attributes": {
            "cool_down_period": {
              "description": "The minimum duration between consecutive autoscale operations.\nIt starts once addition or removal of nodes is fully completed.\nMinimum cool down period is 30m.\nCool down period must be in whole minutes (for example, 30m, 31m, 50m).\nMandatory for successful addition of autoscaling settings in cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_cluster_node_count": {
              "description": "Maximum number of nodes of any type in a cluster.\nMandatory for successful addition of autoscaling settings in cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_cluster_node_count": {
              "description": "Minimum number of nodes of any type in a cluster.\nMandatory for successful addition of autoscaling settings in cluster.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "autoscaling_policies": {
              "block": {
                "attributes": {
                  "autoscale_policy_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "node_type_id": {
                    "description": "The canonical identifier of the node type to add or remove.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "scale_out_size": {
                    "description": "Number of nodes to add to a cluster during a scale-out operation.\nMust be divisible by 2 for stretched clusters.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "consumed_memory_thresholds": {
                    "block": {
                      "attributes": {
                        "scale_in": {
                          "description": "The utilization triggering the scale-in operation in percent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "scale_out": {
                          "description": "The utilization triggering the scale-out operation in percent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Utilization thresholds pertaining to amount of consumed memory.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "cpu_thresholds": {
                    "block": {
                      "attributes": {
                        "scale_in": {
                          "description": "The utilization triggering the scale-in operation in percent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "scale_out": {
                          "description": "The utilization triggering the scale-out operation in percent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Utilization thresholds pertaining to CPU utilization.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "storage_thresholds": {
                    "block": {
                      "attributes": {
                        "scale_in": {
                          "description": "The utilization triggering the scale-in operation in percent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "scale_out": {
                          "description": "The utilization triggering the scale-out operation in percent.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        }
                      },
                      "description": "Utilization thresholds pertaining to amount of consumed storage.",
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description": "The map with autoscaling policies applied to the cluster.\nThe key is the identifier of the policy.\nIt must meet the following requirements:\n  * Only contains 1-63 alphanumeric characters and hyphens\n  * Begins with an alphabetical character\n  * Ends with a non-hyphen character\n  * Not formatted as a UUID\n  * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)\n\nCurrently the map must contain only one element\nthat describes the autoscaling policy for compute nodes.",
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description": "Configuration of the autoscaling applied to this cluster",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "node_type_configs": {
        "block": {
          "attributes": {
            "custom_core_count": {
              "description": "Customized number of cores available to each node of the type.\nThis number must always be one of 'nodeType.availableCustomCoreCounts'.\nIf zero is provided max value from 'nodeType.availableCustomCoreCounts' will be used.\nOnce the customer is created then corecount cannot be changed.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "node_count": {
              "description": "The number of nodes of this type in the cluster.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "node_type_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "The map of cluster node types in this cluster,\nwhere the key is canonical identifier of the node type (corresponds to the NodeType).",
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func GoogleVmwareengineClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineCluster), &result)
	return &result
}

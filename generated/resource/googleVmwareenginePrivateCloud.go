package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareenginePrivateCloud = `{
  "block": {
    "attributes": {
      "deletion_delay_hours": {
        "description": "The number of hours to delay this request. You can set this value to an hour between 0 to 8, where setting it to 0 starts the deletion request immediately. If no value is set, a default value is set at the API Level.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "description": {
        "description": "User-provided description for this private cloud.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hcx": {
        "computed": true,
        "description": "Details about a HCX Cloud Manager appliance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fqdn": "string",
              "internal_ip": "string",
              "state": "string",
              "version": "string"
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
      "location": {
        "description": "The location where the PrivateCloud should reside.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The ID of the PrivateCloud.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nsx": {
        "computed": true,
        "description": "Details about a NSX Manager appliance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fqdn": "string",
              "internal_ip": "string",
              "state": "string",
              "version": "string"
            }
          ]
        ]
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "send_deletion_delay_hours_if_zero": {
        "description": "While set true, deletion_delay_hours value will be sent in the request even for zero value of the field. This field is only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with deletion_delay_hours.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "state": {
        "computed": true,
        "description": "State of the resource. New values may be added to this enum when appropriate.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "Initial type of the private cloud. Possible values: [\"STANDARD\", \"TIME_LIMITED\", \"STRETCHED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "uid": {
        "computed": true,
        "description": "System-generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "vcenter": {
        "computed": true,
        "description": "Details about a vCenter Server management appliance.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fqdn": "string",
              "internal_ip": "string",
              "state": "string",
              "version": "string"
            }
          ]
        ]
      }
    },
    "block_types": {
      "management_cluster": {
        "block": {
          "attributes": {
            "cluster_id": {
              "description": "The user-provided identifier of the new Cluster. The identifier must meet the following requirements:\n  * Only contains 1-63 alphanumeric characters and hyphens\n  * Begins with an alphabetical character\n  * Ends with a non-hyphen character\n  * Not formatted as a UUID\n  * Complies with RFC 1034 (https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)",
              "description_kind": "plain",
              "required": true,
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
                      "description": "The map with autoscaling policies applied to the cluster.\nThe key is the identifier of the policy.\nIt must meet the following requirements:\n * Only contains 1-63 alphanumeric characters and hyphens\n * Begins with an alphabetical character\n * Ends with a non-hyphen character\n * Not formatted as a UUID\n * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)\n\nCurrently the map must contain only one element\nthat describes the autoscaling policy for compute nodes.",
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description": "Configuration of the autoscaling applied to this cluster\nPrivate cloud must have a minimum of 3 nodes to add autoscale settings",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "node_type_configs": {
              "block": {
                "attributes": {
                  "custom_core_count": {
                    "description": "Customized number of cores available to each node of the type.\nThis number must always be one of 'nodeType.availableCustomCoreCounts'.\nIf zero is provided max value from 'nodeType.availableCustomCoreCounts' will be used.\nThis cannot be changed once the PrivateCloud is created.",
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
            "stretched_cluster_config": {
              "block": {
                "attributes": {
                  "preferred_location": {
                    "description": "Zone that will remain operational when connection between the two zones is lost.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "secondary_location": {
                    "description": "Additional zone for a higher level of availability and load balancing.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description": "The stretched cluster configuration for the private cloud.",
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description": "The management cluster for this private cloud. This used for creating and managing the default cluster.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "network_config": {
        "block": {
          "attributes": {
            "dns_server_ip": {
              "computed": true,
              "description": "DNS Server IP of the Private Cloud.",
              "description_kind": "plain",
              "type": "string"
            },
            "management_cidr": {
              "description": "Management CIDR used by VMware management appliances.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "management_ip_address_layout_version": {
              "computed": true,
              "description": "The IP address layout version of the management IP address range.\nPossible versions include:\n* managementIpAddressLayoutVersion=1: Indicates the legacy IP address layout used by some existing private clouds. This is no longer supported for new private clouds\nas it does not support all features.\n* managementIpAddressLayoutVersion=2: Indicates the latest IP address layout\nused by all newly created private clouds. This version supports all current features.",
              "description_kind": "plain",
              "type": "number"
            },
            "vmware_engine_network": {
              "description": "The relative resource name of the VMware Engine network attached to the private cloud.\nSpecify the name in the following form: projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}\nwhere {project} can either be a project number or a project ID.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vmware_engine_network_canonical": {
              "computed": true,
              "description": "The canonical name of the VMware Engine network in\nthe form: projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Network configuration in the consumer project with which the peering has to be done.",
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

func GoogleVmwareenginePrivateCloudSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareenginePrivateCloud), &result)
	return &result
}

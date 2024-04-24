package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareenginePrivateCloud = `{
  "block": {
    "attributes": {
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

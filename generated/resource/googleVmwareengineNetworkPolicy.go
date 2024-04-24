package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineNetworkPolicy = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation time of this resource.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and\nup to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-provided description for this network policy.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_services_cidr": {
        "description": "IP address range in CIDR notation used to create internet access and external IP access.\nAn RFC 1918 CIDR block, with a \"/26\" prefix, is required. The range cannot overlap with any\nprefixes either in the consumer VPC network or in use by the private clouds attached to that VPC network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The resource name of the location (region) to create the new network policy in.\nResource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\nFor example: projects/my-project/locations/us-central1",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The ID of the Network Policy.",
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
      "uid": {
        "computed": true,
        "description": "System-generated unique identifier for the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "Last updated time of this resource.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and up to nine\nfractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "vmware_engine_network": {
        "description": "The relative resource name of the VMware Engine network. Specify the name in the following form:\nprojects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}\ncan either be a project number or a project ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vmware_engine_network_canonical": {
        "computed": true,
        "description": "The canonical name of the VMware Engine network in the form:\nprojects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "external_ip": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "True if the service is enabled; false otherwise.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "state": {
              "computed": true,
              "description": "State of the service. New values may be added to this enum when appropriate.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Network service that allows External IP addresses to be assigned to VMware workloads.\nThis service can only be enabled when internetAccess is also enabled.",
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "internet_access": {
        "block": {
          "attributes": {
            "enabled": {
              "description": "True if the service is enabled; false otherwise.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "state": {
              "computed": true,
              "description": "State of the service. New values may be added to this enum when appropriate.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Network service that allows VMware workloads to access the internet.",
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

func GoogleVmwareengineNetworkPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineNetworkPolicy), &result)
	return &result
}

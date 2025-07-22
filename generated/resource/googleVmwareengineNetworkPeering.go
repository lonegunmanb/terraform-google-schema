package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineNetworkPeering = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "Creation time of this resource.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and\nup to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-provided description for this network peering.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "export_custom_routes": {
        "description": "True if custom routes are exported to the peered network; false otherwise.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "export_custom_routes_with_public_ip": {
        "description": "True if all subnet routes with a public IP address range are exported; false otherwise.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "import_custom_routes": {
        "description": "True if custom routes are imported from the peered network; false otherwise.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "import_custom_routes_with_public_ip": {
        "description": "True if custom routes are imported from the peered network; false otherwise.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description": "The ID of the Network Peering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_network": {
        "description": "The relative resource name of the network to peer with a standard VMware Engine network.\nThe provided network can be a consumer VPC network or another standard VMware Engine network.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_network_type": {
        "description": "The type of the network to peer with the VMware Engine network. Possible values: [\"STANDARD\", \"VMWARE_ENGINE_NETWORK\", \"PRIVATE_SERVICES_ACCESS\", \"NETAPP_CLOUD_VOLUMES\", \"THIRD_PARTY_SERVICE\", \"DELL_POWERSCALE\", \"GOOGLE_CLOUD_NETAPP_VOLUMES\"]",
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
      "state": {
        "computed": true,
        "description": "State of the network peering.\nThis field has a value of 'ACTIVE' when there's a matching configuration in the peer network.\nNew values may be added to this enum when appropriate.",
        "description_kind": "plain",
        "type": "string"
      },
      "state_details": {
        "computed": true,
        "description": "Details about the current state of the network peering.",
        "description_kind": "plain",
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

func GoogleVmwareengineNetworkPeeringSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineNetworkPeering), &result)
	return &result
}

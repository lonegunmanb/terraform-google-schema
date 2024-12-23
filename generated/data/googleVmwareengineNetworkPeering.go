package data

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
        "computed": true,
        "description": "User-provided description for this network peering.",
        "description_kind": "plain",
        "type": "string"
      },
      "export_custom_routes": {
        "computed": true,
        "description": "True if custom routes are exported to the peered network; false otherwise.",
        "description_kind": "plain",
        "type": "bool"
      },
      "export_custom_routes_with_public_ip": {
        "computed": true,
        "description": "True if all subnet routes with a public IP address range are exported; false otherwise.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "import_custom_routes": {
        "computed": true,
        "description": "True if custom routes are imported from the peered network; false otherwise.",
        "description_kind": "plain",
        "type": "bool"
      },
      "import_custom_routes_with_public_ip": {
        "computed": true,
        "description": "True if custom routes are imported from the peered network; false otherwise.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description": "The ID of the Network Peering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_network": {
        "computed": true,
        "description": "The relative resource name of the network to peer with a standard VMware Engine network.\nThe provided network can be a consumer VPC network or another standard VMware Engine network.",
        "description_kind": "plain",
        "type": "string"
      },
      "peer_network_type": {
        "computed": true,
        "description": "The type of the network to peer with the VMware Engine network. Possible values: [\"STANDARD\", \"VMWARE_ENGINE_NETWORK\", \"PRIVATE_SERVICES_ACCESS\", \"NETAPP_CLOUD_VOLUMES\", \"THIRD_PARTY_SERVICE\", \"DELL_POWERSCALE\"]",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
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
        "computed": true,
        "description": "The relative resource name of the VMware Engine network. Specify the name in the following form:\nprojects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}\ncan either be a project number or a project ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "vmware_engine_network_canonical": {
        "computed": true,
        "description": "The canonical name of the VMware Engine network in the form:\nprojects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}",
        "description_kind": "plain",
        "type": "string"
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

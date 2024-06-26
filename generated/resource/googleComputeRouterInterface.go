package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRouterInterface = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "interconnect_attachment": {
        "description": "The name or resource link to the VLAN interconnect for this interface. Changing this forces a new interface to be created. Only one of interconnect_attachment, subnetwork or vpn_tunnel can be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_range": {
        "computed": true,
        "description": "The IP address and range of the interface. The IP range must be in the RFC3927 link-local IP space. Changing this forces a new interface to be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_version": {
        "computed": true,
        "description": "IP version of this interface.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "A unique name for the interface, required by GCE. Changing this forces a new interface to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "private_ip_address": {
        "description": "The regional private internal IP address that is used to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which this interface's router belongs. If it is not provided, the provider project is used. Changing this forces a new interface to be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redundant_interface": {
        "computed": true,
        "description": "The name of the interface that is redundant to this interface. Changing this forces a new interface to be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region this interface's router sits in. If not specified, the project region will be used. Changing this forces a new interface to be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "router": {
        "description": "The name of the router this interface will be attached to. Changing this forces a new interface to be created.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnetwork": {
        "description": "The URI of the subnetwork resource that this interface belongs to, which must be in the same region as the Cloud Router. Changing this forces a new interface to be created. Only one of subnetwork, interconnect_attachment or vpn_tunnel can be specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_tunnel": {
        "description": "The name or resource link to the VPN tunnel this interface will be linked to. Changing this forces a new interface to be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.",
        "description_kind": "plain",
        "optional": true,
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

func GoogleComputeRouterInterfaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRouterInterface), &result)
	return &result
}

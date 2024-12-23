package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeVpnTunnel = `{
  "block": {
    "attributes": {
      "creation_timestamp": {
        "computed": true,
        "description": "Creation timestamp in RFC3339 text format.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "An optional description of this resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "detailed_status": {
        "computed": true,
        "description": "Detailed status message for the VPN tunnel.",
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
      "ike_version": {
        "description": "IKE protocol version to use when establishing the VPN tunnel with\npeer VPN gateway.\nAcceptable IKE versions are 1 or 2. Default version is 2.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "label_fingerprint": {
        "computed": true,
        "description": "The fingerprint used for optimistic locking of this resource.  Used\ninternally during updates.",
        "description_kind": "plain",
        "type": "string"
      },
      "labels": {
        "description": "Labels to apply to this VpnTunnel.\n\n**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.\nPlease refer to the field 'effective_labels' for all of the labels present on the resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "local_traffic_selector": {
        "computed": true,
        "description": "Local traffic selector to use when establishing the VPN tunnel with\npeer VPN gateway. The value should be a CIDR formatted string,\nfor example '192.168.0.0/16'. The ranges should be disjoint.\nOnly IPv4 is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "name": {
        "description": "Name of the resource. The name must be 1-63 characters long, and\ncomply with RFC1035. Specifically, the name must be 1-63\ncharacters long and match the regular expression\n'[a-z]([-a-z0-9]*[a-z0-9])?' which means the first character\nmust be a lowercase letter, and all following characters must\nbe a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peer_external_gateway": {
        "description": "URL of the peer side external VPN gateway to which this VPN tunnel is connected.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_external_gateway_interface": {
        "description": "The interface ID of the external VPN gateway to which this VPN tunnel is connected.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "peer_gcp_gateway": {
        "description": "URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.\nIf provided, the VPN tunnel will automatically use the same vpn_gateway_interface\nID in the peer GCP VPN gateway.\nThis field must reference a 'google_compute_ha_vpn_gateway' resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "peer_ip": {
        "computed": true,
        "description": "IP address of the peer VPN gateway. Only IPv4 is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The region where the tunnel is located. If unset, is set to the region of 'target_vpn_gateway'.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "remote_traffic_selector": {
        "computed": true,
        "description": "Remote traffic selector to use when establishing the VPN tunnel with\npeer VPN gateway. The value should be a CIDR formatted string,\nfor example '192.168.0.0/16'. The ranges should be disjoint.\nOnly IPv4 is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "router": {
        "description": "URL of router resource to be used for dynamic routing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "shared_secret": {
        "description": "Shared secret used to set the secure session between the Cloud VPN\ngateway and the peer VPN gateway.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "shared_secret_hash": {
        "computed": true,
        "description": "Hash of the shared secret.",
        "description_kind": "plain",
        "type": "string"
      },
      "target_vpn_gateway": {
        "description": "URL of the Target VPN gateway with which this VPN tunnel is\nassociated.",
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
      "tunnel_id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "vpn_gateway": {
        "description": "URL of the VPN gateway with which this VPN tunnel is associated.\nThis must be used if a High Availability VPN gateway resource is created.\nThis field must reference a 'google_compute_ha_vpn_gateway' resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vpn_gateway_interface": {
        "description": "The interface ID of the VPN gateway with which this VPN tunnel is associated.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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

func GoogleComputeVpnTunnelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeVpnTunnel), &result)
	return &result
}

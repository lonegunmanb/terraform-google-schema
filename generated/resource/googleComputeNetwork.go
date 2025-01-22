package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeNetwork = `{
  "block": {
    "attributes": {
      "auto_create_subnetworks": {
        "description": "When set to 'true', the network is created in \"auto subnet mode\" and\nit will create a subnet for each region automatically across the\n'10.128.0.0/9' address range.\n\nWhen set to 'false', the network is created in \"custom subnet mode\" so\nthe user can explicitly connect subnetwork resources.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bgp_always_compare_med": {
        "computed": true,
        "description": "Enables/disables the comparison of MED across routes with different Neighbor ASNs.\nThis value can only be set if the --bgp-best-path-selection-mode is STANDARD",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "bgp_best_path_selection_mode": {
        "computed": true,
        "description": "The BGP best selection algorithm to be employed. MODE can be LEGACY or STANDARD. Possible values: [\"LEGACY\", \"STANDARD\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "bgp_inter_region_cost": {
        "computed": true,
        "description": "Choice of the behavior of inter-regional cost and MED in the BPS algorithm. Possible values: [\"DEFAULT\", \"ADD_COST_TO_MED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delete_default_routes_on_create": {
        "description": "If set to 'true', default routes ('0.0.0.0/0') will be deleted\nimmediately after network creation. Defaults to 'false'.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description": "An optional description of this resource. The resource must be\nrecreated to modify this field.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enable_ula_internal_ipv6": {
        "description": "Enable ULA internal ipv6 on this network. Enabling this feature will assign\na /48 from google defined ULA prefix fd20::/20.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateway_ipv4": {
        "computed": true,
        "description": "The gateway address for default routing out of the network. This value\nis selected by GCP.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internal_ipv6_range": {
        "computed": true,
        "description": "When enabling ula internal ipv6, caller optionally can specify the /48 range\nthey want from the google defined ULA prefix fd20::/20. The input must be a\nvalid /48 ULA IPv6 address and must be within the fd20::/20. Operation will\nfail if the speficied /48 is already in used by another resource.\nIf the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mtu": {
        "computed": true,
        "description": "Maximum Transmission Unit in bytes. The default value is 1460 bytes.\nThe minimum value for this field is 1300 and the maximum value is 8896 bytes (jumbo frames).\nNote that packets larger than 1500 bytes (standard Ethernet) can be subject to TCP-MSS clamping or dropped\nwith an ICMP 'Fragmentation-Needed' message if the packets are routed to the Internet or other VPCs\nwith varying MTUs.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Name of the resource. Provided by the client when the resource is\ncreated. The name must be 1-63 characters long, and comply with\nRFC1035. Specifically, the name must be 1-63 characters long and match\nthe regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the\nfirst character must be a lowercase letter, and all following\ncharacters must be a dash, lowercase letter, or digit, except the last\ncharacter, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_firewall_policy_enforcement_order": {
        "description": "Set the order that Firewall Rules and Firewall Policies are evaluated. Default value: \"AFTER_CLASSIC_FIREWALL\" Possible values: [\"BEFORE_CLASSIC_FIREWALL\", \"AFTER_CLASSIC_FIREWALL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_id": {
        "computed": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "numeric_id": {
        "computed": true,
        "deprecated": true,
        "description": "The unique identifier for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "routing_mode": {
        "computed": true,
        "description": "The network-wide routing mode to use. If set to 'REGIONAL', this\nnetwork's cloud routers will only advertise routes with subnetworks\nof this network in the same region as the router. If set to 'GLOBAL',\nthis network's cloud routers will advertise routes with all\nsubnetworks of this network, across regions. Possible values: [\"REGIONAL\", \"GLOBAL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
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

func GoogleComputeNetworkSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeNetwork), &result)
	return &result
}

package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeAddress = `{
  "block": {
    "attributes": {
      "address": {
        "computed": true,
        "description": "The static external IP address represented by this resource.\nThe IP address must be inside the specified subnetwork,\nif any. Set by the API if undefined.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "address_type": {
        "description": "The type of address to reserve.\nNote: if you set this argument's value as 'INTERNAL' you need to leave the 'network_tier' argument unset in that resource block. Default value: \"EXTERNAL\" Possible values: [\"INTERNAL\", \"EXTERNAL\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
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
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_version": {
        "description": "The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: [\"IPV4\", \"IPV6\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ipv6_endpoint_type": {
        "description": "The endpoint type of this address, which should be VM or NETLB. This is\nused for deciding which type of endpoint this address can be used after\nthe external IPv6 address reservation. Possible values: [\"VM\", \"NETLB\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "Name of the resource. The name must be 1-63 characters long, and\ncomply with RFC1035. Specifically, the name must be 1-63 characters\nlong and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'\nwhich means the first character must be a lowercase letter, and all\nfollowing characters must be a dash, lowercase letter, or digit,\nexcept the last character, which cannot be a dash.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network": {
        "description": "The URL of the network in which to reserve the address. This field\ncan only be used with INTERNAL type with the VPC_PEERING and\nIPSEC_INTERCONNECT purposes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_tier": {
        "computed": true,
        "description": "The networking tier used for configuring this address. If this field is not\nspecified, it is assumed to be PREMIUM.\nThis argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: [\"PREMIUM\", \"STANDARD\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "prefix_length": {
        "computed": true,
        "description": "The prefix length if the resource represents an IP range.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "purpose": {
        "computed": true,
        "description": "The purpose of this resource, which can be one of the following values.\n\n* GCE_ENDPOINT for addresses that are used by VM instances, alias IP\nranges, load balancers, and similar resources.\n\n* SHARED_LOADBALANCER_VIP for an address that can be used by multiple\ninternal load balancers.\n\n* VPC_PEERING for addresses that are reserved for VPC peer networks.\n\n* IPSEC_INTERCONNECT for addresses created from a private IP range that\nare reserved for a VLAN attachment in an HA VPN over Cloud Interconnect\nconfiguration. These addresses are regional resources.\n\n* PRIVATE_SERVICE_CONNECT for a private network address that is used to\nconfigure Private Service Connect. Only global internal addresses can use\nthis purpose.\n\n\nThis should only be set when using an Internal address.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "The Region in which the created address should reside.\nIf it is not provided, the provider region is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "self_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "subnetwork": {
        "computed": true,
        "description": "The URL of the subnetwork in which to reserve the address. If an IP\naddress is specified, it must be within the subnetwork's IP range.\nThis field can only be used with INTERNAL type with\nGCE_ENDPOINT/DNS_RESOLVER purposes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "users": {
        "computed": true,
        "description": "The URLs of the resources that are using this address.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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

func GoogleComputeAddressSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeAddress), &result)
	return &result
}

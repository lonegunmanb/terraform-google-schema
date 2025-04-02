package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputePublicDelegatedPrefix = `{
  "block": {
    "attributes": {
      "allocatable_prefix_length": {
        "computed": true,
        "description": "The allocatable prefix length supported by this public delegated prefix. This field is optional and cannot be set for prefixes in DELEGATION mode. It cannot be set for IPv4 prefixes either, and it always defaults to 32.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "ip_cidr_range": {
        "description": "The IP address range, in CIDR format, represented by this public delegated prefix.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_live_migration": {
        "description": "If true, the prefix will be live migrated.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "mode": {
        "description": "Specifies the mode of this IPv6 PDP. MODE must be one of: DELEGATION,\nEXTERNAL_IPV6_FORWARDING_RULE_CREATION and EXTERNAL_IPV6_SUBNETWORK_CREATION. Possible values: [\"DELEGATION\", \"EXTERNAL_IPV6_FORWARDING_RULE_CREATION\", \"EXTERNAL_IPV6_SUBNETWORK_CREATION\"]",
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
      "parent_prefix": {
        "description": "The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.",
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
      "region": {
        "description": "A region where the prefix will reside.",
        "description_kind": "plain",
        "required": true,
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

func GoogleComputePublicDelegatedPrefixSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputePublicDelegatedPrefix), &result)
	return &result
}

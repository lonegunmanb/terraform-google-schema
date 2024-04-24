package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleVmwareengineExternalAccessRule = `{
  "block": {
    "attributes": {
      "action": {
        "description": "The action that the external access rule performs. Possible values: [\"ALLOW\", \"DENY\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "create_time": {
        "computed": true,
        "description": "Creation time of this resource.\nA timestamp in RFC3339 UTC \"Zulu\" format, with nanosecond resolution and\nup to nine fractional digits. Examples: \"2014-10-02T15:01:23Z\" and \"2014-10-02T15:01:23.045123456Z\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "description": "User-provided description for the external access rule.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_ports": {
        "description": "A list of destination ports to which the external access rule applies.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ip_protocol": {
        "description": "The IP protocol to which the external access rule applies.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The ID of the external access rule.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parent": {
        "description": "The resource name of the network policy.\nResource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.\nFor example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "description": "External access rule priority, which determines the external access rule to use when multiple rules apply.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "source_ports": {
        "description": "A list of source ports to which the external access rule applies.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "state": {
        "computed": true,
        "description": "State of the Cluster.",
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
      }
    },
    "block_types": {
      "destination_ip_ranges": {
        "block": {
          "attributes": {
            "external_address": {
              "description": "The name of an 'ExternalAddress' resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address_range": {
              "description": "An IP address range in the CIDR format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "If destination ranges are specified, the external access rule applies only to\ntraffic that has a destination IP address in these ranges.",
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "source_ip_ranges": {
        "block": {
          "attributes": {
            "ip_address": {
              "description": "A single IP address.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ip_address_range": {
              "description": "An IP address range in the CIDR format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "If source ranges are specified, the external access rule applies only to\ntraffic that has a source IP address in these ranges.",
          "description_kind": "plain"
        },
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

func GoogleVmwareengineExternalAccessRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleVmwareengineExternalAccessRule), &result)
	return &result
}

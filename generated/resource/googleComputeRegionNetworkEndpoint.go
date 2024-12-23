package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeRegionNetworkEndpoint = `{
  "block": {
    "attributes": {
      "fqdn": {
        "description": "Fully qualified domain name of network endpoint.\n\nThis can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.",
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
      "ip_address": {
        "description": "IPv4 address external endpoint.\n\nThis can only be specified when network_endpoint_type of the NEG is INTERNET_IP_PORT.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network_endpoint_id": {
        "computed": true,
        "description": "The unique identifier number for the resource. This identifier is defined by the server.",
        "description_kind": "plain",
        "type": "number"
      },
      "port": {
        "description": "Port number of network endpoint.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region": {
        "computed": true,
        "description": "Region where the containing network endpoint group is located.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "region_network_endpoint_group": {
        "description": "The network endpoint group this endpoint is part of.",
        "description_kind": "plain",
        "required": true,
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

func GoogleComputeRegionNetworkEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeRegionNetworkEndpoint), &result)
	return &result
}

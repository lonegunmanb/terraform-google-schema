package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceNetworkingConnection = `{
  "block": {
    "attributes": {
      "deletion_policy": {
        "description": "When set to ABANDON, terraform will abandon management of the resource instead of deleting it. Prevents terraform apply failures with CloudSQL. Note: The resource will still exist.",
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
      "network": {
        "description": "Name of VPC network connected with service producers using VPC peering.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peering": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "reserved_peering_ranges": {
        "description": "Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a different range when connection is already established will not reallocate already provisioned service producer subnetworks.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "service": {
        "description": "Provider peering service that is managing peering connectivity for a service provider organization. For Google services that support this functionality it is 'servicenetworking.googleapis.com'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_on_creation_fail": {
        "description": "When set to true, enforce an update of the reserved peering ranges on the existing service networking connection in case of a new connection creation failure.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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

func GoogleServiceNetworkingConnectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceNetworkingConnection), &result)
	return &result
}

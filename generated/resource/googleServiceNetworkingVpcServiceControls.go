package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleServiceNetworkingVpcServiceControls = `{
  "block": {
    "attributes": {
      "enabled": {
        "description": "Desired VPC Service Controls state service producer VPC network, as\ndescribed at the top of this page.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "network": {
        "description": "The network that the consumer is using to connect with services.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description": "The id of the Google Cloud project containing the consumer network.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service": {
        "description": "The service that is managing peering connectivity for a service\nproducer's organization. For Google services that support this\nfunctionality, this value is 'servicenetworking.googleapis.com'.",
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

func GoogleServiceNetworkingVpcServiceControlsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleServiceNetworkingVpcServiceControls), &result)
	return &result
}

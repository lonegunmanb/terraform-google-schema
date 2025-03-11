package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleComputeFirewallPolicyAssociation = `{
  "block": {
    "attributes": {
      "attachment_target": {
        "description": "The target that the firewall policy is attached to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "firewall_policy": {
        "description": "The firewall policy of the resource.\n\nThis field can be updated to refer to a different Firewall Policy, which will create a new association from that new\nfirewall policy with the flag to override the existing attachmentTarget's policy association.\n\n**Note** Due to potential risks with this operation it is *highly* recommended to use the 'create_before_destroy' life cycle option\non your exisiting firewall policy so as to prevent a situation where your attachment target has no associated policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name for an association.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "short_name": {
        "computed": true,
        "description": "The short name of the firewall policy of the association.",
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

func GoogleComputeFirewallPolicyAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleComputeFirewallPolicyAssociation), &result)
	return &result
}

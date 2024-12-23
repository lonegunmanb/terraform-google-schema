package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleCloudIdentityGroupTransitiveMemberships = `{
  "block": {
    "attributes": {
      "group": {
        "description": "The name of the Group to get memberships from.",
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
      "memberships": {
        "computed": true,
        "description": "List of Cloud Identity group memberships.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "member": "string",
              "preferred_member_key": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "namespace": "string"
                  }
                ]
              ],
              "relation_type": "string",
              "roles": [
                "set",
                [
                  "object",
                  {
                    "role": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleCloudIdentityGroupTransitiveMembershipsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleCloudIdentityGroupTransitiveMemberships), &result)
	return &result
}

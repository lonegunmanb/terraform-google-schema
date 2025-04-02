package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleOrganizationIamCustomRoles = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "org_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "roles": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "deleted": "bool",
              "description": "string",
              "id": "string",
              "name": "string",
              "permissions": [
                "list",
                "string"
              ],
              "role_id": "string",
              "stage": "string",
              "title": "string"
            }
          ]
        ]
      },
      "show_deleted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "view": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleOrganizationIamCustomRolesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleOrganizationIamCustomRoles), &result)
	return &result
}

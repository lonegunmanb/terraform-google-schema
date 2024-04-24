package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleAlloydbUser = `{
  "block": {
    "attributes": {
      "cluster": {
        "description": "Identifies the alloydb cluster. Must be in the format\n'projects/{project}/locations/{location}/clusters/{cluster_id}'",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_roles": {
        "description": "List of database roles this database user has.",
        "description_kind": "plain",
        "optional": true,
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
      "name": {
        "computed": true,
        "description": "Name of the resource in the form of projects/{project}/locations/{location}/clusters/{cluster}/users/{user}.",
        "description_kind": "plain",
        "type": "string"
      },
      "password": {
        "description": "Password for this database user.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "user_id": {
        "description": "The database role name of the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_type": {
        "description": "The type of this user. Possible values: [\"ALLOYDB_BUILT_IN\", \"ALLOYDB_IAM_USER\"]",
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

func GoogleAlloydbUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleAlloydbUser), &result)
	return &result
}

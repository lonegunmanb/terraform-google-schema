package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleBigtableAuthorizedView = `{
  "block": {
    "attributes": {
      "deletion_protection": {
        "computed": true,
        "description": "A field to make the authorized view protected against data loss i.e. when set to PROTECTED, deleting the authorized view, the table containing the authorized view, and the instance containing the authorized view would be prohibited.\nIf not provided, currently deletion protection will be set to UNPROTECTED as it is the API default value. Note this field configs the deletion protection provided by the API in the backend, and should not be confused with Terraform-side deletion protection.",
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
      "instance_name": {
        "description": "The name of the Bigtable instance in which the authorized view belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the authorized view. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description": "The ID of the project in which the resource belongs. If it is not provided, the provider project is used.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "table_name": {
        "description": "The name of the Bigtable table in which the authorized view belongs.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "subset_view": {
        "block": {
          "attributes": {
            "row_prefixes": {
              "description": "Base64-encoded row prefixes to be included in the authorized view. To provide access to all rows, include the empty string as a prefix (\"\").",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "family_subsets": {
              "block": {
                "attributes": {
                  "family_name": {
                    "description": "Name of the column family to be included in the authorized view.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "qualifier_prefixes": {
                    "description": "Base64-encoded prefixes for qualifiers of the column family to be included in the authorized view. Every qualifier starting with one of these prefixes is included in the authorized view. To provide access to all qualifiers, include the empty string as a prefix (\"\").",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "qualifiers": {
                    "description": "Base64-encoded individual exact column qualifiers of the column family to be included in the authorized view.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description": "Subsets of column families to be included in the authorized view.",
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            }
          },
          "description": "An AuthorizedView permitting access to an explicit subset of a Table.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleBigtableAuthorizedViewSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleBigtableAuthorizedView), &result)
	return &result
}

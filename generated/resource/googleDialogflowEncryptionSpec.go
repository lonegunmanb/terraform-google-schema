package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleDialogflowEncryptionSpec = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description": "The location in which the encryptionSpec is to be initialized.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "encryption_spec": {
        "block": {
          "attributes": {
            "kms_key": {
              "description": "The name of customer-managed encryption key that is used to secure a resource and its sub-resources.\nIf empty, the resource is secured by the default Google encryption key.\nOnly the key in the same location as this resource is allowed to be used for encryption.\nFormat: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{key}",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description": "A nested object resource.",
          "description_kind": "plain"
        },
        "max_items": 1,
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

func GoogleDialogflowEncryptionSpecSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleDialogflowEncryptionSpec), &result)
	return &result
}

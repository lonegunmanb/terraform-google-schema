package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerRegionalSecretVersion = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the regional secret version was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_encryption": {
        "computed": true,
        "description": "The customer-managed encryption configuration of the regional secret.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "kms_key_version_name": "string"
            }
          ]
        ]
      },
      "deletion_policy": {
        "description": "The deletion policy for the regional secret version. Setting 'ABANDON' allows the resource\nto be abandoned rather than deleted. Setting 'DISABLE' allows the resource to be\ndisabled rather than deleted. Default is 'DELETE'. Possible values are:\n  * DELETE\n  * DISABLE\n  * ABANDON",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destroy_time": {
        "computed": true,
        "description": "The time at which the regional secret version was destroyed. Only present if state is DESTROYED.",
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description": "The current state of the regional secret version.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "is_secret_data_base64": {
        "description": "If set to 'true', the secret data is expected to be base64-encoded string and would be sent as is.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description": "Location of Secret Manager regional secret resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the regional secret version. Format:\n'projects/{{project}}/locations/{{location}}/secrets/{{secret_id}}/versions/{{version}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "secret": {
        "description": "Secret Manager regional secret resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secret_data": {
        "description": "The secret data. Must be no larger than 64KiB.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "version": {
        "computed": true,
        "description": "The version of the Regional Secret.",
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

func GoogleSecretManagerRegionalSecretVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerRegionalSecretVersion), &result)
	return &result
}

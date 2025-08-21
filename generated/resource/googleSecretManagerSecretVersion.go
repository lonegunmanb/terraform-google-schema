package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleSecretManagerSecretVersion = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the Secret was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "deletion_policy": {
        "description": "The deletion policy for the secret version. Setting 'ABANDON' allows the resource\nto be abandoned rather than deleted. Setting 'DISABLE' allows the resource to be\ndisabled rather than deleted. Default is 'DELETE'. Possible values are:\n  * DELETE\n  * DISABLE\n  * ABANDON",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destroy_time": {
        "computed": true,
        "description": "The time at which the Secret was destroyed. Only present if state is DESTROYED.",
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description": "The current state of the SecretVersion.",
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
      "name": {
        "computed": true,
        "description": "The resource name of the SecretVersion. Format:\n'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "secret": {
        "description": "Secret Manager secret resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secret_data": {
        "description": "The secret data. Must be no larger than 64KiB.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "secret_data_wo": {
        "description": "The secret data. Must be no larger than 64KiB. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes)",
        "description_kind": "plain",
        "optional": true,
        "type": "string",
        "write_only": true
      },
      "secret_data_wo_version": {
        "description": "Triggers update of secret data write-only. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes)",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "version": {
        "computed": true,
        "description": "The version of the Secret.",
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

func GoogleSecretManagerSecretVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleSecretManagerSecretVersion), &result)
	return &result
}

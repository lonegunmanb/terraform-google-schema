package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleIamWorkforcePoolProviderKey = `{
  "block": {
    "attributes": {
      "expire_time": {
        "computed": true,
        "description": "The time after which the key will be permanently deleted and cannot be recovered.\nNote that the key may get purged before this time if the total limit of keys per provider is exceeded.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_id": {
        "description": "The ID to use for the key, which becomes the final component of the resource name. This value must be 4-32 characters, and may contain the characters [a-z0-9-].",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description": "The location for the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Identifier. The resource name of the key.\nFormat: 'locations/{location}/workforcePools/{workforcePoolId}/providers/{providerId}/keys/{keyId}'",
        "description_kind": "plain",
        "type": "string"
      },
      "provider_id": {
        "description": "The ID of the provider.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The state of the key.",
        "description_kind": "plain",
        "type": "string"
      },
      "use": {
        "description": "The purpose of the key. Possible values: [\"ENCRYPTION\"]",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "workforce_pool_id": {
        "description": "The ID of the workforce pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "key_data": {
        "block": {
          "attributes": {
            "format": {
              "computed": true,
              "description": "The format of the key.",
              "description_kind": "plain",
              "type": "string"
            },
            "key": {
              "computed": true,
              "description": "The key data. The format of the key is represented by the format field.",
              "description_kind": "plain",
              "type": "string"
            },
            "key_spec": {
              "description": "The specifications for the key. Possible values: [\"RSA_2048\", \"RSA_3072\", \"RSA_4096\"]",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "not_after_time": {
              "computed": true,
              "description": "Latest timestamp when this key is valid. Attempts to use this key after this time will fail.\nOnly present if the key data represents a X.509 certificate.\n\nUses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits.\nOffsets other than \"Z\" are also accepted.\nExamples: \"2014-10-02T15:01:23Z\", \"2014-10-02T15:01:23.045123456Z\" or \"2014-10-02T15:01:23+05:30\".",
              "description_kind": "plain",
              "type": "string"
            },
            "not_before_time": {
              "computed": true,
              "description": "Earliest timestamp when this key is valid. Attempts to use this key before this time will fail.\nOnly present if the key data represents a X.509 certificate.\n\nUses RFC 3339, where generated output will always be Z-normalized and uses 0, 3, 6 or 9 fractional digits.\nOffsets other than \"Z\" are also accepted.\nExamples: \"2014-10-02T15:01:23Z\", \"2014-10-02T15:01:23.045123456Z\" or \"2014-10-02T15:01:23+05:30\".",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description": "Immutable. Public half of the asymmetric key.",
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

func GoogleIamWorkforcePoolProviderKeySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleIamWorkforcePoolProviderKey), &result)
	return &result
}

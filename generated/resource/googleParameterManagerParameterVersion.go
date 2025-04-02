package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleParameterManagerParameterVersion = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the Parameter Version was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "description": "The current state of Parameter Version. This field is only applicable for updating Parameter Version.",
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
      "kms_key_version": {
        "computed": true,
        "description": "The resource name of the Cloud KMS CryptoKeyVersion used to decrypt parameter version payload. Format\n'projects/{{project}}/locations/global/keyRings/{{key_ring}}/cryptoKeys/{{crypto_key}}/cryptoKeyVersions/{{crypto_key_version}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Parameter Version. Format:\n'projects/{{project}}/locations/global/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "parameter": {
        "description": "Parameter Manager Parameter resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameter_data": {
        "description": "The Parameter data.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "parameter_version_id": {
        "description": "Version ID of the Parameter Version Resource. This must be unique within the Parameter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which the Parameter Version was updated.",
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

func GoogleParameterManagerParameterVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleParameterManagerParameterVersion), &result)
	return &result
}

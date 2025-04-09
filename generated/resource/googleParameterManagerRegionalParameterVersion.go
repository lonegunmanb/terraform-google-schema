package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleParameterManagerRegionalParameterVersion = `{
  "block": {
    "attributes": {
      "create_time": {
        "computed": true,
        "description": "The time at which the Regional Parameter Version was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "disabled": {
        "description": "The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.",
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
        "description": "The resource name of the Cloud KMS CryptoKeyVersion used to decrypt regional parameter version payload. Format\n'projects/{{project}}/locations/{{location}}/keyRings/{{key_ring}}/cryptoKeys/{{crypto_key}}/cryptoKeyVersions/{{crypto_key_version}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description": "Location of Parameter Manager Regional parameter resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name of the Regional Parameter Version. Format:\n'projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'",
        "description_kind": "plain",
        "type": "string"
      },
      "parameter": {
        "description": "Parameter Manager Regional Parameter resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameter_data": {
        "description": "The Regional Parameter data.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "parameter_version_id": {
        "description": "Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "update_time": {
        "computed": true,
        "description": "The time at which the Regional Parameter Version was updated.",
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

func GoogleParameterManagerRegionalParameterVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleParameterManagerRegionalParameterVersion), &result)
	return &result
}

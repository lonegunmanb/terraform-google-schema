package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsCryptoKeyVersion = `{
  "block": {
    "attributes": {
      "algorithm": {
        "computed": true,
        "description": "The CryptoKeyVersionAlgorithm that this CryptoKeyVersion supports.",
        "description_kind": "plain",
        "type": "string"
      },
      "attestation": {
        "computed": true,
        "description": "Statement that was generated and signed by the HSM at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google.\nOnly provided for key versions with protectionLevel HSM.",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cert_chains": [
                "list",
                [
                  "object",
                  {
                    "cavium_certs": [
                      "list",
                      "string"
                    ],
                    "google_card_certs": [
                      "list",
                      "string"
                    ],
                    "google_partition_certs": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "content": "string",
              "external_protection_level_options": [
                "list",
                [
                  "object",
                  {
                    "ekm_connection_key_path": "string",
                    "external_key_uri": "string"
                  }
                ]
              ],
              "format": "string"
            }
          ]
        ]
      },
      "crypto_key": {
        "description": "The name of the cryptoKey associated with the CryptoKeyVersions.\nFormat: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyring}}/cryptoKeys/{{cryptoKey}}''",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "generate_time": {
        "computed": true,
        "description": "The time this CryptoKeyVersion key material was generated",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The resource name for this CryptoKeyVersion.",
        "description_kind": "plain",
        "type": "string"
      },
      "protection_level": {
        "computed": true,
        "description": "The ProtectionLevel describing how crypto operations are performed with this CryptoKeyVersion.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the CryptoKeyVersion. Note: you can only specify this field to manually 'ENABLE' or 'DISABLE' the CryptoKeyVersion,\notherwise the value of this field is always retrieved automatically. Possible values: [\"PENDING_GENERATION\", \"ENABLED\", \"DISABLED\", \"DESTROYED\", \"DESTROY_SCHEDULED\", \"PENDING_IMPORT\", \"IMPORT_FAILED\"]",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "external_protection_level_options": {
        "block": {
          "attributes": {
            "ekm_connection_key_path": {
              "description": "The path to the external key material on the EKM when using EkmConnection e.g., \"v0/my/key\". Set this field instead of externalKeyUri when using an EkmConnection.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "external_key_uri": {
              "description": "The URI for an external resource that this CryptoKeyVersion represents.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description": "ExternalProtectionLevelOptions stores a group of additional fields for configuring a CryptoKeyVersion that are specific to the EXTERNAL protection level and EXTERNAL_VPC protection levels.",
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

func GoogleKmsCryptoKeyVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsCryptoKeyVersion), &result)
	return &result
}

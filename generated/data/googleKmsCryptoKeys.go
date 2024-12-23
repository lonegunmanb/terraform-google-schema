package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsCryptoKeys = `{
  "block": {
    "attributes": {
      "filter": {
        "description": "\n\t\t\t\t\tThe filter argument is used to add a filter query parameter that limits which keys are retrieved by the data source: ?filter={{filter}}.\n\t\t\t\t\tExample values:\n\t\t\t\t\t\n\t\t\t\t\t* \"name:my-key-\" will retrieve keys that contain \"my-key-\" anywhere in their name. Note: names take the form projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}.\n\t\t\t\t\t* \"name=projects/my-project/locations/global/keyRings/my-key-ring/cryptoKeys/my-key-1\" will only retrieve a key with that exact name.\n\t\t\t\t\t\n\t\t\t\t\t[See the documentation about using filters](https://cloud.google.com/kms/docs/sorting-and-filtering)\n\t\t\t\t",
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
      "key_ring": {
        "description": "The key ring that the keys belongs to. Format: 'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "keys": {
        "computed": true,
        "description": "A list of all the retrieved keys from the provided key ring",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "crypto_key_backend": "string",
              "destroy_scheduled_duration": "string",
              "effective_labels": [
                "map",
                "string"
              ],
              "id": "string",
              "import_only": "bool",
              "key_ring": "string",
              "labels": [
                "map",
                "string"
              ],
              "name": "string",
              "primary": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "state": "string"
                  }
                ]
              ],
              "purpose": "string",
              "rotation_period": "string",
              "skip_initial_version_creation": "bool",
              "terraform_labels": [
                "map",
                "string"
              ],
              "version_template": [
                "list",
                [
                  "object",
                  {
                    "algorithm": "string",
                    "protection_level": "string"
                  }
                ]
              ]
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsCryptoKeysSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsCryptoKeys), &result)
	return &result
}

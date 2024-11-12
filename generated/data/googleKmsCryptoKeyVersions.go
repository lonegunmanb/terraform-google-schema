package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsCryptoKeyVersions = `{
  "block": {
    "attributes": {
      "crypto_key": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filter": {
        "description": "\n\t\t\t\t\tThe filter argument is used to add a filter query parameter that limits which cryptoKeyVersions are retrieved by the data source: ?filter={{filter}}.\n\t\t\t\t\tExample values:\n\t\t\t\t\t\n\t\t\t\t\t* \"name:my-cryptokey-version-\" will retrieve cryptoKeyVersions that contain \"my-key-\" anywhere in their name. Note: names take the form projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}/cryptoKeys/{{cryptoKey}}/cryptoKeyVersions/{{cryptoKeyVersion}}.\n\t\t\t\t\t* \"name=projects/my-project/locations/global/keyRings/my-key-ring/cryptoKeys/my-key-1/cryptoKeyVersions/1\" will only retrieve a key with that exact name.\n\t\t\t\t\t\n\t\t\t\t\t[See the documentation about using filters](https://cloud.google.com/kms/docs/sorting-and-filtering)\n\t\t\t\t",
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
      "public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "pem": "string"
            }
          ]
        ]
      },
      "versions": {
        "computed": true,
        "description": "A list of all the retrieved cryptoKeyVersions from the provided crypto key",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "algorithm": "string",
              "crypto_key": "string",
              "id": "string",
              "name": "string",
              "protection_level": "string",
              "public_key": [
                "list",
                [
                  "object",
                  {
                    "algorithm": "string",
                    "pem": "string"
                  }
                ]
              ],
              "state": "string",
              "version": "number"
            }
          ]
        ]
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsCryptoKeyVersionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsCryptoKeyVersions), &result)
	return &result
}

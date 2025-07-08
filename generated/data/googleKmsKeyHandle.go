package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const googleKmsKeyHandle = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kms_key": {
        "computed": true,
        "description": "A reference to a Cloud KMS CryptoKey that can be used for CMEK in the requested\nproduct/project/location, for example\n'projects/1/locations/us-east1/keyRings/foo/cryptoKeys/bar-ffffff'",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description": "The location for the KeyHandle.\nA full list of valid locations can be found by running 'gcloud kms locations list'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The resource name for the KeyHandle.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type_selector": {
        "computed": true,
        "description": "Selector of the resource type where we want to protect resources.\nFor example, 'storage.googleapis.com/Bucket'.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func GoogleKmsKeyHandleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(googleKmsKeyHandle), &result)
	return &result
}
